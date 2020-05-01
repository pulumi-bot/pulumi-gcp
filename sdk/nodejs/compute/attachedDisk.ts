// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Persistent disks can be attached to a compute instance using the `attachedDisk`
 * section within the compute instance configuration.
 * However there may be situations where managing the attached disks via the compute
 * instance config isn't preferable or possible, such as attaching dynamic
 * numbers of disks using the `count` variable.
 * 
 * 
 * To get more information about attaching disks, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
 * * How-to Guides
 *     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
 * 
 * **Note:** When using `gcp.compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attachedDisk"]` on the `gcp.compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultInstance = new gcp.compute.Instance("defaultInstance", {
 *     machineType: "n1-standard-1",
 *     zone: "us-west1-a",
 *     boot_disk: {
 *         initialize_params: {
 *             image: "debian-cloud/debian-9",
 *         },
 *     },
 *     network_interface: [{
 *         network: "default",
 *     }],
 * });
 * const defaultAttachedDisk = new gcp.compute.AttachedDisk("defaultAttachedDisk", {
 *     disk: google_compute_disk["default"].id,
 *     instance: defaultInstance.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_attached_disk.html.markdown.
 */
export class AttachedDisk extends pulumi.CustomResource {
    /**
     * Get an existing AttachedDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachedDiskState, opts?: pulumi.CustomResourceOptions): AttachedDisk {
        return new AttachedDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/attachedDisk:AttachedDisk';

    /**
     * Returns true if the given object is an instance of AttachedDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttachedDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttachedDisk.__pulumiType;
    }

    /**
     * Specifies a unique device name of your choice that is
     * reflected into the /dev/disk/by-id/google-* tree of a Linux operating
     * system running within the instance. This name can be used to
     * reference the device for mounting, resizing, and so on, from within
     * the instance.
     */
    public readonly deviceName!: pulumi.Output<string>;
    /**
     * `name` or `selfLink` of the disk that will be attached.
     */
    public readonly disk!: pulumi.Output<string>;
    /**
     * `name` or `selfLink` of the compute instance that the disk will be attached to.
     * If the `selfLink` is provided then `zone` and `project` are extracted from the
     * self link. If only the name is used then `zone` and `project` must be defined
     * as properties on the resource or provider.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * The mode in which to attach this disk, either READ_WRITE or
     * READ_ONLY. If not specified, the default is to attach the disk in
     * READ_WRITE mode.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The project that the referenced compute instance is a part of. If `instance` is referenced by its
     * `selfLink` the project defined in the link will take precedence.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The zone that the referenced compute instance is located within. If `instance` is referenced by its
     * `selfLink` the zone defined in the link will take precedence.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a AttachedDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachedDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachedDiskArgs | AttachedDiskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttachedDiskState | undefined;
            inputs["deviceName"] = state ? state.deviceName : undefined;
            inputs["disk"] = state ? state.disk : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as AttachedDiskArgs | undefined;
            if (!args || args.disk === undefined) {
                throw new Error("Missing required property 'disk'");
            }
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["disk"] = args ? args.disk : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AttachedDisk.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttachedDisk resources.
 */
export interface AttachedDiskState {
    /**
     * Specifies a unique device name of your choice that is
     * reflected into the /dev/disk/by-id/google-* tree of a Linux operating
     * system running within the instance. This name can be used to
     * reference the device for mounting, resizing, and so on, from within
     * the instance.
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * `name` or `selfLink` of the disk that will be attached.
     */
    readonly disk?: pulumi.Input<string>;
    /**
     * `name` or `selfLink` of the compute instance that the disk will be attached to.
     * If the `selfLink` is provided then `zone` and `project` are extracted from the
     * self link. If only the name is used then `zone` and `project` must be defined
     * as properties on the resource or provider.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * The mode in which to attach this disk, either READ_WRITE or
     * READ_ONLY. If not specified, the default is to attach the disk in
     * READ_WRITE mode.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The project that the referenced compute instance is a part of. If `instance` is referenced by its
     * `selfLink` the project defined in the link will take precedence.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone that the referenced compute instance is located within. If `instance` is referenced by its
     * `selfLink` the zone defined in the link will take precedence.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttachedDisk resource.
 */
export interface AttachedDiskArgs {
    /**
     * Specifies a unique device name of your choice that is
     * reflected into the /dev/disk/by-id/google-* tree of a Linux operating
     * system running within the instance. This name can be used to
     * reference the device for mounting, resizing, and so on, from within
     * the instance.
     */
    readonly deviceName?: pulumi.Input<string>;
    /**
     * `name` or `selfLink` of the disk that will be attached.
     */
    readonly disk: pulumi.Input<string>;
    /**
     * `name` or `selfLink` of the compute instance that the disk will be attached to.
     * If the `selfLink` is provided then `zone` and `project` are extracted from the
     * self link. If only the name is used then `zone` and `project` must be defined
     * as properties on the resource or provider.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * The mode in which to attach this disk, either READ_WRITE or
     * READ_ONLY. If not specified, the default is to attach the disk in
     * READ_WRITE mode.
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * The project that the referenced compute instance is a part of. If `instance` is referenced by its
     * `selfLink` the project defined in the link will take precedence.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The zone that the referenced compute instance is located within. If `instance` is referenced by its
     * `selfLink` the zone defined in the link will take precedence.
     */
    readonly zone?: pulumi.Input<string>;
}
