package core

import (
	"net/http"
	"strings"

	logmd "github.com/oslokommune/okctl/pkg/middleware/logger"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	kit "github.com/go-kit/kit/transport/http"
	"github.com/oslokommune/okctl/pkg/api"
	"github.com/oslokommune/okctl/pkg/transport"
	"github.com/sirupsen/logrus"
)

// Endpoints defines all available endpoints
type Endpoints struct {
	CreateCluster                                 endpoint.Endpoint
	DeleteCluster                                 endpoint.Endpoint
	CreateVpc                                     endpoint.Endpoint
	DeleteVpc                                     endpoint.Endpoint
	CreateExternalSecretsPolicy                   endpoint.Endpoint
	CreateExternalSecretsServiceAccount           endpoint.Endpoint
	CreateExternalSecretsHelmChart                endpoint.Endpoint
	CreateAlbIngressControllerServiceAccount      endpoint.Endpoint
	CreateAlbIngressControllerPolicy              endpoint.Endpoint
	CreateAlbIngressControllerHelmChart           endpoint.Endpoint
	CreateExternalDNSPolicy                       endpoint.Endpoint
	CreateExternalDNSServiceAccount               endpoint.Endpoint
	CreateExternalDNSKubeDeployment               endpoint.Endpoint
	CreateHostedZone                              endpoint.Endpoint
	CreateCertificate                             endpoint.Endpoint
	CreateSecret                                  endpoint.Endpoint
	DeleteSecret                                  endpoint.Endpoint
	CreateArgoCD                                  endpoint.Endpoint
	CreateExternalSecrets                         endpoint.Endpoint
	DeleteExternalSecretsPolicy                   endpoint.Endpoint
	DeleteAlbIngressControllerPolicy              endpoint.Endpoint
	DeleteExternalDNSPolicy                       endpoint.Endpoint
	DeleteHostedZone                              endpoint.Endpoint
	DeleteExternalSecretsServiceAccount           endpoint.Endpoint
	DeleteAlbIngressControllerServiceAccount      endpoint.Endpoint
	DeleteExternalDNSServiceAccount               endpoint.Endpoint
	CreateIdentityPool                            endpoint.Endpoint
	CreateIdentityPoolClient                      endpoint.Endpoint
	CreateIdentityPoolUser                        endpoint.Endpoint
	DeleteIdentityPool                            endpoint.Endpoint
	DeleteIdentityPoolClient                      endpoint.Endpoint
	CreateAWSLoadBalancerControllerServiceAccount endpoint.Endpoint
	DeleteAWSLoadBalancerControllerServiceAccount endpoint.Endpoint
	CreateAWSLoadBalancerControllerPolicy         endpoint.Endpoint
	DeleteAWSLoadBalancerControllerPolicy         endpoint.Endpoint
	CreateAWSLoadBalancerControllerHelmChart      endpoint.Endpoint
	DeleteCertificate                             endpoint.Endpoint
	DeleteNamespace                               endpoint.Endpoint
	DeleteCognitoCertificate                      endpoint.Endpoint
	CreateAutoscalerHelmChart                     endpoint.Endpoint
	CreateAutoscalerServiceAccount                endpoint.Endpoint
	DeleteAutoscalerServiceAccount                endpoint.Endpoint
	CreateAutoscalerPolicy                        endpoint.Endpoint
	DeleteAutoscalerPolicy                        endpoint.Endpoint
}

