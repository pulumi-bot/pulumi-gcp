// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * AccessPolicy is a container for AccessLevels (which define the necessary
 * attributes to use GCP services) and ServicePerimeters (which define
 * regions of services able to freely pass data within a perimeter). An
 * access policy is globally visible within an organization, and the
 * restrictions it specifies apply to all projects within an organization.
 *
 *
 * To get more information about AccessPolicy, see:
 *
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies)
 * * How-to Guides
 *     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
 *
 * ## Example Usage
 *
 * ### Access Context Manager Access Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const access_policy = new gcp.accesscontextmanager.AccessPolicy("access-policy", {
 *     parent: "organizations/123456789",
 *     title: "my policy",
 * });
 * ```
 */
export class AccessPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessPolicyState, opts?: pulumi.CustomResourceOptions): AccessPolicy {
        return new AccessPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/accessPolicy:AccessPolicy';

    /**
     * Returns true if the given object is an instance of AccessPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessPolicy.__pulumiType;
    }

    /**
     * Time the AccessPolicy was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Resource name of the AccessPolicy. Format: {policy_id}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of this AccessPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id}
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Human readable title. Does not affect behavior.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Time the AccessPolicy was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a AccessPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessPolicyArgs | AccessPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AccessPolicyState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AccessPolicyArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["parent"] = args ? args.parent : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AccessPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessPolicy resources.
 */
export interface AccessPolicyState {
    /**
     * Time the AccessPolicy was created in UTC.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Resource name of the AccessPolicy. Format: {policy_id}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of this AccessPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id}
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Human readable title. Does not affect behavior.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * Time the AccessPolicy was updated in UTC.
     */
    readonly updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessPolicy resource.
 */
export interface AccessPolicyArgs {
    /**
     * The parent of this AccessPolicy in the Cloud Resource Hierarchy.
     * Format: organizations/{organization_id}
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Human readable title. Does not affect behavior.
     */
    readonly title: pulumi.Input<string>;
}
