// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class IAMPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMPolicyState, opts?: pulumi.CustomResourceOptions): IAMPolicy {
        return new IAMPolicy(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly etag: pulumi.Output<string>;
    public readonly orgId: pulumi.Output<string>;
    public readonly policyData: pulumi.Output<string>;

    /**
     * Create a IAMPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMPolicyArgs | IAMPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: IAMPolicyState = argsOrState as IAMPolicyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["orgId"] = state ? state.orgId : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as IAMPolicyArgs | undefined;
            if (!args || args.orgId === undefined) {
                throw new Error("Missing required property 'orgId'");
            }
            if (!args || args.policyData === undefined) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["orgId"] = args ? args.orgId : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super("gcp:organizations/iAMPolicy:IAMPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMPolicy resources.
 */
export interface IAMPolicyState {
    readonly etag?: pulumi.Input<string>;
    readonly orgId?: pulumi.Input<string>;
    readonly policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMPolicy resource.
 */
export interface IAMPolicyArgs {
    readonly orgId: pulumi.Input<string>;
    readonly policyData: pulumi.Input<string>;
}