// MakeEndpoints returns the endpoints initialised with their
// corresponding service
func MakeEndpoints(s Services) Endpoints {
	return Endpoints{
		CreateCluster:                                 makeCreateClusterEndpoint(s.Cluster),
		DeleteCluster:                                 makeDeleteClusterEndpoint(s.Cluster),
		CreateVpc:                                     makeCreateVpcEndpoint(s.Vpc),
		DeleteVpc:                                     makeDeleteVpcEndpoint(s.Vpc),
		CreateExternalSecretsPolicy:                   makeCreateExternalSecretsPolicyEndpoint(s.ManagedPolicy),
		CreateExternalSecretsServiceAccount:           makeCreateExternalSecretsServiceAccountEndpoint(s.ServiceAccount),
		CreateExternalSecretsHelmChart:                makeCreateExternalSecretsHelmChartEndpoint(s.Helm),
		CreateAlbIngressControllerServiceAccount:      makeCreateAlbIngressControllerServiceAccountEndpoint(s.ServiceAccount),
		CreateAlbIngressControllerPolicy:              makeCreateAlbIngressControllerPolicyEndpoint(s.ManagedPolicy),
		CreateAlbIngressControllerHelmChart:           makeCreateAlbIngressControllerHelmChartEndpoint(s.Helm),
		CreateExternalDNSPolicy:                       makeCreateExternalDNSPolicyEndpoint(s.ManagedPolicy),
		CreateExternalDNSServiceAccount:               makeCreateExternalDNSServiceAccountEndpoint(s.ServiceAccount),
		CreateExternalDNSKubeDeployment:               makeCreateExternalDNSKubeDeploymentEndpoint(s.Kube),
		CreateHostedZone:                              makeCreateHostedZoneEndpoint(s.Domain),
		CreateCertificate:                             makeCreateCertificateEndpoint(s.Certificate),
		CreateSecret:                                  makeCreateSecret(s.Parameter),
		DeleteSecret:                                  makeDeleteSecret(s.Parameter),
		CreateArgoCD:                                  makeCreateArgoCD(s.Helm),
		CreateExternalSecrets:                         makeCreateExternalSecretsEndpoint(s.Kube),
		DeleteExternalSecretsPolicy:                   makeDeleteExternalSecretsPolicyEndpoint(s.ManagedPolicy),
		DeleteAlbIngressControllerPolicy:              makeDeleteAlbIngressControllerPolicyEndpoint(s.ManagedPolicy),
		DeleteExternalDNSPolicy:                       makeDeleteExternalDNSPolicyEndpoint(s.ManagedPolicy),
		DeleteHostedZone:                              makeDeleteHostedZoneEndpoint(s.Domain),
		DeleteExternalSecretsServiceAccount:           makeDeleteExternalSecretsServiceAccountEndpoint(s.ServiceAccount),
		DeleteAlbIngressControllerServiceAccount:      makeDeleteAlbIngressControllerServiceAccountEndpoint(s.ServiceAccount),
		DeleteExternalDNSServiceAccount:               makeDeleteExternalDNSServiceAccountEndpoint(s.ServiceAccount),
		CreateIdentityPool:                            makeCreateIdentityPoolEndpoint(s.IdentityManager),
		CreateIdentityPoolClient:                      makeCreateIdentityPoolClient(s.IdentityManager),
		CreateIdentityPoolUser:                        makeCreateIdentityPoolUser(s.IdentityManager),
		DeleteIdentityPool:                            makeDeleteIdentityPoolEndpoint(s.IdentityManager),
		DeleteIdentityPoolClient:                      makeDeleteIdentityPoolClientEndpoint(s.IdentityManager),
		CreateAWSLoadBalancerControllerServiceAccount: makeCreateAWSLoadBalancerControllerServiceAccountEndpoint(s.ServiceAccount),
		DeleteAWSLoadBalancerControllerServiceAccount: makeDeleteAWSLoadBalancerControllerServiceAccountEndpoint(s.ServiceAccount),
		CreateAWSLoadBalancerControllerPolicy:         makeCreateAWSLoadBalancerControllerPolicyEndpoint(s.ManagedPolicy),
		DeleteAWSLoadBalancerControllerPolicy:         makeDeleteAWSLoadBalancerControllerPolicyEndpoint(s.ManagedPolicy),
		CreateAWSLoadBalancerControllerHelmChart:      makeCreateAWSLoadBalancerControllerHelmChartEndpoint(s.Helm),
		DeleteCertificate:                             makeDeleteCertificateEndpoint(s.Certificate),
		DeleteNamespace:                               makeDeleteNamespaceEndpoint(s.Kube),
		DeleteCognitoCertificate:                      makeDeleteCognitoCertificateEndpoint(s.Certificate),
		CreateAutoscalerHelmChart:                     makeCreateAutoscalerHelmtChartEndpoint(s.Helm),
		CreateAutoscalerServiceAccount:                makeCreateAutoscalerServiceAccountEndpoint(s.ServiceAccount),
		DeleteAutoscalerServiceAccount:                makeDeleteAutoscalerServiceAccountEndpoint(s.ServiceAccount),
		CreateAutoscalerPolicy:                        makeCreateAutoscalerPolicyEndpoint(s.ManagedPolicy),
		DeleteAutoscalerPolicy:                        makeDeleteAutoscalerPolicyEndpoint(s.ManagedPolicy),
	}
}

