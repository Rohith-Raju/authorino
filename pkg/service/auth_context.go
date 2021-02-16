package service

import (
	"fmt"
	"strings"
	"sync"

	"golang.org/x/net/context"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/3scale-labs/authorino/pkg/config"
	"github.com/3scale-labs/authorino/pkg/config/common"

	envoy_auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
)

var (
	authCtxLog = ctrl.Log.WithName("Authorino").WithName("AuthContext")
)

type EvaluationResponse struct {
	Evaluator common.AuthConfigEvaluator
	Success bool
	Object interface{}
	Error error
}

func newEvaluationResponseSuccess(evaluator common.AuthConfigEvaluator, obj interface{}) EvaluationResponse {
	return EvaluationResponse{
		Evaluator: evaluator,
		Success: true,
		Object: obj,
	}
}

func newEvaluationResponseFailure(evaluator common.AuthConfigEvaluator, err error) EvaluationResponse {
	return EvaluationResponse{
		Evaluator: evaluator,
		Success: false,
		Error: err,
	}
}

// AuthContext holds the context of each auth request, including the request itself (sent by the client),
// the auth config of the requested API and the lists of identity verifications, metadata add-ons and
// authorization policies, and their corresponding results after evaluated
type AuthContext struct {
	ParentContext *context.Context
	Request       *envoy_auth.CheckRequest
	API           *config.APIConfig

	Identity      map[*config.IdentityConfig]interface{}
	Metadata      map[*config.MetadataConfig]interface{}
	Authorization map[*config.AuthorizationConfig]interface{}
}

type evaluateCallback = func(config common.AuthConfigEvaluator, obj interface{})

// NewAuthContext creates an AuthContext instance
func NewAuthContext(parentCtx context.Context, req *envoy_auth.CheckRequest, apiConfig config.APIConfig) AuthContext {

	return AuthContext{
		ParentContext: &parentCtx,
		Request:       req,
		API:           &apiConfig,
		Identity:      make(map[*config.IdentityConfig]interface{}),
		Metadata:      make(map[*config.MetadataConfig]interface{}),
		Authorization: make(map[*config.AuthorizationConfig]interface{}),
	}

}

func (authContext *AuthContext) evaluateAuthConfig(ctx context.Context, config common.AuthConfigEvaluator, cb evaluateCallback) error {
	select {
	case <-ctx.Done():
		authCtxLog.Info("Context aborted", "config", config)
		return nil
	default:
		if authObj, err := config.Call(authContext); err != nil {
			authCtxLog.Info("Failed to evaluate auth object", "config", config, "error", err)
			return err
		} else {
			cb(config, authObj)
			return nil
		}
	}
}

func (authContext *AuthContext) evaluateOneAuthConfig(authConfigs []common.AuthConfigEvaluator, respChannel *chan EvaluationResponse) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(authConfigs))

	successCallback := func(conf common.AuthConfigEvaluator, authObj interface{}) {
		*respChannel <-newEvaluationResponseSuccess(conf, authObj)
	}

	failureCallback := func(conf common.AuthConfigEvaluator, err error) {
		*respChannel <-newEvaluationResponseFailure(conf, err)
	}

	for _, authConfig := range authConfigs {
		objConfig := authConfig
		go func() {
			defer waitGroup.Done()

			if err := authContext.evaluateAuthConfig(ctx, objConfig, successCallback); err != nil {
				failureCallback(objConfig, err)
			} else {
				cancel() // cancels the context if at least one thread succeeds
			}
		}()
	}

	waitGroup.Wait()
}

func (authContext *AuthContext) evaluateAllAuthConfigs(authConfigs []common.AuthConfigEvaluator, respChannel *chan EvaluationResponse) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(authConfigs))

	successCallback := func(conf common.AuthConfigEvaluator, authObj interface{}) {
		*respChannel <-newEvaluationResponseSuccess(conf, authObj)
	}

	failureCallback := func(conf common.AuthConfigEvaluator, err error) {
		*respChannel <-newEvaluationResponseFailure(conf, err)
	}

	for _, authConfig := range authConfigs {
		objConfig := authConfig
		go func() {
			defer waitGroup.Done()

			if err := authContext.evaluateAuthConfig(ctx, objConfig, successCallback); err != nil {
				failureCallback(objConfig, err)
				cancel() // cancels the context if at least one thread fails
			}
		}()
	}

	waitGroup.Wait()
}

