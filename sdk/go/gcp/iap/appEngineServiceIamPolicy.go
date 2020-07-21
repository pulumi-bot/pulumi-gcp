// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineService. Each of these resources serves a different use case:
//
// * `iap.AppEngineServiceIamPolicy`: Authoritative. Sets the IAM policy for the appengineservice and replaces any existing policy already attached.
// * `iap.AppEngineServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineservice are preserved.
// * `iap.AppEngineServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineservice are preserved.
//
// > **Note:** `iap.AppEngineServiceIamPolicy` **cannot** be used in conjunction with `iap.AppEngineServiceIamBinding` and `iap.AppEngineServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.AppEngineServiceIamBinding` resources **can be** used in conjunction with `iap.AppEngineServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_app\_engine\_service\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
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
// 		_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/organizations"
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
// 		_, err = iap.NewAppEngineServiceIamPolicy(ctx, "policy", &iap.AppEngineServiceIamPolicyArgs{
// 			Project:    pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			AppId:      pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Service:    pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_service\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamBinding(ctx, "binding", &iap.AppEngineServiceIamBindingArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineServiceIamBindingConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_app\_engine\_service\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
// 			AppId:   pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/iap"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iap.NewAppEngineServiceIamMember(ctx, "member", &iap.AppEngineServiceIamMemberArgs{
// 			AppId: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Condition: &iap.AppEngineServiceIamMemberConditionArgs{
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 			},
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Project: pulumi.Any(google_app_engine_standard_app_version.Version.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Service: pulumi.Any(google_app_engine_standard_app_version.Version.Service),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AppEngineServiceIamPolicy struct {
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
}

// NewAppEngineServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewAppEngineServiceIamPolicy(ctx *pulumi.Context,
	name string, args *AppEngineServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*AppEngineServiceIamPolicy, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &AppEngineServiceIamPolicyArgs{}
	}
	var resource AppEngineServiceIamPolicy
	err := ctx.RegisterResource("gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineServiceIamPolicy gets an existing AppEngineServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineServiceIamPolicyState, opts ...pulumi.ResourceOption) (*AppEngineServiceIamPolicy, error) {
	var resource AppEngineServiceIamPolicy
	err := ctx.ReadResource("gcp:iap/appEngineServiceIamPolicy:AppEngineServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineServiceIamPolicy resources.
type appEngineServiceIamPolicyState struct {
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
}

type AppEngineServiceIamPolicyState struct {
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
}

func (AppEngineServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamPolicyState)(nil)).Elem()
}

type appEngineServiceIamPolicyArgs struct {
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
}

// The set of arguments for constructing a AppEngineServiceIamPolicy resource.
type AppEngineServiceIamPolicyArgs struct {
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
}

func (AppEngineServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamPolicyArgs)(nil)).Elem()
}