// Handlers defines http handlers for processing requests
type Handlers struct {
	CreateCluster                                 http.Handler
	DeleteCluster                                 http.Handler
	CreateVpc                                     http.Handler
	DeleteVpc                                     http.Handler
	CreateExternalSecretsPolicy                   http.Handler
	CreateExternalSecretsServiceAccount           http.Handler
	CreateExternalSecretsHelmChart                http.Handler
	CreateAlbIngressControllerServiceAccount      http.Handler
	CreateAlbIngressControllerPolicy              http.Handler
	CreateAlbIngressControllerHelmChart           http.Handler
	CreateExternalDNSPolicy                       http.Handler
	CreateExternalDNSServiceAccount               http.Handler
	CreateExternalDNSKubeDeployment               http.Handler
	CreateHostedZone                              http.Handler
	CreateCertificate                             http.Handler
	CreateSecret                                  http.Handler
	DeleteSecret                                  http.Handler
	CreateArgoCD                                  http.Handler
	CreateExternalSecrets                         http.Handler
	DeleteExternalSecretsPolicy                   http.Handler
	DeleteAlbIngressControllerPolicy              http.Handler
	DeleteExternalDNSPolicy                       http.Handler
	DeleteHostedZone                              http.Handler
	DeleteExternalSecretsServiceAccount           http.Handler
	DeleteAlbIngressControllerServiceAccount      http.Handler
	DeleteExternalDNSServiceAccount               http.Handler
	CreateIdentityPool                            http.Handler
	CreateIdentityPoolClient                      http.Handler
	CreateIdentityPoolUser                        http.Handler
	DeleteIdentityPool                            http.Handler
	DeleteIdentityPoolClient                      http.Handler
	CreateAWSLoadBalancerControllerServiceAccount http.Handler
	DeleteAWSLoadBalancerControllerServiceAccount http.Handler
	CreateAWSLoadBalancerControllerPolicy         http.Handler
	DeleteAWSLoadBalancerControllerPolicy         http.Handler
	CreateAWSLoadBalancerControllerHelmChart      http.Handler
	DeleteCertificate                             http.Handler
	DeleteNamespace                               http.Handler
	DeleteCognitoCertificate                      http.Handler
	CreateAutoscalerHelmChart                     http.Handler
	CreateAutoscalerServiceAccount                http.Handler
	DeleteAutoscalerServiceAccount                http.Handler
	CreateAutoscalerPolicy                        http.Handler
	DeleteAutoscalerPolicy                        http.Handler
}

// EncodeResponseType defines a type for responses
type EncodeResponseType string

const (
	// EncodeJSONResponse encodes as json when returning response
	EncodeJSONResponse EncodeResponseType = "json"
	// EncodeYAMLResponse encodes as yaml when returning response
	EncodeYAMLResponse EncodeResponseType = "yaml"
	// EncodeTextResponse encodes as text when returning response
	EncodeTextResponse EncodeResponseType = "text"
)

