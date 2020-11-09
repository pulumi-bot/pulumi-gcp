// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy Web. Each of these resources serves a different use case:
//
// * `iap.WebIamPolicy`: Authoritative. Sets the IAM policy for the web and replaces any existing policy already attached.
// * `iap.WebIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the web are preserved.
// * `iap.WebIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the web are preserved.
//
// > **Note:** `iap.WebIamPolicy` **cannot** be used in conjunction with `iap.WebIamBinding` and `iap.WebIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebIamBinding` resources **can be** used in conjunction with `iap.WebIamMember` resources **only if** they do not grant privilege to the same role.
type WebIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebIamPolicy(ctx *pulumi.Context,
	name string, args *WebIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &WebIamPolicyArgs{}
	}
	var resource WebIamPolicy
	err := ctx.RegisterResource("gcp:iap/webIamPolicy:WebIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebIamPolicy gets an existing WebIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebIamPolicyState, opts ...pulumi.ResourceOption) (*WebIamPolicy, error) {
	var resource WebIamPolicy
	err := ctx.ReadResource("gcp:iap/webIamPolicy:WebIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebIamPolicy resources.
type webIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamPolicyState)(nil)).Elem()
}

type webIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebIamPolicy resource.
type WebIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webIamPolicyArgs)(nil)).Elem()
}

type WebIamPolicyInput interface {
	pulumi.Input

	ToWebIamPolicyOutput() WebIamPolicyOutput
	ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput
}

func (WebIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamPolicy)(nil)).Elem()
}

func (i WebIamPolicy) ToWebIamPolicyOutput() WebIamPolicyOutput {
	return i.ToWebIamPolicyOutputWithContext(context.Background())
}

func (i WebIamPolicy) ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebIamPolicyOutput)
}

type WebIamPolicyOutput struct {
	*pulumi.OutputState
}

func (WebIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebIamPolicyOutput)(nil)).Elem()
}

func (o WebIamPolicyOutput) ToWebIamPolicyOutput() WebIamPolicyOutput {
	return o
}

func (o WebIamPolicyOutput) ToWebIamPolicyOutputWithContext(ctx context.Context) WebIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebIamPolicyOutput{})
}
