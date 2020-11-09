// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
//
// * `kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
// * `kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
// * `kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
//
// > **Note:** `kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `kms.KeyRingIAMBinding` and `kms.KeyRingIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.KeyRingIAMBinding` resources **can be** used in conjunction with `kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `key_ring_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMMember:KeyRingIAMMember key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `key_ring_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMMember:KeyRingIAMMember key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `key_ring_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMMember:KeyRingIAMMember key_ring_iam your-project-id/location-name/key-ring-name
// ```
type KeyRingIAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringOutput `pulumi:"keyRingId"`
	Member    pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewKeyRingIAMMember registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMMember(ctx *pulumi.Context,
	name string, args *KeyRingIAMMemberArgs, opts ...pulumi.ResourceOption) (*KeyRingIAMMember, error) {
	if args == nil || args.KeyRingId == nil {
		return nil, errors.New("missing required argument 'KeyRingId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &KeyRingIAMMemberArgs{}
	}
	var resource KeyRingIAMMember
	err := ctx.RegisterResource("gcp:kms/keyRingIAMMember:KeyRingIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRingIAMMember gets an existing KeyRingIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingIAMMemberState, opts ...pulumi.ResourceOption) (*KeyRingIAMMember, error) {
	var resource KeyRingIAMMember
	err := ctx.ReadResource("gcp:kms/keyRingIAMMember:KeyRingIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRingIAMMember resources.
type keyRingIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *KeyRingIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the key ring's IAM policy.
	Etag *string `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId *string `pulumi:"keyRingId"`
	Member    *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type KeyRingIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMMemberConditionPtrInput
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringPtrInput
	Member    pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (KeyRingIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMMemberState)(nil)).Elem()
}

type keyRingIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *KeyRingIAMMemberCondition `pulumi:"condition"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId string `pulumi:"keyRingId"`
	Member    string `pulumi:"member"`
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a KeyRingIAMMember resource.
type KeyRingIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition KeyRingIAMMemberConditionPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringInput
	Member    pulumi.StringInput
	// The role that should be applied. Only one
	// `kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (KeyRingIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMMemberArgs)(nil)).Elem()
}
