// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud AI Notebooks Instance. Each of these resources serves a different use case:
//
// * `notebooks.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `notebooks.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `notebooks.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `notebooks.InstanceIamPolicy` **cannot** be used in conjunction with `notebooks.InstanceIamBinding` and `notebooks.InstanceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `notebooks.InstanceIamBinding` resources **can be** used in conjunction with `notebooks.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_notebooks\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = notebooks.NewInstanceIamPolicy(ctx, "policy", &notebooks.InstanceIamPolicyArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_notebooks\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamBinding(ctx, "binding", &notebooks.InstanceIamBindingArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
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
// ## google\_notebooks\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamMember(ctx, "member", &notebooks.InstanceIamMemberArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
// 			Member:       pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/instances/{{instance_name}} * {{project}}/{{location}}/{{instance_name}} * {{location}}/{{instance_name}} * {{instance_name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud AI Notebooks instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamPolicy:InstanceIamPolicy editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamPolicy:InstanceIamPolicy editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamPolicy:InstanceIamPolicy editor projects/{{project}}/locations/{{location}}/instances/{{instance_name}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewInstanceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamPolicy(ctx *pulumi.Context,
	name string, args *InstanceIamPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource InstanceIamPolicy
	err := ctx.RegisterResource("gcp:notebooks/instanceIamPolicy:InstanceIamPolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:notebooks/instanceIamPolicy:InstanceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamPolicy resources.
type instanceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type InstanceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (InstanceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamPolicyState)(nil)).Elem()
}

type instanceIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a InstanceIamPolicy resource.
type InstanceIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
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
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamPolicyOutput).ToInstanceIamPolicyPtrOutput()
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
	return reflect.TypeOf((*[]InstanceIamPolicy)(nil))
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
	return reflect.TypeOf((*map[string]InstanceIamPolicy)(nil))
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
