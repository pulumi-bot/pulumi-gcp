// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a VPN gateway managed outside of GCP.
//
// To get more information about ExternalVpnGateway, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)
//
// ## Example Usage
// ### External Vpn Gateway
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
// 			RoutingMode:           pulumi.String("GLOBAL"),
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		haGateway, err := compute.NewHaVpnGateway(ctx, "haGateway", &compute.HaVpnGatewayArgs{
// 			Region:  pulumi.String("us-central1"),
// 			Network: network.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		externalGateway, err := compute.NewExternalVpnGateway(ctx, "externalGateway", &compute.ExternalVpnGatewayArgs{
// 			RedundancyType: pulumi.String("SINGLE_IP_INTERNALLY_REDUNDANT"),
// 			Description:    pulumi.String("An externally managed VPN gateway"),
// 			Interfaces: compute.ExternalVpnGatewayInterfaceArray{
// 				&compute.ExternalVpnGatewayInterfaceArgs{
// 					Id:        pulumi.Int(0),
// 					IpAddress: pulumi.String("8.8.8.8"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "networkSubnet1", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.1.0/24"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     network.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "networkSubnet2", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.2.0/24"),
// 			Region:      pulumi.String("us-west1"),
// 			Network:     network.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1, err := compute.NewRouter(ctx, "router1", &compute.RouterArgs{
// 			Network: network.Name,
// 			Bgp: &compute.RouterBgpArgs{
// 				Asn: pulumi.Int(64514),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel1, err := compute.NewVPNTunnel(ctx, "tunnel1", &compute.VPNTunnelArgs{
// 			Region:                       pulumi.String("us-central1"),
// 			VpnGateway:                   haGateway.ID(),
// 			PeerExternalGateway:          externalGateway.ID(),
// 			PeerExternalGatewayInterface: pulumi.Int(0),
// 			SharedSecret:                 pulumi.String("a secret message"),
// 			Router:                       router1.ID(),
// 			VpnGatewayInterface:          pulumi.Int(0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel2, err := compute.NewVPNTunnel(ctx, "tunnel2", &compute.VPNTunnelArgs{
// 			Region:                       pulumi.String("us-central1"),
// 			VpnGateway:                   haGateway.ID(),
// 			PeerExternalGateway:          externalGateway.ID(),
// 			PeerExternalGatewayInterface: pulumi.Int(0),
// 			SharedSecret:                 pulumi.String("a secret message"),
// 			Router: router1.ID().ApplyT(func(id string) (string, error) {
// 				return fmt.Sprintf("%v%v", " ", id), nil
// 			}).(pulumi.StringOutput),
// 			VpnGatewayInterface: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1Interface1, err := compute.NewRouterInterface(ctx, "router1Interface1", &compute.RouterInterfaceArgs{
// 			Router:    router1.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.0.1/30"),
// 			VpnTunnel: tunnel1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router1Peer1", &compute.RouterPeerArgs{
// 			Router:                  router1.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.0.2"),
// 			PeerAsn:                 pulumi.Int(64515),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router1Interface1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		router1Interface2, err := compute.NewRouterInterface(ctx, "router1Interface2", &compute.RouterInterfaceArgs{
// 			Router:    router1.Name,
// 			Region:    pulumi.String("us-central1"),
// 			IpRange:   pulumi.String("169.254.1.1/30"),
// 			VpnTunnel: tunnel2.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRouterPeer(ctx, "router1Peer2", &compute.RouterPeerArgs{
// 			Router:                  router1.Name,
// 			Region:                  pulumi.String("us-central1"),
// 			PeerIpAddress:           pulumi.String("169.254.1.2"),
// 			PeerAsn:                 pulumi.Int(64515),
// 			AdvertisedRoutePriority: pulumi.Int(100),
// 			Interface:               router1Interface2.Name,
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
// ExternalVpnGateway can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
// ```
type ExternalVpnGateway struct {
	pulumi.CustomResourceState

	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayOutput `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrOutput `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ExternalVpnGatewayArgs, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	if args == nil {
		args = &ExternalVpnGatewayArgs{}
	}

	var resource ExternalVpnGateway
	err := ctx.RegisterResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalVpnGateway gets an existing ExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalVpnGatewayState, opts ...pulumi.ResourceOption) (*ExternalVpnGateway, error) {
	var resource ExternalVpnGateway
	err := ctx.ReadResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalVpnGateway resources.
type externalVpnGatewayState struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
	RedundancyType *string `pulumi:"redundancyType"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type ExternalVpnGatewayState struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (ExternalVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayState)(nil)).Elem()
}

type externalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces []ExternalVpnGatewayInterface `pulumi:"interfaces"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
	RedundancyType *string `pulumi:"redundancyType"`
}

// The set of arguments for constructing a ExternalVpnGateway resource.
type ExternalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// A list of interfaces on this external VPN gateway.
	// Structure is documented below.
	Interfaces ExternalVpnGatewayInterfaceArrayInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates the redundancy type of this external VPN gateway
	// Possible values are `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, and `TWO_IPS_REDUNDANCY`.
	RedundancyType pulumi.StringPtrInput
}

func (ExternalVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalVpnGatewayArgs)(nil)).Elem()
}

type ExternalVpnGatewayInput interface {
	pulumi.Input

	ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput
	ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput
}

func (*ExternalVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalVpnGateway)(nil))
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return i.ToExternalVpnGatewayOutputWithContext(context.Background())
}

func (i *ExternalVpnGateway) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalVpnGatewayOutput)
}

type ExternalVpnGatewayOutput struct {
	*pulumi.OutputState
}

func (ExternalVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalVpnGateway)(nil))
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutput() ExternalVpnGatewayOutput {
	return o
}

func (o ExternalVpnGatewayOutput) ToExternalVpnGatewayOutputWithContext(ctx context.Context) ExternalVpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExternalVpnGatewayOutput{})
}
