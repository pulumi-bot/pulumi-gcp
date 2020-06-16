// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Network endpoint represents a IP address and port combination that is
 * part of a specific network endpoint group (NEG). NEGs are zonals
 * collection of these endpoints for GCP resources within a
 * single subnet. **NOTE**: Network endpoints cannot be created outside of a
 * network endpoint group.
 *
 * To get more information about NetworkEndpoint, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Network Endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myImage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const defaultNetwork = new gcp.compute.Network("defaultNetwork", {autoCreateSubnetworks: false});
 * const defaultSubnetwork = new gcp.compute.Subnetwork("defaultSubnetwork", {
 *     ipCidrRange: "10.0.0.1/16",
 *     region: "us-central1",
 *     network: defaultNetwork.id,
 * });
 * const endpoint_instance = new gcp.compute.Instance("endpoint-instance", {
 *     machineType: "n1-standard-1",
 *     boot_disk: {
 *         initialize_params: {
 *             image: myImage.then(myImage => myImage.selfLink),
 *         },
 *     },
 *     network_interface: [{
 *         subnetwork: defaultSubnetwork.id,
 *         access_config: [{}],
 *     }],
 * });
 * const default_endpoint = new gcp.compute.NetworkEndpoint("default-endpoint", {
 *     networkEndpointGroup: google_compute_network_endpoint_group.neg.name,
 *     instance: endpoint_instance.name,
 *     port: google_compute_network_endpoint_group.neg.default_port,
 *     ipAddress: endpoint_instance.networkInterfaces.apply(networkInterfaces => networkInterfaces[0].networkIp),
 * });
 * const group = new gcp.compute.NetworkEndpointGroup("group", {
 *     network: defaultNetwork.id,
 *     subnetwork: defaultSubnetwork.id,
 *     defaultPort: "90",
 *     zone: "us-central1-a",
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class NetworkEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing NetworkEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkEndpointState, opts?: pulumi.CustomResourceOptions): NetworkEndpoint {
        return new NetworkEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/networkEndpoint:NetworkEndpoint';

    /**
     * Returns true if the given object is an instance of NetworkEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkEndpoint.__pulumiType;
    }

    /**
     * The name for a specific VM instance that the IP address belongs to.
     * This is required for network endpoints of type GCE_VM_IP_PORT.
     * The instance must be in the same zone of network endpoint group.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * IPv4 address of network endpoint. The IP address must belong
     * to a VM in GCE (either the primary IP or as part of an aliased IP
     * range).
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The network endpoint group this endpoint is part of.
     */
    public readonly networkEndpointGroup!: pulumi.Output<string>;
    /**
     * Port number of network endpoint.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Zone where the containing network endpoint group is located.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NetworkEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkEndpointArgs | NetworkEndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkEndpointState | undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["ipAddress"] = state ? state.ipAddress : undefined;
            inputs["networkEndpointGroup"] = state ? state.networkEndpointGroup : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NetworkEndpointArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            if (!args || args.ipAddress === undefined) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if (!args || args.networkEndpointGroup === undefined) {
                throw new Error("Missing required property 'networkEndpointGroup'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            inputs["instance"] = args ? args.instance : undefined;
            inputs["ipAddress"] = args ? args.ipAddress : undefined;
            inputs["networkEndpointGroup"] = args ? args.networkEndpointGroup : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkEndpoint resources.
 */
export interface NetworkEndpointState {
    /**
     * The name for a specific VM instance that the IP address belongs to.
     * This is required for network endpoints of type GCE_VM_IP_PORT.
     * The instance must be in the same zone of network endpoint group.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * IPv4 address of network endpoint. The IP address must belong
     * to a VM in GCE (either the primary IP or as part of an aliased IP
     * range).
     */
    readonly ipAddress?: pulumi.Input<string>;
    /**
     * The network endpoint group this endpoint is part of.
     */
    readonly networkEndpointGroup?: pulumi.Input<string>;
    /**
     * Port number of network endpoint.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Zone where the containing network endpoint group is located.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkEndpoint resource.
 */
export interface NetworkEndpointArgs {
    /**
     * The name for a specific VM instance that the IP address belongs to.
     * This is required for network endpoints of type GCE_VM_IP_PORT.
     * The instance must be in the same zone of network endpoint group.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * IPv4 address of network endpoint. The IP address must belong
     * to a VM in GCE (either the primary IP or as part of an aliased IP
     * range).
     */
    readonly ipAddress: pulumi.Input<string>;
    /**
     * The network endpoint group this endpoint is part of.
     */
    readonly networkEndpointGroup: pulumi.Input<string>;
    /**
     * Port number of network endpoint.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Zone where the containing network endpoint group is located.
     */
    readonly zone?: pulumi.Input<string>;
}
