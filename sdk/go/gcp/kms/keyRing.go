// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A `KeyRing` is a toplevel logical grouping of `CryptoKeys`.
//
// > **Note:** KeyRings cannot be deleted from Google Cloud Platform.
// Destroying a provider-managed KeyRing will remove it from state but
// *will not delete the resource from the project.*
//
// To get more information about KeyRing, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings)
// * How-to Guides
//     * [Creating a key ring](https://cloud.google.com/kms/docs/creating-keys#create_a_key_ring)
//
// ## Example Usage
type KeyRing struct {
	pulumi.CustomResourceState

	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringOutput `pulumi:"location"`
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
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location *string `pulumi:"location"`
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project  *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
}

type KeyRingState struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringPtrInput
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
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location string `pulumi:"location"`
	// The resource name for the KeyRing.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a KeyRing resource.
type KeyRingArgs struct {
	// The location for the KeyRing.
	// A full list of valid locations can be found by running `gcloud kms locations list`.
	Location pulumi.StringInput
	// The resource name for the KeyRing.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (KeyRingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingArgs)(nil)).Elem()
}

type KeyRingInput interface {
	pulumi.Input

	ToKeyRingOutput() KeyRingOutput
	ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput
}

func (KeyRing) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRing)(nil)).Elem()
}

func (i KeyRing) ToKeyRingOutput() KeyRingOutput {
	return i.ToKeyRingOutputWithContext(context.Background())
}

func (i KeyRing) ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyRingOutput)
}

type KeyRingOutput struct {
	*pulumi.OutputState
}

func (KeyRingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyRingOutput)(nil)).Elem()
}

func (o KeyRingOutput) ToKeyRingOutput() KeyRingOutput {
	return o
}

func (o KeyRingOutput) ToKeyRingOutputWithContext(ctx context.Context) KeyRingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KeyRingOutput{})
}
