// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A connectivity test are a static analysis of your resource configurations
// that enables you to evaluate connectivity to and from Google Cloud
// resources in your Virtual Private Cloud (VPC) network.
//
// To get more information about ConnectivityTest, see:
//
// * [API documentation](https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/reference/networkmanagement/rest/v1/projects.locations.global.connectivityTests)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/network-intelligence-center/docs)
//
// ## Example Usage
// ### Network Management Connectivity Test Instances
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/networkmanagement"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vpc, err := compute.NewNetwork(ctx, "vpc", nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		debian9, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		source, err := compute.NewInstance(ctx, "source", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(debian9.Id),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: vpc.ID(),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		destination, err := compute.NewInstance(ctx, "destination", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(debian9.Id),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: vpc.ID(),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = networkmanagement.NewConnectivityTest(ctx, "instance_test", &networkmanagement.ConnectivityTestArgs{
// 			Source: &networkmanagement.ConnectivityTestSourceArgs{
// 				Instance: source.ID(),
// 			},
// 			Destination: &networkmanagement.ConnectivityTestDestinationArgs{
// 				Instance: destination.ID(),
// 			},
// 			Protocol: pulumi.String("TCP"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Network Management Connectivity Test Addresses
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/networkmanagement"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vpc, err := compute.NewNetwork(ctx, "vpc", nil)
// 		if err != nil {
// 			return err
// 		}
// 		subnet, err := compute.NewSubnetwork(ctx, "subnet", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.0.0/16"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewAddress(ctx, "source_addr", &compute.AddressArgs{
// 			Subnetwork:  subnet.ID(),
// 			AddressType: pulumi.String("INTERNAL"),
// 			Address:     pulumi.String("10.0.42.42"),
// 			Region:      pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewAddress(ctx, "dest_addr", &compute.AddressArgs{
// 			Subnetwork:  subnet.ID(),
// 			AddressType: pulumi.String("INTERNAL"),
// 			Address:     pulumi.String("10.0.43.43"),
// 			Region:      pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = networkmanagement.NewConnectivityTest(ctx, "address_test", &networkmanagement.ConnectivityTestArgs{
// 			Source: &networkmanagement.ConnectivityTestSourceArgs{
// 				IpAddress:   source_addr.Address,
// 				ProjectId:   source_addr.Project,
// 				Network:     vpc.ID(),
// 				NetworkType: pulumi.String("GCP_NETWORK"),
// 			},
// 			Destination: &networkmanagement.ConnectivityTestDestinationArgs{
// 				IpAddress: dest_addr.Address,
// 				ProjectId: dest_addr.Project,
// 				Network:   vpc.ID(),
// 			},
// 			Protocol: pulumi.String("UDP"),
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
// ConnectivityTest can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default projects/{{project}}/locations/global/connectivityTests/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:networkmanagement/connectivityTest:ConnectivityTest default {{name}}
// ```
type ConnectivityTest struct {
	pulumi.CustomResourceState

	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination ConnectivityTestDestinationOutput `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayOutput `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source ConnectivityTestSourceOutput `pulumi:"source"`
}

// NewConnectivityTest registers a new resource with the given unique name, arguments, and options.
func NewConnectivityTest(ctx *pulumi.Context,
	name string, args *ConnectivityTestArgs, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource ConnectivityTest
	err := ctx.RegisterResource("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectivityTest gets an existing ConnectivityTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectivityTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityTestState, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	var resource ConnectivityTest
	err := ctx.ReadResource("gcp:networkmanagement/connectivityTest:ConnectivityTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectivityTest resources.
type connectivityTestState struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination *ConnectivityTestDestination `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source *ConnectivityTestSource `pulumi:"source"`
}

type ConnectivityTestState struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination ConnectivityTestDestinationPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Unique name for the connectivity test.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source ConnectivityTestSourcePtrInput
}

func (ConnectivityTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestState)(nil)).Elem()
}

type connectivityTestArgs struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination ConnectivityTestDestination `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique name for the connectivity test.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source ConnectivityTestSource `pulumi:"source"`
}

// The set of arguments for constructing a ConnectivityTest resource.
type ConnectivityTestArgs struct {
	// The user-supplied description of the Connectivity Test.
	// Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test.
	// You can use a combination of destination IP address, Compute
	// Engine VM instance, or VPC network to uniquely identify the
	// destination location.
	// Even if the destination IP address is not unique, the source IP
	// location is unique. Usually, the analysis can infer the destination
	// endpoint from route information.
	// If the destination you specify is a VM instance and the instance has
	// multiple network interfaces, then you must also specify either a
	// destination IP address or VPC network to identify the destination
	// interface.
	// A reachability analysis proceeds even if the destination location
	// is ambiguous. However, the result can include endpoints that you
	// don't intend to test.
	// Structure is documented below.
	Destination ConnectivityTestDestinationInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Unique name for the connectivity test.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// Other projects that may be relevant for reachability analysis.
	// This is applicable to scenarios where a test can cross project
	// boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test.
	// You can use a combination of source IP address, virtual machine
	// (VM) instance, or Compute Engine network to uniquely identify the
	// source location.
	// Examples: If the source IP address is an internal IP address within
	// a Google Cloud Virtual Private Cloud (VPC) network, then you must
	// also specify the VPC network. Otherwise, specify the VM instance,
	// which already contains its internal IP address and VPC network
	// information.
	// If the source of the test is within an on-premises network, then
	// you must provide the destination VPC network.
	// If the source endpoint is a Compute Engine VM instance with multiple
	// network interfaces, the instance itself is not sufficient to
	// identify the endpoint. So, you must also specify the source IP
	// address or VPC network.
	// A reachability analysis proceeds even if the source location is
	// ambiguous. However, the test result may include endpoints that
	// you don't intend to test.
	// Structure is documented below.
	Source ConnectivityTestSourceInput
}

func (ConnectivityTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestArgs)(nil)).Elem()
}

type ConnectivityTestInput interface {
	pulumi.Input

	ToConnectivityTestOutput() ConnectivityTestOutput
	ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput
}

func (*ConnectivityTest) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTest)(nil))
}

func (i *ConnectivityTest) ToConnectivityTestOutput() ConnectivityTestOutput {
	return i.ToConnectivityTestOutputWithContext(context.Background())
}

func (i *ConnectivityTest) ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityTestOutput)
}

type ConnectivityTestOutput struct {
	*pulumi.OutputState
}

func (ConnectivityTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTest)(nil))
}

func (o ConnectivityTestOutput) ToConnectivityTestOutput() ConnectivityTestOutput {
	return o
}

func (o ConnectivityTestOutput) ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectivityTestOutput{})
}
