// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
//
// * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.
//
// > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_bigtable\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/bigtable.user",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigtable.NewInstanceIamPolicy(ctx, "editor", &bigtable.InstanceIamPolicyArgs{
// 			Project:    pulumi.String("your-project"),
// 			Instance:   pulumi.String("your-bigtable-instance"),
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
// ## google\_bigtable\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamBinding(ctx, "editor", &bigtable.InstanceIamBindingArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Role: pulumi.String("roles/bigtable.user"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_bigtable\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigtable"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigtable.NewInstanceIamMember(ctx, "editor", &bigtable.InstanceIamMemberArgs{
// 			Instance: pulumi.String("your-bigtable-instance"),
// 			Member:   pulumi.String("user:jane@example.com"),
// 			Role:     pulumi.String("roles/bigtable.user"),
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
// Instance IAM resources can be imported using the project, instance name, role and/or member.
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance}"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor"
// ```
//
// ```sh
//  $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewInstanceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamPolicy(ctx *pulumi.Context,
	name string, args *InstanceIamPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource InstanceIamPolicy
	err := ctx.RegisterResource("gcp:bigtable/instanceIamPolicy:InstanceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamPolicy gets an existing InstanceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamPolicyState, opts ...pulumi.ResourceOption) (*InstanceIamPolicy, error) {
	var resource InstanceIamPolicy
	err := ctx.ReadResource("gcp:bigtable/instanceIamPolicy:InstanceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamPolicy resources.
type instanceIamPolicyState struct {
	// (Computed) The etag of the instances's IAM policy.
	Etag *string `pulumi:"etag"`
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance *string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
}

type InstanceIamPolicyState struct {
	// (Computed) The etag of the instances's IAM policy.
	Etag pulumi.StringPtrInput
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringPtrInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
}

func (InstanceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamPolicyState)(nil)).Elem()
}

type instanceIamPolicyArgs struct {
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance string `pulumi:"instance"`
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a InstanceIamPolicy resource.
type InstanceIamPolicyArgs struct {
	// The name or relative resource id of the instance to manage IAM policies for.
	Instance pulumi.StringInput
	// The policy data generated by a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The project in which the instance belongs. If it
	// is not provided, a default will be supplied.
	Project pulumi.StringPtrInput
}

func (InstanceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamPolicyArgs)(nil)).Elem()
}

type InstanceIamPolicyInput interface {
	pulumi.Input

	ToInstanceIamPolicyOutput() InstanceIamPolicyOutput
	ToInstanceIamPolicyOutputWithContext(ctx context.Context) InstanceIamPolicyOutput
}

func (*InstanceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamPolicy)(nil))
}

func (i *InstanceIamPolicy) ToInstanceIamPolicyOutput() InstanceIamPolicyOutput {
	return i.ToInstanceIamPolicyOutputWithContext(context.Background())
}

func (i *InstanceIamPolicy) ToInstanceIamPolicyOutputWithContext(ctx context.Context) InstanceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyOutput)
}

func (i *InstanceIamPolicy) ToInstanceIamPolicyPtrOutput() InstanceIamPolicyPtrOutput {
	return i.ToInstanceIamPolicyPtrOutputWithContext(context.Background())
}

func (i *InstanceIamPolicy) ToInstanceIamPolicyPtrOutputWithContext(ctx context.Context) InstanceIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyPtrOutput)
}

type InstanceIamPolicyPtrInput interface {
	pulumi.Input

	ToInstanceIamPolicyPtrOutput() InstanceIamPolicyPtrOutput
	ToInstanceIamPolicyPtrOutputWithContext(ctx context.Context) InstanceIamPolicyPtrOutput
}

type instanceIamPolicyPtrType InstanceIamPolicyArgs

func (*instanceIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamPolicy)(nil))
}

func (i *instanceIamPolicyPtrType) ToInstanceIamPolicyPtrOutput() InstanceIamPolicyPtrOutput {
	return i.ToInstanceIamPolicyPtrOutputWithContext(context.Background())
}

func (i *instanceIamPolicyPtrType) ToInstanceIamPolicyPtrOutputWithContext(ctx context.Context) InstanceIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyPtrOutput)
}

// InstanceIamPolicyArrayInput is an input type that accepts InstanceIamPolicyArray and InstanceIamPolicyArrayOutput values.
// You can construct a concrete instance of `InstanceIamPolicyArrayInput` via:
//
//          InstanceIamPolicyArray{ InstanceIamPolicyArgs{...} }
type InstanceIamPolicyArrayInput interface {
	pulumi.Input

	ToInstanceIamPolicyArrayOutput() InstanceIamPolicyArrayOutput
	ToInstanceIamPolicyArrayOutputWithContext(context.Context) InstanceIamPolicyArrayOutput
}

