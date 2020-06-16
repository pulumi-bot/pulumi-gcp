// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a TargetInstance resource which defines an endpoint instance
 * that terminates traffic of certain protocols. In particular, they are used
 * in Protocol Forwarding, where forwarding rules can send packets to a
 * non-NAT'ed target instance. Each target instance contains a single
 * virtual machine instance that receives and handles traffic from the
 * corresponding forwarding rules.
 *
 * To get more information about TargetInstance, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetInstances)
 * * How-to Guides
 *     * [Using Protocol Forwarding](https://cloud.google.com/compute/docs/protocol-forwarding)
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Target Instance Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const vmimage = gcp.compute.getImage({
 *     family: "debian-9",
 *     project: "debian-cloud",
 * });
 * const target_vm = new gcp.compute.Instance("target-vm", {
 *     machineType: "n1-standard-1",
 *     zone: "us-central1-a",
 *     boot_disk: {
 *         initialize_params: {
 *             image: vmimage.then(vmimage => vmimage.selfLink),
 *         },
 *     },
 *     network_interface: [{
 *         network: "default",
 *     }],
 * });
 * const _default = new gcp.compute.TargetInstance("default", {instance: target_vm.id});
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class TargetInstance extends pulumi.CustomResource {
    /**
     * Get an existing TargetInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetInstanceState, opts?: pulumi.CustomResourceOptions): TargetInstance {
        return new TargetInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetInstance:TargetInstance';

    /**
     * Returns true if the given object is an instance of TargetInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetInstance.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Compute instance VM handling traffic for this target instance.
     * Accepts the instance self-link, relative path
     * (e.g. `projects/project/zones/zone/instances/instance`) or name. If
     * name is given, the zone will default to the given zone or
     * the provider-default zone and the project will default to the
     * provider-level project.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * NAT option controlling how IPs are NAT'ed to the instance.
     * Currently only NO_NAT (default value) is supported.
     */
    public readonly natPolicy!: pulumi.Output<string | undefined>;
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
     * URL of the zone where the target instance resides.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a TargetInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetInstanceArgs | TargetInstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetInstanceState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natPolicy"] = state ? state.natPolicy : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as TargetInstanceArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natPolicy"] = args ? args.natPolicy : undefined;
            inputs["project"] = args ? args.project : undefined;
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
        super(TargetInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetInstance resources.
 */
export interface TargetInstanceState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Compute instance VM handling traffic for this target instance.
     * Accepts the instance self-link, relative path
     * (e.g. `projects/project/zones/zone/instances/instance`) or name. If
     * name is given, the zone will default to the given zone or
     * the provider-default zone and the project will default to the
     * provider-level project.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * NAT option controlling how IPs are NAT'ed to the instance.
     * Currently only NO_NAT (default value) is supported.
     */
    readonly natPolicy?: pulumi.Input<string>;
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
     * URL of the zone where the target instance resides.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetInstance resource.
 */
export interface TargetInstanceArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Compute instance VM handling traffic for this target instance.
     * Accepts the instance self-link, relative path
     * (e.g. `projects/project/zones/zone/instances/instance`) or name. If
     * name is given, the zone will default to the given zone or
     * the provider-default zone and the project will default to the
     * provider-level project.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * NAT option controlling how IPs are NAT'ed to the instance.
     * Currently only NO_NAT (default value) is supported.
     */
    readonly natPolicy?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * URL of the zone where the target instance resides.
     */
    readonly zone?: pulumi.Input<string>;
}
