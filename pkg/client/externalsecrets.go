package client

import (
	"context"

	"github.com/oslokommune/okctl/pkg/client/store"

	"github.com/oslokommune/okctl/pkg/api"
)

// ExternalSecrets is the content of an external-secrets deployment
type ExternalSecrets struct {
	Policy         *api.ManagedPolicy
	ServiceAccount *api.ServiceAccount
	Chart          *api.Helm
}

// CreateExternalSecretsOpts contains the required inputs
type CreateExternalSecretsOpts struct {
	ID api.ID
}

// ExternalSecretsService is an implementation of the business logic
type ExternalSecretsService interface {
	CreateExternalSecrets(ctx context.Context, opts CreateExternalSecretsOpts) (*ExternalSecrets, error)
}

// ExternalSecretsAPI invokes REST API endpoints
type ExternalSecretsAPI interface {
	CreateExternalSecretsPolicy(opts api.CreateExternalSecretsPolicyOpts) (*api.ManagedPolicy, error)
	CreateExternalSecretsServiceAccount(opts api.CreateExternalSecretsServiceAccountOpts) (*api.ServiceAccount, error)
	CreateExternalSecretsHelmChart(opts api.CreateExternalSecretsHelmChartOpts) (*api.Helm, error)
}

// ExternalSecretsStore is a storage layer implementation
type ExternalSecretsStore interface {
	SaveExternalSecrets(externalSecrets *ExternalSecrets) (*store.Report, error)
}

// ExternalSecretsReport is a report layer
type ExternalSecretsReport interface {
	ReportCreateExternalSecrets(secret *ExternalSecrets, report *store.Report) error
}