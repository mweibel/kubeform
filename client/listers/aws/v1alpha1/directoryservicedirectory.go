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

// DirectoryServiceDirectoryLister helps list DirectoryServiceDirectories.
type DirectoryServiceDirectoryLister interface {
	// List lists all DirectoryServiceDirectories in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceDirectory, err error)
	// DirectoryServiceDirectories returns an object that can list and get DirectoryServiceDirectories.
	DirectoryServiceDirectories(namespace string) DirectoryServiceDirectoryNamespaceLister
	DirectoryServiceDirectoryListerExpansion
}

// directoryServiceDirectoryLister implements the DirectoryServiceDirectoryLister interface.
type directoryServiceDirectoryLister struct {
	indexer cache.Indexer
}

// NewDirectoryServiceDirectoryLister returns a new DirectoryServiceDirectoryLister.
func NewDirectoryServiceDirectoryLister(indexer cache.Indexer) DirectoryServiceDirectoryLister {
	return &directoryServiceDirectoryLister{indexer: indexer}
}

// List lists all DirectoryServiceDirectories in the indexer.
func (s *directoryServiceDirectoryLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceDirectory, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryServiceDirectory))
	})
	return ret, err
}

// DirectoryServiceDirectories returns an object that can list and get DirectoryServiceDirectories.
func (s *directoryServiceDirectoryLister) DirectoryServiceDirectories(namespace string) DirectoryServiceDirectoryNamespaceLister {
	return directoryServiceDirectoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectoryServiceDirectoryNamespaceLister helps list and get DirectoryServiceDirectories.
type DirectoryServiceDirectoryNamespaceLister interface {
	// List lists all DirectoryServiceDirectories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceDirectory, err error)
	// Get retrieves the DirectoryServiceDirectory from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DirectoryServiceDirectory, error)
	DirectoryServiceDirectoryNamespaceListerExpansion
}

// directoryServiceDirectoryNamespaceLister implements the DirectoryServiceDirectoryNamespaceLister
// interface.
type directoryServiceDirectoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectoryServiceDirectories in the indexer for a given namespace.
func (s directoryServiceDirectoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceDirectory, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryServiceDirectory))
	})
	return ret, err
}

// Get retrieves the DirectoryServiceDirectory from the indexer for a given namespace and name.
func (s directoryServiceDirectoryNamespaceLister) Get(name string) (*v1alpha1.DirectoryServiceDirectory, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("directoryservicedirectory"), name)
	}
	return obj.(*v1alpha1.DirectoryServiceDirectory), nil
}
