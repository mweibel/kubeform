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

// VpcEndpointServiceAllowedPrincipalLister helps list VpcEndpointServiceAllowedPrincipals.
type VpcEndpointServiceAllowedPrincipalLister interface {
	// List lists all VpcEndpointServiceAllowedPrincipals in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointServiceAllowedPrincipal, err error)
	// VpcEndpointServiceAllowedPrincipals returns an object that can list and get VpcEndpointServiceAllowedPrincipals.
	VpcEndpointServiceAllowedPrincipals(namespace string) VpcEndpointServiceAllowedPrincipalNamespaceLister
	VpcEndpointServiceAllowedPrincipalListerExpansion
}

// vpcEndpointServiceAllowedPrincipalLister implements the VpcEndpointServiceAllowedPrincipalLister interface.
type vpcEndpointServiceAllowedPrincipalLister struct {
	indexer cache.Indexer
}

// NewVpcEndpointServiceAllowedPrincipalLister returns a new VpcEndpointServiceAllowedPrincipalLister.
func NewVpcEndpointServiceAllowedPrincipalLister(indexer cache.Indexer) VpcEndpointServiceAllowedPrincipalLister {
	return &vpcEndpointServiceAllowedPrincipalLister{indexer: indexer}
}

// List lists all VpcEndpointServiceAllowedPrincipals in the indexer.
func (s *vpcEndpointServiceAllowedPrincipalLister) List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointServiceAllowedPrincipal, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcEndpointServiceAllowedPrincipal))
	})
	return ret, err
}

// VpcEndpointServiceAllowedPrincipals returns an object that can list and get VpcEndpointServiceAllowedPrincipals.
func (s *vpcEndpointServiceAllowedPrincipalLister) VpcEndpointServiceAllowedPrincipals(namespace string) VpcEndpointServiceAllowedPrincipalNamespaceLister {
	return vpcEndpointServiceAllowedPrincipalNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcEndpointServiceAllowedPrincipalNamespaceLister helps list and get VpcEndpointServiceAllowedPrincipals.
type VpcEndpointServiceAllowedPrincipalNamespaceLister interface {
	// List lists all VpcEndpointServiceAllowedPrincipals in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointServiceAllowedPrincipal, err error)
	// Get retrieves the VpcEndpointServiceAllowedPrincipal from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VpcEndpointServiceAllowedPrincipal, error)
	VpcEndpointServiceAllowedPrincipalNamespaceListerExpansion
}

// vpcEndpointServiceAllowedPrincipalNamespaceLister implements the VpcEndpointServiceAllowedPrincipalNamespaceLister
// interface.
type vpcEndpointServiceAllowedPrincipalNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcEndpointServiceAllowedPrincipals in the indexer for a given namespace.
func (s vpcEndpointServiceAllowedPrincipalNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcEndpointServiceAllowedPrincipal, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcEndpointServiceAllowedPrincipal))
	})
	return ret, err
}

// Get retrieves the VpcEndpointServiceAllowedPrincipal from the indexer for a given namespace and name.
func (s vpcEndpointServiceAllowedPrincipalNamespaceLister) Get(name string) (*v1alpha1.VpcEndpointServiceAllowedPrincipal, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcendpointserviceallowedprincipal"), name)
	}
	return obj.(*v1alpha1.VpcEndpointServiceAllowedPrincipal), nil
}
