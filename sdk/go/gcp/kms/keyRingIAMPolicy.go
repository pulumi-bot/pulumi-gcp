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
// ~> **Note:** `google_kms_key_ring_iam_policy` **cannot** be used in conjunction with `google_kms_key_ring_iam_binding` and `google_kms_key_ring_iam_member` or they will fight over what your policy should be.
// 
// ~> **Note:** `google_kms_key_ring_iam_binding` resources **can be** used in conjunction with `google_kms_key_ring_iam_member` resources **only if** they do not grant privilege to the same role.
type KeyRingIAMPolicy struct {
	s *pulumi.ResourceState
}

// NewKeyRingIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, args *KeyRingIAMPolicyArgs, opts ...pulumi.ResourceOpt) (*KeyRingIAMPolicy, error) {
	if args == nil || args.KeyRingId == nil {
		return nil, errors.New("missing required argument 'KeyRingId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keyRingId"] = nil
		inputs["policyData"] = nil
	} else {
		inputs["keyRingId"] = args.KeyRingId
		inputs["policyData"] = args.PolicyData
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRingIAMPolicy{s: s}, nil
}

// GetKeyRingIAMPolicy gets an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KeyRingIAMPolicyState, opts ...pulumi.ResourceOpt) (*KeyRingIAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["keyRingId"] = state.KeyRingId
		inputs["policyData"] = state.PolicyData
	}
	s, err := ctx.ReadResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KeyRingIAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KeyRingIAMPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KeyRingIAMPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the key ring's IAM policy.
func (r *KeyRingIAMPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The key ring ID, in the form
// `{project_id}/{location_name}/{key_ring_name}` or
// `{location_name}/{key_ring_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *KeyRingIAMPolicy) KeyRingId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyRingId"])
}

// The policy data generated by
// a `google_iam_policy` data source.
func (r *KeyRingIAMPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// Input properties used for looking up and filtering KeyRingIAMPolicy resources.
type KeyRingIAMPolicyState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag interface{}
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}

// The set of arguments for constructing a KeyRingIAMPolicy resource.
type KeyRingIAMPolicyArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}
