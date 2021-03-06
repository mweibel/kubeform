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

// MainRouteTableAssociationLister helps list MainRouteTableAssociations.
type MainRouteTableAssociationLister interface {
	// List lists all MainRouteTableAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MainRouteTableAssociation, err error)
	// MainRouteTableAssociations returns an object that can list and get MainRouteTableAssociations.
	MainRouteTableAssociations(namespace string) MainRouteTableAssociationNamespaceLister
	MainRouteTableAssociationListerExpansion
}

// mainRouteTableAssociationLister implements the MainRouteTableAssociationLister interface.
type mainRouteTableAssociationLister struct {
	indexer cache.Indexer
}

// NewMainRouteTableAssociationLister returns a new MainRouteTableAssociationLister.
func NewMainRouteTableAssociationLister(indexer cache.Indexer) MainRouteTableAssociationLister {
	return &mainRouteTableAssociationLister{indexer: indexer}
}

// List lists all MainRouteTableAssociations in the indexer.
func (s *mainRouteTableAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.MainRouteTableAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MainRouteTableAssociation))
	})
	return ret, err
}

// MainRouteTableAssociations returns an object that can list and get MainRouteTableAssociations.
func (s *mainRouteTableAssociationLister) MainRouteTableAssociations(namespace string) MainRouteTableAssociationNamespaceLister {
	return mainRouteTableAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MainRouteTableAssociationNamespaceLister helps list and get MainRouteTableAssociations.
type MainRouteTableAssociationNamespaceLister interface {
	// List lists all MainRouteTableAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MainRouteTableAssociation, err error)
	// Get retrieves the MainRouteTableAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MainRouteTableAssociation, error)
	MainRouteTableAssociationNamespaceListerExpansion
}

// mainRouteTableAssociationNamespaceLister implements the MainRouteTableAssociationNamespaceLister
// interface.
type mainRouteTableAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MainRouteTableAssociations in the indexer for a given namespace.
func (s mainRouteTableAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MainRouteTableAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MainRouteTableAssociation))
	})
	return ret, err
}

// Get retrieves the MainRouteTableAssociation from the indexer for a given namespace and name.
func (s mainRouteTableAssociationNamespaceLister) Get(name string) (*v1alpha1.MainRouteTableAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mainroutetableassociation"), name)
	}
	return obj.(*v1alpha1.MainRouteTableAssociation), nil
}
