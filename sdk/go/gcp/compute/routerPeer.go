// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// BGP information that must be configured into the routing stack to
// establish BGP peering. This information must specify the peer ASN
// and either the interface name, IP address, or peer IP address.
// Please refer to RFC4273.
//
// To get more information about RouterBgpPeer, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
// * How-to Guides
//     * [Google Cloud Router](https://cloud.google.com/router/docs/)
//
// ## Example Usage
type RouterPeer struct {
	pulumi.CustomResourceState

	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `CUSTOM`.
	AdvertiseMode pulumi.StringPtrOutput `pulumi:"advertiseMode"`
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
	// * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
	// * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
	AdvertisedGroups pulumi.StringArrayOutput `pulumi:"advertisedGroups"`
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is `CUSTOM` and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIpRanges RouterPeerAdvertisedIpRangeArrayOutput `pulumi:"advertisedIpRanges"`
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority pulumi.IntPtrOutput `pulumi:"advertisedRoutePriority"`
	// Name of the interface the BGP peer is associated with.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The resource that configures and manages this BGP peer. * 'MANAGED_BY_USER' is the default value and can be managed by
	// you or other users * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and managed by Cloud Interconnect,
	// specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type
	// of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
	ManagementType pulumi.StringOutput `pulumi:"managementType"`
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn pulumi.IntOutput `pulumi:"peerAsn"`
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIpAddress pulumi.StringOutput `pulumi:"peerIpAddress"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router pulumi.StringOutput `pulumi:"router"`
}

// NewRouterPeer registers a new resource with the given unique name, arguments, and options.
func NewRouterPeer(ctx *pulumi.Context,
	name string, args *RouterPeerArgs, opts ...pulumi.ResourceOption) (*RouterPeer, error) {
	if args == nil || args.Interface == nil {
		return nil, errors.New("missing required argument 'Interface'")
	}
	if args == nil || args.PeerAsn == nil {
		return nil, errors.New("missing required argument 'PeerAsn'")
	}
	if args == nil || args.PeerIpAddress == nil {
		return nil, errors.New("missing required argument 'PeerIpAddress'")
	}
	if args == nil || args.Router == nil {
		return nil, errors.New("missing required argument 'Router'")
	}
	if args == nil {
		args = &RouterPeerArgs{}
	}
	var resource RouterPeer
	err := ctx.RegisterResource("gcp:compute/routerPeer:RouterPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterPeer gets an existing RouterPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterPeerState, opts ...pulumi.ResourceOption) (*RouterPeer, error) {
	var resource RouterPeer
	err := ctx.ReadResource("gcp:compute/routerPeer:RouterPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterPeer resources.
type routerPeerState struct {
	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `CUSTOM`.
	AdvertiseMode *string `pulumi:"advertiseMode"`
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
	// * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
	// * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
	AdvertisedGroups []string `pulumi:"advertisedGroups"`
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is `CUSTOM` and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIpRanges []RouterPeerAdvertisedIpRange `pulumi:"advertisedIpRanges"`
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority *int `pulumi:"advertisedRoutePriority"`
	// Name of the interface the BGP peer is associated with.
	Interface *string `pulumi:"interface"`
	// IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
	IpAddress *string `pulumi:"ipAddress"`
	// The resource that configures and manages this BGP peer. * 'MANAGED_BY_USER' is the default value and can be managed by
	// you or other users * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and managed by Cloud Interconnect,
	// specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type
	// of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
	ManagementType *string `pulumi:"managementType"`
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn *int `pulumi:"peerAsn"`
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIpAddress *string `pulumi:"peerIpAddress"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router *string `pulumi:"router"`
}

type RouterPeerState struct {
	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `CUSTOM`.
	AdvertiseMode pulumi.StringPtrInput
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
	// * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
	// * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
	AdvertisedGroups pulumi.StringArrayInput
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is `CUSTOM` and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIpRanges RouterPeerAdvertisedIpRangeArrayInput
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority pulumi.IntPtrInput
	// Name of the interface the BGP peer is associated with.
	Interface pulumi.StringPtrInput
	// IP address of the interface inside Google Cloud Platform. Only IPv4 is supported.
	IpAddress pulumi.StringPtrInput
	// The resource that configures and manages this BGP peer. * 'MANAGED_BY_USER' is the default value and can be managed by
	// you or other users * 'MANAGED_BY_ATTACHMENT' is a BGP peer that is configured and managed by Cloud Interconnect,
	// specifically by an InterconnectAttachment of type PARTNER. Google automatically creates, updates, and deletes this type
	// of BGP peer when the PARTNER InterconnectAttachment is created, updated, or deleted.
	ManagementType pulumi.StringPtrInput
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn pulumi.IntPtrInput
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIpAddress pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router pulumi.StringPtrInput
}

