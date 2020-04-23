// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A `KeyRing` is a toplevel logical grouping of `CryptoKeys`.
//
//
// > **Note:** KeyRings cannot be deleted from Google Cloud Platform.
// Destroying a provider-managed KeyRing will remove it from state but
// *will not delete the resource on the server.*
//
//
// To get more information about KeyRing, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings)
// * How-to Guides
//     * [Creating a key ring](https://cloud.google.com/kms/docs/creating-keys#create_a_key_ring)
type KeyRing struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringOutput `pulumi:"location"`
	// -
	// (Required)
	// The resource name for the KeyRing.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project  pulumi.StringOutput `pulumi:"project"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewKeyRing registers a new resource with the given unique name, arguments, and options.
func NewKeyRing(ctx *pulumi.Context,
	name string, args *KeyRingArgs, opts ...pulumi.ResourceOption) (*KeyRing, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil {
		args = &KeyRingArgs{}
	}
	var resource KeyRing
	err := ctx.RegisterResource("gcp:kms/keyRing:KeyRing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRing gets an existing KeyRing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingState, opts ...pulumi.ResourceOption) (*KeyRing, error) {
	var resource KeyRing
	err := ctx.ReadResource("gcp:kms/keyRing:KeyRing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRing resources.
type keyRingState struct {
	// -
	// (Required)
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location *string `pulumi:"location"`
	// -
	// (Required)
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project  *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
}

type KeyRingState struct {
	// -
	// (Required)
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringPtrInput
	// -
	// (Required)
	// The resource name for the KeyRing.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project  pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
}

func (KeyRingState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingState)(nil)).Elem()
}

type keyRingArgs struct {
	// -
	// (Required)
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location string `pulumi:"location"`
	// -
	// (Required)
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a KeyRing resource.
type KeyRingArgs struct {
	// -
	// (Required)
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringInput
	// -
	// (Required)
	// The resource name for the KeyRing.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (KeyRingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingArgs)(nil)).Elem()
}
