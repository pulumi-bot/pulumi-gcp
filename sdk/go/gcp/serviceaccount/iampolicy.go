// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the googleProjectIam set of resources.
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
//
// ## google\_service\_account\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iam.serviceAccountUser",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can interact with"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMPolicy(ctx, "admin_account_iam", &serviceAccount.IAMPolicyArgs{
// 			ServiceAccountId: sa.Name,
// 			PolicyData:       pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_account\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMBinding(ctx, "admin_account_iam", &serviceAccount.IAMBindingArgs{
// 			ServiceAccountId: sa.Name,
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that only Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMBinding(ctx, "admin_account_iam", &serviceAccount.IAMBindingArgs{
// 			Condition: &serviceAccount.IAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			ServiceAccountId: sa.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_account\_iam\_member
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := compute.GetDefaultServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMMember(ctx, "admin_account_iam", &serviceAccount.IAMMemberArgs{
// 			ServiceAccountId: sa.Name,
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			Member:           pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMMember(ctx, "gce_default_account_iam", &serviceAccount.IAMMemberArgs{
// 			ServiceAccountId: pulumi.String(_default.Name),
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			Member: sa.Email.ApplyT(func(email string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 			}).(pulumi.StringOutput),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sa, err := serviceAccount.NewAccount(ctx, "sa", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-service-account"),
// 			DisplayName: pulumi.String("A service account that Jane can use"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewIAMMember(ctx, "admin_account_iam", &serviceAccount.IAMMemberArgs{
// 			Condition: &serviceAccount.IAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:           pulumi.String("user:jane@example.com"),
// 			Role:             pulumi.String("roles/iam.serviceAccountUser"),
// 			ServiceAccountId: sa.Name,
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
// Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).
//
// ```sh
//  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
// ```
//
// ```sh
//  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser"
// ```
//
// ```sh
//  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions
//
// ```sh
//  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
// ```
//
// ```sh
//  $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
// ```
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

type IAMPolicyInput interface {
	pulumi.Input

	ToIAMPolicyOutput() IAMPolicyOutput
	ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput
}

func (*IAMPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicy)(nil))
}

func (i *IAMPolicy) ToIAMPolicyOutput() IAMPolicyOutput {
	return i.ToIAMPolicyOutputWithContext(context.Background())
}

func (i *IAMPolicy) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyOutput)
}

func (i *IAMPolicy) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return i.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *IAMPolicy) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyPtrOutput)
}

type IAMPolicyPtrInput interface {
	pulumi.Input

	ToIAMPolicyPtrOutput() IAMPolicyPtrOutput
	ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput
}

type iampolicyPtrType IAMPolicyArgs

func (*iampolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil))
}

func (i *iampolicyPtrType) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return i.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (i *iampolicyPtrType) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyPtrOutput)
}

// IAMPolicyArrayInput is an input type that accepts IAMPolicyArray and IAMPolicyArrayOutput values.
// You can construct a concrete instance of `IAMPolicyArrayInput` via:
//
//          IAMPolicyArray{ IAMPolicyArgs{...} }
type IAMPolicyArrayInput interface {
	pulumi.Input

	ToIAMPolicyArrayOutput() IAMPolicyArrayOutput
	ToIAMPolicyArrayOutputWithContext(context.Context) IAMPolicyArrayOutput
}

type IAMPolicyArray []IAMPolicyInput

func (IAMPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IAMPolicy)(nil))
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return i.ToIAMPolicyArrayOutputWithContext(context.Background())
}

func (i IAMPolicyArray) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyArrayOutput)
}

// IAMPolicyMapInput is an input type that accepts IAMPolicyMap and IAMPolicyMapOutput values.
// You can construct a concrete instance of `IAMPolicyMapInput` via:
//
//          IAMPolicyMap{ "key": IAMPolicyArgs{...} }
type IAMPolicyMapInput interface {
	pulumi.Input

	ToIAMPolicyMapOutput() IAMPolicyMapOutput
	ToIAMPolicyMapOutputWithContext(context.Context) IAMPolicyMapOutput
}

type IAMPolicyMap map[string]IAMPolicyInput

func (IAMPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IAMPolicy)(nil))
}

func (i IAMPolicyMap) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return i.ToIAMPolicyMapOutputWithContext(context.Background())
}

func (i IAMPolicyMap) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMPolicyMapOutput)
}

type IAMPolicyOutput struct {
	*pulumi.OutputState
}

func (IAMPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMPolicy)(nil))
}

func (o IAMPolicyOutput) ToIAMPolicyOutput() IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyOutputWithContext(ctx context.Context) IAMPolicyOutput {
	return o
}

func (o IAMPolicyOutput) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return o.ToIAMPolicyPtrOutputWithContext(context.Background())
}

func (o IAMPolicyOutput) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return o.ApplyT(func(v IAMPolicy) *IAMPolicy {
		return &v
	}).(IAMPolicyPtrOutput)
}

type IAMPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (IAMPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMPolicy)(nil))
}

func (o IAMPolicyPtrOutput) ToIAMPolicyPtrOutput() IAMPolicyPtrOutput {
	return o
}

func (o IAMPolicyPtrOutput) ToIAMPolicyPtrOutputWithContext(ctx context.Context) IAMPolicyPtrOutput {
	return o
}

type IAMPolicyArrayOutput struct{ *pulumi.OutputState }

func (IAMPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IAMPolicy)(nil))
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutput() IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) ToIAMPolicyArrayOutputWithContext(ctx context.Context) IAMPolicyArrayOutput {
	return o
}

func (o IAMPolicyArrayOutput) Index(i pulumi.IntInput) IAMPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IAMPolicy {
		return vs[0].([]IAMPolicy)[vs[1].(int)]
	}).(IAMPolicyOutput)
}

type IAMPolicyMapOutput struct{ *pulumi.OutputState }

func (IAMPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IAMPolicy)(nil))
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutput() IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) ToIAMPolicyMapOutputWithContext(ctx context.Context) IAMPolicyMapOutput {
	return o
}

func (o IAMPolicyMapOutput) MapIndex(k pulumi.StringInput) IAMPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IAMPolicy {
		return vs[0].(map[string]IAMPolicy)[vs[1].(string)]
	}).(IAMPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(IAMPolicyOutput{})
	pulumi.RegisterOutputType(IAMPolicyPtrOutput{})
	pulumi.RegisterOutputType(IAMPolicyArrayOutput{})
	pulumi.RegisterOutputType(IAMPolicyMapOutput{})
}
