// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ProjectSink extends pulumi.CustomResource {
    /**
     * Get an existing ProjectSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * {{% examples %}}
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * {{% /examples %}}
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * The name of the logging sink.
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
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `uniqueWriterIdentity` to true.
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
        if (opts && opts.id) {
            const state = argsOrState as ProjectSinkState | undefined;
            inputs["bigqueryOptions"] = state ? state.bigqueryOptions : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["uniqueWriterIdentity"] = state ? state.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as ProjectSinkArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            inputs["bigqueryOptions"] = args ? args.bigqueryOptions : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["uniqueWriterIdentity"] = args ? args.uniqueWriterIdentity : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
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
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * {{% examples %}}
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * {{% /examples %}}
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination?: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
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
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `uniqueWriterIdentity` to true.
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
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * {{% examples %}}
     * ```typescript
     * import * as pulumi from "@pulumi/pulumi";
     * ```
     * {{% /examples %}}
     * The writer associated with the sink must have access to write to the above resource.
     */
    readonly destination: pulumi.Input<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    readonly filter?: pulumi.Input<string>;
    /**
     * The name of the logging sink.
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
     * then a unique service account is created and used for this sink. If you wish to publish logs across projects, you
     * must set `uniqueWriterIdentity` to true.
     */
    readonly uniqueWriterIdentity?: pulumi.Input<boolean>;
}
