// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access a Network Endpoint Group's attributes.
 * 
 * The NEG may be found by providing either a `selfLink`, or a `name` and a `zone`.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const neg1 = gcp.compute.getNetworkEndpointGroup({
 *     name: "k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
 *     zone: "us-central1-a",
 * });
 * const neg2 = gcp.compute.getNetworkEndpointGroup({
 *     selfLink: "https://www.googleapis.com/compute/v1/projects/myproject/zones/us-central1-a/networkEndpointGroups/k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_compute_network_endpoint_group.html.markdown.
 */
export function getNetworkEndpointGroup(args?: GetNetworkEndpointGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkEndpointGroupResult> & GetNetworkEndpointGroupResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetNetworkEndpointGroupResult> = pulumi.runtime.invoke("gcp:compute/getNetworkEndpointGroup:getNetworkEndpointGroup", {
        "name": args.name,
        "selfLink": args.selfLink,
        "zone": args.zone,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getNetworkEndpointGroup.
 */
export interface GetNetworkEndpointGroupArgs {
    /**
     * The Network Endpoint Group name.
     * Provide either this or a `selfLink`.
     */
    readonly name?: string;
    /**
     * The Network Endpoint Group self\_link.
     */
    readonly selfLink?: string;
    /**
     * The Network Endpoint Group availability zone.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getNetworkEndpointGroup.
 */
export interface GetNetworkEndpointGroupResult {
    /**
     * The NEG default port.
     */
    readonly defaultPort: number;
    /**
     * The NEG description.
     */
    readonly description: string;
    readonly name?: string;
    /**
     * The network to which all network endpoints in the NEG belong.
     */
    readonly network: string;
    /**
     * Type of network endpoints in this network endpoint group.
     */
    readonly networkEndpointType: string;
    readonly project: string;
    readonly selfLink?: string;
    /**
     * Number of network endpoints in the network endpoint group.
     */
    readonly size: number;
    /**
     * subnetwork to which all network endpoints in the NEG belong.
     */
    readonly subnetwork: string;
    readonly zone?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
