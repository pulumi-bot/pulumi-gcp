// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineVersion. Each of these resources serves a different use case:
//
// * `iap.AppEngineVersionIamPolicy`: Authoritative. Sets the IAM policy for the appengineversion and replaces any existing policy already attached.
// * `iap.AppEngineVersionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineversion are preserved.
// * `iap.AppEngineVersionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineversion are preserved.
//
// > **Note:** `iap.AppEngineVersionIamPolicy` **cannot** be used in conjunction with `iap.AppEngineVersionIamBinding` and `iap.AppEngineVersionIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.AppEngineVersionIamBinding` resources **can be** used in conjunction with `iap.AppEngineVersionIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_app\_engine\_version\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/iap.httpsResourceAccessor",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 					Condition: organizations.GetIAMPolicyBindingCondition{
// 						Title:       "expires_after_2019_12_31",
// 						Description: "Expiring at midnight of 2019-12-31",
// 						Expression:  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewAppEngineVersionIamPolicy(ctx, "policy", &iap.AppEngineVersionIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId:  pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_version\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineVersionIamBinding(ctx, "binding", &iap.AppEngineVersionIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineVersionIamBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_version\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
// 			AppId:     pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineVersionIamMember(ctx, "member", &iap.AppEngineVersionIamMemberArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineVersionIamMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:    pulumi.String("user:jane@example.com"),
// 			Project:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:      pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service:   pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			VersionId: pulumi.Any(google_app_engine_standard_app_version.Version.Version_id),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} * {{project}}/{{appId}}/{{service}}/{{versionId}} * {{appId}}/{{service}}/{{versionId}} * {{version}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineversion IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineVersionIamPolicy struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewAppEngineVersionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppEngineVersionIamPolicy(ctx *pulumi.Context,
	name string, args *AppEngineVersionIamPolicyArgs, opts ...pulumi.ResourceOption) (*AppEngineVersionIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.VersionId == nil {
		return nil, errors.New("invalid value for required argument 'VersionId'")
	}
	var resource AppEngineVersionIamPolicy
	err := ctx.RegisterResource("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineVersionIamPolicy gets an existing AppEngineVersionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineVersionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineVersionIamPolicyState, opts ...pulumi.ResourceOption) (*AppEngineVersionIamPolicy, error) {
	var resource AppEngineVersionIamPolicy
	err := ctx.ReadResource("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineVersionIamPolicy resources.
type appEngineVersionIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId *string `pulumi:"versionId"`
}

type AppEngineVersionIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringPtrInput
}

func (AppEngineVersionIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamPolicyState)(nil)).Elem()
}

type appEngineVersionIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a AppEngineVersionIamPolicy resource.
type AppEngineVersionIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
	// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
	VersionId pulumi.StringInput
}

func (AppEngineVersionIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineVersionIamPolicyArgs)(nil)).Elem()
}

type AppEngineVersionIamPolicyInput interface {
	pulumi.Input

	ToAppEngineVersionIamPolicyOutput() AppEngineVersionIamPolicyOutput
	ToAppEngineVersionIamPolicyOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyOutput
}

func (*AppEngineVersionIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineVersionIamPolicy)(nil))
}

func (i *AppEngineVersionIamPolicy) ToAppEngineVersionIamPolicyOutput() AppEngineVersionIamPolicyOutput {
	return i.ToAppEngineVersionIamPolicyOutputWithContext(context.Background())
}

func (i *AppEngineVersionIamPolicy) ToAppEngineVersionIamPolicyOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamPolicyOutput)
}

func (i *AppEngineVersionIamPolicy) ToAppEngineVersionIamPolicyPtrOutput() AppEngineVersionIamPolicyPtrOutput {
	return i.ToAppEngineVersionIamPolicyPtrOutputWithContext(context.Background())
}

func (i *AppEngineVersionIamPolicy) ToAppEngineVersionIamPolicyPtrOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamPolicyPtrOutput)
}

type AppEngineVersionIamPolicyPtrInput interface {
	pulumi.Input

	ToAppEngineVersionIamPolicyPtrOutput() AppEngineVersionIamPolicyPtrOutput
	ToAppEngineVersionIamPolicyPtrOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyPtrOutput
}

type appEngineVersionIamPolicyPtrType AppEngineVersionIamPolicyArgs

func (*appEngineVersionIamPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamPolicy)(nil))
}

func (i *appEngineVersionIamPolicyPtrType) ToAppEngineVersionIamPolicyPtrOutput() AppEngineVersionIamPolicyPtrOutput {
	return i.ToAppEngineVersionIamPolicyPtrOutputWithContext(context.Background())
}

func (i *appEngineVersionIamPolicyPtrType) ToAppEngineVersionIamPolicyPtrOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamPolicyOutput).ToAppEngineVersionIamPolicyPtrOutput()
}

// AppEngineVersionIamPolicyArrayInput is an input type that accepts AppEngineVersionIamPolicyArray and AppEngineVersionIamPolicyArrayOutput values.
// You can construct a concrete instance of `AppEngineVersionIamPolicyArrayInput` via:
//
//          AppEngineVersionIamPolicyArray{ AppEngineVersionIamPolicyArgs{...} }
type AppEngineVersionIamPolicyArrayInput interface {
	pulumi.Input

	ToAppEngineVersionIamPolicyArrayOutput() AppEngineVersionIamPolicyArrayOutput
	ToAppEngineVersionIamPolicyArrayOutputWithContext(context.Context) AppEngineVersionIamPolicyArrayOutput
}

