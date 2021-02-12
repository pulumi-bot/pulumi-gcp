// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a project-level logging sink. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/),
 * [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs)
 * and
 * [API](https://cloud.google.com/logging/docs/reference/v2/rest/).
 *
 * > **Note:** You must have [granted the "Logs Configuration Writer"](https://cloud.google.com/logging/docs/access-control) IAM role (`roles/logging.configWriter`) to the credentials used with this provider.
 *
 * > **Note** You must [enable the Cloud Resource Manager API](https://console.cloud.google.com/apis/library/cloudresourcemanager.googleapis.com)
 *
 * ## Import
 *
 * Project-level logging sinks can be imported using their URI, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:logging/projectSink:ProjectSink my_sink projects/my-project/sinks/my-sink
 * ```
 */
export class ProjectSink extends pulumi.CustomResource {
    /**
     * Get an existing ProjectSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectSinkState, opts?: pulumi.CustomResourceOptions): ProjectSink {
        return new ProjectSink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/projectSink:ProjectSink';

    /**
     * Returns true if the given object is an instance of ProjectSink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectSink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectSink.__pulumiType;
    }

    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    public readonly bigqueryOptions!: pulumi.Output<outputs.logging.ProjectSinkBigqueryOptions>;
    /**
     * A description of this exclusion.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * If set to True, then this exclusion is disabled and it does not exclude any log entries.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    public readonly exclusions!: pulumi.Output<outputs.logging.ProjectSinkExclusion[] | undefined>;
    /**
     * An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
     * `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
     */
    public readonly uniqueWriterIdentity!: pulumi.Output<boolean | undefined>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity!: pulumi.Output<string>;

    /**
     * Create a ProjectSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectSinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectSinkArgs | ProjectSinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectSinkState | undefined;
            inputs["bigqueryOptions"] = state ? state.bigqueryOptions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["exclusions"] = state ? state.exclusions : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["uniqueWriterIdentity"] = state ? state.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as ProjectSinkArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["bigqueryOptions"] = args ? args.bigqueryOptions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["exclusions"] = args ? args.exclusions : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["uniqueWriterIdentity"] = args ? args.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProjectSink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectSink resources.
 */
export interface ProjectSinkState {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    readonly bigqueryOptions?: pulumi.Input<inputs.logging.ProjectSinkBigqueryOptions>;
    /**
     * A description of this exclusion.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * If set to True, then this exclusion is disabled and it does not exclude any log entries.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    readonly exclusions?: pulumi.Input<pulumi.Input<inputs.logging.ProjectSinkExclusion>[]>;
    /**
     * An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
     * `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
     */
    readonly uniqueWriterIdentity?: pulumi.Input<boolean>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    readonly writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectSink resource.
 */
export interface ProjectSinkArgs {
    /**
     * Options that affect sinks exporting data to BigQuery. Structure documented below.
     */
    readonly bigqueryOptions?: pulumi.Input<inputs.logging.ProjectSinkBigqueryOptions>;
    /**
     * A description of this exclusion.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket . Examples:
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * If set to True, then this exclusion is disabled and it does not exclude any log entries.
     */
    readonly disabled?: pulumi.Input<boolean>;
    /**
     * Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both filter and one of exclusionFilters it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
     */
    readonly exclusions?: pulumi.Input<pulumi.Input<inputs.logging.ProjectSinkExclusion>[]>;
    /**
     * An advanced logs filter that matches the log entries to be excluded. By using the sample function, you can exclude less than 100% of the matching log entries. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * A client-assigned identifier, such as `load-balancer-exclusion`. Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project to create the sink in. If omitted, the project associated with the provider is
     * used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Whether or not to create a unique identity associated with this sink. If `false`
     * (the default), then the `writerIdentity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true`,
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
     * `bigqueryOptions`, you must set `uniqueWriterIdentity` to true.
     */
    readonly uniqueWriterIdentity?: pulumi.Input<boolean>;
}
