// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents an Address resource.
 * 
 * Each virtual machine instance has an ephemeral internal IP address and,
 * optionally, an external IP address. To communicate between instances on
 * the same network, you can use an instance's internal IP address. To
 * communicate with the Internet and instances outside of the same network,
 * you must specify the instance's external IP address.
 * 
 * Internal IP addresses are ephemeral and only belong to an instance for
 * the lifetime of the instance; if the instance is deleted and recreated,
 * the instance is assigned a new internal IP address, either by Compute
 * Engine or by you. External IP addresses can be either ephemeral or
 * static.
 * 
 * 
 * To get more information about Address, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
 *     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
 * 
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=address_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Address Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_address_ip_address = new gcp.compute.Address("ip_address", {
 *     name: "my-address",
 * });
 * ```
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=address_with_subnetwork&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Address With Subnetwork
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_network_default = new gcp.compute.Network("default", {
 *     name: "my-network",
 * });
 * const google_compute_subnetwork_default = new gcp.compute.Subnetwork("default", {
 *     ipCidrRange: "10.0.0.0/16",
 *     name: "my-subnet",
 *     network: google_compute_network_default.selfLink,
 *     region: "us-central1",
 * });
 * const google_compute_address_internal_with_subnet_and_address = new gcp.compute.Address("internal_with_subnet_and_address", {
 *     address: "10.0.42.42",
 *     addressType: "INTERNAL",
 *     name: "my-internal-address",
 *     region: "us-central1",
 *     subnetwork: google_compute_subnetwork_default.selfLink,
 * });
 * ```
 * <div class = "oics-button" style="float: right; margin: 0 0 -15px">
 *   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=instance_with_ip&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
 *     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
 *   </a>
 * </div>
 * ## Example Usage - Instance With Ip
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_compute_image_debian_image = pulumi.output(gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * }));
 * const google_compute_address_static = new gcp.compute.Address("static", {
 *     name: "ipv4-address",
 * });
 * const google_compute_instance_instance_with_ip = new gcp.compute.Instance("instance_with_ip", {
 *     bootDisk: {
 *         initializeParams: {
 *             image: google_compute_image_debian_image.apply(__arg0 => __arg0.selfLink),
 *         },
 *     },
 *     machineType: "f1-micro",
 *     name: "vm-instance",
 *     networkInterfaces: [{
 *         accessConfigs: [{
 *             natIp: google_compute_address_static.address,
 *         }],
 *         network: "default",
 *     }],
 *     zone: "us-central1-a",
 * });
 * ```
 */
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

    /**
     * The IP of the created resource.
     */
    public readonly address: pulumi.Output<string>;
    public readonly addressType: pulumi.Output<string | undefined>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networkTier: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
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
    /**
     * The IP of the created resource.
     */
    readonly address?: pulumi.Input<string>;
    readonly addressType?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Address resource.
 */
export interface AddressArgs {
    /**
     * The IP of the created resource.
     */
    readonly address?: pulumi.Input<string>;
    readonly addressType?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    readonly networkTier?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly subnetwork?: pulumi.Input<string>;
}
