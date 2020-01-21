// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package attestor

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/AttestorAttestationAuthorityNote"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/AttestorAttestationAuthorityNotePublicKey"
	"https:/github.com/pulumi/pulumi-gcp/binaryauthorization/AttestorAttestationAuthorityNotePublicKeyPkixPublicKey"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_attestor.html.markdown.
type Attestor struct {
	pulumi.CustomResourceState

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	AttestationAuthorityNote binaryauthorizationAttestorAttestationAuthorityNote.AttestorAttestationAuthorityNoteOutput `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewAttestor registers a new resource with the given unique name, arguments, and options.
func NewAttestor(ctx *pulumi.Context,
	name string, args *AttestorArgs, opts ...pulumi.ResourceOption) (*Attestor, error) {
	if args == nil || args.AttestationAuthorityNote == nil {
		return nil, errors.New("missing required argument 'AttestationAuthorityNote'")
	}
	if args == nil {
		args = &AttestorArgs{}
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
	AttestationAuthorityNote *binaryauthorizationAttestorAttestationAuthorityNote.AttestorAttestationAuthorityNote `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description *string `pulumi:"description"`
	// The resource name.
	Name *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

type AttestorState struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	AttestationAuthorityNote binaryauthorizationAttestorAttestationAuthorityNote.AttestorAttestationAuthorityNotePtrInput
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (AttestorState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorState)(nil)).Elem()
}

type attestorArgs struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	AttestationAuthorityNote binaryauthorizationAttestorAttestationAuthorityNote.AttestorAttestationAuthorityNote `pulumi:"attestationAuthorityNote"`
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description *string `pulumi:"description"`
	// The resource name.
	Name *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Attestor resource.
type AttestorArgs struct {
	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	AttestationAuthorityNote binaryauthorizationAttestorAttestationAuthorityNote.AttestorAttestationAuthorityNoteInput
	// A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
	Description pulumi.StringPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (AttestorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestorArgs)(nil)).Elem()
}

