// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iamAuditConfig

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-gcp/organizations/IamAuditConfigAuditLogConfig"
)

type IamAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs organizationsIamAuditConfigAuditLogConfig.IamAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIamAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIamAuditConfig(ctx *pulumi.Context,
	name string, args *IamAuditConfigArgs, opts ...pulumi.ResourceOption) (*IamAuditConfig, error) {
	if args == nil || args.AuditLogConfigs == nil {
		return nil, errors.New("missing required argument 'AuditLogConfigs'")
	}
	if args == nil || args.OrgId == nil {
		return nil, errors.New("missing required argument 'OrgId'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &IamAuditConfigArgs{}
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
	AuditLogConfigs []organizationsIamAuditConfigAuditLogConfig.IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	Etag *string `pulumi:"etag"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId *string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IamAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs organizationsIamAuditConfigAuditLogConfig.IamAuditConfigAuditLogConfigArrayInput
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
	AuditLogConfigs []organizationsIamAuditConfigAuditLogConfig.IamAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId string `pulumi:"orgId"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IamAuditConfig resource.
type IamAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs organizationsIamAuditConfigAuditLogConfig.IamAuditConfigAuditLogConfigArrayInput
	// The numeric ID of the organization in which you want to manage the audit logging config.
	OrgId pulumi.StringInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringInput
}

func (IamAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamAuditConfigArgs)(nil)).Elem()
}

