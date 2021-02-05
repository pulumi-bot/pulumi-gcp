// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} * {{project}}/{{region}}/{{taxonomy}} * {{region}}/{{taxonomy}} * {{taxonomy}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog taxonomy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type TaxonomyIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	Region  pulumi.StringOutput `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringOutput `pulumi:"taxonomy"`
}

// NewTaxonomyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyIamPolicy(ctx *pulumi.Context,
	name string, args *TaxonomyIamPolicyArgs, opts ...pulumi.ResourceOption) (*TaxonomyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Taxonomy == nil {
		return nil, errors.New("invalid value for required argument 'Taxonomy'")
	}
	var resource TaxonomyIamPolicy
	err := ctx.RegisterResource("gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyIamPolicy gets an existing TaxonomyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyIamPolicyState, opts ...pulumi.ResourceOption) (*TaxonomyIamPolicy, error) {
	var resource TaxonomyIamPolicy
	err := ctx.ReadResource("gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyIamPolicy resources.
type taxonomyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy *string `pulumi:"taxonomy"`
}

type TaxonomyIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringPtrInput
}

func (TaxonomyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamPolicyState)(nil)).Elem()
}

type taxonomyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy string `pulumi:"taxonomy"`
}

// The set of arguments for constructing a TaxonomyIamPolicy resource.
type TaxonomyIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	Region  pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Taxonomy pulumi.StringInput
}

func (TaxonomyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyIamPolicyArgs)(nil)).Elem()
}

type TaxonomyIamPolicyInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput
	ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput
}

func (*TaxonomyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyIamPolicy)(nil))
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput {
	return i.ToTaxonomyIamPolicyOutputWithContext(context.Background())
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyOutput)
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyPtrOutput() TaxonomyIamPolicyPtrOutput {
	return i.ToTaxonomyIamPolicyPtrOutputWithContext(context.Background())
}

func (i *TaxonomyIamPolicy) ToTaxonomyIamPolicyPtrOutputWithContext(ctx context.Context) TaxonomyIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyPtrOutput)
}

type TaxonomyIamPolicyPtrInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyPtrOutput() TaxonomyIamPolicyPtrOutput
	ToTaxonomyIamPolicyPtrOutputWithContext(ctx context.Context) TaxonomyIamPolicyPtrOutput
}

type taxonomyIamPolicyPtrType TaxonomyIamPolicyArgs

func (*taxonomyIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamPolicy)(nil))
}

func (i *taxonomyIamPolicyPtrType) ToTaxonomyIamPolicyPtrOutput() TaxonomyIamPolicyPtrOutput {
	return i.ToTaxonomyIamPolicyPtrOutputWithContext(context.Background())
}

func (i *taxonomyIamPolicyPtrType) ToTaxonomyIamPolicyPtrOutputWithContext(ctx context.Context) TaxonomyIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyPtrOutput)
}

// TaxonomyIamPolicyArrayInput is an input type that accepts TaxonomyIamPolicyArray and TaxonomyIamPolicyArrayOutput values.
// You can construct a concrete instance of `TaxonomyIamPolicyArrayInput` via:
//
//          TaxonomyIamPolicyArray{ TaxonomyIamPolicyArgs{...} }
type TaxonomyIamPolicyArrayInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput
	ToTaxonomyIamPolicyArrayOutputWithContext(context.Context) TaxonomyIamPolicyArrayOutput
}

type TaxonomyIamPolicyArray []TaxonomyIamPolicyInput

func (TaxonomyIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TaxonomyIamPolicy)(nil))
}

func (i TaxonomyIamPolicyArray) ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput {
	return i.ToTaxonomyIamPolicyArrayOutputWithContext(context.Background())
}

func (i TaxonomyIamPolicyArray) ToTaxonomyIamPolicyArrayOutputWithContext(ctx context.Context) TaxonomyIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyArrayOutput)
}