// MakeHandlers returns all handlers initialised with encoders, decoders, etc
// nolint: funlen
func MakeHandlers(responseType EncodeResponseType, endpoints Endpoints) *Handlers {
	var encoderFn kit.EncodeResponseFunc

	switch responseType {
	case EncodeJSONResponse:
		encoderFn = kit.EncodeJSONResponse
	case EncodeYAMLResponse:
		encoderFn = transport.EncodeYAMLResponse
	case EncodeTextResponse:
		encoderFn = transport.EncodeTextResponse
	}

	newServer := func(e endpoint.Endpoint, decodeFn kit.DecodeRequestFunc) http.Handler {
		return kit.NewServer(e, decodeFn, encoderFn)
	}

	return &Handlers{
		CreateCluster:                                 newServer(endpoints.CreateCluster, decodeClusterCreateRequest),
		DeleteCluster:                                 newServer(endpoints.DeleteCluster, decodeClusterDeleteRequest),
		CreateVpc:                                     newServer(endpoints.CreateVpc, decodeVpcCreateRequest),
		DeleteVpc:                                     newServer(endpoints.DeleteVpc, decodeVpcDeleteRequest),
		CreateExternalSecretsPolicy:                   newServer(endpoints.CreateExternalSecretsPolicy, decodeCreateExternalSecretsPolicyRequest),
		CreateExternalSecretsServiceAccount:           newServer(endpoints.CreateExternalSecretsServiceAccount, decodeCreateExternalSecretsServiceAccount),
		CreateExternalSecretsHelmChart:                newServer(endpoints.CreateExternalSecretsHelmChart, decodeCreateExternalSecretsHelmChart),
		CreateAlbIngressControllerServiceAccount:      newServer(endpoints.CreateAlbIngressControllerServiceAccount, decodeCreateAlbIngressControllerServiceAccount),
		CreateAlbIngressControllerPolicy:              newServer(endpoints.CreateAlbIngressControllerPolicy, decodeCreateAlbIngressControllerPolicyRequest),
		CreateAlbIngressControllerHelmChart:           newServer(endpoints.CreateAlbIngressControllerHelmChart, decodeCreateAlbIngressControllerHelmChart),
		CreateExternalDNSPolicy:                       newServer(endpoints.CreateExternalDNSPolicy, decodeCreateExternalDNSPolicyRequest),
		CreateExternalDNSServiceAccount:               newServer(endpoints.CreateExternalDNSServiceAccount, decodeCreateExternalDNSServiceAccount),
		CreateExternalDNSKubeDeployment:               newServer(endpoints.CreateExternalDNSKubeDeployment, decodeCreateExternalDNSKubeDeployment),
		CreateHostedZone:                              newServer(endpoints.CreateHostedZone, decodeCreateHostedZone),
		CreateCertificate:                             newServer(endpoints.CreateCertificate, decodeCreateCertificate),
		CreateSecret:                                  newServer(endpoints.CreateSecret, decodeCreateSecret),
		DeleteSecret:                                  newServer(endpoints.DeleteSecret, decodeDeleteSecret),
		CreateArgoCD:                                  newServer(endpoints.CreateArgoCD, decodeCreateArgoCD),
		CreateExternalSecrets:                         newServer(endpoints.CreateExternalSecrets, decodeCreateExternalSecrets),
		DeleteExternalSecretsPolicy:                   newServer(endpoints.DeleteExternalSecretsPolicy, decodeIDRequest),
		DeleteAlbIngressControllerPolicy:              newServer(endpoints.DeleteAlbIngressControllerPolicy, decodeIDRequest),
		DeleteExternalDNSPolicy:                       newServer(endpoints.DeleteExternalDNSPolicy, decodeIDRequest),
		DeleteHostedZone:                              newServer(endpoints.DeleteHostedZone, decodeDeleteHostedZone),
		DeleteExternalSecretsServiceAccount:           newServer(endpoints.DeleteExternalSecretsServiceAccount, decodeIDRequest),
		DeleteAlbIngressControllerServiceAccount:      newServer(endpoints.DeleteAlbIngressControllerServiceAccount, decodeIDRequest),
		DeleteExternalDNSServiceAccount:               newServer(endpoints.DeleteExternalDNSServiceAccount, decodeIDRequest),
		CreateIdentityPool:                            newServer(endpoints.CreateIdentityPool, decodeCreateIdentityPool),
		CreateIdentityPoolClient:                      newServer(endpoints.CreateIdentityPoolClient, decodeCreateIdentityPoolClient),
		CreateIdentityPoolUser:                        newServer(endpoints.CreateIdentityPoolUser, decodeCreateIdentityPoolUser),
		DeleteIdentityPool:                            newServer(endpoints.DeleteIdentityPool, decodeDeleteIdentityPool),
		DeleteIdentityPoolClient:                      newServer(endpoints.DeleteIdentityPoolClient, decodeDeleteIdentityPoolClient),
		CreateAWSLoadBalancerControllerServiceAccount: newServer(endpoints.CreateAWSLoadBalancerControllerServiceAccount, decodeCreateAWSLoadBalancerControllerServiceAccount),
		DeleteAWSLoadBalancerControllerServiceAccount: newServer(endpoints.DeleteAWSLoadBalancerControllerServiceAccount, decodeIDRequest),
		CreateAWSLoadBalancerControllerPolicy:         newServer(endpoints.CreateAWSLoadBalancerControllerPolicy, decodeCreateAWSLoadBalancerControllerPolicyRequest),
		DeleteAWSLoadBalancerControllerPolicy:         newServer(endpoints.DeleteAWSLoadBalancerControllerPolicy, decodeIDRequest),
		CreateAWSLoadBalancerControllerHelmChart:      newServer(endpoints.CreateAWSLoadBalancerControllerHelmChart, decodeCreateAWSLoadBalancerControllerHelmChart),
		DeleteCertificate:                             newServer(endpoints.DeleteCertificate, decodeDeleteCertificate),
		DeleteNamespace:                               newServer(endpoints.DeleteNamespace, decodeDeleteNamespace),
		DeleteCognitoCertificate:                      newServer(endpoints.DeleteCognitoCertificate, decodeDeleteCognitoCertificate),
		CreateAutoscalerHelmChart:                     newServer(endpoints.CreateAutoscalerHelmChart, decodeCreateAutoscalerHelmChart),
		CreateAutoscalerServiceAccount:                newServer(endpoints.CreateAutoscalerServiceAccount, decodeCreateAutoscalerServiceAccount),
		DeleteAutoscalerServiceAccount:                newServer(endpoints.DeleteAutoscalerServiceAccount, decodeIDRequest),
		CreateAutoscalerPolicy:                        newServer(endpoints.CreateAutoscalerPolicy, decodeCreateAutoscalerPolicy),
		DeleteAutoscalerPolicy:                        newServer(endpoints.DeleteAutoscalerPolicy, decodeIDRequest),
	}
}

