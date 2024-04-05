/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-api-metamodel/tests/go/generated/servicelogs/v1

import (
	"net/http"
	"path"
)

// ClusterClient is the client of the 'cluster' resource.
//
// Manages a specific Cluster for a cluster log.
type ClusterClient struct {
	transport http.RoundTripper
	path      string
}

// NewClusterClient creates a new client for the 'cluster'
// resource using the given transport to send the requests and receive the
// responses.
func NewClusterClient(transport http.RoundTripper, path string) *ClusterClient {
	return &ClusterClient{
		transport: transport,
		path:      path,
	}
}

// ClusterLogs returns the target 'cluster_logs_UUID' resource.
//
// Reference to the list of cluster logs for a specific cluster uuid.
func (c *ClusterClient) ClusterLogs() *ClusterLogsUUIDClient {
	return NewClusterLogsUUIDClient(
		c.transport,
		path.Join(c.path, "cluster_logs"),
	)
}
