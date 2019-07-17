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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakePubsubSubscriptionIamBindings implements PubsubSubscriptionIamBindingInterface
type FakePubsubSubscriptionIamBindings struct {
	Fake *FakeGoogleV1alpha1
}

var pubsubsubscriptioniambindingsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "pubsubsubscriptioniambindings"}

var pubsubsubscriptioniambindingsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "PubsubSubscriptionIamBinding"}

// Get takes name of the pubsubSubscriptionIamBinding, and returns the corresponding pubsubSubscriptionIamBinding object, and an error if there is any.
func (c *FakePubsubSubscriptionIamBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.PubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pubsubsubscriptioniambindingsResource, name), &v1alpha1.PubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), err
}

// List takes label and field selectors, and returns the list of PubsubSubscriptionIamBindings that match those selectors.
func (c *FakePubsubSubscriptionIamBindings) List(opts v1.ListOptions) (result *v1alpha1.PubsubSubscriptionIamBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pubsubsubscriptioniambindingsResource, pubsubsubscriptioniambindingsKind, opts), &v1alpha1.PubsubSubscriptionIamBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PubsubSubscriptionIamBindingList{ListMeta: obj.(*v1alpha1.PubsubSubscriptionIamBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.PubsubSubscriptionIamBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubsubSubscriptionIamBindings.
func (c *FakePubsubSubscriptionIamBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pubsubsubscriptioniambindingsResource, opts))
}

// Create takes the representation of a pubsubSubscriptionIamBinding and creates it.  Returns the server's representation of the pubsubSubscriptionIamBinding, and an error, if there is any.
func (c *FakePubsubSubscriptionIamBindings) Create(pubsubSubscriptionIamBinding *v1alpha1.PubsubSubscriptionIamBinding) (result *v1alpha1.PubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pubsubsubscriptioniambindingsResource, pubsubSubscriptionIamBinding), &v1alpha1.PubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), err
}

// Update takes the representation of a pubsubSubscriptionIamBinding and updates it. Returns the server's representation of the pubsubSubscriptionIamBinding, and an error, if there is any.
func (c *FakePubsubSubscriptionIamBindings) Update(pubsubSubscriptionIamBinding *v1alpha1.PubsubSubscriptionIamBinding) (result *v1alpha1.PubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pubsubsubscriptioniambindingsResource, pubsubSubscriptionIamBinding), &v1alpha1.PubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubsubSubscriptionIamBindings) UpdateStatus(pubsubSubscriptionIamBinding *v1alpha1.PubsubSubscriptionIamBinding) (*v1alpha1.PubsubSubscriptionIamBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pubsubsubscriptioniambindingsResource, "status", pubsubSubscriptionIamBinding), &v1alpha1.PubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), err
}

// Delete takes name of the pubsubSubscriptionIamBinding and deletes it. Returns an error if one occurs.
func (c *FakePubsubSubscriptionIamBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pubsubsubscriptioniambindingsResource, name), &v1alpha1.PubsubSubscriptionIamBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubsubSubscriptionIamBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pubsubsubscriptioniambindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PubsubSubscriptionIamBindingList{})
	return err
}

// Patch applies the patch and returns the patched pubsubSubscriptionIamBinding.
func (c *FakePubsubSubscriptionIamBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pubsubsubscriptioniambindingsResource, name, pt, data, subresources...), &v1alpha1.PubsubSubscriptionIamBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamBinding), err
}