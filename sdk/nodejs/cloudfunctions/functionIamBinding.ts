// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Functions CloudFunction. Each of these resources serves a different use case:
 *
 * * `gcp.cloudfunctions.FunctionIamPolicy`: Authoritative. Sets the IAM policy for the cloudfunction and replaces any existing policy already attached.
 * * `gcp.cloudfunctions.FunctionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cloudfunction are preserved.
 * * `gcp.cloudfunctions.FunctionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cloudfunction are preserved.
 *
 * > **Note:** `gcp.cloudfunctions.FunctionIamPolicy` **cannot** be used in conjunction with `gcp.cloudfunctions.FunctionIamBinding` and `gcp.cloudfunctions.FunctionIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.cloudfunctions.FunctionIamBinding` resources **can be** used in conjunction with `gcp.cloudfunctions.FunctionIamMember` resources **only if** they do not grant privilege to the same role.
 */
export class FunctionIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing FunctionIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionIamBindingState, opts?: pulumi.CustomResourceOptions): FunctionIamBinding {
        return new FunctionIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudfunctions/functionIamBinding:FunctionIamBinding';

    /**
     * Returns true if the given object is an instance of FunctionIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionIamBinding.__pulumiType;
    }

    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly cloudFunction!: pulumi.Output<string>;
    public readonly condition!: pulumi.Output<outputs.cloudfunctions.FunctionIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a FunctionIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionIamBindingArgs | FunctionIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FunctionIamBindingState | undefined;
            inputs["cloudFunction"] = state ? state.cloudFunction : undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as FunctionIamBindingArgs | undefined;
            if (!args || args.cloudFunction === undefined) {
                throw new Error("Missing required property 'cloudFunction'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["cloudFunction"] = args ? args.cloudFunction : undefined;
            inputs["condition"] = args ? args.condition : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FunctionIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionIamBinding resources.
 */
export interface FunctionIamBindingState {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly cloudFunction?: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.cloudfunctions.FunctionIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionIamBinding resource.
 */
export interface FunctionIamBindingArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly cloudFunction: pulumi.Input<string>;
    readonly condition?: pulumi.Input<inputs.cloudfunctions.FunctionIamBindingCondition>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudfunctions.FunctionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
