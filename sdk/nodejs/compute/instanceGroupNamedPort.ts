// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Mange the named ports setting for a managed instance group without
 * managing the group as whole. This resource is primarily intended for use
 * with GKE-generated groups that shouldn't otherwise be managed by other
 * tools.
 * 
 * 
 * To get more information about InstanceGroupNamedPort, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_group_named_port.html.markdown.
 */
export class InstanceGroupNamedPort extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroupNamedPort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupNamedPortState, opts?: pulumi.CustomResourceOptions): InstanceGroupNamedPort {
        return new InstanceGroupNamedPort(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort';

    /**
     * Returns true if the given object is an instance of InstanceGroupNamedPort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGroupNamedPort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGroupNamedPort.__pulumiType;
    }

    /**
     * -
     * (Required)
     * The name of the instance group.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * -
     * (Required)
     * The name for this named port. The name must be 1-63 characters
     * long, and comply with RFC1035.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * -
     * (Required)
     * The port number, which can be a value between 1 and 65535.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * -
     * (Optional)
     * The zone of the instance group.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceGroupNamedPort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceGroupNamedPortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGroupNamedPortArgs | InstanceGroupNamedPortState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceGroupNamedPortState | undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceGroupNamedPortArgs | undefined;
            if (!args || args.group === undefined) {
                throw new Error("Missing required property 'group'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            inputs["group"] = args ? args.group : undefined;
            inputs["name"] = args ? args.name : undefined;
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
        super(InstanceGroupNamedPort.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroupNamedPort resources.
 */
export interface InstanceGroupNamedPortState {
    /**
     * -
     * (Required)
     * The name of the instance group.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The name for this named port. The name must be 1-63 characters
     * long, and comply with RFC1035.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The port number, which can be a value between 1 and 65535.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * -
     * (Optional)
     * The zone of the instance group.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroupNamedPort resource.
 */
export interface InstanceGroupNamedPortArgs {
    /**
     * -
     * (Required)
     * The name of the instance group.
     */
    readonly group: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The name for this named port. The name must be 1-63 characters
     * long, and comply with RFC1035.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * -
     * (Required)
     * The port number, which can be a value between 1 and 65535.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * -
     * (Optional)
     * The zone of the instance group.
     */
    readonly zone?: pulumi.Input<string>;
}
