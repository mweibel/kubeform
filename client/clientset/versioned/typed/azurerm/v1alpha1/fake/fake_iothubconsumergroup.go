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

// FakeIothubConsumerGroups implements IothubConsumerGroupInterface
type FakeIothubConsumerGroups struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var iothubconsumergroupsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "iothubconsumergroups"}

var iothubconsumergroupsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "IothubConsumerGroup"}

// Get takes name of the iothubConsumerGroup, and returns the corresponding iothubConsumerGroup object, and an error if there is any.
func (c *FakeIothubConsumerGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IothubConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iothubconsumergroupsResource, c.ns, name), &v1alpha1.IothubConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubConsumerGroup), err
}

// List takes label and field selectors, and returns the list of IothubConsumerGroups that match those selectors.
func (c *FakeIothubConsumerGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IothubConsumerGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iothubconsumergroupsResource, iothubconsumergroupsKind, c.ns, opts), &v1alpha1.IothubConsumerGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IothubConsumerGroupList{ListMeta: obj.(*v1alpha1.IothubConsumerGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.IothubConsumerGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iothubConsumerGroups.
func (c *FakeIothubConsumerGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iothubconsumergroupsResource, c.ns, opts))

}

// Create takes the representation of a iothubConsumerGroup and creates it.  Returns the server's representation of the iothubConsumerGroup, and an error, if there is any.
func (c *FakeIothubConsumerGroups) Create(ctx context.Context, iothubConsumerGroup *v1alpha1.IothubConsumerGroup, opts v1.CreateOptions) (result *v1alpha1.IothubConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iothubconsumergroupsResource, c.ns, iothubConsumerGroup), &v1alpha1.IothubConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubConsumerGroup), err
}

// Update takes the representation of a iothubConsumerGroup and updates it. Returns the server's representation of the iothubConsumerGroup, and an error, if there is any.
func (c *FakeIothubConsumerGroups) Update(ctx context.Context, iothubConsumerGroup *v1alpha1.IothubConsumerGroup, opts v1.UpdateOptions) (result *v1alpha1.IothubConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iothubconsumergroupsResource, c.ns, iothubConsumerGroup), &v1alpha1.IothubConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubConsumerGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIothubConsumerGroups) UpdateStatus(ctx context.Context, iothubConsumerGroup *v1alpha1.IothubConsumerGroup, opts v1.UpdateOptions) (*v1alpha1.IothubConsumerGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iothubconsumergroupsResource, "status", c.ns, iothubConsumerGroup), &v1alpha1.IothubConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubConsumerGroup), err
}

// Delete takes name of the iothubConsumerGroup and deletes it. Returns an error if one occurs.
func (c *FakeIothubConsumerGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iothubconsumergroupsResource, c.ns, name), &v1alpha1.IothubConsumerGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIothubConsumerGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iothubconsumergroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IothubConsumerGroupList{})
	return err
}

// Patch applies the patch and returns the patched iothubConsumerGroup.
func (c *FakeIothubConsumerGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IothubConsumerGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iothubconsumergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IothubConsumerGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IothubConsumerGroup), err
}
