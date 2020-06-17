// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a VPN gateway running in GCP. This virtual device is managed
// by Google, but used only by you.
//
//
// To get more information about VpnGateway, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways)
//
// ## Example Usage
//
// ### Target Vpn Gateway Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		network1, err := compute.NewNetwork(ctx, "network1", nil)
// 		if err != nil {
// 			return err
// 		}
// 		targetGateway, err := compute.NewVPNGateway(ctx, "targetGateway", &compute.VPNGatewayArgs{
// 			Network: network1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vpnStaticIp, err := compute.NewAddress(ctx, "vpnStaticIp", nil)
// 		if err != nil {
// 			return err
// 		}
// 		frEsp, err := compute.NewForwardingRule(ctx, "frEsp", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("ESP"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		frUdp500, err := compute.NewForwardingRule(ctx, "frUdp500", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("UDP"),
// 			PortRange:  pulumi.String("500"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		frUdp4500, err := compute.NewForwardingRule(ctx, "frUdp4500", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("UDP"),
// 			PortRange:  pulumi.String("4500"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tunnel1, err := compute.NewVPNTunnel(ctx, "tunnel1", &compute.VPNTunnelArgs{
// 			PeerIp:           pulumi.String("15.0.0.120"),
// 			SharedSecret:     pulumi.String("a secret message"),
// 			TargetVpnGateway: targetGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRoute(ctx, "route1", &compute.RouteArgs{
// 			Network:          network1.Name,
// 			DestRange:        pulumi.String("15.0.0.0/24"),
// 			Priority:         pulumi.Int(1000),
// 			NextHopVpnTunnel: tunnel1.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VPNGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The unique identifier for the resource.
	GatewayId pulumi.IntOutput `pulumi:"gatewayId"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region this gateway should sit in.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewVPNGateway registers a new resource with the given unique name, arguments, and options.
func NewVPNGateway(ctx *pulumi.Context,
	name string, args *VPNGatewayArgs, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	if args == nil {
		args = &VPNGatewayArgs{}
	}
	var resource VPNGateway
	err := ctx.RegisterResource("gcp:compute/vPNGateway:VPNGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPNGateway gets an existing VPNGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPNGatewayState, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	var resource VPNGateway
	err := ctx.ReadResource("gcp:compute/vPNGateway:VPNGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPNGateway resources.
type vpngatewayState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The unique identifier for the resource.
	GatewayId *int `pulumi:"gatewayId"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type VPNGatewayState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The unique identifier for the resource.
	GatewayId pulumi.IntPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (VPNGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayState)(nil)).Elem()
}

type vpngatewayArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this VPN gateway is accepting traffic for.
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region this gateway should sit in.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a VPNGateway resource.
type VPNGatewayArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this VPN gateway is accepting traffic for.
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region this gateway should sit in.
	Region pulumi.StringPtrInput
}

func (VPNGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayArgs)(nil)).Elem()
}
