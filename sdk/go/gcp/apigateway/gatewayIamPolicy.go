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
//  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type GatewayIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput `pulumi:"etag"`
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewGatewayIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewGatewayIamPolicy(ctx *pulumi.Context,
	name string, args *GatewayIamPolicyArgs, opts ...pulumi.ResourceOption) (*GatewayIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource GatewayIamPolicy
	err := ctx.RegisterResource("gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayIamPolicy gets an existing GatewayIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayIamPolicyState, opts ...pulumi.ResourceOption) (*GatewayIamPolicy, error) {
	var resource GatewayIamPolicy
	err := ctx.ReadResource("gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayIamPolicy resources.
type gatewayIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag    *string `pulumi:"etag"`
	Gateway *string `pulumi:"gateway"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

type GatewayIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Gateway pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (GatewayIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamPolicyState)(nil)).Elem()
}

type gatewayIamPolicyArgs struct {
	Gateway string `pulumi:"gateway"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a GatewayIamPolicy resource.
type GatewayIamPolicyArgs struct {
	Gateway pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the gateway for the API.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
}

func (GatewayIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayIamPolicyArgs)(nil)).Elem()
}

type GatewayIamPolicyInput interface {
	pulumi.Input

	ToGatewayIamPolicyOutput() GatewayIamPolicyOutput
	ToGatewayIamPolicyOutputWithContext(ctx context.Context) GatewayIamPolicyOutput
}

func (*GatewayIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamPolicy)(nil))
}

func (i *GatewayIamPolicy) ToGatewayIamPolicyOutput() GatewayIamPolicyOutput {
	return i.ToGatewayIamPolicyOutputWithContext(context.Background())
}

func (i *GatewayIamPolicy) ToGatewayIamPolicyOutputWithContext(ctx context.Context) GatewayIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamPolicyOutput)
}

func (i *GatewayIamPolicy) ToGatewayIamPolicyPtrOutput() GatewayIamPolicyPtrOutput {
	return i.ToGatewayIamPolicyPtrOutputWithContext(context.Background())
}

func (i *GatewayIamPolicy) ToGatewayIamPolicyPtrOutputWithContext(ctx context.Context) GatewayIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamPolicyPtrOutput)
}

type GatewayIamPolicyPtrInput interface {
	pulumi.Input

	ToGatewayIamPolicyPtrOutput() GatewayIamPolicyPtrOutput
	ToGatewayIamPolicyPtrOutputWithContext(ctx context.Context) GatewayIamPolicyPtrOutput
}

type gatewayIamPolicyPtrType GatewayIamPolicyArgs

func (*gatewayIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamPolicy)(nil))
}

func (i *gatewayIamPolicyPtrType) ToGatewayIamPolicyPtrOutput() GatewayIamPolicyPtrOutput {
	return i.ToGatewayIamPolicyPtrOutputWithContext(context.Background())
}

func (i *gatewayIamPolicyPtrType) ToGatewayIamPolicyPtrOutputWithContext(ctx context.Context) GatewayIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamPolicyOutput).ToGatewayIamPolicyPtrOutput()
}

// GatewayIamPolicyArrayInput is an input type that accepts GatewayIamPolicyArray and GatewayIamPolicyArrayOutput values.
// You can construct a concrete instance of `GatewayIamPolicyArrayInput` via:
//
//          GatewayIamPolicyArray{ GatewayIamPolicyArgs{...} }
type GatewayIamPolicyArrayInput interface {
	pulumi.Input

	ToGatewayIamPolicyArrayOutput() GatewayIamPolicyArrayOutput
	ToGatewayIamPolicyArrayOutputWithContext(context.Context) GatewayIamPolicyArrayOutput
}

type GatewayIamPolicyArray []GatewayIamPolicyInput

func (GatewayIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayIamPolicy)(nil))
}

func (i GatewayIamPolicyArray) ToGatewayIamPolicyArrayOutput() GatewayIamPolicyArrayOutput {
	return i.ToGatewayIamPolicyArrayOutputWithContext(context.Background())
}

func (i GatewayIamPolicyArray) ToGatewayIamPolicyArrayOutputWithContext(ctx context.Context) GatewayIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamPolicyArrayOutput)
}

