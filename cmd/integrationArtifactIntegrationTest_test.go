package cmd

import (
	"testing"

	"github.com/SAP/jenkins-library/pkg/mock"
	"github.com/stretchr/testify/assert"
)

type integrationArtifactIntegrationTestMockUtils struct {
	*mock.ExecMockRunner
	*mock.FilesMock
}

func newIntegrationArtifactIntegrationTestTestsUtils() integrationArtifactIntegrationTestMockUtils {
	utils := integrationArtifactIntegrationTestMockUtils{
		ExecMockRunner: &mock.ExecMockRunner{},
		FilesMock:      &mock.FilesMock{},
	}
	return utils
}

func TestRunIntegrationArtifactIntegrationTest(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		// init
		config := integrationArtifactIntegrationTestOptions{
			Host:                  "https://demo",
			OAuthTokenProviderURL: "https://demo/oauth/token",
			Username:              "demouser",
			Password:              "******",
			IntegrationFlowID:     "CPI_IFlow_Call_using_Cert",
			Platform:              "cf",
		}

		utils := newIntegrationArtifactIntegrationTestTestsUtils()
		utils.AddFile("file.txt", []byte("dummy content"))
		httpClient := httpMockCpis{CPIFunction: "IntegrationArtifactGetServiceEndpoint", ResponseBody: ``, TestType: "PositiveAndGetetIntegrationArtifactGetServiceResBody"}

		// test
		err := runIntegrationArtifactIntegrationTest(&config, nil, utils, &httpClient)

		// assert
		assert.NoError(t, err)
	})

	// t.Run("error path", func(t *testing.T) {
	// 	t.Parallel()
	// 	// init
	// 	config := integrationArtifactIntegrationTestOptions{}

	// 	utils := newIntegrationArtifactIntegrationTestTestsUtils()

	// 	// test
	// 	err := runIntegrationArtifactIntegrationTest(&config, nil, utils)

	// 	// assert
	// 	assert.EqualError(t, err, "cannot run without important file")
	// })
}
