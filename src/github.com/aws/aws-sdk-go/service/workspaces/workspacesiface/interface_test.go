// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package workspacesiface_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*workspacesiface.WorkSpacesAPI)(nil), workspaces.New(nil))
}