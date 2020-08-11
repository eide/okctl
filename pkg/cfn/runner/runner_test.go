package runner_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/oslokommune/okctl/pkg/cfn/runner"
	"github.com/oslokommune/okctl/pkg/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		name        string
		runner      *runner.Runner
		expect      interface{}
		expectError bool
	}{
		{
			name: "Should work",
			runner: runner.
				New(mock.DefaultStackName, []byte{},
					mock.NewCloudProvider().
						DescribeStacksEmpty().
						CreateStackSuccess().
						DescribeStacksResponse(cloudformation.StackStatusCreateComplete),
				),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.runner.CreateIfNotExists(10)
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expect, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name        string
		runner      *runner.Runner
		expect      interface{}
		expectError bool
	}{
		{
			name: "Should work",
			runner: runner.
				New(mock.DefaultStackName, []byte{},
					mock.NewCloudProvider().
						DeleteStackSuccess().
						DescribeStacksResponse(cloudformation.StackStatusDeleteComplete),
				),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.runner.Delete()
			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, tc.expect, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}