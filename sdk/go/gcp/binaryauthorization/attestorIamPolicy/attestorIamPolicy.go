// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package attestorIamPolicy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_attestor_iam_policy.html.markdown.
type AttestorIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringOutput `pulumi:"attestor"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewAttestorIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamPolicy(ctx *pulumi.Context,
	name string, args *AttestorIamPolicyArgs, opts ...pulumi.ResourceOption) (*AttestorIamPolicy, error) {
	if args == nil || args.Attestor == nil {
		return nil, errors.New("missing required argument 'Attestor'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &AttestorIamPolicyArgs{}
	}
	var resource AttestorIamPolicy
	err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamPolicy:AttestorIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestorIamPolicy gets an existing AttestorIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorIamPolicyState, opts ...pulumi.ResourceOption) (*AttestorIamPolicy, error) {
	var resource AttestorIamPolicy
	err := ctx.ReadResource("gcp:binaryauthorization/attestorIamPolicy:AttestorIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestorIamPolicy resources.
type attestorIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor *string `pulumi:"attestor"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type AttestorIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamPolicyState)(nil)).Elem()
}

type attestorIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor string `pulumi:"attestor"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AttestorIamPolicy resource.
type AttestorIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorIamPolicyArgs)(nil)).Elem()
}

