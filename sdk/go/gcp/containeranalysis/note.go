// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"context"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/containeranalysis"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/containeranalysis"
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
//
// ## Import
//
// Note can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:containeranalysis/note:Note default projects/{{project}}/notes/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:containeranalysis/note:Note default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:containeranalysis/note:Note default {{name}}
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
	// project.
	// Structure is documented below.
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
	// URLs associated with this note and related metadata.
	// Structure is documented below.
	RelatedUrls NoteRelatedUrlArrayOutput `pulumi:"relatedUrls"`
	// A one sentence description of the note.
	ShortDescription pulumi.StringPtrOutput `pulumi:"shortDescription"`
	// The time this note was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOption) (*Note, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttestationAuthority == nil {
		return nil, errors.New("invalid value for required argument 'AttestationAuthority'")
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
	// project.
	// Structure is documented below.
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
	// URLs associated with this note and related metadata.
	// Structure is documented below.
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
	// project.
	// Structure is documented below.
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
	// URLs associated with this note and related metadata.
	// Structure is documented below.
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
	// project.
	// Structure is documented below.
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
	// URLs associated with this note and related metadata.
	// Structure is documented below.
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
	// project.
	// Structure is documented below.
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
	// URLs associated with this note and related metadata.
	// Structure is documented below.
	RelatedUrls NoteRelatedUrlArrayInput
	// A one sentence description of the note.
	ShortDescription pulumi.StringPtrInput
}

func (NoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteArgs)(nil)).Elem()
}

type NoteInput interface {
	pulumi.Input

	ToNoteOutput() NoteOutput
	ToNoteOutputWithContext(ctx context.Context) NoteOutput
}

func (*Note) ElementType() reflect.Type {
	return reflect.TypeOf((*Note)(nil))
}

func (i *Note) ToNoteOutput() NoteOutput {
	return i.ToNoteOutputWithContext(context.Background())
}

func (i *Note) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteOutput)
}

func (i *Note) ToNotePtrOutput() NotePtrOutput {
	return i.ToNotePtrOutputWithContext(context.Background())
}

func (i *Note) ToNotePtrOutputWithContext(ctx context.Context) NotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotePtrOutput)
}

type NotePtrInput interface {
	pulumi.Input

	ToNotePtrOutput() NotePtrOutput
	ToNotePtrOutputWithContext(ctx context.Context) NotePtrOutput
}

type notePtrType NoteArgs

func (*notePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil))
}

func (i *notePtrType) ToNotePtrOutput() NotePtrOutput {
	return i.ToNotePtrOutputWithContext(context.Background())
}

func (i *notePtrType) ToNotePtrOutputWithContext(ctx context.Context) NotePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteOutput).ToNotePtrOutput()
}

// NoteArrayInput is an input type that accepts NoteArray and NoteArrayOutput values.
// You can construct a concrete instance of `NoteArrayInput` via:
//
//          NoteArray{ NoteArgs{...} }
type NoteArrayInput interface {
	pulumi.Input

	ToNoteArrayOutput() NoteArrayOutput
	ToNoteArrayOutputWithContext(context.Context) NoteArrayOutput
}

type NoteArray []NoteInput

func (NoteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Note)(nil))
}

func (i NoteArray) ToNoteArrayOutput() NoteArrayOutput {
	return i.ToNoteArrayOutputWithContext(context.Background())
}

func (i NoteArray) ToNoteArrayOutputWithContext(ctx context.Context) NoteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteArrayOutput)
}

// NoteMapInput is an input type that accepts NoteMap and NoteMapOutput values.
// You can construct a concrete instance of `NoteMapInput` via:
//
//          NoteMap{ "key": NoteArgs{...} }
type NoteMapInput interface {
	pulumi.Input

	ToNoteMapOutput() NoteMapOutput
	ToNoteMapOutputWithContext(context.Context) NoteMapOutput
}

type NoteMap map[string]NoteInput

func (NoteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Note)(nil))
}

func (i NoteMap) ToNoteMapOutput() NoteMapOutput {
	return i.ToNoteMapOutputWithContext(context.Background())
}

func (i NoteMap) ToNoteMapOutputWithContext(ctx context.Context) NoteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteMapOutput)
}

type NoteOutput struct {
	*pulumi.OutputState
}

func (NoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Note)(nil))
}

func (o NoteOutput) ToNoteOutput() NoteOutput {
	return o
}

func (o NoteOutput) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return o
}

func (o NoteOutput) ToNotePtrOutput() NotePtrOutput {
	return o.ToNotePtrOutputWithContext(context.Background())
}

func (o NoteOutput) ToNotePtrOutputWithContext(ctx context.Context) NotePtrOutput {
	return o.ApplyT(func(v Note) *Note {
		return &v
	}).(NotePtrOutput)
}

type NotePtrOutput struct {
	*pulumi.OutputState
}

func (NotePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil))
}

func (o NotePtrOutput) ToNotePtrOutput() NotePtrOutput {
	return o
}

func (o NotePtrOutput) ToNotePtrOutputWithContext(ctx context.Context) NotePtrOutput {
	return o
}

type NoteArrayOutput struct{ *pulumi.OutputState }

func (NoteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Note)(nil))
}

func (o NoteArrayOutput) ToNoteArrayOutput() NoteArrayOutput {
	return o
}

func (o NoteArrayOutput) ToNoteArrayOutputWithContext(ctx context.Context) NoteArrayOutput {
	return o
}

func (o NoteArrayOutput) Index(i pulumi.IntInput) NoteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Note {
		return vs[0].([]Note)[vs[1].(int)]
	}).(NoteOutput)
}

type NoteMapOutput struct{ *pulumi.OutputState }

func (NoteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Note)(nil))
}

func (o NoteMapOutput) ToNoteMapOutput() NoteMapOutput {
	return o
}

func (o NoteMapOutput) ToNoteMapOutputWithContext(ctx context.Context) NoteMapOutput {
	return o
}

func (o NoteMapOutput) MapIndex(k pulumi.StringInput) NoteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Note {
		return vs[0].(map[string]Note)[vs[1].(string)]
	}).(NoteOutput)
}

func init() {
	pulumi.RegisterOutputType(NoteOutput{})
	pulumi.RegisterOutputType(NotePtrOutput{})
	pulumi.RegisterOutputType(NoteArrayOutput{})
	pulumi.RegisterOutputType(NoteMapOutput{})
}
