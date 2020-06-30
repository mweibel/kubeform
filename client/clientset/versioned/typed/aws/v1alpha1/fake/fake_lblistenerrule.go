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

// FakeLbListenerRules implements LbListenerRuleInterface
type FakeLbListenerRules struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lblistenerrulesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lblistenerrules"}

var lblistenerrulesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LbListenerRule"}

// Get takes name of the lbListenerRule, and returns the corresponding lbListenerRule object, and an error if there is any.
func (c *FakeLbListenerRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lblistenerrulesResource, c.ns, name), &v1alpha1.LbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// List takes label and field selectors, and returns the list of LbListenerRules that match those selectors.
func (c *FakeLbListenerRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LbListenerRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lblistenerrulesResource, lblistenerrulesKind, c.ns, opts), &v1alpha1.LbListenerRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbListenerRuleList{ListMeta: obj.(*v1alpha1.LbListenerRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbListenerRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbListenerRules.
func (c *FakeLbListenerRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lblistenerrulesResource, c.ns, opts))

}

// Create takes the representation of a lbListenerRule and creates it.  Returns the server's representation of the lbListenerRule, and an error, if there is any.
func (c *FakeLbListenerRules) Create(ctx context.Context, lbListenerRule *v1alpha1.LbListenerRule, opts v1.CreateOptions) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lblistenerrulesResource, c.ns, lbListenerRule), &v1alpha1.LbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// Update takes the representation of a lbListenerRule and updates it. Returns the server's representation of the lbListenerRule, and an error, if there is any.
func (c *FakeLbListenerRules) Update(ctx context.Context, lbListenerRule *v1alpha1.LbListenerRule, opts v1.UpdateOptions) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lblistenerrulesResource, c.ns, lbListenerRule), &v1alpha1.LbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbListenerRules) UpdateStatus(ctx context.Context, lbListenerRule *v1alpha1.LbListenerRule, opts v1.UpdateOptions) (*v1alpha1.LbListenerRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lblistenerrulesResource, "status", c.ns, lbListenerRule), &v1alpha1.LbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}

// Delete takes name of the lbListenerRule and deletes it. Returns an error if one occurs.
func (c *FakeLbListenerRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lblistenerrulesResource, c.ns, name), &v1alpha1.LbListenerRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbListenerRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lblistenerrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbListenerRuleList{})
	return err
}

// Patch applies the patch and returns the patched lbListenerRule.
func (c *FakeLbListenerRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lblistenerrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbListenerRule), err
}
