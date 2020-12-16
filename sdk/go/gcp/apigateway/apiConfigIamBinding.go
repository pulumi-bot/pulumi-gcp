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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} * {{project}}/{{api}}/{{api_config}} * {{api}}/{{api_config}} * {{api_config}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway apiconfig IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type ApiConfigIamBinding struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringOutput                   `pulumi:"api"`
	ApiConfig pulumi.StringOutput                   `pulumi:"apiConfig"`
	Condition ApiConfigIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewApiConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamBinding(ctx *pulumi.Context,
	name string, args *ApiConfigIamBindingArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.ApiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfig'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource ApiConfigIamBinding
	err := ctx.RegisterResource("gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamBinding gets an existing ApiConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamBindingState, opts ...pulumi.ResourceOption) (*ApiConfigIamBinding, error) {
	var resource ApiConfigIamBinding
	err := ctx.ReadResource("gcp:apigateway/apiConfigIamBinding:ApiConfigIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamBinding resources.
type apiConfigIamBindingState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       *string                       `pulumi:"api"`
	ApiConfig *string                       `pulumi:"apiConfig"`
	Condition *ApiConfigIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type ApiConfigIamBindingState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringPtrInput
	ApiConfig pulumi.StringPtrInput
	Condition ApiConfigIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (ApiConfigIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamBindingState)(nil)).Elem()
}

type apiConfigIamBindingArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       string                        `pulumi:"api"`
	ApiConfig string                        `pulumi:"apiConfig"`
	Condition *ApiConfigIamBindingCondition `pulumi:"condition"`
	Members   []string                      `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a ApiConfigIamBinding resource.
type ApiConfigIamBindingArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringInput
	ApiConfig pulumi.StringInput
	Condition ApiConfigIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (ApiConfigIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamBindingArgs)(nil)).Elem()
}

type ApiConfigIamBindingInput interface {
	pulumi.Input

	ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput
	ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput
}

func (*ApiConfigIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamBinding)(nil))
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput {
	return i.ToApiConfigIamBindingOutputWithContext(context.Background())
}

func (i *ApiConfigIamBinding) ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamBindingOutput)
}

type ApiConfigIamBindingOutput struct {
	*pulumi.OutputState
}

func (ApiConfigIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamBinding)(nil))
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingOutput() ApiConfigIamBindingOutput {
	return o
}

func (o ApiConfigIamBindingOutput) ToApiConfigIamBindingOutputWithContext(ctx context.Context) ApiConfigIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiConfigIamBindingOutput{})
}
