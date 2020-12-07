// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of audit logging config for a given service for a Google Cloud Platform Organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewIamAuditConfig(ctx, "config", &organizations.IamAuditConfigArgs{
// 			AuditLogConfigs: organizations.IamAuditConfigAuditLogConfigArray{
// 				&organizations.IamAuditConfigAuditLogConfigArgs{
// 					ExemptedMembers: pulumi.StringArray{
// 						pulumi.String("user:joebloggs@hashicorp.com"),
// 					},
// 					LogType: pulumi.String("DATA_READ"),
// 				},
// 			},
// 			OrgId:   pulumi.String("your-organization-id"),
// 			Service: pulumi.String("allServices"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig config "your-organization-id foo.googleapis.com"
// ```
type IamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
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
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IamAuditConfig
	err := ctx.RegisterResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:organizations/iamAuditConfig:IamAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamAuditConfig resources.
type iamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The etag of iam policy
	Etag *string `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId *string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The etag of iam policy
	Etag pulumi.StringPtrInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringPtrInput
}

func (IamAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigState)(nil)).Elem()
}

type iamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IamAuditConfigAuditLogConfigArrayInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
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

func (IamAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*IamAuditConfig)(nil)).Elem()
}

func (i IamAuditConfig) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return i.ToIamAuditConfigOutputWithContext(context.Background())
}

func (i IamAuditConfig) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamAuditConfigOutput)
}

type IamAuditConfigOutput struct {
	*pulumi.OutputState
}

func (IamAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamAuditConfigOutput)(nil)).Elem()
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutput() IamAuditConfigOutput {
	return o
}

func (o IamAuditConfigOutput) ToIamAuditConfigOutputWithContext(ctx context.Context) IamAuditConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IamAuditConfigOutput{})
}
