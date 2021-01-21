// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Security Policy defines an IP blacklist or whitelist that protects load balanced Google Cloud services by denying or permitting traffic from specified IP ranges. For more information
// see the [official documentation](https://cloud.google.com/armor/docs/configure-security-policies)
// and the [API](https://cloud.google.com/compute/docs/reference/rest/beta/securityPolicies).
//
// ## Example Usage
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
// 		_, err := compute.NewSecurityPolicy(ctx, "policy", &compute.SecurityPolicyArgs{
// 			Rules: compute.SecurityPolicyRuleArray{
// 				&compute.SecurityPolicyRuleArgs{
// 					Action:      pulumi.String("deny(403)"),
// 					Description: pulumi.String("Deny access to IPs in 9.9.9.0/24"),
// 					Match: &compute.SecurityPolicyRuleMatchArgs{
// 						Config: &compute.SecurityPolicyRuleMatchConfigArgs{
// 							SrcIpRanges: pulumi.StringArray{
// 								pulumi.String("9.9.9.0/24"),
// 							},
// 						},
// 						VersionedExpr: pulumi.String("SRC_IPS_V1"),
// 					},
// 					Priority: pulumi.Int(1000),
// 				},
// 				&compute.SecurityPolicyRuleArgs{
// 					Action:      pulumi.String("allow"),
// 					Description: pulumi.String("default rule"),
// 					Match: &compute.SecurityPolicyRuleMatchArgs{
// 						Config: &compute.SecurityPolicyRuleMatchConfigArgs{
// 							SrcIpRanges: pulumi.StringArray{
// 								pulumi.String("*"),
// 							},
// 						},
// 						VersionedExpr: pulumi.String("SRC_IPS_V1"),
// 					},
// 					Priority: pulumi.Int(2147483647),
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
// Security policies can be imported using any of the following formats
//
// ```sh
//  $ pulumi import gcp:compute/securityPolicy:SecurityPolicy policy projects/{{project}}/global/securityPolicies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/securityPolicy:SecurityPolicy policy {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/securityPolicy:SecurityPolicy policy {{name}}
// ```
type SecurityPolicy struct {
	pulumi.CustomResourceState

	// An optional description of this rule. Max size is 64.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the security policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rules SecurityPolicyRuleArrayOutput `pulumi:"rules"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecurityPolicy(ctx *pulumi.Context,
	name string, args *SecurityPolicyArgs, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	if args == nil {
		args = &SecurityPolicyArgs{}
	}

	var resource SecurityPolicy
	err := ctx.RegisterResource("gcp:compute/securityPolicy:SecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityPolicy gets an existing SecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityPolicyState, opts ...pulumi.ResourceOption) (*SecurityPolicy, error) {
	var resource SecurityPolicy
	err := ctx.ReadResource("gcp:compute/securityPolicy:SecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityPolicy resources.
type securityPolicyState struct {
	// An optional description of this rule. Max size is 64.
	Description *string `pulumi:"description"`
	// Fingerprint of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the security policy.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rules []SecurityPolicyRule `pulumi:"rules"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type SecurityPolicyState struct {
	// An optional description of this rule. Max size is 64.
	Description pulumi.StringPtrInput
	// Fingerprint of this resource.
	Fingerprint pulumi.StringPtrInput
	// The name of the security policy.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rules SecurityPolicyRuleArrayInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (SecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyState)(nil)).Elem()
}

type securityPolicyArgs struct {
	// An optional description of this rule. Max size is 64.
	Description *string `pulumi:"description"`
	// The name of the security policy.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rules []SecurityPolicyRule `pulumi:"rules"`
}

// The set of arguments for constructing a SecurityPolicy resource.
type SecurityPolicyArgs struct {
	// An optional description of this rule. Max size is 64.
	Description pulumi.StringPtrInput
	// The name of the security policy.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The set of rules that belong to this policy. There must always be a default
	// rule (rule with priority 2147483647 and match "\*"). If no rules are provided when creating a
	// security policy, a default rule with action "allow" will be added. Structure is documented below.
	Rules SecurityPolicyRuleArrayInput
}

func (SecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityPolicyArgs)(nil)).Elem()
}

type SecurityPolicyInput interface {
	pulumi.Input

	ToSecurityPolicyOutput() SecurityPolicyOutput
	ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput
}

func (*SecurityPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicy)(nil))
}

func (i *SecurityPolicy) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return i.ToSecurityPolicyOutputWithContext(context.Background())
}

func (i *SecurityPolicy) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityPolicyOutput)
}

type SecurityPolicyOutput struct {
	*pulumi.OutputState
}

func (SecurityPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityPolicy)(nil))
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutput() SecurityPolicyOutput {
	return o
}

func (o SecurityPolicyOutput) ToSecurityPolicyOutputWithContext(ctx context.Context) SecurityPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecurityPolicyOutput{})
}
