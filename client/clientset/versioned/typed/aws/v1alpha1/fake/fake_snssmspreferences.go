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

// FakeSnsSmsPreferenceses implements SnsSmsPreferencesInterface
type FakeSnsSmsPreferenceses struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var snssmspreferencesesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "snssmspreferenceses"}

var snssmspreferencesesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SnsSmsPreferences"}

// Get takes name of the snsSmsPreferences, and returns the corresponding snsSmsPreferences object, and an error if there is any.
func (c *FakeSnsSmsPreferenceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SnsSmsPreferences, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snssmspreferencesesResource, c.ns, name), &v1alpha1.SnsSmsPreferences{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsSmsPreferences), err
}

// List takes label and field selectors, and returns the list of SnsSmsPreferenceses that match those selectors.
func (c *FakeSnsSmsPreferenceses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnsSmsPreferencesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snssmspreferencesesResource, snssmspreferencesesKind, c.ns, opts), &v1alpha1.SnsSmsPreferencesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnsSmsPreferencesList{ListMeta: obj.(*v1alpha1.SnsSmsPreferencesList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnsSmsPreferencesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snsSmsPreferenceses.
func (c *FakeSnsSmsPreferenceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snssmspreferencesesResource, c.ns, opts))

}

// Create takes the representation of a snsSmsPreferences and creates it.  Returns the server's representation of the snsSmsPreferences, and an error, if there is any.
func (c *FakeSnsSmsPreferenceses) Create(ctx context.Context, snsSmsPreferences *v1alpha1.SnsSmsPreferences, opts v1.CreateOptions) (result *v1alpha1.SnsSmsPreferences, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snssmspreferencesesResource, c.ns, snsSmsPreferences), &v1alpha1.SnsSmsPreferences{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsSmsPreferences), err
}

// Update takes the representation of a snsSmsPreferences and updates it. Returns the server's representation of the snsSmsPreferences, and an error, if there is any.
func (c *FakeSnsSmsPreferenceses) Update(ctx context.Context, snsSmsPreferences *v1alpha1.SnsSmsPreferences, opts v1.UpdateOptions) (result *v1alpha1.SnsSmsPreferences, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snssmspreferencesesResource, c.ns, snsSmsPreferences), &v1alpha1.SnsSmsPreferences{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsSmsPreferences), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnsSmsPreferenceses) UpdateStatus(ctx context.Context, snsSmsPreferences *v1alpha1.SnsSmsPreferences, opts v1.UpdateOptions) (*v1alpha1.SnsSmsPreferences, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snssmspreferencesesResource, "status", c.ns, snsSmsPreferences), &v1alpha1.SnsSmsPreferences{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsSmsPreferences), err
}

// Delete takes name of the snsSmsPreferences and deletes it. Returns an error if one occurs.
func (c *FakeSnsSmsPreferenceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snssmspreferencesesResource, c.ns, name), &v1alpha1.SnsSmsPreferences{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnsSmsPreferenceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snssmspreferencesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnsSmsPreferencesList{})
	return err
}

// Patch applies the patch and returns the patched snsSmsPreferences.
func (c *FakeSnsSmsPreferenceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SnsSmsPreferences, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snssmspreferencesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SnsSmsPreferences{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsSmsPreferences), err
}