// AttachRoutes creates a router and adds handlers to the corresponding routes
// nolint: funlen
func AttachRoutes(handlers *Handlers) http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Route("/clusters", func(r chi.Router) {
			r.Method(http.MethodPost, "/", handlers.CreateCluster)
			r.Method(http.MethodDelete, "/", handlers.DeleteCluster)
		})
		r.Route("/vpcs", func(r chi.Router) {
			r.Method(http.MethodPost, "/", handlers.CreateVpc)
			r.Method(http.MethodDelete, "/", handlers.DeleteVpc)
		})
		r.Route("/managedpolicies", func(r chi.Router) {
			r.Route("/externalsecrets", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalSecretsPolicy)
				r.Method(http.MethodDelete, "/", handlers.DeleteExternalSecretsPolicy)
			})
			r.Route("/albingresscontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAlbIngressControllerPolicy)
				r.Method(http.MethodDelete, "/", handlers.DeleteAlbIngressControllerPolicy)
			})
			r.Route("/awsloadbalancercontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAWSLoadBalancerControllerPolicy)
				r.Method(http.MethodDelete, "/", handlers.DeleteAWSLoadBalancerControllerPolicy)
			})
			r.Route("/externaldns", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalDNSPolicy)
				r.Method(http.MethodDelete, "/", handlers.DeleteExternalDNSPolicy)
			})
			r.Route("/autoscaler", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAutoscalerPolicy)
				r.Method(http.MethodDelete, "/", handlers.DeleteAutoscalerPolicy)
			})
		})
		r.Route("/serviceaccounts", func(r chi.Router) {
			r.Route("/externalsecrets", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalSecretsServiceAccount)
				r.Method(http.MethodDelete, "/", handlers.DeleteExternalSecretsServiceAccount)
			})
			r.Route("/albingresscontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAlbIngressControllerServiceAccount)
				r.Method(http.MethodDelete, "/", handlers.DeleteAlbIngressControllerServiceAccount)
			})
			r.Route("/awsloadbalancercontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAWSLoadBalancerControllerServiceAccount)
				r.Method(http.MethodDelete, "/", handlers.DeleteAWSLoadBalancerControllerServiceAccount)
			})
			r.Route("/externaldns", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalDNSServiceAccount)
				r.Method(http.MethodDelete, "/", handlers.DeleteExternalDNSServiceAccount)
			})
			r.Route("/autoscaler", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAutoscalerServiceAccount)
				r.Method(http.MethodDelete, "/", handlers.DeleteAutoscalerServiceAccount)
			})
		})
		r.Route("/helm", func(r chi.Router) {
			r.Route("/externalsecrets", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalSecretsHelmChart)
			})
			r.Route("/albingresscontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAlbIngressControllerHelmChart)
			})
			r.Route("/awsloadbalancercontroller", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAWSLoadBalancerControllerHelmChart)
			})
			r.Route("/argocd", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateArgoCD)
			})
			r.Route("/autoscaler", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateAutoscalerHelmChart)
			})
		})
		r.Route("/kube", func(r chi.Router) {
			r.Route("/externaldns", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalDNSKubeDeployment)
			})
			r.Route("/externalsecrets", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateExternalSecrets)
			})
			r.Route("/namespaces", func(r chi.Router) {
				r.Method(http.MethodDelete, "/", handlers.DeleteNamespace)
			})
		})
		r.Route("/domains", func(r chi.Router) {
			r.Route("/hostedzones", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateHostedZone)
				r.Method(http.MethodDelete, "/", handlers.DeleteHostedZone)
			})
		})
		r.Route("/certificates", func(r chi.Router) {
			r.Method(http.MethodPost, "/", handlers.CreateCertificate)
			r.Method(http.MethodDelete, "/", handlers.DeleteCertificate)
			r.Route("/cognito", func(r chi.Router) {
				r.Method(http.MethodDelete, "/", handlers.DeleteCognitoCertificate)
			})
		})
		r.Route("/parameters", func(r chi.Router) {
			r.Route("/secret", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateSecret)
				r.Method(http.MethodDelete, "/", handlers.DeleteSecret)
			})
		})
		r.Route("/identitymanagers", func(r chi.Router) {
			r.Route("/pools", func(r chi.Router) {
				r.Method(http.MethodPost, "/", handlers.CreateIdentityPool)
				r.Method(http.MethodDelete, "/", handlers.DeleteIdentityPool)
				r.Route("/clients", func(r chi.Router) {
					r.Method(http.MethodPost, "/", handlers.CreateIdentityPoolClient)
					r.Method(http.MethodDelete, "/", handlers.DeleteIdentityPoolClient)
				})
				r.Route("/users", func(r chi.Router) {
					r.Method(http.MethodPost, "/", handlers.CreateIdentityPoolUser)
				})
			})
		})
	})

	return r
}

