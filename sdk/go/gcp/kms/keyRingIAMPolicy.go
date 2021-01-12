// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
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
// ## google\_kms\_key\_ring\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
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
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewKeyRingIAMPolicy(ctx, "keyRing", &kms.KeyRingIAMPolicyArgs{
// 			KeyRingId:  keyring.ID(),
// 			PolicyData: pulumi.String(admin.PolicyData),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
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
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewKeyRingIAMPolicy(ctx, "keyRing", &kms.KeyRingIAMPolicyArgs{
// 			KeyRingId:  keyring.ID(),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_kms\_key\_ring\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewKeyRingIAMBinding(ctx, "keyRing", &kms.KeyRingIAMBindingArgs{
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewKeyRingIAMBinding(ctx, "keyRing", &kms.KeyRingIAMBindingArgs{
// 			Condition: &kms.KeyRingIAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_kms\_key\_ring\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewKeyRingIAMMember(ctx, "keyRing", &kms.KeyRingIAMMemberArgs{
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kms.NewKeyRingIAMMember(ctx, "keyRing", &kms.KeyRingIAMMemberArgs{
// 			Condition: &kms.KeyRingIAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			KeyRingId: pulumi.String("your-key-ring-id"),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Role:      pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `key_ring_id`, role, and account e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `key_ring_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `key_ring_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy key_ring_iam your-project-id/location-name/key-ring-name
// ```
type KeyRingIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringOutput `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewKeyRingIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, args *KeyRingIAMPolicyArgs, opts ...pulumi.ResourceOption) (*KeyRingIAMPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource KeyRingIAMPolicy
	err := ctx.RegisterResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRingIAMPolicy gets an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingIAMPolicyState, opts ...pulumi.ResourceOption) (*KeyRingIAMPolicy, error) {
	var resource KeyRingIAMPolicy
	err := ctx.ReadResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRingIAMPolicy resources.
type keyRingIAMPolicyState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag *string `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId *string `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type KeyRingIAMPolicyState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (KeyRingIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMPolicyState)(nil)).Elem()
}

type keyRingIAMPolicyArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId string `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a KeyRingIAMPolicy resource.
type KeyRingIAMPolicyArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (KeyRingIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMPolicyArgs)(nil)).Elem()
}

type KeyRingIAMPolicyInput interface {
	pulumi.Input

	ToKeyRingIAMPolicyOutput() KeyRingIAMPolicyOutput
	ToKeyRingIAMPolicyOutputWithContext(ctx context.Context) KeyRingIAMPolicyOutput
}

func (*KeyRingIAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRingIAMPolicy)(nil))
}

func (i *KeyRingIAMPolicy) ToKeyRingIAMPolicyOutput() KeyRingIAMPolicyOutput {
	return i.ToKeyRingIAMPolicyOutputWithContext(context.Background())
}

func (i *KeyRingIAMPolicy) ToKeyRingIAMPolicyOutputWithContext(ctx context.Context) KeyRingIAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMPolicyOutput)
}

func (i *KeyRingIAMPolicy) ToKeyRingIAMPolicyPtrOutput() KeyRingIAMPolicyPtrOutput {
	return i.ToKeyRingIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *KeyRingIAMPolicy) ToKeyRingIAMPolicyPtrOutputWithContext(ctx context.Context) KeyRingIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMPolicyPtrOutput)
}

type KeyRingIAMPolicyPtrInput interface {
	pulumi.Input

	ToKeyRingIAMPolicyPtrOutput() KeyRingIAMPolicyPtrOutput
	ToKeyRingIAMPolicyPtrOutputWithContext(ctx context.Context) KeyRingIAMPolicyPtrOutput
}

type keyRingIAMPolicyPtrType KeyRingIAMPolicyArgs

func (*keyRingIAMPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingIAMPolicy)(nil))
}

func (i *keyRingIAMPolicyPtrType) ToKeyRingIAMPolicyPtrOutput() KeyRingIAMPolicyPtrOutput {
	return i.ToKeyRingIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *keyRingIAMPolicyPtrType) ToKeyRingIAMPolicyPtrOutputWithContext(ctx context.Context) KeyRingIAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMPolicyOutput).ToKeyRingIAMPolicyPtrOutput()
}

// KeyRingIAMPolicyArrayInput is an input type that accepts KeyRingIAMPolicyArray and KeyRingIAMPolicyArrayOutput values.
// You can construct a concrete instance of `KeyRingIAMPolicyArrayInput` via:
//
//          KeyRingIAMPolicyArray{ KeyRingIAMPolicyArgs{...} }
type KeyRingIAMPolicyArrayInput interface {
	pulumi.Input

	ToKeyRingIAMPolicyArrayOutput() KeyRingIAMPolicyArrayOutput
	ToKeyRingIAMPolicyArrayOutputWithContext(context.Context) KeyRingIAMPolicyArrayOutput
}

type KeyRingIAMPolicyArray []KeyRingIAMPolicyInput

