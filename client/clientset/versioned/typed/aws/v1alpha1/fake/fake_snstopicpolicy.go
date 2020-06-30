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

// FakeSnsTopicPolicies implements SnsTopicPolicyInterface
type FakeSnsTopicPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var snstopicpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "snstopicpolicies"}

var snstopicpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SnsTopicPolicy"}

// Get takes name of the snsTopicPolicy, and returns the corresponding snsTopicPolicy object, and an error if there is any.
func (c *FakeSnsTopicPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snstopicpoliciesResource, c.ns, name), &v1alpha1.SnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicPolicy), err
}

// List takes label and field selectors, and returns the list of SnsTopicPolicies that match those selectors.
func (c *FakeSnsTopicPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnsTopicPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snstopicpoliciesResource, snstopicpoliciesKind, c.ns, opts), &v1alpha1.SnsTopicPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnsTopicPolicyList{ListMeta: obj.(*v1alpha1.SnsTopicPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnsTopicPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snsTopicPolicies.
func (c *FakeSnsTopicPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snstopicpoliciesResource, c.ns, opts))

}

// Create takes the representation of a snsTopicPolicy and creates it.  Returns the server's representation of the snsTopicPolicy, and an error, if there is any.
func (c *FakeSnsTopicPolicies) Create(ctx context.Context, snsTopicPolicy *v1alpha1.SnsTopicPolicy, opts v1.CreateOptions) (result *v1alpha1.SnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snstopicpoliciesResource, c.ns, snsTopicPolicy), &v1alpha1.SnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicPolicy), err
}

// Update takes the representation of a snsTopicPolicy and updates it. Returns the server's representation of the snsTopicPolicy, and an error, if there is any.
func (c *FakeSnsTopicPolicies) Update(ctx context.Context, snsTopicPolicy *v1alpha1.SnsTopicPolicy, opts v1.UpdateOptions) (result *v1alpha1.SnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snstopicpoliciesResource, c.ns, snsTopicPolicy), &v1alpha1.SnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnsTopicPolicies) UpdateStatus(ctx context.Context, snsTopicPolicy *v1alpha1.SnsTopicPolicy, opts v1.UpdateOptions) (*v1alpha1.SnsTopicPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snstopicpoliciesResource, "status", c.ns, snsTopicPolicy), &v1alpha1.SnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicPolicy), err
}

// Delete takes name of the snsTopicPolicy and deletes it. Returns an error if one occurs.
func (c *FakeSnsTopicPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snstopicpoliciesResource, c.ns, name), &v1alpha1.SnsTopicPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnsTopicPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snstopicpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnsTopicPolicyList{})
	return err
}

// Patch applies the patch and returns the patched snsTopicPolicy.
func (c *FakeSnsTopicPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SnsTopicPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snstopicpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SnsTopicPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SnsTopicPolicy), err
}
