// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Cloud Router interface. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/routers).
type RouterInterface struct {
	pulumi.CustomResourceState

	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
	// specified.
	InterconnectAttachment pulumi.StringPtrOutput `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringPtrOutput `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which this interface's router belongs. If it
	// is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region this interface's router sits in. If not specified,
	// the project region will be used. Changing this forces a new interface to be
	// created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router pulumi.StringOutput `pulumi:"router"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel` and `interconnectAttachment` can be specified.
	VpnTunnel pulumi.StringPtrOutput `pulumi:"vpnTunnel"`
}

// NewRouterInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterInterface(ctx *pulumi.Context,
	name string, args *RouterInterfaceArgs, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Router == nil {
		return nil, errors.New("invalid value for required argument 'Router'")
	}
	var resource RouterInterface
	err := ctx.RegisterResource("gcp:compute/routerInterface:RouterInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterInterface gets an existing RouterInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterInterfaceState, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	var resource RouterInterface
	err := ctx.ReadResource("gcp:compute/routerInterface:RouterInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterInterface resources.
type routerInterfaceState struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
	// specified.
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange *string `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which this interface's router belongs. If it
	// is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `pulumi:"project"`
	// The region this interface's router sits in. If not specified,
	// the project region will be used. Changing this forces a new interface to be
	// created.
	Region *string `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router *string `pulumi:"router"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel` and `interconnectAttachment` can be specified.
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

type RouterInterfaceState struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
	// specified.
	InterconnectAttachment pulumi.StringPtrInput
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringPtrInput
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which this interface's router belongs. If it
	// is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringPtrInput
	// The region this interface's router sits in. If not specified,
	// the project region will be used. Changing this forces a new interface to be
	// created.
	Region pulumi.StringPtrInput
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router pulumi.StringPtrInput
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel` and `interconnectAttachment` can be specified.
	VpnTunnel pulumi.StringPtrInput
}

func (RouterInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceState)(nil)).Elem()
}

type routerInterfaceArgs struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
	// specified.
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange *string `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which this interface's router belongs. If it
	// is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `pulumi:"project"`
	// The region this interface's router sits in. If not specified,
	// the project region will be used. Changing this forces a new interface to be
	// created.
	Region *string `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router string `pulumi:"router"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel` and `interconnectAttachment` can be specified.
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

// The set of arguments for constructing a RouterInterface resource.
type RouterInterfaceArgs struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel` and `interconnectAttachment` can be
	// specified.
	InterconnectAttachment pulumi.StringPtrInput
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringPtrInput
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which this interface's router belongs. If it
	// is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringPtrInput
	// The region this interface's router sits in. If not specified,
	// the project region will be used. Changing this forces a new interface to be
	// created.
	Region pulumi.StringPtrInput
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	Router pulumi.StringInput
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel` and `interconnectAttachment` can be specified.
	VpnTunnel pulumi.StringPtrInput
}

func (RouterInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceArgs)(nil)).Elem()
}

type RouterInterfaceInput interface {
	pulumi.Input

	ToRouterInterfaceOutput() RouterInterfaceOutput
	ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput
}

func (RouterInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterInterface)(nil)).Elem()
}

func (i RouterInterface) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return i.ToRouterInterfaceOutputWithContext(context.Background())
}

func (i RouterInterface) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceOutput)
}

type RouterInterfaceOutput struct {
	*pulumi.OutputState
}

func (RouterInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterInterfaceOutput)(nil)).Elem()
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return o
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouterInterfaceOutput{})
}
