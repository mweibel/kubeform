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

// FakeSsmMaintenanceWindowTasks implements SsmMaintenanceWindowTaskInterface
type FakeSsmMaintenanceWindowTasks struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ssmmaintenancewindowtasksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmmaintenancewindowtasks"}

var ssmmaintenancewindowtasksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmMaintenanceWindowTask"}

// Get takes name of the ssmMaintenanceWindowTask, and returns the corresponding ssmMaintenanceWindowTask object, and an error if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssmmaintenancewindowtasksResource, c.ns, name), &v1alpha1.SsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// List takes label and field selectors, and returns the list of SsmMaintenanceWindowTasks that match those selectors.
func (c *FakeSsmMaintenanceWindowTasks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SsmMaintenanceWindowTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssmmaintenancewindowtasksResource, ssmmaintenancewindowtasksKind, c.ns, opts), &v1alpha1.SsmMaintenanceWindowTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsmMaintenanceWindowTaskList{ListMeta: obj.(*v1alpha1.SsmMaintenanceWindowTaskList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsmMaintenanceWindowTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssmMaintenanceWindowTasks.
func (c *FakeSsmMaintenanceWindowTasks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssmmaintenancewindowtasksResource, c.ns, opts))

}

// Create takes the representation of a ssmMaintenanceWindowTask and creates it.  Returns the server's representation of the ssmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Create(ctx context.Context, ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask, opts v1.CreateOptions) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssmmaintenancewindowtasksResource, c.ns, ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// Update takes the representation of a ssmMaintenanceWindowTask and updates it. Returns the server's representation of the ssmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Update(ctx context.Context, ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask, opts v1.UpdateOptions) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssmmaintenancewindowtasksResource, c.ns, ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmMaintenanceWindowTasks) UpdateStatus(ctx context.Context, ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask, opts v1.UpdateOptions) (*v1alpha1.SsmMaintenanceWindowTask, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssmmaintenancewindowtasksResource, "status", c.ns, ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// Delete takes name of the ssmMaintenanceWindowTask and deletes it. Returns an error if one occurs.
func (c *FakeSsmMaintenanceWindowTasks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssmmaintenancewindowtasksResource, c.ns, name), &v1alpha1.SsmMaintenanceWindowTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmMaintenanceWindowTasks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssmmaintenancewindowtasksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmMaintenanceWindowTaskList{})
	return err
}

// Patch applies the patch and returns the patched ssmMaintenanceWindowTask.
func (c *FakeSsmMaintenanceWindowTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssmmaintenancewindowtasksResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsmMaintenanceWindowTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}
