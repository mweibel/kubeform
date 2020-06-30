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

// FakeBackupProtectedFileShares implements BackupProtectedFileShareInterface
type FakeBackupProtectedFileShares struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var backupprotectedfilesharesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "backupprotectedfileshares"}

var backupprotectedfilesharesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "BackupProtectedFileShare"}

// Get takes name of the backupProtectedFileShare, and returns the corresponding backupProtectedFileShare object, and an error if there is any.
func (c *FakeBackupProtectedFileShares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupProtectedFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupprotectedfilesharesResource, c.ns, name), &v1alpha1.BackupProtectedFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedFileShare), err
}

// List takes label and field selectors, and returns the list of BackupProtectedFileShares that match those selectors.
func (c *FakeBackupProtectedFileShares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupProtectedFileShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupprotectedfilesharesResource, backupprotectedfilesharesKind, c.ns, opts), &v1alpha1.BackupProtectedFileShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupProtectedFileShareList{ListMeta: obj.(*v1alpha1.BackupProtectedFileShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupProtectedFileShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupProtectedFileShares.
func (c *FakeBackupProtectedFileShares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupprotectedfilesharesResource, c.ns, opts))

}

// Create takes the representation of a backupProtectedFileShare and creates it.  Returns the server's representation of the backupProtectedFileShare, and an error, if there is any.
func (c *FakeBackupProtectedFileShares) Create(ctx context.Context, backupProtectedFileShare *v1alpha1.BackupProtectedFileShare, opts v1.CreateOptions) (result *v1alpha1.BackupProtectedFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupprotectedfilesharesResource, c.ns, backupProtectedFileShare), &v1alpha1.BackupProtectedFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedFileShare), err
}

// Update takes the representation of a backupProtectedFileShare and updates it. Returns the server's representation of the backupProtectedFileShare, and an error, if there is any.
func (c *FakeBackupProtectedFileShares) Update(ctx context.Context, backupProtectedFileShare *v1alpha1.BackupProtectedFileShare, opts v1.UpdateOptions) (result *v1alpha1.BackupProtectedFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupprotectedfilesharesResource, c.ns, backupProtectedFileShare), &v1alpha1.BackupProtectedFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedFileShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupProtectedFileShares) UpdateStatus(ctx context.Context, backupProtectedFileShare *v1alpha1.BackupProtectedFileShare, opts v1.UpdateOptions) (*v1alpha1.BackupProtectedFileShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupprotectedfilesharesResource, "status", c.ns, backupProtectedFileShare), &v1alpha1.BackupProtectedFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedFileShare), err
}

// Delete takes name of the backupProtectedFileShare and deletes it. Returns an error if one occurs.
func (c *FakeBackupProtectedFileShares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupprotectedfilesharesResource, c.ns, name), &v1alpha1.BackupProtectedFileShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupProtectedFileShares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupprotectedfilesharesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupProtectedFileShareList{})
	return err
}

// Patch applies the patch and returns the patched backupProtectedFileShare.
func (c *FakeBackupProtectedFileShares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupProtectedFileShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupprotectedfilesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupProtectedFileShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupProtectedFileShare), err
}