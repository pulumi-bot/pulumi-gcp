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
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)
 * 
 * ## Example Usage - Global Address Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultGlobalAddress = new gcp.compute.GlobalAddress("default", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_global_address.html.markdown.
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

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/globalAddress:GlobalAddress';

    /**
     * Returns true if the given object is an instance of GlobalAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalAddress.__pulumiType;
    }

    public readonly address!: pulumi.Output<string>;
    public readonly addressType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly ipVersion!: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly network!: pulumi.Output<string | undefined>;
    public readonly prefixLength!: pulumi.Output<number | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly purpose!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;

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
            const state = argsOrState as GlobalAddressState | undefined;
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
            inputs["address"] = args ? args.address : undefined;
            inputs["addressType"] = args ? args.addressType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["prefixLength"] = args ? args.prefixLength : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["purpose"] = args ? args.purpose : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GlobalAddress.__pulumiType, name, inputs, opts);
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
    readonly address?: pulumi.Input<string>;
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
