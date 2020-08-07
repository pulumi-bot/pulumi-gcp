// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud IoT Core device.
 *
 * To get more information about Device, see:
 *
 * * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/iot/docs/)
 *
 * ## Example Usage
 * ### Cloudiot Device Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const registry = new gcp.iot.Registry("registry", {});
 * const test_device = new gcp.iot.Device("test-device", {registry: registry.id});
 * ```
 * ### Cloudiot Device Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const registry = new gcp.iot.Registry("registry", {});
 * const test_device = new gcp.iot.Device("test-device", {
 *     registry: registry.id,
 *     credentials: [{
 *         publicKey: {
 *             format: "RSA_PEM",
 *             key: fs.readFileSync("test-fixtures/rsa_public.pem"),
 *         },
 *     }],
 *     blocked: false,
 *     logLevel: "INFO",
 *     metadata: {
 *         test_key_1: "test_value_1",
 *     },
 *     gatewayConfig: {
 *         gatewayType: "NON_GATEWAY",
 *     },
 * });
 * ```
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iot/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    public readonly blocked!: pulumi.Output<boolean | undefined>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
     */
    public /*out*/ readonly config!: pulumi.Output<outputs.iot.DeviceConfig>;
    /**
     * The credentials used to authenticate this device.  Structure is documented below.
     */
    public readonly credentials!: pulumi.Output<outputs.iot.DeviceCredential[] | undefined>;
    /**
     * Gateway-related configuration and state.  Structure is documented below.
     */
    public readonly gatewayConfig!: pulumi.Output<outputs.iot.DeviceGatewayConfig | undefined>;
    /**
     * The last time a cloud-to-device config version acknowledgment was received from the device.
     */
    public /*out*/ readonly lastConfigAckTime!: pulumi.Output<string>;
    /**
     * The last time a cloud-to-device config version was sent to the device.
     */
    public /*out*/ readonly lastConfigSendTime!: pulumi.Output<string>;
    /**
     * The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
     */
    public /*out*/ readonly lastErrorStatus!: pulumi.Output<outputs.iot.DeviceLastErrorStatus>;
    /**
     * The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
     */
    public /*out*/ readonly lastErrorTime!: pulumi.Output<string>;
    /**
     * The last time a telemetry event was received.
     */
    public /*out*/ readonly lastEventTime!: pulumi.Output<string>;
    /**
     * The last time an MQTT PINGREQ was received.
     */
    public /*out*/ readonly lastHeartbeatTime!: pulumi.Output<string>;
    /**
     * The last time a state event was received.
     */
    public /*out*/ readonly lastStateTime!: pulumi.Output<string>;
    /**
     * The logging verbosity for device activity.
     */
    public readonly logLevel!: pulumi.Output<string | undefined>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
     * unique.
     */
    public /*out*/ readonly numId!: pulumi.Output<string>;
    /**
     * The name of the device registry where this device should be created.
     */
    public readonly registry!: pulumi.Output<string>;
    /**
     * The state most recently received from the device.
     */
    public /*out*/ readonly state!: pulumi.Output<outputs.iot.DeviceState>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeviceState | undefined;
            inputs["blocked"] = state ? state.blocked : undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["gatewayConfig"] = state ? state.gatewayConfig : undefined;
            inputs["lastConfigAckTime"] = state ? state.lastConfigAckTime : undefined;
            inputs["lastConfigSendTime"] = state ? state.lastConfigSendTime : undefined;
            inputs["lastErrorStatus"] = state ? state.lastErrorStatus : undefined;
            inputs["lastErrorTime"] = state ? state.lastErrorTime : undefined;
            inputs["lastEventTime"] = state ? state.lastEventTime : undefined;
            inputs["lastHeartbeatTime"] = state ? state.lastHeartbeatTime : undefined;
            inputs["lastStateTime"] = state ? state.lastStateTime : undefined;
            inputs["logLevel"] = state ? state.logLevel : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numId"] = state ? state.numId : undefined;
            inputs["registry"] = state ? state.registry : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            if (!args || args.registry === undefined) {
                throw new Error("Missing required property 'registry'");
            }
            inputs["blocked"] = args ? args.blocked : undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["gatewayConfig"] = args ? args.gatewayConfig : undefined;
            inputs["logLevel"] = args ? args.logLevel : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["registry"] = args ? args.registry : undefined;
            inputs["config"] = undefined /*out*/;
            inputs["lastConfigAckTime"] = undefined /*out*/;
            inputs["lastConfigSendTime"] = undefined /*out*/;
            inputs["lastErrorStatus"] = undefined /*out*/;
            inputs["lastErrorTime"] = undefined /*out*/;
            inputs["lastEventTime"] = undefined /*out*/;
            inputs["lastHeartbeatTime"] = undefined /*out*/;
            inputs["lastStateTime"] = undefined /*out*/;
            inputs["numId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    readonly blocked?: pulumi.Input<boolean>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
     */
    readonly config?: pulumi.Input<inputs.iot.DeviceConfig>;
    /**
     * The credentials used to authenticate this device.  Structure is documented below.
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.DeviceCredential>[]>;
    /**
     * Gateway-related configuration and state.  Structure is documented below.
     */
    readonly gatewayConfig?: pulumi.Input<inputs.iot.DeviceGatewayConfig>;
    /**
     * The last time a cloud-to-device config version acknowledgment was received from the device.
     */
    readonly lastConfigAckTime?: pulumi.Input<string>;
    /**
     * The last time a cloud-to-device config version was sent to the device.
     */
    readonly lastConfigSendTime?: pulumi.Input<string>;
    /**
     * The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
     */
    readonly lastErrorStatus?: pulumi.Input<inputs.iot.DeviceLastErrorStatus>;
    /**
     * The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
     */
    readonly lastErrorTime?: pulumi.Input<string>;
    /**
     * The last time a telemetry event was received.
     */
    readonly lastEventTime?: pulumi.Input<string>;
    /**
     * The last time an MQTT PINGREQ was received.
     */
    readonly lastHeartbeatTime?: pulumi.Input<string>;
    /**
     * The last time a state event was received.
     */
    readonly lastStateTime?: pulumi.Input<string>;
    /**
     * The logging verbosity for device activity.
     */
    readonly logLevel?: pulumi.Input<string>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A server-defined unique numeric ID for the device. This is a more compact way to identify devices, and it is globally
     * unique.
     */
    readonly numId?: pulumi.Input<string>;
    /**
     * The name of the device registry where this device should be created.
     */
    readonly registry?: pulumi.Input<string>;
    /**
     * The state most recently received from the device.
     */
    readonly state?: pulumi.Input<inputs.iot.DeviceState>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    readonly blocked?: pulumi.Input<boolean>;
    /**
     * The credentials used to authenticate this device.  Structure is documented below.
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.DeviceCredential>[]>;
    /**
     * Gateway-related configuration and state.  Structure is documented below.
     */
    readonly gatewayConfig?: pulumi.Input<inputs.iot.DeviceGatewayConfig>;
    /**
     * The logging verbosity for device activity.
     */
    readonly logLevel?: pulumi.Input<string>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the device registry where this device should be created.
     */
    readonly registry: pulumi.Input<string>;
}
