// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A policy is a collection of DNS rules applied to one or more Virtual
// Private Cloud resources.
//
// To get more information about Policy, see:
//
// * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
// * How-to Guides
//     * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)
//
// ## Example Usage
//
// ## Import
//
// Policy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default projects/{{project}}/policies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default {{name}}
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrOutput `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrOutput `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayOutput `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("gcp:dns/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("gcp:dns/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding *bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging *bool `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name *string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks []PolicyNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type PolicyState struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrInput
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrInput
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrInput
	// User assigned name for this policy.
	Name pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding *bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging *bool `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name *string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks []PolicyNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrInput
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrInput
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrInput
	// User assigned name for this policy.
	Name pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
