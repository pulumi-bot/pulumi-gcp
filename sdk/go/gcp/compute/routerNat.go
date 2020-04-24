// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A NAT service created in a router.
//
//
// To get more information about RouterNat, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
// * How-to Guides
//     * [Google Cloud Router](https://cloud.google.com/router/docs/)
type RouterNat struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNatIps pulumi.StringArrayOutput `pulumi:"drainNatIps"`
	// -
	// (Optional)
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec pulumi.IntPtrOutput `pulumi:"icmpIdleTimeoutSec"`
	// -
	// (Optional)
	// Configuration for logging on NAT  Structure is documented below.
	LogConfig RouterNatLogConfigPtrOutput `pulumi:"logConfig"`
	// -
	// (Optional)
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm pulumi.IntPtrOutput `pulumi:"minPortsPerVm"`
	// -
	// (Required)
	// Self-link of subnetwork to NAT
	Name pulumi.StringOutput `pulumi:"name"`
	// -
	// (Required)
	// How external IPs should be allocated for this NAT. Valid values are
	// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	NatIpAllocateOption pulumi.StringOutput `pulumi:"natIpAllocateOption"`
	// -
	// (Optional)
	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NatIps pulumi.StringArrayOutput `pulumi:"natIps"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Optional)
	// Region where the router and NAT reside.
	Region pulumi.StringOutput `pulumi:"region"`
	// -
	// (Required)
	// The name of the Cloud Router in which this NAT will be configured.
	Router pulumi.StringOutput `pulumi:"router"`
	// -
	// (Required)
	// How NAT should be configured per Subnetwork.
	// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	SourceSubnetworkIpRangesToNat pulumi.StringOutput `pulumi:"sourceSubnetworkIpRangesToNat"`
	// -
	// (Optional)
	// One or more subnetwork NAT configurations. Only used if
	// `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`  Structure is documented below.
	Subnetworks RouterNatSubnetworkArrayOutput `pulumi:"subnetworks"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TcpEstablishedIdleTimeoutSec pulumi.IntPtrOutput `pulumi:"tcpEstablishedIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TcpTransitoryIdleTimeoutSec pulumi.IntPtrOutput `pulumi:"tcpTransitoryIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec pulumi.IntPtrOutput `pulumi:"udpIdleTimeoutSec"`
}

