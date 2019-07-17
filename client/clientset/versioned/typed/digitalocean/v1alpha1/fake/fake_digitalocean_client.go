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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/digitalocean/v1alpha1"
)

type FakeDigitaloceanV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDigitaloceanV1alpha1) Cdns() v1alpha1.CdnInterface {
	return &FakeCdns{c}
}

func (c *FakeDigitaloceanV1alpha1) Certificates() v1alpha1.CertificateInterface {
	return &FakeCertificates{c}
}

func (c *FakeDigitaloceanV1alpha1) DatabaseClusters() v1alpha1.DatabaseClusterInterface {
	return &FakeDatabaseClusters{c}
}

func (c *FakeDigitaloceanV1alpha1) Domains() v1alpha1.DomainInterface {
	return &FakeDomains{c}
}

func (c *FakeDigitaloceanV1alpha1) Droplets() v1alpha1.DropletInterface {
	return &FakeDroplets{c}
}

func (c *FakeDigitaloceanV1alpha1) DropletSnapshots() v1alpha1.DropletSnapshotInterface {
	return &FakeDropletSnapshots{c}
}

func (c *FakeDigitaloceanV1alpha1) Firewalls() v1alpha1.FirewallInterface {
	return &FakeFirewalls{c}
}

func (c *FakeDigitaloceanV1alpha1) FloatingIps() v1alpha1.FloatingIpInterface {
	return &FakeFloatingIps{c}
}

func (c *FakeDigitaloceanV1alpha1) FloatingIpAssignments() v1alpha1.FloatingIpAssignmentInterface {
	return &FakeFloatingIpAssignments{c}
}

func (c *FakeDigitaloceanV1alpha1) KubernetesClusters() v1alpha1.KubernetesClusterInterface {
	return &FakeKubernetesClusters{c}
}

func (c *FakeDigitaloceanV1alpha1) KubernetesNodePools() v1alpha1.KubernetesNodePoolInterface {
	return &FakeKubernetesNodePools{c}
}

func (c *FakeDigitaloceanV1alpha1) Loadbalancers() v1alpha1.LoadbalancerInterface {
	return &FakeLoadbalancers{c}
}

func (c *FakeDigitaloceanV1alpha1) Projects() v1alpha1.ProjectInterface {
	return &FakeProjects{c}
}

func (c *FakeDigitaloceanV1alpha1) Records() v1alpha1.RecordInterface {
	return &FakeRecords{c}
}

func (c *FakeDigitaloceanV1alpha1) SpacesBuckets() v1alpha1.SpacesBucketInterface {
	return &FakeSpacesBuckets{c}
}

func (c *FakeDigitaloceanV1alpha1) SshKeys() v1alpha1.SshKeyInterface {
	return &FakeSshKeys{c}
}

func (c *FakeDigitaloceanV1alpha1) Tags() v1alpha1.TagInterface {
	return &FakeTags{c}
}

func (c *FakeDigitaloceanV1alpha1) Volumes() v1alpha1.VolumeInterface {
	return &FakeVolumes{c}
}

func (c *FakeDigitaloceanV1alpha1) VolumeAttachments() v1alpha1.VolumeAttachmentInterface {
	return &FakeVolumeAttachments{c}
}

func (c *FakeDigitaloceanV1alpha1) VolumeSnapshots() v1alpha1.VolumeSnapshotInterface {
	return &FakeVolumeSnapshots{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDigitaloceanV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}