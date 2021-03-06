package core

import (
	"context"

	"github.com/mishudark/errors"
	"github.com/oslokommune/okctl/pkg/api"
)

type managedPolicyService struct {
	provider api.ManagedPolicyCloudProvider
}

func (m *managedPolicyService) CreateAutoscalerPolicy(_ context.Context, opts api.CreateAutoscalerPolicy) (*api.ManagedPolicy, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "validating inputs", errors.Invalid)
	}

	got, err := m.provider.CreateAutoscalerPolicy(opts)
	if err != nil {
		return nil, errors.E(err, "creating autoscaler policy", errors.Internal)
	}

	return got, nil
}

func (m *managedPolicyService) DeleteAutoscalerPolicy(_ context.Context, id api.ID) error {
	err := id.Validate()
	if err != nil {
		return errors.E(err, "validating inputs", errors.Invalid)
	}

	err = m.provider.DeleteAutoscalerPolicy(id)
	if err != nil {
		return errors.E(err, "deleting autoscaler policy", errors.Internal)
	}

	return nil
}

func (m *managedPolicyService) DeleteExternalSecretsPolicy(_ context.Context, id api.ID) error {
	err := id.Validate()
	if err != nil {
		return errors.E(err, "failed to validate id", errors.Invalid)
	}

	err = m.provider.DeleteExternalSecretsPolicy(id)
	if err != nil {
		return errors.E(err, "failed to delete external secrets policy", errors.Internal)
	}

	return nil
}

func (m *managedPolicyService) DeleteAlbIngressControllerPolicy(_ context.Context, id api.ID) error {
	err := id.Validate()
	if err != nil {
		return errors.E(err, "failed to validate id", errors.Invalid)
	}

	err = m.provider.DeleteAlbIngressControllerPolicy(id)
	if err != nil {
		return errors.E(err, "failed to delete alb ingress controller policy", errors.Internal)
	}

	return nil
}

func (m *managedPolicyService) DeleteAWSLoadBalancerControllerPolicy(_ context.Context, id api.ID) error {
	err := id.Validate()
	if err != nil {
		return errors.E(err, "failed to validate id", errors.Invalid)
	}

	err = m.provider.DeleteAWSLoadBalancerControllerPolicy(id)
	if err != nil {
		return errors.E(err, "deleting aws load balancer controller policy", errors.Internal)
	}

	return nil
}

func (m *managedPolicyService) DeleteExternalDNSPolicy(_ context.Context, id api.ID) error {
	err := id.Validate()
	if err != nil {
		return errors.E(err, "failed to validate id", errors.Invalid)
	}

	err = m.provider.DeleteExternalDNSPolicy(id)
	if err != nil {
		return errors.E(err, "failed to delete external dns policy", errors.Internal)
	}

	return nil
}

func (m *managedPolicyService) CreateExternalDNSPolicy(_ context.Context, opts api.CreateExternalDNSPolicyOpts) (*api.ManagedPolicy, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "failed to validate external dns opts")
	}

	got, err := m.provider.CreateExternalDNSPolicy(opts)
	if err != nil {
		return nil, errors.E(err, "failed to create external dns policy")
	}

	return got, nil
}

func (m *managedPolicyService) CreateAlbIngressControllerPolicy(_ context.Context, opts api.CreateAlbIngressControllerPolicyOpts) (*api.ManagedPolicy, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "failed to validate alb ingress controller opts")
	}

	got, err := m.provider.CreateAlbIngressControllerPolicy(opts)
	if err != nil {
		return nil, errors.E(err, "failed to create alb ingress controller policy")
	}

	return got, nil
}

func (m *managedPolicyService) CreateAWSLoadBalancerControllerPolicy(_ context.Context, opts api.CreateAWSLoadBalancerControllerPolicyOpts) (*api.ManagedPolicy, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "validating aws load balancer controller opts")
	}

	got, err := m.provider.CreateAWSLoadBalancerControllerPolicy(opts)
	if err != nil {
		return nil, errors.E(err, "creating aws load balancer controller policy")
	}

	return got, nil
}

func (m *managedPolicyService) CreateExternalSecretsPolicy(_ context.Context, opts api.CreateExternalSecretsPolicyOpts) (*api.ManagedPolicy, error) {
	err := opts.Validate()
	if err != nil {
		return nil, errors.E(err, "failed to validate create external secrets policy options", errors.Invalid)
	}

	got, err := m.provider.CreateExternalSecretsPolicy(opts)
	if err != nil {
		return nil, errors.E(err, "failed to create external secrets policy")
	}

	return got, nil
}

// NewManagedPolicyService returns an initialised managed policy service
func NewManagedPolicyService(provider api.ManagedPolicyCloudProvider) api.ManagedPolicyService {
	return &managedPolicyService{
		provider: provider,
	}
}
