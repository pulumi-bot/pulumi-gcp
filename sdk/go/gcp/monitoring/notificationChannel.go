// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A NotificationChannel is a medium through which an alert is delivered
// when a policy violation is detected. Examples of channels include email, SMS,
// and third-party messaging applications. Fields containing sensitive information
// like authentication tokens or contact info are only partially populated on retrieval.
//
// Notification Channels are designed to be flexible and are made up of a supported `type`
// and labels to configure that channel. Each `type` has specific labels that need to be
// present for that channel to be correctly configured. The labels that are required to be
// present for one channel `type` are often different than those required for another.
// Due to these loose constraints it's often best to set up a channel through the UI
// and import it to the provider when setting up a brand new channel type to determine which
// labels are required.
//
// A list of supported channels per project the `list` endpoint can be
// accessed programmatically or through the api explorer at  https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list .
// This provides the channel type and all of the required labels that must be passed.
//
//
// To get more information about NotificationChannel, see:
//
// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
// * How-to Guides
//     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
//     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
type NotificationChannel struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// -
	// (Required)
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// -
	// (Optional)
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// -
	// (Optional)
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	// Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitiveLabels block, but cannot be configured in both places.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// -
	// (Optional)
	// Different notification type behaviors are configured primarily using the the `labels` field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the `labels` map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
	SensitiveLabels NotificationChannelSensitiveLabelsPtrOutput `pulumi:"sensitiveLabels"`
	// -
	// (Required)
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type pulumi.StringOutput `pulumi:"type"`
	// -
	// (Optional)
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapOutput `pulumi:"userLabels"`
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus pulumi.StringOutput `pulumi:"verificationStatus"`
}

// NewNotificationChannel registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannel(ctx *pulumi.Context,
	name string, args *NotificationChannelArgs, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &NotificationChannelArgs{}
	}
	var resource NotificationChannel
	err := ctx.RegisterResource("gcp:monitoring/notificationChannel:NotificationChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationChannel gets an existing NotificationChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationChannelState, opts ...pulumi.ResourceOption) (*NotificationChannel, error) {
	var resource NotificationChannel
	err := ctx.ReadResource("gcp:monitoring/notificationChannel:NotificationChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationChannel resources.
type notificationChannelState struct {
	// -
	// (Optional)
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `pulumi:"description"`
	// -
	// (Required)
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName *string `pulumi:"displayName"`
	// -
	// (Optional)
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled *bool `pulumi:"enabled"`
	// -
	// (Optional)
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	// Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitiveLabels block, but cannot be configured in both places.
	Labels map[string]string `pulumi:"labels"`
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// Different notification type behaviors are configured primarily using the the `labels` field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the `labels` map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
	SensitiveLabels *NotificationChannelSensitiveLabels `pulumi:"sensitiveLabels"`
	// -
	// (Required)
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type *string `pulumi:"type"`
	// -
	// (Optional)
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus *string `pulumi:"verificationStatus"`
}

type NotificationChannelState struct {
	// -
	// (Optional)
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrInput
	// -
	// (Required)
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName pulumi.StringPtrInput
	// -
	// (Optional)
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrInput
	// -
	// (Optional)
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	// Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitiveLabels block, but cannot be configured in both places.
	Labels pulumi.StringMapInput
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// Different notification type behaviors are configured primarily using the the `labels` field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the `labels` map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
	SensitiveLabels NotificationChannelSensitiveLabelsPtrInput
	// -
	// (Required)
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type pulumi.StringPtrInput
	// -
	// (Optional)
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus pulumi.StringPtrInput
}

func (NotificationChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelState)(nil)).Elem()
}

type notificationChannelArgs struct {
	// -
	// (Optional)
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `pulumi:"description"`
	// -
	// (Required)
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName string `pulumi:"displayName"`
	// -
	// (Optional)
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled *bool `pulumi:"enabled"`
	// -
	// (Optional)
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	// Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitiveLabels block, but cannot be configured in both places.
	Labels map[string]string `pulumi:"labels"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// -
	// (Optional)
	// Different notification type behaviors are configured primarily using the the `labels` field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the `labels` map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
	SensitiveLabels *NotificationChannelSensitiveLabels `pulumi:"sensitiveLabels"`
	// -
	// (Required)
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type string `pulumi:"type"`
	// -
	// (Optional)
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// The set of arguments for constructing a NotificationChannel resource.
type NotificationChannelArgs struct {
	// -
	// (Optional)
	// An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrInput
	// -
	// (Required)
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	DisplayName pulumi.StringInput
	// -
	// (Optional)
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrInput
	// -
	// (Optional)
	// Configuration fields that define the channel and its behavior. The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	// Labels with sensitive data are obfuscated by the API and therefore the provider cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitiveLabels block, but cannot be configured in both places.
	Labels pulumi.StringMapInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// -
	// (Optional)
	// Different notification type behaviors are configured primarily using the the `labels` field on this
	// resource. This block contains the labels which contain secrets or passwords so that they can be marked
	// sensitive and hidden from plan output. The name of the field, eg: password, will be the key
	// in the `labels` map in the api request.
	// Credentials may not be specified in both locations and will cause an error. Changing from one location
	// to a different credential configuration in the config will require an apply to update state.  Structure is documented below.
	SensitiveLabels NotificationChannelSensitiveLabelsPtrInput
	// -
	// (Required)
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	Type pulumi.StringInput
	// -
	// (Optional)
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (NotificationChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationChannelArgs)(nil)).Elem()
}
