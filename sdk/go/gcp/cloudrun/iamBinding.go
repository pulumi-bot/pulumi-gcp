// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudrun

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
//
// * `cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
// * `cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
// * `cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
//
// > **Note:** `cloudrun.IamPolicy` **cannot** be used in conjunction with `cloudrun.IamBinding` and `cloudrun.IamMember` or they will fight over what your policy should be.
//
// > **Note:** `cloudrun.IamBinding` resources **can be** used in conjunction with `cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_cloud\_run\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudrun"
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
// 		_, err = cloudrun.NewIamPolicy(ctx, "policy", &cloudrun.IamPolicyArgs{
// 			Location:   pulumi.Any(google_cloud_run_service.Default.Location),
// 			Project:    pulumi.Any(google_cloud_run_service.Default.Project),
// 			Service:    pulumi.Any(google_cloud_run_service.Default.Name),
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
// ## google\_cloud\_run\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewIamBinding(ctx, "binding", &cloudrun.IamBindingArgs{
// 			Location: pulumi.Any(google_cloud_run_service.Default.Location),
// 			Project:  pulumi.Any(google_cloud_run_service.Default.Project),
// 			Service:  pulumi.Any(google_cloud_run_service.Default.Name),
// 			Role:     pulumi.String("roles/viewer"),
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
// ## google\_cloud\_run\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudrun"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudrun.NewIamMember(ctx, "member", &cloudrun.IamMemberArgs{
// 			Location: pulumi.Any(google_cloud_run_service.Default.Location),
// 			Project:  pulumi.Any(google_cloud_run_service.Default.Project),
// 			Service:  pulumi.Any(google_cloud_run_service.Default.Name),
// 			Role:     pulumi.String("roles/viewer"),
// 			Member:   pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/services/{{service}} * {{project}}/{{location}}/{{service}} * {{location}}/{{service}} * {{service}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Run service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudrun/iamBinding:IamBinding editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudrun/iamBinding:IamBinding editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:cloudrun/iamBinding:IamBinding editor projects/{{project}}/locations/{{location}}/services/{{service}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IamBinding struct {
	pulumi.CustomResourceState

	Condition IamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput      `pulumi:"location"`
	Members  pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIamBinding(ctx *pulumi.Context,
	name string, args *IamBindingArgs, opts ...pulumi.ResourceOption) (*IamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IamBinding
	err := ctx.RegisterResource("gcp:cloudrun/iamBinding:IamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamBinding gets an existing IamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamBindingState, opts ...pulumi.ResourceOption) (*IamBinding, error) {
	var resource IamBinding
	err := ctx.ReadResource("gcp:cloudrun/iamBinding:IamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamBinding resources.
type iamBindingState struct {
	Condition *IamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type IamBindingState struct {
	Condition IamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (IamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingState)(nil)).Elem()
}

type iamBindingArgs struct {
	Condition *IamBindingCondition `pulumi:"condition"`
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location *string  `pulumi:"location"`
	Members  []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamBinding resource.
type IamBindingArgs struct {
	Condition IamBindingConditionPtrInput
	// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Members  pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `cloudrun.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (IamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingArgs)(nil)).Elem()
}

type IamBindingInput interface {
	pulumi.Input

	ToIamBindingOutput() IamBindingOutput
	ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput
}

func (*IamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*IamBinding)(nil))
}

func (i *IamBinding) ToIamBindingOutput() IamBindingOutput {
	return i.ToIamBindingOutputWithContext(context.Background())
}

func (i *IamBinding) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingOutput)
}

func (i *IamBinding) ToIamBindingPtrOutput() IamBindingPtrOutput {
	return i.ToIamBindingPtrOutputWithContext(context.Background())
}

func (i *IamBinding) ToIamBindingPtrOutputWithContext(ctx context.Context) IamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingPtrOutput)
}

type IamBindingPtrInput interface {
	pulumi.Input

	ToIamBindingPtrOutput() IamBindingPtrOutput
	ToIamBindingPtrOutputWithContext(ctx context.Context) IamBindingPtrOutput
}

type iamBindingPtrType IamBindingArgs

func (*iamBindingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamBinding)(nil))
}

func (i *iamBindingPtrType) ToIamBindingPtrOutput() IamBindingPtrOutput {
	return i.ToIamBindingPtrOutputWithContext(context.Background())
}

