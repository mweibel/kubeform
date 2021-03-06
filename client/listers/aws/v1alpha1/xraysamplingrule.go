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

// XraySamplingRuleLister helps list XraySamplingRules.
type XraySamplingRuleLister interface {
	// List lists all XraySamplingRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.XraySamplingRule, err error)
	// XraySamplingRules returns an object that can list and get XraySamplingRules.
	XraySamplingRules(namespace string) XraySamplingRuleNamespaceLister
	XraySamplingRuleListerExpansion
}

// xraySamplingRuleLister implements the XraySamplingRuleLister interface.
type xraySamplingRuleLister struct {
	indexer cache.Indexer
}

// NewXraySamplingRuleLister returns a new XraySamplingRuleLister.
func NewXraySamplingRuleLister(indexer cache.Indexer) XraySamplingRuleLister {
	return &xraySamplingRuleLister{indexer: indexer}
}

// List lists all XraySamplingRules in the indexer.
func (s *xraySamplingRuleLister) List(selector labels.Selector) (ret []*v1alpha1.XraySamplingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XraySamplingRule))
	})
	return ret, err
}

// XraySamplingRules returns an object that can list and get XraySamplingRules.
func (s *xraySamplingRuleLister) XraySamplingRules(namespace string) XraySamplingRuleNamespaceLister {
	return xraySamplingRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// XraySamplingRuleNamespaceLister helps list and get XraySamplingRules.
type XraySamplingRuleNamespaceLister interface {
	// List lists all XraySamplingRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.XraySamplingRule, err error)
	// Get retrieves the XraySamplingRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.XraySamplingRule, error)
	XraySamplingRuleNamespaceListerExpansion
}

// xraySamplingRuleNamespaceLister implements the XraySamplingRuleNamespaceLister
// interface.
type xraySamplingRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all XraySamplingRules in the indexer for a given namespace.
func (s xraySamplingRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.XraySamplingRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XraySamplingRule))
	})
	return ret, err
}

// Get retrieves the XraySamplingRule from the indexer for a given namespace and name.
func (s xraySamplingRuleNamespaceLister) Get(name string) (*v1alpha1.XraySamplingRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("xraysamplingrule"), name)
	}
	return obj.(*v1alpha1.XraySamplingRule), nil
}
