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

// CloudfunctionsFunctionIamBindingLister helps list CloudfunctionsFunctionIamBindings.
type CloudfunctionsFunctionIamBindingLister interface {
	// List lists all CloudfunctionsFunctionIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunctionIamBinding, err error)
	// CloudfunctionsFunctionIamBindings returns an object that can list and get CloudfunctionsFunctionIamBindings.
	CloudfunctionsFunctionIamBindings(namespace string) CloudfunctionsFunctionIamBindingNamespaceLister
	CloudfunctionsFunctionIamBindingListerExpansion
}

// cloudfunctionsFunctionIamBindingLister implements the CloudfunctionsFunctionIamBindingLister interface.
type cloudfunctionsFunctionIamBindingLister struct {
	indexer cache.Indexer
}

// NewCloudfunctionsFunctionIamBindingLister returns a new CloudfunctionsFunctionIamBindingLister.
func NewCloudfunctionsFunctionIamBindingLister(indexer cache.Indexer) CloudfunctionsFunctionIamBindingLister {
	return &cloudfunctionsFunctionIamBindingLister{indexer: indexer}
}

// List lists all CloudfunctionsFunctionIamBindings in the indexer.
func (s *cloudfunctionsFunctionIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudfunctionsFunctionIamBinding))
	})
	return ret, err
}

// CloudfunctionsFunctionIamBindings returns an object that can list and get CloudfunctionsFunctionIamBindings.
func (s *cloudfunctionsFunctionIamBindingLister) CloudfunctionsFunctionIamBindings(namespace string) CloudfunctionsFunctionIamBindingNamespaceLister {
	return cloudfunctionsFunctionIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudfunctionsFunctionIamBindingNamespaceLister helps list and get CloudfunctionsFunctionIamBindings.
type CloudfunctionsFunctionIamBindingNamespaceLister interface {
	// List lists all CloudfunctionsFunctionIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunctionIamBinding, err error)
	// Get retrieves the CloudfunctionsFunctionIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudfunctionsFunctionIamBinding, error)
	CloudfunctionsFunctionIamBindingNamespaceListerExpansion
}

// cloudfunctionsFunctionIamBindingNamespaceLister implements the CloudfunctionsFunctionIamBindingNamespaceLister
// interface.
type cloudfunctionsFunctionIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudfunctionsFunctionIamBindings in the indexer for a given namespace.
func (s cloudfunctionsFunctionIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudfunctionsFunctionIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudfunctionsFunctionIamBinding))
	})
	return ret, err
}

// Get retrieves the CloudfunctionsFunctionIamBinding from the indexer for a given namespace and name.
func (s cloudfunctionsFunctionIamBindingNamespaceLister) Get(name string) (*v1alpha1.CloudfunctionsFunctionIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudfunctionsfunctioniambinding"), name)
	}
	return obj.(*v1alpha1.CloudfunctionsFunctionIamBinding), nil
}
