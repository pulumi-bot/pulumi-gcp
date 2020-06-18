// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Runtime Configurator Config. Each of these resources serves a different use case:
//
// * `runtimeconfig.ConfigIamPolicy`: Authoritative. Sets the IAM policy for the config and replaces any existing policy already attached.
// * `runtimeconfig.ConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the config are preserved.
// * `runtimeconfig.ConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the config are preserved.
//
// > **Note:** `runtimeconfig.ConfigIamPolicy` **cannot** be used in conjunction with `runtimeconfig.ConfigIamBinding` and `runtimeconfig.ConfigIamMember` or they will fight over what your policy should be.
//
// > **Note:** `runtimeconfig.ConfigIamBinding` resources **can be** used in conjunction with `runtimeconfig.ConfigIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_runtimeconfig\_config\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Binding: []map[string]interface{}{
// 				map[string]interface{}{
// 					"role": "roles/viewer",
// 					"members": []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = runtimeconfig.NewConfigIamPolicy(ctx, "policy", &runtimeconfig.ConfigIamPolicyArgs{
// 			Project:    pulumi.String(google_runtimeconfig_config.Config.Project),
// 			Config:     pulumi.String(google_runtimeconfig_config.Config.Name),
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
// ## google\_runtimeconfig\_config\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = runtimeconfig.NewConfigIamBinding(ctx, "binding", &runtimeconfig.ConfigIamBindingArgs{
// 			Project: pulumi.String(google_runtimeconfig_config.Config.Project),
// 			Config:  pulumi.String(google_runtimeconfig_config.Config.Name),
// 			Role:    pulumi.String("roles/viewer"),
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
// ## google\_runtimeconfig\_config\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/runtimeconfig"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = runtimeconfig.NewConfigIamMember(ctx, "member", &runtimeconfig.ConfigIamMemberArgs{
// 			Project: pulumi.String(google_runtimeconfig_config.Config.Project),
// 			Config:  pulumi.String(google_runtimeconfig_config.Config.Name),
// 			Role:    pulumi.String("roles/viewer"),
// 			Member:  pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ConfigIamPolicy struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringOutput `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*ConfigIamPolicy, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &ConfigIamPolicyArgs{}
	}
	var resource ConfigIamPolicy
	err := ctx.RegisterResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigIamPolicy gets an existing ConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigIamPolicyState, opts ...pulumi.ResourceOption) (*ConfigIamPolicy, error) {
	var resource ConfigIamPolicy
	err := ctx.ReadResource("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigIamPolicy resources.
type configIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config *string `pulumi:"config"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type ConfigIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamPolicyState)(nil)).Elem()
}

type configIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config string `pulumi:"config"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a ConfigIamPolicy resource.
type ConfigIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (ConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configIamPolicyArgs)(nil)).Elem()
}
