// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
//
// > **NOTE:** Using this resource will allow you to conceal secret data within your
// resource definitions, but it does not take care of protecting that data in the
// logging output, plan output, or state output.  Please take care to secure your secret
// data outside of resource definitions.
//
// To get more information about SecretCiphertext, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys/encrypt)
// * How-to Guides
//     * [Encrypting and decrypting data with a symmetric key](https://cloud.google.com/kms/docs/encrypt-decrypt)
//
// > **Warning:** All arguments including `plaintext` and `additionalAuthenticatedData` will be stored in the raw
// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
//
// ## Example Usage
// ### Kms Secret Ciphertext Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
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
// 		cryptokey, err := kms.NewCryptoKey(ctx, "cryptokey", &kms.CryptoKeyArgs{
// 			KeyRing:        keyring.ID(),
// 			RotationPeriod: pulumi.String("100000s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		myPassword, err := kms.NewSecretCiphertext(ctx, "myPassword", &kms.SecretCiphertextArgs{
// 			CryptoKey: cryptokey.ID(),
// 			Plaintext: pulumi.String("my-secret-password"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewInstance(ctx, "instance", &compute.InstanceArgs{
// 			MachineType: pulumi.String("e2-medium"),
// 			Zone:        pulumi.String("us-central1-a"),
// 			BootDisk: &compute.InstanceBootDiskArgs{
// 				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
// 					Image: pulumi.String("debian-cloud/debian-9"),
// 				},
// 			},
// 			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
// 				&compute.InstanceNetworkInterfaceArgs{
// 					Network: pulumi.String("default"),
// 					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
// 						nil,
// 					},
// 				},
// 			},
// 			Metadata: pulumi.StringMap{
// 				"password": myPassword.Ciphertext,
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
// This resource does not support import.
type SecretCiphertext struct {
	pulumi.CustomResourceState

	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrOutput `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringOutput `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringOutput `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringOutput `pulumi:"plaintext"`
}

// NewSecretCiphertext registers a new resource with the given unique name, arguments, and options.
func NewSecretCiphertext(ctx *pulumi.Context,
	name string, args *SecretCiphertextArgs, opts ...pulumi.ResourceOption) (*SecretCiphertext, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CryptoKey == nil {
		return nil, errors.New("invalid value for required argument 'CryptoKey'")
	}
	if args.Plaintext == nil {
		return nil, errors.New("invalid value for required argument 'Plaintext'")
	}
	var resource SecretCiphertext
	err := ctx.RegisterResource("gcp:kms/secretCiphertext:SecretCiphertext", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretCiphertext gets an existing SecretCiphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCiphertext(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretCiphertextState, opts ...pulumi.ResourceOption) (*SecretCiphertext, error) {
	var resource SecretCiphertext
	err := ctx.ReadResource("gcp:kms/secretCiphertext:SecretCiphertext", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretCiphertext resources.
type secretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext *string `pulumi:"ciphertext"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey *string `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext *string `pulumi:"plaintext"`
}

type SecretCiphertextState struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// Contains the result of encrypting the provided plaintext, encoded in base64.
	Ciphertext pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringPtrInput
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringPtrInput
}

func (SecretCiphertextState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCiphertextState)(nil)).Elem()
}

type secretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData *string `pulumi:"additionalAuthenticatedData"`
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey string `pulumi:"cryptoKey"`
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext string `pulumi:"plaintext"`
}

// The set of arguments for constructing a SecretCiphertext resource.
type SecretCiphertextArgs struct {
	// The additional authenticated data used for integrity checks during encryption and decryption.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	AdditionalAuthenticatedData pulumi.StringPtrInput
	// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
	CryptoKey pulumi.StringInput
	// The plaintext to be encrypted.
	// **Note**: This property is sensitive and will not be displayed in the plan.
	Plaintext pulumi.StringInput
}

func (SecretCiphertextArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCiphertextArgs)(nil)).Elem()
}

type SecretCiphertextInput interface {
	pulumi.Input

	ToSecretCiphertextOutput() SecretCiphertextOutput
	ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput
}

func (SecretCiphertext) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretCiphertext)(nil)).Elem()
}

func (i SecretCiphertext) ToSecretCiphertextOutput() SecretCiphertextOutput {
	return i.ToSecretCiphertextOutputWithContext(context.Background())
}

func (i SecretCiphertext) ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCiphertextOutput)
}

type SecretCiphertextOutput struct {
	*pulumi.OutputState
}

func (SecretCiphertextOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretCiphertextOutput)(nil)).Elem()
}

func (o SecretCiphertextOutput) ToSecretCiphertextOutput() SecretCiphertextOutput {
	return o
}

func (o SecretCiphertextOutput) ToSecretCiphertextOutputWithContext(ctx context.Context) SecretCiphertextOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretCiphertextOutput{})
}
