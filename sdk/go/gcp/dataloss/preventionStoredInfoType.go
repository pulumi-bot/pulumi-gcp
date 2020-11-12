// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataloss

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows creation of custom info types.
//
// To get more information about StoredInfoType, see:
//
// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.storedInfoTypes)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-stored-infotypes)
//
// ## Example Usage
//
// ## Import
//
// StoredInfoType can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/storedInfoTypes/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType default {{parent}}/{{name}}
// ```
type PreventionStoredInfoType struct {
	pulumi.CustomResourceState

	// A description of the info type.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrOutput `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrOutput `pulumi:"largeCustomDictionary"`
	// Name describing the field.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrOutput `pulumi:"regex"`
}

// NewPreventionStoredInfoType registers a new resource with the given unique name, arguments, and options.
func NewPreventionStoredInfoType(ctx *pulumi.Context,
	name string, args *PreventionStoredInfoTypeArgs, opts ...pulumi.ResourceOption) (*PreventionStoredInfoType, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil {
		args = &PreventionStoredInfoTypeArgs{}
	}
	var resource PreventionStoredInfoType
	err := ctx.RegisterResource("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreventionStoredInfoType gets an existing PreventionStoredInfoType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreventionStoredInfoType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PreventionStoredInfoTypeState, opts ...pulumi.ResourceOption) (*PreventionStoredInfoType, error) {
	var resource PreventionStoredInfoType
	err := ctx.ReadResource("gcp:dataloss/preventionStoredInfoType:PreventionStoredInfoType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PreventionStoredInfoType resources.
type preventionStoredInfoTypeState struct {
	// A description of the info type.
	Description *string `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary *PreventionStoredInfoTypeDictionary `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName *string `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary *PreventionStoredInfoTypeLargeCustomDictionary `pulumi:"largeCustomDictionary"`
	// Name describing the field.
	Name *string `pulumi:"name"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent *string `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex *PreventionStoredInfoTypeRegex `pulumi:"regex"`
}

type PreventionStoredInfoTypeState struct {
	// A description of the info type.
	Description pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrInput
	// User set display name of the info type.
	DisplayName pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrInput
	// Name describing the field.
	Name pulumi.StringPtrInput
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringPtrInput
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrInput
}

func (PreventionStoredInfoTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionStoredInfoTypeState)(nil)).Elem()
}

type preventionStoredInfoTypeArgs struct {
	// A description of the info type.
	Description *string `pulumi:"description"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary *PreventionStoredInfoTypeDictionary `pulumi:"dictionary"`
	// User set display name of the info type.
	DisplayName *string `pulumi:"displayName"`
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary *PreventionStoredInfoTypeLargeCustomDictionary `pulumi:"largeCustomDictionary"`
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent string `pulumi:"parent"`
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex *PreventionStoredInfoTypeRegex `pulumi:"regex"`
}

// The set of arguments for constructing a PreventionStoredInfoType resource.
type PreventionStoredInfoTypeArgs struct {
	// A description of the info type.
	Description pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary PreventionStoredInfoTypeDictionaryPtrInput
	// User set display name of the info type.
	DisplayName pulumi.StringPtrInput
	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary PreventionStoredInfoTypeLargeCustomDictionaryPtrInput
	// The parent of the info type in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringInput
	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex PreventionStoredInfoTypeRegexPtrInput
}

func (PreventionStoredInfoTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionStoredInfoTypeArgs)(nil)).Elem()
}

type PreventionStoredInfoTypeInput interface {
	pulumi.Input

	ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput
	ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput
}

func (PreventionStoredInfoType) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionStoredInfoType)(nil)).Elem()
}

func (i PreventionStoredInfoType) ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput {
	return i.ToPreventionStoredInfoTypeOutputWithContext(context.Background())
}

func (i PreventionStoredInfoType) ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionStoredInfoTypeOutput)
}

type PreventionStoredInfoTypeOutput struct {
	*pulumi.OutputState
}

func (PreventionStoredInfoTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionStoredInfoTypeOutput)(nil)).Elem()
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypeOutput() PreventionStoredInfoTypeOutput {
	return o
}

func (o PreventionStoredInfoTypeOutput) ToPreventionStoredInfoTypeOutputWithContext(ctx context.Context) PreventionStoredInfoTypeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PreventionStoredInfoTypeOutput{})
}