func (KeyRingIAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyRingIAMPolicy)(nil))
}

func (i KeyRingIAMPolicyArray) ToKeyRingIAMPolicyArrayOutput() KeyRingIAMPolicyArrayOutput {
	return i.ToKeyRingIAMPolicyArrayOutputWithContext(context.Background())
}

func (i KeyRingIAMPolicyArray) ToKeyRingIAMPolicyArrayOutputWithContext(ctx context.Context) KeyRingIAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMPolicyArrayOutput)
}

// KeyRingIAMPolicyMapInput is an input type that accepts KeyRingIAMPolicyMap and KeyRingIAMPolicyMapOutput values.
// You can construct a concrete instance of `KeyRingIAMPolicyMapInput` via:
//
//          KeyRingIAMPolicyMap{ "key": KeyRingIAMPolicyArgs{...} }
type KeyRingIAMPolicyMapInput interface {
	pulumi.Input

	ToKeyRingIAMPolicyMapOutput() KeyRingIAMPolicyMapOutput
	ToKeyRingIAMPolicyMapOutputWithContext(context.Context) KeyRingIAMPolicyMapOutput
}

type KeyRingIAMPolicyMap map[string]KeyRingIAMPolicyInput

func (KeyRingIAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyRingIAMPolicy)(nil))
}

func (i KeyRingIAMPolicyMap) ToKeyRingIAMPolicyMapOutput() KeyRingIAMPolicyMapOutput {
	return i.ToKeyRingIAMPolicyMapOutputWithContext(context.Background())
}

func (i KeyRingIAMPolicyMap) ToKeyRingIAMPolicyMapOutputWithContext(ctx context.Context) KeyRingIAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingIAMPolicyMapOutput)
}

type KeyRingIAMPolicyOutput struct {
	*pulumi.OutputState
}

func (KeyRingIAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRingIAMPolicy)(nil))
}

func (o KeyRingIAMPolicyOutput) ToKeyRingIAMPolicyOutput() KeyRingIAMPolicyOutput {
	return o
}

func (o KeyRingIAMPolicyOutput) ToKeyRingIAMPolicyOutputWithContext(ctx context.Context) KeyRingIAMPolicyOutput {
	return o
}

func (o KeyRingIAMPolicyOutput) ToKeyRingIAMPolicyPtrOutput() KeyRingIAMPolicyPtrOutput {
	return o.ToKeyRingIAMPolicyPtrOutputWithContext(context.Background())
}

func (o KeyRingIAMPolicyOutput) ToKeyRingIAMPolicyPtrOutputWithContext(ctx context.Context) KeyRingIAMPolicyPtrOutput {
	return o.ApplyT(func(v KeyRingIAMPolicy) *KeyRingIAMPolicy {
		return &v
	}).(KeyRingIAMPolicyPtrOutput)
}

type KeyRingIAMPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (KeyRingIAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyRingIAMPolicy)(nil))
}

func (o KeyRingIAMPolicyPtrOutput) ToKeyRingIAMPolicyPtrOutput() KeyRingIAMPolicyPtrOutput {
	return o
}

func (o KeyRingIAMPolicyPtrOutput) ToKeyRingIAMPolicyPtrOutputWithContext(ctx context.Context) KeyRingIAMPolicyPtrOutput {
	return o
}

type KeyRingIAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (KeyRingIAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyRingIAMPolicy)(nil))
}

func (o KeyRingIAMPolicyArrayOutput) ToKeyRingIAMPolicyArrayOutput() KeyRingIAMPolicyArrayOutput {
	return o
}

func (o KeyRingIAMPolicyArrayOutput) ToKeyRingIAMPolicyArrayOutputWithContext(ctx context.Context) KeyRingIAMPolicyArrayOutput {
	return o
}

func (o KeyRingIAMPolicyArrayOutput) Index(i pulumi.IntInput) KeyRingIAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyRingIAMPolicy {
		return vs[0].([]KeyRingIAMPolicy)[vs[1].(int)]
	}).(KeyRingIAMPolicyOutput)
}

type KeyRingIAMPolicyMapOutput struct{ *pulumi.OutputState }

func (KeyRingIAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyRingIAMPolicy)(nil))
}

func (o KeyRingIAMPolicyMapOutput) ToKeyRingIAMPolicyMapOutput() KeyRingIAMPolicyMapOutput {
	return o
}

func (o KeyRingIAMPolicyMapOutput) ToKeyRingIAMPolicyMapOutputWithContext(ctx context.Context) KeyRingIAMPolicyMapOutput {
	return o
}

func (o KeyRingIAMPolicyMapOutput) MapIndex(k pulumi.StringInput) KeyRingIAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyRingIAMPolicy {
		return vs[0].(map[string]KeyRingIAMPolicy)[vs[1].(string)]
	}).(KeyRingIAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyRingIAMPolicyOutput{})
	pulumi.RegisterOutputType(KeyRingIAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(KeyRingIAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(KeyRingIAMPolicyMapOutput{})
}