// Services defines all available services
type Services struct {
	Cluster         api.ClusterService
	Vpc             api.VpcService
	ManagedPolicy   api.ManagedPolicyService
	ServiceAccount  api.ServiceAccountService
	Helm            api.HelmService
	Kube            api.KubeService
	Domain          api.DomainService
	Certificate     api.CertificateService
	Parameter       api.ParameterService
	IdentityManager api.IdentityManagerService
}

// EndpointOption makes it easy to enable and disable the endpoint
// middlewares
type EndpointOption func(Endpoints) Endpoints

const (
	clusterTag                   = "clusterService"
	vpcTag                       = "vpc"
	managedPoliciesTag           = "managedPolicies"
	externalSecretsTag           = "externalSecrets"
	serviceAccountsTag           = "serviceAccounts"
	helmTag                      = "helm"
	albIngressControllerTag      = "albingresscontroller"
	awsLoadBalancerControllerTag = "awsloadbalancercontroller"
	externalDNSTag               = "externaldns"
	kubeTag                      = "kube"
	domainTag                    = "domain"
	hostedZoneTag                = "hostedZone"
	certificateTag               = "certificate"
	parameterTag                 = "parameter"
	secretTag                    = "secret"
	argocdTag                    = "argocd"
	identityManagerTag           = "identitymanager"
	identityPoolTag              = "identitypool"
	identityPoolClientTag        = "identitypoolclient"
	identityPoolUserTag          = "identitypooluser"
	namespaceTag                 = "namespace"
	cognitoTag                   = "cognito"
	autoscalerTag                = "autoscaler"
)

