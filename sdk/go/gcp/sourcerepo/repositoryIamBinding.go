// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sourcerepo

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud Pub/Sub Topic. Each of these resources serves a different use case:
//
// * `pubsub.TopicIAMPolicy`: Authoritative. Sets the IAM policy for the topic and replaces any existing policy already attached.
// * `pubsub.TopicIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the topic are preserved.
// * `pubsub.TopicIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the topic are preserved.
//
// > **Note:** `pubsub.TopicIAMPolicy` **cannot** be used in conjunction with `pubsub.TopicIAMBinding` and `pubsub.TopicIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `pubsub.TopicIAMBinding` resources **can be** used in conjunction with `pubsub.TopicIAMMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_pubsub\_topic\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
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
// 		_, err = pubsub.NewTopicIAMPolicy(ctx, "policy", &pubsub.TopicIAMPolicyArgs{
// 			Project:    pulumi.String(google_pubsub_topic.Example.Project),
// 			Topic:      pulumi.String(google_pubsub_topic.Example.Name),
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
// ## google\_pubsub\_topic\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = pubsub.NewTopicIAMBinding(ctx, "binding", &pubsub.TopicIAMBindingArgs{
// 			Project: pulumi.String(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.String(google_pubsub_topic.Example.Name),
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
// ## google\_pubsub\_topic\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = pubsub.NewTopicIAMMember(ctx, "member", &pubsub.TopicIAMMemberArgs{
// 			Project: pulumi.String(google_pubsub_topic.Example.Project),
// 			Topic:   pulumi.String(google_pubsub_topic.Example.Name),
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
type RepositoryIamBinding struct {
	pulumi.CustomResourceState

	Condition RepositoryIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringOutput `pulumi:"project"`
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRepositoryIamBinding registers a new resource with the given unique name, arguments, and options.
func NewRepositoryIamBinding(ctx *pulumi.Context,
	name string, args *RepositoryIamBindingArgs, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &RepositoryIamBindingArgs{}
	}
	var resource RepositoryIamBinding
	err := ctx.RegisterResource("gcp:sourcerepo/repositoryIamBinding:RepositoryIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryIamBinding gets an existing RepositoryIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryIamBindingState, opts ...pulumi.ResourceOption) (*RepositoryIamBinding, error) {
	var resource RepositoryIamBinding
	err := ctx.ReadResource("gcp:sourcerepo/repositoryIamBinding:RepositoryIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryIamBinding resources.
type repositoryIamBindingState struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    *string `pulumi:"project"`
	Repository *string `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type RepositoryIamBindingState struct {
	Condition RepositoryIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringPtrInput
	Repository pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (RepositoryIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingState)(nil)).Elem()
}

type repositoryIamBindingArgs struct {
	Condition *RepositoryIamBindingCondition `pulumi:"condition"`
	Members   []string                       `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    *string `pulumi:"project"`
	Repository string  `pulumi:"repository"`
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a RepositoryIamBinding resource.
type RepositoryIamBindingArgs struct {
	Condition RepositoryIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project    pulumi.StringPtrInput
	Repository pulumi.StringInput
	// The role that should be applied. Only one
	// `pubsub.TopicIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (RepositoryIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryIamBindingArgs)(nil)).Elem()
}
