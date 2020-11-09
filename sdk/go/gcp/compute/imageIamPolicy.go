// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Compute Engine Image. Each of these resources serves a different use case:
//
// * `compute.ImageIamPolicy`: Authoritative. Sets the IAM policy for the image and replaces any existing policy already attached.
// * `compute.ImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the image are preserved.
// * `compute.ImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the image are preserved.
//
// > **Note:** `compute.ImageIamPolicy` **cannot** be used in conjunction with `compute.ImageIamBinding` and `compute.ImageIamMember` or they will fight over what your policy should be.
//
// > **Note:** `compute.ImageIamBinding` resources **can be** used in conjunction with `compute.ImageIamMember` resources **only if** they do not grant privilege to the same role.
type ImageIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image pulumi.StringOutput `pulumi:"image"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewImageIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewImageIamPolicy(ctx *pulumi.Context,
	name string, args *ImageIamPolicyArgs, opts ...pulumi.ResourceOption) (*ImageIamPolicy, error) {
	if args == nil || args.Image == nil {
		return nil, errors.New("missing required argument 'Image'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &ImageIamPolicyArgs{}
	}
	var resource ImageIamPolicy
	err := ctx.RegisterResource("gcp:compute/imageIamPolicy:ImageIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageIamPolicy gets an existing ImageIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageIamPolicyState, opts ...pulumi.ResourceOption) (*ImageIamPolicy, error) {
	var resource ImageIamPolicy
	err := ctx.ReadResource("gcp:compute/imageIamPolicy:ImageIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageIamPolicy resources.
type imageIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	Image *string `pulumi:"image"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type ImageIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Image pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ImageIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamPolicyState)(nil)).Elem()
}

type imageIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Image string `pulumi:"image"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ImageIamPolicy resource.
type ImageIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Image pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ImageIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageIamPolicyArgs)(nil)).Elem()
}

type ImageIamPolicyInput interface {
	pulumi.Input

	ToImageIamPolicyOutput() ImageIamPolicyOutput
	ToImageIamPolicyOutputWithContext(ctx context.Context) ImageIamPolicyOutput
}

func (ImageIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamPolicy)(nil)).Elem()
}

func (i ImageIamPolicy) ToImageIamPolicyOutput() ImageIamPolicyOutput {
	return i.ToImageIamPolicyOutputWithContext(context.Background())
}

func (i ImageIamPolicy) ToImageIamPolicyOutputWithContext(ctx context.Context) ImageIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageIamPolicyOutput)
}

type ImageIamPolicyOutput struct {
	*pulumi.OutputState
}

func (ImageIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageIamPolicyOutput)(nil)).Elem()
}

func (o ImageIamPolicyOutput) ToImageIamPolicyOutput() ImageIamPolicyOutput {
	return o
}

func (o ImageIamPolicyOutput) ToImageIamPolicyOutputWithContext(ctx context.Context) ImageIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImageIamPolicyOutput{})
}
