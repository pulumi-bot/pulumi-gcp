// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Directory Namespace. Each of these resources serves a different use case:
//
// * `servicedirectory.NamespaceIamPolicy`: Authoritative. Sets the IAM policy for the namespace and replaces any existing policy already attached.
// * `servicedirectory.NamespaceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the namespace are preserved.
// * `servicedirectory.NamespaceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the namespace are preserved.
//
// > **Note:** `servicedirectory.NamespaceIamPolicy` **cannot** be used in conjunction with `servicedirectory.NamespaceIamBinding` and `servicedirectory.NamespaceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `servicedirectory.NamespaceIamBinding` resources **can be** used in conjunction with `servicedirectory.NamespaceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_directory\_namespace\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = servicedirectory.NewNamespaceIamPolicy(ctx, "policy", &servicedirectory.NamespaceIamPolicyArgs{
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
// ## google\_service\_directory\_namespace\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewNamespaceIamBinding(ctx, "binding", &servicedirectory.NamespaceIamBindingArgs{
// 			Role: pulumi.String("roles/viewer"),
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
// ## google\_service\_directory\_namespace\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewNamespaceIamMember(ctx, "member", &servicedirectory.NamespaceIamMemberArgs{
// 			Role:   pulumi.String("roles/viewer"),
// 			Member: pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} * {{project}}/{{location}}/{{namespace_id}} * {{location}}/{{namespace_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory namespace IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamMember:NamespaceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamMember:NamespaceIamMember editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamMember:NamespaceIamMember editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NamespaceIamMember struct {
	pulumi.CustomResourceState

	Condition NamespaceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewNamespaceIamMember registers a new resource with the given unique name, arguments, and options.
