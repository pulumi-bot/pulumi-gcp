// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A rule for the OrganizationSecurityPolicy.
//
// To get more information about OrganizationSecurityPolicyRule, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addRule)
// * How-to Guides
//     * [Creating firewall rules](https://cloud.google.com/vpc/docs/using-firewall-policies#create-rules)
//
// ## Example Usage
type OrganizationSecurityPolicyRule struct {
	pulumi.CustomResourceState

	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "gotoNext".
	Action pulumi.StringOutput `pulumi:"action"`
	// A description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created.
	// Possible values are `INGRESS` and `EGRESS`.
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// Structure is documented below.
	Match OrganizationSecurityPolicyRuleMatchOutput `pulumi:"match"`
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// If set to true, the specified action is not enforced.
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// An integer indicating the priority of a rule in the list. The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources pulumi.StringArrayOutput `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayOutput `pulumi:"targetServiceAccounts"`
}

// NewOrganizationSecurityPolicyRule registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSecurityPolicyRule(ctx *pulumi.Context,
	name string, args *OrganizationSecurityPolicyRuleArgs, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	var resource OrganizationSecurityPolicyRule
	err := ctx.RegisterResource("gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSecurityPolicyRule gets an existing OrganizationSecurityPolicyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSecurityPolicyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSecurityPolicyRuleState, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicyRule, error) {
	var resource OrganizationSecurityPolicyRule
	err := ctx.ReadResource("gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSecurityPolicyRule resources.
type organizationSecurityPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "gotoNext".
	Action *string `pulumi:"action"`
	// A description of the rule.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created.
	// Possible values are `INGRESS` and `EGRESS`.
	Direction *string `pulumi:"direction"`
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging *bool `pulumi:"enableLogging"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// Structure is documented below.
	Match *OrganizationSecurityPolicyRuleMatch `pulumi:"match"`
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId *string `pulumi:"policyId"`
	// If set to true, the specified action is not enforced.
	Preview *bool `pulumi:"preview"`
	// An integer indicating the priority of a rule in the list. The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	Priority *int `pulumi:"priority"`
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources []string `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

type OrganizationSecurityPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "gotoNext".
	Action pulumi.StringPtrInput
	// A description of the rule.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. If unspecified an INGRESS rule is created.
	// Possible values are `INGRESS` and `EGRESS`.
	Direction pulumi.StringPtrInput
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging pulumi.BoolPtrInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// Structure is documented below.
	Match OrganizationSecurityPolicyRuleMatchPtrInput
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId pulumi.StringPtrInput
	// If set to true, the specified action is not enforced.
	Preview pulumi.BoolPtrInput
	// An integer indicating the priority of a rule in the list. The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntPtrInput
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources pulumi.StringArrayInput
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (OrganizationSecurityPolicyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyRuleState)(nil)).Elem()
}

type organizationSecurityPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "gotoNext".
	Action string `pulumi:"action"`
	// A description of the rule.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created.
	// Possible values are `INGRESS` and `EGRESS`.
	Direction *string `pulumi:"direction"`
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging *bool `pulumi:"enableLogging"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// Structure is documented below.
	Match OrganizationSecurityPolicyRuleMatch `pulumi:"match"`
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId string `pulumi:"policyId"`
	// If set to true, the specified action is not enforced.
	Preview *bool `pulumi:"preview"`
	// An integer indicating the priority of a rule in the list. The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	Priority int `pulumi:"priority"`
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources []string `pulumi:"targetResources"`
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

// The set of arguments for constructing a OrganizationSecurityPolicyRule resource.
type OrganizationSecurityPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either
	// "allow", "deny" or "gotoNext".
	Action pulumi.StringInput
	// A description of the rule.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. If unspecified an INGRESS rule is created.
	// Possible values are `INGRESS` and `EGRESS`.
	Direction pulumi.StringPtrInput
	// Denotes whether to enable logging for a particular rule.
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	EnableLogging pulumi.BoolPtrInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	// Structure is documented below.
	Match OrganizationSecurityPolicyRuleMatchInput
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	PolicyId pulumi.StringInput
	// If set to true, the specified action is not enforced.
	Preview pulumi.BoolPtrInput
	// An integer indicating the priority of a rule in the list. The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntInput
	// A list of network resource URLs to which this rule applies.
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	TargetResources pulumi.StringArrayInput
	// A list of service accounts indicating the sets of
	// instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (OrganizationSecurityPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyRuleArgs)(nil)).Elem()
}
