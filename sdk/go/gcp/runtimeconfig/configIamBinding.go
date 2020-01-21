// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package runtimeconfig

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_binding.html.markdown.
type ConfigIamBinding struct {
	pulumi.CustomResourceState

	Condition ConfigIamBindingConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringOutput `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
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
	Etag *string `pulumi:"etag"`
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
	Etag pulumi.StringPtrInput
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
	Config string `pulumi:"config"`
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
	Config pulumi.StringInput
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

