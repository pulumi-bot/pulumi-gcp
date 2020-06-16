// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Global Address resource. Global addresses are used for
// HTTP(S) load balancing.
//
// To get more information about GlobalAddress, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
// * How-to Guides
//     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
//
// {{% examples %}}
// ## Example Usage
// {{% example %}}
// ### Global Address Basic
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
// 		_, err = compute.NewGlobalAddress(ctx, "default", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {{% /example %}}
// {{% /examples %}}
type GlobalAddress struct {
	pulumi.CustomResourceState

	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of the address to reserve.
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network.
	AddressType pulumi.StringPtrOutput `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IP Version that will be used by this address. The default value is `IPV4`.
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The purpose of the resource. For global internal addresses it can be
	// * VPC_PEERING - for peer networks
	//   This should only be set when using an Internal address.
	Purpose pulumi.StringPtrOutput `pulumi:"purpose"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewGlobalAddress registers a new resource with the given unique name, arguments, and options.
func NewGlobalAddress(ctx *pulumi.Context,
	name string, args *GlobalAddressArgs, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	if args == nil {
		args = &GlobalAddressArgs{}
	}
	var resource GlobalAddress
	err := ctx.RegisterResource("gcp:compute/globalAddress:GlobalAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalAddress gets an existing GlobalAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalAddressState, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	var resource GlobalAddress
	err := ctx.ReadResource("gcp:compute/globalAddress:GlobalAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalAddress resources.
type globalAddressState struct {
	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address *string `pulumi:"address"`
	// The type of the address to reserve.
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network.
	AddressType *string `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The IP Version that will be used by this address. The default value is `IPV4`.
	IpVersion *string `pulumi:"ipVersion"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network *string `pulumi:"network"`
	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL.
	PrefixLength *int `pulumi:"prefixLength"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. For global internal addresses it can be
	// * VPC_PEERING - for peer networks
	//   This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type GlobalAddressState struct {
	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address pulumi.StringPtrInput
	// The type of the address to reserve.
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network.
	AddressType pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The IP Version that will be used by this address. The default value is `IPV4`.
	IpVersion pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network pulumi.StringPtrInput
	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL.
	PrefixLength pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. For global internal addresses it can be
	// * VPC_PEERING - for peer networks
	//   This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (GlobalAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressState)(nil)).Elem()
}

type globalAddressArgs struct {
	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address *string `pulumi:"address"`
	// The type of the address to reserve.
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network.
	AddressType *string `pulumi:"addressType"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The IP Version that will be used by this address. The default value is `IPV4`.
	IpVersion *string `pulumi:"ipVersion"`
	// Labels to apply to this address.  A list of key->value pairs.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network *string `pulumi:"network"`
	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL.
	PrefixLength *int `pulumi:"prefixLength"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of the resource. For global internal addresses it can be
	// * VPC_PEERING - for peer networks
	//   This should only be set when using an Internal address.
	Purpose *string `pulumi:"purpose"`
}

// The set of arguments for constructing a GlobalAddress resource.
type GlobalAddressArgs struct {
	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address pulumi.StringPtrInput
	// The type of the address to reserve.
	// * EXTERNAL indicates public/external single IP address.
	// * INTERNAL indicates internal IP ranges belonging to some network.
	AddressType pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The IP Version that will be used by this address. The default value is `IPV4`.
	IpVersion pulumi.StringPtrInput
	// Labels to apply to this address.  A list of key->value pairs.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network pulumi.StringPtrInput
	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL.
	PrefixLength pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of the resource. For global internal addresses it can be
	// * VPC_PEERING - for peer networks
	//   This should only be set when using an Internal address.
	Purpose pulumi.StringPtrInput
}

func (GlobalAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressArgs)(nil)).Elem()
}
