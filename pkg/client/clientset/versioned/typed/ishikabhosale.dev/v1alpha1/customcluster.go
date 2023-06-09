/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/ishikabhosale-cs/customcluster/pkg/apis/ishikabhosale.dev/v1alpha1"
	ishikabhosaledevv1alpha1 "github.com/ishikabhosale-cs/customcluster/pkg/client/applyconfiguration/ishikabhosale.dev/v1alpha1"
	scheme "github.com/ishikabhosale-cs/customcluster/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CustomclustersGetter has a method to return a CustomclusterInterface.
// A group's client should implement this interface.
type CustomclustersGetter interface {
	Customclusters(namespace string) CustomclusterInterface
}

// CustomclusterInterface has methods to work with Customcluster resources.
type CustomclusterInterface interface {
	Create(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.CreateOptions) (*v1alpha1.Customcluster, error)
	Update(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.UpdateOptions) (*v1alpha1.Customcluster, error)
	UpdateStatus(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.UpdateOptions) (*v1alpha1.Customcluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Customcluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CustomclusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Customcluster, err error)
	Apply(ctx context.Context, customcluster *ishikabhosaledevv1alpha1.CustomclusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Customcluster, err error)
	ApplyStatus(ctx context.Context, customcluster *ishikabhosaledevv1alpha1.CustomclusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Customcluster, err error)
	CustomclusterExpansion
}

// customclusters implements CustomclusterInterface
type customclusters struct {
	client rest.Interface
	ns     string
}

// newCustomclusters returns a Customclusters
func newCustomclusters(c *IshikabhosaleV1alpha1Client, namespace string) *customclusters {
	return &customclusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the customcluster, and returns the corresponding customcluster object, and an error if there is any.
func (c *customclusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Customcluster, err error) {
	result = &v1alpha1.Customcluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Customclusters that match those selectors.
func (c *customclusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomclusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CustomclusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customclusters.
func (c *customclusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("customclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a customcluster and creates it.  Returns the server's representation of the customcluster, and an error, if there is any.
func (c *customclusters) Create(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.CreateOptions) (result *v1alpha1.Customcluster, err error) {
	result = &v1alpha1.Customcluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("customclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customcluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a customcluster and updates it. Returns the server's representation of the customcluster, and an error, if there is any.
func (c *customclusters) Update(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.UpdateOptions) (result *v1alpha1.Customcluster, err error) {
	result = &v1alpha1.Customcluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("customclusters").
		Name(customcluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customcluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *customclusters) UpdateStatus(ctx context.Context, customcluster *v1alpha1.Customcluster, opts v1.UpdateOptions) (result *v1alpha1.Customcluster, err error) {
	result = &v1alpha1.Customcluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("customclusters").
		Name(customcluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customcluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the customcluster and deletes it. Returns an error if one occurs.
func (c *customclusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customclusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched customcluster.
func (c *customclusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Customcluster, err error) {
	result = &v1alpha1.Customcluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("customclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied customcluster.
func (c *customclusters) Apply(ctx context.Context, customcluster *ishikabhosaledevv1alpha1.CustomclusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Customcluster, err error) {
	if customcluster == nil {
		return nil, fmt.Errorf("customcluster provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(customcluster)
	if err != nil {
		return nil, err
	}
	name := customcluster.Name
	if name == nil {
		return nil, fmt.Errorf("customcluster.Name must be provided to Apply")
	}
	result = &v1alpha1.Customcluster{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("customclusters").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *customclusters) ApplyStatus(ctx context.Context, customcluster *ishikabhosaledevv1alpha1.CustomclusterApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Customcluster, err error) {
	if customcluster == nil {
		return nil, fmt.Errorf("customcluster provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(customcluster)
	if err != nil {
		return nil, err
	}

	name := customcluster.Name
	if name == nil {
		return nil, fmt.Errorf("customcluster.Name must be provided to Apply")
	}

	result = &v1alpha1.Customcluster{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("customclusters").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