func NewNamespaceIamMember(ctx *pulumi.Context,
	name string, args *NamespaceIamMemberArgs, opts ...pulumi.ResourceOption) (*NamespaceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource NamespaceIamMember
	err := ctx.RegisterResource("gcp:servicedirectory/namespaceIamMember:NamespaceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceIamMember gets an existing NamespaceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIamMemberState, opts ...pulumi.ResourceOption) (*NamespaceIamMember, error) {
	var resource NamespaceIamMember
	err := ctx.ReadResource("gcp:servicedirectory/namespaceIamMember:NamespaceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceIamMember resources.
type namespaceIamMemberState struct {
	Condition *NamespaceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type NamespaceIamMemberState struct {
	Condition NamespaceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (NamespaceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamMemberState)(nil)).Elem()
}

type namespaceIamMemberArgs struct {
	Condition *NamespaceIamMemberCondition `pulumi:"condition"`
	Member    string                       `pulumi:"member"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a NamespaceIamMember resource.
type NamespaceIamMemberArgs struct {
	Condition NamespaceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (NamespaceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamMemberArgs)(nil)).Elem()
}

type NamespaceIamMemberInput interface {
	pulumi.Input

	ToNamespaceIamMemberOutput() NamespaceIamMemberOutput
	ToNamespaceIamMemberOutputWithContext(ctx context.Context) NamespaceIamMemberOutput
}

func (*NamespaceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceIamMember)(nil))
}

func (i *NamespaceIamMember) ToNamespaceIamMemberOutput() NamespaceIamMemberOutput {
	return i.ToNamespaceIamMemberOutputWithContext(context.Background())
}

func (i *NamespaceIamMember) ToNamespaceIamMemberOutputWithContext(ctx context.Context) NamespaceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamMemberOutput)
}

func (i *NamespaceIamMember) ToNamespaceIamMemberPtrOutput() NamespaceIamMemberPtrOutput {
	return i.ToNamespaceIamMemberPtrOutputWithContext(context.Background())
}

func (i *NamespaceIamMember) ToNamespaceIamMemberPtrOutputWithContext(ctx context.Context) NamespaceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamMemberPtrOutput)
}

type NamespaceIamMemberPtrInput interface {
	pulumi.Input

	ToNamespaceIamMemberPtrOutput() NamespaceIamMemberPtrOutput
	ToNamespaceIamMemberPtrOutputWithContext(ctx context.Context) NamespaceIamMemberPtrOutput
}

type namespaceIamMemberPtrType NamespaceIamMemberArgs

func (*namespaceIamMemberPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamMember)(nil))
}

func (i *namespaceIamMemberPtrType) ToNamespaceIamMemberPtrOutput() NamespaceIamMemberPtrOutput {
	return i.ToNamespaceIamMemberPtrOutputWithContext(context.Background())
}

func (i *namespaceIamMemberPtrType) ToNamespaceIamMemberPtrOutputWithContext(ctx context.Context) NamespaceIamMemberPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamMemberOutput).ToNamespaceIamMemberPtrOutput()
}

// NamespaceIamMemberArrayInput is an input type that accepts NamespaceIamMemberArray and NamespaceIamMemberArrayOutput values.
// You can construct a concrete instance of `NamespaceIamMemberArrayInput` via:
//
//          NamespaceIamMemberArray{ NamespaceIamMemberArgs{...} }
type NamespaceIamMemberArrayInput interface {
	pulumi.Input

	ToNamespaceIamMemberArrayOutput() NamespaceIamMemberArrayOutput
	ToNamespaceIamMemberArrayOutputWithContext(context.Context) NamespaceIamMemberArrayOutput
}

type NamespaceIamMemberArray []NamespaceIamMemberInput

func (NamespaceIamMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceIamMember)(nil))
}

func (i NamespaceIamMemberArray) ToNamespaceIamMemberArrayOutput() NamespaceIamMemberArrayOutput {
	return i.ToNamespaceIamMemberArrayOutputWithContext(context.Background())
}

func (i NamespaceIamMemberArray) ToNamespaceIamMemberArrayOutputWithContext(ctx context.Context) NamespaceIamMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamMemberArrayOutput)
}

// NamespaceIamMemberMapInput is an input type that accepts NamespaceIamMemberMap and NamespaceIamMemberMapOutput values.
// You can construct a concrete instance of `NamespaceIamMemberMapInput` via:
//
//          NamespaceIamMemberMap{ "key": NamespaceIamMemberArgs{...} }
type NamespaceIamMemberMapInput interface {
	pulumi.Input

	ToNamespaceIamMemberMapOutput() NamespaceIamMemberMapOutput
	ToNamespaceIamMemberMapOutputWithContext(context.Context) NamespaceIamMemberMapOutput
}

type NamespaceIamMemberMap map[string]NamespaceIamMemberInput

func (NamespaceIamMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NamespaceIamMember)(nil))
}

func (i NamespaceIamMemberMap) ToNamespaceIamMemberMapOutput() NamespaceIamMemberMapOutput {
	return i.ToNamespaceIamMemberMapOutputWithContext(context.Background())
}

func (i NamespaceIamMemberMap) ToNamespaceIamMemberMapOutputWithContext(ctx context.Context) NamespaceIamMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamMemberMapOutput)
}

type NamespaceIamMemberOutput struct {
	*pulumi.OutputState
}

func (NamespaceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceIamMember)(nil))
}

func (o NamespaceIamMemberOutput) ToNamespaceIamMemberOutput() NamespaceIamMemberOutput {
	return o
}

func (o NamespaceIamMemberOutput) ToNamespaceIamMemberOutputWithContext(ctx context.Context) NamespaceIamMemberOutput {
	return o
}

func (o NamespaceIamMemberOutput) ToNamespaceIamMemberPtrOutput() NamespaceIamMemberPtrOutput {
	return o.ToNamespaceIamMemberPtrOutputWithContext(context.Background())
}

func (o NamespaceIamMemberOutput) ToNamespaceIamMemberPtrOutputWithContext(ctx context.Context) NamespaceIamMemberPtrOutput {
	return o.ApplyT(func(v NamespaceIamMember) *NamespaceIamMember {
		return &v
	}).(NamespaceIamMemberPtrOutput)
}

type NamespaceIamMemberPtrOutput struct {
	*pulumi.OutputState
}

func (NamespaceIamMemberPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceIamMember)(nil))
}

func (o NamespaceIamMemberPtrOutput) ToNamespaceIamMemberPtrOutput() NamespaceIamMemberPtrOutput {
	return o
}

func (o NamespaceIamMemberPtrOutput) ToNamespaceIamMemberPtrOutputWithContext(ctx context.Context) NamespaceIamMemberPtrOutput {
	return o
}

type NamespaceIamMemberArrayOutput struct{ *pulumi.OutputState }

func (NamespaceIamMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NamespaceIamMember)(nil))
}

func (o NamespaceIamMemberArrayOutput) ToNamespaceIamMemberArrayOutput() NamespaceIamMemberArrayOutput {
	return o
}

func (o NamespaceIamMemberArrayOutput) ToNamespaceIamMemberArrayOutputWithContext(ctx context.Context) NamespaceIamMemberArrayOutput {
	return o
}

func (o NamespaceIamMemberArrayOutput) Index(i pulumi.IntInput) NamespaceIamMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NamespaceIamMember {
		return vs[0].([]NamespaceIamMember)[vs[1].(int)]
	}).(NamespaceIamMemberOutput)
}

type NamespaceIamMemberMapOutput struct{ *pulumi.OutputState }

func (NamespaceIamMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NamespaceIamMember)(nil))
}

func (o NamespaceIamMemberMapOutput) ToNamespaceIamMemberMapOutput() NamespaceIamMemberMapOutput {
	return o
}

func (o NamespaceIamMemberMapOutput) ToNamespaceIamMemberMapOutputWithContext(ctx context.Context) NamespaceIamMemberMapOutput {
	return o
}

func (o NamespaceIamMemberMapOutput) MapIndex(k pulumi.StringInput) NamespaceIamMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NamespaceIamMember {
		return vs[0].(map[string]NamespaceIamMember)[vs[1].(string)]
	}).(NamespaceIamMemberOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceIamMemberOutput{})
	pulumi.RegisterOutputType(NamespaceIamMemberPtrOutput{})
	pulumi.RegisterOutputType(NamespaceIamMemberArrayOutput{})
	pulumi.RegisterOutputType(NamespaceIamMemberMapOutput{})
}
