package helper

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/helm"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// status
var WasDeploySuccessful bool = false

// Deploy your helmChart
// Put here your path to helm chart
func Deploy(t *testing.T, releaseName string, helmPath string, options *helm.Options) {
	helmChartPath, err := filepath.Abs(helmPath)
	require.NoError(t, err)

	helm.Install(t, options, helmChartPath, releaseName)

	WasDeploySuccessful = true
}

// Destroy release function
func Destroy(t *testing.T, releaseName string, options *helm.Options) {
	helm.Delete(t, options, releaseName, true)
}

// Verify server function
func Verify(t *testing.T, status int, url string, bodyw string, retries int) bool {
	sleep := 4 * time.Second
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		nil,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			isOk := statusCode == 200
			isBackEnd := assert.Contains(t, body, bodyw)
			return isOk && isBackEnd
		},
	)
}
