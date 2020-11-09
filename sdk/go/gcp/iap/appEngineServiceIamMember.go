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
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} * {{project}}/{{appId}}/{{service}} * {{appId}}/{{service}} * {{service}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}} roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type AppEngineServiceIamMember struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewAppEngineServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewAppEngineServiceIamMember(ctx *pulumi.Context,
	name string, args *AppEngineServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*AppEngineServiceIamMember, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &AppEngineServiceIamMemberArgs{}
	}
	var resource AppEngineServiceIamMember
	err := ctx.RegisterResource("gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppEngineServiceIamMember gets an existing AppEngineServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppEngineServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppEngineServiceIamMemberState, opts ...pulumi.ResourceOption) (*AppEngineServiceIamMember, error) {
	var resource AppEngineServiceIamMember
	err := ctx.ReadResource("gcp:iap/appEngineServiceIamMember:AppEngineServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppEngineServiceIamMember resources.
type appEngineServiceIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineServiceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service *string `pulumi:"service"`
}

type AppEngineServiceIamMemberState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringPtrInput
}

func (AppEngineServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamMemberState)(nil)).Elem()
}

type appEngineServiceIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *AppEngineServiceIamMemberCondition `pulumi:"condition"`
	Member    string                              `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a AppEngineServiceIamMember resource.
type AppEngineServiceIamMemberArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition AppEngineServiceIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
	Service pulumi.StringInput
}

func (AppEngineServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appEngineServiceIamMemberArgs)(nil)).Elem()
}
