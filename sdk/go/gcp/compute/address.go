// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents an Address resource.
//
// Each virtual machine instance has an ephemeral internal IP address and,
// optionally, an external IP address. To communicate between instances on
// the same network, you can use an instance's internal IP address. To
// communicate with the Internet and instances outside of the same network,
// you must specify the instance's external IP address.
//
// Internal IP addresses are ephemeral and only belong to an instance for
// the lifetime of the instance; if the instance is deleted and recreated,
// the instance is assigned a new internal IP address, either by Compute
// Engine or by you. External IP addresses can be either ephemeral or
// static.
//
//
// To get more information about Address, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
// * How-to Guides
//     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
//     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
//
// ## Example Usage
//
// ### Address Basic
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
// 		_, err = compute.NewAddress(ctx, "ipAddress", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Address With Subnetwork
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
// 		defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", nil)
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
// 		_, err = compute.NewAddress(ctx, "internalWithSubnetAndAddress", &compute.AddressArgs{
// 			Subnetwork:  defaultSubnetwork.ID(),
// 			AddressType: pulumi.String("INTERNAL"),
// 			Address:     pulumi.String("10.0.42.42"),
// 			Region:      pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Address With Gce Endpoint
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
// 		_, err = compute.NewAddress(ctx, "internalWithGceEndpoint", &compute.AddressArgs{
// 			AddressType: pulumi.String("INTERNAL"),
// 			Purpose:     pulumi.String("GCE_ENDPOINT"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Address struct {
	pulumi.CustomResourceState

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve.
	AddressType pulumi.StringPtrOutput `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewAddress registers a new resource with the given unique name, arguments, and options.
func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
	if args == nil {
		args = &AddressArgs{}
	}
	var resource Address
	err := ctx.RegisterResource("gcp:compute/address:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress gets an existing Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("gcp:compute/address:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address resources.
type addressState struct {
	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address *string `pulumi:"address"`
	// The type of address to reserve.
	AddressType *string `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	NetworkTier *string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork *string `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users []string `pulumi:"users"`
}

type AddressState struct {
	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address pulumi.StringPtrInput
	// The type of address to reserve.
	AddressType pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values:
	// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringPtrInput
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayInput
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address *string `pulumi:"address"`
	// The type of address to reserve.
	AddressType *string `pulumi:"addressType"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	NetworkTier *string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork *string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a Address resource.
type AddressArgs struct {
	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address pulumi.StringPtrInput
	// The type of address to reserve.
	AddressType pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values:
	// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	Subnetwork pulumi.StringPtrInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}