// TaxonomyIamPolicyMapInput is an input type that accepts TaxonomyIamPolicyMap and TaxonomyIamPolicyMapOutput values.
// You can construct a concrete instance of `TaxonomyIamPolicyMapInput` via:
//
//          TaxonomyIamPolicyMap{ "key": TaxonomyIamPolicyArgs{...} }
type TaxonomyIamPolicyMapInput interface {
	pulumi.Input

	ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput
	ToTaxonomyIamPolicyMapOutputWithContext(context.Context) TaxonomyIamPolicyMapOutput
}

type TaxonomyIamPolicyMap map[string]TaxonomyIamPolicyInput

func (TaxonomyIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TaxonomyIamPolicy)(nil))
}

func (i TaxonomyIamPolicyMap) ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput {
	return i.ToTaxonomyIamPolicyMapOutputWithContext(context.Background())
}

func (i TaxonomyIamPolicyMap) ToTaxonomyIamPolicyMapOutputWithContext(ctx context.Context) TaxonomyIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyIamPolicyMapOutput)
}

type TaxonomyIamPolicyOutput struct {
	*pulumi.OutputState
}

func (TaxonomyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyIamPolicy)(nil))
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyOutput() TaxonomyIamPolicyOutput {
	return o
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyOutputWithContext(ctx context.Context) TaxonomyIamPolicyOutput {
	return o
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyPtrOutput() TaxonomyIamPolicyPtrOutput {
	return o.ToTaxonomyIamPolicyPtrOutputWithContext(context.Background())
}

func (o TaxonomyIamPolicyOutput) ToTaxonomyIamPolicyPtrOutputWithContext(ctx context.Context) TaxonomyIamPolicyPtrOutput {
	return o.ApplyT(func(v TaxonomyIamPolicy) *TaxonomyIamPolicy {
		return &v
	}).(TaxonomyIamPolicyPtrOutput)
}

type TaxonomyIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (TaxonomyIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyIamPolicy)(nil))
}

func (o TaxonomyIamPolicyPtrOutput) ToTaxonomyIamPolicyPtrOutput() TaxonomyIamPolicyPtrOutput {
	return o
}

func (o TaxonomyIamPolicyPtrOutput) ToTaxonomyIamPolicyPtrOutputWithContext(ctx context.Context) TaxonomyIamPolicyPtrOutput {
	return o
}

type TaxonomyIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (TaxonomyIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaxonomyIamPolicy)(nil))
}

func (o TaxonomyIamPolicyArrayOutput) ToTaxonomyIamPolicyArrayOutput() TaxonomyIamPolicyArrayOutput {
	return o
}

func (o TaxonomyIamPolicyArrayOutput) ToTaxonomyIamPolicyArrayOutputWithContext(ctx context.Context) TaxonomyIamPolicyArrayOutput {
	return o
}

func (o TaxonomyIamPolicyArrayOutput) Index(i pulumi.IntInput) TaxonomyIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TaxonomyIamPolicy {
		return vs[0].([]TaxonomyIamPolicy)[vs[1].(int)]
	}).(TaxonomyIamPolicyOutput)
}

type TaxonomyIamPolicyMapOutput struct{ *pulumi.OutputState }

func (TaxonomyIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TaxonomyIamPolicy)(nil))
}

func (o TaxonomyIamPolicyMapOutput) ToTaxonomyIamPolicyMapOutput() TaxonomyIamPolicyMapOutput {
	return o
}

func (o TaxonomyIamPolicyMapOutput) ToTaxonomyIamPolicyMapOutputWithContext(ctx context.Context) TaxonomyIamPolicyMapOutput {
	return o
}

func (o TaxonomyIamPolicyMapOutput) MapIndex(k pulumi.StringInput) TaxonomyIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TaxonomyIamPolicy {
		return vs[0].(map[string]TaxonomyIamPolicy)[vs[1].(string)]
	}).(TaxonomyIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(TaxonomyIamPolicyOutput{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(TaxonomyIamPolicyMapOutput{})
}
