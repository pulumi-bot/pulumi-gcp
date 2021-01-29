// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A `CryptoKey` represents a logical key that can be used for cryptographic operations.
//
// > **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
// Destroying a provider-managed CryptoKey will remove it from state
// and delete all CryptoKeyVersions, rendering the key unusable, but *will
// not delete the resource from the project.* When the provider destroys these keys,
// any data previously encrypted with these keys will be irrecoverable.
// For this reason, it is strongly recommended that you add lifecycle hooks
// to the resource to prevent accidental destruction.
//
// To get more information about CryptoKey, see:
//
// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
// * How-to Guides
//     * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
//
// ## Example Usage
// ### Kms Crypto Key Basic
//
// ```go
// package main
//
// import (
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
// 		_, err = kms.NewCryptoKey(ctx, "example_key", &kms.CryptoKeyArgs{
// 			KeyRing:        keyring.ID(),
// 			RotationPeriod: pulumi.String("100000s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Kms Crypto Key Asymmetric Sign
//
// ```go
// package main
//
// import (
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
// 		_, err = kms.NewCryptoKey(ctx, "example_asymmetric_sign_key", &kms.CryptoKeyArgs{
// 			KeyRing: keyring.ID(),
// 			Purpose: pulumi.String("ASYMMETRIC_SIGN"),
// 			VersionTemplate: &kms.CryptoKeyVersionTemplateArgs{
// 				Algorithm: pulumi.String("EC_SIGN_P384_SHA384"),
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
// CryptoKey can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
// ```
type CryptoKey struct {
	pulumi.CustomResourceState

	// The KeyRing that this key belongs to.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	KeyRing pulumi.StringOutput `pulumi:"keyRing"`
	// Labels with user-defined metadata to apply to this resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the CryptoKey.
	Name pulumi.StringOutput `pulumi:"name"`
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is `ENCRYPT_DECRYPT`.
	// Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
	Purpose pulumi.StringPtrOutput `pulumi:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter `s` (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod pulumi.StringPtrOutput `pulumi:"rotationPeriod"`
	SelfLink       pulumi.StringOutput    `pulumi:"selfLink"`
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation pulumi.BoolPtrOutput `pulumi:"skipInitialVersionCreation"`
	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate CryptoKeyVersionTemplateOutput `pulumi:"versionTemplate"`
}

// NewCryptoKey registers a new resource with the given unique name, arguments, and options.
func NewCryptoKey(ctx *pulumi.Context,
	name string, args *CryptoKeyArgs, opts ...pulumi.ResourceOption) (*CryptoKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyRing == nil {
		return nil, errors.New("invalid value for required argument 'KeyRing'")
	}
	var resource CryptoKey
	err := ctx.RegisterResource("gcp:kms/cryptoKey:CryptoKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCryptoKey gets an existing CryptoKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCryptoKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CryptoKeyState, opts ...pulumi.ResourceOption) (*CryptoKey, error) {
	var resource CryptoKey
	err := ctx.ReadResource("gcp:kms/cryptoKey:CryptoKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CryptoKey resources.
type cryptoKeyState struct {
	// The KeyRing that this key belongs to.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	KeyRing *string `pulumi:"keyRing"`
	// Labels with user-defined metadata to apply to this resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the CryptoKey.
	Name *string `pulumi:"name"`
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is `ENCRYPT_DECRYPT`.
	// Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
	Purpose *string `pulumi:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter `s` (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod *string `pulumi:"rotationPeriod"`
	SelfLink       *string `pulumi:"selfLink"`
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation *bool `pulumi:"skipInitialVersionCreation"`
	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate *CryptoKeyVersionTemplate `pulumi:"versionTemplate"`
}

type CryptoKeyState struct {
	// The KeyRing that this key belongs to.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	KeyRing pulumi.StringPtrInput
	// Labels with user-defined metadata to apply to this resource.
	Labels pulumi.StringMapInput
	// The resource name for the CryptoKey.
	Name pulumi.StringPtrInput
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is `ENCRYPT_DECRYPT`.
	// Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
	Purpose pulumi.StringPtrInput
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter `s` (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod pulumi.StringPtrInput
	SelfLink       pulumi.StringPtrInput
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation pulumi.BoolPtrInput
	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate CryptoKeyVersionTemplatePtrInput
}

func (CryptoKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyState)(nil)).Elem()
}