func (i *iamBindingPtrType) ToIamBindingPtrOutputWithContext(ctx context.Context) IamBindingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingPtrOutput)
}

// IamBindingArrayInput is an input type that accepts IamBindingArray and IamBindingArrayOutput values.
// You can construct a concrete instance of `IamBindingArrayInput` via:
//
//          IamBindingArray{ IamBindingArgs{...} }
type IamBindingArrayInput interface {
	pulumi.Input

	ToIamBindingArrayOutput() IamBindingArrayOutput
	ToIamBindingArrayOutputWithContext(context.Context) IamBindingArrayOutput
}

type IamBindingArray []IamBindingInput

func (IamBindingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IamBinding)(nil))
}

func (i IamBindingArray) ToIamBindingArrayOutput() IamBindingArrayOutput {
	return i.ToIamBindingArrayOutputWithContext(context.Background())
}

func (i IamBindingArray) ToIamBindingArrayOutputWithContext(ctx context.Context) IamBindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingArrayOutput)
}

// IamBindingMapInput is an input type that accepts IamBindingMap and IamBindingMapOutput values.
// You can construct a concrete instance of `IamBindingMapInput` via:
//
//          IamBindingMap{ "key": IamBindingArgs{...} }
type IamBindingMapInput interface {
	pulumi.Input

	ToIamBindingMapOutput() IamBindingMapOutput
	ToIamBindingMapOutputWithContext(context.Context) IamBindingMapOutput
}

type IamBindingMap map[string]IamBindingInput

func (IamBindingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IamBinding)(nil))
}

func (i IamBindingMap) ToIamBindingMapOutput() IamBindingMapOutput {
	return i.ToIamBindingMapOutputWithContext(context.Background())
}

func (i IamBindingMap) ToIamBindingMapOutputWithContext(ctx context.Context) IamBindingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingMapOutput)
}

type IamBindingOutput struct {
	*pulumi.OutputState
}

func (IamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamBinding)(nil))
}

func (o IamBindingOutput) ToIamBindingOutput() IamBindingOutput {
	return o
}

func (o IamBindingOutput) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return o
}

func (o IamBindingOutput) ToIamBindingPtrOutput() IamBindingPtrOutput {
	return o.ToIamBindingPtrOutputWithContext(context.Background())
}

func (o IamBindingOutput) ToIamBindingPtrOutputWithContext(ctx context.Context) IamBindingPtrOutput {
	return o.ApplyT(func(v IamBinding) *IamBinding {
		return &v
	}).(IamBindingPtrOutput)
}

type IamBindingPtrOutput struct {
	*pulumi.OutputState
}

func (IamBindingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamBinding)(nil))
}

func (o IamBindingPtrOutput) ToIamBindingPtrOutput() IamBindingPtrOutput {
	return o
}

func (o IamBindingPtrOutput) ToIamBindingPtrOutputWithContext(ctx context.Context) IamBindingPtrOutput {
	return o
}

type IamBindingArrayOutput struct{ *pulumi.OutputState }

func (IamBindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamBinding)(nil))
}

func (o IamBindingArrayOutput) ToIamBindingArrayOutput() IamBindingArrayOutput {
	return o
}

func (o IamBindingArrayOutput) ToIamBindingArrayOutputWithContext(ctx context.Context) IamBindingArrayOutput {
	return o
}

func (o IamBindingArrayOutput) Index(i pulumi.IntInput) IamBindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamBinding {
		return vs[0].([]IamBinding)[vs[1].(int)]
	}).(IamBindingOutput)
}

type IamBindingMapOutput struct{ *pulumi.OutputState }

func (IamBindingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamBinding)(nil))
}

func (o IamBindingMapOutput) ToIamBindingMapOutput() IamBindingMapOutput {
	return o
}

func (o IamBindingMapOutput) ToIamBindingMapOutputWithContext(ctx context.Context) IamBindingMapOutput {
	return o
}

func (o IamBindingMapOutput) MapIndex(k pulumi.StringInput) IamBindingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamBinding {
		return vs[0].(map[string]IamBinding)[vs[1].(string)]
	}).(IamBindingOutput)
}

func init() {
	pulumi.RegisterOutputType(IamBindingOutput{})
	pulumi.RegisterOutputType(IamBindingPtrOutput{})
	pulumi.RegisterOutputType(IamBindingArrayOutput{})
	pulumi.RegisterOutputType(IamBindingMapOutput{})
}
