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

// AcmCertificateLister helps list AcmCertificates.
type AcmCertificateLister interface {
	// List lists all AcmCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AcmCertificate, err error)
	// AcmCertificates returns an object that can list and get AcmCertificates.
	AcmCertificates(namespace string) AcmCertificateNamespaceLister
	AcmCertificateListerExpansion
}

// acmCertificateLister implements the AcmCertificateLister interface.
type acmCertificateLister struct {
	indexer cache.Indexer
}

// NewAcmCertificateLister returns a new AcmCertificateLister.
func NewAcmCertificateLister(indexer cache.Indexer) AcmCertificateLister {
	return &acmCertificateLister{indexer: indexer}
}

// List lists all AcmCertificates in the indexer.
func (s *acmCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.AcmCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AcmCertificate))
	})
	return ret, err
}

// AcmCertificates returns an object that can list and get AcmCertificates.
func (s *acmCertificateLister) AcmCertificates(namespace string) AcmCertificateNamespaceLister {
	return acmCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AcmCertificateNamespaceLister helps list and get AcmCertificates.
type AcmCertificateNamespaceLister interface {
	// List lists all AcmCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AcmCertificate, err error)
	// Get retrieves the AcmCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AcmCertificate, error)
	AcmCertificateNamespaceListerExpansion
}

// acmCertificateNamespaceLister implements the AcmCertificateNamespaceLister
// interface.
type acmCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AcmCertificates in the indexer for a given namespace.
func (s acmCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AcmCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AcmCertificate))
	})
	return ret, err
}

// Get retrieves the AcmCertificate from the indexer for a given namespace and name.
func (s acmCertificateNamespaceLister) Get(name string) (*v1alpha1.AcmCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("acmcertificate"), name)
	}
	return obj.(*v1alpha1.AcmCertificate), nil
}
