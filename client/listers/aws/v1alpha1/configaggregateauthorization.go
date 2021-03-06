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

// ConfigAggregateAuthorizationLister helps list ConfigAggregateAuthorizations.
type ConfigAggregateAuthorizationLister interface {
	// List lists all ConfigAggregateAuthorizations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigAggregateAuthorization, err error)
	// ConfigAggregateAuthorizations returns an object that can list and get ConfigAggregateAuthorizations.
	ConfigAggregateAuthorizations(namespace string) ConfigAggregateAuthorizationNamespaceLister
	ConfigAggregateAuthorizationListerExpansion
}

// configAggregateAuthorizationLister implements the ConfigAggregateAuthorizationLister interface.
type configAggregateAuthorizationLister struct {
	indexer cache.Indexer
}

// NewConfigAggregateAuthorizationLister returns a new ConfigAggregateAuthorizationLister.
func NewConfigAggregateAuthorizationLister(indexer cache.Indexer) ConfigAggregateAuthorizationLister {
	return &configAggregateAuthorizationLister{indexer: indexer}
}

// List lists all ConfigAggregateAuthorizations in the indexer.
func (s *configAggregateAuthorizationLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigAggregateAuthorization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigAggregateAuthorization))
	})
	return ret, err
}

// ConfigAggregateAuthorizations returns an object that can list and get ConfigAggregateAuthorizations.
func (s *configAggregateAuthorizationLister) ConfigAggregateAuthorizations(namespace string) ConfigAggregateAuthorizationNamespaceLister {
	return configAggregateAuthorizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigAggregateAuthorizationNamespaceLister helps list and get ConfigAggregateAuthorizations.
type ConfigAggregateAuthorizationNamespaceLister interface {
	// List lists all ConfigAggregateAuthorizations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigAggregateAuthorization, err error)
	// Get retrieves the ConfigAggregateAuthorization from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ConfigAggregateAuthorization, error)
	ConfigAggregateAuthorizationNamespaceListerExpansion
}

// configAggregateAuthorizationNamespaceLister implements the ConfigAggregateAuthorizationNamespaceLister
// interface.
type configAggregateAuthorizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigAggregateAuthorizations in the indexer for a given namespace.
func (s configAggregateAuthorizationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigAggregateAuthorization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigAggregateAuthorization))
	})
	return ret, err
}

// Get retrieves the ConfigAggregateAuthorization from the indexer for a given namespace and name.
func (s configAggregateAuthorizationNamespaceLister) Get(name string) (*v1alpha1.ConfigAggregateAuthorization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configaggregateauthorization"), name)
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), nil
}
