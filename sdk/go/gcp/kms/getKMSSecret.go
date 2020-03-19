// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source allows you to use data encrypted with Google Cloud KMS
// within your resource definitions.
//
// For more information see
// [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
//
// > **NOTE**: Using this data provider will allow you to conceal secret data within your
// resource definitions, but it does not take care of protecting that data in the
// logging output, plan output, or state output.  Please take care to secure your secret
// data outside of resource definitions.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_secret.html.markdown.
func GetKMSSecret(ctx *pulumi.Context, args *GetKMSSecretArgs, opts ...pulumi.InvokeOption) (*GetKMSSecretResult, error) {
	var rv GetKMSSecretResult
	err := ctx.Invoke("gcp:kms/getKMSSecret:getKMSSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKMSSecret.
type GetKMSSecretArgs struct {
	// The ciphertext to be decrypted, encoded in base64
	Ciphertext string `pulumi:"ciphertext"`
	// The id of the CryptoKey that will be used to
	// decrypt the provided ciphertext. This is represented by the format
	// `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
	CryptoKey string `pulumi:"cryptoKey"`
}


// A collection of values returned by getKMSSecret.
type GetKMSSecretResult struct {
	Ciphertext string `pulumi:"ciphertext"`
	CryptoKey string `pulumi:"cryptoKey"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Contains the result of decrypting the provided ciphertext.
	Plaintext string `pulumi:"plaintext"`
}

