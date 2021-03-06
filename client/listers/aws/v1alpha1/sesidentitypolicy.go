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

// SesIdentityPolicyLister helps list SesIdentityPolicies.
type SesIdentityPolicyLister interface {
	// List lists all SesIdentityPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SesIdentityPolicy, err error)
	// SesIdentityPolicies returns an object that can list and get SesIdentityPolicies.
	SesIdentityPolicies(namespace string) SesIdentityPolicyNamespaceLister
	SesIdentityPolicyListerExpansion
}

// sesIdentityPolicyLister implements the SesIdentityPolicyLister interface.
type sesIdentityPolicyLister struct {
	indexer cache.Indexer
}

// NewSesIdentityPolicyLister returns a new SesIdentityPolicyLister.
func NewSesIdentityPolicyLister(indexer cache.Indexer) SesIdentityPolicyLister {
	return &sesIdentityPolicyLister{indexer: indexer}
}

// List lists all SesIdentityPolicies in the indexer.
func (s *sesIdentityPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.SesIdentityPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesIdentityPolicy))
	})
	return ret, err
}

// SesIdentityPolicies returns an object that can list and get SesIdentityPolicies.
func (s *sesIdentityPolicyLister) SesIdentityPolicies(namespace string) SesIdentityPolicyNamespaceLister {
	return sesIdentityPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SesIdentityPolicyNamespaceLister helps list and get SesIdentityPolicies.
type SesIdentityPolicyNamespaceLister interface {
	// List lists all SesIdentityPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SesIdentityPolicy, err error)
	// Get retrieves the SesIdentityPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SesIdentityPolicy, error)
	SesIdentityPolicyNamespaceListerExpansion
}

// sesIdentityPolicyNamespaceLister implements the SesIdentityPolicyNamespaceLister
// interface.
type sesIdentityPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SesIdentityPolicies in the indexer for a given namespace.
func (s sesIdentityPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SesIdentityPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SesIdentityPolicy))
	})
	return ret, err
}

// Get retrieves the SesIdentityPolicy from the indexer for a given namespace and name.
func (s sesIdentityPolicyNamespaceLister) Get(name string) (*v1alpha1.SesIdentityPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sesidentitypolicy"), name)
	}
	return obj.(*v1alpha1.SesIdentityPolicy), nil
}
