// Copyright (C) 2018-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under the
// terms of the under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
//
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package boshdirector

import (
	"github.com/lubronzhan/on-demand-service-broker/config"
)

func (c *Client) GetDNSAddresses(deploymentName string, dnsRequest []config.BindingDNS) (map[string]string, error) {
	addresses := map[string]string{}
	for _, req := range dnsRequest {
		providerId, err := c.dnsRetriever.LinkProviderID(deploymentName, req.InstanceGroup, req.LinkProvider)
		if err != nil {
			return nil, err
		}
		consumerId, err := c.dnsRetriever.CreateLinkConsumer(providerId)
		if err != nil {
			return nil, err
		}

		addr, err := c.dnsRetriever.GetLinkAddress(consumerId, req.Properties.AZS, req.Properties.Status)
		if err != nil {
			return nil, err
		}

		c.dnsRetriever.DeleteLinkConsumer(consumerId)
		addresses[req.Name] = addr
	}
	return addresses, nil
}
