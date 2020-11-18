// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ApiConfigIamPolicy struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringOutput `pulumi:"api"`
	ApiConfig pulumi.StringOutput `pulumi:"apiConfig"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewApiConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ApiConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	if args.ApiConfig == nil {
		return nil, errors.New("invalid value for required argument 'ApiConfig'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource ApiConfigIamPolicy
	err := ctx.RegisterResource("gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamPolicy gets an existing ApiConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamPolicyState, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	var resource ApiConfigIamPolicy
	err := ctx.ReadResource("gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamPolicy resources.
type apiConfigIamPolicyState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       *string `pulumi:"api"`
	ApiConfig *string `pulumi:"apiConfig"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type ApiConfigIamPolicyState struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringPtrInput
	ApiConfig pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyState)(nil)).Elem()
}

type apiConfigIamPolicyArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       string `pulumi:"api"`
	ApiConfig string `pulumi:"apiConfig"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ApiConfigIamPolicy resource.
type ApiConfigIamPolicyArgs struct {
	// The API to attach the config to.
	// Used to find the parent resource to bind the IAM policy to
	Api       pulumi.StringInput
	ApiConfig pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ApiConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyArgs)(nil)).Elem()
}

type ApiConfigIamPolicyInput interface {
	pulumi.Input

	ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput
	ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput
}

func (ApiConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamPolicy)(nil)).Elem()
}

func (i ApiConfigIamPolicy) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return i.ToApiConfigIamPolicyOutputWithContext(context.Background())
}

func (i ApiConfigIamPolicy) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamPolicyOutput)
}

type ApiConfigIamPolicyOutput struct {
	*pulumi.OutputState
}

func (ApiConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiConfigIamPolicyOutput)(nil)).Elem()
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return o
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiConfigIamPolicyOutput{})
}
