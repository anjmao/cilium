// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type CiliumV2Interface interface {
	RESTClient() rest.Interface
	CiliumNetworkPoliciesGetter
}

// CiliumV2Client is used to interact with features provided by the cilium.io group.
type CiliumV2Client struct {
	restClient rest.Interface
}

func (c *CiliumV2Client) CiliumNetworkPolicies(namespace string) CiliumNetworkPolicyInterface {
	return newCiliumNetworkPolicies(c, namespace)
}

// NewForConfig creates a new CiliumV2Client for the given config.
func NewForConfig(c *rest.Config) (*CiliumV2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CiliumV2Client{client}, nil
}

// NewForConfigOrDie creates a new CiliumV2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CiliumV2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CiliumV2Client for the given RESTClient.
func New(c rest.Interface) *CiliumV2Client {
	return &CiliumV2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CiliumV2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
