// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package configIamMember

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/runtimeconfig/ConfigIamMemberCondition"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_member.html.markdown.
type ConfigIamMember struct {
	pulumi.CustomResourceState

	Condition runtimeconfigConfigIamMemberCondition.ConfigIamMemberConditionPtrOutput `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringOutput `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewConfigIamMember registers a new resource with the given unique name, arguments, and options.
func NewConfigIamMember(ctx *pulumi.Context,
	name string, args *ConfigIamMemberArgs, opts ...pulumi.ResourceOption) (*ConfigIamMember, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &ConfigIamMemberArgs{}
	}
	var resource ConfigIamMember
	err := ctx.RegisterResource("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigIamMember gets an existing ConfigIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigIamMemberState, opts ...pulumi.ResourceOption) (*ConfigIamMember, error) {
	var resource ConfigIamMember
	err := ctx.ReadResource("gcp:runtimeconfig/configIamMember:ConfigIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigIamMember resources.
type configIamMemberState struct {
	Condition *runtimeconfigConfigIamMemberCondition.ConfigIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config *string `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ConfigIamMemberState struct {
	Condition runtimeconfigConfigIamMemberCondition.ConfigIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ConfigIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamMemberState)(nil)).Elem()
}

type configIamMemberArgs struct {
	Condition *runtimeconfigConfigIamMemberCondition.ConfigIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	Config string `pulumi:"config"`
	Member string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ConfigIamMember resource.
type ConfigIamMemberArgs struct {
	Condition runtimeconfigConfigIamMemberCondition.ConfigIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringInput
	Member pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ConfigIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamMemberArgs)(nil)).Elem()
}