// NewRouterNat registers a new resource with the given unique name, arguments, and options.
func NewRouterNat(ctx *pulumi.Context,
	name string, args *RouterNatArgs, opts ...pulumi.ResourceOption) (*RouterNat, error) {
	if args == nil || args.NatIpAllocateOption == nil {
		return nil, errors.New("missing required argument 'NatIpAllocateOption'")
	}
	if args == nil || args.Router == nil {
		return nil, errors.New("missing required argument 'Router'")
	}
	if args == nil || args.SourceSubnetworkIpRangesToNat == nil {
		return nil, errors.New("missing required argument 'SourceSubnetworkIpRangesToNat'")
	}
	if args == nil {
		args = &RouterNatArgs{}
	}
	var resource RouterNat
	err := ctx.RegisterResource("gcp:compute/routerNat:RouterNat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterNat gets an existing RouterNat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterNat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterNatState, opts ...pulumi.ResourceOption) (*RouterNat, error) {
	var resource RouterNat
	err := ctx.ReadResource("gcp:compute/routerNat:RouterNat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterNat resources.
type routerNatState struct {
	// -
	// (Optional)
	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNatIps []string `pulumi:"drainNatIps"`
	// -
	// (Optional)
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec *int `pulumi:"icmpIdleTimeoutSec"`
	// -
	// (Optional)
	// Configuration for logging on NAT  Structure is documented below.
	LogConfig *RouterNatLogConfig `pulumi:"logConfig"`
	// -
	// (Optional)
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm *int `pulumi:"minPortsPerVm"`
	// -
	// (Required)
	// Self-link of subnetwork to NAT
	Name *string `pulumi:"name"`
	// -
	// (Required)
	// How external IPs should be allocated for this NAT. Valid values are
	// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	NatIpAllocateOption *string `pulumi:"natIpAllocateOption"`
	// -
	// (Optional)
	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NatIps []string `pulumi:"natIps"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// Region where the router and NAT reside.
	Region *string `pulumi:"region"`
	// -
	// (Required)
	// The name of the Cloud Router in which this NAT will be configured.
	Router *string `pulumi:"router"`
	// -
	// (Required)
	// How NAT should be configured per Subnetwork.
	// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	SourceSubnetworkIpRangesToNat *string `pulumi:"sourceSubnetworkIpRangesToNat"`
	// -
	// (Optional)
	// One or more subnetwork NAT configurations. Only used if
	// `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`  Structure is documented below.
	Subnetworks []RouterNatSubnetwork `pulumi:"subnetworks"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TcpEstablishedIdleTimeoutSec *int `pulumi:"tcpEstablishedIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TcpTransitoryIdleTimeoutSec *int `pulumi:"tcpTransitoryIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec *int `pulumi:"udpIdleTimeoutSec"`
}

type RouterNatState struct {
	// -
	// (Optional)
	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNatIps pulumi.StringArrayInput
	// -
	// (Optional)
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Configuration for logging on NAT  Structure is documented below.
	LogConfig RouterNatLogConfigPtrInput
	// -
	// (Optional)
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm pulumi.IntPtrInput
	// -
	// (Required)
	// Self-link of subnetwork to NAT
	Name pulumi.StringPtrInput
	// -
	// (Required)
	// How external IPs should be allocated for this NAT. Valid values are
	// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	NatIpAllocateOption pulumi.StringPtrInput
	// -
	// (Optional)
	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NatIps pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// Region where the router and NAT reside.
	Region pulumi.StringPtrInput
	// -
	// (Required)
	// The name of the Cloud Router in which this NAT will be configured.
	Router pulumi.StringPtrInput
	// -
	// (Required)
	// How NAT should be configured per Subnetwork.
	// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	SourceSubnetworkIpRangesToNat pulumi.StringPtrInput
	// -
	// (Optional)
	// One or more subnetwork NAT configurations. Only used if
	// `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`  Structure is documented below.
	Subnetworks RouterNatSubnetworkArrayInput
	// -
	// (Optional)
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TcpEstablishedIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TcpTransitoryIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec pulumi.IntPtrInput
}

func (RouterNatState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerNatState)(nil)).Elem()
}

type routerNatArgs struct {
	// -
	// (Optional)
	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNatIps []string `pulumi:"drainNatIps"`
	// -
	// (Optional)
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec *int `pulumi:"icmpIdleTimeoutSec"`
	// -
	// (Optional)
	// Configuration for logging on NAT  Structure is documented below.
	LogConfig *RouterNatLogConfig `pulumi:"logConfig"`
	// -
	// (Optional)
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm *int `pulumi:"minPortsPerVm"`
	// -
	// (Required)
	// Self-link of subnetwork to NAT
	Name *string `pulumi:"name"`
	// -
	// (Required)
	// How external IPs should be allocated for this NAT. Valid values are
	// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	NatIpAllocateOption string `pulumi:"natIpAllocateOption"`
	// -
	// (Optional)
	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NatIps []string `pulumi:"natIps"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// Region where the router and NAT reside.
	Region *string `pulumi:"region"`
	// -
	// (Required)
	// The name of the Cloud Router in which this NAT will be configured.
	Router string `pulumi:"router"`
	// -
	// (Required)
	// How NAT should be configured per Subnetwork.
	// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	SourceSubnetworkIpRangesToNat string `pulumi:"sourceSubnetworkIpRangesToNat"`
	// -
	// (Optional)
	// One or more subnetwork NAT configurations. Only used if
	// `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`  Structure is documented below.
	Subnetworks []RouterNatSubnetwork `pulumi:"subnetworks"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TcpEstablishedIdleTimeoutSec *int `pulumi:"tcpEstablishedIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TcpTransitoryIdleTimeoutSec *int `pulumi:"tcpTransitoryIdleTimeoutSec"`
	// -
	// (Optional)
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec *int `pulumi:"udpIdleTimeoutSec"`
}

// The set of arguments for constructing a RouterNat resource.
type RouterNatArgs struct {
	// -
	// (Optional)
	// A list of URLs of the IP resources to be drained. These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	DrainNatIps pulumi.StringArrayInput
	// -
	// (Optional)
	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Configuration for logging on NAT  Structure is documented below.
	LogConfig RouterNatLogConfigPtrInput
	// -
	// (Optional)
	// Minimum number of ports allocated to a VM from this NAT.
	MinPortsPerVm pulumi.IntPtrInput
	// -
	// (Required)
	// Self-link of subnetwork to NAT
	Name pulumi.StringPtrInput
	// -
	// (Required)
	// How external IPs should be allocated for this NAT. Valid values are
	// `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	// Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	NatIpAllocateOption pulumi.StringInput
	// -
	// (Optional)
	// Self-links of NAT IPs. Only valid if natIpAllocateOption
	// is set to MANUAL_ONLY.
	NatIps pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// Region where the router and NAT reside.
	Region pulumi.StringPtrInput
	// -
	// (Required)
	// The name of the Cloud Router in which this NAT will be configured.
	Router pulumi.StringInput
	// -
	// (Required)
	// How NAT should be configured per Subnetwork.
	// If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	// IP ranges in every Subnetwork are allowed to Nat.
	// If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	// ranges in every Subnetwork are allowed to Nat.
	// `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	// (specified in the field subnetwork below). Note that if this field
	// contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	// ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	// other RouterNat section in any Router for this network in this region.
	SourceSubnetworkIpRangesToNat pulumi.StringInput
	// -
	// (Optional)
	// One or more subnetwork NAT configurations. Only used if
	// `sourceSubnetworkIpRangesToNat` is set to `LIST_OF_SUBNETWORKS`  Structure is documented below.
	Subnetworks RouterNatSubnetworkArrayInput
	// -
	// (Optional)
	// Timeout (in seconds) for TCP established connections.
	// Defaults to 1200s if not set.
	TcpEstablishedIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Timeout (in seconds) for TCP transitory connections.
	// Defaults to 30s if not set.
	TcpTransitoryIdleTimeoutSec pulumi.IntPtrInput
	// -
	// (Optional)
	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec pulumi.IntPtrInput
}

func (RouterNatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerNatArgs)(nil)).Elem()
}
