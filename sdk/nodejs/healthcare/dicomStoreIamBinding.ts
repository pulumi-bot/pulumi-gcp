// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:
 *
 * * `gcp.healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
 * * `gcp.healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
 * * `gcp.healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.
 *
 * > **Note:** `gcp.healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.DicomStoreIamBinding` and `gcp.healthcare.DicomStoreIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_healthcare\_dicom\_store\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const dicomStore = new gcp.healthcare.DicomStoreIamPolicy("dicomStore", {
 *     dicomStoreId: "your-dicom-store-id",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_healthcare\_dicom\_store\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dicomStore = new gcp.healthcare.DicomStoreIamBinding("dicom_store", {
 *     dicomStoreId: "your-dicom-store-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_healthcare\_dicom\_store\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dicomStore = new gcp.healthcare.DicomStoreIamMember("dicom_store", {
 *     dicomStoreId: "your-dicom-store-id",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 */
export class DicomStoreIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing DicomStoreIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DicomStoreIamBindingState, opts?: pulumi.CustomResourceOptions): DicomStoreIamBinding {
        return new DicomStoreIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/dicomStoreIamBinding:DicomStoreIamBinding';

    /**
     * Returns true if the given object is an instance of DicomStoreIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DicomStoreIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DicomStoreIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.healthcare.DicomStoreIamBindingCondition | undefined>;
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly dicomStoreId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the DICOM store's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a DicomStoreIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DicomStoreIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DicomStoreIamBindingArgs | DicomStoreIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DicomStoreIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["dicomStoreId"] = state ? state.dicomStoreId : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as DicomStoreIamBindingArgs | undefined;
            if (!args || args.dicomStoreId === undefined) {
                throw new Error("Missing required property 'dicomStoreId'");
            }
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["dicomStoreId"] = args ? args.dicomStoreId : undefined;
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
        super(DicomStoreIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DicomStoreIamBinding resources.
 */
export interface DicomStoreIamBindingState {
    readonly condition?: pulumi.Input<inputs.healthcare.DicomStoreIamBindingCondition>;
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly dicomStoreId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the DICOM store's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DicomStoreIamBinding resource.
 */
export interface DicomStoreIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.healthcare.DicomStoreIamBindingCondition>;
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly dicomStoreId: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
