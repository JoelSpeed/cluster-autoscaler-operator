/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2020 Red Hat, Inc.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/openshift/cluster-autoscaler-operator/pkg/apis/autoscaling/v1beta1"
	scheme "github.com/openshift/cluster-autoscaler-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineAutoscalersGetter has a method to return a MachineAutoscalerInterface.
// A group's client should implement this interface.
type MachineAutoscalersGetter interface {
	MachineAutoscalers(namespace string) MachineAutoscalerInterface
}

// MachineAutoscalerInterface has methods to work with MachineAutoscaler resources.
type MachineAutoscalerInterface interface {
	Create(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.CreateOptions) (*v1beta1.MachineAutoscaler, error)
	Update(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.UpdateOptions) (*v1beta1.MachineAutoscaler, error)
	UpdateStatus(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.UpdateOptions) (*v1beta1.MachineAutoscaler, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.MachineAutoscaler, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MachineAutoscalerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineAutoscaler, err error)
	MachineAutoscalerExpansion
}

// machineAutoscalers implements MachineAutoscalerInterface
type machineAutoscalers struct {
	client rest.Interface
	ns     string
}

// newMachineAutoscalers returns a MachineAutoscalers
func newMachineAutoscalers(c *AutoscalingV1beta1Client, namespace string) *machineAutoscalers {
	return &machineAutoscalers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineAutoscaler, and returns the corresponding machineAutoscaler object, and an error if there is any.
func (c *machineAutoscalers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MachineAutoscaler, err error) {
	result = &v1beta1.MachineAutoscaler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineAutoscalers that match those selectors.
func (c *machineAutoscalers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MachineAutoscalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MachineAutoscalerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineAutoscalers.
func (c *machineAutoscalers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machineautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineAutoscaler and creates it.  Returns the server's representation of the machineAutoscaler, and an error, if there is any.
func (c *machineAutoscalers) Create(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.CreateOptions) (result *v1beta1.MachineAutoscaler, err error) {
	result = &v1beta1.MachineAutoscaler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machineautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineAutoscaler and updates it. Returns the server's representation of the machineAutoscaler, and an error, if there is any.
func (c *machineAutoscalers) Update(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.UpdateOptions) (result *v1beta1.MachineAutoscaler, err error) {
	result = &v1beta1.MachineAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineautoscalers").
		Name(machineAutoscaler.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *machineAutoscalers) UpdateStatus(ctx context.Context, machineAutoscaler *v1beta1.MachineAutoscaler, opts v1.UpdateOptions) (result *v1beta1.MachineAutoscaler, err error) {
	result = &v1beta1.MachineAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineautoscalers").
		Name(machineAutoscaler.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineAutoscaler).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineAutoscaler and deletes it. Returns an error if one occurs.
func (c *machineAutoscalers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineautoscalers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineAutoscalers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineautoscalers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineAutoscaler.
func (c *machineAutoscalers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineAutoscaler, err error) {
	result = &v1beta1.MachineAutoscaler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machineautoscalers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
