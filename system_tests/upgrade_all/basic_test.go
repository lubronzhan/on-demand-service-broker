// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upgrade_all_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/pborman/uuid"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/cf_helpers"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/service_helpers"
)

var _ = Describe("upgrade-all-service-instances errand", func() {

	var (
		brokerInfo bosh_helpers.BrokerInfo
		uniqueID   string
	)

	BeforeEach(func() {
		uniqueID = uuid.New()[:8]

		brokerInfo = bosh_helpers.DeployAndRegisterBroker(
			"basic-upgrade-"+uniqueID,
			bosh_helpers.BrokerDeploymentOptions{},
			service_helpers.Redis,
			[]string{
				"service_catalog.yml",
				"remove_parallel_upgrade.yml",
			})
	})

	AfterEach(func() {
		bosh_helpers.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
	})

	It("when there are no service instances provisioned, upgrade-all-service-instances runs successfully", func() {
		By("logging stdout to the errand output")
		session := bosh_helpers.RunErrand(brokerInfo.DeploymentName, "upgrade-all-service-instances")
		Expect(session).To(gbytes.Say("STARTING OPERATION"))
		Expect(session).To(gbytes.Say("FINISHED PROCESSING Status: SUCCESS"))
	})

	Context("upgrading some instances in series", func() {

		var appDetailsList []appDetails

		AfterEach(func() {
			for _, appDtls := range appDetailsList {
				cf_helpers.UnbindAndDeleteApp(appDtls.appName, appDtls.serviceName)
			}
		})

		It("succeeds", func() {
			serviceNumber := 2
			appDtlsCh := make(chan appDetails, serviceNumber)
			appPath := cf_helpers.GetAppPath(service_helpers.Redis)

			performInParallel(func() {
				defer GinkgoRecover()

				uuid := uuid.New()[:8]
				serviceName := "service-" + uuid
				appName := "app-" + uuid
				cf_helpers.CreateService(brokerInfo.ServiceOffering, "dedicated-vm", serviceName, "")

				serviceGUID := cf_helpers.ServiceInstanceGUID(serviceName)
				serviceDeploymentName := "service-instance_" + serviceGUID

				By("verifying that the persistence property starts as 'yes'", func() {
					manifest := bosh_helpers.GetManifest(serviceDeploymentName)
					instanceGroupProperties := bosh_helpers.FindInstanceGroupProperties(&manifest, "redis-server")
					Expect(instanceGroupProperties["redis"].(map[interface{}]interface{})["persistence"]).To(Equal("yes"))
				})

				appURL := cf_helpers.PushAndBindApp(appName, serviceName, appPath)
				cf_helpers.PutToTestApp(appURL, "uuid", uuid)

				appDtlsCh <- appDetails{
					uuid:                  uuid,
					appURL:                appURL,
					appName:               appName,
					serviceName:           serviceName,
					serviceDeploymentName: serviceDeploymentName,
				}
			}, serviceNumber)
			close(appDtlsCh)

			for dtls := range appDtlsCh {
				appDetailsList = append(appDetailsList, dtls)
			}

			By("changing the name of instance group and disabling persistence", func() {
				brokerInfo = bosh_helpers.DeployAndRegisterBroker(
					"basic-upgrade-"+uniqueID,
					bosh_helpers.BrokerDeploymentOptions{},
					service_helpers.Redis,
					[]string{
						"service_catalog_updated.yml",
						"remove_parallel_upgrade.yml",
					})
			})

			session := bosh_helpers.RunErrand(brokerInfo.DeploymentName, "upgrade-all-service-instances")
			Expect(session).To(gbytes.Say("STARTING OPERATION"))
			Expect(session).To(gbytes.Say("FINISHED PROCESSING Status: SUCCESS"))

			for _, appDtls := range appDetailsList {
				By("verifying the update changes were applied to the instance", func() {
					manifest := bosh_helpers.GetManifest(appDtls.serviceDeploymentName)
					instanceGroupProperties := bosh_helpers.FindInstanceGroupProperties(&manifest, "redis")
					Expect(instanceGroupProperties["redis"].(map[interface{}]interface{})["persistence"]).To(Equal("no"))
				})

				By("checking apps still have access to the data previously stored in their service", func() {
					Expect(cf_helpers.GetFromTestApp(appDtls.appURL, "uuid")).To(Equal(appDtls.uuid))
				})

			}
		})
	})
})
