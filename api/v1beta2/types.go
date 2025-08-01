/*
 Copyright (c) 2021, 2022 Oracle and/or its affiliates.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package v1beta2

const (
	ControlPlaneRole         = "control-plane"
	ControlPlaneEndpointRole = "control-plane-endpoint"
	WorkerRole               = "worker"
	ServiceLoadBalancerRole  = "service-lb"
	PodRole                  = "pod"
	Private                  = "private"
	Public                   = "public"
	Custom                   = "custom"
)

// OCIClusterSubnetRoles a slice of all the subnet roles for self managed cluster
var OCIClusterSubnetRoles = []Role{ControlPlaneRole, ControlPlaneEndpointRole, WorkerRole, ServiceLoadBalancerRole, Custom}

// OCIManagedClusterSubnetRoles a slice of all the subnet roles for managed cluster
var OCIManagedClusterSubnetRoles = []Role{PodRole, ControlPlaneEndpointRole, WorkerRole, ServiceLoadBalancerRole, Custom}

// NetworkDetails defines the configuration options for the network
type NetworkDetails struct {
	// SubnetId defines the ID of the subnet to use. This parameter takes priority over SubnetName.
	SubnetId *string `json:"subnetId,omitempty"`

	// AssignIPv6 determines whether to assign a IPv6 address to the instance.
	// +optional
	AssignIpv6Ip bool `json:"assignIpv6Ip,omitempty"`

	// AssignPublicIp defines whether the instance should have a public IP address
	AssignPublicIp bool `json:"assignPublicIp,omitempty"`

	// SubnetName defines the subnet name to use for the VNIC
	SubnetName string `json:"subnetName,omitempty"`

	// SkipSourceDestCheck defines whether the source/destination check is disabled on the VNIC.
	SkipSourceDestCheck *bool `json:"skipSourceDestCheck,omitempty"`

	// NSGId defines the ID of the NSG to use. This parameter takes priority over NsgNames.
	// Deprecated, please use NetworkDetails.NSGIds
	NSGId *string `json:"nsgId,omitempty"`

	// NSGIds defines the list of NSG IDs to use. This parameter takes priority over NsgNames.
	NSGIds []string `json:"nsgIds,omitempty"`

	// NsgNames defines a list of the nsg names of the network security groups (NSGs) to add the VNIC to.
	NsgNames []string `json:"nsgNames,omitempty"`

	// HostnameLabel defines the hostname for the VNIC's primary private IP. Used for DNS.
	HostnameLabel *string `json:"hostnameLabel,omitempty"`

	// DisplayName defines a user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty"`

	// AssignPrivateDnsRecord defines whether the VNIC should be assigned a DNS record.
	AssignPrivateDnsRecord *bool `json:"assignPrivateDnsRecord,omitempty"`
}

type VnicAttachment struct {
	// VnicAttachmentId defines the ID of the VnicAttachment
	VnicAttachmentId *string `json:"vnicAttachmentId,omitempty"`

	// AssignPublicIp defines whether the vnic should have a public IP address
	// +optional
	AssignPublicIp bool `json:"assignPublicIp,omitempty"`

	// SubnetName defines the subnet name to use for the VNIC
	// Defaults to the "worker" subnet if not provided
	// +optional
	SubnetName string `json:"subnetName,omitempty"`

	// DisplayName defines a user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `json:"displayName"`

	// NicIndex defines which physical Network Interface Card (NIC) to use
	// You can determine which NICs are active for a shape by reviewing the
	// https://docs.oracle.com/en-us/iaas/Content/Compute/References/computeshapes.htm
	// +optional
	NicIndex *int `json:"nicIndex,omitempty"`
}

// LaunchOptionsBootVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsBootVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsBootVolumeTypeEnum
const (
	LaunchOptionsBootVolumeTypeIscsi           LaunchOptionsBootVolumeTypeEnum = "ISCSI"
	LaunchOptionsBootVolumeTypeScsi            LaunchOptionsBootVolumeTypeEnum = "SCSI"
	LaunchOptionsBootVolumeTypeIde             LaunchOptionsBootVolumeTypeEnum = "IDE"
	LaunchOptionsBootVolumeTypeVfio            LaunchOptionsBootVolumeTypeEnum = "VFIO"
	LaunchOptionsBootVolumeTypeParavirtualized LaunchOptionsBootVolumeTypeEnum = "PARAVIRTUALIZED"
)

// LaunchOptionsFirmwareEnum Enum with underlying type: string
type LaunchOptionsFirmwareEnum string

// Set of constants representing the allowable values for LaunchOptionsFirmwareEnum
const (
	LaunchOptionsFirmwareBios   LaunchOptionsFirmwareEnum = "BIOS"
	LaunchOptionsFirmwareUefi64 LaunchOptionsFirmwareEnum = "UEFI_64"
)

// LaunchOptionsNetworkTypeEnum Enum with underlying type: string
type LaunchOptionsNetworkTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsNetworkTypeEnum
const (
	LaunchOptionsNetworkTypeE1000           LaunchOptionsNetworkTypeEnum = "E1000"
	LaunchOptionsNetworkTypeVfio            LaunchOptionsNetworkTypeEnum = "VFIO"
	LaunchOptionsNetworkTypeParavirtualized LaunchOptionsNetworkTypeEnum = "PARAVIRTUALIZED"
)

// LaunchOptionsRemoteDataVolumeTypeEnum Enum with underlying type: string
type LaunchOptionsRemoteDataVolumeTypeEnum string

// Set of constants representing the allowable values for LaunchOptionsRemoteDataVolumeTypeEnum
const (
	LaunchOptionsRemoteDataVolumeTypeIscsi           LaunchOptionsRemoteDataVolumeTypeEnum = "ISCSI"
	LaunchOptionsRemoteDataVolumeTypeScsi            LaunchOptionsRemoteDataVolumeTypeEnum = "SCSI"
	LaunchOptionsRemoteDataVolumeTypeIde             LaunchOptionsRemoteDataVolumeTypeEnum = "IDE"
	LaunchOptionsRemoteDataVolumeTypeVfio            LaunchOptionsRemoteDataVolumeTypeEnum = "VFIO"
	LaunchOptionsRemoteDataVolumeTypeParavirtualized LaunchOptionsRemoteDataVolumeTypeEnum = "PARAVIRTUALIZED"
)

// LaunchOptions Options for tuning the compatibility and performance of VM shapes. The values that you specify override any
// default values.
type LaunchOptions struct {

	// BootVolumeType defines Emulation type for the boot volume.
	// * `ISCSI` - ISCSI attached block storage device.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage. This is the default option for local data
	// volumes on platform images.
	// * `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block
	// storage volumes on platform images.
	BootVolumeType LaunchOptionsBootVolumeTypeEnum `json:"bootVolumeType,omitempty"`

	// Firmware defines the firmware used to boot VM. Select the option that matches your operating system.
	// * `BIOS` - Boot VM using BIOS style firmware. This is compatible with both 32 bit and 64 bit operating
	// systems that boot using MBR style bootloaders.
	// * `UEFI_64` - Boot VM using UEFI style firmware compatible with 64 bit operating systems. This is the
	// default for platform images.
	Firmware LaunchOptionsFirmwareEnum `json:"firmware,omitempty"`

	// NetworkType defines the emulation type for the physical network interface card (NIC).
	// * `E1000` - Emulated Gigabit ethernet controller. Compatible with Linux e1000 network driver.
	// * `VFIO` - Direct attached Virtual Function network controller. This is the networking type
	// when you launch an instance using hardware-assisted (SR-IOV) networking.
	// * `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
	NetworkType LaunchOptionsNetworkTypeEnum `json:"networkType,omitempty"`

	// RemoteDataVolumeType defines the emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device.
	// * `SCSI` - Emulated SCSI disk.
	// * `IDE` - Emulated IDE disk.
	// * `VFIO` - Direct attached Virtual Function storage. This is the default option for local data
	// volumes on platform images.
	// * `PARAVIRTUALIZED` - Paravirtualized disk. This is the default for boot volumes and remote block
	// storage volumes on platform images.
	RemoteDataVolumeType LaunchOptionsRemoteDataVolumeTypeEnum `json:"remoteDataVolumeType,omitempty"`

	// IsConsistentVolumeNamingEnabled defines whether to enable consistent volume naming feature. Defaults to false.
	IsConsistentVolumeNamingEnabled *bool `json:"isConsistentVolumeNamingEnabled,omitempty"`
}

// InstanceSourceViaImageConfig The configuration options for booting up instances via images
type InstanceSourceViaImageConfig struct {
	// KmsKeyId defines the OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId *string `json:"kmsKeyId,omitempty"`

	// BootVolumeVpusPerGB defines the number of volume performance units (VPUs) that will be applied to this volume per GB,
	// representing the Block Volume service's elastic performance options.
	// See Block Volume Performance Levels (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.
	// Allowed values:
	//   * `10`: Represents Balanced option.
	//   * `20`: Represents Higher Performance option.
	//   * `30`-`120`: Represents the Ultra High Performance option.
	// For volumes with the auto-tuned performance feature enabled, this is set to the default (minimum) VPUs/GB.
	BootVolumeVpusPerGB *int64 `json:"bootVolumeVpusPerGB,omitempty"`
}

// LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum Enum with underlying type: string
type PlatformConfigTypeEnum string

// Set of constants representing the allowable values for LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum
const (
	PlatformConfigTypeAmdRomeBmGpu   PlatformConfigTypeEnum = "AMD_ROME_BM_GPU"
	PlatformConfigTypeAmdRomeBm      PlatformConfigTypeEnum = "AMD_ROME_BM"
	PlatformConfigTypeIntelIcelakeBm PlatformConfigTypeEnum = "INTEL_ICELAKE_BM"
	PlatformConfigTypeAmdvm          PlatformConfigTypeEnum = "AMD_VM"
	PlatformConfigTypeIntelVm        PlatformConfigTypeEnum = "INTEL_VM"
	PlatformConfigTypeIntelSkylakeBm PlatformConfigTypeEnum = "INTEL_SKYLAKE_BM"
	PlatformConfigTypeAmdMilanBm     PlatformConfigTypeEnum = "AMD_MILAN_BM"
)

// PlatformConfig defines the platform config parameters
type PlatformConfig struct {
	// The type of platform configuration. Valid values are
	// * `AMD_ROME_BM_GPU`
	// * `AMD_ROME_BM`
	// * `INTEL_ICELAKE_BM`
	// * `AMD_VM`
	// * `INTEL_VM`
	// * `INTEL_SKYLAKE_BM`
	// * `AMD_MILAN_BM`
	// Based on the enum, exactly one of the specific configuration types must be set
	PlatformConfigType PlatformConfigTypeEnum `json:"platformConfigType,omitempty"`

	// AmdMilanBmPlatformConfig describe AMD Milan BM platform configuration
	AmdMilanBmPlatformConfig AmdMilanBmPlatformConfig `json:"amdMilanBmPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe AMD Rome BM platform configuration
	AmdRomeBmPlatformConfig AmdRomeBmPlatformConfig `json:"amdRomeBmPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe Intel Skylke BM platform configuration
	IntelSkylakeBmPlatformConfig IntelSkylakeBmPlatformConfig `json:"intelSkylakeBmPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe Intel Skylke BM platform configuration
	IntelIcelakeBmPlatformConfig IntelIcelakeBmPlatformConfig `json:"intelIcelakeBmPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe AMD Rome BM platform configuration
	AmdRomeBmGpuPlatformConfig AmdRomeBmGpuPlatformConfig `json:"amdRomeBmGpuPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe Intel VM platform configuration
	IntelVmPlatformConfig IntelVmPlatformConfig `json:"intelVmPlatformConfig,omitempty"`

	// AmdMilanBmPlatformConfig describe AMD VM platform configuration
	AmdVmPlatformConfig AmdVmPlatformConfig `json:"amdVmPlatformConfig,omitempty"`
}

// AmdMilanBmPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdMilanBmPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdMilanBmPlatformConfigNumaNodesPerSocketEnum
const (
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps0 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps1 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps2 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdMilanBmPlatformConfigNumaNodesPerSocketNps4 AmdMilanBmPlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

// AmdMilanBmPlatformConfig The platform configuration used when launching a bare metal instance with one of the following shapes: BM.Standard.E4.128
// or BM.DenseIO.E4.128 (the AMD Milan platform).
type AmdMilanBmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`

	// Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also
	// called simultaneous multithreading (SMT) or Intel Hyper-Threading.
	// Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple
	// independent threads of execution, to better use the resources and increase the efficiency
	// of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which
	// can provide higher or more predictable performance for some workloads.
	IsSymmetricMultiThreadingEnabled *bool `json:"isSymmetricMultiThreadingEnabled,omitempty"`

	// Whether the Access Control Service is enabled on the instance. When enabled,
	// the platform can enforce PCIe device isolation, required for VFIO device pass-through.
	IsAccessControlServiceEnabled *bool `json:"isAccessControlServiceEnabled,omitempty"`

	// Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes
	// or VT-x for Intel shapes.
	AreVirtualInstructionsEnabled *bool `json:"areVirtualInstructionsEnabled,omitempty"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `json:"isInputOutputMemoryManagementUnitEnabled,omitempty"`

	// The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage
	// results in a fractional number of cores, the system rounds up the number of cores across processors
	// and provisions an instance with a whole number of cores.
	// If the applications that you run on the instance use a core-based licensing model and need fewer cores
	// than the full size of the shape, you can disable cores to reduce your licensing costs. The instance
	// itself is billed for the full shape, regardless of whether all cores are enabled.
	PercentageOfCoresEnabled *int `json:"percentageOfCoresEnabled,omitempty"`

	// The number of NUMA nodes per socket (NPS).
	// The following values are supported:
	// * `NPS0`
	// * `NPS1`
	// * `NPS2`
	// * `NPS4`
	NumaNodesPerSocket AmdMilanBmPlatformConfigNumaNodesPerSocketEnum `json:"numaNodesPerSocket,omitempty"`
}

// AmdRomeBmPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdRomeBmPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdRomeBmPlatformConfigNumaNodesPerSocketEnum
const (
	AmdRomeBmPlatformConfigNumaNodesPerSocketNps0 AmdRomeBmPlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdRomeBmPlatformConfigNumaNodesPerSocketNps1 AmdRomeBmPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdRomeBmPlatformConfigNumaNodesPerSocketNps2 AmdRomeBmPlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdRomeBmPlatformConfigNumaNodesPerSocketNps4 AmdRomeBmPlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

// AmdRomeBmPlatformConfig The platform configuration of a bare metal instance that uses the BM.Standard.E3.128 shape (the AMD Rome platform).
type AmdRomeBmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`

	// Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also
	// called simultaneous multithreading (SMT) or Intel Hyper-Threading.
	// Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple
	// independent threads of execution, to better use the resources and increase the efficiency
	// of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which
	// can provide higher or more predictable performance for some workloads.
	IsSymmetricMultiThreadingEnabled *bool `json:"isSymmetricMultiThreadingEnabled,omitempty"`

	// Whether the Access Control Service is enabled on the instance. When enabled,
	// the platform can enforce PCIe device isolation, required for VFIO device pass-through.
	IsAccessControlServiceEnabled *bool `json:"isAccessControlServiceEnabled,omitempty"`

	// Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes
	// or VT-x for Intel shapes.
	AreVirtualInstructionsEnabled *bool `json:"areVirtualInstructionsEnabled,omitempty"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `json:"isInputOutputMemoryManagementUnitEnabled,omitempty"`

	// The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage
	// results in a fractional number of cores, the system rounds up the number of cores across processors
	// and provisions an instance with a whole number of cores.
	// If the applications that you run on the instance use a core-based licensing model and need fewer cores
	// than the full size of the shape, you can disable cores to reduce your licensing costs. The instance
	// itself is billed for the full shape, regardless of whether all cores are enabled.
	PercentageOfCoresEnabled *int `json:"percentageOfCoresEnabled,omitempty"`

	// The number of NUMA nodes per socket (NPS).
	// The following values are supported:
	// * `NPS0`
	// * `NPS1`
	// * `NPS2`
	// * `NPS4`
	NumaNodesPerSocket AmdRomeBmPlatformConfigNumaNodesPerSocketEnum `json:"numaNodesPerSocket,omitempty"`
}

// IntelSkylakeBmPlatformConfig The platform configuration of a bare metal instance that uses one of the following shapes:
// BM.Standard2.52, BM.GPU2.2, BM.GPU3.8, or BM.DenseIO2.52 (the Intel Skylake platform).
type IntelSkylakeBmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`
}

// AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum
const (
	AmdRomeBmGpuPlatformConfigNumaNodesPerSocketNps0 AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS0"
	AmdRomeBmGpuPlatformConfigNumaNodesPerSocketNps1 AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	AmdRomeBmGpuPlatformConfigNumaNodesPerSocketNps2 AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS2"
	AmdRomeBmGpuPlatformConfigNumaNodesPerSocketNps4 AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum = "NPS4"
)

// AmdRomeBmGpuPlatformConfig The platform configuration of a bare metal GPU instance that uses the BM.GPU4.8 shape
// (the AMD Rome platform).
type AmdRomeBmGpuPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`

	// Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also
	// called simultaneous multithreading (SMT) or Intel Hyper-Threading.
	// Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple
	// independent threads of execution, to better use the resources and increase the efficiency
	// of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which
	// can provide higher or more predictable performance for some workloads.
	IsSymmetricMultiThreadingEnabled *bool `json:"isSymmetricMultiThreadingEnabled,omitempty"`

	// Whether the Access Control Service is enabled on the instance. When enabled,
	// the platform can enforce PCIe device isolation, required for VFIO device pass-through.
	IsAccessControlServiceEnabled *bool `json:"isAccessControlServiceEnabled,omitempty"`

	// Whether virtualization instructions are available. For example, Secure Virtual Machine for AMD shapes
	// or VT-x for Intel shapes.
	AreVirtualInstructionsEnabled *bool `json:"areVirtualInstructionsEnabled,omitempty"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `json:"isInputOutputMemoryManagementUnitEnabled,omitempty"`

	// The number of NUMA nodes per socket (NPS).
	// The following values are supported:
	// * `NPS0`
	// * `NPS1`
	// * `NPS2`
	// * `NPS4`
	NumaNodesPerSocket AmdRomeBmGpuPlatformConfigNumaNodesPerSocketEnum `json:"numaNodesPerSocket,omitempty"`
}

// IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum Enum with underlying type: string
type IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum string

// Set of constants representing the allowable values for IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum
const (
	IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps1 IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum = "NPS1"
	IntelIcelakeBmPlatformConfigNumaNodesPerSocketNps2 IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum = "NPS2"
)

// IntelIcelakeBmPlatformConfig The platform configuration of a bare metal instance that uses the BM.Standard3.64 shape or the
// BM.Optimized3.36 shape (the Intel Ice Lake platform).
type IntelIcelakeBmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`

	// Whether symmetric multithreading is enabled on the instance. Symmetric multithreading is also
	// called simultaneous multithreading (SMT) or Intel Hyper-Threading.
	// Intel and AMD processors have two hardware execution threads per core (OCPU). SMT permits multiple
	// independent threads of execution, to better use the resources and increase the efficiency
	// of the CPU. When multithreading is disabled, only one thread is permitted to run on each core, which
	// can provide higher or more predictable performance for some workloads.
	IsSymmetricMultiThreadingEnabled *bool `json:"isSymmetricMultiThreadingEnabled,omitempty"`

	// Whether the input-output memory management unit is enabled.
	IsInputOutputMemoryManagementUnitEnabled *bool `json:"isInputOutputMemoryManagementUnitEnabled,omitempty"`

	// The percentage of cores enabled. Value must be a multiple of 25%. If the requested percentage
	// results in a fractional number of cores, the system rounds up the number of cores across processors
	// and provisions an instance with a whole number of cores.
	// If the applications that you run on the instance use a core-based licensing model and need fewer cores
	// than the full size of the shape, you can disable cores to reduce your licensing costs. The instance
	// itself is billed for the full shape, regardless of whether all cores are enabled.
	PercentageOfCoresEnabled *int `json:"percentageOfCoresEnabled,omitempty"`

	// The number of NUMA nodes per socket (NPS).
	// The following values are supported:
	// * `NPS1`
	// * `NPS2`
	NumaNodesPerSocket IntelIcelakeBmPlatformConfigNumaNodesPerSocketEnum `json:"numaNodesPerSocket,omitempty"`
}

// IntelVmPlatformConfig The platform configuration of a virtual machine instance that uses the Intel platform.
type IntelVmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`
}

// AmdVmPlatformConfig The platform configuration of a virtual machine instance that uses the AMD platform.
type AmdVmPlatformConfig struct {
	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `json:"isMemoryEncryptionEnabled,omitempty"`
}

// InstanceOptions Optional mutable instance options
type InstanceOptions struct {

	// Whether to disable the legacy (/v1) instance metadata service endpoints.
	// Customers who have migrated to /v2 should set this to true for added security.
	// Default is false.
	AreLegacyImdsEndpointsDisabled *bool `json:"areLegacyImdsEndpointsDisabled,omitempty"`
}

// LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum Enum with underlying type: string
type LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum string

// Set of constants representing the allowable values for LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum
const (
	LaunchInstanceAvailabilityConfigDetailsRecoveryActionRestoreInstance LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum = "RESTORE_INSTANCE"
	LaunchInstanceAvailabilityConfigDetailsRecoveryActionStopInstance    LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum = "STOP_INSTANCE"
)

// LaunchInstanceAvailabilityConfig Options for VM migration during infrastructure maintenance events and for defining
// the availability of a VM instance after a maintenance event that impacts the underlying hardware.
type LaunchInstanceAvailabilityConfig struct {

	// IsLiveMigrationPreferred defines whether to live migrate supported VM instances to a healthy physical VM host without
	// disrupting running instances during infrastructure maintenance events. If null, Oracle
	// chooses the best option for migrating the VM during infrastructure maintenance events.
	IsLiveMigrationPreferred *bool `json:"isLiveMigrationPreferred,omitempty"`

	//RecoveryAction defines the lifecycle state for an instance when it is recovered after infrastructure maintenance.
	// * `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event.
	// If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
	// * `STOP_INSTANCE` - The instance is recovered in the stopped state.
	RecoveryAction LaunchInstanceAvailabilityConfigDetailsRecoveryActionEnum `json:"recoveryAction,omitempty"`
}

// PreemptibleInstanceConfig Configuration options for preemptible instances.
type PreemptibleInstanceConfig struct {
	// TerminatePreemptionAction terminates the preemptible instance when it is interrupted for eviction.
	TerminatePreemptionAction *TerminatePreemptionAction `json:"terminatePreemptionAction,omitempty"`
}

// TerminatePreemptionAction Terminates the preemptible instance when it is interrupted for eviction.
type TerminatePreemptionAction struct {

	// PreserveBootVolume defines whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified.
	PreserveBootVolume *bool `json:"preserveBootVolume,omitempty"`
}

// LaunchInstanceAgentConfig Configuration options for the Oracle Cloud Agent software running on the instance.
type LaunchInstanceAgentConfig struct {

	// IsMonitoringDisabled defines whether Oracle Cloud Agent can gather performance metrics and monitor the instance using the
	// monitoring plugins. Default value is false (monitoring plugins are enabled).
	// These are the monitoring plugins: Compute Instance Monitoring
	// and Custom Logs Monitoring.
	// The monitoring plugins are controlled by this parameter and by the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isMonitoringDisabled` is true, all of the monitoring plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isMonitoringDisabled` is false, all of the monitoring plugins are enabled. You
	// can optionally disable individual monitoring plugins by providing a value in the `pluginsConfig`
	// object.
	IsMonitoringDisabled *bool `json:"isMonitoringDisabled,omitempty"`

	// IsManagementDisabled defines whether Oracle Cloud Agent can run all the available management plugins.
	// Default value is false (management plugins are enabled).
	// These are the management plugins: OS Management Service Agent and Compute Instance
	// Run Command.
	// The management plugins are controlled by this parameter and by the per-plugin
	// configuration in the `pluginsConfig` object.
	// - If `isManagementDisabled` is true, all of the management plugins are disabled, regardless of
	// the per-plugin configuration.
	// - If `isManagementDisabled` is false, all of the management plugins are enabled. You
	// can optionally disable individual management plugins by providing a value in the `pluginsConfig`
	// object.
	IsManagementDisabled *bool `json:"isManagementDisabled,omitempty"`

	// AreAllPluginsDisabled defines whether Oracle Cloud Agent can run all the available plugins.
	// This includes the management and monitoring plugins.
	// To get a list of available plugins, use the
	// ListInstanceagentAvailablePlugins
	// operation in the Oracle Cloud Agent API. For more information about the available plugins, see
	// Managing Plugins with Oracle Cloud Agent (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
	AreAllPluginsDisabled *bool `json:"areAllPluginsDisabled,omitempty"`

	// PluginsConfig defines the configuration of plugins associated with this instance.
	PluginsConfig []InstanceAgentPluginConfig `json:"pluginsConfigs,omitempty"`
}

// InstanceAgentPluginConfigDetailsDesiredStateEnum Enum with underlying type: string
type InstanceAgentPluginConfigDetailsDesiredStateEnum string

// Set of constants representing the allowable values for InstanceAgentPluginConfigDetailsDesiredStateEnum
const (
	InstanceAgentPluginConfigDetailsDesiredStateEnabled  InstanceAgentPluginConfigDetailsDesiredStateEnum = "ENABLED"
	InstanceAgentPluginConfigDetailsDesiredStateDisabled InstanceAgentPluginConfigDetailsDesiredStateEnum = "DISABLED"
)

// InstanceAgentPluginConfig defines the configuration of plugins associated with this instance.
type InstanceAgentPluginConfig struct {

	// Name defines the name of the plugin. To get a list of available plugins, use the
	// ListInstanceagentAvailablePlugins
	// operation in the Oracle Cloud Agent API. For more information about the available plugins, see
	// Managing Plugins with Oracle Cloud Agent (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/manage-plugins.htm).
	Name *string `json:"name,omitempty"`

	// DesiredState defines whether the plugin should be enabled or disabled.
	// To enable the monitoring and management plugins, the `isMonitoringDisabled` and
	// `isManagementDisabled` attributes must also be set to false.
	// The following values are supported:
	// * `ENABLED`
	// * `DISABLED`
	DesiredState InstanceAgentPluginConfigDetailsDesiredStateEnum `json:"desiredState,omitempty"`
}

// ShapeConfig defines the configuration options for the compute instance shape
// https://docs.oracle.com/en-us/iaas/api/#/en/iaas/20160918/datatypes/LaunchInstanceShapeConfigDetails
type ShapeConfig struct {
	// The total number of OCPUs available to the instance.
	Ocpus string `json:"ocpus,omitempty"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs string `json:"memoryInGBs,omitempty"`

	// The baseline OCPU utilization for a subcore burstable VM instance. Leave this attribute blank for a
	// non-burstable instance, or explicitly specify non-burstable with `BASELINE_1_1`.
	// The following values are supported:
	// - `BASELINE_1_8` - baseline usage is 1/8 of an OCPU.
	// - `BASELINE_1_2` - baseline usage is 1/2 of an OCPU.
	// - `BASELINE_1_1` - baseline usage is an entire OCPU. This represents a non-burstable instance.
	BaselineOcpuUtilization string `json:"baselineOcpuUtilization,omitempty"`

	// Nvmes defines the number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
	Nvmes *int `json:"nvmes,omitempty"`
}

// EgressSecurityRule A rule for allowing outbound IP packets.
type EgressSecurityRule struct {

	// Conceptually, this is the range of IP addresses that a packet originating from the instance
	// can go to.
	// Allowed values:
	//   * IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`
	//     Note that IPv6 addressing is currently supported only in certain regions. See
	//     IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	//   * The `cidrBlock` value for a Service, if you're
	//     setting up a security list rule for traffic destined for a particular `Service` through
	//     a service gateway. For example: `oci-phx-objectstorage`.
	Destination *string `json:"destination,omitempty"`

	// The transport protocol. Specify either `all` or an IPv4 protocol number as
	// defined in
	// Protocol Numbers (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	Protocol *string `json:"protocol,omitempty"`

	// Type of destination for the rule. The default is `CIDR_BLOCK`.
	// Allowed values:
	//   * `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	//   * `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a
	//     Service (the rule is for traffic destined for a
	//     particular `Service` through a service gateway).
	//   * `NETWORK_SECURITY_GROUP`: If the rule's `destination` is the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a
	//     NetworkSecurityGroup.
	DestinationType EgressSecurityRuleDestinationTypeEnum `json:"destinationType,omitempty"`

	IcmpOptions *IcmpOptions `json:"icmpOptions,omitempty"`

	// A stateless rule allows traffic in one direction. Remember to add a corresponding
	// stateless rule in the other direction if you need to support bidirectional traffic. For
	// example, if egress traffic allows TCP destination port 80, there should be an ingress
	// rule to allow TCP source port 80. Defaults to false, which means the rule is stateful
	// and a corresponding rule is not necessary for bidirectional traffic.
	IsStateless *bool `json:"isStateless,omitempty"`

	TcpOptions *TcpOptions `json:"tcpOptions,omitempty"`

	UdpOptions *UdpOptions `json:"udpOptions,omitempty"`

	// An optional description of your choice for the rule.
	Description *string `json:"description,omitempty"`
}

// IngressSecurityRule A rule for allowing inbound IP packets.
type IngressSecurityRule struct {

	// The transport protocol. Specify either `all` or an IPv4 protocol number as
	// defined in
	// Protocol Numbers (http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	// Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	Protocol *string `json:"protocol,omitempty"`

	// Conceptually, this is the range of IP addresses that a packet coming into the instance
	// can come from.
	// Allowed values:
	//   * IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`.
	//     IPv6 addressing is supported for all commercial and government regions. See
	//     IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	//   * The `cidrBlock` value for a Service, if you're
	//     setting up a security list rule for traffic coming from a particular `Service` through
	//     a service gateway. For example: `oci-phx-objectstorage`.
	Source *string `json:"source,omitempty"`

	IcmpOptions *IcmpOptions `json:"icmpOptions,omitempty"`

	// A stateless rule allows traffic in one direction. Remember to add a corresponding
	// stateless rule in the other direction if you need to support bidirectional traffic. For
	// example, if ingress traffic allows TCP destination port 80, there should be an egress
	// rule to allow TCP source port 80. Defaults to false, which means the rule is stateful
	// and a corresponding rule is not necessary for bidirectional traffic.
	IsStateless *bool `json:"isStateless,omitempty"`

	// Type of source for the rule. The default is `CIDR_BLOCK`.
	//   * `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
	//   * `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a
	//     Service (the rule is for traffic coming from a
	//     particular `Service` through a service gateway).
	//   * `NETWORK_SECURITY_GROUP`: If the rule's `destination` is the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a
	//     NetworkSecurityGroup.
	SourceType IngressSecurityRuleSourceTypeEnum `json:"sourceType,omitempty"`

	TcpOptions *TcpOptions `json:"tcpOptions,omitempty"`

	UdpOptions *UdpOptions `json:"udpOptions,omitempty"`

	// An optional description of your choice for the rule.
	Description *string `json:"description,omitempty"`
}

// IngressSecurityRuleForNSG is IngressSecurityRule for NSG
type IngressSecurityRuleForNSG struct {
	IngressSecurityRule `json:"ingressRule,omitempty"`
}

// EgressSecurityRuleForNSG is EgressSecurityRule for NSG.
type EgressSecurityRuleForNSG struct {
	EgressSecurityRule `json:"egressRule,omitempty"`
}

// IngressSecurityRuleSourceTypeEnum Enum with underlying type: string.
type IngressSecurityRuleSourceTypeEnum string

// Set of constants representing the allowable values for IngressSecurityRuleSourceTypeEnum
const (
	IngressSecurityRuleSourceTypeCidrBlock        IngressSecurityRuleSourceTypeEnum = "CIDR_BLOCK"
	IngressSecurityRuleSourceTypeServiceCidrBlock IngressSecurityRuleSourceTypeEnum = "SERVICE_CIDR_BLOCK"
	IngressSecurityRuleSourceTypeNSG              IngressSecurityRuleSourceTypeEnum = "NETWORK_SECURITY_GROUP"
)

// UdpOptions Optional and valid only for UDP. Use to specify particular destination ports for UDP rules.
// If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
type UdpOptions struct {
	DestinationPortRange *PortRange `json:"destinationPortRange,omitempty"`

	SourcePortRange *PortRange `json:"sourcePortRange,omitempty"`
}

// IcmpOptions Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code
// as defined in:
// - ICMP Parameters (http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
// - ICMPv6 Parameters (https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)
// If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and
// codes are allowed. If you do provide this object, the type is required and the code is optional.
// To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination
// Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify
// multiple codes for a single type, create a separate security list rule for each.
type IcmpOptions struct {

	// The ICMP type.
	Type *int `json:"type,omitempty"`

	// The ICMP code (optional).
	Code *int `json:"code,omitempty"`
}

// TcpOptions Optional and valid only for TCP. Use to specify particular destination ports for TCP rules.
// If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
type TcpOptions struct {
	DestinationPortRange *PortRange `json:"destinationPortRange,omitempty"`

	SourcePortRange *PortRange `json:"sourcePortRange,omitempty"`
}

// PortRange The representation of PortRange.
type PortRange struct {

	// The maximum port number, which must not be less than the minimum port number. To specify
	// a single port number, set both the min and max to the same value.
	Max *int `json:"max,omitempty"`

	// The minimum port number, which must not be greater than the maximum port number.
	Min *int `json:"min,omitempty"`
}

const (
	// EgressSecurityRuleDestinationTypeCidrBlock is the contant for CIDR block security rule destination type
	EgressSecurityRuleDestinationTypeCidrBlock        EgressSecurityRuleDestinationTypeEnum = "CIDR_BLOCK"
	EgressSecurityRuleDestinationTypeServiceCidrBlock EgressSecurityRuleDestinationTypeEnum = "SERVICE_CIDR_BLOCK"
	EgressSecurityRuleDestinationTypeNSG              EgressSecurityRuleDestinationTypeEnum = "NETWORK_SECURITY_GROUP"
)

type EgressSecurityRuleDestinationTypeEnum string

// SecurityList defines the configureation for the security list for network virtual firewall
// https://docs.oracle.com/en-us/iaas/Content/Network/Concepts/securitylists.htm
type SecurityList struct {
	// ID of the SecurityList.
	// +optional
	ID *string `json:"id,omitempty"`
	// SecurityList Name.
	// +optional
	Name string `json:"name"`
	// EgressRules on the SecurityList.
	// +optional
	EgressRules []EgressSecurityRule `json:"egressRules,omitempty"`
	//IngressRules on the SecurityList.
	// +optional
	IngressRules []IngressSecurityRule `json:"ingressRules,omitempty"`
}

// Role defines the unique role of a subnet.
type Role string

type SubnetType string

// Subnet defines the configuration for a network's subnet
// https://docs.oracle.com/en-us/iaas/Content/Network/Tasks/managingVCNs_topic-Overview_of_VCNs_and_Subnets.htm#Overview
type Subnet struct {
	// Role defines the subnet role (eg. control-plane, control-plane-endpoint, service-lb, worker).
	Role Role `json:"role"`
	// Subnet OCID.
	// +optional
	ID *string `json:"id,omitempty"`
	// Subnet Name.
	Name string `json:"name"`
	// Skip specifies whether to skip creating subnets. If set to true(default: false) the ID
	// must be specified by the user to a valid Subnet ID.
	// +optional
	Skip bool `json:"skip,omitempty"`
	// Subnet CIDR.
	// +optional
	CIDR string `json:"cidr,omitempty"`
	// Type defines the subnet type (e.g. public, private).
	// +optional
	Type SubnetType `json:"type,omitempty"`
	// The security list associated with Subnet.
	// +optional
	SecurityList *SecurityList `json:"securityList,omitempty"`

	// DnsLabel DNS label for the subnet, used in conjunction with the VNIC's hostname and
	// VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// +optional
	DnsLabel *string `json:"dnsLabel,omitempty"`

	// Use this to enable IPv6 hextet for this subnet. The VCN must be enabled for IPv6.
	// You can't change this subnet characteristic later. All subnets are /64 in size. The subnet
	// portion of the IPv6 address is the fourth hextet from the left (1111 in the following example).
	// Example: `2001:0db8:0123:1111::/64`
	// +optional
	Ipv6CidrBlockHextet *string `json:"ipv6CidrBlockHextet,omitempty"`
}

// NSG defines configuration for a Network Security Group.
// https://docs.oracle.com/en-us/iaas/Content/Network/Concepts/networksecuritygroups.htm
type NSG struct {
	// NSG OCID.
	// +optional
	ID *string `json:"id,omitempty"`
	// NSG Name.
	Name string `json:"name"`
	// Role defines the NSG role (eg. control-plane, control-plane-endpoint, service-lb, worker).
	Role Role `json:"role,omitempty"`
	// EgressRules on the NSG.
	// +optional
	EgressRules []EgressSecurityRuleForNSG `json:"egressRules,omitempty"`
	// IngressRules on the NSG.
	// +optional
	IngressRules []IngressSecurityRuleForNSG `json:"ingressRules,omitempty"`
}

// VCN dfines the configuration for a Virtual Cloud Network.
// https://docs.oracle.com/en-us/iaas/Content/Network/Concepts/overview.htm
type VCN struct {
	// VCN OCID.
	// +optional
	ID *string `json:"id,omitempty"`
	// VCN Name.
	// +optional
	Name string `json:"name"`

	// Skip specifies whether to skip creating VCN.
	// When is set to true, InternetGateway, NatGateway, ServiceGateway, RouteTable must have also Skip set to true
	// If set to true(default: false) the ID must be specified by the user to a valid VCN ID.
	// +optional
	Skip bool `json:"skip,omitempty"`

	// VCN CIDR.
	// +optional
	// Deprecated, please use NetworkDetails.cidrs
	CIDR string `json:"cidr,omitempty"`

	// VCN CIDRs.
	// +optional
	CIDRS []string `json:"cidrs,omitempty"`

	// Subnets is the configuration for subnets required in the VCN.
	// +optional
	// +listType=map
	// +listMapKey=name
	Subnets []*Subnet `json:"subnets,omitempty"`

	// Configuration for Internet Gateway.
	// +optional
	InternetGateway InternetGateway `json:"internetGateway,omitempty"`

	// Configuration for NAT Gateway.
	// +optional
	NATGateway NATGateway `json:"natGateway,omitempty"`

	// Configuration for Service Gateway.
	// +optional
	ServiceGateway ServiceGateway `json:"serviceGateway,omitempty"`

	// Configuration for Route table.
	// +optional
	RouteTable RouteTable `json:"routeTable,omitempty"`

	// Configuration for NSG management.
	// +optional
	NetworkSecurityGroup NetworkSecurityGroup `json:"networkSecurityGroup,omitempty"`

	// DnsLabel specifies a DNS label for the VCN, used in conjunction with the VNIC's hostname and
	// subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance1.subnet123.vcn1.oraclevcn.com`).
	// +optional
	DnsLabel *string `json:"dnsLabel,omitempty"`

	// Configuration to allow OCI to assign IPv6 Prefix. When true will use a /56 IPv6 global unicast address (GUA) prefix allocated by Oracle
	// +optional
	IsOracleGuaAllocationEnabled *bool `json:"isOracleGuaAllocationEnabled,omitempty"`

	// Configuration to enable IPv6.
	// +optional
	IsIpv6Enabled *bool `json:"isIpv6Enabled,omitempty"`
}

// LoadBalancerType is an enumeration of the supported load balancer types.
type LoadBalancerType string

const (
	// NetworkLoadBalancer is the default load balancer type.
	LoadBalancerTypeNLB LoadBalancerType = LoadBalancerType("nlb")

	// LoadBalancer is the alternative load balancer type.
	LoadBalancerTypeLB LoadBalancerType = LoadBalancerType("lb")
)

// LoadBalancer Configuration
type LoadBalancer struct {
	//LoadBalancer Name.
	// +optional
	Name string `json:"name"`

	// ID of Load Balancer.
	// +optional
	LoadBalancerId *string `json:"loadBalancerId,omitempty"`

	// Type of Load Balancer: NLB (default) or LBaaS.
	// +optional
	LoadBalancerType LoadBalancerType `json:"loadBalancerType,omitempty"`

	// The NLB Spec
	// +optional
	NLBSpec NLBSpec `json:"nlbSpec,omitempty"`
}

// NLBSpec specifies the NLB spec.
type NLBSpec struct {
	// BackendSetDetails specifies the configuration of a network load balancer backend set.
	// +optional
	BackendSetDetails BackendSetDetails `json:"backendSetDetails,omitempty"`
}

// BackendSetDetails specifies the configuration of a network load balancer backend set.
type BackendSetDetails struct {
	// If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends.
	// Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled.
	// The value is false by default.
	// +optional
	IsPreserveSource *bool `json:"isPreserveSource,omitempty"`

	// If enabled, the network load balancer will continue to distribute traffic in the configured distribution in the event all backends are unhealthy.
	// The value is false by default.
	// +optional
	IsFailOpen *bool `json:"isFailOpen,omitempty"`

	// If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy.
	// +optional
	IsInstantFailoverEnabled *bool `json:"isInstantFailoverEnabled,omitempty"`

	// If enabled existing connections will be forwarded to an alternative healthy backend as soon as current backend becomes unhealthy.
	// +optional
	HealthChecker HealthChecker `json:"healthChecker,omitempty"`
}

// HealthChecker The health check policy configuration.
// For more information, see Editing Health Check Policies (https://docs.cloud.oracle.com/Content/NetworkLoadBalancer/HealthCheckPolicies/health-check-policy-management.htm).
type HealthChecker struct {

	// The path against which to run the health check.
	// Example: `/healthcheck`
	// Default value is `/healthz`
	UrlPath *string `json:"urlPath,omitempty"`
}

// NetworkSpec specifies what the OCI networking resources should look like.
type NetworkSpec struct {
	// SkipNetworkManagement defines if the networking spec(VCN related) specified by the user needs to be reconciled(actioned-upon)
	// or used as it is. APIServerLB will still be reconciled.
	// +optional
	SkipNetworkManagement bool `json:"skipNetworkManagement,omitempty"`

	// VCN configuration.
	// +optional
	Vcn VCN `json:"vcn,omitempty"`

	//API Server LB configuration.
	// +optional
	APIServerLB LoadBalancer `json:"apiServerLoadBalancer,omitempty"`

	// VCNPeering configuration.
	// +optional
	VCNPeering *VCNPeering `json:"vcnPeering,omitempty"`
}

// VCNPeering defines the VCN peering details of the workload cluster VCN.
type VCNPeering struct {

	// DRG configuration refers to the DRG which has to be created if required. If management cluster
	// and workload cluster shares the same DRG, this fields is not required to be specified.
	// +optional
	DRG *DRG `json:"drg,omitempty"`

	// PeerRouteRules defines the routing rules which will be added to the private route tables
	// of the workload cluster VCN. The routes defined here will be directed to DRG.
	PeerRouteRules []PeerRouteRule `json:"peerRouteRules,omitempty"`

	// RemotePeeringConnections defines the RPC connections which be established with the
	// workload cluster DRG.
	RemotePeeringConnections []RemotePeeringConnection `json:"remotePeeringConnections,omitempty"`
}

// DRG defines the configuration for a Dynamic Resource Group.
type DRG struct {

	// Manage defines whether the DRG has to be managed(including create). If set to false(the default) the ID
	// has to be specified by the user to a valid DRG ID to which the VCN has to be attached.
	// +optional
	Manage bool `json:"manage,omitempty"`

	// Name is the name of the created DRG.
	// +optional
	Name string `json:"name,omitempty"`

	// ID is the OCID for the created DRG.
	// +optional
	ID *string `json:"id,omitempty"`

	// VcnAttachmentId is the ID of the VCN attachment of the DRG.
	// The workload cluster VCN can be attached to either the management cluster VCN if they are sharing the same DRG
	// or to the workload cluster DRG.
	// +optional
	VcnAttachmentId *string `json:"vcnAttachmentId,omitempty"`
}

// PeerRouteRule defines a Route Rule to be routed via a DRG.
type PeerRouteRule struct {
	// VCNCIDRRange is the CIDR Range of peer VCN to which the
	// workload cluster VCN will be peered. The CIDR range is required to add the route rule
	// in the workload cluster VCN, the route rule will forward any traffic to the CIDR to the DRG.
	// +optional
	VCNCIDRRange string `json:"vcnCIDRRange,omitempty"`
}

// RemotePeeringConnection is used to peer VCNs residing in different regions(typically).
// Remote VCN Peering is explained here - https://docs.oracle.com/en-us/iaas/Content/Network/Tasks/remoteVCNpeering.htm
type RemotePeeringConnection struct {

	// ManagePeerRPC will define if the Peer VCN needs to be managed. If set to true
	// a Remote Peering Connection will be created in the Peer DRG and the connection
	// will be created between local and peer RPC.
	ManagePeerRPC bool `json:"managePeerRPC,omitempty"`

	// A unique name identifying the RPC, please note this is to identify the RPC
	// from other RPC elements, and will not be used in any OCI API call.
	Name string `json:"name,omitempty"`

	// PeerRegionName defined the region name of Peer VCN.
	PeerRegionName string `json:"peerRegionName,omitempty"`

	// PeerDRGId defines the DRG ID of the peer.
	PeerDRGId *string `json:"peerDRGId,omitempty"`

	// PeerRPCConnectionId defines the RPC ID of peer. If ManagePeerRPC is set to true
	// this will be created by Cluster API Provider for OCI, otherwise this has be defined by the
	// user.
	PeerRPCConnectionId *string `json:"peerRPCConnectionId,omitempty"`

	// RPCConnectionId is the connection ID of the connection between peer and local RPC.
	RPCConnectionId *string `json:"rpcConnectionId,omitempty"`
}

// InternetGateway is used to specify the options for creating internet gateway.
type InternetGateway struct {
	// Skip specifies whether to skip creating internet gateway even if any one Subnet is public.
	// In case of VCN being Skipped (Skip field of VCN set to true), this field must be true also
	// +optional
	Skip bool `json:"skip,omitempty"`

	// ID of Internet Gateway.
	// +optional
	Id *string `json:"id,omitempty"`
}

// NATGateway is used to specify the options for creating NAT gateway.
type NATGateway struct {
	// Skip specifies whether to skip creating NAT gateway even if any one Subnet is private.
	// In case of VCN being Skipped (Skip field of VCN set to true), this field must be true also
	// +optional
	Skip bool `json:"skip,omitempty"`

	// ID of Nat Gateway.
	// +optional
	Id *string `json:"id,omitempty"`
}

// ServiceGateway is used to specify the options for creating Service gateway.
type ServiceGateway struct {
	// Skip specifies whether to skip creating Service gateway.
	// In case of VCN being Skipped (Skip field of VCN set to true), this field must be true also
	// +optional
	Skip bool `json:"skip,omitempty"`

	// ID of Service Gateway.
	// +optional
	Id *string `json:"id,omitempty"`
}

// RouteTable is used to specify the options for creating Route table.
type RouteTable struct {
	// Skip specifies whether to skip creating Route table.
	// In case of VCN being Skipped (Skip field of VCN set to true), this field must be true also
	// +optional
	Skip bool `json:"skip,omitempty"`

	// ID of Private Route Table.
	// +optional
	PrivateRouteTableId *string `json:"privateRouteTableId,omitempty"`

	// ID of Public Route Table.
	// +optional
	PublicRouteTableId *string `json:"publicRouteTableId,omitempty"`
}

// NetworkSecurityGroup is used to specify the options for managing network security groups.
type NetworkSecurityGroup struct {
	// Skip specifies whether to skip creating network security groups.
	// +optional
	Skip bool `json:"skip,omitempty"`

	// NetworkSecurityGroup is the configuration for the Network Security Groups required in the VCN.
	// +optional
	// +listType=map
	// +listMapKey=name
	List []*NSG `json:"list,omitempty"`
}

type VolumeType string

const (
	IscsiType           VolumeType = "iscsi"
	ParavirtualizedType VolumeType = "paravirtualized"
)

// EncryptionInTransitTypeEnum Enum with underlying type: string
type EncryptionInTransitTypeEnum string

// Set of constants representing the allowable values for EncryptionInTransitTypeEnum
const (
	EncryptionInTransitTypeNone                  EncryptionInTransitTypeEnum = "NONE"
	EncryptionInTransitTypeBmEncryptionInTransit EncryptionInTransitTypeEnum = "BM_ENCRYPTION_IN_TRANSIT"
)

// LaunchVolumeAttachment specifies the details of the volume attachment.
type LaunchVolumeAttachment struct {
	// The type of volume. Valid value is iscsi.
	Type VolumeType `json:"volumeType,omitempty"`

	// The details of iscsi volume attachment.
	IscsiAttachment           LaunchIscsiVolumeAttachment           `json:"launchIscsiVolumeAttachment,omitempty"`
	ParavirtualizedAttachment LaunchParavirtualizedVolumeAttachment `json:"launchParavirtualizedVolumeAttachment,omitempty"`
}

// LaunchIscsiVolumeAttachment specifies the iscsi volume attachments to create as part of the launch instance operation.
type LaunchIscsiVolumeAttachment struct {
	// The device name. To retrieve a list of devices for a given instance, see ListInstanceDevices.
	Device *string `json:"device,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `json:"isShareable,omitempty"`

	// The OCID of the volume. If CreateVolumeDetails is specified, this field must be omitted from the request.
	VolumeId *string `json:"volumeId,omitempty"`

	// LaunchCreateVolumeFromAttributes The details of the volume to create for CreateVolume operation.
	LaunchCreateVolumeFromAttributes LaunchCreateVolumeFromAttributes `json:"launchCreateVolumeFromAttributes,omitempty"`

	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap *bool `json:"useChap,omitempty"`

	// Whether to enable Oracle Cloud Agent to perform the iSCSI login and logout commands after the volume attach or detach operations for non multipath-enabled iSCSI attachments.
	IsAgentAutoIscsiLoginEnabled *bool `json:"isAgentAutoIscsiLoginEnabled,omitempty"`

	// Refer the top-level definition of encryptionInTransitType.
	// The default value is NONE.
	EncryptionInTransitType EncryptionInTransitTypeEnum `json:"encryptionInTransitType,omitempty"`
}

// LaunchParavirtualizedVolumeAttachment specifies the paravirtualized volume attachments to create as part of the launch instance operation.
type LaunchParavirtualizedVolumeAttachment struct {
	// The device name. To retrieve a list of devices for a given instance, see ListInstanceDevices.
	Device *string `json:"device,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `json:"isShareable,omitempty"`

	// The OCID of the volume. If CreateVolumeDetails is specified, this field must be omitted from the request.
	VolumeId *string `json:"volumeId,omitempty"`

	// LaunchCreateVolumeFromAttributes The details of the volume to create for CreateVolume operation.
	LaunchCreateVolumeFromAttributes LaunchCreateVolumeFromAttributes `json:"launchCreateVolumeFromAttributes,omitempty"`

	// Refer the top-level definition of encryptionInTransitType.
	// The default value is False.
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty"`
}

// LaunchCreateVolumeFromAttributes The details of the volume to create for CreateVolume operation.
type LaunchCreateVolumeFromAttributes struct {

	// The size of the volume in GBs.
	SizeInGBs *int64 `json:"sizeInGBs,omitempty"`

	// The OCID of the compartment that contains the volume. If not provided,
	// it will be inherited from the instance.
	CompartmentId *string `json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty"`

	// The OCID of the Vault service key to assign as the master encryption key
	// for the volume.
	KmsKeyId *string `json:"kmsKeyId,omitempty"`

	// The number of volume performance units (VPUs) that will be applied to this volume per GB,
	// representing the Block Volume service's elastic performance options.
	// See Block Volume Performance Levels (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeperformance.htm#perf_levels) for more information.
	// Allowed values:
	//   * `0`: Represents Lower Cost option.
	//   * `10`: Represents Balanced option.
	//   * `20`: Represents Higher Performance option.
	//   * `30`-`120`: Represents the Ultra High Performance option.
	VpusPerGB *int64 `json:"vpusPerGB,omitempty"`
}

const (
	VCNNativeCNI CNIOptionEnum = "OCI_VCN_IP_NATIVE"
	FlannelCNI   CNIOptionEnum = "FLANNEL_OVERLAY"
)

type CNIOptionEnum string
