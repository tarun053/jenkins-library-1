package cmd

import (
	"fmt"

	"github.com/SAP/jenkins-library/pkg/command"
	piperhttp "github.com/SAP/jenkins-library/pkg/http"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperutils"
	"github.com/SAP/jenkins-library/pkg/telemetry"
)

type integrationArtifactIntegrationTestUtils interface {
	command.ExecRunner

	FileExists(filename string) (bool, error)

	// Add more methods here, or embed additional interfaces, or remove/replace as required.
	// The integrationArtifactIntegrationTestUtils interface should be descriptive of your runtime dependencies,
	// i.e. include everything you need to be able to mock in tests.
	// Unit tests shall be executable in parallel (not depend on global state), and don't (re-)test dependencies.
}

type integrationArtifactIntegrationTestUtilsBundle struct {
	*command.Command
	*piperutils.Files

	// Embed more structs as necessary to implement methods or interfaces you add to integrationArtifactIntegrationTestUtils.
	// Structs embedded in this way must each have a unique set of methods attached.
	// If there is no struct which implements the method you need, attach the method to
	// integrationArtifactIntegrationTestUtilsBundle and forward to the implementation of the dependency.
}

func newIntegrationArtifactIntegrationTestUtils() integrationArtifactIntegrationTestUtils {
	utils := integrationArtifactIntegrationTestUtilsBundle{
		Command: &command.Command{},
		Files:   &piperutils.Files{},
	}
	// Reroute command output to logging framework
	utils.Stdout(log.Writer())
	utils.Stderr(log.Writer())
	return &utils
}

func integrationArtifactIntegrationTest(config integrationArtifactIntegrationTestOptions, telemetryData *telemetry.CustomData) {
	// Utils can be used wherever the command.ExecRunner interface is expected.
	// It can also be used for example as a mavenExecRunner.
	utils := newIntegrationArtifactIntegrationTestUtils()
	httpClient := &piperhttp.Client{}
	// For HTTP calls import  piperhttp "github.com/SAP/jenkins-library/pkg/http"
	// and use a  &piperhttp.Client{} in a custom system
	// Example: step checkmarxExecuteScan.go

	// Error situations should be bubbled up until they reach the line below which will then stop execution
	// through the log.Entry().Fatal() call leading to an os.Exit(1) in the end.
	err := runIntegrationArtifactIntegrationTest(&config, telemetryData, utils, httpClient)
	if err != nil {
		log.Entry().WithError(err).Fatal("step execution failed")
	}
}

func runIntegrationArtifactIntegrationTest(config *integrationArtifactIntegrationTestOptions, telemetryData *telemetry.CustomData, utils integrationArtifactIntegrationTestUtils, httpClient piperhttp.Sender) error {
	getServiceEndpointOptions := integrationArtifactGetServiceEndpointOptions{
		Username:              config.Username,
		Password:              config.Password,
		IntegrationFlowID:     config.IntegrationFlowID,
		Platform:              config.Platform,
		Host:                  config.Host,
		OAuthTokenProviderURL: config.OAuthTokenProviderURL,
	}

	var commonPipelineEnvironment integrationArtifactGetServiceEndpointCommonPipelineEnvironment

	runIntegrationArtifactGetServiceEndpoint(&getServiceEndpointOptions, nil, httpClient, &commonPipelineEnvironment)
	serviceUrl := commonPipelineEnvironment.custom.iFlowServiceEndpoint
	log.Entry().Info("The Service URL : ", serviceUrl)
	log.Entry().WithField("LogField", "Log field content").Info("This is just a demo for a simple step.")

	// Example of calling methods from external dependencies directly on utils:
	exists, err := utils.FileExists("file.txt")
	if err != nil {
		// It is good practice to set an error category.
		// Most likely you want to do this at the place where enough context is known.
		log.SetErrorCategory(log.ErrorConfiguration)
		// Always wrap non-descriptive errors to enrich them with context for when they appear in the log:
		return fmt.Errorf("failed to check for important file: %w", err)
	}
	if !exists {
		log.SetErrorCategory(log.ErrorConfiguration)
		return fmt.Errorf("cannot run without important file")
	}

	return nil
}