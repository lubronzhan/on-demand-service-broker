// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package boshdirector_test

import (
	"errors"

	"github.com/cloudfoundry/bosh-cli/director"
	"github.com/lubronzhan/on-demand-service-broker/boshdirector"
	"github.com/lubronzhan/on-demand-service-broker/boshdirector/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("deployments", func() {
	var (
		fakeDeployment *fakes.FakeBOSHDeployment
	)

	BeforeEach(func() {
		fakeDeployment = new(fakes.FakeBOSHDeployment)
		fakeDeployment.NameReturns("some-deployment")

		fakeDirector.DeploymentsReturns([]director.Deployment{fakeDeployment}, nil)
	})

	It("fetch the deployments", func() {
		expectedDeployments := []boshdirector.Deployment{
			{Name: "some-deployment"},
			{Name: "some-deployment"},
		}
		fakeDirector.DeploymentsReturns([]director.Deployment{fakeDeployment, fakeDeployment}, nil)
		deployments, err := c.GetDeployments(logger)
		Expect(err).NotTo(HaveOccurred())
		Expect(deployments).To(Equal(expectedDeployments))
	})

	It("returns an error if cannot fetch the deployments", func() {
		fakeDirector.DeploymentsReturns(nil, errors.New("oops"))
		_, err := c.GetDeployments(logger)
		Expect(err).To(MatchError(ContainSubstring("Cannot get the list of deployments")))
	})
})
