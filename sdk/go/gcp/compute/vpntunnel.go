// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VPN tunnel resource.
//
// To get more information about VpnTunnel, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
// * How-to Guides
//     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
//     * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)
//
// > **Warning:** All arguments including `sharedSecret` will be stored in the raw
// state as plain-text.
//
// ## Example Usage
// ### Vpn Tunnel Basic
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
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			frEsp,
// 			frUdp500,
// 			frUdp4500,
// 		}))
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
// ### Vpn Tunnel Beta
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
// 		network1, err := compute.NewNetwork(ctx, "network1", nil, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		targetGateway, err := compute.NewVPNGateway(ctx, "targetGateway", &compute.VPNGatewayArgs{
// 			Network: network1.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		vpnStaticIp, err := compute.NewAddress(ctx, "vpnStaticIp", nil, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		frEsp, err := compute.NewForwardingRule(ctx, "frEsp", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("ESP"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		frUdp500, err := compute.NewForwardingRule(ctx, "frUdp500", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("UDP"),
// 			PortRange:  pulumi.String("500"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		frUdp4500, err := compute.NewForwardingRule(ctx, "frUdp4500", &compute.ForwardingRuleArgs{
// 			IpProtocol: pulumi.String("UDP"),
// 			PortRange:  pulumi.String("4500"),
// 			IpAddress:  vpnStaticIp.Address,
// 			Target:     targetGateway.ID(),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		tunnel1, err := compute.NewVPNTunnel(ctx, "tunnel1", &compute.VPNTunnelArgs{
// 			PeerIp:           pulumi.String("15.0.0.120"),
// 			SharedSecret:     pulumi.String("a secret message"),
// 			TargetVpnGateway: targetGateway.ID(),
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		}, pulumi.Provider(google_beta), pulumi.DependsOn([]pulumi.Resource{
// 			frEsp,
// 			frUdp500,
// 			frUdp4500,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewRoute(ctx, "route1", &compute.RouteArgs{
// 			Network:          network1.Name,
// 			DestRange:        pulumi.String("15.0.0.0/24"),
// 			Priority:         pulumi.Int(1000),
// 			NextHopVpnTunnel: tunnel1.ID(),
// 		}, pulumi.Provider(google_beta))
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
// VpnTunnel can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{name}}
// ```
type VPNTunnel struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion pulumi.IntPtrOutput `pulumi:"ikeVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this VpnTunnel.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayOutput `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// `a-z?` which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrOutput `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.IntPtrOutput `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
	// ID in the peer GCP VPN gateway.
	// This field must reference a `compute.HaVpnGateway` resource.
	PeerGcpGateway pulumi.StringPtrOutput `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp pulumi.StringOutput `pulumi:"peerIp"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
	Region pulumi.StringOutput `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayOutput `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router pulumi.StringPtrOutput `pulumi:"router"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SharedSecret pulumi.StringOutput `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringOutput `pulumi:"sharedSecretHash"`
	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVpnGateway pulumi.StringPtrOutput `pulumi:"targetVpnGateway"`
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId pulumi.StringOutput `pulumi:"tunnelId"`
	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a `compute.HaVpnGateway` resource.
	VpnGateway pulumi.StringPtrOutput `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.IntPtrOutput `pulumi:"vpnGatewayInterface"`
}

// NewVPNTunnel registers a new resource with the given unique name, arguments, and options.
func NewVPNTunnel(ctx *pulumi.Context,
	name string, args *VPNTunnelArgs, opts ...pulumi.ResourceOption) (*VPNTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SharedSecret == nil {
		return nil, errors.New("invalid value for required argument 'SharedSecret'")
	}
	var resource VPNTunnel
	err := ctx.RegisterResource("gcp:compute/vPNTunnel:VPNTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPNTunnel gets an existing VPNTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPNTunnelState, opts ...pulumi.ResourceOption) (*VPNTunnel, error) {
	var resource VPNTunnel
	err := ctx.ReadResource("gcp:compute/vPNTunnel:VPNTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPNTunnel resources.
type vpntunnelState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Detailed status message for the VPN tunnel.
	DetailedStatus *string `pulumi:"detailedStatus"`
	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion *int `pulumi:"ikeVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this VpnTunnel.
	Labels map[string]string `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelectors []string `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// `a-z?` which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway *string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *int `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
	// ID in the peer GCP VPN gateway.
	// This field must reference a `compute.HaVpnGateway` resource.
	PeerGcpGateway *string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp *string `pulumi:"peerIp"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
	Region *string `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelectors []string `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router *string `pulumi:"router"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SharedSecret *string `pulumi:"sharedSecret"`
	// Hash of the shared secret.
	SharedSecretHash *string `pulumi:"sharedSecretHash"`
	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVpnGateway *string `pulumi:"targetVpnGateway"`
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId *string `pulumi:"tunnelId"`
	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a `compute.HaVpnGateway` resource.
	VpnGateway *string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface *int `pulumi:"vpnGatewayInterface"`
}

type VPNTunnelState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Detailed status message for the VPN tunnel.
	DetailedStatus pulumi.StringPtrInput
	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion pulumi.IntPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this VpnTunnel.
	Labels pulumi.StringMapInput
	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// `a-z?` which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrInput
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.IntPtrInput
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
	// ID in the peer GCP VPN gateway.
	// This field must reference a `compute.HaVpnGateway` resource.
	PeerGcpGateway pulumi.StringPtrInput
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
	Region pulumi.StringPtrInput
	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayInput
	// URL of router resource to be used for dynamic routing.
	Router pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SharedSecret pulumi.StringPtrInput
	// Hash of the shared secret.
	SharedSecretHash pulumi.StringPtrInput
	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVpnGateway pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	TunnelId pulumi.StringPtrInput
	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a `compute.HaVpnGateway` resource.
	VpnGateway pulumi.StringPtrInput
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.IntPtrInput
}

