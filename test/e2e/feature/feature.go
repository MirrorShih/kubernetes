/*
Copyright 2023 The Kubernetes Authors.

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

// Package feature contains pre-defined features used by test/e2e and/or
// test/e2e_node.
package feature

import (
	"k8s.io/kubernetes/test/e2e/framework"
)

var (
	// Please keep the list in alphabetical order.

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	APIServerIdentity = framework.WithFeature(framework.ValidFeatures.Add("APIServerIdentity"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	AppArmor = framework.WithFeature(framework.ValidFeatures.Add("AppArmor"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	BootstrapTokens = framework.WithFeature(framework.ValidFeatures.Add("BootstrapTokens"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	BoundServiceAccountTokenVolume = framework.WithFeature(framework.ValidFeatures.Add("BoundServiceAccountTokenVolume"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	CloudProvider = framework.WithFeature(framework.ValidFeatures.Add("CloudProvider"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability1 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability1"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability2 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability2"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability3 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability3"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability4 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability4"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability5 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability5"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterAutoscalerScalability6 = framework.WithFeature(framework.ValidFeatures.Add("ClusterAutoscalerScalability6"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterDowngrade = framework.WithFeature(framework.ValidFeatures.Add("ClusterDowngrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterScaleUpBypassScheduler = framework.WithFeature(framework.ValidFeatures.Add("ClusterScaleUpBypassScheduler"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterSizeAutoscalingGpu = framework.WithFeature(framework.ValidFeatures.Add("ClusterSizeAutoscalingGpu"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterSizeAutoscalingScaleDown = framework.WithFeature(framework.ValidFeatures.Add("ClusterSizeAutoscalingScaleDown"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterSizeAutoscalingScaleUp = framework.WithFeature(framework.ValidFeatures.Add("ClusterSizeAutoscalingScaleUp"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterTrustBundle = framework.WithFeature(framework.ValidFeatures.Add("ClusterTrustBundle"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterTrustBundleProjection = framework.WithFeature(framework.ValidFeatures.Add("ClusterTrustBundleProjection"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ClusterUpgrade = framework.WithFeature(framework.ValidFeatures.Add("ClusterUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ComprehensiveNamespaceDraining = framework.WithFeature(framework.ValidFeatures.Add("ComprehensiveNamespaceDraining"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	CPUManager = framework.WithFeature(framework.ValidFeatures.Add("CPUManager"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	CustomMetricsAutoscaling = framework.WithFeature(framework.ValidFeatures.Add("CustomMetricsAutoscaling"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	DeviceManager = framework.WithFeature(framework.ValidFeatures.Add("DeviceManager"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	DevicePluginProbe = framework.WithFeature(framework.ValidFeatures.Add("DevicePluginProbe"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Downgrade = framework.WithFeature(framework.ValidFeatures.Add("Downgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	DynamicResourceAllocation = framework.WithFeature(framework.ValidFeatures.Add("DynamicResourceAllocation"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	EphemeralStorage = framework.WithFeature(framework.ValidFeatures.Add("EphemeralStorage"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Example = framework.WithFeature(framework.ValidFeatures.Add("Example"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ExperimentalResourceUsageTracking = framework.WithFeature(framework.ValidFeatures.Add("ExperimentalResourceUsageTracking"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Flexvolumes = framework.WithFeature(framework.ValidFeatures.Add("Flexvolumes"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GKENodePool = framework.WithFeature(framework.ValidFeatures.Add("GKENodePool"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GPUClusterDowngrade = framework.WithFeature(framework.ValidFeatures.Add("GPUClusterDowngrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GPUClusterUpgrade = framework.WithFeature(framework.ValidFeatures.Add("GPUClusterUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GPUDevicePlugin = framework.WithFeature(framework.ValidFeatures.Add("GPUDevicePlugin"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GPUMasterUpgrade = framework.WithFeature(framework.ValidFeatures.Add("GPUMasterUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	GPUUpgrade = framework.WithFeature(framework.ValidFeatures.Add("GPUUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	HAMaster = framework.WithFeature(framework.ValidFeatures.Add("HAMaster"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	HPA = framework.WithFeature(framework.ValidFeatures.Add("HPA"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	HugePages = framework.WithFeature(framework.ValidFeatures.Add("HugePages"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Ingress = framework.WithFeature(framework.ValidFeatures.Add("Ingress"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	IngressScale = framework.WithFeature(framework.ValidFeatures.Add("IngressScale"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	InPlacePodVerticalScaling = framework.WithFeature(framework.ValidFeatures.Add("InPlacePodVerticalScaling"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	IPv6DualStack = framework.WithFeature(framework.ValidFeatures.Add("IPv6DualStack"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Kind = framework.WithFeature(framework.ValidFeatures.Add("Kind"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	KubeletCredentialProviders = framework.WithFeature(framework.ValidFeatures.Add("KubeletCredentialProviders"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	KubeletSecurity = framework.WithFeature(framework.ValidFeatures.Add("KubeletSecurity"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	KubeProxyDaemonSetDowngrade = framework.WithFeature(framework.ValidFeatures.Add("KubeProxyDaemonSetDowngrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	KubeProxyDaemonSetUpgrade = framework.WithFeature(framework.ValidFeatures.Add("KubeProxyDaemonSetUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	KubeProxyDaemonSetMigration = framework.WithFeature(framework.ValidFeatures.Add("KubeProxyDaemonSetMigration"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	LabelSelector = framework.WithFeature(framework.ValidFeatures.Add("LabelSelector"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	LocalStorageCapacityIsolation = framework.WithFeature(framework.ValidFeatures.Add("LocalStorageCapacityIsolation"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	LocalStorageCapacityIsolationQuota = framework.WithFeature(framework.ValidFeatures.Add("LocalStorageCapacityIsolationQuota"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	MasterUpgrade = framework.WithFeature(framework.ValidFeatures.Add("MasterUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	MemoryManager = framework.WithFeature(framework.ValidFeatures.Add("MemoryManager"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NEG = framework.WithFeature(framework.ValidFeatures.Add("NEG"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NetworkingDNS = framework.WithFeature(framework.ValidFeatures.Add("Networking-DNS"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NetworkingIPv4 = framework.WithFeature(framework.ValidFeatures.Add("Networking-IPv4"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NetworkingIPv6 = framework.WithFeature(framework.ValidFeatures.Add("Networking-IPv6"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NetworkingPerformance = framework.WithFeature(framework.ValidFeatures.Add("Networking-Performance"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NetworkPolicy = framework.WithFeature(framework.ValidFeatures.Add("NetworkPolicy"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NodeAuthenticator = framework.WithFeature(framework.ValidFeatures.Add("NodeAuthenticator"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NodeAuthorizer = framework.WithFeature(framework.ValidFeatures.Add("NodeAuthorizer"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NodeLogQuery = framework.WithFeature(framework.ValidFeatures.Add("NodeLogQuery"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NodeOutOfServiceVolumeDetach = framework.WithFeature(framework.ValidFeatures.Add("NodeOutOfServiceVolumeDetach"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	NoSNAT = framework.WithFeature(framework.ValidFeatures.Add("NoSNAT"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PersistentVolumeLastPhaseTransitionTime = framework.WithFeature(framework.ValidFeatures.Add("PersistentVolumeLastPhaseTransitionTime"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PerformanceDNS = framework.WithFeature(framework.ValidFeatures.Add("PerformanceDNS"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodGarbageCollector = framework.WithFeature(framework.ValidFeatures.Add("PodGarbageCollector"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodHostIPs = framework.WithFeature(framework.ValidFeatures.Add("PodHostIPs"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodLifecycleSleepAction = framework.WithFeature(framework.ValidFeatures.Add("PodLifecycleSleepAction"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodPriority = framework.WithFeature(framework.ValidFeatures.Add("PodPriority"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodReadyToStartContainersCondition = framework.WithFeature(framework.ValidFeatures.Add("PodReadyToStartContainersCondition"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	PodResources = framework.WithFeature(framework.ValidFeatures.Add("PodResources"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Reboot = framework.WithFeature(framework.ValidFeatures.Add("Reboot"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ReclaimPolicy = framework.WithFeature(framework.ValidFeatures.Add("ReclaimPolicy"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	RecoverVolumeExpansionFailure = framework.WithFeature(framework.ValidFeatures.Add("RecoverVolumeExpansionFailure"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Recreate = framework.WithFeature(framework.ValidFeatures.Add("Recreate"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	RegularResourceUsageTracking = framework.WithFeature(framework.ValidFeatures.Add("RegularResourceUsageTracking"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ScopeSelectors = framework.WithFeature(framework.ValidFeatures.Add("ScopeSelectors"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	SCTPConnectivity = framework.WithFeature(framework.ValidFeatures.Add("SCTPConnectivity"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	SeccompDefault = framework.WithFeature(framework.ValidFeatures.Add("SeccompDefault"))

	// Owner: sig-storage
	// This feature marks tests that need all schedulable Linux nodes in the cluster to have SELinux enabled.
	SELinux = framework.WithFeature(framework.ValidFeatures.Add("SELinux"))

	// Owner: sig-storage
	// This feature marks tests that need SELinuxMountReadWriteOncePod feature gate enabled and SELinuxMount **disabled**.
	// This is a temporary feature to allow testing of metrics when SELinuxMount is disabled.
	// TODO: remove when SELinuxMount feature gate is enabled by default.
	SELinuxMountReadWriteOncePodOnly = framework.WithFeature(framework.ValidFeatures.Add("SELinuxMountReadWriteOncePodOnly"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ServiceCIDRs = framework.WithFeature(framework.ValidFeatures.Add("ServiceCIDRs"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	SidecarContainers = framework.WithFeature(framework.ValidFeatures.Add("SidecarContainers"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StackdriverAcceleratorMonitoring = framework.WithFeature(framework.ValidFeatures.Add("StackdriverAcceleratorMonitoring"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StackdriverCustomMetrics = framework.WithFeature(framework.ValidFeatures.Add("StackdriverCustomMetrics"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StackdriverExternalMetrics = framework.WithFeature(framework.ValidFeatures.Add("StackdriverExternalMetrics"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StackdriverMetadataAgent = framework.WithFeature(framework.ValidFeatures.Add("StackdriverMetadataAgent"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StackdriverMonitoring = framework.WithFeature(framework.ValidFeatures.Add("StackdriverMonitoring"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StandaloneMode = framework.WithFeature(framework.ValidFeatures.Add("StandaloneMode"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StatefulSet = framework.WithFeature(framework.ValidFeatures.Add("StatefulSet"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StatefulSetStartOrdinal = framework.WithFeature(framework.ValidFeatures.Add("StatefulSetStartOrdinal"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StatefulUpgrade = framework.WithFeature(framework.ValidFeatures.Add("StatefulUpgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StorageProvider = framework.WithFeature(framework.ValidFeatures.Add("StorageProvider"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	StorageVersionAPI = framework.WithFeature(framework.ValidFeatures.Add("StorageVersionAPI"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	TopologyHints = framework.WithFeature(framework.ValidFeatures.Add("Topology Hints"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	TopologyManager = framework.WithFeature(framework.ValidFeatures.Add("TopologyManager"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	UDP = framework.WithFeature(framework.ValidFeatures.Add("UDP"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Upgrade = framework.WithFeature(framework.ValidFeatures.Add("Upgrade"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	UserNamespacesSupport = framework.WithFeature(framework.ValidFeatures.Add("UserNamespacesSupport"))

	// Owned by SIG Node
	// Can be used when the UserNamespacesPodSecurityStandards kubelet feature
	// gate is enabled to relax the application of Pod Security Standards in a
	// controlled way.
	UserNamespacesPodSecurityStandards = framework.WithFeature(framework.ValidFeatures.Add("UserNamespacesPodSecurityStandards"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	ValidatingAdmissionPolicy = framework.WithFeature(framework.ValidFeatures.Add("ValidatingAdmissionPolicy"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Volumes = framework.WithFeature(framework.ValidFeatures.Add("Volumes"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	VolumeSnapshotDataSource = framework.WithFeature(framework.ValidFeatures.Add("VolumeSnapshotDataSource"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	VolumeSourceXFS = framework.WithFeature(framework.ValidFeatures.Add("VolumeSourceXFS"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Vsphere = framework.WithFeature(framework.ValidFeatures.Add("vsphere"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	WatchList = framework.WithFeature(framework.ValidFeatures.Add("WatchList"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	Windows = framework.WithFeature(framework.ValidFeatures.Add("Windows"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	WindowsHostProcessContainers = framework.WithFeature(framework.ValidFeatures.Add("WindowsHostProcessContainers"))

	// TODO: document the feature (owning SIG, when to use this feature for a test)
	WindowsHyperVContainers = framework.WithFeature(framework.ValidFeatures.Add("WindowsHyperVContainers"))

	// Please keep the list in alphabetical order.
)

func init() {
	// This prevents adding additional ad-hoc features in tests.
	framework.ValidFeatures.Freeze()
}
