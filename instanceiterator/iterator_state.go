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

package instanceiterator

import (
	"errors"
	"fmt"

	"github.com/lubronzhan/on-demand-service-broker/service"
)

type instanceInfo struct {
	status        OperationState
	initialPlan   string
	operation     TriggeredOperation
	couldBeCanary bool
}

type iteratorState struct {
	guids           []string
	states          map[string]instanceInfo
	processCanaries bool
	// Required number of canaries to process.  Use 0 as 'no limit'.
	canaryLimit  int
	pos          int
	allInstances []service.Instance
}

type summary struct {
	orphaned  int
	succeeded int
	busy      int
	deleted   int
	skipped   int
}

func NewIteratorState(canaryInstances, allInstances []service.Instance, canaryLimit int) (*iteratorState, error) {
	us := iteratorState{}

	us.allInstances = allInstances
	us.processCanaries = len(canaryInstances) > 0
	us.canaryLimit = canaryLimit
	us.states = map[string]instanceInfo{}
	for _, i := range allInstances {
		us.guids = append(us.guids, i.GUID)
		us.states[i.GUID] = instanceInfo{status: OperationPending, initialPlan: i.PlanUniqueID}
	}
	for _, i := range canaryInstances {
		info, ok := us.states[i.GUID]
		if !ok {
			return nil, fmt.Errorf("Canary '%s' not in instance list", i.GUID)
		}
		info.couldBeCanary = true
		us.states[i.GUID] = info
	}
	return &us, nil
}

func (is *iteratorState) AllInstances() []service.Instance {
	return is.allInstances
}

func (is *iteratorState) IsProcessingCanaries() bool {
	return is.processCanaries
}

func (is *iteratorState) RewindAndResetBusyInstances() {
	is.pos = 0
	for k, v := range is.states {
		if v.status == OperationInProgress {
			v.status = OperationPending
			is.states[k] = v
		}
	}
}

func (is *iteratorState) HasInstancesToProcess() bool {
	return len(is.GetInstancesInStates(OperationPending, OperationAccepted)) > 0
}

func (is *iteratorState) HasInstancesProcessing() bool {
	return len(is.GetInstancesInStates(OperationAccepted)) > 0
}

func (is *iteratorState) HasFailures() bool {
	return len(is.GetInstancesInStates(OperationFailed)) > 0
}

func (is *iteratorState) InProgressInstances() []service.Instance {
	return is.GetInstancesInStates(OperationAccepted)
}

func (is *iteratorState) CountInProgressInstances() int {
	return len(is.InProgressInstances())
}

func (is *iteratorState) RetryBusyInstances() {
	is.pos = 0
}

func (is *iteratorState) NextPending() (service.Instance, error) {
	for is.pos < len(is.guids) {
		guid := is.guids[is.pos]
		is.pos++
		if is.processable(guid) {
			return service.Instance{GUID: guid, PlanUniqueID: is.states[guid].initialPlan}, nil
		}
	}
	return service.Instance{}, errors.New("Cannot retrieve next pending instance")
}

func (is *iteratorState) GetIteratorIndex() int {
	return len(is.GetInstancesInStates(OperationSucceeded, OperationAccepted, InstanceNotFound, OrphanDeployment)) + 1
}

func (is *iteratorState) GetGUIDsInStates(states ...OperationState) (guids []string) {
	guids = []string{}
	for _, i := range is.GetInstancesInStates(states...) {
		guids = append(guids, i.GUID)
	}
	return
}

func (is *iteratorState) GetInstancesInStates(states ...OperationState) (instances []service.Instance) {
	instances = []service.Instance{}
	for _, guid := range is.guids {
		inst := is.states[guid]
		if is.processCanaries && !inst.couldBeCanary {
			continue
		}
		for _, state := range states {
			if inst.status == state {
				instances = append(instances, service.Instance{GUID: guid, PlanUniqueID: inst.initialPlan})
			}
		}
	}
	return
}

func (is *iteratorState) Summary() summary {
	return summary{
		orphaned:  len(is.GetInstancesInStates(OrphanDeployment)),
		succeeded: len(is.GetInstancesInStates(OperationSucceeded)),
		busy:      len(is.GetInstancesInStates(OperationInProgress)),
		deleted:   len(is.GetInstancesInStates(InstanceNotFound)),
		skipped:   len(is.GetInstancesInStates(OperationSkipped)),
	}
}

func (is *iteratorState) SetState(guid string, state OperationState) error {
	info := is.states[guid]
	info.status = state
	is.states[guid] = info
	return nil
}

func (is *iteratorState) SetOperation(guid string, iteratorOp TriggeredOperation) {
	info := is.states[guid]
	info.operation = iteratorOp
	is.states[guid] = info
}

func (is *iteratorState) GetOperation(guid string) TriggeredOperation {
	return is.states[guid].operation
}

func (is *iteratorState) CurrentPhaseIsComplete() bool {
	if is.processCanaries {
		return is.canariesCompleted()
	}
	return is.allCompleted()
}

func (is *iteratorState) OutstandingCanaryCount() int {
	pending := 0
	triggered := 0

	for _, guid := range is.guids {
		info := is.states[guid]
		if !info.couldBeCanary {
			continue
		}
		if info.status == OperationPending {
			pending++
		} else {
			triggered++
		}
	}

	outstanding := pending
	if is.canaryLimit > 0 {
		outstanding = is.canaryLimit - triggered
	}

	return outstanding
}

func (is *iteratorState) canariesCompleted() bool {
	completedCanaries := 0
	for _, info := range is.states {
		if !info.couldBeCanary {
			continue
		}
		if !isFinalState(info.status) && is.canaryLimit == 0 {
			return false
		}
		if isFinalState(info.status) {
			completedCanaries++
		}
	}
	return completedCanaries >= is.canaryLimit
}

func (is *iteratorState) allCompleted() bool {
	for _, info := range is.states {
		if !isFinalState(info.status) {
			return false
		}
	}
	return true
}

func (is *iteratorState) MarkCanariesCompleted() {
	is.processCanaries = false
	is.pos = 0
}

func (is *iteratorState) CountInstancesInCurrentPhase() int {
	c := 0
	for _, inst := range is.states {
		if is.processCanaries && !inst.couldBeCanary {
			continue
		}
		c++
	}
	return c
}

func (is *iteratorState) processable(guid string) bool {
	return is.doingCanariesAndPendingCanary(guid) ||
		is.notDoingCanariesAndPendingInstance(guid)
}

func (is *iteratorState) doingCanariesAndPendingCanary(guid string) bool {
	return is.processCanaries &&
		is.states[guid].couldBeCanary &&
		is.states[guid].status == OperationPending
}

func (is *iteratorState) notDoingCanariesAndPendingInstance(guid string) bool {
	return !is.processCanaries &&
		is.states[guid].status == OperationPending
}

func isFinalState(status OperationState) bool {
	return status != OperationInProgress && status != OperationPending && status != OperationAccepted
}
