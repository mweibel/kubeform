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

// AppmeshMeshLister helps list AppmeshMeshes.
type AppmeshMeshLister interface {
	// List lists all AppmeshMeshes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshMesh, err error)
	// AppmeshMeshes returns an object that can list and get AppmeshMeshes.
	AppmeshMeshes(namespace string) AppmeshMeshNamespaceLister
	AppmeshMeshListerExpansion
}

// appmeshMeshLister implements the AppmeshMeshLister interface.
type appmeshMeshLister struct {
	indexer cache.Indexer
}

// NewAppmeshMeshLister returns a new AppmeshMeshLister.
func NewAppmeshMeshLister(indexer cache.Indexer) AppmeshMeshLister {
	return &appmeshMeshLister{indexer: indexer}
}

// List lists all AppmeshMeshes in the indexer.
func (s *appmeshMeshLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshMesh, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshMesh))
	})
	return ret, err
}

// AppmeshMeshes returns an object that can list and get AppmeshMeshes.
func (s *appmeshMeshLister) AppmeshMeshes(namespace string) AppmeshMeshNamespaceLister {
	return appmeshMeshNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppmeshMeshNamespaceLister helps list and get AppmeshMeshes.
type AppmeshMeshNamespaceLister interface {
	// List lists all AppmeshMeshes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppmeshMesh, err error)
	// Get retrieves the AppmeshMesh from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppmeshMesh, error)
	AppmeshMeshNamespaceListerExpansion
}

// appmeshMeshNamespaceLister implements the AppmeshMeshNamespaceLister
// interface.
type appmeshMeshNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppmeshMeshes in the indexer for a given namespace.
func (s appmeshMeshNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppmeshMesh, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppmeshMesh))
	})
	return ret, err
}

// Get retrieves the AppmeshMesh from the indexer for a given namespace and name.
func (s appmeshMeshNamespaceLister) Get(name string) (*v1alpha1.AppmeshMesh, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appmeshmesh"), name)
	}
	return obj.(*v1alpha1.AppmeshMesh), nil
}
