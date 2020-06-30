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

	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// S3BucketPublicAccessBlockInformer provides access to a shared informer and lister for
// S3BucketPublicAccessBlocks.
type S3BucketPublicAccessBlockInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.S3BucketPublicAccessBlockLister
}

type s3BucketPublicAccessBlockInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewS3BucketPublicAccessBlockInformer constructs a new informer for S3BucketPublicAccessBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewS3BucketPublicAccessBlockInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredS3BucketPublicAccessBlockInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredS3BucketPublicAccessBlockInformer constructs a new informer for S3BucketPublicAccessBlock type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredS3BucketPublicAccessBlockInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().S3BucketPublicAccessBlocks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().S3BucketPublicAccessBlocks(namespace).Watch(context.TODO(), options)
			},
		},
		&awsv1alpha1.S3BucketPublicAccessBlock{},
		resyncPeriod,
		indexers,
	)
}

func (f *s3BucketPublicAccessBlockInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredS3BucketPublicAccessBlockInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *s3BucketPublicAccessBlockInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.S3BucketPublicAccessBlock{}, f.defaultInformer)
}

func (f *s3BucketPublicAccessBlockInformer) Lister() v1alpha1.S3BucketPublicAccessBlockLister {
	return v1alpha1.NewS3BucketPublicAccessBlockLister(f.Informer().GetIndexer())
}
