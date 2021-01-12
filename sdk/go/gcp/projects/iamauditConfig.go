// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
//
// * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
// * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
// * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
// * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.
//
// > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.
//
// > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
//    IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.
//
// ## google\_project\_iam\_policy
//
// > **Be careful!** You can accidentally lock yourself out of your project
//    using this resource. Deleting a `projects.IAMPolicy` removes access
//    from anyone without organization-level access to the project. Proceed with caution.
//    It's not recommended to use `projects.IAMPolicy` with your provider project
//    to avoid locking yourself out, and it should generally only be used with projects
//    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
//    applying the change.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/editor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			Project:    pulumi.String("your-project-id"),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 						Title:       "expires_after_2019_12_31",
// 					},
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Role: "roles/compute.admin",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = projects.NewIAMPolicy(ctx, "project", &projects.IAMPolicyArgs{
// 			PolicyData: pulumi.String(admin.PolicyData),
// 			Project:    pulumi.String("your-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMBinding(ctx, "project", &projects.IAMBindingArgs{
// 			Condition: &projects.IAMBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/container.admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/editor"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// With IAM Conditions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMMember(ctx, "project", &projects.IAMMemberArgs{
// 			Condition: &projects.IAMMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.String("your-project-id"),
// 			Role:    pulumi.String("roles/firebase.admin"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_project\_iam\_audit\_config
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/projects"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := projects.NewIAMAuditConfig(ctx, "project", &projects.IAMAuditConfigArgs{
// 			AuditLogConfigs: projects.IAMAuditConfigAuditLogConfigArray{
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					LogType: pulumi.String("ADMIN_READ"),
// 				},
// 				&projects.IAMAuditConfigAuditLogConfigArgs{
// 					ExemptedMembers: pulumi.StringArray{
// 						pulumi.String("user:joebloggs@hashicorp.com"),
// 					},
// 					LogType: pulumi.String("DATA_READ"),
// 				},
// 			},
// 			Project: pulumi.String("your-project-id"),
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
// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
//
// This member resource can be imported using the `project_id`, role, and member e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id roles/viewer user:foo@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiers; the resource in question and the role.
//
// This binding resource can be imported using the `project_id` and role, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question.
//
// This policy resource can be imported using the `project_id`.
//
// ```sh
//  $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project your-project-id
// ```
//
//  IAM audit config imports use the identifier of the resource in question and the service, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id foo.googleapis.com"
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IAMAuditConfig struct {
	pulumi.CustomResourceState

	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayOutput `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringOutput `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewIAMAuditConfig registers a new resource with the given unique name, arguments, and options.
func NewIAMAuditConfig(ctx *pulumi.Context,
	name string, args *IAMAuditConfigArgs, opts ...pulumi.ResourceOption) (*IAMAuditConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuditLogConfigs == nil {
		return nil, errors.New("invalid value for required argument 'AuditLogConfigs'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	var resource IAMAuditConfig
	err := ctx.RegisterResource("gcp:projects/iAMAuditConfig:IAMAuditConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMAuditConfig gets an existing IAMAuditConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMAuditConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMAuditConfigState, opts ...pulumi.ResourceOption) (*IAMAuditConfig, error) {
	var resource IAMAuditConfig
	err := ctx.ReadResource("gcp:projects/iAMAuditConfig:IAMAuditConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMAuditConfig resources.
type iamauditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IAMAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// (Computed) The etag of the project's IAM policy.
	Etag *string `pulumi:"etag"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project *string `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service *string `pulumi:"service"`
}

type IAMAuditConfigState struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayInput
	// (Computed) The etag of the project's IAM policy.
	Etag pulumi.StringPtrInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringPtrInput
}

func (IAMAuditConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamauditConfigState)(nil)).Elem()
}

type iamauditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs []IAMAuditConfigAuditLogConfig `pulumi:"auditLogConfigs"`
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project *string `pulumi:"project"`
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a IAMAuditConfig resource.
type IAMAuditConfigArgs struct {
	// The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
	AuditLogConfigs IAMAuditConfigAuditLogConfigArrayInput
	// The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
	// Required for `projects.IAMPolicy` - you must explicitly set the project, and it
	// will not be inferred from the provider.
	Project pulumi.StringPtrInput
	// Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
	Service pulumi.StringInput
}

func (IAMAuditConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamauditConfigArgs)(nil)).Elem()
}

type IAMAuditConfigInput interface {
	pulumi.Input

	ToIAMAuditConfigOutput() IAMAuditConfigOutput
	ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput
}

func (*IAMAuditConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMAuditConfig)(nil))
}

func (i *IAMAuditConfig) ToIAMAuditConfigOutput() IAMAuditConfigOutput {
	return i.ToIAMAuditConfigOutputWithContext(context.Background())
}

func (i *IAMAuditConfig) ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigOutput)
}

func (i *IAMAuditConfig) ToIAMAuditConfigPtrOutput() IAMAuditConfigPtrOutput {
	return i.ToIAMAuditConfigPtrOutputWithContext(context.Background())
}

func (i *IAMAuditConfig) ToIAMAuditConfigPtrOutputWithContext(ctx context.Context) IAMAuditConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigPtrOutput)
}

type IAMAuditConfigPtrInput interface {
	pulumi.Input

	ToIAMAuditConfigPtrOutput() IAMAuditConfigPtrOutput
	ToIAMAuditConfigPtrOutputWithContext(ctx context.Context) IAMAuditConfigPtrOutput
}

