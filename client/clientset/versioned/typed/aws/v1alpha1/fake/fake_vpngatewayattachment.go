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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpnGatewayAttachments implements VpnGatewayAttachmentInterface
type FakeVpnGatewayAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpngatewayattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpngatewayattachments"}

var vpngatewayattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpnGatewayAttachment"}

// Get takes name of the vpnGatewayAttachment, and returns the corresponding vpnGatewayAttachment object, and an error if there is any.
func (c *FakeVpnGatewayAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpngatewayattachmentsResource, c.ns, name), &v1alpha1.VpnGatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayAttachment), err
}

// List takes label and field selectors, and returns the list of VpnGatewayAttachments that match those selectors.
func (c *FakeVpnGatewayAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpnGatewayAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpngatewayattachmentsResource, vpngatewayattachmentsKind, c.ns, opts), &v1alpha1.VpnGatewayAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpnGatewayAttachmentList{ListMeta: obj.(*v1alpha1.VpnGatewayAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpnGatewayAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpnGatewayAttachments.
func (c *FakeVpnGatewayAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpngatewayattachmentsResource, c.ns, opts))

}

// Create takes the representation of a vpnGatewayAttachment and creates it.  Returns the server's representation of the vpnGatewayAttachment, and an error, if there is any.
func (c *FakeVpnGatewayAttachments) Create(ctx context.Context, vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment, opts v1.CreateOptions) (result *v1alpha1.VpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpngatewayattachmentsResource, c.ns, vpnGatewayAttachment), &v1alpha1.VpnGatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayAttachment), err
}

// Update takes the representation of a vpnGatewayAttachment and updates it. Returns the server's representation of the vpnGatewayAttachment, and an error, if there is any.
func (c *FakeVpnGatewayAttachments) Update(ctx context.Context, vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment, opts v1.UpdateOptions) (result *v1alpha1.VpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpngatewayattachmentsResource, c.ns, vpnGatewayAttachment), &v1alpha1.VpnGatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpnGatewayAttachments) UpdateStatus(ctx context.Context, vpnGatewayAttachment *v1alpha1.VpnGatewayAttachment, opts v1.UpdateOptions) (*v1alpha1.VpnGatewayAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpngatewayattachmentsResource, "status", c.ns, vpnGatewayAttachment), &v1alpha1.VpnGatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayAttachment), err
}

// Delete takes name of the vpnGatewayAttachment and deletes it. Returns an error if one occurs.
func (c *FakeVpnGatewayAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpngatewayattachmentsResource, c.ns, name), &v1alpha1.VpnGatewayAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpnGatewayAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpngatewayattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpnGatewayAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched vpnGatewayAttachment.
func (c *FakeVpnGatewayAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpnGatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpngatewayattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpnGatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayAttachment), err
}
