package ingress

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

func IngressCheck(ingressName string, releaseName string, kubectlOptions *k8s.KubectlOptions, retries int) func(t *testing.T) {
	return func(t *testing.T) {
		ingressName := fmt.Sprintf("%s-%s", releaseName, ingressName)

		k8s.WaitUntilIngressAvailable(t, kubectlOptions, ingressName, retries, 5*time.Second)
	}
}
