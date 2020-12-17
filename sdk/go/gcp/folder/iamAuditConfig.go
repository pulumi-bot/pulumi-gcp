// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type IamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Folder pulumi.StringOutput `pulumi:"folder"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIamAuditConfig(ctx *pulumi.Context,
	name string, args *IamAuditConfigArgs, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IamAuditConfig
	err := ctx.RegisterResource("gcp:folder/iamAuditConfig:IamAuditConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamAuditConfig gets an existing IamAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamAuditConfigState, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	var resource IamAuditConfig
	err := ctx.ReadResource("gcp:folder/iamAuditConfig:IamAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamAuditConfig resources.
type iamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag   *string `pulumi:"etag"`
	Folder *string `pulumi:"folder"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service *string `pulumi:"service"`
}

type IamAuditConfigState struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The etag of iam policy
	Etag   pulumi.StringPtrInput
	Folder pulumi.StringPtrInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringPtrInput
}

func (IamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigState)(nil)).Elem()
}

type iamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	Folder          string                         `pulumi:"folder"`
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission. This can be specified multiple times.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	Folder          pulumi.StringInput
	// Service which will be enabled for audit logging. The special value allServices covers all services.
	Service pulumi.StringInput
}

func (IamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigArgs)(nil)).Elem()
}

type IamAuditConfigInput interface {
	pulumi.Input

	ToIamAuditConfigOutput() IamAuditConfigOutput
	ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput
}

func (*IamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*IamAuditConfig)(nil))
}

func (i *IamAuditConfig) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return i.ToIamAuditConfigOutputWithContext(context.Background())
}

func (i *IamAuditConfig) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigOutput)
}

func (i *IamAuditConfig) ToIamAuditConfigPtrOutput() IamAuditConfigPtrOutput {
	return i.ToIamAuditConfigPtrOutputWithContext(context.Background())
}

func (i *IamAuditConfig) ToIamAuditConfigPtrOutputWithContext(ctx context.Context) IamAuditConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigPtrOutput)
}

type IamAuditConfigPtrInput interface {
	pulumi.Input

	ToIamAuditConfigPtrOutput() IamAuditConfigPtrOutput
	ToIamAuditConfigPtrOutputWithContext(ctx context.Context) IamAuditConfigPtrOutput
}

type iamAuditConfigPtrType IamAuditConfigArgs

func (*iamAuditConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil))
}

func (i *iamAuditConfigPtrType) ToIamAuditConfigPtrOutput() IamAuditConfigPtrOutput {
	return i.ToIamAuditConfigPtrOutputWithContext(context.Background())
}

func (i *iamAuditConfigPtrType) ToIamAuditConfigPtrOutputWithContext(ctx context.Context) IamAuditConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigOutput).ToIamAuditConfigPtrOutput()
}

// IamAuditConfigArrayInput is an input type that accepts IamAuditConfigArray and IamAuditConfigArrayOutput values.
// You can construct a concrete instance of `IamAuditConfigArrayInput` via:
//
//          IamAuditConfigArray{ IamAuditConfigArgs{...} }
type IamAuditConfigArrayInput interface {
	pulumi.Input

	ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput
	ToIamAuditConfigArrayOutputWithContext(context.Context) IamAuditConfigArrayOutput
}

type IamAuditConfigArray []IamAuditConfigInput

func (IamAuditConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamAuditConfig)(nil))
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return i.ToIamAuditConfigArrayOutputWithContext(context.Background())
}

func (i IamAuditConfigArray) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigArrayOutput)
}

// IamAuditConfigMapInput is an input type that accepts IamAuditConfigMap and IamAuditConfigMapOutput values.
// You can construct a concrete instance of `IamAuditConfigMapInput` via:
//
//          IamAuditConfigMap{ "key": IamAuditConfigArgs{...} }
type IamAuditConfigMapInput interface {
	pulumi.Input

	ToIamAuditConfigMapOutput() IamAuditConfigMapOutput
	ToIamAuditConfigMapOutputWithContext(context.Context) IamAuditConfigMapOutput
}

type IamAuditConfigMap map[string]IamAuditConfigInput

func (IamAuditConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamAuditConfig)(nil))
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return i.ToIamAuditConfigMapOutputWithContext(context.Background())
}

func (i IamAuditConfigMap) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigMapOutput)
}

type IamAuditConfigOutput struct {
	*pulumi.OutputState
}

func (IamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamAuditConfig)(nil))
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToIamAuditConfigPtrOutput() IamAuditConfigPtrOutput {
	return o.ToIamAuditConfigPtrOutputWithContext(context.Background())
}

func (o IamAuditConfigOutput) ToIamAuditConfigPtrOutputWithContext(ctx context.Context) IamAuditConfigPtrOutput {
	return o.ApplyT(func(v IamAuditConfig) *IamAuditConfig {
		return &v
	}).(IamAuditConfigPtrOutput)
}

type IamAuditConfigPtrOutput struct {
	*pulumi.OutputState
}

func (IamAuditConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamAuditConfig)(nil))
}

func (o IamAuditConfigPtrOutput) ToIamAuditConfigPtrOutput() IamAuditConfigPtrOutput {
	return o
}

func (o IamAuditConfigPtrOutput) ToIamAuditConfigPtrOutputWithContext(ctx context.Context) IamAuditConfigPtrOutput {
	return o
}

type IamAuditConfigArrayOutput struct{ *pulumi.OutputState }

func (IamAuditConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamAuditConfig)(nil))
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutput() IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) ToIamAuditConfigArrayOutputWithContext(ctx context.Context) IamAuditConfigArrayOutput {
	return o
}

func (o IamAuditConfigArrayOutput) Index(i pulumi.IntInput) IamAuditConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamAuditConfig {
		return vs[0].([]IamAuditConfig)[vs[1].(int)]
	}).(IamAuditConfigOutput)
}

type IamAuditConfigMapOutput struct{ *pulumi.OutputState }

func (IamAuditConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamAuditConfig)(nil))
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutput() IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) ToIamAuditConfigMapOutputWithContext(ctx context.Context) IamAuditConfigMapOutput {
	return o
}

func (o IamAuditConfigMapOutput) MapIndex(k pulumi.StringInput) IamAuditConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamAuditConfig {
		return vs[0].(map[string]IamAuditConfig)[vs[1].(string)]
	}).(IamAuditConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(IamAuditConfigOutput{})
	pulumi.RegisterOutputType(IamAuditConfigPtrOutput{})
	pulumi.RegisterOutputType(IamAuditConfigArrayOutput{})
	pulumi.RegisterOutputType(IamAuditConfigMapOutput{})
}
