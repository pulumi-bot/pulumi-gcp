// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/organizations/PolicyBooleanPolicy"
	"https:/github.com/pulumi/pulumi-gcp/organizations/PolicyListPolicy"
	"https:/github.com/pulumi/pulumi-gcp/organizations/PolicyListPolicyAllow"
	"https:/github.com/pulumi/pulumi-gcp/organizations/PolicyListPolicyDeny"
	"https:/github.com/pulumi/pulumi-gcp/organizations/PolicyRestorePolicy"
)

// Allows management of Organization policies for a Google Organization. For more information see
// [the official
// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
// [API](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/setOrgPolicy).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/organization_policy.html.markdown.
type Policy struct {
	pulumi.CustomResourceState

	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy organizationsPolicyBooleanPolicy.PolicyBooleanPolicyPtrOutput `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringOutput `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy organizationsPolicyListPolicy.PolicyListPolicyPtrOutput `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy organizationsPolicyRestorePolicy.PolicyRestorePolicyPtrOutput `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.Constraint == nil {
		return nil, errors.New("missing required argument 'Constraint'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("gcp:organizations/policy:Policy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:organizations/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *organizationsPolicyBooleanPolicy.PolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint *string `pulumi:"constraint"`
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag *string `pulumi:"etag"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *organizationsPolicyListPolicy.PolicyListPolicy `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId *string `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *organizationsPolicyRestorePolicy.PolicyRestorePolicy `pulumi:"restorePolicy"`
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime *string `pulumi:"updateTime"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

type PolicyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy organizationsPolicyBooleanPolicy.PolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringPtrInput
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringPtrInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy organizationsPolicyListPolicy.PolicyListPolicyPtrInput
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringPtrInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy organizationsPolicyRestorePolicy.PolicyRestorePolicyPtrInput
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime pulumi.StringPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy *organizationsPolicyBooleanPolicy.PolicyBooleanPolicy `pulumi:"booleanPolicy"`
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint string `pulumi:"constraint"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy *organizationsPolicyListPolicy.PolicyListPolicy `pulumi:"listPolicy"`
	// The numeric ID of the organization to set the policy for.
	OrgId string `pulumi:"orgId"`
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy *organizationsPolicyRestorePolicy.PolicyRestorePolicy `pulumi:"restorePolicy"`
	// Version of the Policy. Default version is 0.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy organizationsPolicyBooleanPolicy.PolicyBooleanPolicyPtrInput
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint pulumi.StringInput
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy organizationsPolicyListPolicy.PolicyListPolicyPtrInput
	// The numeric ID of the organization to set the policy for.
	OrgId pulumi.StringInput
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy organizationsPolicyRestorePolicy.PolicyRestorePolicyPtrInput
	// Version of the Policy. Default version is 0.
	Version pulumi.IntPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

