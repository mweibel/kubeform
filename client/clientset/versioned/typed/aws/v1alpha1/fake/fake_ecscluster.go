/*
Copyright 2019 The Kubeform Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeEcsClusters implements EcsClusterInterface
type FakeEcsClusters struct {
	Fake *FakeAwsV1alpha1
}

var ecsclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ecsclusters"}

var ecsclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EcsCluster"}

// Get takes name of the ecsCluster, and returns the corresponding ecsCluster object, and an error if there is any.
func (c *FakeEcsClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.EcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ecsclustersResource, name), &v1alpha1.EcsCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsCluster), err
}

// List takes label and field selectors, and returns the list of EcsClusters that match those selectors.
func (c *FakeEcsClusters) List(opts v1.ListOptions) (result *v1alpha1.EcsClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ecsclustersResource, ecsclustersKind, opts), &v1alpha1.EcsClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EcsClusterList{ListMeta: obj.(*v1alpha1.EcsClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.EcsClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ecsClusters.
func (c *FakeEcsClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ecsclustersResource, opts))
}

// Create takes the representation of a ecsCluster and creates it.  Returns the server's representation of the ecsCluster, and an error, if there is any.
func (c *FakeEcsClusters) Create(ecsCluster *v1alpha1.EcsCluster) (result *v1alpha1.EcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ecsclustersResource, ecsCluster), &v1alpha1.EcsCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsCluster), err
}

// Update takes the representation of a ecsCluster and updates it. Returns the server's representation of the ecsCluster, and an error, if there is any.
func (c *FakeEcsClusters) Update(ecsCluster *v1alpha1.EcsCluster) (result *v1alpha1.EcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ecsclustersResource, ecsCluster), &v1alpha1.EcsCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEcsClusters) UpdateStatus(ecsCluster *v1alpha1.EcsCluster) (*v1alpha1.EcsCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ecsclustersResource, "status", ecsCluster), &v1alpha1.EcsCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsCluster), err
}

// Delete takes name of the ecsCluster and deletes it. Returns an error if one occurs.
func (c *FakeEcsClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ecsclustersResource, name), &v1alpha1.EcsCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEcsClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ecsclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EcsClusterList{})
	return err
}

// Patch applies the patch and returns the patched ecsCluster.
func (c *FakeEcsClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ecsclustersResource, name, pt, data, subresources...), &v1alpha1.EcsCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EcsCluster), err
}