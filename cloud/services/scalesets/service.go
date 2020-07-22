/*
Copyright 2020 The Kubernetes Authors.

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
package scalesets

import (
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/loadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/services/resourceskus"
)

// Service provides operations on azure resources
type Service struct {
	Client
	ResourceSkusClient  resourceskus.Client
	LoadBalancersClient loadbalancers.Client
}

// NewService creates a new service.
func NewService(auth azure.Authorizer) *Service {
	return &Service{
		Client:              NewClient(auth),
		ResourceSkusClient:  resourceskus.NewClient(auth),
		LoadBalancersClient: loadbalancers.NewClient(auth),
	}
}