// InstrumentEndpoints adds instrumentation to the endpoints
// nolint: lll
func InstrumentEndpoints(logger *logrus.Logger) EndpointOption {
	return func(endpoints Endpoints) Endpoints {
		return Endpoints{
			CreateCluster:                                 logmd.Logging(logger, clusterTag, "create")(endpoints.CreateCluster),
			DeleteCluster:                                 logmd.Logging(logger, clusterTag, "delete")(endpoints.DeleteCluster),
			CreateVpc:                                     logmd.Logging(logger, vpcTag, "create")(endpoints.CreateVpc),
			DeleteVpc:                                     logmd.Logging(logger, vpcTag, "delete")(endpoints.DeleteVpc),
			CreateExternalSecretsPolicy:                   logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, externalSecretsTag}, "/"), "create")(endpoints.CreateExternalSecretsPolicy),
			CreateExternalSecretsServiceAccount:           logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, externalSecretsTag}, "/"), "create")(endpoints.CreateExternalSecretsServiceAccount),
			CreateExternalSecretsHelmChart:                logmd.Logging(logger, strings.Join([]string{helmTag, externalSecretsTag}, "/"), "create")(endpoints.CreateExternalSecretsHelmChart),
			CreateAlbIngressControllerServiceAccount:      logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, albIngressControllerTag}, "/"), "create")(endpoints.CreateAlbIngressControllerServiceAccount),
			CreateAlbIngressControllerPolicy:              logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, albIngressControllerTag}, "/"), "create")(endpoints.CreateAlbIngressControllerPolicy),
			CreateAlbIngressControllerHelmChart:           logmd.Logging(logger, strings.Join([]string{helmTag, albIngressControllerTag}, "/"), "create")(endpoints.CreateAlbIngressControllerHelmChart),
			CreateExternalDNSPolicy:                       logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, externalDNSTag}, "/"), "create")(endpoints.CreateExternalDNSPolicy),
			CreateExternalDNSServiceAccount:               logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, externalDNSTag}, "/"), "create")(endpoints.CreateExternalDNSServiceAccount),
			CreateExternalDNSKubeDeployment:               logmd.Logging(logger, strings.Join([]string{kubeTag, externalDNSTag}, "/"), "create")(endpoints.CreateExternalDNSKubeDeployment),
			CreateHostedZone:                              logmd.Logging(logger, strings.Join([]string{domainTag, hostedZoneTag}, "/"), "create")(endpoints.CreateHostedZone),
			CreateCertificate:                             logmd.Logging(logger, certificateTag, "create")(endpoints.CreateCertificate),
			CreateSecret:                                  logmd.Logging(logger, strings.Join([]string{parameterTag, secretTag}, "/"), "create")(endpoints.CreateSecret),
			DeleteSecret:                                  logmd.Logging(logger, strings.Join([]string{parameterTag, secretTag}, "/"), "delete")(endpoints.DeleteSecret),
			CreateArgoCD:                                  logmd.Logging(logger, strings.Join([]string{helmTag, argocdTag}, "/"), "create")(endpoints.CreateArgoCD),
			CreateExternalSecrets:                         logmd.Logging(logger, strings.Join([]string{kubeTag, externalSecretsTag}, "/"), "create")(endpoints.CreateExternalSecrets),
			DeleteExternalSecretsPolicy:                   logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, externalSecretsTag}, "/"), "delete")(endpoints.DeleteExternalSecretsPolicy),
			DeleteAlbIngressControllerPolicy:              logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, albIngressControllerTag}, "/"), "delete")(endpoints.DeleteAlbIngressControllerPolicy),
			DeleteExternalDNSPolicy:                       logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, externalDNSTag}, "/"), "delete")(endpoints.DeleteExternalDNSPolicy),
			DeleteHostedZone:                              logmd.Logging(logger, strings.Join([]string{domainTag, hostedZoneTag}, "/"), "delete")(endpoints.DeleteHostedZone),
			DeleteExternalSecretsServiceAccount:           logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, externalSecretsTag}, "/"), "delete")(endpoints.DeleteExternalSecretsServiceAccount),
			DeleteAlbIngressControllerServiceAccount:      logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, albIngressControllerTag}, "/"), "delete")(endpoints.DeleteAlbIngressControllerServiceAccount),
			DeleteExternalDNSServiceAccount:               logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, externalDNSTag}, "/"), "delete")(endpoints.DeleteExternalDNSServiceAccount),
			CreateIdentityPool:                            logmd.Logging(logger, strings.Join([]string{identityManagerTag, identityPoolTag}, "/"), "create")(endpoints.CreateIdentityPool),
			CreateIdentityPoolClient:                      logmd.Logging(logger, strings.Join([]string{identityManagerTag, identityPoolTag, identityPoolClientTag}, "/"), "create")(endpoints.CreateIdentityPoolClient),
			CreateIdentityPoolUser:                        logmd.Logging(logger, strings.Join([]string{identityManagerTag, identityPoolTag, identityPoolUserTag}, "/"), "create")(endpoints.CreateIdentityPoolUser),
			DeleteIdentityPool:                            logmd.Logging(logger, strings.Join([]string{identityManagerTag, identityPoolTag}, "/"), "delete")(endpoints.DeleteIdentityPool),
			DeleteIdentityPoolClient:                      logmd.Logging(logger, strings.Join([]string{identityManagerTag, identityPoolClientTag}, "/"), "delete")(endpoints.DeleteIdentityPoolClient),
			CreateAWSLoadBalancerControllerServiceAccount: logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, awsLoadBalancerControllerTag}, "/"), "create")(endpoints.CreateAWSLoadBalancerControllerServiceAccount),
			DeleteAWSLoadBalancerControllerServiceAccount: logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, awsLoadBalancerControllerTag}, "/"), "delete")(endpoints.DeleteAWSLoadBalancerControllerServiceAccount),
			CreateAWSLoadBalancerControllerPolicy:         logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, awsLoadBalancerControllerTag}, "/"), "create")(endpoints.CreateAWSLoadBalancerControllerPolicy),
			DeleteAWSLoadBalancerControllerPolicy:         logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, awsLoadBalancerControllerTag}, "/"), "delete")(endpoints.DeleteAWSLoadBalancerControllerPolicy),
			CreateAWSLoadBalancerControllerHelmChart:      logmd.Logging(logger, strings.Join([]string{helmTag, awsLoadBalancerControllerTag}, "/"), "create")(endpoints.CreateAWSLoadBalancerControllerHelmChart),
			DeleteCertificate:                             logmd.Logging(logger, certificateTag, "delete")(endpoints.DeleteCertificate),
			DeleteNamespace:                               logmd.Logging(logger, strings.Join([]string{kubeTag, namespaceTag}, "/"), "delete")(endpoints.DeleteNamespace),
			DeleteCognitoCertificate:                      logmd.Logging(logger, strings.Join([]string{certificateTag, cognitoTag}, "/"), "delete")(endpoints.DeleteCognitoCertificate),
			CreateAutoscalerHelmChart:                     logmd.Logging(logger, strings.Join([]string{helmTag, autoscalerTag}, "/"), "create")(endpoints.CreateAutoscalerHelmChart),
			CreateAutoscalerServiceAccount:                logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, autoscalerTag}, "/"), "create")(endpoints.CreateAutoscalerServiceAccount),
			DeleteAutoscalerServiceAccount:                logmd.Logging(logger, strings.Join([]string{serviceAccountsTag, autoscalerTag}, "/"), "delete")(endpoints.DeleteAutoscalerServiceAccount),
			CreateAutoscalerPolicy:                        logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, autoscalerTag}, "/"), "create")(endpoints.CreateAutoscalerPolicy),
			DeleteAutoscalerPolicy:                        logmd.Logging(logger, strings.Join([]string{managedPoliciesTag, autoscalerTag}, "/"), "delete")(endpoints.DeleteAutoscalerPolicy),
		}
	}
}

// GenerateEndpoints is a convenience function for decorating endpoints
func GenerateEndpoints(s Services, opts ...EndpointOption) Endpoints {
	endpoints := MakeEndpoints(s)
	for _, opt := range opts {
		endpoints = opt(endpoints)
	}

	return endpoints
}
