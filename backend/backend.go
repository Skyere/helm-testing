package backend

import (
	"fmt"
	"testing"
	"tests/helper"
	"tests/services"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

func AvailabilityBackend(releaseName string, serviceName string, url string, kubectlOptions *k8s.KubectlOptions) func(t *testing.T) {
	return func(t *testing.T) {

		backServiceName := fmt.Sprintf("back-end-service-%s", releaseName)
		backService := k8s.GetService(t, kubectlOptions, backServiceName)

		backUrl := fmt.Sprintf("http://%s/dev/api/clubs", k8s.GetServiceEndpoint(t, kubectlOptions, backService, 80))

		helper.Verify(t, 200, backUrl, "city", 10)
	}
}
