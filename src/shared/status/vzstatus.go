/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

// This file contains all possible reasons for the Vizier components.
// These are reported to the operator via statusz endpoints.

var reasonToMessageMap = map[VizierReason]string{
	"":                             "",
	CloudConnectorFailedToConnect:  "Cloud connector failed to connect to Pixie Cloud. Please check your cloud address.",
	CloudConnectorInvalidDeployKey: "Invalid deploy key specified. Please check that the deploy key is correct.",
	CloudConnectorBasicQueryFailed: "Unable to run basic healthcheck query on cluster.",
}

// GetMessageFromReason gets the human-readable message for a vizier status reason.
func GetMessageFromReason(reason VizierReason) string {
	if msg, ok := reasonToMessageMap[reason]; ok {
		return msg
	}
	return ""
}

// VizierReason is the reason that Vizier is in its current state.
// All VizierReason values should be included in this file.
type VizierReason string

const (
	// CloudConnectorFailedToConnect is a status for when the cloud connector is unable to connect to the specified cloud addr.
	CloudConnectorFailedToConnect VizierReason = "CloudConnectFailed"
	// CloudConnectorInvalidDeployKey is a status for when the cloud connector has an invalid deploy key. This will prevent
	// the Vizier from properly registering.
	CloudConnectorInvalidDeployKey VizierReason = "InvalidDeployKey"
	// CloudConnectorBasicQueryFailed is a status for when the cloud connector is fully connected, but fails to run basic queries.
	CloudConnectorBasicQueryFailed VizierReason = "BasicQueryFailed"
	// CloudConnectorPodPending is a status when a cloud connector that is still in the Pending Kubernetes Phase.
	CloudConnectorPodPending VizierReason = "CloudConnectorPodPending"
	// CloudConnectorPodFailed is the status when a cloud connector pod is in the Failed Kubernetes Phase.
	CloudConnectorPodFailed VizierReason = "CloudConnectorPodFailed"
	// CloudConnectorMissing is the status when a cloud connector doesn't exist for a cluster.
	CloudConnectorMissing VizierReason = "CloudConnectorMissing"
)
