// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    public readonly cluster: pulumi.Output<{ clusterId?: string, numNodes?: number, storageType?: string, zone: string } | undefined>;
    public readonly clusterId: pulumi.Output<string | undefined>;
    public readonly displayName: pulumi.Output<string>;
    public readonly instanceType: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly numNodes: pulumi.Output<number | undefined>;
    public readonly project: pulumi.Output<string>;
    public readonly storageType: pulumi.Output<string | undefined>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numNodes"] = state ? state.numNodes : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["storageType"] = state ? state.storageType : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["numNodes"] = args ? args.numNodes : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        super("gcp:bigtable/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    readonly cluster?: pulumi.Input<{ clusterId?: pulumi.Input<string>, numNodes?: pulumi.Input<number>, storageType?: pulumi.Input<string>, zone?: pulumi.Input<string> }>;
    readonly clusterId?: pulumi.Input<string>;
    readonly displayName?: pulumi.Input<string>;
    readonly instanceType?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly numNodes?: pulumi.Input<number>;
    readonly project?: pulumi.Input<string>;
    readonly storageType?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    readonly cluster?: pulumi.Input<{ clusterId?: pulumi.Input<string>, numNodes?: pulumi.Input<number>, storageType?: pulumi.Input<string>, zone?: pulumi.Input<string> }>;
    readonly clusterId?: pulumi.Input<string>;
    readonly displayName?: pulumi.Input<string>;
    readonly instanceType?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly numNodes?: pulumi.Input<number>;
    readonly project?: pulumi.Input<string>;
    readonly storageType?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}
