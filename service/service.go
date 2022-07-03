package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

// status
var WasServiceCheck bool = false

func ServiceCheck(serviceName string, releaseName string, kubectlOptions *k8s.KubectlOptions) func(t *testing.T) {
	return func(t *testing.T) {
		ServiceName := fmt.Sprintf("%s-%s", releaseName, serviceName)

		k8s.WaitUntilServiceAvailable(t, kubectlOptions, ServiceName, 10, 5*time.Second)

		WasServiceCheck = true
	}
}
