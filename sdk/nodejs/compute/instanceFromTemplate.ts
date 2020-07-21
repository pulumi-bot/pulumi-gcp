// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 *
 * This resource is specifically to create a compute instance from a given
 * `sourceInstanceTemplate`. To create an instance without a template, use the
 * `gcp.compute.Instance` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tplInstanceTemplate = new gcp.compute.InstanceTemplate("tplInstanceTemplate", {
 *     machineType: "n1-standard-1",
 *     disks: [{
 *         sourceImage: "debian-cloud/debian-9",
 *         autoDelete: true,
 *         diskSizeGb: 100,
 *         boot: true,
 *     }],
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     canIpForward: true,
 * });
 * const tplInstanceFromTemplate = new gcp.compute.InstanceFromTemplate("tplInstanceFromTemplate", {
 *     zone: "us-central1-a",
 *     sourceInstanceTemplate: tplInstanceTemplate.id,
 *     canIpForward: false,
 *     labels: {
 *         my_key: "my_value",
 *     },
 * });
 * ```
 */
export class InstanceFromTemplate extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFromTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions): InstanceFromTemplate {
        return new InstanceFromTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceFromTemplate:InstanceFromTemplate';

    /**
     * Returns true if the given object is an instance of InstanceFromTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFromTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFromTemplate.__pulumiType;
    }

    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    public readonly allowStoppingForUpdate!: pulumi.Output<boolean>;
    /**
     * List of disks attached to the instance
     */
    public readonly attachedDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateAttachedDisk[]>;
    /**
     * The boot disk for the instance.
     */
    public readonly bootDisk!: pulumi.Output<outputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    public readonly canIpForward!: pulumi.Output<boolean>;
    /**
     * The CPU platform used by this instance.
     */
    public /*out*/ readonly cpuPlatform!: pulumi.Output<string>;
    /**
     * Current status of the instance.
     */
    public /*out*/ readonly currentStatus!: pulumi.Output<string>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * A brief description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    public readonly desiredStatus!: pulumi.Output<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    public readonly enableDisplay!: pulumi.Output<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.InstanceFromTemplateGuestAccelerator[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The server-assigned unique identifier of this instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * The unique fingerprint of the labels.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The machine type to create.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The unique fingerprint of the metadata.
     */
    public /*out*/ readonly metadataFingerprint!: pulumi.Output<string>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    public readonly metadataStartupScript!: pulumi.Output<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    public readonly minCpuPlatform!: pulumi.Output<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The networks attached to the instance.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.InstanceFromTemplateNetworkInterface[]>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
     * instance to recreate. Currently a max of 1 resource policy is supported.
     */
    public readonly resourcePolicies!: pulumi.Output<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    public readonly scheduling!: pulumi.Output<outputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    public readonly scratchDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateScratchDisk[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The service account to attach to the instance.
     */
    public readonly serviceAccount!: pulumi.Output<outputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    public readonly sourceInstanceTemplate!: pulumi.Output<string>;
    /**
     * The list of tags attached to the instance.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The unique fingerprint of the tags.
     */
    public /*out*/ readonly tagsFingerprint!: pulumi.Output<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceFromTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFromTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceFromTemplateArgs | InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceFromTemplateState | undefined;
            inputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = state ? state.attachedDisks : undefined;
            inputs["bootDisk"] = state ? state.bootDisk : undefined;
            inputs["canIpForward"] = state ? state.canIpForward : undefined;
            inputs["cpuPlatform"] = state ? state.cpuPlatform : undefined;
            inputs["currentStatus"] = state ? state.currentStatus : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["desiredStatus"] = state ? state.desiredStatus : undefined;
            inputs["enableDisplay"] = state ? state.enableDisplay : undefined;
            inputs["guestAccelerators"] = state ? state.guestAccelerators : undefined;
            inputs["hostname"] = state ? state.hostname : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["machineType"] = state ? state.machineType : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["metadataFingerprint"] = state ? state.metadataFingerprint : undefined;
            inputs["metadataStartupScript"] = state ? state.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = state ? state.minCpuPlatform : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["resourcePolicies"] = state ? state.resourcePolicies : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["scratchDisks"] = state ? state.scratchDisks : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["shieldedInstanceConfig"] = state ? state.shieldedInstanceConfig : undefined;
            inputs["sourceInstanceTemplate"] = state ? state.sourceInstanceTemplate : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceFromTemplateArgs | undefined;
            if (!args || args.sourceInstanceTemplate === undefined) {
                throw new Error("Missing required property 'sourceInstanceTemplate'");
            }
            inputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = args ? args.attachedDisks : undefined;
            inputs["bootDisk"] = args ? args.bootDisk : undefined;
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["desiredStatus"] = args ? args.desiredStatus : undefined;
            inputs["enableDisplay"] = args ? args.enableDisplay : undefined;
            inputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["metadataStartupScript"] = args ? args.metadataStartupScript : undefined;
            inputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["scratchDisks"] = args ? args.scratchDisks : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            inputs["sourceInstanceTemplate"] = args ? args.sourceInstanceTemplate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["currentStatus"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["metadataFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["tagsFingerprint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceFromTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceFromTemplate resources.
 */
export interface InstanceFromTemplateState {
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * List of disks attached to the instance
     */
    readonly attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    /**
     * The boot disk for the instance.
     */
    readonly bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * The CPU platform used by this instance.
     */
    readonly cpuPlatform?: pulumi.Input<string>;
    /**
     * Current status of the instance.
     */
    readonly currentStatus?: pulumi.Input<string>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * A brief description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    readonly desiredStatus?: pulumi.Input<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    readonly enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * The server-assigned unique identifier of this instance.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The unique fingerprint of the labels.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    readonly machineType?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique fingerprint of the metadata.
     */
    readonly metadataFingerprint?: pulumi.Input<string>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    readonly metadataStartupScript?: pulumi.Input<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The networks attached to the instance.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
     * instance to recreate. Currently a max of 1 resource policy is supported.
     */
    readonly resourcePolicies?: pulumi.Input<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    readonly scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The service account to attach to the instance.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate?: pulumi.Input<string>;
    /**
     * The list of tags attached to the instance.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique fingerprint of the tags.
     */
    readonly tagsFingerprint?: pulumi.Input<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceFromTemplate resource.
 */
export interface InstanceFromTemplateArgs {
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * List of disks attached to the instance
     */
    readonly attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    /**
     * The boot disk for the instance.
     */
    readonly bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * A brief description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    readonly desiredStatus?: pulumi.Input<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    readonly enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    readonly hostname?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    readonly machineType?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    readonly metadataStartupScript?: pulumi.Input<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The networks attached to the instance.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the
     * instance to recreate. Currently a max of 1 resource policy is supported.
     */
    readonly resourcePolicies?: pulumi.Input<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    readonly scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    /**
     * The service account to attach to the instance.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate: pulumi.Input<string>;
    /**
     * The list of tags attached to the instance.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}
