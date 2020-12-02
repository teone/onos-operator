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

package v1beta1

import (
	"github.com/onosproject/onos-operator/pkg/apis/clientset/versioned/scheme"
	v1beta1 "github.com/onosproject/onos-operator/pkg/apis/topo/v1beta1"
	rest "k8s.io/client-go/rest"
)

type TopoV1beta1Interface interface {
	RESTClient() rest.Interface
	EntitiesGetter
	KindsGetter
	RelationsGetter
}

// TopoV1beta1Client is used to interact with features provided by the topo.onosproject.org group.
type TopoV1beta1Client struct {
	restClient rest.Interface
}

func (c *TopoV1beta1Client) Entities(namespace string) EntityInterface {
	return newEntities(c, namespace)
}

func (c *TopoV1beta1Client) Kinds(namespace string) KindInterface {
	return newKinds(c, namespace)
}

func (c *TopoV1beta1Client) Relations(namespace string) RelationInterface {
	return newRelations(c, namespace)
}

// NewForConfig creates a new TopoV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*TopoV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TopoV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new TopoV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TopoV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TopoV1beta1Client for the given RESTClient.
func New(c rest.Interface) *TopoV1beta1Client {
	return &TopoV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TopoV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
