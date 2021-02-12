// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A NotificationChannel is a medium through which an alert is delivered
 * when a policy violation is detected. Examples of channels include email, SMS,
 * and third-party messaging applications. Fields containing sensitive information
 * like authentication tokens or contact info are only partially populated on retrieval.
 *
 * Notification Channels are designed to be flexible and are made up of a supported `type`
 * and labels to configure that channel. Each `type` has specific labels that need to be
 * present for that channel to be correctly configured. The labels that are required to be
 * present for one channel `type` are often different than those required for another.
 * Due to these loose constraints it's often best to set up a channel through the UI
 * and import it to the provider when setting up a brand new channel type to determine which
 * labels are required.
 *
 * A list of supported channels per project the `list` endpoint can be
 * accessed programmatically or through the api explorer at  https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list .
 * This provides the channel type and all of the required labels that must be passed.
 *
 * To get more information about NotificationChannel, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
 * * How-to Guides
 *     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 * ## Example Usage
 * ### Notification Channel Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.monitoring.NotificationChannel("basic", {
 *     displayName: "Test Notification Channel",
 *     labels: {
 *         email_address: "fake_email@blahblah.com",
 *     },
 *     type: "email",
 * });
 * ```
 * ### Notification Channel Sensitive
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultNotificationChannel = new gcp.monitoring.NotificationChannel("default", {
 *     displayName: "Test Slack Channel",
 *     labels: {
 *         channel_name: "#foobar",
 *     },
 *     sensitiveLabels: {
 *         authToken: "one",
 *     },
 *     type: "slack",
 * });
 * ```
 *
 * ## Import
 *
 * NotificationChannel can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/notificationChannel:NotificationChannel default {{name}}
 * ```
 */
export class NotificationChannel extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationChannelState, opts?: pulumi.CustomResourceOptions): NotificationChannel {
        return new NotificationChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/notificationChannel:NotificationChannel';

    /**
     * Returns true if the given object is an instance of NotificationChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannel.__pulumiType;
    }

    /**
     * An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration fields that define the channel and its behavior. The
     * permissible and required labels are specified in the
     * NotificationChannelDescriptor corresponding to the type field.
     * Labels with sensitive data are obfuscated by the API and therefore the provider cannot
     * determine if there are upstream changes to these fields. They can also be configured via
     * the sensitiveLabels block, but cannot be configured in both places.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
     * [CHANNEL_ID] is automatically assigned by the server on creation.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Different notification type behaviors are configured primarily using the the `labels` field on this
     * resource. This block contains the labels which contain secrets or passwords so that they can be marked
     * sensitive and hidden from plan output. The name of the field, eg: password, will be the key
     * in the `labels` map in the api request.
     * Credentials may not be specified in both locations and will cause an error. Changing from one location
     * to a different credential configuration in the config will require an apply to update state.
     * Structure is documented below.
     */
    public readonly sensitiveLabels!: pulumi.Output<outputs.monitoring.NotificationChannelSensitiveLabels | undefined>;
    /**
     * The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    public readonly userLabels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
     * operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
     * non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
     * works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
     * verification or that this specific channel has been exempted from verification because it was created prior to
     * verification being required for channels of this type.This field cannot be modified using a standard
     * UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
     */
    public /*out*/ readonly verificationStatus!: pulumi.Output<string>;

    /**
     * Create a NotificationChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationChannelArgs | NotificationChannelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationChannelState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["sensitiveLabels"] = state ? state.sensitiveLabels : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["userLabels"] = state ? state.userLabels : undefined;
            inputs["verificationStatus"] = state ? state.verificationStatus : undefined;
        } else {
            const args = argsOrState as NotificationChannelArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["sensitiveLabels"] = args ? args.sensitiveLabels : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["userLabels"] = args ? args.userLabels : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["verificationStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationChannel resources.
 */
export interface NotificationChannelState {
    /**
     * An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Configuration fields that define the channel and its behavior. The
     * permissible and required labels are specified in the
     * NotificationChannelDescriptor corresponding to the type field.
     * Labels with sensitive data are obfuscated by the API and therefore the provider cannot
     * determine if there are upstream changes to these fields. They can also be configured via
     * the sensitiveLabels block, but cannot be configured in both places.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
     * [CHANNEL_ID] is automatically assigned by the server on creation.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Different notification type behaviors are configured primarily using the the `labels` field on this
     * resource. This block contains the labels which contain secrets or passwords so that they can be marked
     * sensitive and hidden from plan output. The name of the field, eg: password, will be the key
     * in the `labels` map in the api request.
     * Credentials may not be specified in both locations and will cause an error. Changing from one location
     * to a different credential configuration in the config will require an apply to update state.
     * Structure is documented below.
     */
    readonly sensitiveLabels?: pulumi.Input<inputs.monitoring.NotificationChannelSensitiveLabels>;
    /**
     * The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
     */
    readonly type?: pulumi.Input<string>;
    /**
     * User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    readonly userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
     * operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
     * non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
     * works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
     * verification or that this specific channel has been exempted from verification because it was created prior to
     * verification being required for channels of this type.This field cannot be modified using a standard
     * UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
     */
    readonly verificationStatus?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotificationChannel resource.
 */
export interface NotificationChannelArgs {
    /**
     * An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Configuration fields that define the channel and its behavior. The
     * permissible and required labels are specified in the
     * NotificationChannelDescriptor corresponding to the type field.
     * Labels with sensitive data are obfuscated by the API and therefore the provider cannot
     * determine if there are upstream changes to these fields. They can also be configured via
     * the sensitiveLabels block, but cannot be configured in both places.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Different notification type behaviors are configured primarily using the the `labels` field on this
     * resource. This block contains the labels which contain secrets or passwords so that they can be marked
     * sensitive and hidden from plan output. The name of the field, eg: password, will be the key
     * in the `labels` map in the api request.
     * Credentials may not be specified in both locations and will cause an error. Changing from one location
     * to a different credential configuration in the config will require an apply to update state.
     * Structure is documented below.
     */
    readonly sensitiveLabels?: pulumi.Input<inputs.monitoring.NotificationChannelSensitiveLabels>;
    /**
     * The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
     */
    readonly type: pulumi.Input<string>;
    /**
     * User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
     */
    readonly userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
