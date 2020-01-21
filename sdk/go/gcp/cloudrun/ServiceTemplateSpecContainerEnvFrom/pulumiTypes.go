// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ServiceTemplateSpecContainerEnvFrom

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference"
)

type ServiceTemplateSpecContainerEnvFrom struct {
	ConfigMapRef *cloudrunServiceTemplateSpecContainerEnvFromConfigMapRef.ServiceTemplateSpecContainerEnvFromConfigMapRef `pulumi:"configMapRef"`
	Prefix *string `pulumi:"prefix"`
	SecretRef *cloudrunServiceTemplateSpecContainerEnvFromSecretRef.ServiceTemplateSpecContainerEnvFromSecretRef `pulumi:"secretRef"`
}

type ServiceTemplateSpecContainerEnvFromInput interface {
	pulumi.Input

	ToServiceTemplateSpecContainerEnvFromOutput() ServiceTemplateSpecContainerEnvFromOutput
	ToServiceTemplateSpecContainerEnvFromOutputWithContext(context.Context) ServiceTemplateSpecContainerEnvFromOutput
}

type ServiceTemplateSpecContainerEnvFromArgs struct {
	ConfigMapRef cloudrunServiceTemplateSpecContainerEnvFromConfigMapRef.ServiceTemplateSpecContainerEnvFromConfigMapRefPtrInput `pulumi:"configMapRef"`
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	SecretRef cloudrunServiceTemplateSpecContainerEnvFromSecretRef.ServiceTemplateSpecContainerEnvFromSecretRefPtrInput `pulumi:"secretRef"`
}

func (ServiceTemplateSpecContainerEnvFromArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpecContainerEnvFrom)(nil)).Elem()
}

func (i ServiceTemplateSpecContainerEnvFromArgs) ToServiceTemplateSpecContainerEnvFromOutput() ServiceTemplateSpecContainerEnvFromOutput {
	return i.ToServiceTemplateSpecContainerEnvFromOutputWithContext(context.Background())
}

func (i ServiceTemplateSpecContainerEnvFromArgs) ToServiceTemplateSpecContainerEnvFromOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerEnvFromOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateSpecContainerEnvFromOutput)
}

type ServiceTemplateSpecContainerEnvFromArrayInput interface {
	pulumi.Input

	ToServiceTemplateSpecContainerEnvFromArrayOutput() ServiceTemplateSpecContainerEnvFromArrayOutput
	ToServiceTemplateSpecContainerEnvFromArrayOutputWithContext(context.Context) ServiceTemplateSpecContainerEnvFromArrayOutput
}

type ServiceTemplateSpecContainerEnvFromArray []ServiceTemplateSpecContainerEnvFromInput

func (ServiceTemplateSpecContainerEnvFromArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTemplateSpecContainerEnvFrom)(nil)).Elem()
}

func (i ServiceTemplateSpecContainerEnvFromArray) ToServiceTemplateSpecContainerEnvFromArrayOutput() ServiceTemplateSpecContainerEnvFromArrayOutput {
	return i.ToServiceTemplateSpecContainerEnvFromArrayOutputWithContext(context.Background())
}

func (i ServiceTemplateSpecContainerEnvFromArray) ToServiceTemplateSpecContainerEnvFromArrayOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerEnvFromArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateSpecContainerEnvFromArrayOutput)
}

type ServiceTemplateSpecContainerEnvFromOutput struct { *pulumi.OutputState }

func (ServiceTemplateSpecContainerEnvFromOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplateSpecContainerEnvFrom)(nil)).Elem()
}

func (o ServiceTemplateSpecContainerEnvFromOutput) ToServiceTemplateSpecContainerEnvFromOutput() ServiceTemplateSpecContainerEnvFromOutput {
	return o
}

func (o ServiceTemplateSpecContainerEnvFromOutput) ToServiceTemplateSpecContainerEnvFromOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerEnvFromOutput {
	return o
}

func (o ServiceTemplateSpecContainerEnvFromOutput) ConfigMapRef() cloudrunServiceTemplateSpecContainerEnvFromConfigMapRef.ServiceTemplateSpecContainerEnvFromConfigMapRefPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainerEnvFrom) *cloudrunServiceTemplateSpecContainerEnvFromConfigMapRef.ServiceTemplateSpecContainerEnvFromConfigMapRef { return v.ConfigMapRef }).(cloudrunServiceTemplateSpecContainerEnvFromConfigMapRef.ServiceTemplateSpecContainerEnvFromConfigMapRefPtrOutput)
}

func (o ServiceTemplateSpecContainerEnvFromOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainerEnvFrom) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o ServiceTemplateSpecContainerEnvFromOutput) SecretRef() cloudrunServiceTemplateSpecContainerEnvFromSecretRef.ServiceTemplateSpecContainerEnvFromSecretRefPtrOutput {
	return o.ApplyT(func (v ServiceTemplateSpecContainerEnvFrom) *cloudrunServiceTemplateSpecContainerEnvFromSecretRef.ServiceTemplateSpecContainerEnvFromSecretRef { return v.SecretRef }).(cloudrunServiceTemplateSpecContainerEnvFromSecretRef.ServiceTemplateSpecContainerEnvFromSecretRefPtrOutput)
}

type ServiceTemplateSpecContainerEnvFromArrayOutput struct { *pulumi.OutputState}

func (ServiceTemplateSpecContainerEnvFromArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceTemplateSpecContainerEnvFrom)(nil)).Elem()
}

func (o ServiceTemplateSpecContainerEnvFromArrayOutput) ToServiceTemplateSpecContainerEnvFromArrayOutput() ServiceTemplateSpecContainerEnvFromArrayOutput {
	return o
}

func (o ServiceTemplateSpecContainerEnvFromArrayOutput) ToServiceTemplateSpecContainerEnvFromArrayOutputWithContext(ctx context.Context) ServiceTemplateSpecContainerEnvFromArrayOutput {
	return o
}

func (o ServiceTemplateSpecContainerEnvFromArrayOutput) Index(i pulumi.IntInput) ServiceTemplateSpecContainerEnvFromOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceTemplateSpecContainerEnvFrom {
		return vs[0].([]ServiceTemplateSpecContainerEnvFrom)[vs[1].(int)]
	}).(ServiceTemplateSpecContainerEnvFromOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTemplateSpecContainerEnvFromOutput{})
	pulumi.RegisterOutputType(ServiceTemplateSpecContainerEnvFromArrayOutput{})
}
