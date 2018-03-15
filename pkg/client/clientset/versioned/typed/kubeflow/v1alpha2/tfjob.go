// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1alpha2

import (
	v1alpha2 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1alpha2"
	scheme "github.com/kubeflow/tf-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TFJobsGetter has a method to return a TFJobInterface.
// A group's client should implement this interface.
type TFJobsGetter interface {
	TFJobs(namespace string) TFJobInterface
}

// TFJobInterface has methods to work with TFJob resources.
type TFJobInterface interface {
	Create(*v1alpha2.TFJob) (*v1alpha2.TFJob, error)
	Update(*v1alpha2.TFJob) (*v1alpha2.TFJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.TFJob, error)
	List(opts v1.ListOptions) (*v1alpha2.TFJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.TFJob, err error)
	TFJobExpansion
}

// tFJobs implements TFJobInterface
type tFJobs struct {
	client rest.Interface
	ns     string
}

// newTFJobs returns a TFJobs
func newTFJobs(c *KubeflowV1alpha2Client, namespace string) *tFJobs {
	return &tFJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tFJob, and returns the corresponding tFJob object, and an error if there is any.
func (c *tFJobs) Get(name string, options v1.GetOptions) (result *v1alpha2.TFJob, err error) {
	result = &v1alpha2.TFJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TFJobs that match those selectors.
func (c *tFJobs) List(opts v1.ListOptions) (result *v1alpha2.TFJobList, err error) {
	result = &v1alpha2.TFJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tFJobs.
func (c *tFJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a tFJob and creates it.  Returns the server's representation of the tFJob, and an error, if there is any.
func (c *tFJobs) Create(tFJob *v1alpha2.TFJob) (result *v1alpha2.TFJob, err error) {
	result = &v1alpha2.TFJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tfjobs").
		Body(tFJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tFJob and updates it. Returns the server's representation of the tFJob, and an error, if there is any.
func (c *tFJobs) Update(tFJob *v1alpha2.TFJob) (result *v1alpha2.TFJob, err error) {
	result = &v1alpha2.TFJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(tFJob.Name).
		Body(tFJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the tFJob and deletes it. Returns an error if one occurs.
func (c *tFJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tFJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tFJob.
func (c *tFJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.TFJob, err error) {
	result = &v1alpha2.TFJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tfjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
