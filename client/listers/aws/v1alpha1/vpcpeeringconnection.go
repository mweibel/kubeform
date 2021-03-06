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

// VpcPeeringConnectionLister helps list VpcPeeringConnections.
type VpcPeeringConnectionLister interface {
	// List lists all VpcPeeringConnections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnection, err error)
	// VpcPeeringConnections returns an object that can list and get VpcPeeringConnections.
	VpcPeeringConnections(namespace string) VpcPeeringConnectionNamespaceLister
	VpcPeeringConnectionListerExpansion
}

// vpcPeeringConnectionLister implements the VpcPeeringConnectionLister interface.
type vpcPeeringConnectionLister struct {
	indexer cache.Indexer
}

// NewVpcPeeringConnectionLister returns a new VpcPeeringConnectionLister.
func NewVpcPeeringConnectionLister(indexer cache.Indexer) VpcPeeringConnectionLister {
	return &vpcPeeringConnectionLister{indexer: indexer}
}

// List lists all VpcPeeringConnections in the indexer.
func (s *vpcPeeringConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcPeeringConnection))
	})
	return ret, err
}

// VpcPeeringConnections returns an object that can list and get VpcPeeringConnections.
func (s *vpcPeeringConnectionLister) VpcPeeringConnections(namespace string) VpcPeeringConnectionNamespaceLister {
	return vpcPeeringConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcPeeringConnectionNamespaceLister helps list and get VpcPeeringConnections.
type VpcPeeringConnectionNamespaceLister interface {
	// List lists all VpcPeeringConnections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnection, err error)
	// Get retrieves the VpcPeeringConnection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VpcPeeringConnection, error)
	VpcPeeringConnectionNamespaceListerExpansion
}

// vpcPeeringConnectionNamespaceLister implements the VpcPeeringConnectionNamespaceLister
// interface.
type vpcPeeringConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcPeeringConnections in the indexer for a given namespace.
func (s vpcPeeringConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcPeeringConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcPeeringConnection))
	})
	return ret, err
}

// Get retrieves the VpcPeeringConnection from the indexer for a given namespace and name.
func (s vpcPeeringConnectionNamespaceLister) Get(name string) (*v1alpha1.VpcPeeringConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcpeeringconnection"), name)
	}
	return obj.(*v1alpha1.VpcPeeringConnection), nil
}
