// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a Global Address resource. Global addresses are used for
 * HTTP(S) load balancing.
 * 
 * 
 * To get more information about GlobalAddress, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/latest/globalAddresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=global_address_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Global Address Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_global_address_default = new gcp.compute.GlobalAddress("default", {
 *     name: "global-appserver-ip",
 * });
 * ```
 */
export class GlobalAddress extends pulumi.CustomResource {
    /**
     * Get an existing GlobalAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalAddressState, opts?: pulumi.CustomResourceOptions): GlobalAddress {
        return new GlobalAddress(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly address: pulumi.Output<string>;
    public readonly addressType: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly ipVersion: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly network: pulumi.Output<string | undefined>;
    public readonly prefixLength: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly purpose: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;

    /**
     * Create a GlobalAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GlobalAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalAddressArgs | GlobalAddressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: GlobalAddressState = argsOrState as GlobalAddressState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["addressType"] = state ? state.addressType : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["prefixLength"] = state ? state.prefixLength : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["purpose"] = state ? state.purpose : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
        } else {
            const args = argsOrState as GlobalAddressArgs | undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["prefixLength"] = args ? args.prefixLength : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["address"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/globalAddress:GlobalAddress", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalAddress resources.
 */
export interface GlobalAddressState {
    readonly address?: pulumi.Input<string>;
    readonly addressType?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly ipVersion?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly prefixLength?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly purpose?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GlobalAddress resource.
 */
export interface GlobalAddressArgs {
    readonly addressType?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly ipVersion?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly prefixLength?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly purpose?: pulumi.Input<string>;
}
