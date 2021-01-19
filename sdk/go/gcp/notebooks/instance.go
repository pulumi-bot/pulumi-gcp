// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Cloud AI Platform Notebook instance.
//
// > **Note:** Due to limitations of the Notebooks Instance API, many fields
// in this resource do not properly detect drift. These fields will also not
// appear in state once imported.
//
// To get more information about Instance, see:
//
// * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
//
// ## Example Usage
// ### Notebook Instance Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstance(ctx, "instance", &notebooks.InstanceArgs{
// 			Location:    pulumi.String("us-west1-a"),
// 			MachineType: pulumi.String("e2-medium"),
// 			VmImage: &notebooks.InstanceVmImageArgs{
// 				ImageFamily: pulumi.String("tf-latest-cpu"),
// 				Project:     pulumi.String("deeplearning-platform-release"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Notebook Instance Basic Container
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstance(ctx, "instance", &notebooks.InstanceArgs{
// 			ContainerImage: &notebooks.InstanceContainerImageArgs{
// 				Repository: pulumi.String("gcr.io/deeplearning-platform-release/base-cpu"),
// 				Tag:        pulumi.String("latest"),
// 			},
// 			Location:    pulumi.String("us-west1-a"),
// 			MachineType: pulumi.String("e2-medium"),
// 			Metadata: pulumi.StringMap{
// 				"proxy-mode": pulumi.String("service_account"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Notebook Instance Basic Gpu
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstance(ctx, "instance", &notebooks.InstanceArgs{
// 			AcceleratorConfig: &notebooks.InstanceAcceleratorConfigArgs{
// 				CoreCount: pulumi.Int(1),
// 				Type:      pulumi.String("NVIDIA_TESLA_T4"),
// 			},
// 			InstallGpuDriver: pulumi.Bool(true),
// 			Location:         pulumi.String("us-west1-a"),
// 			MachineType:      pulumi.String("n1-standard-1"),
// 			VmImage: &notebooks.InstanceVmImageArgs{
// 				ImageFamily: pulumi.String("tf-latest-gpu"),
// 				Project:     pulumi.String("deeplearning-platform-release"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Notebook Instance Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myNetwork, err := compute.LookupNetwork(ctx, &compute.LookupNetworkArgs{
// 			Name: "default",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "default"
// 		opt1 := "us-central1"
// 		mySubnetwork, err := compute.LookupSubnetwork(ctx, &compute.LookupSubnetworkArgs{
// 			Name:   &opt0,
// 			Region: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = notebooks.NewInstance(ctx, "instance", &notebooks.InstanceArgs{
// 			Location:    pulumi.String("us-central1-a"),
// 			MachineType: pulumi.String("e2-medium"),
// 			VmImage: &notebooks.InstanceVmImageArgs{
// 				Project:     pulumi.String("deeplearning-platform-release"),
// 				ImageFamily: pulumi.String("tf-latest-cpu"),
// 			},
// 			InstanceOwners: pulumi.String(pulumi.String{
// 				pulumi.String("admin@hashicorptest.com"),
// 			}),
// 			ServiceAccount:   pulumi.String("emailAddress:my@service-account.com"),
// 			InstallGpuDriver: pulumi.Bool(true),
// 			BootDiskType:     pulumi.String("PD_SSD"),
// 			BootDiskSizeGb:   pulumi.Int(110),
// 			NoPublicIp:       pulumi.Bool(true),
// 			NoProxyAccess:    pulumi.Bool(true),
// 			Network:          pulumi.String(myNetwork.Id),
// 			Subnet:           pulumi.String(mySubnetwork.Id),
// 			Labels: pulumi.StringMap{
// 				"k": pulumi.String("val"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Instance can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:notebooks/instance:Instance default projects/{{project}}/locations/{{location}}/instances/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/instance:Instance default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:notebooks/instance:Instance default {{location}}/{{name}}
// ```
type Instance struct {
	pulumi.CustomResourceState

	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	AcceleratorConfig InstanceAcceleratorConfigPtrOutput `pulumi:"acceleratorConfig"`
	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	BootDiskSizeGb pulumi.IntPtrOutput `pulumi:"bootDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	BootDiskType pulumi.StringPtrOutput `pulumi:"bootDiskType"`
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage InstanceContainerImagePtrOutput `pulumi:"containerImage"`
	// Instance creation time
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath pulumi.StringPtrOutput `pulumi:"customGpuDriverPath"`
	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	DataDiskSizeGb pulumi.IntPtrOutput `pulumi:"dataDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	DataDiskType pulumi.StringPtrOutput `pulumi:"dataDiskType"`
	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, and `CMEK`.
	DiskEncryption pulumi.StringPtrOutput `pulumi:"diskEncryption"`
	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver pulumi.BoolPtrOutput `pulumi:"installGpuDriver"`
	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	InstanceOwners pulumi.StringPtrOutput `pulumi:"instanceOwners"`
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	KmsKey pulumi.StringPtrOutput `pulumi:"kmsKey"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A reference to the zone where the machine resides.
	Location pulumi.StringOutput `pulumi:"location"`
	// A reference to a machine type which defines VM kind.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name specified for the Notebook instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	Network pulumi.StringOutput `pulumi:"network"`
	// the notebook instance will not register with the proxy..
	NoProxyAccess pulumi.BoolPtrOutput `pulumi:"noProxyAccess"`
	// no public IP will be assigned to this instance.
	NoPublicIp pulumi.BoolPtrOutput `pulumi:"noPublicIp"`
	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk pulumi.BoolPtrOutput `pulumi:"noRemoveDataDisk"`
	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	PostStartupScript pulumi.StringPtrOutput `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringOutput `pulumi:"project"`
	// The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri pulumi.StringOutput `pulumi:"proxyUri"`
	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The state of this instance.
	State pulumi.StringOutput `pulumi:"state"`
	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Instance update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage InstanceVmImagePtrOutput `pulumi:"vmImage"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MachineType == nil {
		return nil, errors.New("invalid value for required argument 'MachineType'")
	}
	var resource Instance
	err := ctx.RegisterResource("gcp:notebooks/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("gcp:notebooks/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	AcceleratorConfig *InstanceAcceleratorConfig `pulumi:"acceleratorConfig"`
	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	BootDiskSizeGb *int `pulumi:"bootDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	BootDiskType *string `pulumi:"bootDiskType"`
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *InstanceContainerImage `pulumi:"containerImage"`
	// Instance creation time
	CreateTime *string `pulumi:"createTime"`
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath *string `pulumi:"customGpuDriverPath"`
	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	DataDiskSizeGb *int `pulumi:"dataDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	DataDiskType *string `pulumi:"dataDiskType"`
	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, and `CMEK`.
	DiskEncryption *string `pulumi:"diskEncryption"`
	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver *bool `pulumi:"installGpuDriver"`
	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	InstanceOwners *string `pulumi:"instanceOwners"`
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	KmsKey *string `pulumi:"kmsKey"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A reference to the zone where the machine resides.
	Location *string `pulumi:"location"`
	// A reference to a machine type which defines VM kind.
	MachineType *string `pulumi:"machineType"`
	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Metadata map[string]string `pulumi:"metadata"`
	// The name specified for the Notebook instance.
	Name *string `pulumi:"name"`
	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	Network *string `pulumi:"network"`
	// the notebook instance will not register with the proxy..
	NoProxyAccess *bool `pulumi:"noProxyAccess"`
	// no public IP will be assigned to this instance.
	NoPublicIp *bool `pulumi:"noPublicIp"`
	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk *bool `pulumi:"noRemoveDataDisk"`
	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri *string `pulumi:"proxyUri"`
	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The state of this instance.
	State *string `pulumi:"state"`
	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	Subnet *string `pulumi:"subnet"`
	// Instance update time.
	UpdateTime *string `pulumi:"updateTime"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *InstanceVmImage `pulumi:"vmImage"`
}

type InstanceState struct {
	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	AcceleratorConfig InstanceAcceleratorConfigPtrInput
	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	BootDiskSizeGb pulumi.IntPtrInput
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	BootDiskType pulumi.StringPtrInput
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage InstanceContainerImagePtrInput
	// Instance creation time
	CreateTime pulumi.StringPtrInput
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath pulumi.StringPtrInput
	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	DataDiskSizeGb pulumi.IntPtrInput
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	DataDiskType pulumi.StringPtrInput
	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, and `CMEK`.
	DiskEncryption pulumi.StringPtrInput
	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver pulumi.BoolPtrInput
	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	InstanceOwners pulumi.StringPtrInput
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	KmsKey pulumi.StringPtrInput
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringPtrInput
	// A reference to a machine type which defines VM kind.
	MachineType pulumi.StringPtrInput
	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Metadata pulumi.StringMapInput
	// The name specified for the Notebook instance.
	Name pulumi.StringPtrInput
	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	Network pulumi.StringPtrInput
	// the notebook instance will not register with the proxy..
	NoProxyAccess pulumi.BoolPtrInput
	// no public IP will be assigned to this instance.
	NoPublicIp pulumi.BoolPtrInput
	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk pulumi.BoolPtrInput
	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri pulumi.StringPtrInput
	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	ServiceAccount pulumi.StringPtrInput
	// The state of this instance.
	State pulumi.StringPtrInput
	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	Subnet pulumi.StringPtrInput
	// Instance update time.
	UpdateTime pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage InstanceVmImagePtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	AcceleratorConfig *InstanceAcceleratorConfig `pulumi:"acceleratorConfig"`
	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	BootDiskSizeGb *int `pulumi:"bootDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	BootDiskType *string `pulumi:"bootDiskType"`
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage *InstanceContainerImage `pulumi:"containerImage"`
	// Instance creation time
	CreateTime *string `pulumi:"createTime"`
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath *string `pulumi:"customGpuDriverPath"`
	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	DataDiskSizeGb *int `pulumi:"dataDiskSizeGb"`
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	DataDiskType *string `pulumi:"dataDiskType"`
	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, and `CMEK`.
	DiskEncryption *string `pulumi:"diskEncryption"`
	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver *bool `pulumi:"installGpuDriver"`
	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	InstanceOwners *string `pulumi:"instanceOwners"`
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	KmsKey *string `pulumi:"kmsKey"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// A reference to the zone where the machine resides.
	Location string `pulumi:"location"`
	// A reference to a machine type which defines VM kind.
	MachineType string `pulumi:"machineType"`
	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Metadata map[string]string `pulumi:"metadata"`
	// The name specified for the Notebook instance.
	Name *string `pulumi:"name"`
	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	Network *string `pulumi:"network"`
	// the notebook instance will not register with the proxy..
	NoProxyAccess *bool `pulumi:"noProxyAccess"`
	// no public IP will be assigned to this instance.
	NoPublicIp *bool `pulumi:"noPublicIp"`
	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk *bool `pulumi:"noRemoveDataDisk"`
	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	PostStartupScript *string `pulumi:"postStartupScript"`
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project *string `pulumi:"project"`
	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	Subnet *string `pulumi:"subnet"`
	// Instance update time.
	UpdateTime *string `pulumi:"updateTime"`
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage *InstanceVmImage `pulumi:"vmImage"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The hardware accelerator used on this instance. If you use accelerators,
	// make sure that your configuration has enough vCPUs and memory to support the
	// machineType you have selected.
	// Structure is documented below.
	AcceleratorConfig InstanceAcceleratorConfigPtrInput
	// The size of the boot disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB.
	// If not specified, this defaults to 100.
	BootDiskSizeGb pulumi.IntPtrInput
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	BootDiskType pulumi.StringPtrInput
	// Use a container image to start the notebook instance.
	// Structure is documented below.
	ContainerImage InstanceContainerImagePtrInput
	// Instance creation time
	CreateTime pulumi.StringPtrInput
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	// If not specified, we'll automatically choose from official GPU drivers.
	CustomGpuDriverPath pulumi.StringPtrInput
	// The size of the data disk in GB attached to this instance,
	// up to a maximum of 64000 GB (64 TB).
	// You can choose the size of the data disk based on how big your notebooks and data are.
	// If not specified, this defaults to 100.
	DataDiskSizeGb pulumi.IntPtrInput
	// Possible disk types for notebook instances.
	// Possible values are `DISK_TYPE_UNSPECIFIED`, `PD_STANDARD`, `PD_SSD`, and `PD_BALANCED`.
	DataDiskType pulumi.StringPtrInput
	// Disk encryption method used on the boot and data disks, defaults to GMEK.
	// Possible values are `DISK_ENCRYPTION_UNSPECIFIED`, `GMEK`, and `CMEK`.
	DiskEncryption pulumi.StringPtrInput
	// Whether the end user authorizes Google Cloud to install GPU driver
	// on this instance. If this field is empty or set to false, the GPU driver
	// won't be installed. Only applicable to instances with GPUs.
	InstallGpuDriver pulumi.BoolPtrInput
	// The list of owners of this instance after creation.
	// Format: alias@example.com.
	// Currently supports one owner only.
	// If not specified, all of the service account users of
	// your VM instance's service account can use the instance.
	InstanceOwners pulumi.StringPtrInput
	// The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK.
	// Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}
	KmsKey pulumi.StringPtrInput
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// A reference to the zone where the machine resides.
	Location pulumi.StringInput
	// A reference to a machine type which defines VM kind.
	MachineType pulumi.StringInput
	// Custom metadata to apply to this instance.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Metadata pulumi.StringMapInput
	// The name specified for the Notebook instance.
	Name pulumi.StringPtrInput
	// The name of the VPC that this instance is in.
	// Format: projects/{project_id}/global/networks/{network_id}
	Network pulumi.StringPtrInput
	// the notebook instance will not register with the proxy..
	NoProxyAccess pulumi.BoolPtrInput
	// no public IP will be assigned to this instance.
	NoPublicIp pulumi.BoolPtrInput
	// If true, the data disk will not be auto deleted when deleting the instance.
	NoRemoveDataDisk pulumi.BoolPtrInput
	// Path to a Bash script that automatically runs after a
	// notebook instance fully boots up. The path must be a URL
	// or Cloud Storage path (gs://path-to-file/file-name).
	PostStartupScript pulumi.StringPtrInput
	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	Project pulumi.StringPtrInput
	// The service account on this instance, giving access to other
	// Google Cloud services. You can use any service account within
	// the same project, but you must have the service account user
	// permission to use the instance. If not specified,
	// the Compute Engine default service account is used.
	ServiceAccount pulumi.StringPtrInput
	// The name of the subnet that this instance is in.
	// Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}
	Subnet pulumi.StringPtrInput
	// Instance update time.
	UpdateTime pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	// Structure is documented below.
	VmImage InstanceVmImagePtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

func (i *Instance) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *Instance) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

type InstancePtrInput interface {
	pulumi.Input

	ToInstancePtrOutput() InstancePtrOutput
	ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput
}

type instancePtrType InstanceArgs

func (*instancePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (i *instancePtrType) ToInstancePtrOutput() InstancePtrOutput {
	return i.ToInstancePtrOutputWithContext(context.Background())
}

func (i *instancePtrType) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePtrOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//          InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Instance)(nil))
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//          InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Instance)(nil))
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct {
	*pulumi.OutputState
}

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Instance)(nil))
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o.ToInstancePtrOutputWithContext(context.Background())
}

func (o InstanceOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o.ApplyT(func(v Instance) *Instance {
		return &v
	}).(InstancePtrOutput)
}

type InstancePtrOutput struct {
	*pulumi.OutputState
}

func (InstancePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil))
}

func (o InstancePtrOutput) ToInstancePtrOutput() InstancePtrOutput {
	return o
}

func (o InstancePtrOutput) ToInstancePtrOutputWithContext(ctx context.Context) InstancePtrOutput {
	return o
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Instance)(nil))
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Instance {
		return vs[0].([]Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Instance)(nil))
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Instance {
		return vs[0].(map[string]Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstancePtrOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
