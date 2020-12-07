// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Network endpoint groups (NEGs) are zonal resources that represent
// collections of IP address and port combinations for GCP resources within a
// single subnet. Each IP address and port combination is called a network
// endpoint.
//
// Network endpoint groups can be used as backends in backend services for
// HTTP(S), TCP proxy, and SSL proxy load balancers. You cannot use NEGs as a
// backend with internal load balancers. Because NEG backends allow you to
// specify IP addresses and ports, you can distribute traffic in a granular
// fashion among applications or containers running within VM instances.
//
// Recreating a network endpoint group that's in use by another resource will give a
// `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
// to avoid this type of error.
//
// To get more information about NetworkEndpointGroup, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
//
// ## Example Usage
// ### Network Endpoint Group
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		defaultSubnetwork, err := compute.NewSubnetwork(ctx, "defaultSubnetwork", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.0.0/16"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     defaultNetwork.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkEndpointGroup(ctx, "neg", &compute.NetworkEndpointGroupArgs{
// 			Network:     defaultNetwork.ID(),
// 			Subnetwork:  defaultSubnetwork.ID(),
// 			DefaultPort: pulumi.Int(90),
// 			Zone:        pulumi.String("us-central1-a"),
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
// NetworkEndpointGroup can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{project}}/{{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{zone}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/networkEndpointGroup:NetworkEndpointGroup default {{name}}
// ```
type NetworkEndpointGroup struct {
	pulumi.CustomResourceState

	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort pulumi.IntPtrOutput `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	Network pulumi.StringOutput `pulumi:"network"`
	// Type of network endpoints in this network endpoint group.
	// Default value is `GCE_VM_IP_PORT`.
	// Possible values are `GCE_VM_IP_PORT`.
	NetworkEndpointType pulumi.StringPtrOutput `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Number of network endpoints in the network endpoint group.
	Size pulumi.IntOutput `pulumi:"size"`
	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrOutput `pulumi:"subnetwork"`
	// Zone where the network endpoint group is located.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNetworkEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkEndpointGroup(ctx *pulumi.Context,
	name string, args *NetworkEndpointGroupArgs, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	var resource NetworkEndpointGroup
	err := ctx.RegisterResource("gcp:compute/networkEndpointGroup:NetworkEndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkEndpointGroup gets an existing NetworkEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkEndpointGroupState, opts ...pulumi.ResourceOption) (*NetworkEndpointGroup, error) {
	var resource NetworkEndpointGroup
	err := ctx.ReadResource("gcp:compute/networkEndpointGroup:NetworkEndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkEndpointGroup resources.
type networkEndpointGroupState struct {
	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	Network *string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group.
	// Default value is `GCE_VM_IP_PORT`.
	// Possible values are `GCE_VM_IP_PORT`.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Number of network endpoints in the network endpoint group.
	Size *int `pulumi:"size"`
	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
	// Zone where the network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

type NetworkEndpointGroupState struct {
	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	Network pulumi.StringPtrInput
	// Type of network endpoints in this network endpoint group.
	// Default value is `GCE_VM_IP_PORT`.
	// Possible values are `GCE_VM_IP_PORT`.
	NetworkEndpointType pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Number of network endpoints in the network endpoint group.
	Size pulumi.IntPtrInput
	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
	// Zone where the network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (NetworkEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupState)(nil)).Elem()
}

type networkEndpointGroupArgs struct {
	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort *int `pulumi:"defaultPort"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `pulumi:"description"`
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	Network string `pulumi:"network"`
	// Type of network endpoints in this network endpoint group.
	// Default value is `GCE_VM_IP_PORT`.
	// Possible values are `GCE_VM_IP_PORT`.
	NetworkEndpointType *string `pulumi:"networkEndpointType"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork *string `pulumi:"subnetwork"`
	// Zone where the network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a NetworkEndpointGroup resource.
type NetworkEndpointGroupArgs struct {
	// The default port used if the port number is not specified in the
	// network endpoint.
	DefaultPort pulumi.IntPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description pulumi.StringPtrInput
	// Name of the resource; provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network to which all network endpoints in the NEG belong.
	// Uses "default" project network if unspecified.
	Network pulumi.StringInput
	// Type of network endpoints in this network endpoint group.
	// Default value is `GCE_VM_IP_PORT`.
	// Possible values are `GCE_VM_IP_PORT`.
	NetworkEndpointType pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork pulumi.StringPtrInput
	// Zone where the network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (NetworkEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkEndpointGroupArgs)(nil)).Elem()
}

type NetworkEndpointGroupInput interface {
	pulumi.Input

	ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput
	ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput
}

func (NetworkEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpointGroup)(nil)).Elem()
}

func (i NetworkEndpointGroup) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return i.ToNetworkEndpointGroupOutputWithContext(context.Background())
}

func (i NetworkEndpointGroup) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkEndpointGroupOutput)
}

type NetworkEndpointGroupOutput struct {
	*pulumi.OutputState
}

func (NetworkEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkEndpointGroupOutput)(nil)).Elem()
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutput() NetworkEndpointGroupOutput {
	return o
}

func (o NetworkEndpointGroupOutput) ToNetworkEndpointGroupOutputWithContext(ctx context.Context) NetworkEndpointGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkEndpointGroupOutput{})
}
