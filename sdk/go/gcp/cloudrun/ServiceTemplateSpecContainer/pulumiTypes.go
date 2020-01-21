// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ServiceTemplateSpecContainer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnv"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFrom"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerResources"
)

type ServiceTemplateSpecContainer struct {
	Args []string `pulumi:"args"`
	Commands []string `pulumi:"commands"`
	EnvFroms []cloudrunServiceTemplateSpecContainerEnvFrom.ServiceTemplateSpecContainerEnvFrom `pulumi:"envFroms"`
	Envs []cloudrunServiceTemplateSpecContainerEnv.ServiceTemplateSpecContainerEnv `pulumi:"envs"`
	Image string `pulumi:"image"`
	Resources *cloudrunServiceTemplateSpecContainerResources.ServiceTemplateSpecContainerResources `pulumi:"resources"`
	WorkingDir *string `pulumi:"workingDir"`
}

type ServiceTemplateSpecContainerInput interface {
	pulumi.Input

	ToServiceTemplateSpecContainerOutput() ServiceTemplateSpecContainerOutput
	ToServiceTemplateSpecContainerOutputWithContext(context.Context) ServiceTemplateSpecContainerOutput
}

type ServiceTemplateSpecContainerArgs struct {
	Args pulumi.StringArrayInput `pulumi:"args"`
	Commands pulumi.StringArrayInput `pulumi:"commands"`
	EnvFroms cloudrunServiceTemplateSpecContainerEnvFrom.ServiceTemplateSpecContainerEnvFromArrayInput `pulumi:"envFroms"`
	Envs cloudrunServiceTemplateSpecContainerEnv.ServiceTemplateSpecContainerEnvArrayInput `pulumi:"envs"`
	Image pulumi.StringInput `pulumi:"image"`
	Resources cloudrunServiceTemplateSpecContainerResources.ServiceTemplateSpecContainerResourcesPtrInput `pulumi:"resources"`
	WorkingDir pulumi.StringPtrInput `pulumi:"workingDir"`
}

func (ServiceTemplateSpecContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpecContainer)(nil)).Elem()
}

func (i ServiceTemplateSpecContainerArgs) ToServiceTemplateSpecContainerOutput() ServiceTemplateSpecContainerOutput {
	return i.ToServiceTemplateSpecContainerOutputWithContext(context.Background())
}

func (i ServiceTemplateSpecContainerArgs) ToServiceTemplateSpecContainerOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateSpecContainerOutput)
}

type ServiceTemplateSpecContainerArrayInput interface {
	pulumi.Input

	ToServiceTemplateSpecContainerArrayOutput() ServiceTemplateSpecContainerArrayOutput
	ToServiceTemplateSpecContainerArrayOutputWithContext(context.Context) ServiceTemplateSpecContainerArrayOutput
}

type ServiceTemplateSpecContainerArray []ServiceTemplateSpecContainerInput

func (ServiceTemplateSpecContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTemplateSpecContainer)(nil)).Elem()
}

func (i ServiceTemplateSpecContainerArray) ToServiceTemplateSpecContainerArrayOutput() ServiceTemplateSpecContainerArrayOutput {
	return i.ToServiceTemplateSpecContainerArrayOutputWithContext(context.Background())
}

func (i ServiceTemplateSpecContainerArray) ToServiceTemplateSpecContainerArrayOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateSpecContainerArrayOutput)
}

type ServiceTemplateSpecContainerOutput struct { *pulumi.OutputState }

func (ServiceTemplateSpecContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpecContainer)(nil)).Elem()
}

func (o ServiceTemplateSpecContainerOutput) ToServiceTemplateSpecContainerOutput() ServiceTemplateSpecContainerOutput {
	return o
}

func (o ServiceTemplateSpecContainerOutput) ToServiceTemplateSpecContainerOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerOutput {
	return o
}

func (o ServiceTemplateSpecContainerOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o ServiceTemplateSpecContainerOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o ServiceTemplateSpecContainerOutput) EnvFroms() cloudrunServiceTemplateSpecContainerEnvFrom.ServiceTemplateSpecContainerEnvFromArrayOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) []cloudrunServiceTemplateSpecContainerEnvFrom.ServiceTemplateSpecContainerEnvFrom { return v.EnvFroms }).(cloudrunServiceTemplateSpecContainerEnvFrom.ServiceTemplateSpecContainerEnvFromArrayOutput)
}

func (o ServiceTemplateSpecContainerOutput) Envs() cloudrunServiceTemplateSpecContainerEnv.ServiceTemplateSpecContainerEnvArrayOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) []cloudrunServiceTemplateSpecContainerEnv.ServiceTemplateSpecContainerEnv { return v.Envs }).(cloudrunServiceTemplateSpecContainerEnv.ServiceTemplateSpecContainerEnvArrayOutput)
}

func (o ServiceTemplateSpecContainerOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) string { return v.Image }).(pulumi.StringOutput)
}

func (o ServiceTemplateSpecContainerOutput) Resources() cloudrunServiceTemplateSpecContainerResources.ServiceTemplateSpecContainerResourcesPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) *cloudrunServiceTemplateSpecContainerResources.ServiceTemplateSpecContainerResources { return v.Resources }).(cloudrunServiceTemplateSpecContainerResources.ServiceTemplateSpecContainerResourcesPtrOutput)
}

func (o ServiceTemplateSpecContainerOutput) WorkingDir() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainer) *string { return v.WorkingDir }).(pulumi.StringPtrOutput)
}

type ServiceTemplateSpecContainerArrayOutput struct { *pulumi.OutputState}

func (ServiceTemplateSpecContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTemplateSpecContainer)(nil)).Elem()
}

func (o ServiceTemplateSpecContainerArrayOutput) ToServiceTemplateSpecContainerArrayOutput() ServiceTemplateSpecContainerArrayOutput {
	return o
}

func (o ServiceTemplateSpecContainerArrayOutput) ToServiceTemplateSpecContainerArrayOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerArrayOutput {
	return o
}

func (o ServiceTemplateSpecContainerArrayOutput) Index(i pulumi.IntInput) ServiceTemplateSpecContainerOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceTemplateSpecContainer {
		return vs[0].([]ServiceTemplateSpecContainer)[vs[1].(int)]
	}).(ServiceTemplateSpecContainerOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTemplateSpecContainerOutput{})
	pulumi.RegisterOutputType(ServiceTemplateSpecContainerArrayOutput{})
}
