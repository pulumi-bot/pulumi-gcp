// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containeranalysis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An occurrence is an instance of a Note, or type of analysis that
// can be done for a resource.
//
// To get more information about Occurrence, see:
//
// * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/container-analysis/)
//
// ## Example Usage
//
// ## Import
//
// Occurrence can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:containeranalysis/occurence:Occurence default projects/{{project}}/occurrences/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:containeranalysis/occurence:Occurence default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:containeranalysis/occurence:Occurence default {{name}}
// ```
type Occurence struct {
	pulumi.CustomResourceState

	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	// Structure is documented below.
	Attestation OccurenceAttestationOutput `pulumi:"attestation"`
	// The time when the repository was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The note kind which explicitly denotes which of the occurrence details are specified. This field can be used as a filter
	// in list requests.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the occurrence.
	Name pulumi.StringOutput `pulumi:"name"`
	// The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName pulumi.StringOutput `pulumi:"noteName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A description of actions that can be taken to remedy the note.
	Remediation pulumi.StringPtrOutput `pulumi:"remediation"`
	// Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceUri pulumi.StringOutput `pulumi:"resourceUri"`
	// The time when the repository was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOccurence registers a new resource with the given unique name, arguments, and options.
func NewOccurence(ctx *pulumi.Context,
	name string, args *OccurenceArgs, opts ...pulumi.ResourceOption) (*Occurence, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Attestation == nil {
		return nil, errors.New("invalid value for required argument 'Attestation'")
	}
	if args.NoteName == nil {
		return nil, errors.New("invalid value for required argument 'NoteName'")
	}
	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	var resource Occurence
	err := ctx.RegisterResource("gcp:containeranalysis/occurence:Occurence", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOccurence gets an existing Occurence resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOccurence(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OccurenceState, opts ...pulumi.ResourceOption) (*Occurence, error) {
	var resource Occurence
	err := ctx.ReadResource("gcp:containeranalysis/occurence:Occurence", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Occurence resources.
type occurenceState struct {
	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	// Structure is documented below.
	Attestation *OccurenceAttestation `pulumi:"attestation"`
	// The time when the repository was created.
	CreateTime *string `pulumi:"createTime"`
	// The note kind which explicitly denotes which of the occurrence details are specified. This field can be used as a filter
	// in list requests.
	Kind *string `pulumi:"kind"`
	// The name of the occurrence.
	Name *string `pulumi:"name"`
	// The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName *string `pulumi:"noteName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A description of actions that can be taken to remedy the note.
	Remediation *string `pulumi:"remediation"`
	// Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceUri *string `pulumi:"resourceUri"`
	// The time when the repository was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type OccurenceState struct {
	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	// Structure is documented below.
	Attestation OccurenceAttestationPtrInput
	// The time when the repository was created.
	CreateTime pulumi.StringPtrInput
	// The note kind which explicitly denotes which of the occurrence details are specified. This field can be used as a filter
	// in list requests.
	Kind pulumi.StringPtrInput
	// The name of the occurrence.
	Name pulumi.StringPtrInput
	// The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A description of actions that can be taken to remedy the note.
	Remediation pulumi.StringPtrInput
	// Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceUri pulumi.StringPtrInput
	// The time when the repository was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (OccurenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*occurenceState)(nil)).Elem()
}

type occurenceArgs struct {
	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	// Structure is documented below.
	Attestation OccurenceAttestation `pulumi:"attestation"`
	// The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName string `pulumi:"noteName"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A description of actions that can be taken to remedy the note.
	Remediation *string `pulumi:"remediation"`
	// Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a Occurence resource.
type OccurenceArgs struct {
	// Occurrence that represents a single "attestation". The authenticity
	// of an attestation can be verified using the attached signature.
	// If the verifier trusts the public key of the signer, then verifying
	// the signature is sufficient to establish trust. In this circumstance,
	// the authority to which this attestation is attached is primarily
	// useful for lookup (how to find this attestation if you already
	// know the authority and artifact to be verified) and intent (for
	// which authority this attestation was intended to sign.
	// Structure is documented below.
	Attestation OccurenceAttestationInput
	// The analysis note associated with this occurrence, in the form of
	// projects/[PROJECT]/notes/[NOTE_ID]. This field can be used as a
	// filter in list requests.
	NoteName pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A description of actions that can be taken to remedy the note.
	Remediation pulumi.StringPtrInput
	// Required. Immutable. A URI that represents the resource for which
	// the occurrence applies. For example,
	// https://gcr.io/project/image@sha256:123abc for a Docker image.
	ResourceUri pulumi.StringInput
}

func (OccurenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*occurenceArgs)(nil)).Elem()
}

type OccurenceInput interface {
	pulumi.Input

	ToOccurenceOutput() OccurenceOutput
	ToOccurenceOutputWithContext(ctx context.Context) OccurenceOutput
}

func (*Occurence) ElementType() reflect.Type {
	return reflect.TypeOf((*Occurence)(nil))
}

func (i *Occurence) ToOccurenceOutput() OccurenceOutput {
	return i.ToOccurenceOutputWithContext(context.Background())
}

func (i *Occurence) ToOccurenceOutputWithContext(ctx context.Context) OccurenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurenceOutput)
}

func (i *Occurence) ToOccurencePtrOutput() OccurencePtrOutput {
	return i.ToOccurencePtrOutputWithContext(context.Background())
}

func (i *Occurence) ToOccurencePtrOutputWithContext(ctx context.Context) OccurencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurencePtrOutput)
}

