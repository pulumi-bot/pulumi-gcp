// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Secret Manager secret's version. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets.versions).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/secretmanager"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := secretmanager.LookupSecretVersion(ctx, &secretmanager.LookupSecretVersionArgs{
// 			Secret: "my-secret",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSecretVersion(ctx *pulumi.Context, args *LookupSecretVersionArgs, opts ...pulumi.InvokeOption) (*LookupSecretVersionResult, error) {
	var rv LookupSecretVersionResult
	err := ctx.Invoke("gcp:secretmanager/getSecretVersion:getSecretVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionArgs struct {
	// The project to get the secret version for. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The secret to get the secret version for.
	Secret string `pulumi:"secret"`
	// The version of the secret to get. If it
	// is not provided, the latest version is retrieved.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getSecretVersion.
type LookupSecretVersionResult struct {
	// The time at which the Secret was created.
	CreateTime string `pulumi:"createTime"`
	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime string `pulumi:"destroyTime"`
	// True if the current state of the SecretVersion is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The resource name of the SecretVersion. Format:
	// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	Secret  string `pulumi:"secret"`
	// The secret data. No larger than 64KiB.
	SecretData string `pulumi:"secretData"`
	Version    string `pulumi:"version"`
}

func LookupSecretVersionApply(ctx *pulumi.Context, args LookupSecretVersionApplyInput, opts ...pulumi.InvokeOption) LookupSecretVersionResultOutput {
	return args.ToLookupSecretVersionApplyOutput().ApplyT(func(v LookupSecretVersionArgs) (LookupSecretVersionResult, error) {
		r, err := LookupSecretVersion(ctx, &v, opts...)
		return *r, err

	}).(LookupSecretVersionResultOutput)
}

// LookupSecretVersionApplyInput is an input type that accepts LookupSecretVersionApplyArgs and LookupSecretVersionApplyOutput values.
// You can construct a concrete instance of `LookupSecretVersionApplyInput` via:
//
//          LookupSecretVersionApplyArgs{...}
type LookupSecretVersionApplyInput interface {
	pulumi.Input

	ToLookupSecretVersionApplyOutput() LookupSecretVersionApplyOutput
	ToLookupSecretVersionApplyOutputWithContext(context.Context) LookupSecretVersionApplyOutput
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionApplyArgs struct {
	// The project to get the secret version for. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput `pulumi:"project"`
	// The secret to get the secret version for.
	Secret pulumi.StringInput `pulumi:"secret"`
	// The version of the secret to get. If it
	// is not provided, the latest version is retrieved.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (LookupSecretVersionApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretVersionArgs)(nil)).Elem()
}

func (i LookupSecretVersionApplyArgs) ToLookupSecretVersionApplyOutput() LookupSecretVersionApplyOutput {
	return i.ToLookupSecretVersionApplyOutputWithContext(context.Background())
}

func (i LookupSecretVersionApplyArgs) ToLookupSecretVersionApplyOutputWithContext(ctx context.Context) LookupSecretVersionApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupSecretVersionApplyOutput)
}

// A collection of arguments for invoking getSecretVersion.
type LookupSecretVersionApplyOutput struct{ *pulumi.OutputState }

func (LookupSecretVersionApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretVersionArgs)(nil)).Elem()
}

func (o LookupSecretVersionApplyOutput) ToLookupSecretVersionApplyOutput() LookupSecretVersionApplyOutput {
	return o
}

func (o LookupSecretVersionApplyOutput) ToLookupSecretVersionApplyOutputWithContext(ctx context.Context) LookupSecretVersionApplyOutput {
	return o
}

// The project to get the secret version for. If it
// is not provided, the provider project is used.
func (o LookupSecretVersionApplyOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionArgs) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The secret to get the secret version for.
func (o LookupSecretVersionApplyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionArgs) string { return v.Secret }).(pulumi.StringOutput)
}

// The version of the secret to get. If it
// is not provided, the latest version is retrieved.
func (o LookupSecretVersionApplyOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretVersionArgs) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getSecretVersion.
type LookupSecretVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSecretVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretVersionResult)(nil)).Elem()
}

func (o LookupSecretVersionResultOutput) ToLookupSecretVersionResultOutput() LookupSecretVersionResultOutput {
	return o
}

func (o LookupSecretVersionResultOutput) ToLookupSecretVersionResultOutputWithContext(ctx context.Context) LookupSecretVersionResultOutput {
	return o
}

// The time at which the Secret was created.
func (o LookupSecretVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which the Secret was destroyed. Only present if state is DESTROYED.
func (o LookupSecretVersionResultOutput) DestroyTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.DestroyTime }).(pulumi.StringOutput)
}

// True if the current state of the SecretVersion is enabled.
func (o LookupSecretVersionResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name of the SecretVersion. Format:
// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
func (o LookupSecretVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecretVersionResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupSecretVersionResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Secret }).(pulumi.StringOutput)
}

// The secret data. No larger than 64KiB.
func (o LookupSecretVersionResultOutput) SecretData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.SecretData }).(pulumi.StringOutput)
}

func (o LookupSecretVersionResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretVersionResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretVersionApplyOutput{})
	pulumi.RegisterOutputType(LookupSecretVersionResultOutput{})
}
