// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
 * consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
 * for service accounts.
 *
 *
 * To get more information about HmacKey, see:
 *
 * * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)
 *
 * > **Warning:** All arguments including the `secret` value will be stored in the raw
 * state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
 * On import, the `secret` value will not be retrieved.
 *
 * > **Warning:** All arguments including `secret` will be stored in the raw
 * state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
 *
 * ## Example Usage - Storage Hmac Key
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serviceAccount = new gcp.serviceAccount.Account("serviceAccount", {accountId: "my-svc-acc"});
 * const key = new gcp.storage.HmacKey("key", {serviceAccountEmail: serviceAccount.email});
 * ```
 */
export class HmacKey extends pulumi.CustomResource {
    /**
     * Get an existing HmacKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HmacKeyState, opts?: pulumi.CustomResourceOptions): HmacKey {
        return new HmacKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/hmacKey:HmacKey';

    /**
     * Returns true if the given object is an instance of HmacKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HmacKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HmacKey.__pulumiType;
    }

    /**
     * The access ID of the HMAC Key.
     */
    public /*out*/ readonly accessId!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * HMAC secret key material.
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * The email address of the key's associated service account.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * The state of the key. Can be set to one of ACTIVE, INACTIVE.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * 'The creation time of the HMAC key in RFC 3339 format. '
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * 'The last modification time of the HMAC key metadata in RFC 3339 format.'
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;

    /**
     * Create a HmacKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HmacKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HmacKeyArgs | HmacKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as HmacKeyState | undefined;
            inputs["accessId"] = state ? state.accessId : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["updated"] = state ? state.updated : undefined;
        } else {
            const args = argsOrState as HmacKeyArgs | undefined;
            if (!args || args.serviceAccountEmail === undefined) {
                throw new Error("Missing required property 'serviceAccountEmail'");
            }
            inputs["project"] = args ? args.project : undefined;
            inputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["accessId"] = undefined /*out*/;
            inputs["secret"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(HmacKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HmacKey resources.
 */
export interface HmacKeyState {
    /**
     * The access ID of the HMAC Key.
     */
    readonly accessId?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * HMAC secret key material.
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * The email address of the key's associated service account.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The state of the key. Can be set to one of ACTIVE, INACTIVE.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * 'The creation time of the HMAC key in RFC 3339 format. '
     */
    readonly timeCreated?: pulumi.Input<string>;
    /**
     * 'The last modification time of the HMAC key metadata in RFC 3339 format.'
     */
    readonly updated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HmacKey resource.
 */
export interface HmacKeyArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The email address of the key's associated service account.
     */
    readonly serviceAccountEmail: pulumi.Input<string>;
    /**
     * The state of the key. Can be set to one of ACTIVE, INACTIVE.
     */
    readonly state?: pulumi.Input<string>;
}