// GatewayIamPolicyMapInput is an input type that accepts GatewayIamPolicyMap and GatewayIamPolicyMapOutput values.
// You can construct a concrete instance of `GatewayIamPolicyMapInput` via:
//
//          GatewayIamPolicyMap{ "key": GatewayIamPolicyArgs{...} }
type GatewayIamPolicyMapInput interface {
	pulumi.Input

	ToGatewayIamPolicyMapOutput() GatewayIamPolicyMapOutput
	ToGatewayIamPolicyMapOutputWithContext(context.Context) GatewayIamPolicyMapOutput
}

type GatewayIamPolicyMap map[string]GatewayIamPolicyInput

func (GatewayIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GatewayIamPolicy)(nil))
}

func (i GatewayIamPolicyMap) ToGatewayIamPolicyMapOutput() GatewayIamPolicyMapOutput {
	return i.ToGatewayIamPolicyMapOutputWithContext(context.Background())
}

func (i GatewayIamPolicyMap) ToGatewayIamPolicyMapOutputWithContext(ctx context.Context) GatewayIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayIamPolicyMapOutput)
}

type GatewayIamPolicyOutput struct {
	*pulumi.OutputState
}

func (GatewayIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayIamPolicy)(nil))
}

func (o GatewayIamPolicyOutput) ToGatewayIamPolicyOutput() GatewayIamPolicyOutput {
	return o
}

func (o GatewayIamPolicyOutput) ToGatewayIamPolicyOutputWithContext(ctx context.Context) GatewayIamPolicyOutput {
	return o
}

func (o GatewayIamPolicyOutput) ToGatewayIamPolicyPtrOutput() GatewayIamPolicyPtrOutput {
	return o.ToGatewayIamPolicyPtrOutputWithContext(context.Background())
}

func (o GatewayIamPolicyOutput) ToGatewayIamPolicyPtrOutputWithContext(ctx context.Context) GatewayIamPolicyPtrOutput {
	return o.ApplyT(func(v GatewayIamPolicy) *GatewayIamPolicy {
		return &v
	}).(GatewayIamPolicyPtrOutput)
}

type GatewayIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (GatewayIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayIamPolicy)(nil))
}

func (o GatewayIamPolicyPtrOutput) ToGatewayIamPolicyPtrOutput() GatewayIamPolicyPtrOutput {
	return o
}

func (o GatewayIamPolicyPtrOutput) ToGatewayIamPolicyPtrOutputWithContext(ctx context.Context) GatewayIamPolicyPtrOutput {
	return o
}

type GatewayIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (GatewayIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayIamPolicy)(nil))
}

func (o GatewayIamPolicyArrayOutput) ToGatewayIamPolicyArrayOutput() GatewayIamPolicyArrayOutput {
	return o
}

func (o GatewayIamPolicyArrayOutput) ToGatewayIamPolicyArrayOutputWithContext(ctx context.Context) GatewayIamPolicyArrayOutput {
	return o
}

func (o GatewayIamPolicyArrayOutput) Index(i pulumi.IntInput) GatewayIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayIamPolicy {
		return vs[0].([]GatewayIamPolicy)[vs[1].(int)]
	}).(GatewayIamPolicyOutput)
}

type GatewayIamPolicyMapOutput struct{ *pulumi.OutputState }

func (GatewayIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GatewayIamPolicy)(nil))
}

func (o GatewayIamPolicyMapOutput) ToGatewayIamPolicyMapOutput() GatewayIamPolicyMapOutput {
	return o
}

func (o GatewayIamPolicyMapOutput) ToGatewayIamPolicyMapOutputWithContext(ctx context.Context) GatewayIamPolicyMapOutput {
	return o
}

func (o GatewayIamPolicyMapOutput) MapIndex(k pulumi.StringInput) GatewayIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GatewayIamPolicy {
		return vs[0].(map[string]GatewayIamPolicy)[vs[1].(string)]
	}).(GatewayIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayIamPolicyOutput{})
	pulumi.RegisterOutputType(GatewayIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(GatewayIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(GatewayIamPolicyMapOutput{})
}
