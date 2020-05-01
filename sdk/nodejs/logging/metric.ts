// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Logs-based metric can also be used to extract values from logs and create a a distribution
 * of the values. The distribution records the statistics of the extracted values along with
 * an optional histogram of the values as specified by the bucket options.
 * 
 * 
 * To get more information about Metric, see:
 * 
 * * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/logging/docs/apis)
 * 
 * ## Example Usage - Logging Metric Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     bucketOptions: {
 *         linearBuckets: {
 *             numFiniteBuckets: 3,
 *             offset: 1,
 *             width: 1,
 *         },
 *     },
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     labelExtractors: {
 *         mass: "EXTRACT(jsonPayload.request)",
 *         sku: "EXTRACT(jsonPayload.id)",
 *     },
 *     metricDescriptor: {
 *         displayName: "My metric",
 *         labels: [
 *             {
 *                 description: "amount of matter",
 *                 key: "mass",
 *                 valueType: "STRING",
 *             },
 *             {
 *                 description: "Identifying number for item",
 *                 key: "sku",
 *                 valueType: "INT64",
 *             },
 *         ],
 *         metricKind: "DELTA",
 *         unit: "1",
 *         valueType: "DISTRIBUTION",
 *     },
 *     name: "my-(custom)/metric",
 *     valueExtractor: "EXTRACT(jsonPayload.request)",
 * });
 * ```
 * ## Example Usage - Logging Metric Counter Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     metricDescriptor: {
 *         metricKind: "DELTA",
 *         valueType: "INT64",
 *     },
 *     name: "my-(custom)/metric",
 * });
 * ```
 * ## Example Usage - Logging Metric Counter Labels
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     labelExtractors: {
 *         mass: "EXTRACT(jsonPayload.request)",
 *     },
 *     metricDescriptor: {
 *         labels: [{
 *             description: "amount of matter",
 *             key: "mass",
 *             valueType: "STRING",
 *         }],
 *         metricKind: "DELTA",
 *         valueType: "INT64",
 *     },
 *     name: "my-(custom)/metric",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_metric.html.markdown.
 */
export class Metric extends pulumi.CustomResource {
    /**
     * Get an existing Metric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricState, opts?: pulumi.CustomResourceOptions): Metric {
        return new Metric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/metric:Metric';

    /**
     * Returns true if the given object is an instance of Metric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Metric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Metric.__pulumiType;
    }

    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
     */
    public readonly bucketOptions!: pulumi.Output<outputs.logging.MetricBucketOptions | undefined>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    public readonly labelExtractors!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The metric descriptor associated with the logs-based metric.  Structure is documented below.
     */
    public readonly metricDescriptor!: pulumi.Output<outputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    public readonly valueExtractor!: pulumi.Output<string | undefined>;

    /**
     * Create a Metric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricArgs | MetricState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MetricState | undefined;
            inputs["bucketOptions"] = state ? state.bucketOptions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["labelExtractors"] = state ? state.labelExtractors : undefined;
            inputs["metricDescriptor"] = state ? state.metricDescriptor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["valueExtractor"] = state ? state.valueExtractor : undefined;
        } else {
            const args = argsOrState as MetricArgs | undefined;
            if (!args || args.filter === undefined) {
                throw new Error("Missing required property 'filter'");
            }
            if (!args || args.metricDescriptor === undefined) {
                throw new Error("Missing required property 'metricDescriptor'");
            }
            inputs["bucketOptions"] = args ? args.bucketOptions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["labelExtractors"] = args ? args.labelExtractors : undefined;
            inputs["metricDescriptor"] = args ? args.metricDescriptor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["valueExtractor"] = args ? args.valueExtractor : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Metric.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Metric resources.
 */
export interface MetricState {
    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
     */
    readonly bucketOptions?: pulumi.Input<inputs.logging.MetricBucketOptions>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    readonly labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The metric descriptor associated with the logs-based metric.  Structure is documented below.
     */
    readonly metricDescriptor?: pulumi.Input<inputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    readonly valueExtractor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Metric resource.
 */
export interface MetricArgs {
    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.  Structure is documented below.
     */
    readonly bucketOptions?: pulumi.Input<inputs.logging.MetricBucketOptions>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     */
    readonly filter: pulumi.Input<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    readonly labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The metric descriptor associated with the logs-based metric.  Structure is documented below.
     */
    readonly metricDescriptor: pulumi.Input<inputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    readonly valueExtractor?: pulumi.Input<string>;
}
