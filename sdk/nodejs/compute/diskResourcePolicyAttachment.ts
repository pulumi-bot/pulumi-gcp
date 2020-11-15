// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Adds existing resource policies to a disk. You can only add one policy
 * which will be applied to this disk for scheduling snapshot creation.
 *
 * > **Note:** This resource does not support regional disks (`gcp.compute.RegionDisk`). For regional disks, please refer to the `gcp.compute.RegionDiskResourcePolicyAttachment` resource.
 *
 * ## Example Usage
 *
 * ## Import
 *
 * DiskResourcePolicyAttachment can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment default projects/{{project}}/zones/{{zone}}/disks/{{disk}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment default {{project}}/{{zone}}/{{disk}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment default {{zone}}/{{disk}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment default {{disk}}/{{name}}
 * ```
 */
export class DiskResourcePolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing DiskResourcePolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskResourcePolicyAttachmentState, opts?: pulumi.CustomResourceOptions): DiskResourcePolicyAttachment {
        return new DiskResourcePolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment';

    /**
     * Returns true if the given object is an instance of DiskResourcePolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskResourcePolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskResourcePolicyAttachment.__pulumiType;
    }

    /**
     * The name of the disk in which the resource policies are attached to.
     */
    public readonly disk!: pulumi.Output<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A reference to the zone where the disk resides.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a DiskResourcePolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskResourcePolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskResourcePolicyAttachmentArgs | DiskResourcePolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DiskResourcePolicyAttachmentState | undefined;
            inputs["disk"] = state ? state.disk : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DiskResourcePolicyAttachmentArgs | undefined;
            if (!args || args.disk === undefined) {
                throw new Error("Missing required property 'disk'");
            }
            inputs["disk"] = args ? args.disk : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DiskResourcePolicyAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskResourcePolicyAttachment resources.
 */
export interface DiskResourcePolicyAttachmentState {
    /**
     * The name of the disk in which the resource policies are attached to.
     */
    readonly disk?: pulumi.Input<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the zone where the disk resides.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DiskResourcePolicyAttachment resource.
 */
export interface DiskResourcePolicyAttachmentArgs {
    /**
     * The name of the disk in which the resource policies are attached to.
     */
    readonly disk: pulumi.Input<string>;
    /**
     * The resource policy to be attached to the disk for scheduling snapshot
     * creation. Do not specify the self link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A reference to the zone where the disk resides.
     */
    readonly zone?: pulumi.Input<string>;
}
