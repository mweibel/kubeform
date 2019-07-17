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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeSsmMaintenanceWindowTasks implements SsmMaintenanceWindowTaskInterface
type FakeSsmMaintenanceWindowTasks struct {
	Fake *FakeAwsV1alpha1
}

var ssmmaintenancewindowtasksResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmmaintenancewindowtasks"}

var ssmmaintenancewindowtasksKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmMaintenanceWindowTask"}

// Get takes name of the ssmMaintenanceWindowTask, and returns the corresponding ssmMaintenanceWindowTask object, and an error if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ssmmaintenancewindowtasksResource, name), &v1alpha1.SsmMaintenanceWindowTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// List takes label and field selectors, and returns the list of SsmMaintenanceWindowTasks that match those selectors.
func (c *FakeSsmMaintenanceWindowTasks) List(opts v1.ListOptions) (result *v1alpha1.SsmMaintenanceWindowTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ssmmaintenancewindowtasksResource, ssmmaintenancewindowtasksKind, opts), &v1alpha1.SsmMaintenanceWindowTaskList{})
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
func (c *FakeSsmMaintenanceWindowTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ssmmaintenancewindowtasksResource, opts))
}

// Create takes the representation of a ssmMaintenanceWindowTask and creates it.  Returns the server's representation of the ssmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Create(ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ssmmaintenancewindowtasksResource, ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// Update takes the representation of a ssmMaintenanceWindowTask and updates it. Returns the server's representation of the ssmMaintenanceWindowTask, and an error, if there is any.
func (c *FakeSsmMaintenanceWindowTasks) Update(ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ssmmaintenancewindowtasksResource, ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmMaintenanceWindowTasks) UpdateStatus(ssmMaintenanceWindowTask *v1alpha1.SsmMaintenanceWindowTask) (*v1alpha1.SsmMaintenanceWindowTask, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ssmmaintenancewindowtasksResource, "status", ssmMaintenanceWindowTask), &v1alpha1.SsmMaintenanceWindowTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}

// Delete takes name of the ssmMaintenanceWindowTask and deletes it. Returns an error if one occurs.
func (c *FakeSsmMaintenanceWindowTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ssmmaintenancewindowtasksResource, name), &v1alpha1.SsmMaintenanceWindowTask{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmMaintenanceWindowTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ssmmaintenancewindowtasksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmMaintenanceWindowTaskList{})
	return err
}

// Patch applies the patch and returns the patched ssmMaintenanceWindowTask.
func (c *FakeSsmMaintenanceWindowTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmMaintenanceWindowTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ssmmaintenancewindowtasksResource, name, pt, data, subresources...), &v1alpha1.SsmMaintenanceWindowTask{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindowTask), err
}