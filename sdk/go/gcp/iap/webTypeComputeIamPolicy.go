// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeCompute. Each of these resources serves a different use case:
//
// * `iap.WebTypeComputeIamPolicy`: Authoritative. Sets the IAM policy for the webtypecompute and replaces any existing policy already attached.
// * `iap.WebTypeComputeIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypecompute are preserved.
// * `iap.WebTypeComputeIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypecompute are preserved.
//
// > **Note:** `iap.WebTypeComputeIamPolicy` **cannot** be used in conjunction with `iap.WebTypeComputeIamBinding` and `iap.WebTypeComputeIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeComputeIamBinding` resources **can be** used in conjunction with `iap.WebTypeComputeIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_web\_type\_compute\_iam\_policy
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
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/iap.httpsResourceAccessor",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
// 			Project:    pulumi.String(google_project_service.Project_service.Project),
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
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/iap.httpsResourceAccessor",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 					"condition": map[string]interface{}{
// 						"title":       "expires_after_2019_12_31",
// 						"description": "Expiring at midnight of 2019-12-31",
// 						"expression":  "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iap.NewWebTypeComputeIamPolicy(ctx, "policy", &iap.WebTypeComputeIamPolicyArgs{
// 			Project:    pulumi.String(google_project_service.Project_service.Project),
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_type\_compute\_iam\_binding
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
// 		_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
// 			Project: pulumi.String(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
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
// 		_, err := iap.NewWebTypeComputeIamBinding(ctx, "binding", &iap.WebTypeComputeIamBindingArgs{
// 			Project: pulumi.String(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.WebTypeComputeIamBindingConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_type\_compute\_iam\_member
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
// 		_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
// 			Project: pulumi.String(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
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
// 		_, err := iap.NewWebTypeComputeIamMember(ctx, "member", &iap.WebTypeComputeIamMemberArgs{
// 			Project: pulumi.String(google_project_service.Project_service.Project),
// 			Role:    pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 			Condition: &iap.WebTypeComputeIamMemberConditionArgs{
// 				Title:       pulumi.String("expires_after_2019_12_31"),
// 				Description: pulumi.String("Expiring at midnight of 2019-12-31"),
// 				Expression:  pulumi.String("request.time < timestamp(\"2020-01-01T00:00:00Z\")"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type WebTypeComputeIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebTypeComputeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebTypeComputeIamPolicy(ctx *pulumi.Context,
	name string, args *WebTypeComputeIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebTypeComputeIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &WebTypeComputeIamPolicyArgs{}
	}
	var resource WebTypeComputeIamPolicy
	err := ctx.RegisterResource("gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeComputeIamPolicy gets an existing WebTypeComputeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeComputeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeComputeIamPolicyState, opts ...pulumi.ResourceOption) (*WebTypeComputeIamPolicy, error) {
	var resource WebTypeComputeIamPolicy
	err := ctx.ReadResource("gcp:iap/webTypeComputeIamPolicy:WebTypeComputeIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeComputeIamPolicy resources.
type webTypeComputeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebTypeComputeIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeComputeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamPolicyState)(nil)).Elem()
}

type webTypeComputeIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebTypeComputeIamPolicy resource.
type WebTypeComputeIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeComputeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamPolicyArgs)(nil)).Elem()
}
