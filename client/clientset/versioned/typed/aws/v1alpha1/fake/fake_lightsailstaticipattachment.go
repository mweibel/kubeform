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

// FakeLightsailStaticIPAttachments implements LightsailStaticIPAttachmentInterface
type FakeLightsailStaticIPAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lightsailstaticipattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lightsailstaticipattachments"}

var lightsailstaticipattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LightsailStaticIPAttachment"}

// Get takes name of the lightsailStaticIPAttachment, and returns the corresponding lightsailStaticIPAttachment object, and an error if there is any.
func (c *FakeLightsailStaticIPAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LightsailStaticIPAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lightsailstaticipattachmentsResource, c.ns, name), &v1alpha1.LightsailStaticIPAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIPAttachment), err
}

// List takes label and field selectors, and returns the list of LightsailStaticIPAttachments that match those selectors.
func (c *FakeLightsailStaticIPAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LightsailStaticIPAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lightsailstaticipattachmentsResource, lightsailstaticipattachmentsKind, c.ns, opts), &v1alpha1.LightsailStaticIPAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LightsailStaticIPAttachmentList{ListMeta: obj.(*v1alpha1.LightsailStaticIPAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.LightsailStaticIPAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lightsailStaticIPAttachments.
func (c *FakeLightsailStaticIPAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lightsailstaticipattachmentsResource, c.ns, opts))

}

// Create takes the representation of a lightsailStaticIPAttachment and creates it.  Returns the server's representation of the lightsailStaticIPAttachment, and an error, if there is any.
func (c *FakeLightsailStaticIPAttachments) Create(ctx context.Context, lightsailStaticIPAttachment *v1alpha1.LightsailStaticIPAttachment, opts v1.CreateOptions) (result *v1alpha1.LightsailStaticIPAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lightsailstaticipattachmentsResource, c.ns, lightsailStaticIPAttachment), &v1alpha1.LightsailStaticIPAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIPAttachment), err
}

// Update takes the representation of a lightsailStaticIPAttachment and updates it. Returns the server's representation of the lightsailStaticIPAttachment, and an error, if there is any.
func (c *FakeLightsailStaticIPAttachments) Update(ctx context.Context, lightsailStaticIPAttachment *v1alpha1.LightsailStaticIPAttachment, opts v1.UpdateOptions) (result *v1alpha1.LightsailStaticIPAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lightsailstaticipattachmentsResource, c.ns, lightsailStaticIPAttachment), &v1alpha1.LightsailStaticIPAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIPAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLightsailStaticIPAttachments) UpdateStatus(ctx context.Context, lightsailStaticIPAttachment *v1alpha1.LightsailStaticIPAttachment, opts v1.UpdateOptions) (*v1alpha1.LightsailStaticIPAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lightsailstaticipattachmentsResource, "status", c.ns, lightsailStaticIPAttachment), &v1alpha1.LightsailStaticIPAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIPAttachment), err
}

// Delete takes name of the lightsailStaticIPAttachment and deletes it. Returns an error if one occurs.
func (c *FakeLightsailStaticIPAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lightsailstaticipattachmentsResource, c.ns, name), &v1alpha1.LightsailStaticIPAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLightsailStaticIPAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lightsailstaticipattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LightsailStaticIPAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched lightsailStaticIPAttachment.
func (c *FakeLightsailStaticIPAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LightsailStaticIPAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lightsailstaticipattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LightsailStaticIPAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIPAttachment), err
}
