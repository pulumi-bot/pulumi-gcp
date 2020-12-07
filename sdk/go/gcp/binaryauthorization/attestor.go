// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"context"
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
// ### Binary Authorization Attestor Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		note, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("Attestor Note"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = binaryauthorization.NewAttestor(ctx, "attestor", &binaryauthorization.AttestorArgs{
// 			AttestationAuthorityNote: &binaryauthorization.AttestorAttestationAuthorityNoteArgs{
// 				NoteReference: note.Name,
// 				PublicKeys: binaryauthorization.AttestorAttestationAuthorityNotePublicKeyArray{
// 					&binaryauthorization.AttestorAttestationAuthorityNotePublicKeyArgs{
// 						AsciiArmoredPgpPublicKey: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "mQENBFtP0doBCADF+joTiXWKVuP8kJt3fgpBSjT9h8ezMfKA4aXZctYLx5wslWQl\n", "bB7Iu2ezkECNzoEeU7WxUe8a61pMCh9cisS9H5mB2K2uM4Jnf8tgFeXn3akJDVo0\n", "oR1IC+Dp9mXbRSK3MAvKkOwWlG99sx3uEdvmeBRHBOO+grchLx24EThXFOyP9Fk6\n", "V39j6xMjw4aggLD15B4V0v9JqBDdJiIYFzszZDL6pJwZrzcP0z8JO4rTZd+f64bD\n", "Mpj52j/pQfA8lZHOaAgb1OrthLdMrBAjoDjArV4Ek7vSbrcgYWcI6BhsQrFoxKdX\n", "83TZKai55ZCfCLIskwUIzA1NLVwyzCS+fSN/ABEBAAG0KCJUZXN0IEF0dGVzdG9y\n", "IiA8ZGFuYWhvZmZtYW5AZ29vZ2xlLmNvbT6JAU4EEwEIADgWIQRfWkqHt6hpTA1L\n", "uY060eeM4dc66AUCW0/R2gIbLwULCQgHAgYVCgkICwIEFgIDAQIeAQIXgAAKCRA6\n", "0eeM4dc66HdpCAC4ot3b0OyxPb0Ip+WT2U0PbpTBPJklesuwpIrM4Lh0N+1nVRLC\n", "51WSmVbM8BiAFhLbN9LpdHhds1kUrHF7+wWAjdR8sqAj9otc6HGRM/3qfa2qgh+U\n", "WTEk/3us/rYSi7T7TkMuutRMIa1IkR13uKiW56csEMnbOQpn9rDqwIr5R8nlZP5h\n", "MAU9vdm1DIv567meMqTaVZgR3w7bck2P49AO8lO5ERFpVkErtu/98y+rUy9d789l\n", "+OPuS1NGnxI1YKsNaWJF4uJVuvQuZ1twrhCbGNtVorO2U12+cEq+YtUxj7kmdOC1\n", "qoIRW6y0+UlAc+MbqfL0ziHDOAmcqz1GnROg\n", "=6Bvm\n")),
// 					},
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
// ### Binary Authorization Attestor Kms
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/binaryauthorization"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/containeranalysis"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
// 			Location: pulumi.String("global"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKey(ctx, "crypto_key", &kms.CryptoKeyArgs{
// 			KeyRing: keyring.ID(),
// 			Purpose: pulumi.String("ASYMMETRIC_SIGN"),
// 			VersionTemplate: &kms.CryptoKeyVersionTemplateArgs{
// 				Algorithm: pulumi.String("RSA_SIGN_PKCS1_4096_SHA512"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		note, err := containeranalysis.NewNote(ctx, "note", &containeranalysis.NoteArgs{
// 			AttestationAuthority: &containeranalysis.NoteAttestationAuthorityArgs{
// 				Hint: &containeranalysis.NoteAttestationAuthorityHintArgs{
// 					HumanReadableName: pulumi.String("Attestor Note"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = binaryauthorization.NewAttestor(ctx, "attestor", &binaryauthorization.AttestorArgs{
// 			AttestationAuthorityNote: &binaryauthorization.AttestorAttestationAuthorityNoteArgs{
// 				NoteReference: note.Name,
// 				PublicKeys: binaryauthorization.AttestorAttestationAuthorityNotePublicKeyArray{
// 					&binaryauthorization.AttestorAttestationAuthorityNotePublicKeyArgs{
// 						Id: version.ApplyT(func(version kms.GetKMSCryptoKeyVersionResult) (string, error) {
// 							return version.Id, nil
// 						}).(pulumi.StringOutput),
// 						PkixPublicKey: &binaryauthorization.AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs{
// 							PublicKeyPem: version.ApplyT(func(version kms.GetKMSCryptoKeyVersionResult) (string, error) {
// 								return version.PublicKeys[0].Pem, nil
// 							}).(pulumi.StringOutput),
// 							SignatureAlgorithm: version.ApplyT(func(version kms.GetKMSCryptoKeyVersionResult) (string, error) {
// 								return version.PublicKeys[0].Algorithm, nil
// 							}).(pulumi.StringOutput),
// 						},
// 					},
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
//
// ## Import
//
// Attestor can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:binaryauthorization/attestor:Attestor default projects/{{project}}/attestors/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:binaryauthorization/attestor:Attestor default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:binaryauthorization/attestor:Attestor default {{name}}
// ```
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

type AttestorInput interface {
	pulumi.Input

	ToAttestorOutput() AttestorOutput
	ToAttestorOutputWithContext(ctx context.Context) AttestorOutput
}

func (Attestor) ElementType() reflect.Type {
	return reflect.TypeOf((*Attestor)(nil)).Elem()
}

func (i Attestor) ToAttestorOutput() AttestorOutput {
	return i.ToAttestorOutputWithContext(context.Background())
}

func (i Attestor) ToAttestorOutputWithContext(ctx context.Context) AttestorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestorOutput)
}

type AttestorOutput struct {
	*pulumi.OutputState
}

func (AttestorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttestorOutput)(nil)).Elem()
}

func (o AttestorOutput) ToAttestorOutput() AttestorOutput {
	return o
}

func (o AttestorOutput) ToAttestorOutputWithContext(ctx context.Context) AttestorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AttestorOutput{})
}
