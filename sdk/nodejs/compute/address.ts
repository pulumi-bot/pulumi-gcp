// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Address extends pulumi.CustomResource {
    /**
     * Get an existing Address resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressState, opts?: pulumi.CustomResourceOptions): Address {
        return new Address(name, <any>state, { ...opts, id: id });
    }

    public readonly address: pulumi.Output<string>;
    public readonly addressType: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networkTier: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly subnetwork: pulumi.Output<string>;
    public /*out*/ readonly users: pulumi.Output<string[]>;

    /**
     * Create a Address resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressArgs | AddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AddressState = argsOrState as AddressState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["addressType"] = state ? state.addressType : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkTier"] = state ? state.networkTier : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as AddressArgs | undefined;
            inputs["address"] = args ? args.address : undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkTier"] = args ? args.networkTier : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["users"] = undefined /*out*/;
        }
        super("gcp:compute/address:Address", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Address resources.
 */
export interface AddressState {
    readonly address?: pulumi.Input<string>;
    readonly addressType?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly selfLink?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Address resource.
 */
export interface AddressArgs {
    readonly address?: pulumi.Input<string>;
    readonly addressType?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
}
