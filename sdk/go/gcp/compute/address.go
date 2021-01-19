// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
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
// To get more information about Address, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
// * How-to Guides
//     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
//     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
//
// ## Example Usage
// ### Address Basic
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
// 		_, err := compute.NewAddress(ctx, "ipAddress", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Address With Subnetwork
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
// ### Address With Gce Endpoint
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
// 		_, err := compute.NewAddress(ctx, "internalWithGceEndpoint", &compute.AddressArgs{
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
// ### Instance With Ip
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
// 		static, err := compute.NewAddress(ctx, "static", nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := "debian-9"
// 		opt1 := "debian-cloud"
// 		debianImage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
// 			Family:  &opt0,
// 			Project: &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "instanceWithIp", &compute.InstanceArgs{
// 			MachineType: pulumi.String("f1-micro"),
// 			Zone:        pulumi.String("us-central1-a"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String(debianImage.SelfLink),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						&compute.InstanceNetworkInterfaceAccessConfigArgs{
// 							NatIp: static.Address,
// 						},
// 					},
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
//
// ## Import
//
// Address can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/address:Address default projects/{{project}}/regions/{{region}}/addresses/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/address:Address default {{project}}/{{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/address:Address default {{region}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/address:Address default {{name}}
// ```
type Address struct {
	pulumi.CustomResourceState

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve.
	// Default value is `EXTERNAL`.
	// Possible values are `INTERNAL` and `EXTERNAL`.
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
	// Possible values are `PREMIUM` and `STANDARD`.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//   This should only be set when using an Internal address.
	//   Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
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
	// Default value is `EXTERNAL`.
	// Possible values are `INTERNAL` and `EXTERNAL`.
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
	// Possible values are `PREMIUM` and `STANDARD`.
	NetworkTier *string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//   This should only be set when using an Internal address.
	//   Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
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
	// Default value is `EXTERNAL`.
	// Possible values are `INTERNAL` and `EXTERNAL`.
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
	// Possible values are `PREMIUM` and `STANDARD`.
	NetworkTier pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values:
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//   This should only be set when using an Internal address.
	//   Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
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
	// Default value is `EXTERNAL`.
	// Possible values are `INTERNAL` and `EXTERNAL`.
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
	// Possible values are `PREMIUM` and `STANDARD`.
	NetworkTier *string `pulumi:"networkTier"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values:
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//   This should only be set when using an Internal address.
	//   Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
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
	// Default value is `EXTERNAL`.
	// Possible values are `INTERNAL` and `EXTERNAL`.
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
	// Possible values are `PREMIUM` and `STANDARD`.
	NetworkTier pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values:
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//   This should only be set when using an Internal address.
	//   Possible values are `GCE_ENDPOINT`, `VPC_PEERING`, and `SHARED_LOADBALANCER_VIP`.
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

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil))
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

func (i *Address) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *Address) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
}

type AddressPtrInput interface {
	pulumi.Input

	ToAddressPtrOutput() AddressPtrOutput
	ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput
}

type addressPtrType AddressArgs

func (*addressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil))
}

func (i *addressPtrType) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *addressPtrType) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
}

// AddressArrayInput is an input type that accepts AddressArray and AddressArrayOutput values.
// You can construct a concrete instance of `AddressArrayInput` via:
//
//          AddressArray{ AddressArgs{...} }
type AddressArrayInput interface {
	pulumi.Input

	ToAddressArrayOutput() AddressArrayOutput
	ToAddressArrayOutputWithContext(context.Context) AddressArrayOutput
}

type AddressArray []AddressInput

func (AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Address)(nil))
}

func (i AddressArray) ToAddressArrayOutput() AddressArrayOutput {
	return i.ToAddressArrayOutputWithContext(context.Background())
}

func (i AddressArray) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressArrayOutput)
}

// AddressMapInput is an input type that accepts AddressMap and AddressMapOutput values.
// You can construct a concrete instance of `AddressMapInput` via:
//
//          AddressMap{ "key": AddressArgs{...} }
type AddressMapInput interface {
	pulumi.Input

	ToAddressMapOutput() AddressMapOutput
	ToAddressMapOutputWithContext(context.Context) AddressMapOutput
}

type AddressMap map[string]AddressInput

func (AddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Address)(nil))
}

func (i AddressMap) ToAddressMapOutput() AddressMapOutput {
	return i.ToAddressMapOutputWithContext(context.Background())
}

func (i AddressMap) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressMapOutput)
}

type AddressOutput struct {
	*pulumi.OutputState
}

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil))
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

func (o AddressOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o.ToAddressPtrOutputWithContext(context.Background())
}

func (o AddressOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o.ApplyT(func(v Address) *Address {
		return &v
	}).(AddressPtrOutput)
}

type AddressPtrOutput struct {
	*pulumi.OutputState
}

func (AddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil))
}

func (o AddressPtrOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o
}

type AddressArrayOutput struct{ *pulumi.OutputState }

func (AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Address)(nil))
}

func (o AddressArrayOutput) ToAddressArrayOutput() AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) Index(i pulumi.IntInput) AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Address {
		return vs[0].([]Address)[vs[1].(int)]
	}).(AddressOutput)
}

type AddressMapOutput struct{ *pulumi.OutputState }

func (AddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Address)(nil))
}

func (o AddressMapOutput) ToAddressMapOutput() AddressMapOutput {
	return o
}

func (o AddressMapOutput) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return o
}

func (o AddressMapOutput) MapIndex(k pulumi.StringInput) AddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Address {
		return vs[0].(map[string]Address)[vs[1].(string)]
	}).(AddressOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressPtrOutput{})
	pulumi.RegisterOutputType(AddressArrayOutput{})
	pulumi.RegisterOutputType(AddressMapOutput{})
}
