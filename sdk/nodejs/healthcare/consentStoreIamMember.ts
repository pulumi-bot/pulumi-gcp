// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* {{dataset}}/consentStores/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Healthcare consentstore IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor "{{dataset}}/consentStores/{{consent_store}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember editor {{dataset}}/consentStores/{{consent_store}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ConsentStoreIamMember extends pulumi.CustomResource {
    /**
     * Get an existing ConsentStoreIamMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConsentStoreIamMemberState, opts?: pulumi.CustomResourceOptions): ConsentStoreIamMember {
        return new ConsentStoreIamMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/consentStoreIamMember:ConsentStoreIamMember';

    /**
     * Returns true if the given object is an instance of ConsentStoreIamMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConsentStoreIamMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConsentStoreIamMember.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.healthcare.ConsentStoreIamMemberCondition | undefined>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly consentStoreId!: pulumi.Output<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a ConsentStoreIamMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConsentStoreIamMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConsentStoreIamMemberArgs | ConsentStoreIamMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConsentStoreIamMemberState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["consentStoreId"] = state ? state.consentStoreId : undefined;
            inputs["dataset"] = state ? state.dataset : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as ConsentStoreIamMemberArgs | undefined;
            if ((!args || args.consentStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'consentStoreId'");
            }
            if ((!args || args.dataset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataset'");
            }
            if ((!args || args.member === undefined) && !opts.urn) {
                throw new Error("Missing required property 'member'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["consentStoreId"] = args ? args.consentStoreId : undefined;
            inputs["dataset"] = args ? args.dataset : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ConsentStoreIamMember.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConsentStoreIamMember resources.
 */
export interface ConsentStoreIamMemberState {
    readonly condition?: pulumi.Input<inputs.healthcare.ConsentStoreIamMemberCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly consentStoreId?: pulumi.Input<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly dataset?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConsentStoreIamMember resource.
 */
export interface ConsentStoreIamMemberArgs {
    readonly condition?: pulumi.Input<inputs.healthcare.ConsentStoreIamMemberCondition>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly consentStoreId: pulumi.Input<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly dataset: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.ConsentStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
