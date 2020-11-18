// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GatewayIamMember struct {
	pulumi.CustomResourceState

	Condition GatewayIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput `pulumi:"etag"`
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	Member  pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewGatewayIamMember registers a new resource with the given unique name, arguments, and options.
func NewGatewayIamMember(ctx *pulumi.Context,
	name string, args *GatewayIamMemberArgs, opts ...pulumi.ResourceOption) (*GatewayIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource GatewayIamMember
	err := ctx.RegisterResource("gcp:apigateway/gatewayIamMember:GatewayIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayIamMember gets an existing GatewayIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayIamMemberState, opts ...pulumi.ResourceOption) (*GatewayIamMember, error) {
	var resource GatewayIamMember
	err := ctx.ReadResource("gcp:apigateway/gatewayIamMember:GatewayIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayIamMember resources.
type gatewayIamMemberState struct {
	Condition *GatewayIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string `pulumi:"etag"`
	Gateway *string `pulumi:"gateway"`
	Member  *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type GatewayIamMemberState struct {
	Condition GatewayIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Gateway pulumi.StringPtrInput
	Member  pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (GatewayIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamMemberState)(nil)).Elem()
}

type gatewayIamMemberArgs struct {
	Condition *GatewayIamMemberCondition `pulumi:"condition"`
	Gateway   string                     `pulumi:"gateway"`
	Member    string                     `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a GatewayIamMember resource.
type GatewayIamMemberArgs struct {
	Condition GatewayIamMemberConditionPtrInput
	Gateway   pulumi.StringInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.GatewayIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (GatewayIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamMemberArgs)(nil)).Elem()
}

type GatewayIamMemberInput interface {
	pulumi.Input

	ToGatewayIamMemberOutput() GatewayIamMemberOutput
	ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput
}

func (GatewayIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamMember)(nil)).Elem()
}

func (i GatewayIamMember) ToGatewayIamMemberOutput() GatewayIamMemberOutput {
	return i.ToGatewayIamMemberOutputWithContext(context.Background())
}

func (i GatewayIamMember) ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamMemberOutput)
}

type GatewayIamMemberOutput struct {
	*pulumi.OutputState
}

func (GatewayIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamMemberOutput)(nil)).Elem()
}

func (o GatewayIamMemberOutput) ToGatewayIamMemberOutput() GatewayIamMemberOutput {
	return o
}

func (o GatewayIamMemberOutput) ToGatewayIamMemberOutputWithContext(ctx context.Context) GatewayIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayIamMemberOutput{})
}
