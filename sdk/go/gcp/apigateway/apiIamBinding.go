// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ApiIamBinding struct {
	pulumi.CustomResourceState

	Api       pulumi.StringOutput             `pulumi:"api"`
	Condition ApiIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiIamBinding registers a new resource with the given unique name, arguments, and options.
func NewApiIamBinding(ctx *pulumi.Context,
	name string, args *ApiIamBindingArgs, opts ...pulumi.ResourceOption) (*ApiIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ApiIamBinding
	err := ctx.RegisterResource("gcp:apigateway/apiIamBinding:ApiIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIamBinding gets an existing ApiIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIamBindingState, opts ...pulumi.ResourceOption) (*ApiIamBinding, error) {
	var resource ApiIamBinding
	err := ctx.ReadResource("gcp:apigateway/apiIamBinding:ApiIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIamBinding resources.
type apiIamBindingState struct {
	Api       *string                 `pulumi:"api"`
	Condition *ApiIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiIamBindingState struct {
	Api       pulumi.StringPtrInput
	Condition ApiIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamBindingState)(nil)).Elem()
}

type apiIamBindingArgs struct {
	Api       string                  `pulumi:"api"`
	Condition *ApiIamBindingCondition `pulumi:"condition"`
	Members   []string                `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiIamBinding resource.
type ApiIamBindingArgs struct {
	Api       pulumi.StringInput
	Condition ApiIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIamBindingArgs)(nil)).Elem()
}

type ApiIamBindingInput interface {
	pulumi.Input

	ToApiIamBindingOutput() ApiIamBindingOutput
	ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput
}

func (ApiIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIamBinding)(nil)).Elem()
}

func (i ApiIamBinding) ToApiIamBindingOutput() ApiIamBindingOutput {
	return i.ToApiIamBindingOutputWithContext(context.Background())
}

func (i ApiIamBinding) ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIamBindingOutput)
}

type ApiIamBindingOutput struct {
	*pulumi.OutputState
}

func (ApiIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiIamBindingOutput)(nil)).Elem()
}

func (o ApiIamBindingOutput) ToApiIamBindingOutput() ApiIamBindingOutput {
	return o
}

func (o ApiIamBindingOutput) ToApiIamBindingOutputWithContext(ctx context.Context) ApiIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiIamBindingOutput{})
}
