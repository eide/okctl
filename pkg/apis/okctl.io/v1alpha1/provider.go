package v1alpha1

import (
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type CloudProvider interface {
	EC2() ec2iface.EC2API
	CloudFormation() cloudformationiface.CloudFormationAPI
	Region() string
}
