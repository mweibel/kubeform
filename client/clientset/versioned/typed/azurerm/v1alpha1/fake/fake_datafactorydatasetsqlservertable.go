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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeDataFactoryDatasetSqlServerTables implements DataFactoryDatasetSqlServerTableInterface
type FakeDataFactoryDatasetSqlServerTables struct {
	Fake *FakeAzurermV1alpha1
}

var datafactorydatasetsqlservertablesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactorydatasetsqlservertables"}

var datafactorydatasetsqlservertablesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactoryDatasetSqlServerTable"}

// Get takes name of the dataFactoryDatasetSqlServerTable, and returns the corresponding dataFactoryDatasetSqlServerTable object, and an error if there is any.
func (c *FakeDataFactoryDatasetSqlServerTables) Get(name string, options v1.GetOptions) (result *v1alpha1.DataFactoryDatasetSqlServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(datafactorydatasetsqlservertablesResource, name), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSqlServerTable), err
}

// List takes label and field selectors, and returns the list of DataFactoryDatasetSqlServerTables that match those selectors.
func (c *FakeDataFactoryDatasetSqlServerTables) List(opts v1.ListOptions) (result *v1alpha1.DataFactoryDatasetSqlServerTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(datafactorydatasetsqlservertablesResource, datafactorydatasetsqlservertablesKind, opts), &v1alpha1.DataFactoryDatasetSqlServerTableList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryDatasetSqlServerTableList{ListMeta: obj.(*v1alpha1.DataFactoryDatasetSqlServerTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryDatasetSqlServerTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactoryDatasetSqlServerTables.
func (c *FakeDataFactoryDatasetSqlServerTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(datafactorydatasetsqlservertablesResource, opts))
}

// Create takes the representation of a dataFactoryDatasetSqlServerTable and creates it.  Returns the server's representation of the dataFactoryDatasetSqlServerTable, and an error, if there is any.
func (c *FakeDataFactoryDatasetSqlServerTables) Create(dataFactoryDatasetSqlServerTable *v1alpha1.DataFactoryDatasetSqlServerTable) (result *v1alpha1.DataFactoryDatasetSqlServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(datafactorydatasetsqlservertablesResource, dataFactoryDatasetSqlServerTable), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSqlServerTable), err
}

// Update takes the representation of a dataFactoryDatasetSqlServerTable and updates it. Returns the server's representation of the dataFactoryDatasetSqlServerTable, and an error, if there is any.
func (c *FakeDataFactoryDatasetSqlServerTables) Update(dataFactoryDatasetSqlServerTable *v1alpha1.DataFactoryDatasetSqlServerTable) (result *v1alpha1.DataFactoryDatasetSqlServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(datafactorydatasetsqlservertablesResource, dataFactoryDatasetSqlServerTable), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSqlServerTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactoryDatasetSqlServerTables) UpdateStatus(dataFactoryDatasetSqlServerTable *v1alpha1.DataFactoryDatasetSqlServerTable) (*v1alpha1.DataFactoryDatasetSqlServerTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(datafactorydatasetsqlservertablesResource, "status", dataFactoryDatasetSqlServerTable), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSqlServerTable), err
}

// Delete takes name of the dataFactoryDatasetSqlServerTable and deletes it. Returns an error if one occurs.
func (c *FakeDataFactoryDatasetSqlServerTables) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(datafactorydatasetsqlservertablesResource, name), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactoryDatasetSqlServerTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(datafactorydatasetsqlservertablesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryDatasetSqlServerTableList{})
	return err
}

// Patch applies the patch and returns the patched dataFactoryDatasetSqlServerTable.
func (c *FakeDataFactoryDatasetSqlServerTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataFactoryDatasetSqlServerTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(datafactorydatasetsqlservertablesResource, name, pt, data, subresources...), &v1alpha1.DataFactoryDatasetSqlServerTable{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactoryDatasetSqlServerTable), err
}