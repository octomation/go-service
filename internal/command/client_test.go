package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.octolab.org/template/service/internal/command"
	"go.octolab.org/template/service/internal/config"
)

func TestNewClient(t *testing.T) {
	root := command.NewClient(new(config.Service))
	require.NotNil(t, root)
	assert.NotEmpty(t, root.Use)
	assert.NotEmpty(t, root.Short)
	assert.NotEmpty(t, root.Long)
	assert.False(t, root.SilenceErrors)
	assert.True(t, root.SilenceUsage)
}
