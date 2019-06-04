// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of a Google Cloud Platform folder. For more information see 
 * [the official documentation](https://cloud.google.com/resource-manager/docs/creating-managing-folders)
 * and 
 * [API](https://cloud.google.com/resource-manager/reference/rest/v2/folders).
 * 
 * A folder can contain projects, other folders, or a combination of both. You can use folders to group projects under an organization in a hierarchy. For example, your organization might contain multiple departments, each with its own set of Cloud Platform resources. Folders allows you to group these resources on a per-department basis. Folders are used to group resources that share common IAM policies.
 * 
 * Folders created live inside an Organization. See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
 * 
 * The service account used to run Terraform when creating a `google_folder`
 * resource must have `roles/resourcemanager.folderCreator`. See the
 * [Access Control for Folders Using IAM](https://cloud.google.com/resource-manager/docs/access-control-folders)
 * doc for more information.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * // Top-level folder under an organization.
 * const department1 = new gcp.organizations.Folder("department1", {
 *     displayName: "Department 1",
 *     parent: "organizations/1234567",
 * });
 * // Folder nested under another folder.
 * const team_abc = new gcp.organizations.Folder("team-abc", {
 *     displayName: "Team ABC",
 *     parent: department1.name,
 * });
 * ```
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState, opts?: pulumi.CustomResourceOptions): Folder {
        return new Folder(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:organizations/folder:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Folder.__pulumiType;
    }

    /**
     * Timestamp when the Folder was created. Assigned by the server.
     * A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The folder’s display name.
     * A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
     */
    public /*out*/ readonly lifecycleState!: pulumi.Output<string>;
    /**
     * The resource name of the Folder. Its format is folders/{folder_id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource name of the parent Folder or Organization.
     * Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
     */
    public readonly parent!: pulumi.Output<string>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FolderArgs | FolderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FolderState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["lifecycleState"] = state ? state.lifecycleState : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
        } else {
            const args = argsOrState as FolderArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["lifecycleState"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        super(Folder.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Folder resources.
 */
export interface FolderState {
    /**
     * Timestamp when the Folder was created. Assigned by the server.
     * A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * The folder’s display name.
     * A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The lifecycle state of the folder such as `ACTIVE` or `DELETE_REQUESTED`.
     */
    readonly lifecycleState?: pulumi.Input<string>;
    /**
     * The resource name of the Folder. Its format is folders/{folder_id}.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource name of the parent Folder or Organization.
     * Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
     */
    readonly parent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * The folder’s display name.
     * A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * The resource name of the parent Folder or Organization.
     * Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
     */
    readonly parent: pulumi.Input<string>;
}
