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

// OpsworksNodejsAppLayerLister helps list OpsworksNodejsAppLayers.
type OpsworksNodejsAppLayerLister interface {
	// List lists all OpsworksNodejsAppLayers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksNodejsAppLayer, err error)
	// OpsworksNodejsAppLayers returns an object that can list and get OpsworksNodejsAppLayers.
	OpsworksNodejsAppLayers(namespace string) OpsworksNodejsAppLayerNamespaceLister
	OpsworksNodejsAppLayerListerExpansion
}

// opsworksNodejsAppLayerLister implements the OpsworksNodejsAppLayerLister interface.
type opsworksNodejsAppLayerLister struct {
	indexer cache.Indexer
}

// NewOpsworksNodejsAppLayerLister returns a new OpsworksNodejsAppLayerLister.
func NewOpsworksNodejsAppLayerLister(indexer cache.Indexer) OpsworksNodejsAppLayerLister {
	return &opsworksNodejsAppLayerLister{indexer: indexer}
}

// List lists all OpsworksNodejsAppLayers in the indexer.
func (s *opsworksNodejsAppLayerLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksNodejsAppLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksNodejsAppLayer))
	})
	return ret, err
}

// OpsworksNodejsAppLayers returns an object that can list and get OpsworksNodejsAppLayers.
func (s *opsworksNodejsAppLayerLister) OpsworksNodejsAppLayers(namespace string) OpsworksNodejsAppLayerNamespaceLister {
	return opsworksNodejsAppLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsworksNodejsAppLayerNamespaceLister helps list and get OpsworksNodejsAppLayers.
type OpsworksNodejsAppLayerNamespaceLister interface {
	// List lists all OpsworksNodejsAppLayers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.OpsworksNodejsAppLayer, err error)
	// Get retrieves the OpsworksNodejsAppLayer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.OpsworksNodejsAppLayer, error)
	OpsworksNodejsAppLayerNamespaceListerExpansion
}

// opsworksNodejsAppLayerNamespaceLister implements the OpsworksNodejsAppLayerNamespaceLister
// interface.
type opsworksNodejsAppLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsworksNodejsAppLayers in the indexer for a given namespace.
func (s opsworksNodejsAppLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsworksNodejsAppLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsworksNodejsAppLayer))
	})
	return ret, err
}

// Get retrieves the OpsworksNodejsAppLayer from the indexer for a given namespace and name.
func (s opsworksNodejsAppLayerNamespaceLister) Get(name string) (*v1alpha1.OpsworksNodejsAppLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsworksnodejsapplayer"), name)
	}
	return obj.(*v1alpha1.OpsworksNodejsAppLayer), nil
}
