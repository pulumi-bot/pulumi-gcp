// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Gives dataset access for a single entity. This resource is intended to be used in cases where
 * it is not possible to compile a full list of access blocks to include in a
 * `gcp.bigquery.Dataset` resource, to enable them to be added separately.
 *
 * > **Note:** If this resource is used alongside a `gcp.bigquery.Dataset` resource, the
 * dataset resource must either have no defined `access` blocks or a `lifecycle` block with
 * `ignoreChanges = [access]` so they don't fight over which accesses should be on the dataset.
 *
 * To get more information about DatasetAccess, see:
 *
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets)
 * * How-to Guides
 *     * [Controlling access to datasets](https://cloud.google.com/bigquery/docs/dataset-access-controls)
 *
 * ## Example Usage
 * ### Bigquery Dataset Access Basic User
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dataset = new gcp.bigquery.Dataset("dataset", {datasetId: "example_dataset"});
 * const bqowner = new gcp.serviceAccount.Account("bqowner", {accountId: "bqowner"});
 * const access = new gcp.bigquery.DatasetAccess("access", {
 *     datasetId: dataset.datasetId,
 *     role: "OWNER",
 *     userByEmail: bqowner.email,
 * });
 * ```
 * ### Bigquery Dataset Access View
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _private = new gcp.bigquery.Dataset("private", {datasetId: "example_dataset"});
 * const publicDataset = new gcp.bigquery.Dataset("publicDataset", {datasetId: "example_dataset2"});
 * const publicTable = new gcp.bigquery.Table("publicTable", {
 *     datasetId: publicDataset.datasetId,
 *     tableId: "example_table",
 *     view: {
 *         query: "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
 *         useLegacySql: false,
 *     },
 * });
 * const access = new gcp.bigquery.DatasetAccess("access", {
 *     datasetId: _private.datasetId,
 *     view: {
 *         projectId: publicTable.project,
 *         datasetId: publicDataset.datasetId,
 *         tableId: publicTable.tableId,
 *     },
 * });
 * ```
 */
export class DatasetAccess extends pulumi.CustomResource {
    /**
     * Get an existing DatasetAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetAccessState, opts?: pulumi.CustomResourceOptions): DatasetAccess {
        return new DatasetAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/datasetAccess:DatasetAccess';

    /**
     * Returns true if the given object is an instance of DatasetAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatasetAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatasetAccess.__pulumiType;
    }

    /**
     * The ID of the dataset containing this table.
     */
    public readonly datasetId!: pulumi.Output<string>;
    /**
     * A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * An email address of a Google Group to grant access to.
     */
    public readonly groupByEmail!: pulumi.Output<string | undefined>;
    /**
     * Some other type of member that appears in the IAM Policy but isn't a user,
     * group, domain, or special group. For example: `allUsers`
     */
    public readonly iamMember!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Describes the rights granted to the user specified by the other
     * member of the access object. Primitive, Predefined and custom
     * roles are supported. Predefined roles that have equivalent
     * primitive roles are swapped by the API to their Primitive
     * counterparts, and will show a diff post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * A special group to grant access to. Possible values include:
     */
    public readonly specialGroup!: pulumi.Output<string | undefined>;
    /**
     * An email address of a user to grant access to. For example:
     * fred@example.com
     */
    public readonly userByEmail!: pulumi.Output<string | undefined>;
    /**
     * A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.  Structure is documented below.
     */
    public readonly view!: pulumi.Output<outputs.bigquery.DatasetAccessView | undefined>;

    /**
     * Create a DatasetAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetAccessArgs | DatasetAccessState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatasetAccessState | undefined;
            inputs["datasetId"] = state ? state.datasetId : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["groupByEmail"] = state ? state.groupByEmail : undefined;
            inputs["iamMember"] = state ? state.iamMember : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["specialGroup"] = state ? state.specialGroup : undefined;
            inputs["userByEmail"] = state ? state.userByEmail : undefined;
            inputs["view"] = state ? state.view : undefined;
        } else {
            const args = argsOrState as DatasetAccessArgs | undefined;
            if (!args || args.datasetId === undefined) {
                throw new Error("Missing required property 'datasetId'");
            }
            inputs["datasetId"] = args ? args.datasetId : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["groupByEmail"] = args ? args.groupByEmail : undefined;
            inputs["iamMember"] = args ? args.iamMember : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["specialGroup"] = args ? args.specialGroup : undefined;
            inputs["userByEmail"] = args ? args.userByEmail : undefined;
            inputs["view"] = args ? args.view : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatasetAccess.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatasetAccess resources.
 */
export interface DatasetAccessState {
    /**
     * The ID of the dataset containing this table.
     */
    readonly datasetId?: pulumi.Input<string>;
    /**
     * A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * An email address of a Google Group to grant access to.
     */
    readonly groupByEmail?: pulumi.Input<string>;
    /**
     * Some other type of member that appears in the IAM Policy but isn't a user,
     * group, domain, or special group. For example: `allUsers`
     */
    readonly iamMember?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Describes the rights granted to the user specified by the other
     * member of the access object. Primitive, Predefined and custom
     * roles are supported. Predefined roles that have equivalent
     * primitive roles are swapped by the API to their Primitive
     * counterparts, and will show a diff post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     */
    readonly role?: pulumi.Input<string>;
    /**
     * A special group to grant access to. Possible values include:
     */
    readonly specialGroup?: pulumi.Input<string>;
    /**
     * An email address of a user to grant access to. For example:
     * fred@example.com
     */
    readonly userByEmail?: pulumi.Input<string>;
    /**
     * A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.  Structure is documented below.
     */
    readonly view?: pulumi.Input<inputs.bigquery.DatasetAccessView>;
}

/**
 * The set of arguments for constructing a DatasetAccess resource.
 */
export interface DatasetAccessArgs {
    /**
     * The ID of the dataset containing this table.
     */
    readonly datasetId: pulumi.Input<string>;
    /**
     * A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * An email address of a Google Group to grant access to.
     */
    readonly groupByEmail?: pulumi.Input<string>;
    /**
     * Some other type of member that appears in the IAM Policy but isn't a user,
     * group, domain, or special group. For example: `allUsers`
     */
    readonly iamMember?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Describes the rights granted to the user specified by the other
     * member of the access object. Primitive, Predefined and custom
     * roles are supported. Predefined roles that have equivalent
     * primitive roles are swapped by the API to their Primitive
     * counterparts, and will show a diff post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     */
    readonly role?: pulumi.Input<string>;
    /**
     * A special group to grant access to. Possible values include:
     */
    readonly specialGroup?: pulumi.Input<string>;
    /**
     * An email address of a user to grant access to. For example:
     * fred@example.com
     */
    readonly userByEmail?: pulumi.Input<string>;
    /**
     * A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.  Structure is documented below.
     */
    readonly view?: pulumi.Input<inputs.bigquery.DatasetAccessView>;
}
