// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Storage Bucket. Each of these resources serves a different use case:
//
// * `storage.BucketIAMPolicy`: Authoritative. Sets the IAM policy for the bucket and replaces any existing policy already attached.
// * `storage.BucketIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the bucket are preserved.
// * `storage.BucketIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the bucket are preserved.
//
// > **Note:** `storage.BucketIAMPolicy` **cannot** be used in conjunction with `storage.BucketIAMBinding` and `storage.BucketIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `storage.BucketIAMBinding` resources **can be** used in conjunction with `storage.BucketIAMMember` resources **only if** they do not grant privilege to the same role.
type BucketIAMMember struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBucketIAMMember registers a new resource with the given unique name, arguments, and options.
func NewBucketIAMMember(ctx *pulumi.Context,
	name string, args *BucketIAMMemberArgs, opts ...pulumi.ResourceOption) (*BucketIAMMember, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &BucketIAMMemberArgs{}
	}
	var resource BucketIAMMember
	err := ctx.RegisterResource("gcp:storage/bucketIAMMember:BucketIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIAMMember gets an existing BucketIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIAMMemberState, opts ...pulumi.ResourceOption) (*BucketIAMMember, error) {
	var resource BucketIAMMember
	err := ctx.ReadResource("gcp:storage/bucketIAMMember:BucketIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIAMMember resources.
type bucketIAMMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BucketIAMMemberState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BucketIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMMemberState)(nil)).Elem()
}

type bucketIAMMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMMemberCondition `pulumi:"condition"`
	Member    string                    `pulumi:"member"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BucketIAMMember resource.
type BucketIAMMemberArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BucketIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMMemberArgs)(nil)).Elem()
}
