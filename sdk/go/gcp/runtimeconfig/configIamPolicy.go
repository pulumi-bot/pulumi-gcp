// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Runtime Configurator Config. Each of these resources serves a different use case:
//
// * `runtimeconfig.ConfigIamPolicy`: Authoritative. Sets the IAM policy for the config and replaces any existing policy already attached.
// * `runtimeconfig.ConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the config are preserved.
// * `runtimeconfig.ConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the config are preserved.
//
// > **Note:** `runtimeconfig.ConfigIamPolicy` **cannot** be used in conjunction with `runtimeconfig.ConfigIamBinding` and `runtimeconfig.ConfigIamMember` or they will fight over what your policy should be.
//
// > **Note:** `runtimeconfig.ConfigIamBinding` resources **can be** used in conjunction with `runtimeconfig.ConfigIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_runtimeconfig\_config\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/runtimeconfig"
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
// 		_, err = runtimeconfig.NewConfigIamPolicy(ctx, "policy", &runtimeconfig.ConfigIamPolicyArgs{
// 			Project:    pulumi.Any(google_runtimeconfig_config.Config.Project),
// 			Config:     pulumi.Any(google_runtimeconfig_config.Config.Name),
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
// ## google\_runtimeconfig\_config\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := runtimeconfig.NewConfigIamBinding(ctx, "binding", &runtimeconfig.ConfigIamBindingArgs{
// 			Project: pulumi.Any(google_runtimeconfig_config.Config.Project),
// 			Config:  pulumi.Any(google_runtimeconfig_config.Config.Name),
// 			Role:    pulumi.String("roles/viewer"),
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
// ## google\_runtimeconfig\_config\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := runtimeconfig.NewConfigIamMember(ctx, "member", &runtimeconfig.ConfigIamMemberArgs{
// 			Project: pulumi.Any(google_runtimeconfig_config.Config.Project),
// 			Config:  pulumi.Any(google_runtimeconfig_config.Config.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Member:  pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/configs/{{config}} * {{project}}/{{config}} * {{config}} Any variables not passed in the import command will be taken from the provider configuration. Runtime Configurator config IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy editor "projects/{{project}}/configs/{{config}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy editor "projects/{{project}}/configs/{{config}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy editor projects/{{project}}/configs/{{config}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConfigIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringOutput `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*ConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource ConfigIamPolicy
	err := ctx.RegisterResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigIamPolicy gets an existing ConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigIamPolicyState, opts ...pulumi.ResourceOption) (*ConfigIamPolicy, error) {
	var resource ConfigIamPolicy
	err := ctx.ReadResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigIamPolicy resources.
type configIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config *string `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConfigIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamPolicyState)(nil)).Elem()
}

type configIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config string `pulumi:"config"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ConfigIamPolicy resource.
type ConfigIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamPolicyArgs)(nil)).Elem()
}

type ConfigIamPolicyInput interface {
	pulumi.Input

	ToConfigIamPolicyOutput() ConfigIamPolicyOutput
	ToConfigIamPolicyOutputWithContext(ctx context.Context) ConfigIamPolicyOutput
}

func (*ConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamPolicy)(nil))
}

func (i *ConfigIamPolicy) ToConfigIamPolicyOutput() ConfigIamPolicyOutput {
	return i.ToConfigIamPolicyOutputWithContext(context.Background())
}

func (i *ConfigIamPolicy) ToConfigIamPolicyOutputWithContext(ctx context.Context) ConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamPolicyOutput)
}

func (i *ConfigIamPolicy) ToConfigIamPolicyPtrOutput() ConfigIamPolicyPtrOutput {
	return i.ToConfigIamPolicyPtrOutputWithContext(context.Background())
}

func (i *ConfigIamPolicy) ToConfigIamPolicyPtrOutputWithContext(ctx context.Context) ConfigIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamPolicyPtrOutput)
}

type ConfigIamPolicyPtrInput interface {
	pulumi.Input

	ToConfigIamPolicyPtrOutput() ConfigIamPolicyPtrOutput
	ToConfigIamPolicyPtrOutputWithContext(ctx context.Context) ConfigIamPolicyPtrOutput
}

type configIamPolicyPtrType ConfigIamPolicyArgs

func (*configIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamPolicy)(nil))
}

func (i *configIamPolicyPtrType) ToConfigIamPolicyPtrOutput() ConfigIamPolicyPtrOutput {
	return i.ToConfigIamPolicyPtrOutputWithContext(context.Background())
}

func (i *configIamPolicyPtrType) ToConfigIamPolicyPtrOutputWithContext(ctx context.Context) ConfigIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamPolicyOutput).ToConfigIamPolicyPtrOutput()
}

