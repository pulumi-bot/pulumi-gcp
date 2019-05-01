// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for storage bucket. Each of these resources serves a different use case:
 * 
 * * `google_storage_bucket_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the storage bucket are preserved.
 * * `google_storage_bucket_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the storage bucket are preserved.
 * * `google_storage_bucket_iam_policy`: Setting a policy removes all other permissions on the bucket, and if done incorrectly, there's a real chance you will lock yourself out of the bucket. If possible for your use case, using multiple google_storage_bucket_iam_binding resources will be much safer. See the usage example on how to work with policy correctly.
 * 
 * 
 * > **Note:** `google_storage_bucket_iam_binding` resources **can be** used in conjunction with `google_storage_bucket_iam_member` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_storage\_bucket\_iam\_binding
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const binding = new gcp.storage.BucketIAMBinding("binding", {
 *     bucket: "your-bucket-name",
 *     members: ["user:jane@example.com"],
 *     role: "roles/storage.objectViewer",
 * });
 * ```
 * 
 * ## google\_storage\_bucket\_iam\_member
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const member = new gcp.storage.BucketIAMMember("member", {
 *     bucket: "your-bucket-name",
 *     member: "user:jane@example.com",
 *     role: "roles/storage.objectViewer",
 * });
 * ```
 * 
 * ## google\_storage\_bucket\_iam\_policy
 * 
 * When applying a policy that does not include the roles listed below, you lose the default permissions which google adds to your bucket:
 * * `roles/storage.legacyBucketOwner`
 * * `roles/storage.legacyBucketReader`
 * 
 * If this happens only an entity with `roles/storage.admin` privileges can repair this bucket's policies. It is recommended to include the above roles in policies to get the same behaviour as with the other two options.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const foo_policy = pulumi.output(gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         members: ["group:yourgroup@example.com"],
 *         role: "roles/your-role",
 *     }],
 * }));
 * const member = new gcp.storage.BucketIAMPolicy("member", {
 *     bucket: "your-bucket-name",
 *     policyData: foo_policy.apply(foo_policy => foo_policy.policyData),
 * });
 * ```
 */
export class BucketIAMMember extends pulumi.CustomResource {
    /**
     * Get an existing BucketIAMMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketIAMMemberState, opts?: pulumi.CustomResourceOptions): BucketIAMMember {
        return new BucketIAMMember(name, <any>state, { ...opts, id: id });
    }

    /**
     * The name of the bucket it applies to.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the storage bucket's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    public readonly member!: pulumi.Output<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a BucketIAMMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketIAMMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketIAMMemberArgs | BucketIAMMemberState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketIAMMemberState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["member"] = state ? state.member : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as BucketIAMMemberArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.member === undefined) {
                throw new Error("Missing required property 'member'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["member"] = args ? args.member : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:storage/bucketIAMMember:BucketIAMMember", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketIAMMember resources.
 */
export interface BucketIAMMemberState {
    /**
     * The name of the bucket it applies to.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the storage bucket's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    readonly member?: pulumi.Input<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketIAMMember resource.
 */
export interface BucketIAMMemberArgs {
    /**
     * The name of the bucket it applies to.
     */
    readonly bucket: pulumi.Input<string>;
    readonly member: pulumi.Input<string>;
    /**
     * The role that should be applied. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
