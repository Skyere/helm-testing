package ingress

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
)

func IngressCheck(releaseName string, kubectlOptions *k8s.KubectlOptions) func(t *testing.T) {
	return func(t *testing.T) {
		ingressName := fmt.Sprintf("app-dns-%s", releaseName)

		k8s.WaitUntilIngressAvailable(t, kubectlOptions, ingressName, 10, 5*time.Second)
	}
}
