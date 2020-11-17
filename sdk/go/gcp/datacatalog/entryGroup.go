// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datacatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.
//
// To get more information about EntryGroup, see:
//
// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
//
// ## Example Usage
type EntryGroup struct {
	pulumi.CustomResourceState

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupId pulumi.StringOutput `pulumi:"entryGroupId"`
	// The resource name of the entry group in URL format. Example:
	// projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// EntryGroup location region.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewEntryGroup registers a new resource with the given unique name, arguments, and options.
func NewEntryGroup(ctx *pulumi.Context,
	name string, args *EntryGroupArgs, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	if args == nil || args.EntryGroupId == nil {
		return nil, errors.New("missing required argument 'EntryGroupId'")
	}
	if args == nil {
		args = &EntryGroupArgs{}
	}
	var resource EntryGroup
	err := ctx.RegisterResource("gcp:datacatalog/entryGroup:EntryGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroup gets an existing EntryGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupState, opts ...pulumi.ResourceOption) (*EntryGroup, error) {
	var resource EntryGroup
	err := ctx.ReadResource("gcp:datacatalog/entryGroup:EntryGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroup resources.
type entryGroupState struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupId *string `pulumi:"entryGroupId"`
	// The resource name of the entry group in URL format. Example:
	// projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// EntryGroup location region.
	Region *string `pulumi:"region"`
}

type EntryGroupState struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrInput
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupId pulumi.StringPtrInput
	// The resource name of the entry group in URL format. Example:
	// projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// EntryGroup location region.
	Region pulumi.StringPtrInput
}

func (EntryGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupState)(nil)).Elem()
}

type entryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description *string `pulumi:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName *string `pulumi:"displayName"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupId string `pulumi:"entryGroupId"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// EntryGroup location region.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a EntryGroup resource.
type EntryGroupArgs struct {
	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description pulumi.StringPtrInput
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName pulumi.StringPtrInput
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupId pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// EntryGroup location region.
	Region pulumi.StringPtrInput
}

func (EntryGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupArgs)(nil)).Elem()
}

type EntryGroupInput interface {
	pulumi.Input

	ToEntryGroupOutput() EntryGroupOutput
	ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput
}

func (EntryGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroup)(nil)).Elem()
}

func (i EntryGroup) ToEntryGroupOutput() EntryGroupOutput {
	return i.ToEntryGroupOutputWithContext(context.Background())
}

func (i EntryGroup) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupOutput)
}

type EntryGroupOutput struct {
	*pulumi.OutputState
}

func (EntryGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryGroupOutput)(nil)).Elem()
}

func (o EntryGroupOutput) ToEntryGroupOutput() EntryGroupOutput {
	return o
}

func (o EntryGroupOutput) ToEntryGroupOutputWithContext(ctx context.Context) EntryGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EntryGroupOutput{})
}
