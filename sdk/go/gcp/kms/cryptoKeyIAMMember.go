// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
//
// * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
// * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
// * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
//
// > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
//
// With IAM Conditions:
//
// With IAM Conditions:
//
// With IAM Conditions:
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
//
// These bindings can be imported using the `crypto_key_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `crypto_key_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember crypto_key your-project-id/location-name/key-ring-name/key-name
// ```
type CryptoKeyIAMMember struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrOutput `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringOutput `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewCryptoKeyIAMMember registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, args *CryptoKeyIAMMemberArgs, opts ...pulumi.ResourceOption) (*CryptoKeyIAMMember, error) {
	if args == nil || args.CryptoKeyId == nil {
		return nil, errors.New("missing required argument 'CryptoKeyId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &CryptoKeyIAMMemberArgs{}
	}
	var resource CryptoKeyIAMMember
	err := ctx.RegisterResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKeyIAMMember gets an existing CryptoKeyIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyIAMMemberState, opts ...pulumi.ResourceOption) (*CryptoKeyIAMMember, error) {
	var resource CryptoKeyIAMMember
	err := ctx.ReadResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKeyIAMMember resources.
type cryptoKeyIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMMemberCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId *string `pulumi:"cryptoKeyId"`
	// (Computed) The etag of the project's IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type CryptoKeyIAMMemberState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringPtrInput
	// (Computed) The etag of the project's IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (CryptoKeyIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMMemberState)(nil)).Elem()
}

type cryptoKeyIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *CryptoKeyIAMMemberCondition `pulumi:"condition"`
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId string `pulumi:"cryptoKeyId"`
	Member      string `pulumi:"member"`
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a CryptoKeyIAMMember resource.
type CryptoKeyIAMMemberArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition CryptoKeyIAMMemberConditionPtrInput
	// The crypto key ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId pulumi.StringInput
	Member      pulumi.StringInput
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (CryptoKeyIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyIAMMemberArgs)(nil)).Elem()
}

type CryptoKeyIAMMemberInput interface {
	pulumi.Input

	ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput
	ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput
}

func (CryptoKeyIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMMember)(nil)).Elem()
}

func (i CryptoKeyIAMMember) ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput {
	return i.ToCryptoKeyIAMMemberOutputWithContext(context.Background())
}

func (i CryptoKeyIAMMember) ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyIAMMemberOutput)
}

type CryptoKeyIAMMemberOutput struct {
	*pulumi.OutputState
}

func (CryptoKeyIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKeyIAMMemberOutput)(nil)).Elem()
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberOutput() CryptoKeyIAMMemberOutput {
	return o
}

func (o CryptoKeyIAMMemberOutput) ToCryptoKeyIAMMemberOutputWithContext(ctx context.Context) CryptoKeyIAMMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CryptoKeyIAMMemberOutput{})
}
