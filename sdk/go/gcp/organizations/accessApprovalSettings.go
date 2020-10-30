// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccessApprovalSettings struct {
	pulumi.CustomResourceState

	// This field will always be unset for the organization since organizations do not have ancestors.
	EnrolledAncestor pulumi.BoolOutput `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can be done for individual services. A maximum of 10 enrolled services will be enforced, to be expanded as
	// the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayOutput `pulumi:"enrolledServices"`
	// The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
	// ID of the organization of the access approval settings.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
}

// NewAccessApprovalSettings registers a new resource with the given unique name, arguments, and options.
func NewAccessApprovalSettings(ctx *pulumi.Context,
	name string, args *AccessApprovalSettingsArgs, opts ...pulumi.ResourceOption) (*AccessApprovalSettings, error) {
	if args == nil || args.EnrolledServices == nil {
		return nil, errors.New("missing required argument 'EnrolledServices'")
	}
	if args == nil || args.OrganizationId == nil {
		return nil, errors.New("missing required argument 'OrganizationId'")
	}
	if args == nil {
		args = &AccessApprovalSettingsArgs{}
	}
	var resource AccessApprovalSettings
	err := ctx.RegisterResource("gcp:organizations/accessApprovalSettings:AccessApprovalSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessApprovalSettings gets an existing AccessApprovalSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessApprovalSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessApprovalSettingsState, opts ...pulumi.ResourceOption) (*AccessApprovalSettings, error) {
	var resource AccessApprovalSettings
	err := ctx.ReadResource("gcp:organizations/accessApprovalSettings:AccessApprovalSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessApprovalSettings resources.
type accessApprovalSettingsState struct {
	// This field will always be unset for the organization since organizations do not have ancestors.
	EnrolledAncestor *bool `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can be done for individual services. A maximum of 10 enrolled services will be enforced, to be expanded as
	// the set of supported services is expanded.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"
	Name *string `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// ID of the organization of the access approval settings.
	OrganizationId *string `pulumi:"organizationId"`
}

type AccessApprovalSettingsState struct {
	// This field will always be unset for the organization since organizations do not have ancestors.
	EnrolledAncestor pulumi.BoolPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can be done for individual services. A maximum of 10 enrolled services will be enforced, to be expanded as
	// the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"
	Name pulumi.StringPtrInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
	// ID of the organization of the access approval settings.
	OrganizationId pulumi.StringPtrInput
}

func (AccessApprovalSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsState)(nil)).Elem()
}

type accessApprovalSettingsArgs struct {
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can be done for individual services. A maximum of 10 enrolled services will be enforced, to be expanded as
	// the set of supported services is expanded.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
	// ID of the organization of the access approval settings.
	OrganizationId string `pulumi:"organizationId"`
}

// The set of arguments for constructing a AccessApprovalSettings resource.
type AccessApprovalSettingsArgs struct {
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can be done for individual services. A maximum of 10 enrolled services will be enforced, to be expanded as
	// the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
	// ID of the organization of the access approval settings.
	OrganizationId pulumi.StringInput
}

func (AccessApprovalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsArgs)(nil)).Elem()
}

type AccessApprovalSettingsInput interface {
	pulumi.Input

	ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput
	ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput
}

func (AccessApprovalSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessApprovalSettings)(nil)).Elem()
}

func (i AccessApprovalSettings) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return i.ToAccessApprovalSettingsOutputWithContext(context.Background())
}

func (i AccessApprovalSettings) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsOutput)
}

type AccessApprovalSettingsOutput struct {
	*pulumi.OutputState
}

func (AccessApprovalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessApprovalSettingsOutput)(nil)).Elem()
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return o
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccessApprovalSettingsOutput{})
}
