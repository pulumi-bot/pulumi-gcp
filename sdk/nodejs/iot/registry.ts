// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 *  Creates a device registry in Google's Cloud IoT Core platform. For more information see
 * [the official documentation](https://cloud.google.com/iot/docs/) and
 * [API](https://cloud.google.com/iot/docs/reference/cloudiot/rest/v1/projects.locations.registries).
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const defaultDevicestatus = new gcp.pubsub.Topic("default-devicestatus", {});
 * const defaultTelemetry = new gcp.pubsub.Topic("default-telemetry", {});
 * const defaultRegistry = new gcp.iot.Registry("default-registry", {
 *     event_notification_configs: [{
 *         pubsubTopicName: default_telemetry.id,
 *     }],
 *     stateNotificationConfig: {
 *         pubsub_topic_name: default_devicestatus.id,
 *     },
 *     httpConfig: {
 *         http_enabled_state: "HTTP_ENABLED",
 *     },
 *     mqttConfig: {
 *         mqtt_enabled_state: "MQTT_ENABLED",
 *     },
 *     credentials: [{
 *         publicKeyCertificate: {
 *             format: "X509_CERTIFICATE_PEM",
 *             certificate: fs.readFileSync("rsa_cert.pem"),
 *         },
 *     }],
 * });
 * ```
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryState, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iot/registry:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * List of public key certificates to authenticate devices. Structure is documented below. 
     */
    public readonly credentials!: pulumi.Output<outputs.iot.RegistryCredential[] | undefined>;
    /**
     * List of configurations for event notification, such as
     * PubSub topics to publish device events to. Structure is documented below.
     */
    public readonly eventNotificationConfigs!: pulumi.Output<outputs.iot.RegistryEventNotificationConfigItem[]>;
    /**
     * Activate or deactivate HTTP. Structure is documented below.
     */
    public readonly httpConfig!: pulumi.Output<outputs.iot.RegistryHttpConfig>;
    public readonly logLevel!: pulumi.Output<string | undefined>;
    /**
     * Activate or deactivate MQTT. Structure is documented below.
     */
    public readonly mqttConfig!: pulumi.Output<outputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     * Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A PubSub topic to publish device state updates. Structure is documented below.
     */
    public readonly stateNotificationConfig!: pulumi.Output<outputs.iot.RegistryStateNotificationConfig | undefined>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs | RegistryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegistryState | undefined;
            inputs["credentials"] = state ? state.credentials : undefined;
            inputs["eventNotificationConfigs"] = state ? state.eventNotificationConfigs : undefined;
            inputs["httpConfig"] = state ? state.httpConfig : undefined;
            inputs["logLevel"] = state ? state.logLevel : undefined;
            inputs["mqttConfig"] = state ? state.mqttConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["stateNotificationConfig"] = state ? state.stateNotificationConfig : undefined;
        } else {
            const args = argsOrState as RegistryArgs | undefined;
            inputs["credentials"] = args ? args.credentials : undefined;
            inputs["eventNotificationConfigs"] = args ? args.eventNotificationConfigs : undefined;
            inputs["httpConfig"] = args ? args.httpConfig : undefined;
            inputs["logLevel"] = args ? args.logLevel : undefined;
            inputs["mqttConfig"] = args ? args.mqttConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["stateNotificationConfig"] = args ? args.stateNotificationConfig : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "gcp:kms/registry:Registry" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Registry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registry resources.
 */
export interface RegistryState {
    /**
     * List of public key certificates to authenticate devices. Structure is documented below. 
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.RegistryCredential>[]>;
    /**
     * List of configurations for event notification, such as
     * PubSub topics to publish device events to. Structure is documented below.
     */
    readonly eventNotificationConfigs?: pulumi.Input<pulumi.Input<inputs.iot.RegistryEventNotificationConfigItem>[]>;
    /**
     * Activate or deactivate HTTP. Structure is documented below.
     */
    readonly httpConfig?: pulumi.Input<inputs.iot.RegistryHttpConfig>;
    readonly logLevel?: pulumi.Input<string>;
    /**
     * Activate or deactivate MQTT. Structure is documented below.
     */
    readonly mqttConfig?: pulumi.Input<inputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A PubSub topic to publish device state updates. Structure is documented below.
     */
    readonly stateNotificationConfig?: pulumi.Input<inputs.iot.RegistryStateNotificationConfig>;
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * List of public key certificates to authenticate devices. Structure is documented below. 
     */
    readonly credentials?: pulumi.Input<pulumi.Input<inputs.iot.RegistryCredential>[]>;
    /**
     * List of configurations for event notification, such as
     * PubSub topics to publish device events to. Structure is documented below.
     */
    readonly eventNotificationConfigs?: pulumi.Input<pulumi.Input<inputs.iot.RegistryEventNotificationConfigItem>[]>;
    /**
     * Activate or deactivate HTTP. Structure is documented below.
     */
    readonly httpConfig?: pulumi.Input<inputs.iot.RegistryHttpConfig>;
    readonly logLevel?: pulumi.Input<string>;
    /**
     * Activate or deactivate MQTT. Structure is documented below.
     */
    readonly mqttConfig?: pulumi.Input<inputs.iot.RegistryMqttConfig>;
    /**
     * A unique name for the resource, required by device registry.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The Region in which the created address should reside. If it is not provided, the provider region is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A PubSub topic to publish device state updates. Structure is documented below.
     */
    readonly stateNotificationConfig?: pulumi.Input<inputs.iot.RegistryStateNotificationConfig>;
}
