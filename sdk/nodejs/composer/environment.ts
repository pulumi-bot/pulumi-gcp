// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentState, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:composer/environment:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    public readonly config!: pulumi.Output<{ airflowUri: string, dagGcsPrefix: string, gkeCluster: string, nodeConfig: { diskSizeGb: number, machineType: string, network: string, oauthScopes: string[], serviceAccount: string, subnetwork?: string, tags?: string[], zone: string }, nodeCount: number, softwareConfig: { airflowConfigOverrides?: {[key: string]: string}, envVariables?: {[key: string]: string}, imageVersion: string, pypiPackages?: {[key: string]: string}, pythonVersion: string } }>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs | EnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EnvironmentState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as EnvironmentArgs | undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Environment resources.
 */
export interface EnvironmentState {
    readonly config?: pulumi.Input<{ airflowUri?: pulumi.Input<string>, dagGcsPrefix?: pulumi.Input<string>, gkeCluster?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, machineType?: pulumi.Input<string>, network?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, serviceAccount?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone: pulumi.Input<string> }>, nodeCount?: pulumi.Input<number>, softwareConfig?: pulumi.Input<{ airflowConfigOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, imageVersion?: pulumi.Input<string>, pypiPackages?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pythonVersion?: pulumi.Input<string> }> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    readonly config?: pulumi.Input<{ airflowUri?: pulumi.Input<string>, dagGcsPrefix?: pulumi.Input<string>, gkeCluster?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, machineType?: pulumi.Input<string>, network?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, serviceAccount?: pulumi.Input<string>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone: pulumi.Input<string> }>, nodeCount?: pulumi.Input<number>, softwareConfig?: pulumi.Input<{ airflowConfigOverrides?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, imageVersion?: pulumi.Input<string>, pypiPackages?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, pythonVersion?: pulumi.Input<string> }> }>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
}