func (VPNTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpntunnelState)(nil)).Elem()
}

type vpntunnelArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion *int `pulumi:"ikeVersion"`
	// Labels to apply to this VpnTunnel.
	Labels map[string]string `pulumi:"labels"`
	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelectors []string `pulumi:"localTrafficSelectors"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// `a-z?` which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway *string `pulumi:"peerExternalGateway"`
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface *int `pulumi:"peerExternalGatewayInterface"`
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
	// ID in the peer GCP VPN gateway.
	// This field must reference a `compute.HaVpnGateway` resource.
	PeerGcpGateway *string `pulumi:"peerGcpGateway"`
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp *string `pulumi:"peerIp"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
	Region *string `pulumi:"region"`
	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelectors []string `pulumi:"remoteTrafficSelectors"`
	// URL of router resource to be used for dynamic routing.
	Router *string `pulumi:"router"`
	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SharedSecret string `pulumi:"sharedSecret"`
	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVpnGateway *string `pulumi:"targetVpnGateway"`
	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a `compute.HaVpnGateway` resource.
	VpnGateway *string `pulumi:"vpnGateway"`
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface *int `pulumi:"vpnGatewayInterface"`
}

// The set of arguments for constructing a VPNTunnel resource.
type VPNTunnelArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	IkeVersion pulumi.IntPtrInput
	// Labels to apply to this VpnTunnel.
	Labels pulumi.StringMapInput
	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	LocalTrafficSelectors pulumi.StringArrayInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// `a-z?` which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGateway pulumi.StringPtrInput
	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	PeerExternalGatewayInterface pulumi.IntPtrInput
	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpnGatewayInterface
	// ID in the peer GCP VPN gateway.
	// This field must reference a `compute.HaVpnGateway` resource.
	PeerGcpGateway pulumi.StringPtrInput
	// IP address of the peer VPN gateway. Only IPv4 is supported.
	PeerIp pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region where the tunnel is located. If unset, is set to the region of `targetVpnGateway`.
	Region pulumi.StringPtrInput
	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example `192.168.0.0/16`. The ranges should be disjoint.
	// Only IPv4 is supported.
	RemoteTrafficSelectors pulumi.StringArrayInput
	// URL of router resource to be used for dynamic routing.
	Router pulumi.StringPtrInput
	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	SharedSecret pulumi.StringInput
	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	TargetVpnGateway pulumi.StringPtrInput
	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a `compute.HaVpnGateway` resource.
	VpnGateway pulumi.StringPtrInput
	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	VpnGatewayInterface pulumi.IntPtrInput
}

func (VPNTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpntunnelArgs)(nil)).Elem()
}

type VPNTunnelInput interface {
	pulumi.Input

	ToVPNTunnelOutput() VPNTunnelOutput
	ToVPNTunnelOutputWithContext(ctx context.Context) VPNTunnelOutput
}

func (VPNTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNTunnel)(nil)).Elem()
}

func (i VPNTunnel) ToVPNTunnelOutput() VPNTunnelOutput {
	return i.ToVPNTunnelOutputWithContext(context.Background())
}

func (i VPNTunnel) ToVPNTunnelOutputWithContext(ctx context.Context) VPNTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNTunnelOutput)
}

type VPNTunnelOutput struct {
	*pulumi.OutputState
}

func (VPNTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNTunnelOutput)(nil)).Elem()
}

func (o VPNTunnelOutput) ToVPNTunnelOutput() VPNTunnelOutput {
	return o
}

func (o VPNTunnelOutput) ToVPNTunnelOutputWithContext(ctx context.Context) VPNTunnelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VPNTunnelOutput{})
}
