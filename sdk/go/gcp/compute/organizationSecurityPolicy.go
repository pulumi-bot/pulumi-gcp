// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Organization security policies are used to control incoming/outgoing traffic.
//
// To get more information about OrganizationSecurityPolicy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies)
// * How-to Guides
//     * [Creating a firewall policy](https://cloud.google.com/vpc/docs/using-firewall-policies#create-policy)
//
// ## Example Usage
type OrganizationSecurityPolicy struct {
	pulumi.CustomResourceState

	// A textual description for the organization security policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A textual name of the security policy.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// The unique identifier for the resource. This identifier is defined by the server.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL".
	// Default value is `FIREWALL`.
	// Possible values are `FIREWALL`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewOrganizationSecurityPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSecurityPolicy(ctx *pulumi.Context,
	name string, args *OrganizationSecurityPolicyArgs, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource OrganizationSecurityPolicy
	err := ctx.RegisterResource("gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSecurityPolicy gets an existing OrganizationSecurityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSecurityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSecurityPolicyState, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicy, error) {
	var resource OrganizationSecurityPolicy
	err := ctx.ReadResource("gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSecurityPolicy resources.
type organizationSecurityPolicyState struct {
	// A textual description for the organization security policy.
	Description *string `pulumi:"description"`
	// A textual name of the security policy.
	DisplayName *string `pulumi:"displayName"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}
	Parent *string `pulumi:"parent"`
	// The unique identifier for the resource. This identifier is defined by the server.
	PolicyId *string `pulumi:"policyId"`
	// The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL".
	// Default value is `FIREWALL`.
	// Possible values are `FIREWALL`.
	Type *string `pulumi:"type"`
}

type OrganizationSecurityPolicyState struct {
	// A textual description for the organization security policy.
	Description pulumi.StringPtrInput
	// A textual name of the security policy.
	DisplayName pulumi.StringPtrInput
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringPtrInput
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}
	Parent pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	PolicyId pulumi.StringPtrInput
	// The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL".
	// Default value is `FIREWALL`.
	// Possible values are `FIREWALL`.
	Type pulumi.StringPtrInput
}

func (OrganizationSecurityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyState)(nil)).Elem()
}

type organizationSecurityPolicyArgs struct {
	// A textual description for the organization security policy.
	Description *string `pulumi:"description"`
	// A textual name of the security policy.
	DisplayName string `pulumi:"displayName"`
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}
	Parent string `pulumi:"parent"`
	// The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL".
	// Default value is `FIREWALL`.
	// Possible values are `FIREWALL`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OrganizationSecurityPolicy resource.
type OrganizationSecurityPolicyArgs struct {
	// A textual description for the organization security policy.
	Description pulumi.StringPtrInput
	// A textual name of the security policy.
	DisplayName pulumi.StringInput
	// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
	// Format: organizations/{organization_id} or folders/{folder_id}
	Parent pulumi.StringInput
	// The type indicates the intended use of the security policy.
	// For organization security policies, the only supported type
	// is "FIREWALL".
	// Default value is `FIREWALL`.
	// Possible values are `FIREWALL`.
	Type pulumi.StringPtrInput
}

func (OrganizationSecurityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyArgs)(nil)).Elem()
}
