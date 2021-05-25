// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This message configures which resources and services to monitor for availability.
 *
 * To get more information about UptimeCheckConfig, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
 *
 * > **Warning:** All arguments including `http_check.auth_info.password` will be stored in the raw
 * state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
 *
 * ## Example Usage
 * ### Uptime Check Config Http
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
 *         body: "Zm9vJTI1M0RiYXI=",
 *         contentType: "URL_ENCODED",
 *         path: "/some-path",
 *         port: 8010,
 *         requestMethod: "POST",
 *     },
 *     monitoredResource: {
 *         labels: {
 *             host: "192.168.1.1",
 *             project_id: "my-project-name",
 *         },
 *         type: "uptime_url",
 *     },
 *     timeout: "60s",
 * });
 * ```
 * ### Uptime Check Config Https
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const https = new gcp.monitoring.UptimeCheckConfig("https", {
 *     contentMatchers: [{
 *         content: "example",
 *     }],
 *     displayName: "https-uptime-check",
 *     httpCheck: {
 *         path: "/some-path",
 *         port: 443,
 *         useSsl: true,
 *         validateSsl: true,
 *     },
 *     monitoredResource: {
 *         labels: {
 *             host: "192.168.1.1",
 *             project_id: "my-project-name",
 *         },
 *         type: "uptime_url",
 *     },
 *     timeout: "60s",
 * });
 * ```
 * ### Uptime Check Tcp
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
 *     timeout: "60s",
 *     tcpCheck: {
 *         port: 888,
 *     },
 *     resourceGroup: {
 *         resourceType: "INSTANCE",
 *         groupId: check.name,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * UptimeCheckConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig default {{name}}
 * ```
 */
export class UptimeCheckConfig extends pulumi.CustomResource {
    /**
     * Get an existing UptimeCheckConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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

    /**
     * The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
     * Structure is documented below.
     */
    public readonly contentMatchers!: pulumi.Output<outputs.monitoring.UptimeCheckConfigContentMatcher[] | undefined>;
    /**
     * A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Contains information needed to make an HTTP or HTTPS check.
     * Structure is documented below.
     */
    public readonly httpCheck!: pulumi.Output<outputs.monitoring.UptimeCheckConfigHttpCheck | undefined>;
    /**
     * The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
     * Structure is documented below.
     */
    public readonly monitoredResource!: pulumi.Output<outputs.monitoring.UptimeCheckConfigMonitoredResource | undefined>;
    /**
     * A unique resource name for this UptimeCheckConfig. The format is
     * projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
     */
    public readonly period!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The group resource associated with the configuration.
     * Structure is documented below.
     */
    public readonly resourceGroup!: pulumi.Output<outputs.monitoring.UptimeCheckConfigResourceGroup | undefined>;
    /**
     * The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
     */
    public readonly selectedRegions!: pulumi.Output<string[] | undefined>;
    /**
     * Contains information needed to make a TCP check.
     * Structure is documented below.
     */
    public readonly tcpCheck!: pulumi.Output<outputs.monitoring.UptimeCheckConfigTcpCheck | undefined>;
    /**
     * The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
     */
    public readonly timeout!: pulumi.Output<string>;
    /**
     * The id of the uptime check
     */
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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UptimeCheckConfigState | undefined;
            inputs["contentMatchers"] = state ? state.contentMatchers : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["httpCheck"] = state ? state.httpCheck : undefined;
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
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            inputs["contentMatchers"] = args ? args.contentMatchers : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["httpCheck"] = args ? args.httpCheck : undefined;
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
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UptimeCheckConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UptimeCheckConfig resources.
 */
export interface UptimeCheckConfigState {
    /**
     * The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
     * Structure is documented below.
     */
    contentMatchers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigContentMatcher>[]>;
    /**
     * A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Contains information needed to make an HTTP or HTTPS check.
     * Structure is documented below.
     */
    httpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigHttpCheck>;
    /**
     * The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
     * Structure is documented below.
     */
    monitoredResource?: pulumi.Input<inputs.monitoring.UptimeCheckConfigMonitoredResource>;
    /**
     * A unique resource name for this UptimeCheckConfig. The format is
     * projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].
     */
    name?: pulumi.Input<string>;
    /**
     * How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
     */
    period?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The group resource associated with the configuration.
     * Structure is documented below.
     */
    resourceGroup?: pulumi.Input<inputs.monitoring.UptimeCheckConfigResourceGroup>;
    /**
     * The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
     */
    selectedRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Contains information needed to make a TCP check.
     * Structure is documented below.
     */
    tcpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigTcpCheck>;
    /**
     * The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
     */
    timeout?: pulumi.Input<string>;
    /**
     * The id of the uptime check
     */
    uptimeCheckId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UptimeCheckConfig resource.
 */
export interface UptimeCheckConfigArgs {
    /**
     * The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
     * Structure is documented below.
     */
    contentMatchers?: pulumi.Input<pulumi.Input<inputs.monitoring.UptimeCheckConfigContentMatcher>[]>;
    /**
     * A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
     */
    displayName: pulumi.Input<string>;
    /**
     * Contains information needed to make an HTTP or HTTPS check.
     * Structure is documented below.
     */
    httpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigHttpCheck>;
    /**
     * The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks:  uptimeUrl  gceInstance  gaeApp  awsEc2Instance  awsElbLoadBalancer
     * Structure is documented below.
     */
    monitoredResource?: pulumi.Input<inputs.monitoring.UptimeCheckConfigMonitoredResource>;
    /**
     * How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
     */
    period?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The group resource associated with the configuration.
     * Structure is documented below.
     */
    resourceGroup?: pulumi.Input<inputs.monitoring.UptimeCheckConfigResourceGroup>;
    /**
     * The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
     */
    selectedRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Contains information needed to make a TCP check.
     * Structure is documented below.
     */
    tcpCheck?: pulumi.Input<inputs.monitoring.UptimeCheckConfigTcpCheck>;
    /**
     * The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration
     */
    timeout: pulumi.Input<string>;
}
