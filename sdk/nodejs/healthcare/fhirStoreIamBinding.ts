// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
 *
 * * `gcp.healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
 * * `gcp.healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
 * * `gcp.healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
 *
 * > **Note:** `gcp.healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.FhirStoreIamBinding` and `gcp.healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
 */
export class FhirStoreIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing FhirStoreIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FhirStoreIamBindingState, opts?: pulumi.CustomResourceOptions): FhirStoreIamBinding {
        return new FhirStoreIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding';

    /**
     * Returns true if the given object is an instance of FhirStoreIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FhirStoreIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FhirStoreIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.healthcare.FhirStoreIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the FHIR store's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The FHIR store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
     * `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly fhirStoreId!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a FhirStoreIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FhirStoreIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FhirStoreIamBindingArgs | FhirStoreIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FhirStoreIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["fhirStoreId"] = state ? state.fhirStoreId : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as FhirStoreIamBindingArgs | undefined;
            if (!args || args.fhirStoreId === undefined) {
                throw new Error("Missing required property 'fhirStoreId'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["fhirStoreId"] = args ? args.fhirStoreId : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FhirStoreIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FhirStoreIamBinding resources.
 */
export interface FhirStoreIamBindingState {
    readonly condition?: pulumi.Input<inputs.healthcare.FhirStoreIamBindingCondition>;
    /**
     * (Computed) The etag of the FHIR store's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The FHIR store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
     * `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly fhirStoreId?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FhirStoreIamBinding resource.
 */
export interface FhirStoreIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.healthcare.FhirStoreIamBindingCondition>;
    /**
     * The FHIR store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
     * `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly fhirStoreId: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
