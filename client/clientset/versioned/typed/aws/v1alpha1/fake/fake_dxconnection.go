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

// FakeDxConnections implements DxConnectionInterface
type FakeDxConnections struct {
	Fake *FakeAwsV1alpha1
}

var dxconnectionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dxconnections"}

var dxconnectionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DxConnection"}

// Get takes name of the dxConnection, and returns the corresponding dxConnection object, and an error if there is any.
func (c *FakeDxConnections) Get(name string, options v1.GetOptions) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dxconnectionsResource, name), &v1alpha1.DxConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// List takes label and field selectors, and returns the list of DxConnections that match those selectors.
func (c *FakeDxConnections) List(opts v1.ListOptions) (result *v1alpha1.DxConnectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dxconnectionsResource, dxconnectionsKind, opts), &v1alpha1.DxConnectionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DxConnectionList{ListMeta: obj.(*v1alpha1.DxConnectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.DxConnectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dxConnections.
func (c *FakeDxConnections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dxconnectionsResource, opts))
}

// Create takes the representation of a dxConnection and creates it.  Returns the server's representation of the dxConnection, and an error, if there is any.
func (c *FakeDxConnections) Create(dxConnection *v1alpha1.DxConnection) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dxconnectionsResource, dxConnection), &v1alpha1.DxConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// Update takes the representation of a dxConnection and updates it. Returns the server's representation of the dxConnection, and an error, if there is any.
func (c *FakeDxConnections) Update(dxConnection *v1alpha1.DxConnection) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dxconnectionsResource, dxConnection), &v1alpha1.DxConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDxConnections) UpdateStatus(dxConnection *v1alpha1.DxConnection) (*v1alpha1.DxConnection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dxconnectionsResource, "status", dxConnection), &v1alpha1.DxConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}

// Delete takes name of the dxConnection and deletes it. Returns an error if one occurs.
func (c *FakeDxConnections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dxconnectionsResource, name), &v1alpha1.DxConnection{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDxConnections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dxconnectionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DxConnectionList{})
	return err
}

// Patch applies the patch and returns the patched dxConnection.
func (c *FakeDxConnections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DxConnection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dxconnectionsResource, name, pt, data, subresources...), &v1alpha1.DxConnection{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DxConnection), err
}