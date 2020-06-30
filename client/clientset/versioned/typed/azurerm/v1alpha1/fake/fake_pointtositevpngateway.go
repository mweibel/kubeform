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

// FakePointToSiteVPNGateways implements PointToSiteVPNGatewayInterface
type FakePointToSiteVPNGateways struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var pointtositevpngatewaysResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "pointtositevpngateways"}

var pointtositevpngatewaysKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "PointToSiteVPNGateway"}

// Get takes name of the pointToSiteVPNGateway, and returns the corresponding pointToSiteVPNGateway object, and an error if there is any.
func (c *FakePointToSiteVPNGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PointToSiteVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pointtositevpngatewaysResource, c.ns, name), &v1alpha1.PointToSiteVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PointToSiteVPNGateway), err
}

// List takes label and field selectors, and returns the list of PointToSiteVPNGateways that match those selectors.
func (c *FakePointToSiteVPNGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PointToSiteVPNGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pointtositevpngatewaysResource, pointtositevpngatewaysKind, c.ns, opts), &v1alpha1.PointToSiteVPNGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PointToSiteVPNGatewayList{ListMeta: obj.(*v1alpha1.PointToSiteVPNGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.PointToSiteVPNGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pointToSiteVPNGateways.
func (c *FakePointToSiteVPNGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pointtositevpngatewaysResource, c.ns, opts))

}

// Create takes the representation of a pointToSiteVPNGateway and creates it.  Returns the server's representation of the pointToSiteVPNGateway, and an error, if there is any.
func (c *FakePointToSiteVPNGateways) Create(ctx context.Context, pointToSiteVPNGateway *v1alpha1.PointToSiteVPNGateway, opts v1.CreateOptions) (result *v1alpha1.PointToSiteVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pointtositevpngatewaysResource, c.ns, pointToSiteVPNGateway), &v1alpha1.PointToSiteVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PointToSiteVPNGateway), err
}

// Update takes the representation of a pointToSiteVPNGateway and updates it. Returns the server's representation of the pointToSiteVPNGateway, and an error, if there is any.
func (c *FakePointToSiteVPNGateways) Update(ctx context.Context, pointToSiteVPNGateway *v1alpha1.PointToSiteVPNGateway, opts v1.UpdateOptions) (result *v1alpha1.PointToSiteVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pointtositevpngatewaysResource, c.ns, pointToSiteVPNGateway), &v1alpha1.PointToSiteVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PointToSiteVPNGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePointToSiteVPNGateways) UpdateStatus(ctx context.Context, pointToSiteVPNGateway *v1alpha1.PointToSiteVPNGateway, opts v1.UpdateOptions) (*v1alpha1.PointToSiteVPNGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pointtositevpngatewaysResource, "status", c.ns, pointToSiteVPNGateway), &v1alpha1.PointToSiteVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PointToSiteVPNGateway), err
}

// Delete takes name of the pointToSiteVPNGateway and deletes it. Returns an error if one occurs.
func (c *FakePointToSiteVPNGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pointtositevpngatewaysResource, c.ns, name), &v1alpha1.PointToSiteVPNGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePointToSiteVPNGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pointtositevpngatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PointToSiteVPNGatewayList{})
	return err
}

// Patch applies the patch and returns the patched pointToSiteVPNGateway.
func (c *FakePointToSiteVPNGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PointToSiteVPNGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pointtositevpngatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.PointToSiteVPNGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PointToSiteVPNGateway), err
}
