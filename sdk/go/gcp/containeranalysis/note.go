// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Container Analysis note is a high-level piece of metadata that
// describes a type of analysis that can be done for a resource.
//
// To get more information about Note, see:
//
// * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/container-analysis/)
//     * [Creating Attestations (Occurrences)](https://cloud.google.com/binary-authorization/docs/making-attestations)
//
// ## Example Usage
// ### Container Analysis Note Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("Attestor Note"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Container Analysis Note Attestation Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("Attestor Note"),
// 				},
// 			},
// 			ExpirationTime:  pulumi.String("2120-10-02T15:01:23.045123456Z"),
// 			LongDescription: pulumi.String("a longer description of test note"),
// 			RelatedUrls: containeranalysis.NoteRelatedUrlArray{
// 				&containeranalysis.NoteRelatedUrlArgs{
// 					Label: pulumi.String("foo"),
// 					Url:   pulumi.String("some.url"),
// 				},
// 				&containeranalysis.NoteRelatedUrlArgs{
// 					Url: pulumi.String("google.com"),
// 				},
// 			},
// 			ShortDescription: pulumi.String("test note"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Note struct {
	pulumi.CustomResourceState

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
	// The time this note was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime pulumi.StringPtrOutput `pulumi:"expirationTime"`
	// The type of analysis this note describes
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A detailed description of the note
	LongDescription pulumi.StringPtrOutput `pulumi:"longDescription"`
	// The name of the note.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Names of other notes related to this note.
	RelatedNoteNames pulumi.StringArrayOutput `pulumi:"relatedNoteNames"`
	// URLs associated with this note and related metadata.  Structure is documented below.
	RelatedUrls NoteRelatedUrlArrayOutput `pulumi:"relatedUrls"`
	// A one sentence description of the note.
	ShortDescription pulumi.StringPtrOutput `pulumi:"shortDescription"`
	// The time this note was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
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
	// The time this note was created.
	CreateTime *string `pulumi:"createTime"`
	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// The type of analysis this note describes
	Kind *string `pulumi:"kind"`
	// A detailed description of the note
	LongDescription *string `pulumi:"longDescription"`
	// The name of the note.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Names of other notes related to this note.
	RelatedNoteNames []string `pulumi:"relatedNoteNames"`
	// URLs associated with this note and related metadata.  Structure is documented below.
	RelatedUrls []NoteRelatedUrl `pulumi:"relatedUrls"`
	// A one sentence description of the note.
	ShortDescription *string `pulumi:"shortDescription"`
	// The time this note was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type NoteState struct {
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
	// The time this note was created.
	CreateTime pulumi.StringPtrInput
	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// The type of analysis this note describes
	Kind pulumi.StringPtrInput
	// A detailed description of the note
	LongDescription pulumi.StringPtrInput
	// The name of the note.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Names of other notes related to this note.
	RelatedNoteNames pulumi.StringArrayInput
	// URLs associated with this note and related metadata.  Structure is documented below.
	RelatedUrls NoteRelatedUrlArrayInput
	// A one sentence description of the note.
	ShortDescription pulumi.StringPtrInput
	// The time this note was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (NoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteState)(nil)).Elem()
}

type noteArgs struct {
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
	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// A detailed description of the note
	LongDescription *string `pulumi:"longDescription"`
	// The name of the note.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Names of other notes related to this note.
	RelatedNoteNames []string `pulumi:"relatedNoteNames"`
	// URLs associated with this note and related metadata.  Structure is documented below.
	RelatedUrls []NoteRelatedUrl `pulumi:"relatedUrls"`
	// A one sentence description of the note.
	ShortDescription *string `pulumi:"shortDescription"`
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
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
	// Time of expiration for this note. Leave empty if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// A detailed description of the note
	LongDescription pulumi.StringPtrInput
	// The name of the note.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Names of other notes related to this note.
	RelatedNoteNames pulumi.StringArrayInput
	// URLs associated with this note and related metadata.  Structure is documented below.
	RelatedUrls NoteRelatedUrlArrayInput
	// A one sentence description of the note.
	ShortDescription pulumi.StringPtrInput
}

func (NoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteArgs)(nil)).Elem()
}
