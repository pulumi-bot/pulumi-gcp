// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_notification_channel.html.markdown.
type NotificationChannel struct {
	s *pulumi.ResourceState
}

// NewNotificationChannel registers a new resource with the given unique name, arguments, and options.
func NewNotificationChannel(ctx *pulumi.Context,
	name string, args *NotificationChannelArgs, opts ...pulumi.ResourceOpt) (*NotificationChannel, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["displayName"] = nil
		inputs["enabled"] = nil
		inputs["labels"] = nil
		inputs["project"] = nil
		inputs["type"] = nil
		inputs["userLabels"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["displayName"] = args.DisplayName
		inputs["enabled"] = args.Enabled
		inputs["labels"] = args.Labels
		inputs["project"] = args.Project
		inputs["type"] = args.Type
		inputs["userLabels"] = args.UserLabels
	}
	inputs["name"] = nil
	inputs["verificationStatus"] = nil
	s, err := ctx.RegisterResource("gcp:monitoring/notificationChannel:NotificationChannel", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NotificationChannel{s: s}, nil
}

// GetNotificationChannel gets an existing NotificationChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NotificationChannelState, opts ...pulumi.ResourceOpt) (*NotificationChannel, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["displayName"] = state.DisplayName
		inputs["enabled"] = state.Enabled
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["type"] = state.Type
		inputs["userLabels"] = state.UserLabels
		inputs["verificationStatus"] = state.VerificationStatus
	}
	s, err := ctx.ReadResource("gcp:monitoring/notificationChannel:NotificationChannel", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NotificationChannel{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NotificationChannel) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NotificationChannel) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An optional human-readable description of this notification channel. This description may provide additional details,
// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
func (r *NotificationChannel) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
// is limited to 512 Unicode characters.
func (r *NotificationChannel) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
// notifications to a particular channel without removing the channel from all alerting policies that reference the
// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
// same set of alerting policies on the channel at some point in the future.
func (r *NotificationChannel) Enabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enabled"])
}

// Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the
// NotificationChannelDescriptor corresponding to the type field. **Note**: Some NotificationChannelDescriptor labels are
// sensitive and the API will return an partially-obfuscated value. For example, for '"type": "slack"' channels, an
// 'auth_token' label with value "SECRET" will be obfuscated as "**CRET". In order to avoid a diff, Terraform will use the
// state value if it appears that the obfuscated value matches the state value in length/unobfuscated characters. However,
// Terraform will not detect a diff if the obfuscated portion of the value was changed outside of Terraform.
func (r *NotificationChannel) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
// [CHANNEL_ID] is automatically assigned by the server on creation.
func (r *NotificationChannel) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *NotificationChannel) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
// valid values such as "email", "slack", etc...
func (r *NotificationChannel) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
// begin with a letter.
func (r *NotificationChannel) UserLabels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["userLabels"])
}

// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
// verification or that this specific channel has been exempted from verification because it was created prior to
// verification being required for channels of this type.This field cannot be modified using a standard
// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
func (r *NotificationChannel) VerificationStatus() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["verificationStatus"])
}

// Input properties used for looking up and filtering NotificationChannel resources.
type NotificationChannelState struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description interface{}
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and
	// unique name in order to make it easier to identify the channels in your project, though this is not enforced. The
	// display name is limited to 512 Unicode characters.
	DisplayName interface{}
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled interface{}
	// Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field. **Note**: Some NotificationChannelDescriptor labels are
	// sensitive and the API will return an partially-obfuscated value. For example, for '"type": "slack"' channels, an
	// 'auth_token' label with value "SECRET" will be obfuscated as "**CRET". In order to avoid a diff, Terraform will use the
	// state value if it appears that the obfuscated value matches the state value in length/unobfuscated characters. However,
	// Terraform will not detect a diff if the obfuscated portion of the value was changed outside of Terraform.
	Labels interface{}
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	// The [CHANNEL_ID] is automatically assigned by the server on creation.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type interface{}
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels interface{}
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus interface{}
}

// The set of arguments for constructing a NotificationChannel resource.
type NotificationChannelArgs struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description interface{}
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and
	// unique name in order to make it easier to identify the channels in your project, though this is not enforced. The
	// display name is limited to 512 Unicode characters.
	DisplayName interface{}
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled interface{}
	// Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field. **Note**: Some NotificationChannelDescriptor labels are
	// sensitive and the API will return an partially-obfuscated value. For example, for '"type": "slack"' channels, an
	// 'auth_token' label with value "SECRET" will be obfuscated as "**CRET". In order to avoid a diff, Terraform will use the
	// state value if it appears that the obfuscated value matches the state value in length/unobfuscated characters. However,
	// Terraform will not detect a diff if the obfuscated portion of the value was changed outside of Terraform.
	Labels interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type interface{}
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys
	// must begin with a letter.
	UserLabels interface{}
}
