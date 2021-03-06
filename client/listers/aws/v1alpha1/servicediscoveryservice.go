/*
Copyright The Kubeform Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceDiscoveryServiceLister helps list ServiceDiscoveryServices.
type ServiceDiscoveryServiceLister interface {
	// List lists all ServiceDiscoveryServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceDiscoveryService, err error)
	// ServiceDiscoveryServices returns an object that can list and get ServiceDiscoveryServices.
	ServiceDiscoveryServices(namespace string) ServiceDiscoveryServiceNamespaceLister
	ServiceDiscoveryServiceListerExpansion
}

// serviceDiscoveryServiceLister implements the ServiceDiscoveryServiceLister interface.
type serviceDiscoveryServiceLister struct {
	indexer cache.Indexer
}

// NewServiceDiscoveryServiceLister returns a new ServiceDiscoveryServiceLister.
func NewServiceDiscoveryServiceLister(indexer cache.Indexer) ServiceDiscoveryServiceLister {
	return &serviceDiscoveryServiceLister{indexer: indexer}
}

// List lists all ServiceDiscoveryServices in the indexer.
func (s *serviceDiscoveryServiceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceDiscoveryService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceDiscoveryService))
	})
	return ret, err
}

// ServiceDiscoveryServices returns an object that can list and get ServiceDiscoveryServices.
func (s *serviceDiscoveryServiceLister) ServiceDiscoveryServices(namespace string) ServiceDiscoveryServiceNamespaceLister {
	return serviceDiscoveryServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceDiscoveryServiceNamespaceLister helps list and get ServiceDiscoveryServices.
type ServiceDiscoveryServiceNamespaceLister interface {
	// List lists all ServiceDiscoveryServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceDiscoveryService, err error)
	// Get retrieves the ServiceDiscoveryService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceDiscoveryService, error)
	ServiceDiscoveryServiceNamespaceListerExpansion
}

// serviceDiscoveryServiceNamespaceLister implements the ServiceDiscoveryServiceNamespaceLister
// interface.
type serviceDiscoveryServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceDiscoveryServices in the indexer for a given namespace.
func (s serviceDiscoveryServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceDiscoveryService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceDiscoveryService))
	})
	return ret, err
}

// Get retrieves the ServiceDiscoveryService from the indexer for a given namespace and name.
func (s serviceDiscoveryServiceNamespaceLister) Get(name string) (*v1alpha1.ServiceDiscoveryService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicediscoveryservice"), name)
	}
	return obj.(*v1alpha1.ServiceDiscoveryService), nil
}
