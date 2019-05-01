// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new bucket ACL in Google cloud storage service (GCS). For more information see 
 * [the official documentation](https://cloud.google.com/storage/docs/access-control/lists) 
 * and 
 * [API](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls).
 * 
 * ## Example Usage
 * 
 * Example creating an ACL on a bucket with one owner, and one reader.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const image_store = new gcp.storage.Bucket("image-store", {
 *     location: "EU",
 * });
 * const image_store_acl = new gcp.storage.BucketACL("image-store-acl", {
 *     bucket: image_store.name,
 *     roleEntities: [
 *         "OWNER:user-my.email@gmail.com",
 *         "READER:group-mygroup",
 *     ],
 * });
 * ```
 */
export class BucketACL extends pulumi.CustomResource {
    /**
     * Get an existing BucketACL resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketACLState, opts?: pulumi.CustomResourceOptions): BucketACL {
        return new BucketACL(name, <any>state, { ...opts, id: id });
    }

    /**
     * The name of the bucket it applies to.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    public readonly defaultAcl!: pulumi.Output<string | undefined>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
     */
    public readonly predefinedAcl!: pulumi.Output<string | undefined>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
     */
    public readonly roleEntities!: pulumi.Output<string[]>;

    /**
     * Create a BucketACL resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketACLArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketACLArgs | BucketACLState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketACLState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["defaultAcl"] = state ? state.defaultAcl : undefined;
            inputs["predefinedAcl"] = state ? state.predefinedAcl : undefined;
            inputs["roleEntities"] = state ? state.roleEntities : undefined;
        } else {
            const args = argsOrState as BucketACLArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["defaultAcl"] = args ? args.defaultAcl : undefined;
            inputs["predefinedAcl"] = args ? args.predefinedAcl : undefined;
            inputs["roleEntities"] = args ? args.roleEntities : undefined;
        }
        super("gcp:storage/bucketACL:BucketACL", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketACL resources.
 */
export interface BucketACLState {
    /**
     * The name of the bucket it applies to.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    readonly defaultAcl?: pulumi.Input<string>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
     */
    readonly predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
     */
    readonly roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a BucketACL resource.
 */
export interface BucketACLArgs {
    /**
     * The name of the bucket it applies to.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * Configure this ACL to be the default ACL.
     */
    readonly defaultAcl?: pulumi.Input<string>;
    /**
     * The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if `role_entity` is not.
     */
    readonly predefinedAcl?: pulumi.Input<string>;
    /**
     * List of role/entity pairs in the form `ROLE:entity`. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls)  for more details. Must be set if `predefined_acl` is not.
     */
    readonly roleEntities?: pulumi.Input<pulumi.Input<string>[]>;
}
