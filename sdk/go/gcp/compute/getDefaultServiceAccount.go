// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve default service account for this project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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

func GetDefaultServiceAccountOutput(ctx *pulumi.Context, args GetDefaultServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetDefaultServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDefaultServiceAccountResult, error) {
			args := v.(GetDefaultServiceAccountArgs)
			r, err := GetDefaultServiceAccount(ctx, &args, opts...)
			return *r, err
		}).(GetDefaultServiceAccountResultOutput)
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountOutputArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetDefaultServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetDefaultServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDefaultServiceAccountResult)(nil)).Elem()
}

func (o GetDefaultServiceAccountResultOutput) ToGetDefaultServiceAccountResultOutput() GetDefaultServiceAccountResultOutput {
	return o
}

func (o GetDefaultServiceAccountResultOutput) ToGetDefaultServiceAccountResultOutputWithContext(ctx context.Context) GetDefaultServiceAccountResultOutput {
	return o
}

// The display name for the service account.
func (o GetDefaultServiceAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Email address of the default service account used by VMs running in this project
func (o GetDefaultServiceAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDefaultServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The fully-qualified name of the service account.
func (o GetDefaultServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDefaultServiceAccountResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.Project }).(pulumi.StringOutput)
}

// The unique id of the service account.
func (o GetDefaultServiceAccountResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDefaultServiceAccountResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDefaultServiceAccountResultOutput{})
}
