package lifecycle_test

import (
	"os"
	"testing"

	"github.com/lubronzhan/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	. "github.com/lubronzhan/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/lubronzhan/on-demand-service-broker/system_tests/test_helpers/env_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRedisLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redis Lifecycle Suite")
}

var (
	dopplerAddress    string
	deploymentOptions BrokerDeploymentOptions
	metricsOpsFile    string
)

var _ = BeforeSuite(func() {
	err := env_helpers.ValidateEnvVars("DOPPLER_ADDRESS", "CF_CLIENT_ID", "CF_CLIENT_SECRET", "CF_UAA_URL")
	Expect(err).ToNot(HaveOccurred(), "Doppler address must be set")

	dopplerAddress = os.Getenv("DOPPLER_ADDRESS")
	legacyMetrics := os.Getenv("LEGACY_SERVICE_METRICS")
	metricsOpsFile = "service_metrics.yml"
	if legacyMetrics == "true" {
		metricsOpsFile = "service_metrics_with_metron_agent.yml"
	}

	deploymentOptions = bosh_helpers.BrokerDeploymentOptions{
		ServiceMetrics: true,
		BrokerTLS:      true,
	}
})
