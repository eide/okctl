package client

import (
	"context"

	"github.com/oslokommune/okctl/pkg/api"
	"github.com/oslokommune/okctl/pkg/client/store"
)

// IdentityManagerService orchestrates the creation of an identity pool
type IdentityManagerService interface {
	CreateIdentityPool(ctx context.Context, opts api.CreateIdentityPoolOpts) (*api.IdentityPool, error)
	CreateIdentityPoolClient(ctx context.Context, opts api.CreateIdentityPoolClientOpts) (*api.IdentityPoolClient, error)
}

// IdentityManagerAPI invokes the API calls for creating an identity pool
type IdentityManagerAPI interface {
	CreateIdentityPool(opts api.CreateIdentityPoolOpts) (*api.IdentityPool, error)
	CreateIdentityPoolClient(opts api.CreateIdentityPoolClientOpts) (*api.IdentityPoolClient, error)
}

// IdentityManagerStore stores the data
type IdentityManagerStore interface {
	SaveIdentityPool(pool *api.IdentityPool) (*store.Report, error)
	SaveIdentityPoolClient(client *api.IdentityPoolClient) (*store.Report, error)
}

// IdentityManagerState implements the state layer
type IdentityManagerState interface {
	SaveIdentityPool(pool *api.IdentityPool) (*store.Report, error)
	SaveIdentityPoolClient(client *api.IdentityPoolClient) (*store.Report, error)
}

// IdentityManagerReport provides output of the actions
type IdentityManagerReport interface {
	ReportIdentityPool(pool *api.IdentityPool, reports []*store.Report) error
	ReportIdentityPoolClient(client *api.IdentityPoolClient, reports []*store.Report) error
}