type OccurencePtrInput interface {
	pulumi.Input

	ToOccurencePtrOutput() OccurencePtrOutput
	ToOccurencePtrOutputWithContext(ctx context.Context) OccurencePtrOutput
}

type occurencePtrType OccurenceArgs

func (*occurencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Occurence)(nil))
}

func (i *occurencePtrType) ToOccurencePtrOutput() OccurencePtrOutput {
	return i.ToOccurencePtrOutputWithContext(context.Background())
}

func (i *occurencePtrType) ToOccurencePtrOutputWithContext(ctx context.Context) OccurencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurenceOutput).ToOccurencePtrOutput()
}

// OccurenceArrayInput is an input type that accepts OccurenceArray and OccurenceArrayOutput values.
// You can construct a concrete instance of `OccurenceArrayInput` via:
//
//          OccurenceArray{ OccurenceArgs{...} }
type OccurenceArrayInput interface {
	pulumi.Input

	ToOccurenceArrayOutput() OccurenceArrayOutput
	ToOccurenceArrayOutputWithContext(context.Context) OccurenceArrayOutput
}

type OccurenceArray []OccurenceInput

func (OccurenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Occurence)(nil))
}

func (i OccurenceArray) ToOccurenceArrayOutput() OccurenceArrayOutput {
	return i.ToOccurenceArrayOutputWithContext(context.Background())
}

func (i OccurenceArray) ToOccurenceArrayOutputWithContext(ctx context.Context) OccurenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurenceArrayOutput)
}

// OccurenceMapInput is an input type that accepts OccurenceMap and OccurenceMapOutput values.
// You can construct a concrete instance of `OccurenceMapInput` via:
//
//          OccurenceMap{ "key": OccurenceArgs{...} }
type OccurenceMapInput interface {
	pulumi.Input

	ToOccurenceMapOutput() OccurenceMapOutput
	ToOccurenceMapOutputWithContext(context.Context) OccurenceMapOutput
}

type OccurenceMap map[string]OccurenceInput

func (OccurenceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Occurence)(nil))
}

func (i OccurenceMap) ToOccurenceMapOutput() OccurenceMapOutput {
	return i.ToOccurenceMapOutputWithContext(context.Background())
}

func (i OccurenceMap) ToOccurenceMapOutputWithContext(ctx context.Context) OccurenceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurenceMapOutput)
}

type OccurenceOutput struct {
	*pulumi.OutputState
}

func (OccurenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Occurence)(nil))
}

func (o OccurenceOutput) ToOccurenceOutput() OccurenceOutput {
	return o
}

func (o OccurenceOutput) ToOccurenceOutputWithContext(ctx context.Context) OccurenceOutput {
	return o
}

func (o OccurenceOutput) ToOccurencePtrOutput() OccurencePtrOutput {
	return o.ToOccurencePtrOutputWithContext(context.Background())
}

func (o OccurenceOutput) ToOccurencePtrOutputWithContext(ctx context.Context) OccurencePtrOutput {
	return o.ApplyT(func(v Occurence) *Occurence {
		return &v
	}).(OccurencePtrOutput)
}

type OccurencePtrOutput struct {
	*pulumi.OutputState
}

func (OccurencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Occurence)(nil))
}

func (o OccurencePtrOutput) ToOccurencePtrOutput() OccurencePtrOutput {
	return o
}

func (o OccurencePtrOutput) ToOccurencePtrOutputWithContext(ctx context.Context) OccurencePtrOutput {
	return o
}

type OccurenceArrayOutput struct{ *pulumi.OutputState }

func (OccurenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Occurence)(nil))
}

func (o OccurenceArrayOutput) ToOccurenceArrayOutput() OccurenceArrayOutput {
	return o
}

func (o OccurenceArrayOutput) ToOccurenceArrayOutputWithContext(ctx context.Context) OccurenceArrayOutput {
	return o
}

func (o OccurenceArrayOutput) Index(i pulumi.IntInput) OccurenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Occurence {
		return vs[0].([]Occurence)[vs[1].(int)]
	}).(OccurenceOutput)
}

type OccurenceMapOutput struct{ *pulumi.OutputState }

func (OccurenceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Occurence)(nil))
}

func (o OccurenceMapOutput) ToOccurenceMapOutput() OccurenceMapOutput {
	return o
}

func (o OccurenceMapOutput) ToOccurenceMapOutputWithContext(ctx context.Context) OccurenceMapOutput {
	return o
}

func (o OccurenceMapOutput) MapIndex(k pulumi.StringInput) OccurenceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Occurence {
		return vs[0].(map[string]Occurence)[vs[1].(string)]
	}).(OccurenceOutput)
}

func init() {
	pulumi.RegisterOutputType(OccurenceOutput{})
	pulumi.RegisterOutputType(OccurencePtrOutput{})
	pulumi.RegisterOutputType(OccurenceArrayOutput{})
	pulumi.RegisterOutputType(OccurenceMapOutput{})
}
