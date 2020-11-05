// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the googleProjectIam set of resources.
//
// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
//
// * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
// * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
// * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
//
// > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.
//
// > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
type IAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the service account IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
}

// NewIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewIAMPolicy(ctx *pulumi.Context,
	name string, args *IAMPolicyArgs, opts ...pulumi.ResourceOption) (*IAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	var resource IAMPolicy
	err := ctx.RegisterResource("gcp:serviceAccount/iAMPolicy:IAMPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:serviceAccount/iAMPolicy:IAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMPolicy resources.
type iampolicyState struct {
	// (Computed) The etag of the service account IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId *string `pulumi:"serviceAccountId"`
}

type IAMPolicyState struct {
	// (Computed) The etag of the service account IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringPtrInput
}

func (IAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyState)(nil)).Elem()
}

type iampolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a IAMPolicy resource.
type IAMPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The fully-qualified name of the service account to apply policy to.
	ServiceAccountId pulumi.StringInput
}

func (IAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iampolicyArgs)(nil)).Elem()
}
