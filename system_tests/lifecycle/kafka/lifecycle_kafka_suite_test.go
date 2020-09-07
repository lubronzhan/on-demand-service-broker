package lifecycle_test

import (
	"testing"

	. "github.com/lubronzhan/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/lubronzhan/on-demand-service-broker/system_tests/test_helpers/service_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
)

func TestKafkaLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafka Lifecycle Suite")
}

var (
	brokerInfo BrokerInfo
)

var _ = BeforeSuite(func() {
	uniqueID := uuid.New()[:6]
	brokerInfo = BrokerInfo{}

	brokerInfo = DeployAndRegisterBroker(
		"-kafka-lifecycle-"+uniqueID,
		BrokerDeploymentOptions{},
		service_helpers.Kafka,
		[]string{"basic_service_catalog.yml"},
	)
})

var _ = AfterSuite(func() {
	DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
})
