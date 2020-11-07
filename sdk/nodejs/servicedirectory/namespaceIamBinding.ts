// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Service Directory Namespace. Each of these resources serves a different use case:
 *
 * * `gcp.servicedirectory.NamespaceIamPolicy`: Authoritative. Sets the IAM policy for the namespace and replaces any existing policy already attached.
 * * `gcp.servicedirectory.NamespaceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the namespace are preserved.
 * * `gcp.servicedirectory.NamespaceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the namespace are preserved.
 *
 * > **Note:** `gcp.servicedirectory.NamespaceIamPolicy` **cannot** be used in conjunction with `gcp.servicedirectory.NamespaceIamBinding` and `gcp.servicedirectory.NamespaceIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.servicedirectory.NamespaceIamBinding` resources **can be** used in conjunction with `gcp.servicedirectory.NamespaceIamMember` resources **only if** they do not grant privilege to the same role.
 */
export class NamespaceIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing NamespaceIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceIamBindingState, opts?: pulumi.CustomResourceOptions): NamespaceIamBinding {
        return new NamespaceIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding';

    /**
     * Returns true if the given object is an instance of NamespaceIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamespaceIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamespaceIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.servicedirectory.NamespaceIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a NamespaceIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceIamBindingArgs | NamespaceIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NamespaceIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as NamespaceIamBindingArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NamespaceIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NamespaceIamBinding resources.
 */
export interface NamespaceIamBindingState {
    readonly condition?: pulumi.Input<inputs.servicedirectory.NamespaceIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NamespaceIamBinding resource.
 */
export interface NamespaceIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.servicedirectory.NamespaceIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
