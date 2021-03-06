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

// OrganizationsPolicyLister helps list OrganizationsPolicies.
type OrganizationsPolicyLister interface {
	// List lists all OrganizationsPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationsPolicy, err error)
	// OrganizationsPolicies returns an object that can list and get OrganizationsPolicies.
	OrganizationsPolicies(namespace string) OrganizationsPolicyNamespaceLister
	OrganizationsPolicyListerExpansion
}

// organizationsPolicyLister implements the OrganizationsPolicyLister interface.
type organizationsPolicyLister struct {
	indexer cache.Indexer
}

// NewOrganizationsPolicyLister returns a new OrganizationsPolicyLister.
func NewOrganizationsPolicyLister(indexer cache.Indexer) OrganizationsPolicyLister {
	return &organizationsPolicyLister{indexer: indexer}
}

// List lists all OrganizationsPolicies in the indexer.
func (s *organizationsPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationsPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationsPolicy))
	})
	return ret, err
}

// OrganizationsPolicies returns an object that can list and get OrganizationsPolicies.
func (s *organizationsPolicyLister) OrganizationsPolicies(namespace string) OrganizationsPolicyNamespaceLister {
	return organizationsPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationsPolicyNamespaceLister helps list and get OrganizationsPolicies.
type OrganizationsPolicyNamespaceLister interface {
	// List lists all OrganizationsPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationsPolicy, err error)
	// Get retrieves the OrganizationsPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OrganizationsPolicy, error)
	OrganizationsPolicyNamespaceListerExpansion
}

// organizationsPolicyNamespaceLister implements the OrganizationsPolicyNamespaceLister
// interface.
type organizationsPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationsPolicies in the indexer for a given namespace.
func (s organizationsPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationsPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationsPolicy))
	})
	return ret, err
}

// Get retrieves the OrganizationsPolicy from the indexer for a given namespace and name.
func (s organizationsPolicyNamespaceLister) Get(name string) (*v1alpha1.OrganizationsPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationspolicy"), name)
	}
	return obj.(*v1alpha1.OrganizationsPolicy), nil
}
