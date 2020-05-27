// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Cloud TPU instance.
 *
 *
 * To get more information about Node, see:
 *
 * * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/tpu/docs/)
 *
 * ## Example Usage - TPU Node Basic
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const available = gcp.tpu.getTensorflowVersions({});
 * const tpu = new gcp.tpu.Node("tpu", {
 *     zone: "us-central1-b",
 *     acceleratorType: "v3-8",
 *     tensorflowVersion: available.then(available => available.versions[0]),
 *     cidrBlock: "10.2.0.0/29",
 * });
 * ```
 * {{% /example %}}
 * ## Example Usage - TPU Node Full
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const available = gcp.tpu.getTensorflowVersions({});
 * const tpu = new gcp.tpu.Node("tpu", {
 *     zone: "us-central1-b",
 *     acceleratorType: "v3-8",
 *     cidrBlock: "10.3.0.0/29",
 *     tensorflowVersion: available.then(available => available.versions[0]),
 *     description: "Google Provider test TPU",
 *     network: "default",
 *     labels: {
 *         foo: "bar",
 *     },
 *     scheduling_config: {
 *         preemptible: true,
 *     },
 * });
 * ```
 *
 * {{% /example %}}
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeState, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:tpu/node:Node';

    /**
     * Returns true if the given object is an instance of Node.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Node {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Node.__pulumiType;
    }

    /**
     * The type of hardware accelerators associated with this node.
     */
    public readonly acceleratorType!: pulumi.Output<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP
     * address. This CIDR block must be a /29 block; the Compute Engine
     * networks API forbids a smaller block, and using a larger block would
     * be wasteful (a node can only consume one IP address).
     * Errors will occur if the CIDR block has already been used for a
     * currently existing TPU node, the CIDR block conflicts with any
     * subnetworks in the user's provided network, or the provided network
     * is peered with another network that is using that CIDR block.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource labels to represent user provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The immutable name of the TPU.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of a network to peer the TPU node to. It must be a
     * preexisting Compute Engine network inside of the project on which
     * this API has been activated. If none is provided, "default" will be
     * used.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of
     * the node first reach out to the first (index 0) entry.
     */
    public /*out*/ readonly networkEndpoints!: pulumi.Output<outputs.tpu.NodeNetworkEndpoint[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Sets the scheduling options for this TPU instance.  Structure is documented below.
     */
    public readonly schedulingConfig!: pulumi.Output<outputs.tpu.NodeSchedulingConfig | undefined>;
    /**
     * The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
     * Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
     */
    public /*out*/ readonly serviceAccount!: pulumi.Output<string>;
    /**
     * The version of Tensorflow running in the Node.
     */
    public readonly tensorflowVersion!: pulumi.Output<string>;
    /**
     * The GCP location for the TPU.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeArgs | NodeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NodeState | undefined;
            inputs["acceleratorType"] = state ? state.acceleratorType : undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkEndpoints"] = state ? state.networkEndpoints : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schedulingConfig"] = state ? state.schedulingConfig : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["tensorflowVersion"] = state ? state.tensorflowVersion : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NodeArgs | undefined;
            if (!args || args.acceleratorType === undefined) {
                throw new Error("Missing required property 'acceleratorType'");
            }
            if (!args || args.cidrBlock === undefined) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if (!args || args.tensorflowVersion === undefined) {
                throw new Error("Missing required property 'tensorflowVersion'");
            }
            if (!args || args.zone === undefined) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["acceleratorType"] = args ? args.acceleratorType : undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedulingConfig"] = args ? args.schedulingConfig : undefined;
            inputs["tensorflowVersion"] = args ? args.tensorflowVersion : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["networkEndpoints"] = undefined /*out*/;
            inputs["serviceAccount"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Node.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Node resources.
 */
export interface NodeState {
    /**
     * The type of hardware accelerators associated with this node.
     */
    readonly acceleratorType?: pulumi.Input<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP
     * address. This CIDR block must be a /29 block; the Compute Engine
     * networks API forbids a smaller block, and using a larger block would
     * be wasteful (a node can only consume one IP address).
     * Errors will occur if the CIDR block has already been used for a
     * currently existing TPU node, the CIDR block conflicts with any
     * subnetworks in the user's provided network, or the provided network
     * is peered with another network that is using that CIDR block.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The immutable name of the TPU.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of a network to peer the TPU node to. It must be a
     * preexisting Compute Engine network inside of the project on which
     * this API has been activated. If none is provided, "default" will be
     * used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of
     * the node first reach out to the first (index 0) entry.
     */
    readonly networkEndpoints?: pulumi.Input<pulumi.Input<inputs.tpu.NodeNetworkEndpoint>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Sets the scheduling options for this TPU instance.  Structure is documented below.
     */
    readonly schedulingConfig?: pulumi.Input<inputs.tpu.NodeSchedulingConfig>;
    /**
     * The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
     * Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
     */
    readonly serviceAccount?: pulumi.Input<string>;
    /**
     * The version of Tensorflow running in the Node.
     */
    readonly tensorflowVersion?: pulumi.Input<string>;
    /**
     * The GCP location for the TPU.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * The type of hardware accelerators associated with this node.
     */
    readonly acceleratorType: pulumi.Input<string>;
    /**
     * The CIDR block that the TPU node will use when selecting an IP
     * address. This CIDR block must be a /29 block; the Compute Engine
     * networks API forbids a smaller block, and using a larger block would
     * be wasteful (a node can only consume one IP address).
     * Errors will occur if the CIDR block has already been used for a
     * currently existing TPU node, the CIDR block conflicts with any
     * subnetworks in the user's provided network, or the provided network
     * is peered with another network that is using that CIDR block.
     */
    readonly cidrBlock: pulumi.Input<string>;
    /**
     * The user-supplied description of the TPU. Maximum of 512 characters.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The immutable name of the TPU.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of a network to peer the TPU node to. It must be a
     * preexisting Compute Engine network inside of the project on which
     * this API has been activated. If none is provided, "default" will be
     * used.
     */
    readonly network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Sets the scheduling options for this TPU instance.  Structure is documented below.
     */
    readonly schedulingConfig?: pulumi.Input<inputs.tpu.NodeSchedulingConfig>;
    /**
     * The version of Tensorflow running in the Node.
     */
    readonly tensorflowVersion: pulumi.Input<string>;
    /**
     * The GCP location for the TPU.
     */
    readonly zone: pulumi.Input<string>;
}
