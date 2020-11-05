// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An association for the OrganizationSecurityPolicy.
//
// To get more information about OrganizationSecurityPolicyAssociation, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
// * How-to Guides
//     * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)
//
// ## Example Usage
type OrganizationSecurityPolicyAssociation struct {
	pulumi.CustomResourceState

	// The resource that the security policy is attached to.
	AttachmentId pulumi.StringOutput `pulumi:"attachmentId"`
	// The display name of the security policy of the association.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name for an association.
	Name pulumi.StringOutput `pulumi:"name"`
	// The security policy ID of the association.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
}

// NewOrganizationSecurityPolicyAssociation registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSecurityPolicyAssociation(ctx *pulumi.Context,
	name string, args *OrganizationSecurityPolicyAssociationArgs, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicyAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'AttachmentId'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	var resource OrganizationSecurityPolicyAssociation
	err := ctx.RegisterResource("gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSecurityPolicyAssociation gets an existing OrganizationSecurityPolicyAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSecurityPolicyAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSecurityPolicyAssociationState, opts ...pulumi.ResourceOption) (*OrganizationSecurityPolicyAssociation, error) {
	var resource OrganizationSecurityPolicyAssociation
	err := ctx.ReadResource("gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSecurityPolicyAssociation resources.
type organizationSecurityPolicyAssociationState struct {
	// The resource that the security policy is attached to.
	AttachmentId *string `pulumi:"attachmentId"`
	// The display name of the security policy of the association.
	DisplayName *string `pulumi:"displayName"`
	// The name for an association.
	Name *string `pulumi:"name"`
	// The security policy ID of the association.
	PolicyId *string `pulumi:"policyId"`
}

type OrganizationSecurityPolicyAssociationState struct {
	// The resource that the security policy is attached to.
	AttachmentId pulumi.StringPtrInput
	// The display name of the security policy of the association.
	DisplayName pulumi.StringPtrInput
	// The name for an association.
	Name pulumi.StringPtrInput
	// The security policy ID of the association.
	PolicyId pulumi.StringPtrInput
}

func (OrganizationSecurityPolicyAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyAssociationState)(nil)).Elem()
}

type organizationSecurityPolicyAssociationArgs struct {
	// The resource that the security policy is attached to.
	AttachmentId string `pulumi:"attachmentId"`
	// The name for an association.
	Name *string `pulumi:"name"`
	// The security policy ID of the association.
	PolicyId string `pulumi:"policyId"`
}

// The set of arguments for constructing a OrganizationSecurityPolicyAssociation resource.
type OrganizationSecurityPolicyAssociationArgs struct {
	// The resource that the security policy is attached to.
	AttachmentId pulumi.StringInput
	// The name for an association.
	Name pulumi.StringPtrInput
	// The security policy ID of the association.
	PolicyId pulumi.StringInput
}

func (OrganizationSecurityPolicyAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityPolicyAssociationArgs)(nil)).Elem()
}
