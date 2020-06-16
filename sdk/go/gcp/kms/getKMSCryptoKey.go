// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides access to a Google Cloud Platform KMS CryptoKey. For more information see
// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key)
// and
// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys).
//
// A CryptoKey is an interface to key material which can be used to encrypt and decrypt data. A CryptoKey belongs to a
// Google Cloud KMS KeyRing.
//
// {{% examples %}}
// ## Example Usage
// {{% example %}}
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myKeyRing, err := kms.LookupKMSKeyRing(ctx, &kms.LookupKMSKeyRingArgs{
// 			Name:     "my-key-ring",
// 			Location: "us-central1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err := kms.LookupKMSCryptoKey(ctx, &kms.LookupKMSCryptoKeyArgs{
// 			Name:    "my-crypto-key",
// 			KeyRing: myKeyRing.SelfLink,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// {{% /example %}}
// {{% /examples %}}
func GetKMSCryptoKey(ctx *pulumi.Context, args *GetKMSCryptoKeyArgs, opts ...pulumi.InvokeOption) (*GetKMSCryptoKeyResult, error) {
	var rv GetKMSCryptoKeyResult
	err := ctx.Invoke("gcp:kms/getKMSCryptoKey:getKMSCryptoKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKMSCryptoKey.
type GetKMSCryptoKeyArgs struct {
	// The `selfLink` of the Google Cloud Platform KeyRing to which the key belongs.
	KeyRing string `pulumi:"keyRing"`
	// The CryptoKey's name.
	// A CryptoKey’s name belonging to the specified Google Cloud Platform KeyRing and match the regular expression `[a-zA-Z0-9_-]{1,63}`
	Name string `pulumi:"name"`
}

// A collection of values returned by getKMSCryptoKey.
type GetKMSCryptoKeyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string            `pulumi:"id"`
	KeyRing string            `pulumi:"keyRing"`
	Labels  map[string]string `pulumi:"labels"`
	Name    string            `pulumi:"name"`
	// Defines the cryptographic capabilities of the key.
	Purpose string `pulumi:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as
	// the primary. The first rotation will take place after the specified period. The rotation period has the format
	// of a decimal number with up to 9 fractional digits, followed by the letter s (seconds).
	RotationPeriod string `pulumi:"rotationPeriod"`
	// The self link of the created CryptoKey. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}`.
	SelfLink         string                           `pulumi:"selfLink"`
	VersionTemplates []GetKMSCryptoKeyVersionTemplate `pulumi:"versionTemplates"`
}