type InstanceIamPolicyArray []InstanceIamPolicyInput

func (InstanceIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*InstanceIamPolicy)(nil))
}

func (i InstanceIamPolicyArray) ToInstanceIamPolicyArrayOutput() InstanceIamPolicyArrayOutput {
	return i.ToInstanceIamPolicyArrayOutputWithContext(context.Background())
}

func (i InstanceIamPolicyArray) ToInstanceIamPolicyArrayOutputWithContext(ctx context.Context) InstanceIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyArrayOutput)
}

// InstanceIamPolicyMapInput is an input type that accepts InstanceIamPolicyMap and InstanceIamPolicyMapOutput values.
// You can construct a concrete instance of `InstanceIamPolicyMapInput` via:
//
//          InstanceIamPolicyMap{ "key": InstanceIamPolicyArgs{...} }
type InstanceIamPolicyMapInput interface {
	pulumi.Input

	ToInstanceIamPolicyMapOutput() InstanceIamPolicyMapOutput
	ToInstanceIamPolicyMapOutputWithContext(context.Context) InstanceIamPolicyMapOutput
}

type InstanceIamPolicyMap map[string]InstanceIamPolicyInput

func (InstanceIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*InstanceIamPolicy)(nil))
}

func (i InstanceIamPolicyMap) ToInstanceIamPolicyMapOutput() InstanceIamPolicyMapOutput {
	return i.ToInstanceIamPolicyMapOutputWithContext(context.Background())
}

func (i InstanceIamPolicyMap) ToInstanceIamPolicyMapOutputWithContext(ctx context.Context) InstanceIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyMapOutput)
}

type InstanceIamPolicyOutput struct {
	*pulumi.OutputState
}

func (InstanceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamPolicy)(nil))
}

func (o InstanceIamPolicyOutput) ToInstanceIamPolicyOutput() InstanceIamPolicyOutput {
	return o
}

func (o InstanceIamPolicyOutput) ToInstanceIamPolicyOutputWithContext(ctx context.Context) InstanceIamPolicyOutput {
	return o
}

func (o InstanceIamPolicyOutput) ToInstanceIamPolicyPtrOutput() InstanceIamPolicyPtrOutput {
	return o.ToInstanceIamPolicyPtrOutputWithContext(context.Background())
}

func (o InstanceIamPolicyOutput) ToInstanceIamPolicyPtrOutputWithContext(ctx context.Context) InstanceIamPolicyPtrOutput {
	return o.ApplyT(func(v InstanceIamPolicy) *InstanceIamPolicy {
		return &v
	}).(InstanceIamPolicyPtrOutput)
}

type InstanceIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIamPolicy)(nil))
}

func (o InstanceIamPolicyPtrOutput) ToInstanceIamPolicyPtrOutput() InstanceIamPolicyPtrOutput {
	return o
}

func (o InstanceIamPolicyPtrOutput) ToInstanceIamPolicyPtrOutputWithContext(ctx context.Context) InstanceIamPolicyPtrOutput {
	return o
}

type InstanceIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (InstanceIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceIamPolicy)(nil))
}

func (o InstanceIamPolicyArrayOutput) ToInstanceIamPolicyArrayOutput() InstanceIamPolicyArrayOutput {
	return o
}

func (o InstanceIamPolicyArrayOutput) ToInstanceIamPolicyArrayOutputWithContext(ctx context.Context) InstanceIamPolicyArrayOutput {
	return o
}

func (o InstanceIamPolicyArrayOutput) Index(i pulumi.IntInput) InstanceIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceIamPolicy {
		return vs[0].([]InstanceIamPolicy)[vs[1].(int)]
	}).(InstanceIamPolicyOutput)
}

type InstanceIamPolicyMapOutput struct{ *pulumi.OutputState }

func (InstanceIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceIamPolicy)(nil))
}

func (o InstanceIamPolicyMapOutput) ToInstanceIamPolicyMapOutput() InstanceIamPolicyMapOutput {
	return o
}

func (o InstanceIamPolicyMapOutput) ToInstanceIamPolicyMapOutputWithContext(ctx context.Context) InstanceIamPolicyMapOutput {
	return o
}

func (o InstanceIamPolicyMapOutput) MapIndex(k pulumi.StringInput) InstanceIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceIamPolicy {
		return vs[0].(map[string]InstanceIamPolicy)[vs[1].(string)]
	}).(InstanceIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceIamPolicyOutput{})
	pulumi.RegisterOutputType(InstanceIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(InstanceIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(InstanceIamPolicyMapOutput{})
}
