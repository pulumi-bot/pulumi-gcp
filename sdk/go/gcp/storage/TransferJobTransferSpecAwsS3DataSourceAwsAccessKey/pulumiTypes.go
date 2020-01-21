// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TransferJobTransferSpecAwsS3DataSourceAwsAccessKey

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TransferJobTransferSpecAwsS3DataSourceAwsAccessKey struct {
	AccessKeyId string `pulumi:"accessKeyId"`
	SecretAccessKey string `pulumi:"secretAccessKey"`
}

type TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyInput interface {
	pulumi.Input

	ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput() TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput
	ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputWithContext(context.Context) TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput
}

type TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs struct {
	AccessKeyId pulumi.StringInput `pulumi:"accessKeyId"`
	SecretAccessKey pulumi.StringInput `pulumi:"secretAccessKey"`
}

func (TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJobTransferSpecAwsS3DataSourceAwsAccessKey)(nil)).Elem()
}

func (i TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs) ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput() TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput {
	return i.ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputWithContext(context.Background())
}

func (i TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs) ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputWithContext(ctx context.Context) TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput)
}

type TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput struct { *pulumi.OutputState }

func (TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJobTransferSpecAwsS3DataSourceAwsAccessKey)(nil)).Elem()
}

func (o TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput) ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput() TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput {
	return o
}

func (o TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput) ToTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputWithContext(ctx context.Context) TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput {
	return o
}

func (o TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput) AccessKeyId() pulumi.StringOutput {
	return o.ApplyT(func (v TransferJobTransferSpecAwsS3DataSourceAwsAccessKey) string { return v.AccessKeyId }).(pulumi.StringOutput)
}

func (o TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput) SecretAccessKey() pulumi.StringOutput {
	return o.ApplyT(func (v TransferJobTransferSpecAwsS3DataSourceAwsAccessKey) string { return v.SecretAccessKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutput{})
}