type iamauditConfigPtrType IAMAuditConfigArgs

func (*iamauditConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAuditConfig)(nil))
}

func (i *iamauditConfigPtrType) ToIAMAuditConfigPtrOutput() IAMAuditConfigPtrOutput {
	return i.ToIAMAuditConfigPtrOutputWithContext(context.Background())
}

func (i *iamauditConfigPtrType) ToIAMAuditConfigPtrOutputWithContext(ctx context.Context) IAMAuditConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigOutput).ToIAMAuditConfigPtrOutput()
}

// IAMAuditConfigArrayInput is an input type that accepts IAMAuditConfigArray and IAMAuditConfigArrayOutput values.
// You can construct a concrete instance of `IAMAuditConfigArrayInput` via:
//
//          IAMAuditConfigArray{ IAMAuditConfigArgs{...} }
type IAMAuditConfigArrayInput interface {
	pulumi.Input

	ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput
	ToIAMAuditConfigArrayOutputWithContext(context.Context) IAMAuditConfigArrayOutput
}

type IAMAuditConfigArray []IAMAuditConfigInput

func (IAMAuditConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IAMAuditConfig)(nil))
}

func (i IAMAuditConfigArray) ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput {
	return i.ToIAMAuditConfigArrayOutputWithContext(context.Background())
}

func (i IAMAuditConfigArray) ToIAMAuditConfigArrayOutputWithContext(ctx context.Context) IAMAuditConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigArrayOutput)
}

// IAMAuditConfigMapInput is an input type that accepts IAMAuditConfigMap and IAMAuditConfigMapOutput values.
// You can construct a concrete instance of `IAMAuditConfigMapInput` via:
//
//          IAMAuditConfigMap{ "key": IAMAuditConfigArgs{...} }
type IAMAuditConfigMapInput interface {
	pulumi.Input

	ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput
	ToIAMAuditConfigMapOutputWithContext(context.Context) IAMAuditConfigMapOutput
}

type IAMAuditConfigMap map[string]IAMAuditConfigInput

func (IAMAuditConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IAMAuditConfig)(nil))
}

func (i IAMAuditConfigMap) ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput {
	return i.ToIAMAuditConfigMapOutputWithContext(context.Background())
}

func (i IAMAuditConfigMap) ToIAMAuditConfigMapOutputWithContext(ctx context.Context) IAMAuditConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAuditConfigMapOutput)
}

type IAMAuditConfigOutput struct {
	*pulumi.OutputState
}

func (IAMAuditConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IAMAuditConfig)(nil))
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigOutput() IAMAuditConfigOutput {
	return o
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigOutputWithContext(ctx context.Context) IAMAuditConfigOutput {
	return o
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigPtrOutput() IAMAuditConfigPtrOutput {
	return o.ToIAMAuditConfigPtrOutputWithContext(context.Background())
}

func (o IAMAuditConfigOutput) ToIAMAuditConfigPtrOutputWithContext(ctx context.Context) IAMAuditConfigPtrOutput {
	return o.ApplyT(func(v IAMAuditConfig) *IAMAuditConfig {
		return &v
	}).(IAMAuditConfigPtrOutput)
}

type IAMAuditConfigPtrOutput struct {
	*pulumi.OutputState
}

func (IAMAuditConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAuditConfig)(nil))
}

func (o IAMAuditConfigPtrOutput) ToIAMAuditConfigPtrOutput() IAMAuditConfigPtrOutput {
	return o
}

func (o IAMAuditConfigPtrOutput) ToIAMAuditConfigPtrOutputWithContext(ctx context.Context) IAMAuditConfigPtrOutput {
	return o
}

type IAMAuditConfigArrayOutput struct{ *pulumi.OutputState }

func (IAMAuditConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IAMAuditConfig)(nil))
}

func (o IAMAuditConfigArrayOutput) ToIAMAuditConfigArrayOutput() IAMAuditConfigArrayOutput {
	return o
}

func (o IAMAuditConfigArrayOutput) ToIAMAuditConfigArrayOutputWithContext(ctx context.Context) IAMAuditConfigArrayOutput {
	return o
}

func (o IAMAuditConfigArrayOutput) Index(i pulumi.IntInput) IAMAuditConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IAMAuditConfig {
		return vs[0].([]IAMAuditConfig)[vs[1].(int)]
	}).(IAMAuditConfigOutput)
}

type IAMAuditConfigMapOutput struct{ *pulumi.OutputState }

func (IAMAuditConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IAMAuditConfig)(nil))
}

func (o IAMAuditConfigMapOutput) ToIAMAuditConfigMapOutput() IAMAuditConfigMapOutput {
	return o
}

func (o IAMAuditConfigMapOutput) ToIAMAuditConfigMapOutputWithContext(ctx context.Context) IAMAuditConfigMapOutput {
	return o
}

func (o IAMAuditConfigMapOutput) MapIndex(k pulumi.StringInput) IAMAuditConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IAMAuditConfig {
		return vs[0].(map[string]IAMAuditConfig)[vs[1].(string)]
	}).(IAMAuditConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(IAMAuditConfigOutput{})
	pulumi.RegisterOutputType(IAMAuditConfigPtrOutput{})
	pulumi.RegisterOutputType(IAMAuditConfigArrayOutput{})
	pulumi.RegisterOutputType(IAMAuditConfigMapOutput{})
}
