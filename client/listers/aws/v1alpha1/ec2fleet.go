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

// Ec2FleetLister helps list Ec2Fleets.
type Ec2FleetLister interface {
	// List lists all Ec2Fleets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2Fleet, err error)
	// Ec2Fleets returns an object that can list and get Ec2Fleets.
	Ec2Fleets(namespace string) Ec2FleetNamespaceLister
	Ec2FleetListerExpansion
}

// ec2FleetLister implements the Ec2FleetLister interface.
type ec2FleetLister struct {
	indexer cache.Indexer
}

// NewEc2FleetLister returns a new Ec2FleetLister.
func NewEc2FleetLister(indexer cache.Indexer) Ec2FleetLister {
	return &ec2FleetLister{indexer: indexer}
}

// List lists all Ec2Fleets in the indexer.
func (s *ec2FleetLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2Fleet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2Fleet))
	})
	return ret, err
}

// Ec2Fleets returns an object that can list and get Ec2Fleets.
func (s *ec2FleetLister) Ec2Fleets(namespace string) Ec2FleetNamespaceLister {
	return ec2FleetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Ec2FleetNamespaceLister helps list and get Ec2Fleets.
type Ec2FleetNamespaceLister interface {
	// List lists all Ec2Fleets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2Fleet, err error)
	// Get retrieves the Ec2Fleet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Ec2Fleet, error)
	Ec2FleetNamespaceListerExpansion
}

// ec2FleetNamespaceLister implements the Ec2FleetNamespaceLister
// interface.
type ec2FleetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ec2Fleets in the indexer for a given namespace.
func (s ec2FleetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2Fleet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2Fleet))
	})
	return ret, err
}

// Get retrieves the Ec2Fleet from the indexer for a given namespace and name.
func (s ec2FleetNamespaceLister) Get(name string) (*v1alpha1.Ec2Fleet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ec2fleet"), name)
	}
	return obj.(*v1alpha1.Ec2Fleet), nil
}
