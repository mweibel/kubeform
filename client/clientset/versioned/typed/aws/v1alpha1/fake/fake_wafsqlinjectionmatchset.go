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

// FakeWafSqlInjectionMatchSets implements WafSqlInjectionMatchSetInterface
type FakeWafSqlInjectionMatchSets struct {
	Fake *FakeAwsV1alpha1
}

var wafsqlinjectionmatchsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "wafsqlinjectionmatchsets"}

var wafsqlinjectionmatchsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "WafSqlInjectionMatchSet"}

// Get takes name of the wafSqlInjectionMatchSet, and returns the corresponding wafSqlInjectionMatchSet object, and an error if there is any.
func (c *FakeWafSqlInjectionMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(wafsqlinjectionmatchsetsResource, name), &v1alpha1.WafSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafSqlInjectionMatchSet), err
}

// List takes label and field selectors, and returns the list of WafSqlInjectionMatchSets that match those selectors.
func (c *FakeWafSqlInjectionMatchSets) List(opts v1.ListOptions) (result *v1alpha1.WafSqlInjectionMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(wafsqlinjectionmatchsetsResource, wafsqlinjectionmatchsetsKind, opts), &v1alpha1.WafSqlInjectionMatchSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WafSqlInjectionMatchSetList{ListMeta: obj.(*v1alpha1.WafSqlInjectionMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.WafSqlInjectionMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested wafSqlInjectionMatchSets.
func (c *FakeWafSqlInjectionMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(wafsqlinjectionmatchsetsResource, opts))
}

// Create takes the representation of a wafSqlInjectionMatchSet and creates it.  Returns the server's representation of the wafSqlInjectionMatchSet, and an error, if there is any.
func (c *FakeWafSqlInjectionMatchSets) Create(wafSqlInjectionMatchSet *v1alpha1.WafSqlInjectionMatchSet) (result *v1alpha1.WafSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(wafsqlinjectionmatchsetsResource, wafSqlInjectionMatchSet), &v1alpha1.WafSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafSqlInjectionMatchSet), err
}

// Update takes the representation of a wafSqlInjectionMatchSet and updates it. Returns the server's representation of the wafSqlInjectionMatchSet, and an error, if there is any.
func (c *FakeWafSqlInjectionMatchSets) Update(wafSqlInjectionMatchSet *v1alpha1.WafSqlInjectionMatchSet) (result *v1alpha1.WafSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(wafsqlinjectionmatchsetsResource, wafSqlInjectionMatchSet), &v1alpha1.WafSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafSqlInjectionMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWafSqlInjectionMatchSets) UpdateStatus(wafSqlInjectionMatchSet *v1alpha1.WafSqlInjectionMatchSet) (*v1alpha1.WafSqlInjectionMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(wafsqlinjectionmatchsetsResource, "status", wafSqlInjectionMatchSet), &v1alpha1.WafSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafSqlInjectionMatchSet), err
}

// Delete takes name of the wafSqlInjectionMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeWafSqlInjectionMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(wafsqlinjectionmatchsetsResource, name), &v1alpha1.WafSqlInjectionMatchSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWafSqlInjectionMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(wafsqlinjectionmatchsetsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WafSqlInjectionMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched wafSqlInjectionMatchSet.
func (c *FakeWafSqlInjectionMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafSqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(wafsqlinjectionmatchsetsResource, name, pt, data, subresources...), &v1alpha1.WafSqlInjectionMatchSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WafSqlInjectionMatchSet), err
}