// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows creation and management of a single member for a single binding within
// the IAM policy for an existing Google Cloud KMS crypto key.
// 
// ~> **Note:** This resource _must not_ be used in conjunction with
//    `google_kms_crypto_key_iam_policy` or they will fight over what your policy
//    should be. Similarly, roles controlled by `google_kms_crypto_key_iam_binding`
//    should not be assigned to using `google_kms_crypto_key_iam_member`.
type CryptoKeyIAMMember struct {
	s *pulumi.ResourceState
}

// NewCryptoKeyIAMMember registers a new resource with the given unique name, arguments, and options.
func NewCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, args *CryptoKeyIAMMemberArgs, opts ...pulumi.ResourceOpt) (*CryptoKeyIAMMember, error) {
	if args == nil || args.CryptoKeyId == nil {
		return nil, errors.New("missing required argument 'CryptoKeyId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cryptoKeyId"] = nil
		inputs["member"] = nil
		inputs["role"] = nil
	} else {
		inputs["cryptoKeyId"] = args.CryptoKeyId
		inputs["member"] = args.Member
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyIAMMember{s: s}, nil
}

// GetCryptoKeyIAMMember gets an existing CryptoKeyIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKeyIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CryptoKeyIAMMemberState, opts ...pulumi.ResourceOpt) (*CryptoKeyIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cryptoKeyId"] = state.CryptoKeyId
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:kms/cryptoKeyIAMMember:CryptoKeyIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CryptoKeyIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CryptoKeyIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CryptoKeyIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The key ring ID, in the form
// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
// the provider's project setting will be used as a fallback.
func (r *CryptoKeyIAMMember) CryptoKeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cryptoKeyId"])
}

// (Computed) The etag of the project's IAM policy.
func (r *CryptoKeyIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The user that the role should apply to.
func (r *CryptoKeyIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The role that should be applied. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *CryptoKeyIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering CryptoKeyIAMMember resources.
type CryptoKeyIAMMemberState struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId interface{}
	// (Computed) The etag of the project's IAM policy.
	Etag interface{}
	// The user that the role should apply to.
	Member interface{}
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a CryptoKeyIAMMember resource.
type CryptoKeyIAMMemberArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
	// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
	// the provider's project setting will be used as a fallback.
	CryptoKeyId interface{}
	// The user that the role should apply to.
	Member interface{}
	// The role that should be applied. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
