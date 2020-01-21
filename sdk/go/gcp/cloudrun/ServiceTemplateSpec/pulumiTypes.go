// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ServiceTemplateSpec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainer"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnv"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFrom"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerResources"
)

type ServiceTemplateSpec struct {
	ContainerConcurrency *int `pulumi:"containerConcurrency"`
	Containers []cloudrunServiceTemplateSpecContainer.ServiceTemplateSpecContainer `pulumi:"containers"`
	ServiceAccountName *string `pulumi:"serviceAccountName"`
	ServingState *string `pulumi:"servingState"`
}

type ServiceTemplateSpecInput interface {
	pulumi.Input

	ToServiceTemplateSpecOutput() ServiceTemplateSpecOutput
	ToServiceTemplateSpecOutputWithContext(context.Context) ServiceTemplateSpecOutput
}

type ServiceTemplateSpecArgs struct {
	ContainerConcurrency pulumi.IntPtrInput `pulumi:"containerConcurrency"`
	Containers cloudrunServiceTemplateSpecContainer.ServiceTemplateSpecContainerArrayInput `pulumi:"containers"`
	ServiceAccountName pulumi.StringPtrInput `pulumi:"serviceAccountName"`
	ServingState pulumi.StringPtrInput `pulumi:"servingState"`
}

func (ServiceTemplateSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpec)(nil)).Elem()
}

func (i ServiceTemplateSpecArgs) ToServiceTemplateSpecOutput() ServiceTemplateSpecOutput {
	return i.ToServiceTemplateSpecOutputWithContext(context.Background())
}

func (i ServiceTemplateSpecArgs) ToServiceTemplateSpecOutputWithContext(ctx context.Context) ServiceTemplateSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateSpecOutput)
}

type ServiceTemplateSpecOutput struct { *pulumi.OutputState }

func (ServiceTemplateSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpec)(nil)).Elem()
}

func (o ServiceTemplateSpecOutput) ToServiceTemplateSpecOutput() ServiceTemplateSpecOutput {
	return o
}

func (o ServiceTemplateSpecOutput) ToServiceTemplateSpecOutputWithContext(ctx context.Context) ServiceTemplateSpecOutput {
	return o
}

func (o ServiceTemplateSpecOutput) ContainerConcurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpec) *int { return v.ContainerConcurrency }).(pulumi.IntPtrOutput)
}

func (o ServiceTemplateSpecOutput) Containers() cloudrunServiceTemplateSpecContainer.ServiceTemplateSpecContainerArrayOutput {
	return o.ApplyT(func (v ServiceTemplateSpec) []cloudrunServiceTemplateSpecContainer.ServiceTemplateSpecContainer { return v.Containers }).(cloudrunServiceTemplateSpecContainer.ServiceTemplateSpecContainerArrayOutput)
}

func (o ServiceTemplateSpecOutput) ServiceAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpec) *string { return v.ServiceAccountName }).(pulumi.StringPtrOutput)
}

func (o ServiceTemplateSpecOutput) ServingState() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpec) *string { return v.ServingState }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTemplateSpecOutput{})
}
