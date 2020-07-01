// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
//
// * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
// * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
// * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
//
// > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_iap\_web\_backend\_service\_iam\_policy
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
// 		_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			PolicyData:        pulumi.String(admin.PolicyData),
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
// 		_, err = iap.NewWebBackendServiceIamPolicy(ctx, "policy", &iap.WebBackendServiceIamPolicyArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			PolicyData:        pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## google\_iap\_web\_backend\_service\_iam\_binding
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
// 		_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
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
// 		_, err := iap.NewWebBackendServiceIamBinding(ctx, "binding", &iap.WebBackendServiceIamBindingArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 			Condition: &iap.WebBackendServiceIamBindingConditionArgs{
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
// ## google\_iap\_web\_backend\_service\_iam\_member
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
// 		_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:            pulumi.String("user:jane@example.com"),
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
// 		_, err := iap.NewWebBackendServiceIamMember(ctx, "member", &iap.WebBackendServiceIamMemberArgs{
// 			Project:           pulumi.Any(google_compute_backend_service.Default.Project),
// 			WebBackendService: pulumi.Any(google_compute_backend_service.Default.Name),
// 			Role:              pulumi.String("roles/iap.httpsResourceAccessor"),
// 			Member:            pulumi.String("user:jane@example.com"),
// 			Condition: &iap.WebBackendServiceIamMemberConditionArgs{
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
type WebBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringOutput `pulumi:"webBackendService"`
}

// NewWebBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil || args.WebBackendService == nil {
		return nil, errors.New("missing required argument 'WebBackendService'")
	}
	if args == nil {
		args = &WebBackendServiceIamPolicyArgs{}
	}
	var resource WebBackendServiceIamPolicy
	err := ctx.RegisterResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebBackendServiceIamPolicy gets an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*WebBackendServiceIamPolicy, error) {
	var resource WebBackendServiceIamPolicy
	err := ctx.ReadResource("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebBackendServiceIamPolicy resources.
type webBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService *string `pulumi:"webBackendService"`
}

type WebBackendServiceIamPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringPtrInput
}

func (WebBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyState)(nil)).Elem()
}

type webBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService string `pulumi:"webBackendService"`
}

// The set of arguments for constructing a WebBackendServiceIamPolicy resource.
type WebBackendServiceIamPolicyArgs struct {
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService pulumi.StringInput
}

func (WebBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webBackendServiceIamPolicyArgs)(nil)).Elem()
}
