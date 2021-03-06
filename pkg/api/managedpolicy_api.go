package api

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ManagedPolicy contains all state for a policy
type ManagedPolicy struct {
	ID                     ID
	StackName              string
	PolicyARN              string
	CloudFormationTemplate []byte
}

// CreateExternalSecretsPolicyOpts contains the options
// that are required for creating an external secrets policy
type CreateExternalSecretsPolicyOpts struct {
	ID ID
}

// Validate determines if the options are valid
func (o CreateExternalSecretsPolicyOpts) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required),
	)
}

// CreateAlbIngressControllerPolicyOpts contains the input
type CreateAlbIngressControllerPolicyOpts struct {
	ID ID
}

// Validate the input
func (o CreateAlbIngressControllerPolicyOpts) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required),
	)
}

// CreateAWSLoadBalancerControllerPolicyOpts contains the input
type CreateAWSLoadBalancerControllerPolicyOpts struct {
	ID ID
}

// Validate the input
func (o CreateAWSLoadBalancerControllerPolicyOpts) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required),
	)
}

// CreateExternalDNSPolicyOpts contains the input
type CreateExternalDNSPolicyOpts struct {
	ID ID
}

// Validate the input
func (o CreateExternalDNSPolicyOpts) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required),
	)
}

// CreateAutoscalerPolicy contains all required inputs
type CreateAutoscalerPolicy struct {
	ID ID
}

// Validate the inputs
func (o CreateAutoscalerPolicy) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.ID, validation.Required),
	)
}

// ManagedPolicyService defines the service layer for managed policies
type ManagedPolicyService interface {
	CreateExternalSecretsPolicy(ctx context.Context, opts CreateExternalSecretsPolicyOpts) (*ManagedPolicy, error)
	DeleteExternalSecretsPolicy(ctx context.Context, id ID) error
	CreateAlbIngressControllerPolicy(ctx context.Context, opts CreateAlbIngressControllerPolicyOpts) (*ManagedPolicy, error)
	DeleteAlbIngressControllerPolicy(ctx context.Context, id ID) error
	CreateAWSLoadBalancerControllerPolicy(ctx context.Context, opts CreateAWSLoadBalancerControllerPolicyOpts) (*ManagedPolicy, error)
	DeleteAWSLoadBalancerControllerPolicy(ctx context.Context, id ID) error
	CreateExternalDNSPolicy(ctx context.Context, opts CreateExternalDNSPolicyOpts) (*ManagedPolicy, error)
	DeleteExternalDNSPolicy(ctx context.Context, id ID) error
	CreateAutoscalerPolicy(ctx context.Context, opts CreateAutoscalerPolicy) (*ManagedPolicy, error)
	DeleteAutoscalerPolicy(ctx context.Context, id ID) error
}

// ManagedPolicyCloudProvider defines the cloud provider layer for managed policies
type ManagedPolicyCloudProvider interface {
	CreateExternalSecretsPolicy(opts CreateExternalSecretsPolicyOpts) (*ManagedPolicy, error)
	DeleteExternalSecretsPolicy(id ID) error
	CreateAlbIngressControllerPolicy(opts CreateAlbIngressControllerPolicyOpts) (*ManagedPolicy, error)
	DeleteAlbIngressControllerPolicy(id ID) error
	CreateAWSLoadBalancerControllerPolicy(opts CreateAWSLoadBalancerControllerPolicyOpts) (*ManagedPolicy, error)
	DeleteAWSLoadBalancerControllerPolicy(id ID) error
	CreateExternalDNSPolicy(opts CreateExternalDNSPolicyOpts) (*ManagedPolicy, error)
	DeleteExternalDNSPolicy(id ID) error
	CreateAutoscalerPolicy(opts CreateAutoscalerPolicy) (*ManagedPolicy, error)
	DeleteAutoscalerPolicy(id ID) error
}

// ManagedPolicyStore defines the storage layer for managed policies
type ManagedPolicyStore interface {
	SaveExternalSecretsPolicy(policy *ManagedPolicy) error
	GetExternalSecretsPolicy() (*ManagedPolicy, error)
	SaveAlbIngressControllerPolicy(policy *ManagedPolicy) error
	GetAlbIngressControllerPolicy() (*ManagedPolicy, error)
	SaveAWSLoadBalancerControllerPolicy(policy *ManagedPolicy) error
	GetAWSLoadBalancerControllerPolicy() (*ManagedPolicy, error)
	SaveExternalDNSPolicy(policy *ManagedPolicy) error
	GetExternalDNSPolicy() (*ManagedPolicy, error)
}
