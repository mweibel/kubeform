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

// RuntimeconfigConfigIamBindingLister helps list RuntimeconfigConfigIamBindings.
type RuntimeconfigConfigIamBindingLister interface {
	// List lists all RuntimeconfigConfigIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigConfigIamBinding, err error)
	// RuntimeconfigConfigIamBindings returns an object that can list and get RuntimeconfigConfigIamBindings.
	RuntimeconfigConfigIamBindings(namespace string) RuntimeconfigConfigIamBindingNamespaceLister
	RuntimeconfigConfigIamBindingListerExpansion
}

// runtimeconfigConfigIamBindingLister implements the RuntimeconfigConfigIamBindingLister interface.
type runtimeconfigConfigIamBindingLister struct {
	indexer cache.Indexer
}

// NewRuntimeconfigConfigIamBindingLister returns a new RuntimeconfigConfigIamBindingLister.
func NewRuntimeconfigConfigIamBindingLister(indexer cache.Indexer) RuntimeconfigConfigIamBindingLister {
	return &runtimeconfigConfigIamBindingLister{indexer: indexer}
}

// List lists all RuntimeconfigConfigIamBindings in the indexer.
func (s *runtimeconfigConfigIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeconfigConfigIamBinding))
	})
	return ret, err
}

// RuntimeconfigConfigIamBindings returns an object that can list and get RuntimeconfigConfigIamBindings.
func (s *runtimeconfigConfigIamBindingLister) RuntimeconfigConfigIamBindings(namespace string) RuntimeconfigConfigIamBindingNamespaceLister {
	return runtimeconfigConfigIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuntimeconfigConfigIamBindingNamespaceLister helps list and get RuntimeconfigConfigIamBindings.
type RuntimeconfigConfigIamBindingNamespaceLister interface {
	// List lists all RuntimeconfigConfigIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigConfigIamBinding, err error)
	// Get retrieves the RuntimeconfigConfigIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RuntimeconfigConfigIamBinding, error)
	RuntimeconfigConfigIamBindingNamespaceListerExpansion
}

// runtimeconfigConfigIamBindingNamespaceLister implements the RuntimeconfigConfigIamBindingNamespaceLister
// interface.
type runtimeconfigConfigIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RuntimeconfigConfigIamBindings in the indexer for a given namespace.
func (s runtimeconfigConfigIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RuntimeconfigConfigIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuntimeconfigConfigIamBinding))
	})
	return ret, err
}

// Get retrieves the RuntimeconfigConfigIamBinding from the indexer for a given namespace and name.
func (s runtimeconfigConfigIamBindingNamespaceLister) Get(name string) (*v1alpha1.RuntimeconfigConfigIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("runtimeconfigconfigiambinding"), name)
	}
	return obj.(*v1alpha1.RuntimeconfigConfigIamBinding), nil
}
