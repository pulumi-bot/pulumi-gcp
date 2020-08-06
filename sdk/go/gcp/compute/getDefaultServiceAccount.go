// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve default service account for this project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_default, err := compute.GetDefaultServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("defaultAccount", _default.Email)
// 		return nil
// 	})
// }
// ```
func GetDefaultServiceAccount(ctx *pulumi.Context, args *GetDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetDefaultServiceAccountResult, error) {
	var rv GetDefaultServiceAccountResult
	err := ctx.Invoke("gcp:compute/getDefaultServiceAccount:getDefaultServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResult struct {
	// The display name for the service account.
	DisplayName string `pulumi:"displayName"`
	// Email address of the default service account used by VMs running in this project
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The fully-qualified name of the service account.
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId string `pulumi:"uniqueId"`
}
