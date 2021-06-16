// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a group of dissimilar Compute Engine virtual machine instances.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
 *
 * ## Example Usage
 * ### Empty Instance Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const test = new gcp.compute.InstanceGroup("test", {
 *     description: "Test instance group",
 *     zone: "us-central1-a",
 *     network: google_compute_network["default"].id,
 * });
 * ```
 * ### Example Usage - With instances and named ports
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const webservers = new gcp.compute.InstanceGroup("webservers", {
 *     description: "Test instance group",
 *     instances: [
 *         google_compute_instance.test.id,
 *         google_compute_instance.test2.id,
 *     ],
 *     namedPorts: [
 *         {
 *             name: "http",
 *             port: "8080",
 *         },
 *         {
 *             name: "https",
 *             port: "8443",
 *         },
 *     ],
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * ## Import
 *
 * Instance group can be imported using the `zone` and `name` with an optional `project`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers us-central1-a/terraform-webservers
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers big-project/us-central1-a/terraform-webservers
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers projects/big-project/zones/us-central1-a/instanceGroups/terraform-webservers
 * ```
 */
export class InstanceGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupState, opts?: pulumi.CustomResourceOptions): InstanceGroup {
        return new InstanceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceGroup:InstanceGroup';

    /**
     * Returns true if the given object is an instance of InstanceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGroup.__pulumiType;
    }

    /**
     * An optional textual description of the instance
     * group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of instances in the group. They should be given
     * as either selfLink or id. When adding instances they must all be in the same
     * network and zone as the instance group.
     */
    public readonly instances!: pulumi.Output<string[]>;
    /**
     * The name which the port will be mapped to.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.InstanceGroupNamedPort[] | undefined>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The number of instances in the group.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The zone that this instance group should be created in.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGroupArgs | InstanceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namedPorts"] = state ? state.namedPorts : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namedPorts"] = args ? args.namedPorts : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["selfLink"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InstanceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroup resources.
 */
export interface InstanceGroupState {
    /**
     * An optional textual description of the instance
     * group.
     */
    description?: pulumi.Input<string>;
    /**
     * List of instances in the group. They should be given
     * as either selfLink or id. When adding instances they must all be in the same
     * network and zone as the instance group.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name which the port will be mapped to.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.InstanceGroupNamedPort>[]>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The number of instances in the group.
     */
    size?: pulumi.Input<number>;
    /**
     * The zone that this instance group should be created in.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroup resource.
 */
export interface InstanceGroupArgs {
    /**
     * An optional textual description of the instance
     * group.
     */
    description?: pulumi.Input<string>;
    /**
     * List of instances in the group. They should be given
     * as either selfLink or id. When adding instances they must all be in the same
     * network and zone as the instance group.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name which the port will be mapped to.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.InstanceGroupNamedPort>[]>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The zone that this instance group should be created in.
     */
    zone?: pulumi.Input<string>;
}
