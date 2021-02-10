// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A VPC network is a virtual version of the traditional physical networks
// that exist within and between physical data centers. A VPC network
// provides connectivity for your Compute Engine virtual machine (VM)
// instances, Container Engine containers, App Engine Flex services, and
// other network-related resources.
//
// Each GCP project contains one or more VPC networks. Each VPC network is a
// global entity spanning all GCP regions. This global VPC network allows VM
// instances and other resources to communicate with each other via internal,
// private IP addresses.
//
// Each VPC network is subdivided into subnets, and each subnet is contained
// within a single region. You can have more than one subnet in a region for
// a given VPC network. Each subnet has a contiguous private RFC1918 IP
// space. You create instances, containers, and the like in these subnets.
// When you create an instance, you must create it in a subnet, and the
// instance draws its internal IP address from that subnet.
//
// Virtual machine (VM) instances in a VPC network can communicate with
// instances in all other subnets of the same VPC network, regardless of
// region, using their RFC1918 private IP addresses. You can isolate portions
// of the network, even entire subnets, using firewall rules.
//
// To get more information about Subnetwork, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/subnetworks)
// * How-to Guides
//     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
//     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
//
// ## Example Usage
// ### Subnetwork Basic
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
// 		_, err := compute.NewNetwork(ctx, "custom_test", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network_with_private_secondary_ip_ranges", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.2.0.0/16"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     custom_test.ID(),
// 			SecondaryIpRanges: compute.SubnetworkSecondaryIpRangeArray{
// 				&compute.SubnetworkSecondaryIpRangeArgs{
// 					RangeName:   pulumi.String("tf-test-secondary-range-update1"),
// 					IpCidrRange: pulumi.String("192.168.10.0/24"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Subnetwork Logging Config
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
// 		_, err := compute.NewNetwork(ctx, "custom_test", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "subnet_with_logging", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.2.0.0/16"),
// 			Region:      pulumi.String("us-central1"),
// 			Network:     custom_test.ID(),
// 			LogConfig: &compute.SubnetworkLogConfigArgs{
// 				AggregationInterval: pulumi.String("INTERVAL_10_MIN"),
// 				FlowSampling:        pulumi.Float64(0.5),
// 				Metadata:            pulumi.String("INCLUDE_ALL_METADATA"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Subnetwork Internal L7lb
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
// 		_, err := compute.NewNetwork(ctx, "custom_test", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewSubnetwork(ctx, "network_for_l7lb", &compute.SubnetworkArgs{
// 			IpCidrRange: pulumi.String("10.0.0.0/22"),
// 			Region:      pulumi.String("us-central1"),
// 			Purpose:     pulumi.String("INTERNAL_HTTPS_LOAD_BALANCER"),
// 			Role:        pulumi.String("ACTIVE"),
// 			Network:     custom_test.ID(),
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
// Subnetwork can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/subnetwork:Subnetwork default projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/subnetwork:Subnetwork default {{name}}
// ```
type Subnetwork struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringOutput `pulumi:"gatewayAddress"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
	// Structure is documented below.
	LogConfig SubnetworkLogConfigPtrOutput `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringOutput `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrOutput `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringOutput `pulumi:"privateIpv6GoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are `ACTIVE` and `BACKUP`.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayOutput `pulumi:"secondaryIpRanges"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSubnetwork registers a new resource with the given unique name, arguments, and options.
