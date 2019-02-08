// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The DefaultObjectAccessControls resources represent the Access Control
 * Lists (ACLs) applied to a new object within a Google Cloud Storage bucket
 * when no ACL was provided for that object. ACLs let you specify who has
 * access to your bucket contents and to what extent.
 * 
 * There are two roles that can be assigned to an entity:
 * 
 * READERs can get an object, though the acl property will not be revealed.
 * OWNERs are READERs, and they can get the acl property, update an object,
 * and call all objectAccessControls methods on the object. The owner of an
 * object is always an OWNER.
 * For more information, see Access Control, with the caveat that this API
 * uses READER and OWNER instead of READ and FULL_CONTROL.
 * 
 * 
 * To get more information about DefaultObjectAccessControl, see:
 * 
 * * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/defaultObjectAccessControls)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/storage/docs/access-control/create-manage-lists)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=storage_default_object_access_control_public&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Storage Default Object Access Control Public
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const bucket = new gcp.storage.Bucket("bucket", {});
 * const publicRule = new gcp.storage.DefaultObjectAccessControl("public_rule", {
 *     bucket: bucket.name,
 *     entity: "allUsers",
 *     role: "READER",
 * });
 * ```
 */
export class DefaultObjectAccessControl extends pulumi.CustomResource {
    /**
     * Get an existing DefaultObjectAccessControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultObjectAccessControlState, opts?: pulumi.CustomResourceOptions): DefaultObjectAccessControl {
        return new DefaultObjectAccessControl(name, <any>state, { ...opts, id: id });
    }

    public readonly bucket: pulumi.Output<string>;
    public /*out*/ readonly domain: pulumi.Output<string>;
    public /*out*/ readonly email: pulumi.Output<string>;
    public readonly entity: pulumi.Output<string>;
    public /*out*/ readonly entityId: pulumi.Output<string>;
    public /*out*/ readonly generation: pulumi.Output<number>;
    public readonly object: pulumi.Output<string | undefined>;
    public /*out*/ readonly projectTeam: pulumi.Output<{ projectNumber?: string, team?: string }>;
    public readonly role: pulumi.Output<string>;

    /**
     * Create a DefaultObjectAccessControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultObjectAccessControlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultObjectAccessControlArgs | DefaultObjectAccessControlState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DefaultObjectAccessControlState = argsOrState as DefaultObjectAccessControlState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["entity"] = state ? state.entity : undefined;
            inputs["entityId"] = state ? state.entityId : undefined;
            inputs["generation"] = state ? state.generation : undefined;
            inputs["object"] = state ? state.object : undefined;
            inputs["projectTeam"] = state ? state.projectTeam : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DefaultObjectAccessControlArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.entity === undefined) {
                throw new Error("Missing required property 'entity'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["entity"] = args ? args.entity : undefined;
            inputs["object"] = args ? args.object : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["domain"] = undefined /*out*/;
            inputs["email"] = undefined /*out*/;
            inputs["entityId"] = undefined /*out*/;
            inputs["generation"] = undefined /*out*/;
            inputs["projectTeam"] = undefined /*out*/;
        }
        super("gcp:storage/defaultObjectAccessControl:DefaultObjectAccessControl", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultObjectAccessControl resources.
 */
export interface DefaultObjectAccessControlState {
    readonly bucket?: pulumi.Input<string>;
    readonly domain?: pulumi.Input<string>;
    readonly email?: pulumi.Input<string>;
    readonly entity?: pulumi.Input<string>;
    readonly entityId?: pulumi.Input<string>;
    readonly generation?: pulumi.Input<number>;
    readonly object?: pulumi.Input<string>;
    readonly projectTeam?: pulumi.Input<{ projectNumber?: pulumi.Input<string>, team?: pulumi.Input<string> }>;
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultObjectAccessControl resource.
 */
export interface DefaultObjectAccessControlArgs {
    readonly bucket: pulumi.Input<string>;
    readonly entity: pulumi.Input<string>;
    readonly object?: pulumi.Input<string>;
    readonly role: pulumi.Input<string>;
}
