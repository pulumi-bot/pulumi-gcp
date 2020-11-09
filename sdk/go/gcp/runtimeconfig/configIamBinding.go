// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/configs/{{config}} * {{project}}/{{config}} * {{config}} Any variables not passed in the import command will be taken from the provider configuration. Runtime Configurator config IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamBinding:ConfigIamBinding editor "projects/{{project}}/configs/{{config}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamBinding:ConfigIamBinding editor "projects/{{project}}/configs/{{config}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:runtimeconfig/configIamBinding:ConfigIamBinding editor projects/{{project}}/configs/{{config}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ConfigIamBinding struct {
	pulumi.CustomResourceState

	Condition ConfigIamBindingConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringOutput `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewConfigIamBinding(ctx *pulumi.Context,
	name string, args *ConfigIamBindingArgs, opts ...pulumi.ResourceOption) (*ConfigIamBinding, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ConfigIamBindingArgs{}
	}
	var resource ConfigIamBinding
	err := ctx.RegisterResource("gcp:runtimeconfig/configIamBinding:ConfigIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigIamBinding gets an existing ConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigIamBindingState, opts ...pulumi.ResourceOption) (*ConfigIamBinding, error) {
	var resource ConfigIamBinding
	err := ctx.ReadResource("gcp:runtimeconfig/configIamBinding:ConfigIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigIamBinding resources.
type configIamBindingState struct {
	Condition *ConfigIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config *string `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ConfigIamBindingState struct {
	Condition ConfigIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ConfigIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamBindingState)(nil)).Elem()
}

type configIamBindingArgs struct {
	Condition *ConfigIamBindingCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config  string   `pulumi:"config"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ConfigIamBinding resource.
type ConfigIamBindingArgs struct {
	Condition ConfigIamBindingConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Config  pulumi.StringInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ConfigIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamBindingArgs)(nil)).Elem()
}
