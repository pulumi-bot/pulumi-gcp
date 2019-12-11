// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/data_fusion_instance.html.markdown.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datafusion/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * An optional description of the instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Option to enable Stackdriver Logging.
     */
    public readonly enableStackdriverLogging!: pulumi.Output<boolean | undefined>;
    /**
     * Option to enable Stackdriver Monitoring.
     */
    public readonly enableStackdriverMonitoring!: pulumi.Output<boolean | undefined>;
    /**
     * The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine VMs.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network configuration options. These are required when a private Data Fusion instance is to be created.
     */
    public readonly networkConfig!: pulumi.Output<outputs.datafusion.InstanceNetworkConfig | undefined>;
    /**
     * Map of additional options used to configure the behavior of Data Fusion instance.
     */
    public readonly options!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have
     * private IP addresses and will not be able to access the public internet.
     */
    public readonly privateInstance!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region of the Data Fusion instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Endpoint on which the Data Fusion UI and REST APIs are accessible.
     */
    public /*out*/ readonly serviceEndpoint!: pulumi.Output<string>;
    /**
     * The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is running
     * and ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted - UPGRADING:
     * Instance is being upgraded - RESTARTING: Instance is being restarted
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of this Data Fusion instance if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * Represents the type of Data Fusion instance. Each type is configured with the default settings for processing and
     * memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using
     * point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support
     * for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have
     * more features available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Current version of the Data Fusion.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableStackdriverLogging"] = state ? state.enableStackdriverLogging : undefined;
            inputs["enableStackdriverMonitoring"] = state ? state.enableStackdriverMonitoring : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkConfig"] = state ? state.networkConfig : undefined;
            inputs["options"] = state ? state.options : undefined;
            inputs["privateInstance"] = state ? state.privateInstance : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceEndpoint"] = state ? state.serviceEndpoint : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["stateMessage"] = state ? state.stateMessage : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["enableStackdriverLogging"] = args ? args.enableStackdriverLogging : undefined;
            inputs["enableStackdriverMonitoring"] = args ? args.enableStackdriverMonitoring : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkConfig"] = args ? args.networkConfig : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["privateInstance"] = args ? args.privateInstance : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["serviceEndpoint"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateMessage"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * An optional description of the instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Option to enable Stackdriver Logging.
     */
    readonly enableStackdriverLogging?: pulumi.Input<boolean>;
    /**
     * Option to enable Stackdriver Monitoring.
     */
    readonly enableStackdriverMonitoring?: pulumi.Input<boolean>;
    /**
     * The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine VMs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network configuration options. These are required when a private Data Fusion instance is to be created.
     */
    readonly networkConfig?: pulumi.Input<inputs.datafusion.InstanceNetworkConfig>;
    /**
     * Map of additional options used to configure the behavior of Data Fusion instance.
     */
    readonly options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have
     * private IP addresses and will not be able to access the public internet.
     */
    readonly privateInstance?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region of the Data Fusion instance.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Endpoint on which the Data Fusion UI and REST APIs are accessible.
     */
    readonly serviceEndpoint?: pulumi.Input<string>;
    /**
     * The current state of this Data Fusion instance. - CREATING: Instance is being created - RUNNING: Instance is running
     * and ready for requests - FAILED: Instance creation failed - DELETING: Instance is being deleted - UPGRADING:
     * Instance is being upgraded - RESTARTING: Instance is being restarted
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Additional information about the current state of this Data Fusion instance if available.
     */
    readonly stateMessage?: pulumi.Input<string>;
    /**
     * Represents the type of Data Fusion instance. Each type is configured with the default settings for processing and
     * memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using
     * point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support
     * for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have
     * more features available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * Current version of the Data Fusion.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * An optional description of the instance.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Option to enable Stackdriver Logging.
     */
    readonly enableStackdriverLogging?: pulumi.Input<boolean>;
    /**
     * Option to enable Stackdriver Monitoring.
     */
    readonly enableStackdriverMonitoring?: pulumi.Input<boolean>;
    /**
     * The resource labels for instance to use to annotate any related underlying resources, such as Compute Engine VMs.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Network configuration options. These are required when a private Data Fusion instance is to be created.
     */
    readonly networkConfig?: pulumi.Input<inputs.datafusion.InstanceNetworkConfig>;
    /**
     * Map of additional options used to configure the behavior of Data Fusion instance.
     */
    readonly options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have
     * private IP addresses and will not be able to access the public internet.
     */
    readonly privateInstance?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region of the Data Fusion instance.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Represents the type of Data Fusion instance. Each type is configured with the default settings for processing and
     * memory. - BASIC: Basic Data Fusion instance. In Basic type, the user will be able to create data pipelines using
     * point and click UI. However, there are certain limitations, such as fewer number of concurrent pipelines, no support
     * for streaming pipelines, etc. - ENTERPRISE: Enterprise Data Fusion instance. In Enterprise type, the user will have
     * more features available, such as support for streaming pipelines, higher number of concurrent pipelines, etc.
     */
    readonly type: pulumi.Input<string>;
}
