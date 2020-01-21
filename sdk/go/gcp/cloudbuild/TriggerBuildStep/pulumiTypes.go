// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package TriggerBuildStep

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudbuild/TriggerBuildStepVolume"
)

type TriggerBuildStep struct {
	Args []string `pulumi:"args"`
	Dir *string `pulumi:"dir"`
	Entrypoint *string `pulumi:"entrypoint"`
	Envs []string `pulumi:"envs"`
	Id *string `pulumi:"id"`
	Name string `pulumi:"name"`
	SecretEnvs []string `pulumi:"secretEnvs"`
	Timeout *string `pulumi:"timeout"`
	Timing *string `pulumi:"timing"`
	Volumes []cloudbuildTriggerBuildStepVolume.TriggerBuildStepVolume `pulumi:"volumes"`
	WaitFors []string `pulumi:"waitFors"`
}

type TriggerBuildStepInput interface {
	pulumi.Input

	ToTriggerBuildStepOutput() TriggerBuildStepOutput
	ToTriggerBuildStepOutputWithContext(context.Context) TriggerBuildStepOutput
}

type TriggerBuildStepArgs struct {
	Args pulumi.StringArrayInput `pulumi:"args"`
	Dir pulumi.StringPtrInput `pulumi:"dir"`
	Entrypoint pulumi.StringPtrInput `pulumi:"entrypoint"`
	Envs pulumi.StringArrayInput `pulumi:"envs"`
	Id pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringInput `pulumi:"name"`
	SecretEnvs pulumi.StringArrayInput `pulumi:"secretEnvs"`
	Timeout pulumi.StringPtrInput `pulumi:"timeout"`
	Timing pulumi.StringPtrInput `pulumi:"timing"`
	Volumes cloudbuildTriggerBuildStepVolume.TriggerBuildStepVolumeArrayInput `pulumi:"volumes"`
	WaitFors pulumi.StringArrayInput `pulumi:"waitFors"`
}

func (TriggerBuildStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerBuildStep)(nil)).Elem()
}

func (i TriggerBuildStepArgs) ToTriggerBuildStepOutput() TriggerBuildStepOutput {
	return i.ToTriggerBuildStepOutputWithContext(context.Background())
}

func (i TriggerBuildStepArgs) ToTriggerBuildStepOutputWithContext(ctx context.Context) TriggerBuildStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerBuildStepOutput)
}

type TriggerBuildStepArrayInput interface {
	pulumi.Input

	ToTriggerBuildStepArrayOutput() TriggerBuildStepArrayOutput
	ToTriggerBuildStepArrayOutputWithContext(context.Context) TriggerBuildStepArrayOutput
}

type TriggerBuildStepArray []TriggerBuildStepInput

func (TriggerBuildStepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerBuildStep)(nil)).Elem()
}

func (i TriggerBuildStepArray) ToTriggerBuildStepArrayOutput() TriggerBuildStepArrayOutput {
	return i.ToTriggerBuildStepArrayOutputWithContext(context.Background())
}

func (i TriggerBuildStepArray) ToTriggerBuildStepArrayOutputWithContext(ctx context.Context) TriggerBuildStepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerBuildStepArrayOutput)
}

type TriggerBuildStepOutput struct { *pulumi.OutputState }

func (TriggerBuildStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TriggerBuildStep)(nil)).Elem()
}

func (o TriggerBuildStepOutput) ToTriggerBuildStepOutput() TriggerBuildStepOutput {
	return o
}

func (o TriggerBuildStepOutput) ToTriggerBuildStepOutputWithContext(ctx context.Context) TriggerBuildStepOutput {
	return o
}

func (o TriggerBuildStepOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func (v TriggerBuildStep) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o TriggerBuildStepOutput) Dir() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerBuildStep) *string { return v.Dir }).(pulumi.StringPtrOutput)
}

func (o TriggerBuildStepOutput) Entrypoint() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerBuildStep) *string { return v.Entrypoint }).(pulumi.StringPtrOutput)
}

func (o TriggerBuildStepOutput) Envs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v TriggerBuildStep) []string { return v.Envs }).(pulumi.StringArrayOutput)
}

func (o TriggerBuildStepOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerBuildStep) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o TriggerBuildStepOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v TriggerBuildStep) string { return v.Name }).(pulumi.StringOutput)
}

func (o TriggerBuildStepOutput) SecretEnvs() pulumi.StringArrayOutput {
	return o.ApplyT(func (v TriggerBuildStep) []string { return v.SecretEnvs }).(pulumi.StringArrayOutput)
}

func (o TriggerBuildStepOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerBuildStep) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

func (o TriggerBuildStepOutput) Timing() pulumi.StringPtrOutput {
	return o.ApplyT(func (v TriggerBuildStep) *string { return v.Timing }).(pulumi.StringPtrOutput)
}

func (o TriggerBuildStepOutput) Volumes() cloudbuildTriggerBuildStepVolume.TriggerBuildStepVolumeArrayOutput {
	return o.ApplyT(func (v TriggerBuildStep) []cloudbuildTriggerBuildStepVolume.TriggerBuildStepVolume { return v.Volumes }).(cloudbuildTriggerBuildStepVolume.TriggerBuildStepVolumeArrayOutput)
}

func (o TriggerBuildStepOutput) WaitFors() pulumi.StringArrayOutput {
	return o.ApplyT(func (v TriggerBuildStep) []string { return v.WaitFors }).(pulumi.StringArrayOutput)
}

type TriggerBuildStepArrayOutput struct { *pulumi.OutputState}

func (TriggerBuildStepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TriggerBuildStep)(nil)).Elem()
}

func (o TriggerBuildStepArrayOutput) ToTriggerBuildStepArrayOutput() TriggerBuildStepArrayOutput {
	return o
}

func (o TriggerBuildStepArrayOutput) ToTriggerBuildStepArrayOutputWithContext(ctx context.Context) TriggerBuildStepArrayOutput {
	return o
}

func (o TriggerBuildStepArrayOutput) Index(i pulumi.IntInput) TriggerBuildStepOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) TriggerBuildStep {
		return vs[0].([]TriggerBuildStep)[vs[1].(int)]
	}).(TriggerBuildStepOutput)
}

func init() {
	pulumi.RegisterOutputType(TriggerBuildStepOutput{})
	pulumi.RegisterOutputType(TriggerBuildStepArrayOutput{})
}
