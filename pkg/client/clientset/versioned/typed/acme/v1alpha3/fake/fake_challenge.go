/*
Copyright 2020 The cert-manager Authors.

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

	v1alpha3 "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChallenges implements ChallengeInterface
type FakeChallenges struct {
	Fake *FakeAcmeV1alpha3
	ns   string
}

var challengesResource = schema.GroupVersionResource{Group: "acme.cert-manager.io", Version: "v1alpha3", Resource: "challenges"}

var challengesKind = schema.GroupVersionKind{Group: "acme.cert-manager.io", Version: "v1alpha3", Kind: "Challenge"}

// Get takes name of the challenge, and returns the corresponding challenge object, and an error if there is any.
func (c *FakeChallenges) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.Challenge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(challengesResource, c.ns, name), &v1alpha3.Challenge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.Challenge), err
}

// List takes label and field selectors, and returns the list of Challenges that match those selectors.
func (c *FakeChallenges) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.ChallengeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(challengesResource, challengesKind, c.ns, opts), &v1alpha3.ChallengeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.ChallengeList{ListMeta: obj.(*v1alpha3.ChallengeList).ListMeta}
	for _, item := range obj.(*v1alpha3.ChallengeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested challenges.
func (c *FakeChallenges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(challengesResource, c.ns, opts))

}

// Create takes the representation of a challenge and creates it.  Returns the server's representation of the challenge, and an error, if there is any.
func (c *FakeChallenges) Create(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.CreateOptions) (result *v1alpha3.Challenge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(challengesResource, c.ns, challenge), &v1alpha3.Challenge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.Challenge), err
}

// Update takes the representation of a challenge and updates it. Returns the server's representation of the challenge, and an error, if there is any.
func (c *FakeChallenges) Update(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (result *v1alpha3.Challenge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(challengesResource, c.ns, challenge), &v1alpha3.Challenge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.Challenge), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChallenges) UpdateStatus(ctx context.Context, challenge *v1alpha3.Challenge, opts v1.UpdateOptions) (*v1alpha3.Challenge, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(challengesResource, "status", c.ns, challenge), &v1alpha3.Challenge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.Challenge), err
}

// Delete takes name of the challenge and deletes it. Returns an error if one occurs.
func (c *FakeChallenges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(challengesResource, c.ns, name), &v1alpha3.Challenge{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChallenges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(challengesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.ChallengeList{})
	return err
}

// Patch applies the patch and returns the patched challenge.
func (c *FakeChallenges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.Challenge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(challengesResource, c.ns, name, pt, data, subresources...), &v1alpha3.Challenge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.Challenge), err
}
