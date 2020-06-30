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

package versioned

import (
	"fmt"

	awsv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/aws/v1alpha1"
	azurermv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/azurerm/v1alpha1"
	digitaloceanv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/digitalocean/v1alpha1"
	googlev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/google/v1alpha1"
	linodev1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/linode/v1alpha1"
	modulesv1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/modules/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AwsV1alpha1() awsv1alpha1.AwsV1alpha1Interface
	AzurermV1alpha1() azurermv1alpha1.AzurermV1alpha1Interface
	DigitaloceanV1alpha1() digitaloceanv1alpha1.DigitaloceanV1alpha1Interface
	GoogleV1alpha1() googlev1alpha1.GoogleV1alpha1Interface
	LinodeV1alpha1() linodev1alpha1.LinodeV1alpha1Interface
	ModulesV1alpha1() modulesv1alpha1.ModulesV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	awsV1alpha1          *awsv1alpha1.AwsV1alpha1Client
	azurermV1alpha1      *azurermv1alpha1.AzurermV1alpha1Client
	digitaloceanV1alpha1 *digitaloceanv1alpha1.DigitaloceanV1alpha1Client
	googleV1alpha1       *googlev1alpha1.GoogleV1alpha1Client
	linodeV1alpha1       *linodev1alpha1.LinodeV1alpha1Client
	modulesV1alpha1      *modulesv1alpha1.ModulesV1alpha1Client
}

// AwsV1alpha1 retrieves the AwsV1alpha1Client
func (c *Clientset) AwsV1alpha1() awsv1alpha1.AwsV1alpha1Interface {
	return c.awsV1alpha1
}

// AzurermV1alpha1 retrieves the AzurermV1alpha1Client
func (c *Clientset) AzurermV1alpha1() azurermv1alpha1.AzurermV1alpha1Interface {
	return c.azurermV1alpha1
}

// DigitaloceanV1alpha1 retrieves the DigitaloceanV1alpha1Client
func (c *Clientset) DigitaloceanV1alpha1() digitaloceanv1alpha1.DigitaloceanV1alpha1Interface {
	return c.digitaloceanV1alpha1
}

// GoogleV1alpha1 retrieves the GoogleV1alpha1Client
func (c *Clientset) GoogleV1alpha1() googlev1alpha1.GoogleV1alpha1Interface {
	return c.googleV1alpha1
}

// LinodeV1alpha1 retrieves the LinodeV1alpha1Client
func (c *Clientset) LinodeV1alpha1() linodev1alpha1.LinodeV1alpha1Interface {
	return c.linodeV1alpha1
}

// ModulesV1alpha1 retrieves the ModulesV1alpha1Client
func (c *Clientset) ModulesV1alpha1() modulesv1alpha1.ModulesV1alpha1Interface {
	return c.modulesV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.awsV1alpha1, err = awsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.azurermV1alpha1, err = azurermv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.digitaloceanV1alpha1, err = digitaloceanv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.googleV1alpha1, err = googlev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.linodeV1alpha1, err = linodev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.modulesV1alpha1, err = modulesv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.awsV1alpha1 = awsv1alpha1.NewForConfigOrDie(c)
	cs.azurermV1alpha1 = azurermv1alpha1.NewForConfigOrDie(c)
	cs.digitaloceanV1alpha1 = digitaloceanv1alpha1.NewForConfigOrDie(c)
	cs.googleV1alpha1 = googlev1alpha1.NewForConfigOrDie(c)
	cs.linodeV1alpha1 = linodev1alpha1.NewForConfigOrDie(c)
	cs.modulesV1alpha1 = modulesv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.awsV1alpha1 = awsv1alpha1.New(c)
	cs.azurermV1alpha1 = azurermv1alpha1.New(c)
	cs.digitaloceanV1alpha1 = digitaloceanv1alpha1.New(c)
	cs.googleV1alpha1 = googlev1alpha1.New(c)
	cs.linodeV1alpha1 = linodev1alpha1.New(c)
	cs.modulesV1alpha1 = modulesv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
