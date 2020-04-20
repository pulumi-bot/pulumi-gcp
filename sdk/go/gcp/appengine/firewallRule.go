// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A single firewall rule that is evaluated against incoming traffic
// and provides an action to take on matched requests.
//
//
// To get more information about FirewallRule, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)
type FirewallRule struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// The action to take if this rule matches.
	Action pulumi.StringOutput `pulumi:"action"`
	// -
	// (Optional)
	// An optional string description of this rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Optional)
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Required)
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringOutput `pulumi:"sourceRange"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil || args.Action == nil {
		return nil, errors.New("missing required argument 'Action'")
	}
	if args == nil || args.SourceRange == nil {
		return nil, errors.New("missing required argument 'SourceRange'")
	}
	if args == nil {
		args = &FirewallRuleArgs{}
	}
	var resource FirewallRule
	err := ctx.RegisterResource("gcp:appengine/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("gcp:appengine/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// -
	// (Required)
	// The action to take if this rule matches.
	Action *string `pulumi:"action"`
	// -
	// (Optional)
	// An optional string description of this rule.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Required)
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange *string `pulumi:"sourceRange"`
}

type FirewallRuleState struct {
	// -
	// (Required)
	// The action to take if this rule matches.
	Action pulumi.StringPtrInput
	// -
	// (Optional)
	// An optional string description of this rule.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Required)
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// -
	// (Required)
	// The action to take if this rule matches.
	Action string `pulumi:"action"`
	// -
	// (Optional)
	// An optional string description of this rule.
	Description *string `pulumi:"description"`
	// -
	// (Optional)
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Required)
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange string `pulumi:"sourceRange"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// -
	// (Required)
	// The action to take if this rule matches.
	Action pulumi.StringInput
	// -
	// (Optional)
	// An optional string description of this rule.
	Description pulumi.StringPtrInput
	// -
	// (Optional)
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Required)
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}
