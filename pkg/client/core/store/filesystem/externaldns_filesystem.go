package filesystem

import (
	"fmt"

	"github.com/oslokommune/okctl/pkg/api"

	"github.com/oslokommune/okctl/pkg/client"
	"github.com/oslokommune/okctl/pkg/client/store"
	"github.com/spf13/afero"
)

type externalDNSStore struct {
	policy         Paths
	serviceAccount Paths
	kube           Paths
	fs             *afero.Afero
}

// ExternalDNSKube contains the stored state for a kube deployment
type ExternalDNSKube struct {
	ID           api.ID
	HostedZoneID string
	DomainFilter string
	Manifests    []string
}

func (s *externalDNSStore) RemoveExternalDNS(_ api.ID) (*store.Report, error) {
	kube := &ExternalDNSKube{}

	report, err := store.NewFileSystem(s.policy.BaseDir, s.fs).
		Remove(s.policy.OutputFile).
		Remove(s.policy.CloudFormationFile).
		AlterStore(store.SetBaseDir(s.serviceAccount.BaseDir)).
		Remove(s.serviceAccount.OutputFile).
		Remove(s.serviceAccount.ConfigFile).
		AlterStore(store.SetBaseDir(s.kube.BaseDir)).
		GetStruct(s.kube.OutputFile, kube, store.FromJSON(), store.WithWriteIfNotExists([]byte("{}"))).
		ProcessGetStruct(s.kube.OutputFile, func(_ interface{}, operations store.Operations) error {
			for _, name := range kube.Manifests {
				operations.Remove(name)
			}

			return nil
		}).
		Remove(s.kube.OutputFile).
		Do()
	if err != nil {
		return nil, err
	}

	// Delete base directory, ignore error if directory is not empty
	_, basedir := store.NewFileSystem(s.policy.BaseDir, s.fs).
		Remove("").
		Do()
	if basedir != nil {
		return report, err
	}

	return report, err
}

func (s *externalDNSStore) SaveExternalDNS(d *client.ExternalDNS) (*store.Report, error) {
	policy := &ManagedPolicy{
		ID:        d.Policy.ID,
		StackName: d.Policy.StackName,
		PolicyARN: d.Policy.PolicyARN,
	}

	account := &ServiceAccount{
		ID:        d.ServiceAccount.ID,
		PolicyArn: d.ServiceAccount.PolicyArn,
	}

	kube := &ExternalDNSKube{
		ID:           d.Kube.ID,
		HostedZoneID: d.Kube.HostedZoneID,
		DomainFilter: d.Kube.DomainFilter,
		Manifests:    make([]string, len(d.Kube.Manifests)),
	}

	manifests := make([]store.AddStoreBytes, len(d.Kube.Manifests))

	i := 0

	for name, data := range d.Kube.Manifests {
		name := name
		data := data

		manifests[i] = store.AddStoreBytes{
			Name: name,
			Data: data,
		}

		kube.Manifests[i] = name

		i++
	}

	report, err := store.NewFileSystem(s.policy.BaseDir, s.fs).
		// Policy
		StoreStruct(s.policy.OutputFile, policy, store.ToJSON()).
		StoreBytes(s.policy.CloudFormationFile, d.Policy.CloudFormationTemplate).
		// ServiceAccount
		AlterStore(store.SetBaseDir(s.serviceAccount.BaseDir)).
		StoreStruct(s.serviceAccount.OutputFile, account, store.ToJSON()).
		StoreStruct(s.serviceAccount.ConfigFile, d.ServiceAccount.Config, store.ToJSON()).
		// Kube
		AlterStore(store.SetBaseDir(s.kube.BaseDir)).
		StoreStruct(s.kube.OutputFile, kube, store.ToJSON()).
		AddStoreBytes(manifests...).
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to store external dns: %w", err)
	}

	return report, nil
}

// NewExternalDNSStore returns an initialised store
func NewExternalDNSStore(policy, serviceAccount, kube Paths, fs *afero.Afero) client.ExternalDNSStore {
	return &externalDNSStore{
		policy:         policy,
		serviceAccount: serviceAccount,
		kube:           kube,
		fs:             fs,
	}
}