func NewSubnetwork(ctx *pulumi.Context,
	name string, args *SubnetworkArgs, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpCidrRange == nil {
		return nil, errors.New("invalid value for required argument 'IpCidrRange'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	var resource Subnetwork
	err := ctx.RegisterResource("gcp:compute/subnetwork:Subnetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetwork gets an existing Subnetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkState, opts ...pulumi.ResourceOption) (*Subnetwork, error) {
	var resource Subnetwork
	err := ctx.ReadResource("gcp:compute/subnetwork:Subnetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnetwork resources.
type subnetworkState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint *string `pulumi:"fingerprint"`
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress *string `pulumi:"gatewayAddress"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
	// Structure is documented below.
	LogConfig *SubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network *string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
	Purpose *string `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are `ACTIVE` and `BACKUP`.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges []SubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SubnetworkState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	//
	// Deprecated: This field is not useful for users, and has been removed as an output.
	Fingerprint pulumi.StringPtrInput
	// The gateway address for default routes to reach destination addresses outside this subnetwork.
	GatewayAddress pulumi.StringPtrInput
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringPtrInput
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
	// Structure is documented below.
	LogConfig SubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringPtrInput
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
	Purpose pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are `ACTIVE` and `BACKUP`.
	Role pulumi.StringPtrInput
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SubnetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkState)(nil)).Elem()
}

type subnetworkArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description *string `pulumi:"description"`
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange string `pulumi:"ipCidrRange"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
	// Structure is documented below.
	LogConfig *SubnetworkLogConfig `pulumi:"logConfig"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network string `pulumi:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess *bool `pulumi:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess *string `pulumi:"privateIpv6GoogleAccess"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
	Purpose *string `pulumi:"purpose"`
	// The GCP region for this subnetwork.
	Region *string `pulumi:"region"`
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are `ACTIVE` and `BACKUP`.
	Role *string `pulumi:"role"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges []SubnetworkSecondaryIpRange `pulumi:"secondaryIpRanges"`
}

// The set of arguments for constructing a Subnetwork resource.
type SubnetworkArgs struct {
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	Description pulumi.StringPtrInput
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCidrRange pulumi.StringInput
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
	// subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`
	// Structure is documented below.
	LogConfig SubnetworkLogConfigPtrInput
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?` which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network pulumi.StringInput
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	PrivateIpGoogleAccess pulumi.BoolPtrInput
	// The private IPv6 google access type for the VMs in this subnet.
	PrivateIpv6GoogleAccess pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. This field can be either PRIVATE
	// or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
	// INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
	// reserved for Internal HTTP(S) Load Balancing. If unspecified, the
	// purpose defaults to PRIVATE.
	// If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
	// Possible values are `INTERNAL_HTTPS_LOAD_BALANCER` and `PRIVATE`.
	Purpose pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	Region pulumi.StringPtrInput
	// The role of subnetwork. Currently, this field is only used when
	// purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
	// or BACKUP. An ACTIVE subnetwork is one that is currently being used
	// for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
	// is ready to be promoted to ACTIVE or is currently draining.
	// Possible values are `ACTIVE` and `BACKUP`.
	Role pulumi.StringPtrInput
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges. Structure is documented below.
	SecondaryIpRanges SubnetworkSecondaryIpRangeArrayInput
}

func (SubnetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkArgs)(nil)).Elem()
}

type SubnetworkInput interface {
	pulumi.Input

	ToSubnetworkOutput() SubnetworkOutput
	ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput
}

func (*Subnetwork) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnetwork)(nil))
}

func (i *Subnetwork) ToSubnetworkOutput() SubnetworkOutput {
	return i.ToSubnetworkOutputWithContext(context.Background())
}

func (i *Subnetwork) ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkOutput)
}

func (i *Subnetwork) ToSubnetworkPtrOutput() SubnetworkPtrOutput {
	return i.ToSubnetworkPtrOutputWithContext(context.Background())
}

func (i *Subnetwork) ToSubnetworkPtrOutputWithContext(ctx context.Context) SubnetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkPtrOutput)
}

type SubnetworkPtrInput interface {
	pulumi.Input

	ToSubnetworkPtrOutput() SubnetworkPtrOutput
	ToSubnetworkPtrOutputWithContext(ctx context.Context) SubnetworkPtrOutput
}

type subnetworkPtrType SubnetworkArgs

func (*subnetworkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnetwork)(nil))
}

func (i *subnetworkPtrType) ToSubnetworkPtrOutput() SubnetworkPtrOutput {
	return i.ToSubnetworkPtrOutputWithContext(context.Background())
}

func (i *subnetworkPtrType) ToSubnetworkPtrOutputWithContext(ctx context.Context) SubnetworkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkPtrOutput)
}

// SubnetworkArrayInput is an input type that accepts SubnetworkArray and SubnetworkArrayOutput values.
// You can construct a concrete instance of `SubnetworkArrayInput` via:
//
//          SubnetworkArray{ SubnetworkArgs{...} }
type SubnetworkArrayInput interface {
	pulumi.Input

	ToSubnetworkArrayOutput() SubnetworkArrayOutput
	ToSubnetworkArrayOutputWithContext(context.Context) SubnetworkArrayOutput
}

type SubnetworkArray []SubnetworkInput

func (SubnetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Subnetwork)(nil))
}

func (i SubnetworkArray) ToSubnetworkArrayOutput() SubnetworkArrayOutput {
	return i.ToSubnetworkArrayOutputWithContext(context.Background())
}

func (i SubnetworkArray) ToSubnetworkArrayOutputWithContext(ctx context.Context) SubnetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkArrayOutput)
}

// SubnetworkMapInput is an input type that accepts SubnetworkMap and SubnetworkMapOutput values.
// You can construct a concrete instance of `SubnetworkMapInput` via:
//
//          SubnetworkMap{ "key": SubnetworkArgs{...} }
type SubnetworkMapInput interface {
	pulumi.Input

	ToSubnetworkMapOutput() SubnetworkMapOutput
	ToSubnetworkMapOutputWithContext(context.Context) SubnetworkMapOutput
}

type SubnetworkMap map[string]SubnetworkInput

func (SubnetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Subnetwork)(nil))
}

func (i SubnetworkMap) ToSubnetworkMapOutput() SubnetworkMapOutput {
	return i.ToSubnetworkMapOutputWithContext(context.Background())
}

func (i SubnetworkMap) ToSubnetworkMapOutputWithContext(ctx context.Context) SubnetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkMapOutput)
}

type SubnetworkOutput struct {
	*pulumi.OutputState
}

func (SubnetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnetwork)(nil))
}

func (o SubnetworkOutput) ToSubnetworkOutput() SubnetworkOutput {
	return o
}

func (o SubnetworkOutput) ToSubnetworkOutputWithContext(ctx context.Context) SubnetworkOutput {
	return o
}

func (o SubnetworkOutput) ToSubnetworkPtrOutput() SubnetworkPtrOutput {
	return o.ToSubnetworkPtrOutputWithContext(context.Background())
}

func (o SubnetworkOutput) ToSubnetworkPtrOutputWithContext(ctx context.Context) SubnetworkPtrOutput {
	return o.ApplyT(func(v Subnetwork) *Subnetwork {
		return &v
	}).(SubnetworkPtrOutput)
}

type SubnetworkPtrOutput struct {
	*pulumi.OutputState
}

func (SubnetworkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnetwork)(nil))
}

func (o SubnetworkPtrOutput) ToSubnetworkPtrOutput() SubnetworkPtrOutput {
	return o
}

func (o SubnetworkPtrOutput) ToSubnetworkPtrOutputWithContext(ctx context.Context) SubnetworkPtrOutput {
	return o
}

type SubnetworkArrayOutput struct{ *pulumi.OutputState }

func (SubnetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnetwork)(nil))
}

func (o SubnetworkArrayOutput) ToSubnetworkArrayOutput() SubnetworkArrayOutput {
	return o
}

func (o SubnetworkArrayOutput) ToSubnetworkArrayOutputWithContext(ctx context.Context) SubnetworkArrayOutput {
	return o
}

func (o SubnetworkArrayOutput) Index(i pulumi.IntInput) SubnetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subnetwork {
		return vs[0].([]Subnetwork)[vs[1].(int)]
	}).(SubnetworkOutput)
}

type SubnetworkMapOutput struct{ *pulumi.OutputState }

func (SubnetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Subnetwork)(nil))
}

func (o SubnetworkMapOutput) ToSubnetworkMapOutput() SubnetworkMapOutput {
	return o
}

func (o SubnetworkMapOutput) ToSubnetworkMapOutputWithContext(ctx context.Context) SubnetworkMapOutput {
	return o
}

func (o SubnetworkMapOutput) MapIndex(k pulumi.StringInput) SubnetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Subnetwork {
		return vs[0].(map[string]Subnetwork)[vs[1].(string)]
	}).(SubnetworkOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetworkOutput{})
	pulumi.RegisterOutputType(SubnetworkPtrOutput{})
	pulumi.RegisterOutputType(SubnetworkArrayOutput{})
	pulumi.RegisterOutputType(SubnetworkMapOutput{})
}
