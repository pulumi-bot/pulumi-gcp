// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
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
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
// 			Location: pulumi.String("global"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		key, err := kms.NewCryptoKey(ctx, "key", &kms.CryptoKeyArgs{
// 			KeyRing:        keyring.ID(),
// 			RotationPeriod: pulumi.String("100000s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKeyIAMPolicy(ctx, "cryptoKey", &kms.CryptoKeyIAMPolicyArgs{
// 			CryptoKeyId: key.ID(),
// 			PolicyData:  pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/cloudkms.cryptoKeyEncrypter",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKey", &kms.CryptoKeyIAMBindingArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &kms.CryptoKeyIAMBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewCryptoKeyIAMMember(ctx, "cryptoKey", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypter"),
// 			Member:      pulumi.String("user:jane@example.com"),
// 			Condition: &kms.CryptoKeyIAMMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
