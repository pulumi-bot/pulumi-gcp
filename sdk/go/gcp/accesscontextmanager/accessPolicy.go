// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package accesscontextmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// AccessPolicy is a container for AccessLevels (which define the necessary
// attributes to use GCP services) and ServicePerimeters (which define
// regions of services able to freely pass data within a perimeter). An
// access policy is globally visible within an organization, and the
// restrictions it specifies apply to all projects within an organization.
//
// To get more information about AccessPolicy, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
// you must specify a `billingProject` and set `userProjectOverride` to true
// in the provider configuration. Otherwise the ACM API will return a 403 error.
// Your account must have the `serviceusage.services.use` permission on the
// `billingProject` you defined.
//
// ## Example Usage
//
// ## Import
//
// AccessPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:accesscontextmanager/accessPolicy:AccessPolicy default {{name}}
// ```
type AccessPolicy struct {
	pulumi.CustomResourceState

	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Human readable title. Does not affect behavior.
	Title pulumi.StringOutput `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessPolicy(ctx *pulumi.Context,
	name string, args *AccessPolicyArgs, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil {
		args = &AccessPolicyArgs{}
	}
	var resource AccessPolicy
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessPolicy:AccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPolicy gets an existing AccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPolicyState, opts ...pulumi.ResourceOption) (*AccessPolicy, error) {
	var resource AccessPolicy
	err := ctx.ReadResource("gcp:accesscontextmanager/accessPolicy:AccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPolicy resources.
type accessPolicyState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name *string `pulumi:"name"`
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent *string `pulumi:"parent"`
	// Human readable title. Does not affect behavior.
	Title *string `pulumi:"title"`
	// Time the AccessPolicy was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type AccessPolicyState struct {
	// Time the AccessPolicy was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Resource name of the AccessPolicy. Format: {policy_id}
	Name pulumi.StringPtrInput
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringPtrInput
	// Human readable title. Does not affect behavior.
	Title pulumi.StringPtrInput
	// Time the AccessPolicy was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (AccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyState)(nil)).Elem()
}

type accessPolicyArgs struct {
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent string `pulumi:"parent"`
	// Human readable title. Does not affect behavior.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a AccessPolicy resource.
type AccessPolicyArgs struct {
	// The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id}
	Parent pulumi.StringInput
	// Human readable title. Does not affect behavior.
	Title pulumi.StringInput
}

func (AccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPolicyArgs)(nil)).Elem()
}
