// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// OAuth brand data. Only "Organization Internal" brands can be created
// programatically via API. To convert it into an external brands
// please use the GCP Console.
//
// > **Note:** Brands can only be created once for a Google Cloud
// project and the underlying Google API doesn't not support DELETE or PATCH methods.
// Destroying a provider-managed Brand will remove it from state
// but *will not delete it from Google Cloud.*
//
// To get more information about Brand, see:
//
// * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands)
// * How-to Guides
//     * [Setting up IAP Brand](https://cloud.google.com/iap/docs/tutorial-gce#set_up_iap)
//
// ## Example Usage
type Brand struct {
	pulumi.CustomResourceState

	// Application name displayed on OAuth consent screen.
	ApplicationTitle pulumi.StringOutput `pulumi:"applicationTitle"`
	// Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
	// identification corresponds to the project number as only one brand per project can be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly pulumi.BoolOutput `pulumi:"orgInternalOnly"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringOutput `pulumi:"supportEmail"`
}

// NewBrand registers a new resource with the given unique name, arguments, and options.
func NewBrand(ctx *pulumi.Context,
	name string, args *BrandArgs, opts ...pulumi.ResourceOption) (*Brand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationTitle == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationTitle'")
	}
	if args.SupportEmail == nil {
		return nil, errors.New("invalid value for required argument 'SupportEmail'")
	}
	var resource Brand
	err := ctx.RegisterResource("gcp:iap/brand:Brand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBrand gets an existing Brand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBrand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BrandState, opts ...pulumi.ResourceOption) (*Brand, error) {
	var resource Brand
	err := ctx.ReadResource("gcp:iap/brand:Brand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Brand resources.
type brandState struct {
	// Application name displayed on OAuth consent screen.
	ApplicationTitle *string `pulumi:"applicationTitle"`
	// Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
	// identification corresponds to the project number as only one brand per project can be created.
	Name *string `pulumi:"name"`
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly *bool `pulumi:"orgInternalOnly"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail *string `pulumi:"supportEmail"`
}

type BrandState struct {
	// Application name displayed on OAuth consent screen.
	ApplicationTitle pulumi.StringPtrInput
	// Output only. Identifier of the brand, in the format 'projects/{project_number}/brands/{brand_id}'. NOTE: The brand
	// identification corresponds to the project number as only one brand per project can be created.
	Name pulumi.StringPtrInput
	// Whether the brand is only intended for usage inside the GSuite organization only.
	OrgInternalOnly pulumi.BoolPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringPtrInput
}

func (BrandState) ElementType() reflect.Type {
	return reflect.TypeOf((*brandState)(nil)).Elem()
}

type brandArgs struct {
	// Application name displayed on OAuth consent screen.
	ApplicationTitle string `pulumi:"applicationTitle"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail string `pulumi:"supportEmail"`
}

// The set of arguments for constructing a Brand resource.
type BrandArgs struct {
	// Application name displayed on OAuth consent screen.
	ApplicationTitle pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Support email displayed on the OAuth consent screen. Can be either a
	// user or group email. When a user email is specified, the caller must
	// be the user with the associated email address. When a group email is
	// specified, the caller can be either a user or a service account which
	// is an owner of the specified group in Cloud Identity.
	SupportEmail pulumi.StringInput
}

func (BrandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*brandArgs)(nil)).Elem()
}

type BrandInput interface {
	pulumi.Input

	ToBrandOutput() BrandOutput
	ToBrandOutputWithContext(ctx context.Context) BrandOutput
}

func (Brand) ElementType() reflect.Type {
	return reflect.TypeOf((*Brand)(nil)).Elem()
}

func (i Brand) ToBrandOutput() BrandOutput {
	return i.ToBrandOutputWithContext(context.Background())
}

func (i Brand) ToBrandOutputWithContext(ctx context.Context) BrandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BrandOutput)
}

type BrandOutput struct {
	*pulumi.OutputState
}

func (BrandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BrandOutput)(nil)).Elem()
}

func (o BrandOutput) ToBrandOutput() BrandOutput {
	return o
}

func (o BrandOutput) ToBrandOutputWithContext(ctx context.Context) BrandOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BrandOutput{})
}
