// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/gateways/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway gateway IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/gatewayIamBinding:GatewayIamBinding editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type GatewayIamBinding struct {
	pulumi.CustomResourceState

	Condition GatewayIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Gateway pulumi.StringOutput      `pulumi:"gateway"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
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

// NewGatewayIamBinding registers a new resource with the given unique name, arguments, and options.
func NewGatewayIamBinding(ctx *pulumi.Context,
	name string, args *GatewayIamBindingArgs, opts ...pulumi.ResourceOption) (*GatewayIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource GatewayIamBinding
	err := ctx.RegisterResource("gcp:apigateway/gatewayIamBinding:GatewayIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayIamBinding gets an existing GatewayIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayIamBindingState, opts ...pulumi.ResourceOption) (*GatewayIamBinding, error) {
	var resource GatewayIamBinding
	err := ctx.ReadResource("gcp:apigateway/gatewayIamBinding:GatewayIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayIamBinding resources.
type gatewayIamBindingState struct {
	Condition *GatewayIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Gateway *string  `pulumi:"gateway"`
	Members []string `pulumi:"members"`
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

type GatewayIamBindingState struct {
	Condition GatewayIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Gateway pulumi.StringPtrInput
	Members pulumi.StringArrayInput
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

func (GatewayIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamBindingState)(nil)).Elem()
}

type gatewayIamBindingArgs struct {
	Condition *GatewayIamBindingCondition `pulumi:"condition"`
	Gateway   string                      `pulumi:"gateway"`
	Members   []string                    `pulumi:"members"`
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

// The set of arguments for constructing a GatewayIamBinding resource.
type GatewayIamBindingArgs struct {
	Condition GatewayIamBindingConditionPtrInput
	Gateway   pulumi.StringInput
	Members   pulumi.StringArrayInput
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

func (GatewayIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamBindingArgs)(nil)).Elem()
}

type GatewayIamBindingInput interface {
	pulumi.Input

	ToGatewayIamBindingOutput() GatewayIamBindingOutput
	ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput
}

func (*GatewayIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamBinding)(nil))
}

func (i *GatewayIamBinding) ToGatewayIamBindingOutput() GatewayIamBindingOutput {
	return i.ToGatewayIamBindingOutputWithContext(context.Background())
}

func (i *GatewayIamBinding) ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamBindingOutput)
}

type GatewayIamBindingOutput struct {
	*pulumi.OutputState
}

func (GatewayIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamBinding)(nil))
}

func (o GatewayIamBindingOutput) ToGatewayIamBindingOutput() GatewayIamBindingOutput {
	return o
}

func (o GatewayIamBindingOutput) ToGatewayIamBindingOutputWithContext(ctx context.Context) GatewayIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayIamBindingOutput{})
}