type cryptoKeyArgs struct {
	// The KeyRing that this key belongs to.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	KeyRing string `pulumi:"keyRing"`
	// Labels with user-defined metadata to apply to this resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the CryptoKey.
	Name *string `pulumi:"name"`
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is `ENCRYPT_DECRYPT`.
	// Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
	Purpose *string `pulumi:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter `s` (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod *string `pulumi:"rotationPeriod"`
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation *bool `pulumi:"skipInitialVersionCreation"`
	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate *CryptoKeyVersionTemplate `pulumi:"versionTemplate"`
}

// The set of arguments for constructing a CryptoKey resource.
type CryptoKeyArgs struct {
	// The KeyRing that this key belongs to.
	// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
	KeyRing pulumi.StringInput
	// Labels with user-defined metadata to apply to this resource.
	Labels pulumi.StringMapInput
	// The resource name for the CryptoKey.
	Name pulumi.StringPtrInput
	// The immutable purpose of this CryptoKey. See the
	// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
	// for possible inputs.
	// Default value is `ENCRYPT_DECRYPT`.
	// Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, and `ASYMMETRIC_DECRYPT`.
	Purpose pulumi.StringPtrInput
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter `s` (seconds). It must be greater than a day (ie, 86400).
	RotationPeriod pulumi.StringPtrInput
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	// You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
	SkipInitialVersionCreation pulumi.BoolPtrInput
	// A template describing settings for new crypto key versions.
	// Structure is documented below.
	VersionTemplate CryptoKeyVersionTemplatePtrInput
}

func (CryptoKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cryptoKeyArgs)(nil)).Elem()
}

type CryptoKeyInput interface {
	pulumi.Input

	ToCryptoKeyOutput() CryptoKeyOutput
	ToCryptoKeyOutputWithContext(ctx context.Context) CryptoKeyOutput
}

func (*CryptoKey) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKey)(nil))
}

func (i *CryptoKey) ToCryptoKeyOutput() CryptoKeyOutput {
	return i.ToCryptoKeyOutputWithContext(context.Background())
}

func (i *CryptoKey) ToCryptoKeyOutputWithContext(ctx context.Context) CryptoKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyOutput)
}

func (i *CryptoKey) ToCryptoKeyPtrOutput() CryptoKeyPtrOutput {
	return i.ToCryptoKeyPtrOutputWithContext(context.Background())
}

func (i *CryptoKey) ToCryptoKeyPtrOutputWithContext(ctx context.Context) CryptoKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyPtrOutput)
}

type CryptoKeyPtrInput interface {
	pulumi.Input

	ToCryptoKeyPtrOutput() CryptoKeyPtrOutput
	ToCryptoKeyPtrOutputWithContext(ctx context.Context) CryptoKeyPtrOutput
}

type cryptoKeyPtrType CryptoKeyArgs

func (*cryptoKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKey)(nil))
}

func (i *cryptoKeyPtrType) ToCryptoKeyPtrOutput() CryptoKeyPtrOutput {
	return i.ToCryptoKeyPtrOutputWithContext(context.Background())
}

func (i *cryptoKeyPtrType) ToCryptoKeyPtrOutputWithContext(ctx context.Context) CryptoKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyPtrOutput)
}

// CryptoKeyArrayInput is an input type that accepts CryptoKeyArray and CryptoKeyArrayOutput values.
// You can construct a concrete instance of `CryptoKeyArrayInput` via:
//
//          CryptoKeyArray{ CryptoKeyArgs{...} }
type CryptoKeyArrayInput interface {
	pulumi.Input

	ToCryptoKeyArrayOutput() CryptoKeyArrayOutput
	ToCryptoKeyArrayOutputWithContext(context.Context) CryptoKeyArrayOutput
}

type CryptoKeyArray []CryptoKeyInput

func (CryptoKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CryptoKey)(nil))
}

func (i CryptoKeyArray) ToCryptoKeyArrayOutput() CryptoKeyArrayOutput {
	return i.ToCryptoKeyArrayOutputWithContext(context.Background())
}

func (i CryptoKeyArray) ToCryptoKeyArrayOutputWithContext(ctx context.Context) CryptoKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyArrayOutput)
}

// CryptoKeyMapInput is an input type that accepts CryptoKeyMap and CryptoKeyMapOutput values.
// You can construct a concrete instance of `CryptoKeyMapInput` via:
//
//          CryptoKeyMap{ "key": CryptoKeyArgs{...} }
type CryptoKeyMapInput interface {
	pulumi.Input

	ToCryptoKeyMapOutput() CryptoKeyMapOutput
	ToCryptoKeyMapOutputWithContext(context.Context) CryptoKeyMapOutput
}

type CryptoKeyMap map[string]CryptoKeyInput

func (CryptoKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CryptoKey)(nil))
}

func (i CryptoKeyMap) ToCryptoKeyMapOutput() CryptoKeyMapOutput {
	return i.ToCryptoKeyMapOutputWithContext(context.Background())
}

func (i CryptoKeyMap) ToCryptoKeyMapOutputWithContext(ctx context.Context) CryptoKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CryptoKeyMapOutput)
}

type CryptoKeyOutput struct {
	*pulumi.OutputState
}

func (CryptoKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CryptoKey)(nil))
}

func (o CryptoKeyOutput) ToCryptoKeyOutput() CryptoKeyOutput {
	return o
}

func (o CryptoKeyOutput) ToCryptoKeyOutputWithContext(ctx context.Context) CryptoKeyOutput {
	return o
}

func (o CryptoKeyOutput) ToCryptoKeyPtrOutput() CryptoKeyPtrOutput {
	return o.ToCryptoKeyPtrOutputWithContext(context.Background())
}

func (o CryptoKeyOutput) ToCryptoKeyPtrOutputWithContext(ctx context.Context) CryptoKeyPtrOutput {
	return o.ApplyT(func(v CryptoKey) *CryptoKey {
		return &v
	}).(CryptoKeyPtrOutput)
}

type CryptoKeyPtrOutput struct {
	*pulumi.OutputState
}

func (CryptoKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CryptoKey)(nil))
}

func (o CryptoKeyPtrOutput) ToCryptoKeyPtrOutput() CryptoKeyPtrOutput {
	return o
}

func (o CryptoKeyPtrOutput) ToCryptoKeyPtrOutputWithContext(ctx context.Context) CryptoKeyPtrOutput {
	return o
}

type CryptoKeyArrayOutput struct{ *pulumi.OutputState }

func (CryptoKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CryptoKey)(nil))
}

func (o CryptoKeyArrayOutput) ToCryptoKeyArrayOutput() CryptoKeyArrayOutput {
	return o
}

func (o CryptoKeyArrayOutput) ToCryptoKeyArrayOutputWithContext(ctx context.Context) CryptoKeyArrayOutput {
	return o
}

func (o CryptoKeyArrayOutput) Index(i pulumi.IntInput) CryptoKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CryptoKey {
		return vs[0].([]CryptoKey)[vs[1].(int)]
	}).(CryptoKeyOutput)
}

type CryptoKeyMapOutput struct{ *pulumi.OutputState }

func (CryptoKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CryptoKey)(nil))
}

func (o CryptoKeyMapOutput) ToCryptoKeyMapOutput() CryptoKeyMapOutput {
	return o
}

func (o CryptoKeyMapOutput) ToCryptoKeyMapOutputWithContext(ctx context.Context) CryptoKeyMapOutput {
	return o
}

func (o CryptoKeyMapOutput) MapIndex(k pulumi.StringInput) CryptoKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CryptoKey {
		return vs[0].(map[string]CryptoKey)[vs[1].(string)]
	}).(CryptoKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(CryptoKeyOutput{})
	pulumi.RegisterOutputType(CryptoKeyPtrOutput{})
	pulumi.RegisterOutputType(CryptoKeyArrayOutput{})
	pulumi.RegisterOutputType(CryptoKeyMapOutput{})
}
