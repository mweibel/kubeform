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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SourcerepoRepositoryIamPolicyLister helps list SourcerepoRepositoryIamPolicies.
type SourcerepoRepositoryIamPolicyLister interface {
	// List lists all SourcerepoRepositoryIamPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepositoryIamPolicy, err error)
	// SourcerepoRepositoryIamPolicies returns an object that can list and get SourcerepoRepositoryIamPolicies.
	SourcerepoRepositoryIamPolicies(namespace string) SourcerepoRepositoryIamPolicyNamespaceLister
	SourcerepoRepositoryIamPolicyListerExpansion
}

// sourcerepoRepositoryIamPolicyLister implements the SourcerepoRepositoryIamPolicyLister interface.
type sourcerepoRepositoryIamPolicyLister struct {
	indexer cache.Indexer
}

// NewSourcerepoRepositoryIamPolicyLister returns a new SourcerepoRepositoryIamPolicyLister.
func NewSourcerepoRepositoryIamPolicyLister(indexer cache.Indexer) SourcerepoRepositoryIamPolicyLister {
	return &sourcerepoRepositoryIamPolicyLister{indexer: indexer}
}

// List lists all SourcerepoRepositoryIamPolicies in the indexer.
func (s *sourcerepoRepositoryIamPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepositoryIamPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcerepoRepositoryIamPolicy))
	})
	return ret, err
}

// SourcerepoRepositoryIamPolicies returns an object that can list and get SourcerepoRepositoryIamPolicies.
func (s *sourcerepoRepositoryIamPolicyLister) SourcerepoRepositoryIamPolicies(namespace string) SourcerepoRepositoryIamPolicyNamespaceLister {
	return sourcerepoRepositoryIamPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourcerepoRepositoryIamPolicyNamespaceLister helps list and get SourcerepoRepositoryIamPolicies.
type SourcerepoRepositoryIamPolicyNamespaceLister interface {
	// List lists all SourcerepoRepositoryIamPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepositoryIamPolicy, err error)
	// Get retrieves the SourcerepoRepositoryIamPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SourcerepoRepositoryIamPolicy, error)
	SourcerepoRepositoryIamPolicyNamespaceListerExpansion
}

// sourcerepoRepositoryIamPolicyNamespaceLister implements the SourcerepoRepositoryIamPolicyNamespaceLister
// interface.
type sourcerepoRepositoryIamPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourcerepoRepositoryIamPolicies in the indexer for a given namespace.
func (s sourcerepoRepositoryIamPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SourcerepoRepositoryIamPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcerepoRepositoryIamPolicy))
	})
	return ret, err
}

// Get retrieves the SourcerepoRepositoryIamPolicy from the indexer for a given namespace and name.
func (s sourcerepoRepositoryIamPolicyNamespaceLister) Get(name string) (*v1alpha1.SourcerepoRepositoryIamPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sourcereporepositoryiampolicy"), name)
	}
	return obj.(*v1alpha1.SourcerepoRepositoryIamPolicy), nil
}
