package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

// status
var WasServiceCheck bool = false

func ServiceCheck(serviceName string, releaseName string, kubectlOptions *k8s.KubectlOptions, retries int) func(t *testing.T) {
	return func(t *testing.T) {
		ServiceName := fmt.Sprintf("%s-%s", releaseName, serviceName)

		k8s.WaitUntilServiceAvailable(t, kubectlOptions, ServiceName, retries, 5*time.Second)

		WasServiceCheck = true
	}
}

func GetServiceEndpoint(t *testing.T, serviceName string, releaseName string, servicePort int, kubectlOptions *k8s.KubectlOptions) string {
	service := fmt.Sprintf("%s-%s", releaseName, serviceName)
	serviceInfo := k8s.GetService(t, kubectlOptions, service)
	return k8s.GetServiceEndpoint(t, kubectlOptions, serviceInfo, servicePort)
}
