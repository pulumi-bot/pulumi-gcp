// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This message configures which resources and services to monitor for availability.
 * 
 * 
 * To get more information about UptimeCheckConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
 * 
 * ## Example Usage - Uptime Check Config Http
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const http = new gcp.monitoring.UptimeCheckConfig("http", {
 *     contentMatchers: [{
 *         content: "example",
 *     }],
 *     displayName: "http-uptime-check",
 *     httpCheck: {
 *         path: "/some-path",
 *         port: 8010,
 *     },
 *     monitoredResource: {
 *         labels: {
 *             host: "192.168.1.1",
 *             project_id: "example",
 *         },
 *         type: "uptimeUrl",
 *     },
 *     timeout: "60s",
 * });
 * ```
 * ## Example Usage - Uptime Check Tcp
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const check = new gcp.monitoring.Group("check", {
 *     displayName: "uptime-check-group",
 *     filter: "resource.metadata.name=has_substring(\"foo\")",
 * });
 * const tcpGroup = new gcp.monitoring.UptimeCheckConfig("tcpGroup", {
 *     displayName: "tcp-uptime-check",
 *     resourceGroup: {
 *         groupId: check.name,
 *         resourceType: "INSTANCE",
 *     },
 *     tcpCheck: {
 *         port: 888,
 *     },
 *     timeout: "60s",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/monitoring_uptime_check_config.html.markdown.
 */
export class UptimeCheckConfig extends pulumi.CustomResource {
    /**
     * Get an existing UptimeCheckConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UptimeCheckConfigState, opts?: pulumi.CustomResourceOptions): UptimeCheckConfig {
        return new UptimeCheckConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig';

    /**
     * Returns true if the given object is an instance of UptimeCheckConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UptimeCheckConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UptimeCheckConfig.__pulumiType;
    }

    public readonly contentMatchers!: pulumi.Output<outputs.monitoring.UptimeCheckConfigContentMatcher[] | undefined>;
    public readonly displayName!: pulumi.Output<string>;
    public readonly httpCheck!: pulumi.Output<outputs.monitoring.UptimeCheckConfigHttpCheck | undefined>;
    public readonly internalCheckers!: pulumi.Output<outputs.monitoring.UptimeCheckConfigInternalChecker[]>;
    public readonly isInternal!: pulumi.Output<boolean>;
    public readonly monitoredResource!: pulumi.Output<outputs.monitoring.UptimeCheckConfigMonitoredResource | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly resourceGroup!: pulumi.Output<outputs.monitoring.UptimeCheckConfigResourceGroup | undefined>;
    public readonly selectedRegions!: pulumi.Output<string[] | undefined>;
    public readonly tcpCheck!: pulumi.Output<outputs.monitoring.UptimeCheckConfigTcpCheck | undefined>;
    public readonly timeout!: pulumi.Output<string>;
    public /*out*/ readonly uptimeCheckId!: pulumi.Output<string>;

    /**
     * Create a UptimeCheckConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UptimeCheckConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UptimeCheckConfigArgs | UptimeCheckConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UptimeCheckConfigState | undefined;
            inputs["contentMatchers"] = state ? state.contentMatchers : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["httpCheck"] = state ? state.httpCheck : undefined;
            inputs["internalCheckers"] = state ? state.internalCheckers : undefined;
            inputs["isInternal"] = state ? state.isInternal : undefined;
            inputs["monitoredResource"] = state ? state.monitoredResource : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["resourceGroup"] = state ? state.resourceGroup : undefined;
            inputs["selectedRegions"] = state ? state.selectedRegions : undefined;
            inputs["tcpCheck"] = state ? state.tcpCheck : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["uptimeCheckId"] = state ? state.uptimeCheckId : undefined;
        } else {
            const args = argsOrState as UptimeCheckConfigArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.timeout === undefined) {
                throw new Error("Missing required property 'timeout'");
            }
            inputs["contentMatchers"] = args ? args.contentMatchers : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["httpCheck"] = args ? args.httpCheck : undefined;
            inputs["internalCheckers"] = args ? args.internalCheckers : undefined;
            inputs["isInternal"] = args ? args.isInternal : undefined;
            inputs["monitoredResource"] = args ? args.monitoredResource : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["selectedRegions"] = args ? args.selectedRegions : undefined;
            inputs["tcpCheck"] = args ? args.tcpCheck : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["uptimeCheckId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UptimeCheckConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UptimeCheckConfig resources.
 */
export interface UptimeCheckConfigState {
    readonly contentMatchers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigContentMatcher>[]>;
    readonly displayName?: pulumi.Input<string>;
    readonly httpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigHttpCheck>;
    readonly internalCheckers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigInternalChecker>[]>;
    readonly isInternal?: pulumi.Input<boolean>;
    readonly monitoredResource?: pulumi.Input<inputs.monitoring.UptimeCheckConfigMonitoredResource>;
    readonly name?: pulumi.Input<string>;
    readonly period?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly resourceGroup?: pulumi.Input<inputs.monitoring.UptimeCheckConfigResourceGroup>;
    readonly selectedRegions?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tcpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigTcpCheck>;
    readonly timeout?: pulumi.Input<string>;
    readonly uptimeCheckId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UptimeCheckConfig resource.
 */
export interface UptimeCheckConfigArgs {
    readonly contentMatchers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigContentMatcher>[]>;
    readonly displayName: pulumi.Input<string>;
    readonly httpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigHttpCheck>;
    readonly internalCheckers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigInternalChecker>[]>;
    readonly isInternal?: pulumi.Input<boolean>;
    readonly monitoredResource?: pulumi.Input<inputs.monitoring.UptimeCheckConfigMonitoredResource>;
    readonly period?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly resourceGroup?: pulumi.Input<inputs.monitoring.UptimeCheckConfigResourceGroup>;
    readonly selectedRegions?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tcpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigTcpCheck>;
    readonly timeout: pulumi.Input<string>;
}
