package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationArtifactIntegrationTestCommand(t *testing.T) {
	t.Parallel()

	testCmd := IntegrationArtifactIntegrationTestCommand()

	// only high level testing performed - details are tested in step generation procedure
	assert.Equal(t, "integrationArtifactIntegrationTest", testCmd.Use, "command name incorrect")

}