func (RouterPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPeerState)(nil)).Elem()
}

type routerPeerArgs struct {
	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `CUSTOM`.
	AdvertiseMode *string `pulumi:"advertiseMode"`
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
	// * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
	// * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
	AdvertisedGroups []string `pulumi:"advertisedGroups"`
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is `CUSTOM` and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIpRanges []RouterPeerAdvertisedIpRange `pulumi:"advertisedIpRanges"`
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority *int `pulumi:"advertisedRoutePriority"`
	// Name of the interface the BGP peer is associated with.
	Interface string `pulumi:"interface"`
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn int `pulumi:"peerAsn"`
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIpAddress string `pulumi:"peerIpAddress"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router string `pulumi:"router"`
}

// The set of arguments for constructing a RouterPeer resource.
type RouterPeerArgs struct {
	// User-specified flag to indicate which mode to use for advertisement.
	// Valid values of this enum field are: `DEFAULT`, `CUSTOM`
	// Default value is `DEFAULT`.
	// Possible values are `DEFAULT` and `CUSTOM`.
	AdvertiseMode pulumi.StringPtrInput
	// User-specified list of prefix groups to advertise in custom
	// mode, which can take one of the following options:
	// * `ALL_SUBNETS`: Advertises all available subnets, including peer VPC subnets.
	// * `ALL_VPC_SUBNETS`: Advertises the router's own VPC subnets.
	// * `ALL_PEER_VPC_SUBNETS`: Advertises peer subnets of the router's VPC network.
	AdvertisedGroups pulumi.StringArrayInput
	// User-specified list of individual IP ranges to advertise in
	// custom mode. This field can only be populated if advertiseMode
	// is `CUSTOM` and is advertised to all peers of the router. These IP
	// ranges will be advertised in addition to any specified groups.
	// Leave this field blank to advertise no custom IP ranges.
	// Structure is documented below.
	AdvertisedIpRanges RouterPeerAdvertisedIpRangeArrayInput
	// The priority of routes advertised to this BGP peer.
	// Where there is more than one matching route of maximum
	// length, the routes with the lowest priority value win.
	AdvertisedRoutePriority pulumi.IntPtrInput
	// Name of the interface the BGP peer is associated with.
	Interface pulumi.StringInput
	// Name of this BGP peer. The name must be 1-63 characters long,
	// and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Peer BGP Autonomous System Number (ASN).
	// Each BGP interface may use a different value.
	PeerAsn pulumi.IntInput
	// IP address of the BGP interface outside Google Cloud Platform.
	// Only IPv4 is supported.
	PeerIpAddress pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router and BgpPeer reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The name of the Cloud Router in which this BgpPeer will be configured.
	Router pulumi.StringInput
}

func (RouterPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerPeerArgs)(nil)).Elem()
}

type RouterPeerInput interface {
	pulumi.Input

	ToRouterPeerOutput() RouterPeerOutput
	ToRouterPeerOutputWithContext(ctx context.Context) RouterPeerOutput
}

func (RouterPeer) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterPeer)(nil)).Elem()
}

func (i RouterPeer) ToRouterPeerOutput() RouterPeerOutput {
	return i.ToRouterPeerOutputWithContext(context.Background())
}

func (i RouterPeer) ToRouterPeerOutputWithContext(ctx context.Context) RouterPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterPeerOutput)
}

type RouterPeerOutput struct {
	*pulumi.OutputState
}

func (RouterPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouterPeerOutput)(nil)).Elem()
}

func (o RouterPeerOutput) ToRouterPeerOutput() RouterPeerOutput {
	return o
}

func (o RouterPeerOutput) ToRouterPeerOutputWithContext(ctx context.Context) RouterPeerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouterPeerOutput{})
}
