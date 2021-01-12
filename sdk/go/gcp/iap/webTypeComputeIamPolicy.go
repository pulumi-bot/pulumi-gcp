// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypecompute IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy editor projects/{{project}}/iap_web/compute
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeComputeIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebTypeComputeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebTypeComputeIamPolicy(ctx *pulumi.Context,
	name string, args *WebTypeComputeIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebTypeComputeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	var resource WebTypeComputeIamPolicy
	err := ctx.RegisterResource("gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeComputeIamPolicy gets an existing WebTypeComputeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeComputeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeComputeIamPolicyState, opts ...pulumi.ResourceOption) (*WebTypeComputeIamPolicy, error) {
	var resource WebTypeComputeIamPolicy
	err := ctx.ReadResource("gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeComputeIamPolicy resources.
type webTypeComputeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebTypeComputeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeComputeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamPolicyState)(nil)).Elem()
}

type webTypeComputeIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebTypeComputeIamPolicy resource.
type WebTypeComputeIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeComputeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamPolicyArgs)(nil)).Elem()
}

type WebTypeComputeIamPolicyInput interface {
	pulumi.Input

	ToWebTypeComputeIamPolicyOutput() WebTypeComputeIamPolicyOutput
	ToWebTypeComputeIamPolicyOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyOutput
}

func (*WebTypeComputeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeComputeIamPolicy)(nil))
}

func (i *WebTypeComputeIamPolicy) ToWebTypeComputeIamPolicyOutput() WebTypeComputeIamPolicyOutput {
	return i.ToWebTypeComputeIamPolicyOutputWithContext(context.Background())
}

func (i *WebTypeComputeIamPolicy) ToWebTypeComputeIamPolicyOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamPolicyOutput)
}

func (i *WebTypeComputeIamPolicy) ToWebTypeComputeIamPolicyPtrOutput() WebTypeComputeIamPolicyPtrOutput {
	return i.ToWebTypeComputeIamPolicyPtrOutputWithContext(context.Background())
}

func (i *WebTypeComputeIamPolicy) ToWebTypeComputeIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamPolicyPtrOutput)
}

type WebTypeComputeIamPolicyPtrInput interface {
	pulumi.Input

	ToWebTypeComputeIamPolicyPtrOutput() WebTypeComputeIamPolicyPtrOutput
	ToWebTypeComputeIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyPtrOutput
}

type webTypeComputeIamPolicyPtrType WebTypeComputeIamPolicyArgs

func (*webTypeComputeIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamPolicy)(nil))
}

func (i *webTypeComputeIamPolicyPtrType) ToWebTypeComputeIamPolicyPtrOutput() WebTypeComputeIamPolicyPtrOutput {
	return i.ToWebTypeComputeIamPolicyPtrOutputWithContext(context.Background())
}

func (i *webTypeComputeIamPolicyPtrType) ToWebTypeComputeIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamPolicyOutput).ToWebTypeComputeIamPolicyPtrOutput()
}

// WebTypeComputeIamPolicyArrayInput is an input type that accepts WebTypeComputeIamPolicyArray and WebTypeComputeIamPolicyArrayOutput values.
// You can construct a concrete instance of `WebTypeComputeIamPolicyArrayInput` via:
//
//          WebTypeComputeIamPolicyArray{ WebTypeComputeIamPolicyArgs{...} }
type WebTypeComputeIamPolicyArrayInput interface {
	pulumi.Input

	ToWebTypeComputeIamPolicyArrayOutput() WebTypeComputeIamPolicyArrayOutput
	ToWebTypeComputeIamPolicyArrayOutputWithContext(context.Context) WebTypeComputeIamPolicyArrayOutput
}

type WebTypeComputeIamPolicyArray []WebTypeComputeIamPolicyInput

func (WebTypeComputeIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTypeComputeIamPolicy)(nil))
}

func (i WebTypeComputeIamPolicyArray) ToWebTypeComputeIamPolicyArrayOutput() WebTypeComputeIamPolicyArrayOutput {
	return i.ToWebTypeComputeIamPolicyArrayOutputWithContext(context.Background())
}

func (i WebTypeComputeIamPolicyArray) ToWebTypeComputeIamPolicyArrayOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamPolicyArrayOutput)
}

