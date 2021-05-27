// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an active folder within GCP by `displayName` and `parent`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.GetActiveFolder(ctx, &organizations.GetActiveFolderArgs{
// 			DisplayName: "Department 1",
// 			Parent:      "organizations/1234567",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetActiveFolder(ctx *pulumi.Context, args *GetActiveFolderArgs, opts ...pulumi.InvokeOption) (*GetActiveFolderResult, error) {
	var rv GetActiveFolderResult
	err := ctx.Invoke("gcp:organizations/getActiveFolder:getActiveFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActiveFolder.
type GetActiveFolderArgs struct {
	// The folder's display name.
	DisplayName string `pulumi:"displayName"`
	// The resource name of the parent Folder or Organization.
	Parent string `pulumi:"parent"`
}

// A collection of values returned by getActiveFolder.
type GetActiveFolderResult struct {
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The resource name of the Folder. This uniquely identifies the folder.
	Name   string `pulumi:"name"`
	Parent string `pulumi:"parent"`
}

func GetActiveFolderOutput(ctx *pulumi.Context, args GetActiveFolderOutputArgs, opts ...pulumi.InvokeOption) GetActiveFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActiveFolderResult, error) {
			args := v.(GetActiveFolderArgs)
			r, err := GetActiveFolder(ctx, &args, opts...)
			return *r, err
		}).(GetActiveFolderResultOutput)
}

// A collection of arguments for invoking getActiveFolder.
type GetActiveFolderOutputArgs struct {
	// The folder's display name.
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// The resource name of the parent Folder or Organization.
	Parent pulumi.StringInput `pulumi:"parent"`
}

func (GetActiveFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveFolderArgs)(nil)).Elem()
}

// A collection of values returned by getActiveFolder.
type GetActiveFolderResultOutput struct{ *pulumi.OutputState }

func (GetActiveFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveFolderResult)(nil)).Elem()
}

func (o GetActiveFolderResultOutput) ToGetActiveFolderResultOutput() GetActiveFolderResultOutput {
	return o
}

func (o GetActiveFolderResultOutput) ToGetActiveFolderResultOutputWithContext(ctx context.Context) GetActiveFolderResultOutput {
	return o
}

func (o GetActiveFolderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveFolderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetActiveFolderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveFolderResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name of the Folder. This uniquely identifies the folder.
func (o GetActiveFolderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveFolderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetActiveFolderResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveFolderResult) string { return v.Parent }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActiveFolderResultOutput{})
}
