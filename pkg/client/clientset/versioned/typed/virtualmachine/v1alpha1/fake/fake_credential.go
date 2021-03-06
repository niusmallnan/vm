/*
Copyright 2018 Rancher Labs, Inc.

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
package fake

import (
	v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentials implements CredentialInterface
type FakeCredentials struct {
	Fake *FakeVirtualmachineV1alpha1
}

var credentialsResource = schema.GroupVersionResource{Group: "virtualmachine.rancher.com", Version: "v1alpha1", Resource: "credentials"}

var credentialsKind = schema.GroupVersionKind{Group: "virtualmachine.rancher.com", Version: "v1alpha1", Kind: "Credential"}

// Get takes name of the credential, and returns the corresponding credential object, and an error if there is any.
func (c *FakeCredentials) Get(name string, options v1.GetOptions) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(credentialsResource, name), &v1alpha1.Credential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// List takes label and field selectors, and returns the list of Credentials that match those selectors.
func (c *FakeCredentials) List(opts v1.ListOptions) (result *v1alpha1.CredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(credentialsResource, credentialsKind, opts), &v1alpha1.CredentialList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialList{}
	for _, item := range obj.(*v1alpha1.CredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentials.
func (c *FakeCredentials) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(credentialsResource, opts))
}

// Create takes the representation of a credential and creates it.  Returns the server's representation of the credential, and an error, if there is any.
func (c *FakeCredentials) Create(credential *v1alpha1.Credential) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(credentialsResource, credential), &v1alpha1.Credential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// Update takes the representation of a credential and updates it. Returns the server's representation of the credential, and an error, if there is any.
func (c *FakeCredentials) Update(credential *v1alpha1.Credential) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(credentialsResource, credential), &v1alpha1.Credential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// Delete takes name of the credential and deletes it. Returns an error if one occurs.
func (c *FakeCredentials) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(credentialsResource, name), &v1alpha1.Credential{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentials) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(credentialsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialList{})
	return err
}

// Patch applies the patch and returns the patched credential.
func (c *FakeCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(credentialsResource, name, data, subresources...), &v1alpha1.Credential{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}
