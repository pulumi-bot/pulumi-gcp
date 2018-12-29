// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
// 
// * `google_kms_key_ring_iam_policy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
// * `google_kms_key_ring_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
// * `google_kms_key_ring_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
// 
// > **Note:** `google_kms_key_ring_iam_policy` **cannot** be used in conjunction with `google_kms_key_ring_iam_binding` and `google_kms_key_ring_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_kms_key_ring_iam_binding` resources **can be** used in conjunction with `google_kms_key_ring_iam_member` resources **only if** they do not grant privilege to the same role.
type KeyRingIAMBinding struct {
	s *pulumi.ResourceState
}

// NewKeyRingIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMBinding(ctx *pulumi.Context,
	name string, args *KeyRingIAMBindingArgs, opts ...pulumi.ResourceOpt) (*KeyRingIAMBinding, error) {
	if args == nil || args.KeyRingId == nil {
		return nil, errors.New("missing required argument 'KeyRingId'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keyRingId"] = nil
		inputs["members"] = nil
		inputs["role"] = nil
	} else {
		inputs["keyRingId"] = args.KeyRingId
		inputs["members"] = args.Members
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRingIAMBinding{s: s}, nil
}

// GetKeyRingIAMBinding gets an existing KeyRingIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyRingIAMBindingState, opts ...pulumi.ResourceOpt) (*KeyRingIAMBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["keyRingId"] = state.KeyRingId
		inputs["members"] = state.Members
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRingIAMBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KeyRingIAMBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KeyRingIAMBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the key ring's IAM policy.
func (r *KeyRingIAMBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The key ring ID, in the form
// `{project_id}/{location_name}/{key_ring_name}` or
// `{location_name}/{key_ring_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *KeyRingIAMBinding) KeyRingId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyRingId"])
}

func (r *KeyRingIAMBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The role that should be applied. Only one
// `google_kms_key_ring_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *KeyRingIAMBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering KeyRingIAMBinding resources.
type KeyRingIAMBindingState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag interface{}
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId interface{}
	Members interface{}
	// The role that should be applied. Only one
	// `google_kms_key_ring_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a KeyRingIAMBinding resource.
type KeyRingIAMBindingArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId interface{}
	Members interface{}
	// The role that should be applied. Only one
	// `google_kms_key_ring_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
