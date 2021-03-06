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

// FakeConfigOrganizationManagedRules implements ConfigOrganizationManagedRuleInterface
type FakeConfigOrganizationManagedRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var configorganizationmanagedrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "configorganizationmanagedrules"}

var configorganizationmanagedrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ConfigOrganizationManagedRule"}

// Get takes name of the configOrganizationManagedRule, and returns the corresponding configOrganizationManagedRule object, and an error if there is any.
func (c *FakeConfigOrganizationManagedRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigOrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configorganizationmanagedrulesResource, c.ns, name), &v1alpha1.ConfigOrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigOrganizationManagedRule), err
}

// List takes label and field selectors, and returns the list of ConfigOrganizationManagedRules that match those selectors.
func (c *FakeConfigOrganizationManagedRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigOrganizationManagedRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configorganizationmanagedrulesResource, configorganizationmanagedrulesKind, c.ns, opts), &v1alpha1.ConfigOrganizationManagedRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigOrganizationManagedRuleList{ListMeta: obj.(*v1alpha1.ConfigOrganizationManagedRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigOrganizationManagedRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configOrganizationManagedRules.
func (c *FakeConfigOrganizationManagedRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configorganizationmanagedrulesResource, c.ns, opts))

}

// Create takes the representation of a configOrganizationManagedRule and creates it.  Returns the server's representation of the configOrganizationManagedRule, and an error, if there is any.
func (c *FakeConfigOrganizationManagedRules) Create(ctx context.Context, configOrganizationManagedRule *v1alpha1.ConfigOrganizationManagedRule, opts v1.CreateOptions) (result *v1alpha1.ConfigOrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configorganizationmanagedrulesResource, c.ns, configOrganizationManagedRule), &v1alpha1.ConfigOrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigOrganizationManagedRule), err
}

// Update takes the representation of a configOrganizationManagedRule and updates it. Returns the server's representation of the configOrganizationManagedRule, and an error, if there is any.
func (c *FakeConfigOrganizationManagedRules) Update(ctx context.Context, configOrganizationManagedRule *v1alpha1.ConfigOrganizationManagedRule, opts v1.UpdateOptions) (result *v1alpha1.ConfigOrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configorganizationmanagedrulesResource, c.ns, configOrganizationManagedRule), &v1alpha1.ConfigOrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigOrganizationManagedRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigOrganizationManagedRules) UpdateStatus(ctx context.Context, configOrganizationManagedRule *v1alpha1.ConfigOrganizationManagedRule, opts v1.UpdateOptions) (*v1alpha1.ConfigOrganizationManagedRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configorganizationmanagedrulesResource, "status", c.ns, configOrganizationManagedRule), &v1alpha1.ConfigOrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigOrganizationManagedRule), err
}

// Delete takes name of the configOrganizationManagedRule and deletes it. Returns an error if one occurs.
func (c *FakeConfigOrganizationManagedRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configorganizationmanagedrulesResource, c.ns, name), &v1alpha1.ConfigOrganizationManagedRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigOrganizationManagedRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configorganizationmanagedrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigOrganizationManagedRuleList{})
	return err
}

// Patch applies the patch and returns the patched configOrganizationManagedRule.
func (c *FakeConfigOrganizationManagedRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigOrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configorganizationmanagedrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigOrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigOrganizationManagedRule), err
}
