// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
type BucketIAMPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewBucketIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketIAMPolicy(ctx *pulumi.Context,
	name string, args *BucketIAMPolicyArgs, opts ...pulumi.ResourceOption) (*BucketIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource BucketIAMPolicy
	err := ctx.RegisterResource("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIAMPolicy gets an existing BucketIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIAMPolicyState, opts ...pulumi.ResourceOption) (*BucketIAMPolicy, error) {
	var resource BucketIAMPolicy
	err := ctx.ReadResource("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIAMPolicy resources.
type bucketIAMPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type BucketIAMPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (BucketIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMPolicyState)(nil)).Elem()
}

type bucketIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a BucketIAMPolicy resource.
type BucketIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (BucketIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMPolicyArgs)(nil)).Elem()
}
