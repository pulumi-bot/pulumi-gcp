// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/PolicyAdmissionWhitelistPattern"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/PolicyClusterAdmissionRule"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/PolicyDefaultAdmissionRule"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_policy.html.markdown.
type Policy struct {
	pulumi.CustomResourceState

	// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the
	// image's admission requests will always be permitted regardless of your admission rules.
	AdmissionWhitelistPatterns binaryauthorizationPolicyAdmissionWhitelistPattern.PolicyAdmissionWhitelistPatternArrayOutput `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request
	// must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be
	// denied. There can be at most one admission rule per cluster spec. Identifier format: '{{location}}.{{clusterId}}'. A
	// location is either a compute zone (e.g. 'us-central1-a') or a region (e.g. 'us-central1').
	ClusterAdmissionRules binaryauthorizationPolicyClusterAdmissionRule.PolicyClusterAdmissionRuleArrayOutput `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission rule.
	DefaultAdmissionRule binaryauthorizationPolicyDefaultAdmissionRule.PolicyDefaultAdmissionRuleOutput `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not
	// covered by the global policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringOutput `pulumi:"globalPolicyEvaluationMode"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.DefaultAdmissionRule == nil {
		return nil, errors.New("missing required argument 'DefaultAdmissionRule'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("gcp:binaryauthorization/policy:Policy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:binaryauthorization/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the
	// image's admission requests will always be permitted regardless of your admission rules.
	AdmissionWhitelistPatterns []binaryauthorizationPolicyAdmissionWhitelistPattern.PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request
	// must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be
	// denied. There can be at most one admission rule per cluster spec. Identifier format: '{{location}}.{{clusterId}}'. A
	// location is either a compute zone (e.g. 'us-central1-a') or a region (e.g. 'us-central1').
	ClusterAdmissionRules []binaryauthorizationPolicyClusterAdmissionRule.PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission rule.
	DefaultAdmissionRule *binaryauthorizationPolicyDefaultAdmissionRule.PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not
	// covered by the global policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	Project *string `pulumi:"project"`
}

type PolicyState struct {
	// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the
	// image's admission requests will always be permitted regardless of your admission rules.
	AdmissionWhitelistPatterns binaryauthorizationPolicyAdmissionWhitelistPattern.PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request
	// must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be
	// denied. There can be at most one admission rule per cluster spec. Identifier format: '{{location}}.{{clusterId}}'. A
	// location is either a compute zone (e.g. 'us-central1-a') or a region (e.g. 'us-central1').
	ClusterAdmissionRules binaryauthorizationPolicyClusterAdmissionRule.PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission rule.
	DefaultAdmissionRule binaryauthorizationPolicyDefaultAdmissionRule.PolicyDefaultAdmissionRulePtrInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not
	// covered by the global policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the
	// image's admission requests will always be permitted regardless of your admission rules.
	AdmissionWhitelistPatterns []binaryauthorizationPolicyAdmissionWhitelistPattern.PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request
	// must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be
	// denied. There can be at most one admission rule per cluster spec. Identifier format: '{{location}}.{{clusterId}}'. A
	// location is either a compute zone (e.g. 'us-central1-a') or a region (e.g. 'us-central1').
	ClusterAdmissionRules []binaryauthorizationPolicyClusterAdmissionRule.PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission rule.
	DefaultAdmissionRule binaryauthorizationPolicyDefaultAdmissionRule.PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not
	// covered by the global policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the
	// image's admission requests will always be permitted regardless of your admission rules.
	AdmissionWhitelistPatterns binaryauthorizationPolicyAdmissionWhitelistPattern.PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request
	// must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be
	// denied. There can be at most one admission rule per cluster spec. Identifier format: '{{location}}.{{clusterId}}'. A
	// location is either a compute zone (e.g. 'us-central1-a') or a region (e.g. 'us-central1').
	ClusterAdmissionRules binaryauthorizationPolicyClusterAdmissionRule.PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission rule.
	DefaultAdmissionRule binaryauthorizationPolicyDefaultAdmissionRule.PolicyDefaultAdmissionRuleInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not
	// covered by the global policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

