// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a NodeGroup resource to manage a group of sole-tenant nodes.
 *
 * To get more information about NodeGroup, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
 * * How-to Guides
 *     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
 *
 * > **Warning:** Due to limitations of the API, this provider cannot update the
 * number of nodes in a node group and changes to node group size either
 * through provider config or through external changes will cause
 * the provider to delete and recreate the node group.
 *
 * ## Example Usage
 * ### Node Group Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const central1a = gcp.compute.getNodeTypes({
 *     zone: "us-central1-a",
 * });
 * const soletenant_tmpl = new gcp.compute.NodeTemplate("soletenant-tmpl", {
 *     region: "us-central1",
 *     nodeType: central1a.then(central1a => central1a.names[0]),
 * });
 * const nodes = new gcp.compute.NodeGroup("nodes", {
 *     zone: "us-central1-a",
 *     description: "example google_compute_node_group for the Google Provider",
 *     size: 1,
 *     nodeTemplate: soletenant_tmpl.id,
 * });
 * ```
 * ### Node Group Autoscaling Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const central1a = gcp.compute.getNodeTypes({
 *     zone: "us-central1-a",
 * });
 * const soletenant_tmpl = new gcp.compute.NodeTemplate("soletenant-tmpl", {
 *     region: "us-central1",
 *     nodeType: central1a.then(central1a => central1a.names[0]),
 * });
 * const nodes = new gcp.compute.NodeGroup("nodes", {
 *     zone: "us-central1-a",
 *     description: "example google_compute_node_group for the Google Provider",
 *     size: 1,
 *     nodeTemplate: soletenant_tmpl.id,
 *     autoscaling_policy: {
 *         mode: "ON",
 *         minNodes: 1,
 *         maxNodes: 10,
 *     },
 * });
 * ```
 */
export class NodeGroup extends pulumi.CustomResource {
    /**
     * Get an existing NodeGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeGroupState, opts?: pulumi.CustomResourceOptions): NodeGroup {
        return new NodeGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/nodeGroup:NodeGroup';

    /**
     * Returns true if the given object is an instance of NodeGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeGroup.__pulumiType;
    }

    /**
     * -
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
     */
    public readonly autoscalingPolicy!: pulumi.Output<outputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional textual description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    public readonly nodeTemplate!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The total number of nodes in the node group.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Zone where this node group is located
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NodeGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeGroupArgs | NodeGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NodeGroupState | undefined;
            inputs["autoscalingPolicy"] = state ? state.autoscalingPolicy : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeTemplate"] = state ? state.nodeTemplate : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NodeGroupArgs | undefined;
            if (!args || args.nodeTemplate === undefined) {
                throw new Error("Missing required property 'nodeTemplate'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["autoscalingPolicy"] = args ? args.autoscalingPolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeTemplate"] = args ? args.nodeTemplate : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NodeGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodeGroup resources.
 */
export interface NodeGroupState {
    /**
     * -
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
     */
    readonly autoscalingPolicy?: pulumi.Input<inputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional textual description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    readonly nodeTemplate?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The total number of nodes in the node group.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * Zone where this node group is located
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodeGroup resource.
 */
export interface NodeGroupArgs {
    /**
     * -
     * If you use sole-tenant nodes for your workloads, you can use the node
     * group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
     */
    readonly autoscalingPolicy?: pulumi.Input<inputs.compute.NodeGroupAutoscalingPolicy>;
    /**
     * An optional textual description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The URL of the node template to which this node group belongs.
     */
    readonly nodeTemplate: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The total number of nodes in the node group.
     */
    readonly size: pulumi.Input<number>;
    /**
     * Zone where this node group is located
     */
    readonly zone?: pulumi.Input<string>;
}
