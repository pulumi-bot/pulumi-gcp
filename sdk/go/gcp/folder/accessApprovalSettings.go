// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package folder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AccessApprovalSettings struct {
	pulumi.CustomResourceState

	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor pulumi.BoolOutput `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayOutput `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringOutput `pulumi:"folderId"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayOutput `pulumi:"notificationEmails"`
}

// NewAccessApprovalSettings registers a new resource with the given unique name, arguments, and options.
func NewAccessApprovalSettings(ctx *pulumi.Context,
	name string, args *AccessApprovalSettingsArgs, opts ...pulumi.ResourceOption) (*AccessApprovalSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnrolledServices == nil {
		return nil, errors.New("invalid value for required argument 'EnrolledServices'")
	}
	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	var resource AccessApprovalSettings
	err := ctx.RegisterResource("gcp:folder/accessApprovalSettings:AccessApprovalSettings", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:folder/accessApprovalSettings:AccessApprovalSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessApprovalSettings resources.
type accessApprovalSettingsState struct {
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor *bool `pulumi:"enrolledAncestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId *string `pulumi:"folderId"`
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name *string `pulumi:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

type AccessApprovalSettingsState struct {
	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors
	// of the Folder.
	EnrolledAncestor pulumi.BoolPtrInput
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringPtrInput
	// The resource name of the settings. Format is "folders/{folder_id}/accessApprovalSettings"
	Name pulumi.StringPtrInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
}

func (AccessApprovalSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsState)(nil)).Elem()
}

type accessApprovalSettingsArgs struct {
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices []AccessApprovalSettingsEnrolledService `pulumi:"enrolledServices"`
	// ID of the folder of the access approval settings.
	FolderId string `pulumi:"folderId"`
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails []string `pulumi:"notificationEmails"`
}

// The set of arguments for constructing a AccessApprovalSettings resource.
type AccessApprovalSettingsArgs struct {
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled. Access requests for the
	// resource given by name against any of these services contained here will be required to have explicit approval.
	// Enrollment can only be done on an all or nothing basis. A maximum of 10 enrolled services will be enforced, to be
	// expanded as the set of supported services is expanded.
	EnrolledServices AccessApprovalSettingsEnrolledServiceArrayInput
	// ID of the folder of the access approval settings.
	FolderId pulumi.StringInput
	// A list of email addresses to which notifications relating to approval requests should be sent. Notifications relating to
	// a resource will be sent to all emails in the settings of ancestor resources of that resource. A maximum of 50 email
	// addresses are allowed.
	NotificationEmails pulumi.StringArrayInput
}

func (AccessApprovalSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessApprovalSettingsArgs)(nil)).Elem()
}

type AccessApprovalSettingsInput interface {
	pulumi.Input

	ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput
	ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput
}

func (*AccessApprovalSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessApprovalSettings)(nil))
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return i.ToAccessApprovalSettingsOutputWithContext(context.Background())
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsOutput)
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsPtrOutput() AccessApprovalSettingsPtrOutput {
	return i.ToAccessApprovalSettingsPtrOutputWithContext(context.Background())
}

func (i *AccessApprovalSettings) ToAccessApprovalSettingsPtrOutputWithContext(ctx context.Context) AccessApprovalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsPtrOutput)
}

type AccessApprovalSettingsPtrInput interface {
	pulumi.Input

	ToAccessApprovalSettingsPtrOutput() AccessApprovalSettingsPtrOutput
	ToAccessApprovalSettingsPtrOutputWithContext(ctx context.Context) AccessApprovalSettingsPtrOutput
}

type accessApprovalSettingsPtrType AccessApprovalSettingsArgs

func (*accessApprovalSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessApprovalSettings)(nil))
}

func (i *accessApprovalSettingsPtrType) ToAccessApprovalSettingsPtrOutput() AccessApprovalSettingsPtrOutput {
	return i.ToAccessApprovalSettingsPtrOutputWithContext(context.Background())
}

func (i *accessApprovalSettingsPtrType) ToAccessApprovalSettingsPtrOutputWithContext(ctx context.Context) AccessApprovalSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsOutput).ToAccessApprovalSettingsPtrOutput()
}

// AccessApprovalSettingsArrayInput is an input type that accepts AccessApprovalSettingsArray and AccessApprovalSettingsArrayOutput values.
// You can construct a concrete instance of `AccessApprovalSettingsArrayInput` via:
//
//          AccessApprovalSettingsArray{ AccessApprovalSettingsArgs{...} }
type AccessApprovalSettingsArrayInput interface {
	pulumi.Input

	ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput
	ToAccessApprovalSettingsArrayOutputWithContext(context.Context) AccessApprovalSettingsArrayOutput
}

type AccessApprovalSettingsArray []AccessApprovalSettingsInput

func (AccessApprovalSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessApprovalSettings)(nil))
}

func (i AccessApprovalSettingsArray) ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput {
	return i.ToAccessApprovalSettingsArrayOutputWithContext(context.Background())
}

func (i AccessApprovalSettingsArray) ToAccessApprovalSettingsArrayOutputWithContext(ctx context.Context) AccessApprovalSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsArrayOutput)
}

// AccessApprovalSettingsMapInput is an input type that accepts AccessApprovalSettingsMap and AccessApprovalSettingsMapOutput values.
// You can construct a concrete instance of `AccessApprovalSettingsMapInput` via:
//
//          AccessApprovalSettingsMap{ "key": AccessApprovalSettingsArgs{...} }
type AccessApprovalSettingsMapInput interface {
	pulumi.Input

	ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput
	ToAccessApprovalSettingsMapOutputWithContext(context.Context) AccessApprovalSettingsMapOutput
}

type AccessApprovalSettingsMap map[string]AccessApprovalSettingsInput

func (AccessApprovalSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessApprovalSettings)(nil))
}

func (i AccessApprovalSettingsMap) ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput {
	return i.ToAccessApprovalSettingsMapOutputWithContext(context.Background())
}

func (i AccessApprovalSettingsMap) ToAccessApprovalSettingsMapOutputWithContext(ctx context.Context) AccessApprovalSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessApprovalSettingsMapOutput)
}

type AccessApprovalSettingsOutput struct {
	*pulumi.OutputState
}

func (AccessApprovalSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessApprovalSettings)(nil))
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutput() AccessApprovalSettingsOutput {
	return o
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsOutputWithContext(ctx context.Context) AccessApprovalSettingsOutput {
	return o
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsPtrOutput() AccessApprovalSettingsPtrOutput {
	return o.ToAccessApprovalSettingsPtrOutputWithContext(context.Background())
}

func (o AccessApprovalSettingsOutput) ToAccessApprovalSettingsPtrOutputWithContext(ctx context.Context) AccessApprovalSettingsPtrOutput {
	return o.ApplyT(func(v AccessApprovalSettings) *AccessApprovalSettings {
		return &v
	}).(AccessApprovalSettingsPtrOutput)
}

type AccessApprovalSettingsPtrOutput struct {
	*pulumi.OutputState
}

func (AccessApprovalSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessApprovalSettings)(nil))
}

func (o AccessApprovalSettingsPtrOutput) ToAccessApprovalSettingsPtrOutput() AccessApprovalSettingsPtrOutput {
	return o
}

func (o AccessApprovalSettingsPtrOutput) ToAccessApprovalSettingsPtrOutputWithContext(ctx context.Context) AccessApprovalSettingsPtrOutput {
	return o
}

type AccessApprovalSettingsArrayOutput struct{ *pulumi.OutputState }

func (AccessApprovalSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessApprovalSettings)(nil))
}

func (o AccessApprovalSettingsArrayOutput) ToAccessApprovalSettingsArrayOutput() AccessApprovalSettingsArrayOutput {
	return o
}

func (o AccessApprovalSettingsArrayOutput) ToAccessApprovalSettingsArrayOutputWithContext(ctx context.Context) AccessApprovalSettingsArrayOutput {
	return o
}

func (o AccessApprovalSettingsArrayOutput) Index(i pulumi.IntInput) AccessApprovalSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessApprovalSettings {
		return vs[0].([]AccessApprovalSettings)[vs[1].(int)]
	}).(AccessApprovalSettingsOutput)
}

type AccessApprovalSettingsMapOutput struct{ *pulumi.OutputState }

func (AccessApprovalSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AccessApprovalSettings)(nil))
}

func (o AccessApprovalSettingsMapOutput) ToAccessApprovalSettingsMapOutput() AccessApprovalSettingsMapOutput {
	return o
}

func (o AccessApprovalSettingsMapOutput) ToAccessApprovalSettingsMapOutputWithContext(ctx context.Context) AccessApprovalSettingsMapOutput {
	return o
}

func (o AccessApprovalSettingsMapOutput) MapIndex(k pulumi.StringInput) AccessApprovalSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AccessApprovalSettings {
		return vs[0].(map[string]AccessApprovalSettings)[vs[1].(string)]
	}).(AccessApprovalSettingsOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessApprovalSettingsOutput{})
	pulumi.RegisterOutputType(AccessApprovalSettingsPtrOutput{})
	pulumi.RegisterOutputType(AccessApprovalSettingsArrayOutput{})
	pulumi.RegisterOutputType(AccessApprovalSettingsMapOutput{})
}
