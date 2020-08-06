// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the email address of a project's unique BigQuery service account.
//
// Each Google Cloud project has a unique service account used by BigQuery. When using
// BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
// this account needs to be granted the
// `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
//
// For more information see
// [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/kms"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bqSa, err := bigquery.GetDefaultServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kms.NewCryptoKeyIAMMember(ctx, "keySaUser", &kms.CryptoKeyIAMMemberArgs{
// 			CryptoKeyId: pulumi.Any(google_kms_crypto_key.Key.Id),
// 			Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
// 			Member:      pulumi.String(fmt.Sprintf("%v%v", "serviceAccount:", bqSa.Email)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDefaultServiceAccount(ctx *pulumi.Context, args *GetDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetDefaultServiceAccountResult, error) {
	var rv GetDefaultServiceAccountResult
	err := ctx.Invoke("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountArgs struct {
	// The project the unique service account was created for. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResult struct {
	// The email address of the service account. This value is often used to refer to the service account
	// in order to grant IAM permissions.
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Project string `pulumi:"project"`
}
