package core

import (
	"context"

	"github.com/oslokommune/okctl/pkg/config/state"

	"github.com/oslokommune/okctl/pkg/spinner"

	"github.com/oslokommune/okctl/pkg/api"
	"github.com/oslokommune/okctl/pkg/client"
)

type parameterService struct {
	spinner spinner.Spinner
	api     client.ParameterAPI
	store   client.ParameterStore
	report  client.ParameterReport
}

func (s *parameterService) DeleteAllsecrets(ctx context.Context, cluster state.Cluster) error {
	clients := cluster.IdentityPool.Clients
	repos := cluster.Github.Repositories

	argoCdSecretPath := cluster.ArgoCD.SecretKey.Path
	if argoCdSecretPath != "" {
		err := s.DeleteSecret(ctx, api.DeleteSecretOpts{
			Name: argoCdSecretPath,
		})
		if err != nil {
			return err
		}
	}

	for c := range clients {
		clientSecretPath := clients[c].ClientSecret.Path
		if clientSecretPath != "" {
			err := s.DeleteSecret(ctx, api.DeleteSecretOpts{
				Name: clientSecretPath,
			})
			if err != nil {
				return err
			}
		}
	}

	for k := range repos {
		repoSecretPath := repos[k].DeployKey.PrivateKeySecret.Path
		if repoSecretPath != "" {
			err := s.DeleteSecret(ctx, api.DeleteSecretOpts{
				Name: repoSecretPath,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *parameterService) DeleteSecret(ctx context.Context, opts api.DeleteSecretOpts) error {
	err := s.spinner.Start("parameter")
	if err != nil {
		return err
	}

	defer func() {
		err = s.spinner.Stop()
	}()

	err = s.api.DeleteSecret(api.DeleteSecretOpts{
		Name: opts.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *parameterService) CreateSecret(_ context.Context, opts api.CreateSecretOpts) (*api.SecretParameter, error) {
	err := s.spinner.Start("parameter")
	if err != nil {
		return nil, err
	}

	defer func() {
		err = s.spinner.Stop()
	}()

	secret, err := s.api.CreateSecret(opts)
	if err != nil {
		return nil, err
	}

	report, err := s.store.SaveSecret(secret)
	if err != nil {
		return nil, err
	}

	err = s.report.SaveSecret(secret, report)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// NewParameterService returns an initialised service
func NewParameterService(spinner spinner.Spinner, api client.ParameterAPI, store client.ParameterStore, report client.ParameterReport) client.ParameterService {
	return &parameterService{
		spinner: spinner,
		api:     api,
		store:   store,
		report:  report,
	}
}
