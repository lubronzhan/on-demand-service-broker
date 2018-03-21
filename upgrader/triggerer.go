package upgrader

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/service"
)

type UpgradeTriggerer struct {
	brokerServices BrokerServices
	instanceLister InstanceLister
	logger         Listener
}

func NewTriggerer(brokerServices BrokerServices, instanceLister InstanceLister, listener Listener) *UpgradeTriggerer {
	return &UpgradeTriggerer{
		brokerServices: brokerServices,
		instanceLister: instanceLister,
		logger:         listener,
	}
}

func (t *UpgradeTriggerer) TriggerUpgrade(instance service.Instance) (services.UpgradeOperation, error) {
	latestInstance, err := t.instanceLister.LatestInstanceInfo(instance)
	if err != nil {
		if err == service.InstanceNotFound {
			return services.UpgradeOperation{Type: services.InstanceNotFound}, nil
		}
		latestInstance = instance
		t.logger.FailedToRefreshInstanceInfo(instance.GUID)
	}

	operation, err := t.brokerServices.UpgradeInstance(latestInstance)
	if err != nil {
		return services.UpgradeOperation{}, fmt.Errorf("Upgrade failed for service instance %s: %s", instance.GUID, err)
	}

	return operation, nil
}
