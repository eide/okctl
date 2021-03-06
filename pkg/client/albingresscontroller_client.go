package client

import (
	"context"

	"github.com/oslokommune/okctl/pkg/api"
	"github.com/oslokommune/okctl/pkg/client/store"
)

// ALBIngressController contains the state of an
// alb ingress controller deployment
type ALBIngressController struct {
	Policy         *api.ManagedPolicy
	ServiceAccount *api.ServiceAccount
	Chart          *api.Helm
}

// CreateALBIngressControllerOpts defines the required inputs
type CreateALBIngressControllerOpts struct {
	ID    api.ID
	VPCID string
}

// ALBIngressControllerService defines the service layer
type ALBIngressControllerService interface {
	CreateALBIngressController(ctx context.Context, opts CreateALBIngressControllerOpts) (*ALBIngressController, error)
	DeleteALBIngressController(ctx context.Context, id api.ID) error
}

// ALBIngressControllerAPI defines the API layer
type ALBIngressControllerAPI interface {
	CreateAlbIngressControllerPolicy(opts api.CreateAlbIngressControllerPolicyOpts) (*api.ManagedPolicy, error)
	DeleteAlbIngressControllerPolicy(id api.ID) error
	CreateAlbIngressControllerServiceAccount(opts api.CreateAlbIngressControllerServiceAccountOpts) (*api.ServiceAccount, error)
	DeleteAlbIngressControllerServiceAccount(id api.ID) error
	CreateAlbIngressControllerHelmChart(opts api.CreateAlbIngressControllerHelmChartOpts) (*api.Helm, error)
}

// ALBIngressControllerStore defines the storage layer
type ALBIngressControllerStore interface {
	SaveALBIngressController(controller *ALBIngressController) (*store.Report, error)
	RemoveALBIngressController(id api.ID) (*store.Report, error)
}

// ALBIngressControllerReport defines the report layer
type ALBIngressControllerReport interface {
	ReportCreateALBIngressController(controller *ALBIngressController, report *store.Report) error
	ReportDeleteALBIngressController(report *store.Report) error
}
