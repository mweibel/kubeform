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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	googlev1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BigtableInstanceIamMemberInformer provides access to a shared informer and lister for
// BigtableInstanceIamMembers.
type BigtableInstanceIamMemberInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BigtableInstanceIamMemberLister
}

type bigtableInstanceIamMemberInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBigtableInstanceIamMemberInformer constructs a new informer for BigtableInstanceIamMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBigtableInstanceIamMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBigtableInstanceIamMemberInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBigtableInstanceIamMemberInformer constructs a new informer for BigtableInstanceIamMember type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBigtableInstanceIamMemberInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().BigtableInstanceIamMembers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().BigtableInstanceIamMembers(namespace).Watch(context.TODO(), options)
			},
		},
		&googlev1alpha1.BigtableInstanceIamMember{},
		resyncPeriod,
		indexers,
	)
}

func (f *bigtableInstanceIamMemberInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBigtableInstanceIamMemberInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bigtableInstanceIamMemberInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&googlev1alpha1.BigtableInstanceIamMember{}, f.defaultInformer)
}

func (f *bigtableInstanceIamMemberInformer) Lister() v1alpha1.BigtableInstanceIamMemberLister {
	return v1alpha1.NewBigtableInstanceIamMemberLister(f.Informer().GetIndexer())
}
