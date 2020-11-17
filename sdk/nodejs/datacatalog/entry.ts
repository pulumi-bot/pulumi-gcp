// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Entry Metadata. A Data Catalog Entry resource represents another resource in Google Cloud Platform
 * (such as a BigQuery dataset or a Pub/Sub topic) or outside of Google Cloud Platform. Clients can use
 * the linkedResource field in the Entry resource to refer to the original resource ID of the source system.
 *
 * An Entry resource contains resource details, such as its schema. An Entry can also be used to attach
 * flexible metadata, such as a Tag.
 *
 * To get more information about Entry, see:
 *
 * * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
 *
 * ## Example Usage
 */
export class Entry extends pulumi.CustomResource {
    /**
     * Get an existing Entry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntryState, opts?: pulumi.CustomResourceOptions): Entry {
        return new Entry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/entry:Entry';

    /**
     * Returns true if the given object is an instance of Entry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Entry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Entry.__pulumiType;
    }

    /**
     * Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
     * https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
     */
    public /*out*/ readonly bigqueryDateShardedSpecs!: pulumi.Output<outputs.datacatalog.EntryBigqueryDateShardedSpec[]>;
    /**
     * Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
     */
    public /*out*/ readonly bigqueryTableSpecs!: pulumi.Output<outputs.datacatalog.EntryBigqueryTableSpec[]>;
    /**
     * Entry description, which can consist of several sentences or paragraphs that describe entry contents.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Display information such as title and description. A short name to identify the entry,
     * for example, "Analytics Data - Jan 2011".
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The name of the entry group this entry is in.
     */
    public readonly entryGroup!: pulumi.Output<string>;
    /**
     * The id of the entry to create.
     */
    public readonly entryId!: pulumi.Output<string>;
    /**
     * Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
     * Structure is documented below.
     */
    public readonly gcsFilesetSpec!: pulumi.Output<outputs.datacatalog.EntryGcsFilesetSpec | undefined>;
    /**
     * This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
     */
    public /*out*/ readonly integratedSystem!: pulumi.Output<string>;
    /**
     * The resource this metadata entry refers to.
     * For Google Cloud Platform resources, linkedResource is the full name of the resource.
     * For example, the linkedResource for a table resource from BigQuery is:
     * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
     * Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
     * this field is optional and defaults to an empty string.
     */
    public readonly linkedResource!: pulumi.Output<string>;
    /**
     * The Data Catalog resource name of the entry in URL format. Example:
     * projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
     * child resources may not actually be stored in the location in this name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
     * attached to it. See
     * https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
     * for what fields this schema can contain.
     */
    public readonly schema!: pulumi.Output<string | undefined>;
    /**
     * The type of the entry. Only used for Entries with types in the EntryType enum.
     * Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
     * Possible values are `FILESET`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * This field indicates the entry's source system that Data Catalog does not integrate with.
     * userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
     * and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    public readonly userSpecifiedSystem!: pulumi.Output<string | undefined>;
    /**
     * Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
     * When creating an entry, users should check the enum values first, if nothing matches the entry
     * to be created, then provide a custom value, for example "mySpecialType".
     * userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
     * numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    public readonly userSpecifiedType!: pulumi.Output<string | undefined>;

    /**
     * Create a Entry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntryArgs | EntryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EntryState | undefined;
            inputs["bigqueryDateShardedSpecs"] = state ? state.bigqueryDateShardedSpecs : undefined;
            inputs["bigqueryTableSpecs"] = state ? state.bigqueryTableSpecs : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["entryGroup"] = state ? state.entryGroup : undefined;
            inputs["entryId"] = state ? state.entryId : undefined;
            inputs["gcsFilesetSpec"] = state ? state.gcsFilesetSpec : undefined;
            inputs["integratedSystem"] = state ? state.integratedSystem : undefined;
            inputs["linkedResource"] = state ? state.linkedResource : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["userSpecifiedSystem"] = state ? state.userSpecifiedSystem : undefined;
            inputs["userSpecifiedType"] = state ? state.userSpecifiedType : undefined;
        } else {
            const args = argsOrState as EntryArgs | undefined;
            if (!args || args.entryGroup === undefined) {
                throw new Error("Missing required property 'entryGroup'");
            }
            if (!args || args.entryId === undefined) {
                throw new Error("Missing required property 'entryId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["entryGroup"] = args ? args.entryGroup : undefined;
            inputs["entryId"] = args ? args.entryId : undefined;
            inputs["gcsFilesetSpec"] = args ? args.gcsFilesetSpec : undefined;
            inputs["linkedResource"] = args ? args.linkedResource : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["userSpecifiedSystem"] = args ? args.userSpecifiedSystem : undefined;
            inputs["userSpecifiedType"] = args ? args.userSpecifiedType : undefined;
            inputs["bigqueryDateShardedSpecs"] = undefined /*out*/;
            inputs["bigqueryTableSpecs"] = undefined /*out*/;
            inputs["integratedSystem"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Entry.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Entry resources.
 */
export interface EntryState {
    /**
     * Specification for a group of BigQuery tables with name pattern [prefix]YYYYMMDD. Context:
     * https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding.
     */
    readonly bigqueryDateShardedSpecs?: pulumi.Input<pulumi.Input<inputs.datacatalog.EntryBigqueryDateShardedSpec>[]>;
    /**
     * Specification that applies to a BigQuery table. This is only valid on entries of type TABLE.
     */
    readonly bigqueryTableSpecs?: pulumi.Input<pulumi.Input<inputs.datacatalog.EntryBigqueryTableSpec>[]>;
    /**
     * Entry description, which can consist of several sentences or paragraphs that describe entry contents.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display information such as title and description. A short name to identify the entry,
     * for example, "Analytics Data - Jan 2011".
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the entry group this entry is in.
     */
    readonly entryGroup?: pulumi.Input<string>;
    /**
     * The id of the entry to create.
     */
    readonly entryId?: pulumi.Input<string>;
    /**
     * Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
     * Structure is documented below.
     */
    readonly gcsFilesetSpec?: pulumi.Input<inputs.datacatalog.EntryGcsFilesetSpec>;
    /**
     * This field indicates the entry's source system that Data Catalog integrates with, such as BigQuery or Pub/Sub.
     */
    readonly integratedSystem?: pulumi.Input<string>;
    /**
     * The resource this metadata entry refers to.
     * For Google Cloud Platform resources, linkedResource is the full name of the resource.
     * For example, the linkedResource for a table resource from BigQuery is:
     * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
     * Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
     * this field is optional and defaults to an empty string.
     */
    readonly linkedResource?: pulumi.Input<string>;
    /**
     * The Data Catalog resource name of the entry in URL format. Example:
     * projects/{project_id}/locations/{location}/entryGroups/{entryGroupId}/entries/{entryId}. Note that this Entry and its
     * child resources may not actually be stored in the location in this name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
     * attached to it. See
     * https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
     * for what fields this schema can contain.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * The type of the entry. Only used for Entries with types in the EntryType enum.
     * Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
     * Possible values are `FILESET`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * This field indicates the entry's source system that Data Catalog does not integrate with.
     * userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
     * and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    readonly userSpecifiedSystem?: pulumi.Input<string>;
    /**
     * Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
     * When creating an entry, users should check the enum values first, if nothing matches the entry
     * to be created, then provide a custom value, for example "mySpecialType".
     * userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
     * numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    readonly userSpecifiedType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Entry resource.
 */
export interface EntryArgs {
    /**
     * Entry description, which can consist of several sentences or paragraphs that describe entry contents.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Display information such as title and description. A short name to identify the entry,
     * for example, "Analytics Data - Jan 2011".
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The name of the entry group this entry is in.
     */
    readonly entryGroup: pulumi.Input<string>;
    /**
     * The id of the entry to create.
     */
    readonly entryId: pulumi.Input<string>;
    /**
     * Specification that applies to a Cloud Storage fileset. This is only valid on entries of type FILESET.
     * Structure is documented below.
     */
    readonly gcsFilesetSpec?: pulumi.Input<inputs.datacatalog.EntryGcsFilesetSpec>;
    /**
     * The resource this metadata entry refers to.
     * For Google Cloud Platform resources, linkedResource is the full name of the resource.
     * For example, the linkedResource for a table resource from BigQuery is:
     * //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
     * Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
     * this field is optional and defaults to an empty string.
     */
    readonly linkedResource?: pulumi.Input<string>;
    /**
     * Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema
     * attached to it. See
     * https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema
     * for what fields this schema can contain.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * The type of the entry. Only used for Entries with types in the EntryType enum.
     * Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType.
     * Possible values are `FILESET`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * This field indicates the entry's source system that Data Catalog does not integrate with.
     * userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
     * and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    readonly userSpecifiedSystem?: pulumi.Input<string>;
    /**
     * Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
     * When creating an entry, users should check the enum values first, if nothing matches the entry
     * to be created, then provide a custom value, for example "mySpecialType".
     * userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
     * numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
     */
    readonly userSpecifiedType?: pulumi.Input<string>;
}
