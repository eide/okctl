package core // nolint: dupl

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/oslokommune/okctl/pkg/api"
)

func makeCreateExternalSecretsPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateExternalSecretsPolicy(ctx, request.(api.CreateExternalSecretsPolicyOpts))
	}
}

func makeCreateAlbIngressControllerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateAlbIngressControllerPolicy(ctx, request.(api.CreateAlbIngressControllerPolicyOpts))
	}
}

func makeCreateAWSLoadBalancerControllerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateAWSLoadBalancerControllerPolicy(ctx, request.(api.CreateAWSLoadBalancerControllerPolicyOpts))
	}
}

func makeCreateExternalDNSPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateExternalDNSPolicy(ctx, request.(api.CreateExternalDNSPolicyOpts))
	}
}

func makeDeleteExternalSecretsPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &Empty{}, s.DeleteExternalSecretsPolicy(ctx, request.(api.ID))
	}
}

func makeDeleteAlbIngressControllerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &Empty{}, s.DeleteAlbIngressControllerPolicy(ctx, request.(api.ID))
	}
}

func makeDeleteAWSLoadBalancerControllerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &Empty{}, s.DeleteAWSLoadBalancerControllerPolicy(ctx, request.(api.ID))
	}
}

func makeDeleteExternalDNSPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &Empty{}, s.DeleteExternalDNSPolicy(ctx, request.(api.ID))
	}
}

func makeCreateAutoscalerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateAutoscalerPolicy(ctx, request.(api.CreateAutoscalerPolicy))
	}
}

func makeDeleteAutoscalerPolicyEndpoint(s api.ManagedPolicyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &Empty{}, s.DeleteAutoscalerPolicy(ctx, request.(api.ID))
	}
}
