// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Packet Mirroring mirrors traffic to and from particular VM instances.
// You can use the collected traffic to help you detect security threats
// and monitor application performance.
//
// To get more information about PacketMirroring, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/packetMirroring)
// * How-to Guides
//     * [Using Packet Mirroring](https://cloud.google.com/vpc/docs/using-packet-mirroring#creating)
//
// ## Example Usage
//
// ## Import
//
// PacketMirroring can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/packetMirroring:PacketMirroring default projects/{{project}}/regions/{{region}}/packetMirrorings/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/packetMirroring:PacketMirroring default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/packetMirroring:PacketMirroring default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/packetMirroring:PacketMirroring default {{name}}
// ```
type PacketMirroring struct {
	pulumi.CustomResourceState

	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have isMirroringCollector
	// set to true.
	// Structure is documented below.
	CollectorIlb PacketMirroringCollectorIlbOutput `pulumi:"collectorIlb"`
	// A human-readable description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter PacketMirroringFilterPtrOutput `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources PacketMirroringMirroredResourcesOutput `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network PacketMirroringNetworkOutput `pulumi:"network"`
	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPacketMirroring registers a new resource with the given unique name, arguments, and options.
func NewPacketMirroring(ctx *pulumi.Context,
	name string, args *PacketMirroringArgs, opts ...pulumi.ResourceOption) (*PacketMirroring, error) {
	if args == nil || args.CollectorIlb == nil {
		return nil, errors.New("missing required argument 'CollectorIlb'")
	}
	if args == nil || args.MirroredResources == nil {
		return nil, errors.New("missing required argument 'MirroredResources'")
	}
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	if args == nil {
		args = &PacketMirroringArgs{}
	}
	var resource PacketMirroring
	err := ctx.RegisterResource("gcp:compute/packetMirroring:PacketMirroring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPacketMirroring gets an existing PacketMirroring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPacketMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketMirroringState, opts ...pulumi.ResourceOption) (*PacketMirroring, error) {
	var resource PacketMirroring
	err := ctx.ReadResource("gcp:compute/packetMirroring:PacketMirroring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketMirroring resources.
type packetMirroringState struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have isMirroringCollector
	// set to true.
	// Structure is documented below.
	CollectorIlb *PacketMirroringCollectorIlb `pulumi:"collectorIlb"`
	// A human-readable description of the rule.
	Description *string `pulumi:"description"`
	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter *PacketMirroringFilter `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources *PacketMirroringMirroredResources `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name *string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network *PacketMirroringNetwork `pulumi:"network"`
	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

type PacketMirroringState struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have isMirroringCollector
	// set to true.
	// Structure is documented below.
	CollectorIlb PacketMirroringCollectorIlbPtrInput
	// A human-readable description of the rule.
	Description pulumi.StringPtrInput
	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter PacketMirroringFilterPtrInput
	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources PacketMirroringMirroredResourcesPtrInput
	// The name of the packet mirroring rule
	Name pulumi.StringPtrInput
	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network PacketMirroringNetworkPtrInput
	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (PacketMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetMirroringState)(nil)).Elem()
}

type packetMirroringArgs struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have isMirroringCollector
	// set to true.
	// Structure is documented below.
	CollectorIlb PacketMirroringCollectorIlb `pulumi:"collectorIlb"`
	// A human-readable description of the rule.
	Description *string `pulumi:"description"`
	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter *PacketMirroringFilter `pulumi:"filter"`
	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources PacketMirroringMirroredResources `pulumi:"mirroredResources"`
	// The name of the packet mirroring rule
	Name *string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network PacketMirroringNetwork `pulumi:"network"`
	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a PacketMirroring resource.
type PacketMirroringArgs struct {
	// The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
	// that will be used as collector for mirrored traffic. The
	// specified forwarding rule must have isMirroringCollector
	// set to true.
	// Structure is documented below.
	CollectorIlb PacketMirroringCollectorIlbInput
	// A human-readable description of the rule.
	Description pulumi.StringPtrInput
	// A filter for mirrored traffic.  If unset, all traffic is mirrored.
	// Structure is documented below.
	Filter PacketMirroringFilterPtrInput
	// A means of specifying which resources to mirror.
	// Structure is documented below.
	MirroredResources PacketMirroringMirroredResourcesInput
	// The name of the packet mirroring rule
	Name pulumi.StringPtrInput
	// Specifies the mirrored VPC network. Only packets in this network
	// will be mirrored. All mirrored VMs should have a NIC in the given
	// network. All mirrored subnetworks should belong to the given network.
	// Structure is documented below.
	Network PacketMirroringNetworkInput
	// Since only one rule can be active at a time, priority is
	// used to break ties in the case of two rules that apply to
	// the same instances.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
}

func (PacketMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetMirroringArgs)(nil)).Elem()
}

type PacketMirroringInput interface {
	pulumi.Input

	ToPacketMirroringOutput() PacketMirroringOutput
	ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput
}

func (PacketMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketMirroring)(nil)).Elem()
}

func (i PacketMirroring) ToPacketMirroringOutput() PacketMirroringOutput {
	return i.ToPacketMirroringOutputWithContext(context.Background())
}

func (i PacketMirroring) ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketMirroringOutput)
}

type PacketMirroringOutput struct {
	*pulumi.OutputState
}

func (PacketMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketMirroringOutput)(nil)).Elem()
}

func (o PacketMirroringOutput) ToPacketMirroringOutput() PacketMirroringOutput {
	return o
}

func (o PacketMirroringOutput) ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PacketMirroringOutput{})
}
