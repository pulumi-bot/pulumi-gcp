// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A policy for container image binary authorization.
//
// To get more information about Policy, see:
//
// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/binary-authorization/)
type Policy struct {
	pulumi.CustomResourceState

	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.  Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayOutput `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayOutput `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.  Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRuleOutput `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringOutput `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
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
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.  Structure is documented below.
	AdmissionWhitelistPatterns []PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.  Structure is documented below.
	DefaultAdmissionRule *PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type PolicyState struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.  Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission
	// rule.  Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRulePtrInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.  Structure is documented below.
	AdmissionWhitelistPatterns []PolicyAdmissionWhitelistPattern `pulumi:"admissionWhitelistPatterns"`
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules []PolicyClusterAdmissionRule `pulumi:"clusterAdmissionRules"`
	// Default admission rule for a cluster without a per-cluster admission
	// rule.  Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRule `pulumi:"defaultAdmissionRule"`
	// A descriptive comment.
	Description *string `pulumi:"description"`
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode *string `pulumi:"globalPolicyEvaluationMode"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A whitelist of image patterns to exclude from admission rules. If an
	// image's name matches a whitelist pattern, the image's admission
	// requests will always be permitted regardless of your admission rules.  Structure is documented below.
	AdmissionWhitelistPatterns PolicyAdmissionWhitelistPatternArrayInput
	// Per-cluster admission rules. An admission rule specifies either that
	// all container images used in a pod creation request must be attested
	// to by one or more attestors, that all pod creations will be allowed,
	// or that all pod creations will be denied. There can be at most one
	// admission rule per cluster spec.
	ClusterAdmissionRules PolicyClusterAdmissionRuleArrayInput
	// Default admission rule for a cluster without a per-cluster admission
	// rule.  Structure is documented below.
	DefaultAdmissionRule PolicyDefaultAdmissionRuleInput
	// A descriptive comment.
	Description pulumi.StringPtrInput
	// Controls the evaluation of a Google-maintained global admission policy
	// for common system-level images. Images not covered by the global
	// policy will be subject to the project admission policy.
	GlobalPolicyEvaluationMode pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
