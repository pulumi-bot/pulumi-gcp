// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Represents a NodeTemplate resource. Node templates specify properties
 * for creating sole-tenant nodes, such as node type, vCPU and memory
 * requirements, node affinity labels, and region.
 *
 * To get more information about NodeTemplate, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
 * * How-to Guides
 *     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
 *
 * ## Example Usage
 * ### Node Template Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const template = new gcp.compute.NodeTemplate("template", {
 *     nodeType: "n1-node-96-624",
 *     region: "us-central1",
 * });
 * ```
 * ### Node Template Server Binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const central1a = pulumi.output(gcp.compute.getNodeTypes({
 *     zone: "us-central1-a",
 * }, { async: true }));
 * const template = new gcp.compute.NodeTemplate("template", {
 *     nodeAffinityLabels: {
 *         foo: "baz",
 *     },
 *     nodeType: "n1-node-96-624",
 *     region: "us-central1",
 *     serverBinding: {
 *         type: "RESTART_NODE_ON_MINIMAL_SERVERS",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * NodeTemplate can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{name}}
 * ```
 */
export class NodeTemplate extends pulumi.CustomResource {
    /**
     * Get an existing NodeTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeTemplateState, opts?: pulumi.CustomResourceOptions): NodeTemplate {
        return new NodeTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/nodeTemplate:NodeTemplate';

    /**
     * Returns true if the given object is an instance of NodeTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodeTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodeTemplate.__pulumiType;
    }

    /**
     * CPU overcommit.
     * Default value is `NONE`.
     * Possible values are `ENABLED` and `NONE`.
     */
    public readonly cpuOvercommitType!: pulumi.Output<string | undefined>;
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
     * Labels to use for node affinity, which will be used in
     * instance scheduling.
     */
    public readonly nodeAffinityLabels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Node type to use for nodes group that are created from this template.
     * Only one of nodeTypeFlexibility and nodeType can be specified.
     */
    public readonly nodeType!: pulumi.Output<string | undefined>;
    /**
     * Flexible properties for the desired node type. Node groups that
     * use this node template will create nodes of a type that matches
     * these properties. Only one of nodeTypeFlexibility and nodeType can
     * be specified.
     * Structure is documented below.
     */
    public readonly nodeTypeFlexibility!: pulumi.Output<outputs.compute.NodeTemplateNodeTypeFlexibility | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region where nodes using the node template will be created.
     * If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The server binding policy for nodes using this template. Determines
     * where the nodes should restart following a maintenance event.
     * Structure is documented below.
     */
    public readonly serverBinding!: pulumi.Output<outputs.compute.NodeTemplateServerBinding>;

    /**
     * Create a NodeTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NodeTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeTemplateArgs | NodeTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeTemplateState | undefined;
            inputs["cpuOvercommitType"] = state ? state.cpuOvercommitType : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeAffinityLabels"] = state ? state.nodeAffinityLabels : undefined;
            inputs["nodeType"] = state ? state.nodeType : undefined;
            inputs["nodeTypeFlexibility"] = state ? state.nodeTypeFlexibility : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serverBinding"] = state ? state.serverBinding : undefined;
        } else {
            const args = argsOrState as NodeTemplateArgs | undefined;
            inputs["cpuOvercommitType"] = args ? args.cpuOvercommitType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeAffinityLabels"] = args ? args.nodeAffinityLabels : undefined;
            inputs["nodeType"] = args ? args.nodeType : undefined;
            inputs["nodeTypeFlexibility"] = args ? args.nodeTypeFlexibility : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serverBinding"] = args ? args.serverBinding : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NodeTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodeTemplate resources.
 */
export interface NodeTemplateState {
    /**
     * CPU overcommit.
     * Default value is `NONE`.
     * Possible values are `ENABLED` and `NONE`.
     */
    cpuOvercommitType?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional textual description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Labels to use for node affinity, which will be used in
     * instance scheduling.
     */
    nodeAffinityLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Node type to use for nodes group that are created from this template.
     * Only one of nodeTypeFlexibility and nodeType can be specified.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Flexible properties for the desired node type. Node groups that
     * use this node template will create nodes of a type that matches
     * these properties. Only one of nodeTypeFlexibility and nodeType can
     * be specified.
     * Structure is documented below.
     */
    nodeTypeFlexibility?: pulumi.Input<inputs.compute.NodeTemplateNodeTypeFlexibility>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region where nodes using the node template will be created.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The server binding policy for nodes using this template. Determines
     * where the nodes should restart following a maintenance event.
     * Structure is documented below.
     */
    serverBinding?: pulumi.Input<inputs.compute.NodeTemplateServerBinding>;
}

/**
 * The set of arguments for constructing a NodeTemplate resource.
 */
export interface NodeTemplateArgs {
    /**
     * CPU overcommit.
     * Default value is `NONE`.
     * Possible values are `ENABLED` and `NONE`.
     */
    cpuOvercommitType?: pulumi.Input<string>;
    /**
     * An optional textual description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Labels to use for node affinity, which will be used in
     * instance scheduling.
     */
    nodeAffinityLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Node type to use for nodes group that are created from this template.
     * Only one of nodeTypeFlexibility and nodeType can be specified.
     */
    nodeType?: pulumi.Input<string>;
    /**
     * Flexible properties for the desired node type. Node groups that
     * use this node template will create nodes of a type that matches
     * these properties. Only one of nodeTypeFlexibility and nodeType can
     * be specified.
     * Structure is documented below.
     */
    nodeTypeFlexibility?: pulumi.Input<inputs.compute.NodeTemplateNodeTypeFlexibility>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region where nodes using the node template will be created.
     * If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The server binding policy for nodes using this template. Determines
     * where the nodes should restart following a maintenance event.
     * Structure is documented below.
     */
    serverBinding?: pulumi.Input<inputs.compute.NodeTemplateServerBinding>;
}
