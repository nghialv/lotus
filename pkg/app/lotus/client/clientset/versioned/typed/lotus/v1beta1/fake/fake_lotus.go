// Copyright (c) 2018 Lotus Load
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/lotusload/lotus/pkg/app/lotus/apis/lotus/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLotuses implements LotusInterface
type FakeLotuses struct {
	Fake *FakeLotusV1beta1
	ns   string
}

var lotusesResource = schema.GroupVersionResource{Group: "lotus.lotusload.com", Version: "v1beta1", Resource: "lotuses"}

var lotusesKind = schema.GroupVersionKind{Group: "lotus.lotusload.com", Version: "v1beta1", Kind: "Lotus"}

// Get takes name of the lotus, and returns the corresponding lotus object, and an error if there is any.
func (c *FakeLotuses) Get(name string, options v1.GetOptions) (result *v1beta1.Lotus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lotusesResource, c.ns, name), &v1beta1.Lotus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Lotus), err
}

// List takes label and field selectors, and returns the list of Lotuses that match those selectors.
func (c *FakeLotuses) List(opts v1.ListOptions) (result *v1beta1.LotusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lotusesResource, lotusesKind, c.ns, opts), &v1beta1.LotusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.LotusList{ListMeta: obj.(*v1beta1.LotusList).ListMeta}
	for _, item := range obj.(*v1beta1.LotusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lotuses.
func (c *FakeLotuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lotusesResource, c.ns, opts))

}

// Create takes the representation of a lotus and creates it.  Returns the server's representation of the lotus, and an error, if there is any.
func (c *FakeLotuses) Create(lotus *v1beta1.Lotus) (result *v1beta1.Lotus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lotusesResource, c.ns, lotus), &v1beta1.Lotus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Lotus), err
}

// Update takes the representation of a lotus and updates it. Returns the server's representation of the lotus, and an error, if there is any.
func (c *FakeLotuses) Update(lotus *v1beta1.Lotus) (result *v1beta1.Lotus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lotusesResource, c.ns, lotus), &v1beta1.Lotus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Lotus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLotuses) UpdateStatus(lotus *v1beta1.Lotus) (*v1beta1.Lotus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lotusesResource, "status", c.ns, lotus), &v1beta1.Lotus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Lotus), err
}

// Delete takes name of the lotus and deletes it. Returns an error if one occurs.
func (c *FakeLotuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lotusesResource, c.ns, name), &v1beta1.Lotus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLotuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lotusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.LotusList{})
	return err
}

// Patch applies the patch and returns the patched lotus.
func (c *FakeLotuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Lotus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lotusesResource, c.ns, name, pt, data, subresources...), &v1beta1.Lotus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Lotus), err
}
