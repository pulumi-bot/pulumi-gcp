// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a folder-level logging sink. For more information see
 * [the official documentation](https://cloud.google.com/logging/docs/) and
 * [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
 * 
 * Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
 * granted to the credentials used with terraform.
 */
export class FolderSink extends pulumi.CustomResource {
    /**
     * Get an existing FolderSink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderSinkState, opts?: pulumi.CustomResourceOptions): FolderSink {
        return new FolderSink(name, <any>state, { ...opts, id: id });
    }

    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
     * The writer associated with the sink must have access to write to the above resource.
     */
    public readonly destination: pulumi.Output<string>;
    /**
     * The filter to apply when exporting logs. Only log entries that match the filter are exported.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
     * write a filter.
     */
    public readonly filter: pulumi.Output<string | undefined>;
    /**
     * The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
     * accepted.
     */
    public readonly folder: pulumi.Output<string>;
    /**
     * Whether or not to include children folders in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
     */
    public readonly includeChildren: pulumi.Output<boolean | undefined>;
    /**
     * The name of the logging sink.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    public /*out*/ readonly writerIdentity: pulumi.Output<string>;

    /**
     * Create a FolderSink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderSinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FolderSinkArgs | FolderSinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: FolderSinkState = argsOrState as FolderSinkState | undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["includeChildren"] = state ? state.includeChildren : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["writerIdentity"] = state ? state.writerIdentity : undefined;
        } else {
            const args = argsOrState as FolderSinkArgs | undefined;
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.folder === undefined) {
                throw new Error("Missing required property 'folder'");
            }
            inputs["destination"] = args ? args.destination : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["includeChildren"] = args ? args.includeChildren : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["writerIdentity"] = undefined /*out*/;
        }
        super("gcp:logging/folderSink:FolderSink", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FolderSink resources.
 */
export interface FolderSinkState {
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
     * The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
     * accepted.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * Whether or not to include children folders in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
     */
    readonly includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The identity associated with this sink. This identity must be granted write access to the
     * configured `destination`.
     */
    readonly writerIdentity?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FolderSink resource.
 */
export interface FolderSinkArgs {
    /**
     * The destination of the sink (or, in other words, where logs are written to). Can be a
     * Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
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
     * The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is
     * accepted.
     */
    readonly folder: pulumi.Input<string>;
    /**
     * Whether or not to include children folders in the sink export. If true, logs
     * associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
     */
    readonly includeChildren?: pulumi.Input<boolean>;
    /**
     * The name of the logging sink.
     */
    readonly name?: pulumi.Input<string>;
}