type AppEngineVersionIamPolicyArray []AppEngineVersionIamPolicyInput

func (AppEngineVersionIamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppEngineVersionIamPolicy)(nil))
}

func (i AppEngineVersionIamPolicyArray) ToAppEngineVersionIamPolicyArrayOutput() AppEngineVersionIamPolicyArrayOutput {
	return i.ToAppEngineVersionIamPolicyArrayOutputWithContext(context.Background())
}

func (i AppEngineVersionIamPolicyArray) ToAppEngineVersionIamPolicyArrayOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamPolicyArrayOutput)
}

// AppEngineVersionIamPolicyMapInput is an input type that accepts AppEngineVersionIamPolicyMap and AppEngineVersionIamPolicyMapOutput values.
// You can construct a concrete instance of `AppEngineVersionIamPolicyMapInput` via:
//
//          AppEngineVersionIamPolicyMap{ "key": AppEngineVersionIamPolicyArgs{...} }
type AppEngineVersionIamPolicyMapInput interface {
	pulumi.Input

	ToAppEngineVersionIamPolicyMapOutput() AppEngineVersionIamPolicyMapOutput
	ToAppEngineVersionIamPolicyMapOutputWithContext(context.Context) AppEngineVersionIamPolicyMapOutput
}

type AppEngineVersionIamPolicyMap map[string]AppEngineVersionIamPolicyInput

func (AppEngineVersionIamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppEngineVersionIamPolicy)(nil))
}

func (i AppEngineVersionIamPolicyMap) ToAppEngineVersionIamPolicyMapOutput() AppEngineVersionIamPolicyMapOutput {
	return i.ToAppEngineVersionIamPolicyMapOutputWithContext(context.Background())
}

func (i AppEngineVersionIamPolicyMap) ToAppEngineVersionIamPolicyMapOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppEngineVersionIamPolicyMapOutput)
}

type AppEngineVersionIamPolicyOutput struct {
	*pulumi.OutputState
}

func (AppEngineVersionIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppEngineVersionIamPolicy)(nil))
}

func (o AppEngineVersionIamPolicyOutput) ToAppEngineVersionIamPolicyOutput() AppEngineVersionIamPolicyOutput {
	return o
}

func (o AppEngineVersionIamPolicyOutput) ToAppEngineVersionIamPolicyOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyOutput {
	return o
}

func (o AppEngineVersionIamPolicyOutput) ToAppEngineVersionIamPolicyPtrOutput() AppEngineVersionIamPolicyPtrOutput {
	return o.ToAppEngineVersionIamPolicyPtrOutputWithContext(context.Background())
}

func (o AppEngineVersionIamPolicyOutput) ToAppEngineVersionIamPolicyPtrOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyPtrOutput {
	return o.ApplyT(func(v AppEngineVersionIamPolicy) *AppEngineVersionIamPolicy {
		return &v
	}).(AppEngineVersionIamPolicyPtrOutput)
}

type AppEngineVersionIamPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (AppEngineVersionIamPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppEngineVersionIamPolicy)(nil))
}

func (o AppEngineVersionIamPolicyPtrOutput) ToAppEngineVersionIamPolicyPtrOutput() AppEngineVersionIamPolicyPtrOutput {
	return o
}

func (o AppEngineVersionIamPolicyPtrOutput) ToAppEngineVersionIamPolicyPtrOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyPtrOutput {
	return o
}

type AppEngineVersionIamPolicyArrayOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppEngineVersionIamPolicy)(nil))
}

func (o AppEngineVersionIamPolicyArrayOutput) ToAppEngineVersionIamPolicyArrayOutput() AppEngineVersionIamPolicyArrayOutput {
	return o
}

func (o AppEngineVersionIamPolicyArrayOutput) ToAppEngineVersionIamPolicyArrayOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyArrayOutput {
	return o
}

func (o AppEngineVersionIamPolicyArrayOutput) Index(i pulumi.IntInput) AppEngineVersionIamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppEngineVersionIamPolicy {
		return vs[0].([]AppEngineVersionIamPolicy)[vs[1].(int)]
	}).(AppEngineVersionIamPolicyOutput)
}

type AppEngineVersionIamPolicyMapOutput struct{ *pulumi.OutputState }

func (AppEngineVersionIamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppEngineVersionIamPolicy)(nil))
}

func (o AppEngineVersionIamPolicyMapOutput) ToAppEngineVersionIamPolicyMapOutput() AppEngineVersionIamPolicyMapOutput {
	return o
}

func (o AppEngineVersionIamPolicyMapOutput) ToAppEngineVersionIamPolicyMapOutputWithContext(ctx context.Context) AppEngineVersionIamPolicyMapOutput {
	return o
}

func (o AppEngineVersionIamPolicyMapOutput) MapIndex(k pulumi.StringInput) AppEngineVersionIamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppEngineVersionIamPolicy {
		return vs[0].(map[string]AppEngineVersionIamPolicy)[vs[1].(string)]
	}).(AppEngineVersionIamPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(AppEngineVersionIamPolicyOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamPolicyPtrOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamPolicyArrayOutput{})
	pulumi.RegisterOutputType(AppEngineVersionIamPolicyMapOutput{})
}
