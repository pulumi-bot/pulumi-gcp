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
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tplInstanceTemplate = new gcp.compute.InstanceTemplate("tplInstanceTemplate", {
 *     machineType: "n1-standard-1",
 *     disk: [{
 *         sourceImage: "debian-cloud/debian-9",
 *         autoDelete: true,
 *         diskSizeGb: 100,
 *         boot: true,
 *     }],
 *     network_interface: [{
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
 * {{% /example %}}
 * {{% /examples %}}
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

    public readonly allowStoppingForUpdate!: pulumi.Output<boolean>;
    public readonly attachedDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateAttachedDisk[]>;
    public readonly bootDisk!: pulumi.Output<outputs.compute.InstanceFromTemplateBootDisk>;
    public readonly canIpForward!: pulumi.Output<boolean>;
    public /*out*/ readonly cpuPlatform!: pulumi.Output<string>;
    public /*out*/ readonly currentStatus!: pulumi.Output<string>;
    public readonly deletionProtection!: pulumi.Output<boolean>;
    public readonly description!: pulumi.Output<string>;
    public readonly desiredStatus!: pulumi.Output<string>;
    public readonly enableDisplay!: pulumi.Output<boolean>;
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.InstanceFromTemplateGuestAccelerator[]>;
    public readonly hostname!: pulumi.Output<string>;
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly machineType!: pulumi.Output<string>;
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly metadataFingerprint!: pulumi.Output<string>;
    public readonly metadataStartupScript!: pulumi.Output<string>;
    public readonly minCpuPlatform!: pulumi.Output<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.InstanceFromTemplateNetworkInterface[]>;
    public readonly project!: pulumi.Output<string>;
    public readonly resourcePolicies!: pulumi.Output<string>;
    public readonly scheduling!: pulumi.Output<outputs.compute.InstanceFromTemplateScheduling>;
    public readonly scratchDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateScratchDisk[]>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly serviceAccount!: pulumi.Output<outputs.compute.InstanceFromTemplateServiceAccount>;
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    public readonly sourceInstanceTemplate!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[]>;
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
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    readonly bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly cpuPlatform?: pulumi.Input<string>;
    readonly currentStatus?: pulumi.Input<string>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly desiredStatus?: pulumi.Input<string>;
    readonly enableDisplay?: pulumi.Input<boolean>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly instanceId?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataFingerprint?: pulumi.Input<string>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    readonly project?: pulumi.Input<string>;
    readonly resourcePolicies?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    readonly selfLink?: pulumi.Input<string>;
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
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
    readonly allowStoppingForUpdate?: pulumi.Input<boolean>;
    readonly attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    readonly bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    readonly canIpForward?: pulumi.Input<boolean>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    readonly description?: pulumi.Input<string>;
    readonly desiredStatus?: pulumi.Input<string>;
    readonly enableDisplay?: pulumi.Input<boolean>;
    readonly guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    readonly hostname?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly machineType?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly metadataStartupScript?: pulumi.Input<string>;
    readonly minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    readonly project?: pulumi.Input<string>;
    readonly resourcePolicies?: pulumi.Input<string>;
    readonly scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    readonly scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    readonly serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    readonly shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    readonly sourceInstanceTemplate: pulumi.Input<string>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    readonly zone?: pulumi.Input<string>;
}
