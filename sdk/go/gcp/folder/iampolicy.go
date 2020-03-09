// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package folder

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of the IAM policy for an existing Google Cloud
// Platform folder.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_folder_iam_policy.html.markdown.
type IAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing
	// policy applied to the folder.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil || args.Folder == nil {
		return nil, errors.New("missing required argument 'Folder'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &IAMPolicyArgs{}
	}
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:folder/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMPolicy gets an existing IAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMPolicyState, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	var resource IAMPolicy
	err := ctx.ReadResource("gcp:folder/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	// (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag *string `pulumi:"etag"`
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder *string `pulumi:"folder"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing
	// policy applied to the folder.
	PolicyData *string `pulumi:"policyData"`
}

type IAMPolicyState struct {
	// (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag pulumi.StringPtrInput
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringPtrInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing
	// policy applied to the folder.
	PolicyData pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder string `pulumi:"folder"`
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing
	// policy applied to the folder.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
	Folder pulumi.StringInput
	// The `organizations.getIAMPolicy` data source that represents
	// the IAM policy that will be applied to the folder. This policy overrides any existing
	// policy applied to the folder.
	PolicyData pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}

