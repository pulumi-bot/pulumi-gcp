// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
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
//
// ## google\_storage\_bucket\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/storage.admin",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketIAMPolicy(ctx, "policy", &storage.BucketIAMPolicyArgs{
// 			Bucket:     pulumi.String(google_storage_bucket.Default.Name),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/storage.admin",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 					"condition": map[string]interface{}{
// 						"title":       "expires_after_2019_12_31",
// 						"description": "Expiring at midnight of 2019-12-31",
// 						"expression":  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketIAMPolicy(ctx, "policy", &storage.BucketIAMPolicyArgs{
// 			Bucket:     pulumi.String(google_storage_bucket.Default.Name),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_storage\_bucket\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = storage.NewBucketIAMBinding(ctx, "binding", &storage.BucketIAMBindingArgs{
// 			Bucket: pulumi.String(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = storage.NewBucketIAMBinding(ctx, "binding", &storage.BucketIAMBindingArgs{
// 			Bucket: pulumi.String(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &storage.BucketIAMBindingConditionArgs{
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
// ## google\_storage\_bucket\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = storage.NewBucketIAMMember(ctx, "member", &storage.BucketIAMMemberArgs{
// 			Bucket: pulumi.String(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Member: pulumi.String("user:jane@example.com"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = storage.NewBucketIAMMember(ctx, "member", &storage.BucketIAMMemberArgs{
// 			Bucket: pulumi.String(google_storage_bucket.Default.Name),
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Member: pulumi.String("user:jane@example.com"),
// 			Condition: &storage.BucketIAMMemberConditionArgs{
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
type BucketIAMBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBucketIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewBucketIAMBinding(ctx *pulumi.Context,
	name string, args *BucketIAMBindingArgs, opts ...pulumi.ResourceOption) (*BucketIAMBinding, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &BucketIAMBindingArgs{}
	}
	var resource BucketIAMBinding
	err := ctx.RegisterResource("gcp:storage/bucketIAMBinding:BucketIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIAMBinding gets an existing BucketIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIAMBindingState, opts ...pulumi.ResourceOption) (*BucketIAMBinding, error) {
	var resource BucketIAMBinding
	err := ctx.ReadResource("gcp:storage/bucketIAMBinding:BucketIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIAMBinding resources.
type bucketIAMBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BucketIAMBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BucketIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMBindingState)(nil)).Elem()
}

type bucketIAMBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMBindingCondition `pulumi:"condition"`
	Members   []string                   `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BucketIAMBinding resource.
type BucketIAMBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BucketIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMBindingArgs)(nil)).Elem()
}
