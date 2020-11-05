// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An attestor that attests to container image artifacts.
//
// To get more information about Attestor, see:
//
// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/binary-authorization/)
//
// ## Example Usage
type Attestor struct {
	pulumi.CustomResourceState

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote AttestorAttestationAuthorityNoteOutput `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewAttestor registers a new resource with the given unique name, arguments, and options.
func NewAttestor(ctx *pulumi.Context,
	name string, args *AttestorArgs, opts ...pulumi.ResourceOption) (*Attestor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AttestationAuthorityNote == nil {
		return nil, errors.New("invalid value for required argument 'AttestationAuthorityNote'")
	}
	var resource Attestor
	err := ctx.RegisterResource("gcp:binaryauthorization/attestor:Attestor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestor gets an existing Attestor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestorState, opts ...pulumi.ResourceOption) (*Attestor, error) {
	var resource Attestor
	err := ctx.ReadResource("gcp:binaryauthorization/attestor:Attestor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attestor resources.
type attestorState struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote *AttestorAttestationAuthorityNote `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description *string `pulumi:"description"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type AttestorState struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote AttestorAttestationAuthorityNotePtrInput
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorState)(nil)).Elem()
}

type attestorArgs struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote AttestorAttestationAuthorityNote `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description *string `pulumi:"description"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Attestor resource.
type AttestorArgs struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	// Structure is documented below.
	AttestationAuthorityNote AttestorAttestationAuthorityNoteInput
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	Description pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttestorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorArgs)(nil)).Elem()
}
