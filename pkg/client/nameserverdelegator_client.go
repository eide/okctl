package client

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/oslokommune/okctl/pkg/api"
)

// NameserverRecord defines necessary information required to request a nameserver delegation
type NameserverRecord struct {
	FQDN        string
	Nameservers []string
}

// Validate ensures a NameserverRecord contains required information
func (n NameserverRecord) Validate() error {
	return validation.ValidateStruct(&n,
		validation.Field(&n.FQDN, validation.Required),
		validation.Field(&n.Nameservers, validation.Length(1, 0)),
	)
}

// CreateNameserverDelegationRequestOpts contains the required information a NameserverRecordDelegationService needs to do create a delegation
// request
type CreateNameserverDelegationRequestOpts struct {
	ClusterID api.ID

	PrimaryHostedZoneFQDN string
	Nameservers           []string
}

/*NameserverRecordDelegationService defines required functionality for requesting a nameserver delegation record in
the top level domain.

If a team wants 'team.oslo.systems', okctl will create that domain which will get its own nameservers assigned.
The top level domain 'oslo.systems' then needs to delegate DNS inquiries for 'team.oslo.systems' to the assigned
nameservers. This is the delegation this service should handle.
*/
type NameserverRecordDelegationService interface {
	// CreateNameserverRecordDelegationRequest creates a request for a NS record in the top level domain
	CreateNameserverRecordDelegationRequest(opts *CreateNameserverDelegationRequestOpts) (record *NameserverRecord, err error)
}

// HostedZoneDelegationSetter defines a function used to set a hosted zone as delegated in state and store
type HostedZoneDelegationSetter func(domain string, isDelegated bool) error
