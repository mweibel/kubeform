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

// FakeSfnStateMachines implements SfnStateMachineInterface
type FakeSfnStateMachines struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sfnstatemachinesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sfnstatemachines"}

var sfnstatemachinesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SfnStateMachine"}

// Get takes name of the sfnStateMachine, and returns the corresponding sfnStateMachine object, and an error if there is any.
func (c *FakeSfnStateMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sfnstatemachinesResource, c.ns, name), &v1alpha1.SfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SfnStateMachine), err
}

// List takes label and field selectors, and returns the list of SfnStateMachines that match those selectors.
func (c *FakeSfnStateMachines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SfnStateMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sfnstatemachinesResource, sfnstatemachinesKind, c.ns, opts), &v1alpha1.SfnStateMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SfnStateMachineList{ListMeta: obj.(*v1alpha1.SfnStateMachineList).ListMeta}
	for _, item := range obj.(*v1alpha1.SfnStateMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sfnStateMachines.
func (c *FakeSfnStateMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sfnstatemachinesResource, c.ns, opts))

}

// Create takes the representation of a sfnStateMachine and creates it.  Returns the server's representation of the sfnStateMachine, and an error, if there is any.
func (c *FakeSfnStateMachines) Create(ctx context.Context, sfnStateMachine *v1alpha1.SfnStateMachine, opts v1.CreateOptions) (result *v1alpha1.SfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sfnstatemachinesResource, c.ns, sfnStateMachine), &v1alpha1.SfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SfnStateMachine), err
}

// Update takes the representation of a sfnStateMachine and updates it. Returns the server's representation of the sfnStateMachine, and an error, if there is any.
func (c *FakeSfnStateMachines) Update(ctx context.Context, sfnStateMachine *v1alpha1.SfnStateMachine, opts v1.UpdateOptions) (result *v1alpha1.SfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sfnstatemachinesResource, c.ns, sfnStateMachine), &v1alpha1.SfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SfnStateMachine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSfnStateMachines) UpdateStatus(ctx context.Context, sfnStateMachine *v1alpha1.SfnStateMachine, opts v1.UpdateOptions) (*v1alpha1.SfnStateMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sfnstatemachinesResource, "status", c.ns, sfnStateMachine), &v1alpha1.SfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SfnStateMachine), err
}

// Delete takes name of the sfnStateMachine and deletes it. Returns an error if one occurs.
func (c *FakeSfnStateMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sfnstatemachinesResource, c.ns, name), &v1alpha1.SfnStateMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSfnStateMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sfnstatemachinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SfnStateMachineList{})
	return err
}

// Patch applies the patch and returns the patched sfnStateMachine.
func (c *FakeSfnStateMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SfnStateMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sfnstatemachinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SfnStateMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SfnStateMachine), err
}
