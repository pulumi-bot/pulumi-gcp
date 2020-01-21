// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ServiceTemplate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateMetadata"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpec"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainer"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnv"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFrom"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRef"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference"
	"https:/github.com/pulumi/pulumi-gcp/cloudrun/ServiceTemplateSpecContainerResources"
)

type ServiceTemplate struct {
	Metadata *cloudrunServiceTemplateMetadata.ServiceTemplateMetadata `pulumi:"metadata"`
	Spec cloudrunServiceTemplateSpec.ServiceTemplateSpec `pulumi:"spec"`
}

type ServiceTemplateInput interface {
	pulumi.Input

	ToServiceTemplateOutput() ServiceTemplateOutput
	ToServiceTemplateOutputWithContext(context.Context) ServiceTemplateOutput
}

type ServiceTemplateArgs struct {
	Metadata cloudrunServiceTemplateMetadata.ServiceTemplateMetadataPtrInput `pulumi:"metadata"`
	Spec cloudrunServiceTemplateSpec.ServiceTemplateSpecInput `pulumi:"spec"`
}

func (ServiceTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplate)(nil)).Elem()
}

func (i ServiceTemplateArgs) ToServiceTemplateOutput() ServiceTemplateOutput {
	return i.ToServiceTemplateOutputWithContext(context.Background())
}

func (i ServiceTemplateArgs) ToServiceTemplateOutputWithContext(ctx context.Context) ServiceTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateOutput)
}

func (i ServiceTemplateArgs) ToServiceTemplatePtrOutput() ServiceTemplatePtrOutput {
	return i.ToServiceTemplatePtrOutputWithContext(context.Background())
}

func (i ServiceTemplateArgs) ToServiceTemplatePtrOutputWithContext(ctx context.Context) ServiceTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplateOutput).ToServiceTemplatePtrOutputWithContext(ctx)
}

type ServiceTemplatePtrInput interface {
	pulumi.Input

	ToServiceTemplatePtrOutput() ServiceTemplatePtrOutput
	ToServiceTemplatePtrOutputWithContext(context.Context) ServiceTemplatePtrOutput
}

type serviceTemplatePtrType ServiceTemplateArgs

func ServiceTemplatePtr(v *ServiceTemplateArgs) ServiceTemplatePtrInput {	return (*serviceTemplatePtrType)(v)
}

func (*serviceTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTemplate)(nil)).Elem()
}

func (i *serviceTemplatePtrType) ToServiceTemplatePtrOutput() ServiceTemplatePtrOutput {
	return i.ToServiceTemplatePtrOutputWithContext(context.Background())
}

func (i *serviceTemplatePtrType) ToServiceTemplatePtrOutputWithContext(ctx context.Context) ServiceTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTemplatePtrOutput)
}

type ServiceTemplateOutput struct { *pulumi.OutputState }

func (ServiceTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceTemplate)(nil)).Elem()
}

func (o ServiceTemplateOutput) ToServiceTemplateOutput() ServiceTemplateOutput {
	return o
}

func (o ServiceTemplateOutput) ToServiceTemplateOutputWithContext(ctx context.Context) ServiceTemplateOutput {
	return o
}

func (o ServiceTemplateOutput) ToServiceTemplatePtrOutput() ServiceTemplatePtrOutput {
	return o.ToServiceTemplatePtrOutputWithContext(context.Background())
}

func (o ServiceTemplateOutput) ToServiceTemplatePtrOutputWithContext(ctx context.Context) ServiceTemplatePtrOutput {
	return o.ApplyT(func(v ServiceTemplate) *ServiceTemplate {
		return &v
	}).(ServiceTemplatePtrOutput)
}
func (o ServiceTemplateOutput) Metadata() cloudrunServiceTemplateMetadata.ServiceTemplateMetadataPtrOutput {
	return o.ApplyT(func (v ServiceTemplate) *cloudrunServiceTemplateMetadata.ServiceTemplateMetadata { return v.Metadata }).(cloudrunServiceTemplateMetadata.ServiceTemplateMetadataPtrOutput)
}

func (o ServiceTemplateOutput) Spec() cloudrunServiceTemplateSpec.ServiceTemplateSpecOutput {
	return o.ApplyT(func (v ServiceTemplate) cloudrunServiceTemplateSpec.ServiceTemplateSpec { return v.Spec }).(cloudrunServiceTemplateSpec.ServiceTemplateSpecOutput)
}

type ServiceTemplatePtrOutput struct { *pulumi.OutputState}

func (ServiceTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTemplate)(nil)).Elem()
}

func (o ServiceTemplatePtrOutput) ToServiceTemplatePtrOutput() ServiceTemplatePtrOutput {
	return o
}

func (o ServiceTemplatePtrOutput) ToServiceTemplatePtrOutputWithContext(ctx context.Context) ServiceTemplatePtrOutput {
	return o
}

func (o ServiceTemplatePtrOutput) Elem() ServiceTemplateOutput {
	return o.ApplyT(func (v *ServiceTemplate) ServiceTemplate { return *v }).(ServiceTemplateOutput)
}

func (o ServiceTemplatePtrOutput) Metadata() cloudrunServiceTemplateMetadata.ServiceTemplateMetadataPtrOutput {
	return o.ApplyT(func (v ServiceTemplate) *cloudrunServiceTemplateMetadata.ServiceTemplateMetadata { return v.Metadata }).(cloudrunServiceTemplateMetadata.ServiceTemplateMetadataPtrOutput)
}

func (o ServiceTemplatePtrOutput) Spec() cloudrunServiceTemplateSpec.ServiceTemplateSpecOutput {
	return o.ApplyT(func (v ServiceTemplate) cloudrunServiceTemplateSpec.ServiceTemplateSpec { return v.Spec }).(cloudrunServiceTemplateSpec.ServiceTemplateSpecOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTemplateOutput{})
	pulumi.RegisterOutputType(ServiceTemplatePtrOutput{})
}
