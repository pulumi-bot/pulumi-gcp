// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Lien represents an encumbrance on the actions that can be performed on a resource.
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Resource Manager Lien
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.organizations.Project("project", {
 *     projectId: "staging-project",
 * });
 * const lien = new gcp.resourcemanager.Lien("lien", {
 *     origin: "machine-readable-explanation",
 *     parent: pulumi.interpolate`projects/${project.number}`,
 *     reason: "This project is an important environment",
 *     restrictions: ["resourcemanager.projects.delete"],
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class Lien extends pulumi.CustomResource {
    /**
     * Get an existing Lien resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LienState, opts?: pulumi.CustomResourceOptions): Lien {
        return new Lien(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:resourcemanager/lien:Lien';

    /**
     * Returns true if the given object is an instance of Lien.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lien {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lien.__pulumiType;
    }

    /**
     * Time of creation
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A system-generated unique identifier for this Lien.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A stable, user-visible/meaningful string identifying the origin
     * of the Lien, intended to be inspected programmatically. Maximum length of
     * 200 characters.
     */
    public readonly origin!: pulumi.Output<string>;
    /**
     * A reference to the resource this Lien is attached to.
     * The server will validate the parent against those for which Liens are supported.
     * Since a variety of objects can have Liens against them, you must provide the type
     * prefix (e.g. "projects/my-project-name").
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Concise user-visible strings indicating why an action cannot be performed
     * on a resource. Maximum length of 200 characters.
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * The types of operations which should be blocked as a result of this Lien.
     * Each value should correspond to an IAM permission. The server will validate
     * the permissions against those for which Liens are supported.  An empty
     * list is meaningless and will be rejected.
     * e.g. ['resourcemanager.projects.delete']
     */
    public readonly restrictions!: pulumi.Output<string[]>;

    /**
     * Create a Lien resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LienArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LienArgs | LienState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LienState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["origin"] = state ? state.origin : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["reason"] = state ? state.reason : undefined;
            inputs["restrictions"] = state ? state.restrictions : undefined;
        } else {
            const args = argsOrState as LienArgs | undefined;
            if (!args || args.origin === undefined) {
                throw new Error("Missing required property 'origin'");
            }
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            if (!args || args.reason === undefined) {
                throw new Error("Missing required property 'reason'");
            }
            if (!args || args.restrictions === undefined) {
                throw new Error("Missing required property 'restrictions'");
            }
            inputs["origin"] = args ? args.origin : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["reason"] = args ? args.reason : undefined;
            inputs["restrictions"] = args ? args.restrictions : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Lien.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lien resources.
 */
export interface LienState {
    /**
     * Time of creation
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * A system-generated unique identifier for this Lien.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A stable, user-visible/meaningful string identifying the origin
     * of the Lien, intended to be inspected programmatically. Maximum length of
     * 200 characters.
     */
    readonly origin?: pulumi.Input<string>;
    /**
     * A reference to the resource this Lien is attached to.
     * The server will validate the parent against those for which Liens are supported.
     * Since a variety of objects can have Liens against them, you must provide the type
     * prefix (e.g. "projects/my-project-name").
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Concise user-visible strings indicating why an action cannot be performed
     * on a resource. Maximum length of 200 characters.
     */
    readonly reason?: pulumi.Input<string>;
    /**
     * The types of operations which should be blocked as a result of this Lien.
     * Each value should correspond to an IAM permission. The server will validate
     * the permissions against those for which Liens are supported.  An empty
     * list is meaningless and will be rejected.
     * e.g. ['resourcemanager.projects.delete']
     */
    readonly restrictions?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Lien resource.
 */
export interface LienArgs {
    /**
     * A stable, user-visible/meaningful string identifying the origin
     * of the Lien, intended to be inspected programmatically. Maximum length of
     * 200 characters.
     */
    readonly origin: pulumi.Input<string>;
    /**
     * A reference to the resource this Lien is attached to.
     * The server will validate the parent against those for which Liens are supported.
     * Since a variety of objects can have Liens against them, you must provide the type
     * prefix (e.g. "projects/my-project-name").
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Concise user-visible strings indicating why an action cannot be performed
     * on a resource. Maximum length of 200 characters.
     */
    readonly reason: pulumi.Input<string>;
    /**
     * The types of operations which should be blocked as a result of this Lien.
     * Each value should correspond to an IAM permission. The server will validate
     * the permissions against those for which Liens are supported.  An empty
     * list is meaningless and will be rejected.
     * e.g. ['resourcemanager.projects.delete']
     */
    readonly restrictions: pulumi.Input<pulumi.Input<string>[]>;
}
