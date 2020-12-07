// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Gets an SSL Policy within GCE from its name, for use with Target HTTPS and Target SSL Proxies.
//     For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.LookupSSLPolicy(ctx, &compute.LookupSSLPolicyArgs{
// 			Name: "production-ssl-policy",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSSLPolicy(ctx *pulumi.Context, args *LookupSSLPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSSLPolicyResult, error) {
	var rv LookupSSLPolicyResult
	err := ctx.Invoke("gcp:compute/getSSLPolicy:getSSLPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSSLPolicy.
type LookupSSLPolicyArgs struct {
	// The name of the SSL Policy.
	Name string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getSSLPolicy.
type LookupSSLPolicyResult struct {
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// If the `profile` is `CUSTOM`, these are the custom encryption
	// ciphers supported by the profile. If the `profile` is *not* `CUSTOM`, this
	// attribute will be empty.
	CustomFeatures []string `pulumi:"customFeatures"`
	// Description of this SSL Policy.
	Description string `pulumi:"description"`
	// The set of enabled encryption ciphers as a result of the policy config
	EnabledFeatures []string `pulumi:"enabledFeatures"`
	// Fingerprint of this resource.
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The minimum supported TLS version of this policy.
	MinTlsVersion string `pulumi:"minTlsVersion"`
	Name          string `pulumi:"name"`
	// The Google-curated or custom profile used by this policy.
	Profile string  `pulumi:"profile"`
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
}
