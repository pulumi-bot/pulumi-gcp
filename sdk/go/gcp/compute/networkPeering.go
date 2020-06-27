// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a network peering within GCE. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/networks).
//
// > Both network must create a peering with each other for the peering
// to be functional.
//
// > Subnets IP ranges across peered VPC networks cannot overlap.
//
// ## Example Usage
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
// 		_default, err := compute.NewNetwork(ctx, "_default", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		other, err := compute.NewNetwork(ctx, "other", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkPeering(ctx, "peering1", &compute.NetworkPeeringArgs{
// 			Network:     _default.ID(),
// 			PeerNetwork: other.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetworkPeering(ctx, "peering2", &compute.NetworkPeeringArgs{
// 			Network:     other.ID(),
// 			PeerNetwork: _default.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkPeering struct {
	pulumi.CustomResourceState

	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrOutput `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to export the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrOutput `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrOutput `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary network of the peering.
	Network pulumi.StringOutput `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringOutput `pulumi:"peerNetwork"`
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State pulumi.StringOutput `pulumi:"state"`
	// Details about the current state of the peering.
	StateDetails pulumi.StringOutput `pulumi:"stateDetails"`
}

// NewNetworkPeering registers a new resource with the given unique name, arguments, and options.
func NewNetworkPeering(ctx *pulumi.Context,
	name string, args *NetworkPeeringArgs, opts ...pulumi.ResourceOption) (*NetworkPeering, error) {
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	if args == nil || args.PeerNetwork == nil {
		return nil, errors.New("missing required argument 'PeerNetwork'")
	}
	if args == nil {
		args = &NetworkPeeringArgs{}
	}
	var resource NetworkPeering
	err := ctx.RegisterResource("gcp:compute/networkPeering:NetworkPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPeering gets an existing NetworkPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPeeringState, opts ...pulumi.ResourceOption) (*NetworkPeering, error) {
	var resource NetworkPeering
	err := ctx.ReadResource("gcp:compute/networkPeering:NetworkPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPeering resources.
type networkPeeringState struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp *bool `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to export the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp *bool `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name *string `pulumi:"name"`
	// The primary network of the peering.
	Network *string `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork *string `pulumi:"peerNetwork"`
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State *string `pulumi:"state"`
	// Details about the current state of the peering.
	StateDetails *string `pulumi:"stateDetails"`
}

type NetworkPeeringState struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Whether to export the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Name of the peering.
	Name pulumi.StringPtrInput
	// The primary network of the peering.
	Network pulumi.StringPtrInput
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringPtrInput
	// State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
	// `ACTIVE` when there's a matching configuration in the peer network.
	State pulumi.StringPtrInput
	// Details about the current state of the peering.
	StateDetails pulumi.StringPtrInput
}

func (NetworkPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringState)(nil)).Elem()
}

type networkPeeringArgs struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes *bool `pulumi:"exportCustomRoutes"`
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp *bool `pulumi:"exportSubnetRoutesWithPublicIp"`
	// Whether to export the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes *bool `pulumi:"importCustomRoutes"`
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp *bool `pulumi:"importSubnetRoutesWithPublicIp"`
	// Name of the peering.
	Name *string `pulumi:"name"`
	// The primary network of the peering.
	Network string `pulumi:"network"`
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork string `pulumi:"peerNetwork"`
}

// The set of arguments for constructing a NetworkPeering resource.
type NetworkPeeringArgs struct {
	// Whether to export the custom routes to the peer network. Defaults to `false`.
	ExportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
	ExportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Whether to export the custom routes from the peer network. Defaults to `false`.
	ImportCustomRoutes pulumi.BoolPtrInput
	// Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
	ImportSubnetRoutesWithPublicIp pulumi.BoolPtrInput
	// Name of the peering.
	Name pulumi.StringPtrInput
	// The primary network of the peering.
	Network pulumi.StringInput
	// The peer network in the peering. The peer network
	// may belong to a different project.
	PeerNetwork pulumi.StringInput
}

func (NetworkPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPeeringArgs)(nil)).Elem()
}
