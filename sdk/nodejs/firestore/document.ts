// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * In Cloud Firestore, the unit of storage is the document. A document is a lightweight record
 * that contains fields, which map to values. Each document is identified by a name.
 *
 * To get more information about Document, see:
 *
 * * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/firestore/docs/manage-data/add-data)
 *
 * > **Warning:** This resource creates a Firestore Document on a project that already has
 * Firestore enabled. If you haven't already enabled it, you can create a
 * `gcp.appengine.Application` resource with `databaseType` set to
 * `"CLOUD_FIRESTORE"` to do so. Your Firestore location will be the same as
 * the App Engine location specified.
 *
 * ## Example Usage
 * ### Firestore Document Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const mydoc = new gcp.firestore.Document("mydoc", {
 *     collection: "somenewcollection",
 *     documentId: "my-doc-%{random_suffix}",
 *     fields: "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}",
 *     project: "my-project-name",
 * });
 * ```
 * ### Firestore Document Nested Document
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const mydoc = new gcp.firestore.Document("mydoc", {
 *     collection: "somenewcollection",
 *     documentId: "my-doc-%{random_suffix}",
 *     fields: "{\"something\":{\"mapValue\":{\"fields\":{\"akey\":{\"stringValue\":\"avalue\"}}}}}",
 *     project: "my-project-name",
 * });
 * const subDocument = new gcp.firestore.Document("sub_document", {
 *     collection: pulumi.interpolate`${mydoc.path}/subdocs`,
 *     documentId: "bitcoinkey",
 *     fields: "{\"something\":{\"mapValue\":{\"fields\":{\"ayo\":{\"stringValue\":\"val2\"}}}}}",
 *     project: "my-project-name",
 * });
 * const subSubDocument = new gcp.firestore.Document("sub_sub_document", {
 *     collection: pulumi.interpolate`${subDocument.path}/subsubdocs`,
 *     documentId: "asecret",
 *     fields: "{\"something\":{\"mapValue\":{\"fields\":{\"secret\":{\"stringValue\":\"hithere\"}}}}}",
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * Document can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:firestore/document:Document default {{name}}
 * ```
 */
export class Document extends pulumi.CustomResource {
    /**
     * Get an existing Document resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentState, opts?: pulumi.CustomResourceOptions): Document {
        return new Document(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:firestore/document:Document';

    /**
     * Returns true if the given object is an instance of Document.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Document {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Document.__pulumiType;
    }

    /**
     * The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
     */
    public readonly collection!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The Firestore database id. Defaults to `"(default)"`.
     */
    public readonly database!: pulumi.Output<string | undefined>;
    /**
     * The client-assigned document ID to use for this document during creation.
     */
    public readonly documentId!: pulumi.Output<string>;
    /**
     * The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     */
    public readonly fields!: pulumi.Output<string>;
    /**
     * A server defined name for this index. Format:
     * 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A relative path to the collection this document exists within
     */
    public /*out*/ readonly path!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Last update timestamp in RFC3339 format.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Document resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentArgs | DocumentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentState | undefined;
            inputs["collection"] = state ? state.collection : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["documentId"] = state ? state.documentId : undefined;
            inputs["fields"] = state ? state.fields : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as DocumentArgs | undefined;
            if ((!args || args.collection === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collection'");
            }
            if ((!args || args.documentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'documentId'");
            }
            if ((!args || args.fields === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fields'");
            }
            inputs["collection"] = args ? args.collection : undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["documentId"] = args ? args.documentId : undefined;
            inputs["fields"] = args ? args.fields : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["path"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Document.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Document resources.
 */
export interface DocumentState {
    /**
     * The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
     */
    collection?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The Firestore database id. Defaults to `"(default)"`.
     */
    database?: pulumi.Input<string>;
    /**
     * The client-assigned document ID to use for this document during creation.
     */
    documentId?: pulumi.Input<string>;
    /**
     * The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     */
    fields?: pulumi.Input<string>;
    /**
     * A server defined name for this index. Format:
     * 'projects/{{project_id}}/databases/{{database_id}}/documents/{{path}}/{{document_id}}'
     */
    name?: pulumi.Input<string>;
    /**
     * A relative path to the collection this document exists within
     */
    path?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Last update timestamp in RFC3339 format.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Document resource.
 */
export interface DocumentArgs {
    /**
     * The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
     */
    collection: pulumi.Input<string>;
    /**
     * The Firestore database id. Defaults to `"(default)"`.
     */
    database?: pulumi.Input<string>;
    /**
     * The client-assigned document ID to use for this document during creation.
     */
    documentId: pulumi.Input<string>;
    /**
     * The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
     */
    fields: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