// ConfigIamPolicyArrayInput is an input type that accepts ConfigIamPolicyArray and ConfigIamPolicyArrayOutput values.
// You can construct a concrete instance of `ConfigIamPolicyArrayInput` via:
//
//          ConfigIamPolicyArray{ ConfigIamPolicyArgs{...} }
type ConfigIamPolicyArrayInput interface {
	pulumi.Input

	ToConfigIamPolicyArrayOutput() ConfigIamPolicyArrayOutput
	ToConfigIamPolicyArrayOutputWithContext(context.Context) ConfigIamPolicyArrayOutput
}

type ConfigIamPolicyArray []ConfigIamPolicyInput

func (ConfigIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigIamPolicy)(nil))
}

func (i ConfigIamPolicyArray) ToConfigIamPolicyArrayOutput() ConfigIamPolicyArrayOutput {
	return i.ToConfigIamPolicyArrayOutputWithContext(context.Background())
}

func (i ConfigIamPolicyArray) ToConfigIamPolicyArrayOutputWithContext(ctx context.Context) ConfigIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamPolicyArrayOutput)
}

// ConfigIamPolicyMapInput is an input type that accepts ConfigIamPolicyMap and ConfigIamPolicyMapOutput values.
// You can construct a concrete instance of `ConfigIamPolicyMapInput` via:
//
//          ConfigIamPolicyMap{ "key": ConfigIamPolicyArgs{...} }
type ConfigIamPolicyMapInput interface {
	pulumi.Input

	ToConfigIamPolicyMapOutput() ConfigIamPolicyMapOutput
	ToConfigIamPolicyMapOutputWithContext(context.Context) ConfigIamPolicyMapOutput
}

type ConfigIamPolicyMap map[string]ConfigIamPolicyInput

func (ConfigIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConfigIamPolicy)(nil))
}

func (i ConfigIamPolicyMap) ToConfigIamPolicyMapOutput() ConfigIamPolicyMapOutput {
	return i.ToConfigIamPolicyMapOutputWithContext(context.Background())
}

func (i ConfigIamPolicyMap) ToConfigIamPolicyMapOutputWithContext(ctx context.Context) ConfigIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigIamPolicyMapOutput)
}

type ConfigIamPolicyOutput struct {
	*pulumi.OutputState
}

func (ConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigIamPolicy)(nil))
}

func (o ConfigIamPolicyOutput) ToConfigIamPolicyOutput() ConfigIamPolicyOutput {
	return o
}

func (o ConfigIamPolicyOutput) ToConfigIamPolicyOutputWithContext(ctx context.Context) ConfigIamPolicyOutput {
	return o
}

func (o ConfigIamPolicyOutput) ToConfigIamPolicyPtrOutput() ConfigIamPolicyPtrOutput {
	return o.ToConfigIamPolicyPtrOutputWithContext(context.Background())
}

func (o ConfigIamPolicyOutput) ToConfigIamPolicyPtrOutputWithContext(ctx context.Context) ConfigIamPolicyPtrOutput {
	return o.ApplyT(func(v ConfigIamPolicy) *ConfigIamPolicy {
		return &v
	}).(ConfigIamPolicyPtrOutput)
}

type ConfigIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (ConfigIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigIamPolicy)(nil))
}

func (o ConfigIamPolicyPtrOutput) ToConfigIamPolicyPtrOutput() ConfigIamPolicyPtrOutput {
	return o
}

func (o ConfigIamPolicyPtrOutput) ToConfigIamPolicyPtrOutputWithContext(ctx context.Context) ConfigIamPolicyPtrOutput {
	return o
}

type ConfigIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (ConfigIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigIamPolicy)(nil))
}

func (o ConfigIamPolicyArrayOutput) ToConfigIamPolicyArrayOutput() ConfigIamPolicyArrayOutput {
	return o
}

func (o ConfigIamPolicyArrayOutput) ToConfigIamPolicyArrayOutputWithContext(ctx context.Context) ConfigIamPolicyArrayOutput {
	return o
}

func (o ConfigIamPolicyArrayOutput) Index(i pulumi.IntInput) ConfigIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigIamPolicy {
		return vs[0].([]ConfigIamPolicy)[vs[1].(int)]
	}).(ConfigIamPolicyOutput)
}

type ConfigIamPolicyMapOutput struct{ *pulumi.OutputState }

func (ConfigIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConfigIamPolicy)(nil))
}

func (o ConfigIamPolicyMapOutput) ToConfigIamPolicyMapOutput() ConfigIamPolicyMapOutput {
	return o
}

func (o ConfigIamPolicyMapOutput) ToConfigIamPolicyMapOutputWithContext(ctx context.Context) ConfigIamPolicyMapOutput {
	return o
}

func (o ConfigIamPolicyMapOutput) MapIndex(k pulumi.StringInput) ConfigIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConfigIamPolicy {
		return vs[0].(map[string]ConfigIamPolicy)[vs[1].(string)]
	}).(ConfigIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigIamPolicyOutput{})
	pulumi.RegisterOutputType(ConfigIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConfigIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(ConfigIamPolicyMapOutput{})
}
