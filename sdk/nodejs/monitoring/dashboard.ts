// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Stackdriver dashboard. Dashboards define the content and layout of pages in the Stackdriver web application.
 *
 * To get more information about Dashboards, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/dashboards)
 *
 * ## Example Usage
 *
 * ### Monitoring Dashboard Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dashboard = new gcp.monitoring.Dashboard("dashboard", {
 *     dashboardJson: `{
 *   "displayName": "Demo Dashboard",
 *   "gridLayout": {
 *     "widgets": [
 *       {
 *         "blank": {}
 *       }
 *     ]
 *   }
 * }
 *
 * `,
 * });
 * ```
 *
 * ### Monitoring Dashboard GridLayout
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dashboard = new gcp.monitoring.Dashboard("dashboard", {
 *     dashboardJson: `{
 *   "displayName": "Grid Layout Example",
 *   "gridLayout": {
 *     "columns": "2",
 *     "widgets": [
 *       {
 *         "title": "Widget 1",
 *         "xyChart": {
 *           "dataSets": [{
 *             "timeSeriesQuery": {
 *               "timeSeriesFilter": {
 *                 "filter": "metric.type=\\"agent.googleapis.com/nginx/connections/accepted_count\\"",
 *                 "aggregation": {
 *                   "perSeriesAligner": "ALIGN_RATE"
 *                 }
 *               },
 *               "unitOverride": "1"
 *             },
 *             "plotType": "LINE"
 *           }],
 *           "timeshiftDuration": "0s",
 *           "yAxis": {
 *             "label": "y1Axis",
 *             "scale": "LINEAR"
 *           }
 *         }
 *       },
 *       {
 *         "text": {
 *           "content": "Widget 2",
 *           "format": "MARKDOWN"
 *         }
 *       },
 *       {
 *         "title": "Widget 3",
 *         "xyChart": {
 *           "dataSets": [{
 *             "timeSeriesQuery": {
 *               "timeSeriesFilter": {
 *                 "filter": "metric.type=\\"agent.googleapis.com/nginx/connections/accepted_count\\"",
 *                 "aggregation": {
 *                   "perSeriesAligner": "ALIGN_RATE"
 *                 }
 *               },
 *               "unitOverride": "1"
 *             },
 *             "plotType": "STACKED_BAR"
 *           }],
 *           "timeshiftDuration": "0s",
 *           "yAxis": {
 *             "label": "y1Axis",
 *             "scale": "LINEAR"
 *           }
 *         }
 *       }
 *     ]
 *   }
 * }
 *
 * `,
 * });
 * ```
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    /**
     * The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
     * The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
     */
    public readonly dashboardJson!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DashboardState | undefined;
            inputs["dashboardJson"] = state ? state.dashboardJson : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if (!args || args.dashboardJson === undefined) {
                throw new Error("Missing required property 'dashboardJson'");
            }
            inputs["dashboardJson"] = args ? args.dashboardJson : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
     * The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
     */
    readonly dashboardJson?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * The JSON representation of a dashboard, following the format at https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
     * The representation of an existing dashboard can be found by using the [API Explorer](https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards/get)
     */
    readonly dashboardJson: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