// WebTypeComputeIamPolicyMapInput is an input type that accepts WebTypeComputeIamPolicyMap and WebTypeComputeIamPolicyMapOutput values.
// You can construct a concrete instance of `WebTypeComputeIamPolicyMapInput` via:
//
//          WebTypeComputeIamPolicyMap{ "key": WebTypeComputeIamPolicyArgs{...} }
type WebTypeComputeIamPolicyMapInput interface {
	pulumi.Input

	ToWebTypeComputeIamPolicyMapOutput() WebTypeComputeIamPolicyMapOutput
	ToWebTypeComputeIamPolicyMapOutputWithContext(context.Context) WebTypeComputeIamPolicyMapOutput
}

type WebTypeComputeIamPolicyMap map[string]WebTypeComputeIamPolicyInput

func (WebTypeComputeIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebTypeComputeIamPolicy)(nil))
}

func (i WebTypeComputeIamPolicyMap) ToWebTypeComputeIamPolicyMapOutput() WebTypeComputeIamPolicyMapOutput {
	return i.ToWebTypeComputeIamPolicyMapOutputWithContext(context.Background())
}

func (i WebTypeComputeIamPolicyMap) ToWebTypeComputeIamPolicyMapOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamPolicyMapOutput)
}

type WebTypeComputeIamPolicyOutput struct {
	*pulumi.OutputState
}

func (WebTypeComputeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeComputeIamPolicy)(nil))
}

func (o WebTypeComputeIamPolicyOutput) ToWebTypeComputeIamPolicyOutput() WebTypeComputeIamPolicyOutput {
	return o
}

func (o WebTypeComputeIamPolicyOutput) ToWebTypeComputeIamPolicyOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyOutput {
	return o
}

func (o WebTypeComputeIamPolicyOutput) ToWebTypeComputeIamPolicyPtrOutput() WebTypeComputeIamPolicyPtrOutput {
	return o.ToWebTypeComputeIamPolicyPtrOutputWithContext(context.Background())
}

func (o WebTypeComputeIamPolicyOutput) ToWebTypeComputeIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyPtrOutput {
	return o.ApplyT(func(v WebTypeComputeIamPolicy) *WebTypeComputeIamPolicy {
		return &v
	}).(WebTypeComputeIamPolicyPtrOutput)
}

type WebTypeComputeIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (WebTypeComputeIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTypeComputeIamPolicy)(nil))
}

func (o WebTypeComputeIamPolicyPtrOutput) ToWebTypeComputeIamPolicyPtrOutput() WebTypeComputeIamPolicyPtrOutput {
	return o
}

func (o WebTypeComputeIamPolicyPtrOutput) ToWebTypeComputeIamPolicyPtrOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyPtrOutput {
	return o
}

type WebTypeComputeIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebTypeComputeIamPolicy)(nil))
}

func (o WebTypeComputeIamPolicyArrayOutput) ToWebTypeComputeIamPolicyArrayOutput() WebTypeComputeIamPolicyArrayOutput {
	return o
}

func (o WebTypeComputeIamPolicyArrayOutput) ToWebTypeComputeIamPolicyArrayOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyArrayOutput {
	return o
}

func (o WebTypeComputeIamPolicyArrayOutput) Index(i pulumi.IntInput) WebTypeComputeIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebTypeComputeIamPolicy {
		return vs[0].([]WebTypeComputeIamPolicy)[vs[1].(int)]
	}).(WebTypeComputeIamPolicyOutput)
}

type WebTypeComputeIamPolicyMapOutput struct{ *pulumi.OutputState }

func (WebTypeComputeIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebTypeComputeIamPolicy)(nil))
}

func (o WebTypeComputeIamPolicyMapOutput) ToWebTypeComputeIamPolicyMapOutput() WebTypeComputeIamPolicyMapOutput {
	return o
}

func (o WebTypeComputeIamPolicyMapOutput) ToWebTypeComputeIamPolicyMapOutputWithContext(ctx context.Context) WebTypeComputeIamPolicyMapOutput {
	return o
}

func (o WebTypeComputeIamPolicyMapOutput) MapIndex(k pulumi.StringInput) WebTypeComputeIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebTypeComputeIamPolicy {
		return vs[0].(map[string]WebTypeComputeIamPolicy)[vs[1].(string)]
	}).(WebTypeComputeIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(WebTypeComputeIamPolicyOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(WebTypeComputeIamPolicyMapOutput{})
}
