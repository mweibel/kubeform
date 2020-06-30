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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachineDataDiskAttachments implements VirtualMachineDataDiskAttachmentInterface
type FakeVirtualMachineDataDiskAttachments struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var virtualmachinedatadiskattachmentsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "virtualmachinedatadiskattachments"}

var virtualmachinedatadiskattachmentsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "VirtualMachineDataDiskAttachment"}

// Get takes name of the virtualMachineDataDiskAttachment, and returns the corresponding virtualMachineDataDiskAttachment object, and an error if there is any.
func (c *FakeVirtualMachineDataDiskAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinedatadiskattachmentsResource, c.ns, name), &v1alpha1.VirtualMachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineDataDiskAttachment), err
}

// List takes label and field selectors, and returns the list of VirtualMachineDataDiskAttachments that match those selectors.
func (c *FakeVirtualMachineDataDiskAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineDataDiskAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinedatadiskattachmentsResource, virtualmachinedatadiskattachmentsKind, c.ns, opts), &v1alpha1.VirtualMachineDataDiskAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineDataDiskAttachmentList{ListMeta: obj.(*v1alpha1.VirtualMachineDataDiskAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineDataDiskAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineDataDiskAttachments.
func (c *FakeVirtualMachineDataDiskAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinedatadiskattachmentsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineDataDiskAttachment and creates it.  Returns the server's representation of the virtualMachineDataDiskAttachment, and an error, if there is any.
func (c *FakeVirtualMachineDataDiskAttachments) Create(ctx context.Context, virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinedatadiskattachmentsResource, c.ns, virtualMachineDataDiskAttachment), &v1alpha1.VirtualMachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineDataDiskAttachment), err
}

// Update takes the representation of a virtualMachineDataDiskAttachment and updates it. Returns the server's representation of the virtualMachineDataDiskAttachment, and an error, if there is any.
func (c *FakeVirtualMachineDataDiskAttachments) Update(ctx context.Context, virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinedatadiskattachmentsResource, c.ns, virtualMachineDataDiskAttachment), &v1alpha1.VirtualMachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineDataDiskAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineDataDiskAttachments) UpdateStatus(ctx context.Context, virtualMachineDataDiskAttachment *v1alpha1.VirtualMachineDataDiskAttachment, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineDataDiskAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinedatadiskattachmentsResource, "status", c.ns, virtualMachineDataDiskAttachment), &v1alpha1.VirtualMachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineDataDiskAttachment), err
}

// Delete takes name of the virtualMachineDataDiskAttachment and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineDataDiskAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinedatadiskattachmentsResource, c.ns, name), &v1alpha1.VirtualMachineDataDiskAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineDataDiskAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachinedatadiskattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineDataDiskAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineDataDiskAttachment.
func (c *FakeVirtualMachineDataDiskAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineDataDiskAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinedatadiskattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineDataDiskAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineDataDiskAttachment), err
}
