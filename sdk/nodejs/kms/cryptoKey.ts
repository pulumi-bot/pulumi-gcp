// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_crypto_key.html.markdown.
 */
export class CryptoKey extends pulumi.CustomResource {
    /**
     * Get an existing CryptoKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CryptoKeyState, opts?: pulumi.CustomResourceOptions): CryptoKey {
        return new CryptoKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/cryptoKey:CryptoKey';

    /**
     * Returns true if the given object is an instance of CryptoKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CryptoKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CryptoKey.__pulumiType;
    }

    public readonly keyRing!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly purpose!: pulumi.Output<string | undefined>;
    public readonly rotationPeriod!: pulumi.Output<string | undefined>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly versionTemplate!: pulumi.Output<{ algorithm: string, protectionLevel?: string }>;

    /**
     * Create a CryptoKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CryptoKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CryptoKeyArgs | CryptoKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CryptoKeyState | undefined;
            inputs["keyRing"] = state ? state.keyRing : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["purpose"] = state ? state.purpose : undefined;
            inputs["rotationPeriod"] = state ? state.rotationPeriod : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["versionTemplate"] = state ? state.versionTemplate : undefined;
        } else {
            const args = argsOrState as CryptoKeyArgs | undefined;
            if (!args || args.keyRing === undefined) {
                throw new Error("Missing required property 'keyRing'");
            }
            inputs["keyRing"] = args ? args.keyRing : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            inputs["versionTemplate"] = args ? args.versionTemplate : undefined;
            inputs["selfLink"] = undefined /*out*/;
        }
        super(CryptoKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CryptoKey resources.
 */
export interface CryptoKeyState {
    readonly keyRing?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly purpose?: pulumi.Input<string>;
    readonly rotationPeriod?: pulumi.Input<string>;
    readonly selfLink?: pulumi.Input<string>;
    readonly versionTemplate?: pulumi.Input<{ algorithm: pulumi.Input<string>, protectionLevel?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a CryptoKey resource.
 */
export interface CryptoKeyArgs {
    readonly keyRing: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly purpose?: pulumi.Input<string>;
    readonly rotationPeriod?: pulumi.Input<string>;
    readonly versionTemplate?: pulumi.Input<{ algorithm: pulumi.Input<string>, protectionLevel?: pulumi.Input<string> }>;
}