func (authContext *AuthContext) evaluateAnyAuthConfig(authConfigs []common.AuthConfigEvaluator, respChannel *chan EvaluationResponse) {
	ctx := context.Background()
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(authConfigs))

	successCallback := func(conf common.AuthConfigEvaluator, authObj interface{}) {
		*respChannel <-newEvaluationResponseSuccess(conf, authObj)
	}

	failureCallback := func(conf common.AuthConfigEvaluator, err error) {
		*respChannel <-newEvaluationResponseFailure(conf, err)
	}

	for _, authConfig := range authConfigs {
		objConfig := authConfig
		go func() {
			defer waitGroup.Done()

			if err := authContext.evaluateAuthConfig(ctx, objConfig, successCallback); err != nil {
				failureCallback(objConfig, err)
			}
		}()
	}

	waitGroup.Wait()
}

func (authContext *AuthContext) evaluateIdentityConfigs() error {
	configs := authContext.API.IdentityConfigs
	respChannel := make(chan EvaluationResponse, len(configs))

	go func(){
		defer close(respChannel)
		authContext.evaluateOneAuthConfig(configs, &respChannel)
	}()

	var lastError error

	for resp := range respChannel {
		conf, _ := resp.Evaluator.(*config.IdentityConfig)
		obj := resp.Object

		if resp.Success {
			authContext.Identity[conf] = obj
			authCtxLog.Info("Identity", "config", conf, "authObj", obj)
			return nil
		} else {
			lastError = resp.Error
			authCtxLog.Info("Identity", "config", conf, "error", lastError)
		}
	}

	return lastError
}

func (authContext *AuthContext) evaluateMetadataConfigs() {
	configs := authContext.API.MetadataConfigs
	respChannel := make(chan EvaluationResponse, len(configs))

	go func(){
		defer close(respChannel)
		authContext.evaluateAnyAuthConfig(configs, &respChannel)
	}()

	for resp := range respChannel {
		conf, _ := resp.Evaluator.(*config.MetadataConfig)
		obj := resp.Object

		if resp.Success {
			authContext.Metadata[conf] = obj
			authCtxLog.Info("Metadata", "config", conf, "authObj", obj)
		} else {
			authCtxLog.Info("Metadata", "config", conf, "error", resp.Error)
		}
	}
}

func (authContext *AuthContext) evaluateAuthorizationConfigs() error {
	configs := authContext.API.AuthorizationConfigs
	respChannel := make(chan EvaluationResponse, len(configs))

	go func(){
		defer close(respChannel)
		authContext.evaluateAllAuthConfigs(configs, &respChannel)
	}()

	for resp := range respChannel {
		conf, _ := resp.Evaluator.(*config.AuthorizationConfig)
		obj := resp.Object

		if resp.Success {
			authContext.Authorization[conf] = obj
			authCtxLog.Info("Authorization", "config", conf, "authObj", obj)
		} else {
			err := resp.Error
			authCtxLog.Info("Authorization", "config", conf, "error", err)
			return err
		}
	}

	return nil
}

// Evaluate evaluates all steps of the auth pipeline (identity → metadata → policy enforcement)
func (authContext *AuthContext) Evaluate() error {
	// identity
	if err := authContext.evaluateIdentityConfigs(); err != nil {
		return err
	}

	// metadata
	authContext.evaluateMetadataConfigs()

	// policy enforcement (authorization)
	if err := authContext.evaluateAuthorizationConfigs(); err != nil {
		return err
	}

	return nil
}

func (authContext *AuthContext) GetParentContext() *context.Context {
	return authContext.ParentContext
}

func (authContext *AuthContext) GetRequest() *envoy_auth.CheckRequest {
	return authContext.Request
}

func (authContext *AuthContext) GetAPI() interface{} {
	return authContext.API
}

func (authContext *AuthContext) GetIdentity() interface{} {
	var id interface{}
	for _, v := range authContext.Identity {
		if v != nil {
			id = v
			break
		}
	}
	return id
}

func (authContext *AuthContext) GetMetadata() map[string]interface{} {
	m := make(map[string]interface{})
	for metadataCfg, metadataObj := range authContext.Metadata {
		t, _ := metadataCfg.GetType()
		m[t] = metadataObj // FIXME: It will override instead of including all the metadata values of the same type
	}
	return m
}

func (authContext *AuthContext) FindIdentityByName(name string) (interface{}, error) {
	for identityConfig := range authContext.Identity {
		if identityConfig.OIDC.Name == name {
			return identityConfig.OIDC, nil
		}
	}
	return nil, fmt.Errorf("cannot find OIDC token")
}

func (authContext *AuthContext) AuthorizationToken() (string, error) {
	authHeader, authHeaderOK := authContext.Request.Attributes.Request.Http.Headers["authorization"]

	var splitToken []string

	if authHeaderOK {
		splitToken = strings.Split(authHeader, "Bearer ")
	}
	if !authHeaderOK || len(splitToken) != 2 {
		return "", fmt.Errorf("authorization header malformed or not provided")
	}

	return splitToken[1], nil // FIXME: Indexing may panic because because of 'nil' slice
}
