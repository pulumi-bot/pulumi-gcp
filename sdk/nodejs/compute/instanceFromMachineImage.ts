// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class InstanceFromMachineImage extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFromMachineImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceFromMachineImageState, opts?: pulumi.CustomResourceOptions): InstanceFromMachineImage {
        return new InstanceFromMachineImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceFromMachineImage:InstanceFromMachineImage';

    /**
     * Returns true if the given object is an instance of InstanceFromMachineImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFromMachineImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFromMachineImage.__pulumiType;
    }

    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    public readonly allowStoppingForUpdate!: pulumi.Output<boolean>;
    /**
     * List of disks attached to the instance
     */
    public /*out*/ readonly attachedDisks!: pulumi.Output<outputs.compute.InstanceFromMachineImageAttachedDisk[]>;
    /**
     * The boot disk for the instance.
     */
    public /*out*/ readonly bootDisks!: pulumi.Output<outputs.compute.InstanceFromMachineImageBootDisk[]>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    public readonly canIpForward!: pulumi.Output<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    public readonly confidentialInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromMachineImageConfidentialInstanceConfig>;
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
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.InstanceFromMachineImageGuestAccelerator[]>;
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
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.InstanceFromMachineImageNetworkInterface[]>;
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
    public readonly scheduling!: pulumi.Output<outputs.compute.InstanceFromMachineImageScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    public /*out*/ readonly scratchDisks!: pulumi.Output<outputs.compute.InstanceFromMachineImageScratchDisk[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The service account to attach to the instance.
     */
    public readonly serviceAccount!: pulumi.Output<outputs.compute.InstanceFromMachineImageServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromMachineImageShieldedInstanceConfig>;
    /**
     * Name or self link of a machine
     * image to create the instance based on.
     */
    public readonly sourceMachineImage!: pulumi.Output<string>;
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
     * Create a InstanceFromMachineImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFromMachineImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceFromMachineImageArgs | InstanceFromMachineImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceFromMachineImageState | undefined;
            inputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            inputs["attachedDisks"] = state ? state.attachedDisks : undefined;
            inputs["bootDisks"] = state ? state.bootDisks : undefined;
            inputs["canIpForward"] = state ? state.canIpForward : undefined;
            inputs["confidentialInstanceConfig"] = state ? state.confidentialInstanceConfig : undefined;
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
            inputs["sourceMachineImage"] = state ? state.sourceMachineImage : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceFromMachineImageArgs | undefined;
            if (!args || args.sourceMachineImage === undefined) {
                throw new Error("Missing required property 'sourceMachineImage'");
            }
            inputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["confidentialInstanceConfig"] = args ? args.confidentialInstanceConfig : undefined;
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
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            inputs["sourceMachineImage"] = args ? args.sourceMachineImage : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["attachedDisks"] = undefined /*out*/;
            inputs["bootDisks"] = undefined /*out*/;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["currentStatus"] = undefined /*out*/;
            inputs["instanceId"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["metadataFingerprint"] = undefined /*out*/;
            inputs["scratchDisks"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["tagsFingerprint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceFromMachineImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceFromMachineImage resources.
 */
export interface InstanceFromMachineImageState {
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * List of disks attached to the instance
     */
    readonly attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageAttachedDisk>[]>;
    /**
     * The boot disk for the instance.
     */
    readonly bootDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageBootDisk>[]>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    readonly confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromMachineImageConfidentialInstanceConfig>;
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
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageGuestAccelerator>[]>;
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
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageNetworkInterface>[]>;
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
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromMachineImageScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    readonly scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageScratchDisk>[]>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * The service account to attach to the instance.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromMachineImageServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromMachineImageShieldedInstanceConfig>;
    /**
     * Name or self link of a machine
     * image to create the instance based on.
     */
    readonly sourceMachineImage?: pulumi.Input<string>;
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
 * The set of arguments for constructing a InstanceFromMachineImage resource.
 */
export interface InstanceFromMachineImageArgs {
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    readonly canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    readonly confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromMachineImageConfidentialInstanceConfig>;
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
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageGuestAccelerator>[]>;
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
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromMachineImageNetworkInterface>[]>;
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
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromMachineImageScheduling>;
    /**
     * The service account to attach to the instance.
     */
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromMachineImageServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromMachineImageShieldedInstanceConfig>;
    /**
     * Name or self link of a machine
     * image to create the instance based on.
     */
    readonly sourceMachineImage: pulumi.Input<string>;
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
