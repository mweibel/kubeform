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

// SagemakerEndpointConfigurationLister helps list SagemakerEndpointConfigurations.
type SagemakerEndpointConfigurationLister interface {
	// List lists all SagemakerEndpointConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpointConfiguration, err error)
	// SagemakerEndpointConfigurations returns an object that can list and get SagemakerEndpointConfigurations.
	SagemakerEndpointConfigurations(namespace string) SagemakerEndpointConfigurationNamespaceLister
	SagemakerEndpointConfigurationListerExpansion
}

// sagemakerEndpointConfigurationLister implements the SagemakerEndpointConfigurationLister interface.
type sagemakerEndpointConfigurationLister struct {
	indexer cache.Indexer
}

// NewSagemakerEndpointConfigurationLister returns a new SagemakerEndpointConfigurationLister.
func NewSagemakerEndpointConfigurationLister(indexer cache.Indexer) SagemakerEndpointConfigurationLister {
	return &sagemakerEndpointConfigurationLister{indexer: indexer}
}

// List lists all SagemakerEndpointConfigurations in the indexer.
func (s *sagemakerEndpointConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpointConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerEndpointConfiguration))
	})
	return ret, err
}

// SagemakerEndpointConfigurations returns an object that can list and get SagemakerEndpointConfigurations.
func (s *sagemakerEndpointConfigurationLister) SagemakerEndpointConfigurations(namespace string) SagemakerEndpointConfigurationNamespaceLister {
	return sagemakerEndpointConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SagemakerEndpointConfigurationNamespaceLister helps list and get SagemakerEndpointConfigurations.
type SagemakerEndpointConfigurationNamespaceLister interface {
	// List lists all SagemakerEndpointConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpointConfiguration, err error)
	// Get retrieves the SagemakerEndpointConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SagemakerEndpointConfiguration, error)
	SagemakerEndpointConfigurationNamespaceListerExpansion
}

// sagemakerEndpointConfigurationNamespaceLister implements the SagemakerEndpointConfigurationNamespaceLister
// interface.
type sagemakerEndpointConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SagemakerEndpointConfigurations in the indexer for a given namespace.
func (s sagemakerEndpointConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpointConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerEndpointConfiguration))
	})
	return ret, err
}

// Get retrieves the SagemakerEndpointConfiguration from the indexer for a given namespace and name.
func (s sagemakerEndpointConfigurationNamespaceLister) Get(name string) (*v1alpha1.SagemakerEndpointConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sagemakerendpointconfiguration"), name)
	}
	return obj.(*v1alpha1.SagemakerEndpointConfiguration), nil
}
