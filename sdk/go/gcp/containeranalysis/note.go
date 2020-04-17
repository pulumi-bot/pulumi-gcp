// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a detailed description of a Note.
//
//
// To get more information about Note, see:
//
// * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/container-analysis/)
type Note struct {
	pulumi.CustomResourceState

	// -
	// (Required)
	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.  Structure is documented below.
	AttestationAuthority NoteAttestationAuthorityOutput `pulumi:"attestationAuthority"`
	// -
	// (Required)
	// The name of the note.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOption) (*Note, error) {
	if args == nil || args.AttestationAuthority == nil {
		return nil, errors.New("missing required argument 'AttestationAuthority'")
	}
	if args == nil {
		args = &NoteArgs{}
	}
	var resource Note
	err := ctx.RegisterResource("gcp:containeranalysis/note:Note", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteState, opts ...pulumi.ResourceOption) (*Note, error) {
	var resource Note
	err := ctx.ReadResource("gcp:containeranalysis/note:Note", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Note resources.
type noteState struct {
	// -
	// (Required)
	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.  Structure is documented below.
	AttestationAuthority *NoteAttestationAuthority `pulumi:"attestationAuthority"`
	// -
	// (Required)
	// The name of the note.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type NoteState struct {
	// -
	// (Required)
	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.  Structure is documented below.
	AttestationAuthority NoteAttestationAuthorityPtrInput
	// -
	// (Required)
	// The name of the note.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteState)(nil)).Elem()
}

type noteArgs struct {
	// -
	// (Required)
	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.  Structure is documented below.
	AttestationAuthority NoteAttestationAuthority `pulumi:"attestationAuthority"`
	// -
	// (Required)
	// The name of the note.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	// -
	// (Required)
	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.  Structure is documented below.
	AttestationAuthority NoteAttestationAuthorityInput
	// -
	// (Required)
	// The name of the note.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (NoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteArgs)(nil)).Elem()
}
