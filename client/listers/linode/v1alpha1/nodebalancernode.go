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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodebalancerNodeLister helps list NodebalancerNodes.
type NodebalancerNodeLister interface {
	// List lists all NodebalancerNodes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NodebalancerNode, err error)
	// NodebalancerNodes returns an object that can list and get NodebalancerNodes.
	NodebalancerNodes(namespace string) NodebalancerNodeNamespaceLister
	NodebalancerNodeListerExpansion
}

// nodebalancerNodeLister implements the NodebalancerNodeLister interface.
type nodebalancerNodeLister struct {
	indexer cache.Indexer
}

// NewNodebalancerNodeLister returns a new NodebalancerNodeLister.
func NewNodebalancerNodeLister(indexer cache.Indexer) NodebalancerNodeLister {
	return &nodebalancerNodeLister{indexer: indexer}
}

// List lists all NodebalancerNodes in the indexer.
func (s *nodebalancerNodeLister) List(selector labels.Selector) (ret []*v1alpha1.NodebalancerNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodebalancerNode))
	})
	return ret, err
}

// NodebalancerNodes returns an object that can list and get NodebalancerNodes.
func (s *nodebalancerNodeLister) NodebalancerNodes(namespace string) NodebalancerNodeNamespaceLister {
	return nodebalancerNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodebalancerNodeNamespaceLister helps list and get NodebalancerNodes.
type NodebalancerNodeNamespaceLister interface {
	// List lists all NodebalancerNodes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NodebalancerNode, err error)
	// Get retrieves the NodebalancerNode from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NodebalancerNode, error)
	NodebalancerNodeNamespaceListerExpansion
}

// nodebalancerNodeNamespaceLister implements the NodebalancerNodeNamespaceLister
// interface.
type nodebalancerNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodebalancerNodes in the indexer for a given namespace.
func (s nodebalancerNodeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NodebalancerNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodebalancerNode))
	})
	return ret, err
}

// Get retrieves the NodebalancerNode from the indexer for a given namespace and name.
func (s nodebalancerNodeNamespaceLister) Get(name string) (*v1alpha1.NodebalancerNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodebalancernode"), name)
	}
	return obj.(*v1alpha1.NodebalancerNode), nil
}
