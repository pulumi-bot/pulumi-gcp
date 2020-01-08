// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_key_ring_iam_policy.html.markdown.
 */
export class KeyRingIAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyRingIAMPolicyState, opts?: pulumi.CustomResourceOptions): KeyRingIAMPolicy {
        return new KeyRingIAMPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy';

    /**
     * Returns true if the given object is an instance of KeyRingIAMPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyRingIAMPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyRingIAMPolicy.__pulumiType;
    }

    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly keyRingId!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a KeyRingIAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyRingIAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyRingIAMPolicyArgs | KeyRingIAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KeyRingIAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["keyRingId"] = state ? state.keyRingId : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as KeyRingIAMPolicyArgs | undefined;
            if (!args || args.keyRingId === undefined) {
                throw new Error("Missing required property 'keyRingId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["keyRingId"] = args ? args.keyRingId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(KeyRingIAMPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyRingIAMPolicy resources.
 */
export interface KeyRingIAMPolicyState {
    /**
     * (Computed) The etag of the key ring's IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly keyRingId?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KeyRingIAMPolicy resource.
 */
export interface KeyRingIAMPolicyArgs {
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    readonly keyRingId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: pulumi.Input<string>;
}
