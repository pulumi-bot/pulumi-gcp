// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get info about a Google Cloud IAP Client.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "foobar"
// 		project, err := organizations.LookupProject(ctx, &organizations.LookupProjectArgs{
// 			ProjectId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.LookupClient(ctx, &iap.LookupClientArgs{
// 			Brand:    fmt.Sprintf("%v%v%v", "projects/", project.Number, "/brands/[BRAND_NUMBER]"),
// 			ClientId: FOO.Apps.Googleusercontent.Com,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupClient(ctx *pulumi.Context, args *LookupClientArgs, opts ...pulumi.InvokeOption) (*LookupClientResult, error) {
	var rv LookupClientResult
	err := ctx.Invoke("gcp:iap/getClient:getClient", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClient.
type LookupClientArgs struct {
	// The name of the brand.
	Brand string `pulumi:"brand"`
	// The clientId of the brand.
	ClientId string `pulumi:"clientId"`
}

// A collection of values returned by getClient.
type LookupClientResult struct {
	Brand       string `pulumi:"brand"`
	ClientId    string `pulumi:"clientId"`
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Secret string `pulumi:"secret"`
}

func LookupClientOutput(ctx *pulumi.Context, args LookupClientOutputArgs, opts ...pulumi.InvokeOption) LookupClientResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClientResult, error) {
			args := v.(LookupClientArgs)
			r, err := LookupClient(ctx, &args, opts...)
			return *r, err
		}).(LookupClientResultOutput)
}

// A collection of arguments for invoking getClient.
type LookupClientOutputArgs struct {
	// The name of the brand.
	Brand pulumi.StringInput `pulumi:"brand"`
	// The clientId of the brand.
	ClientId pulumi.StringInput `pulumi:"clientId"`
}

func (LookupClientOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientArgs)(nil)).Elem()
}

// A collection of values returned by getClient.
type LookupClientResultOutput struct{ *pulumi.OutputState }

func (LookupClientResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClientResult)(nil)).Elem()
}

func (o LookupClientResultOutput) ToLookupClientResultOutput() LookupClientResultOutput {
	return o
}

func (o LookupClientResultOutput) ToLookupClientResultOutputWithContext(ctx context.Context) LookupClientResultOutput {
	return o
}

func (o LookupClientResultOutput) Brand() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Brand }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClientResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClientResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClientResult) string { return v.Secret }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClientResultOutput{})
}
