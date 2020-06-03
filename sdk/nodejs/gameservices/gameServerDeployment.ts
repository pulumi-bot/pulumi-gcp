// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A game server deployment resource.
 *
 * To get more information about GameServerDeployment, see:
 *
 * * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/game-servers/docs)
 *
 * ## Example Usage - Game Service Deployment Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.gameservices.GameServerDeployment("default", {
 *     deploymentId: "tf-test-deployment",
 *     description: "a deployment description",
 * });
 * ```
 */
export class GameServerDeployment extends pulumi.CustomResource {
    /**
     * Get an existing GameServerDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GameServerDeploymentState, opts?: pulumi.CustomResourceOptions): GameServerDeployment {
        return new GameServerDeployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gameservices/gameServerDeployment:GameServerDeployment';

    /**
     * Returns true if the given object is an instance of GameServerDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerDeployment.__pulumiType;
    }

    /**
     * A unique id for the deployment.
     */
    public readonly deploymentId!: pulumi.Output<string>;
    /**
     * Human readable description of the game server deployment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The labels associated with this game server deployment. Each label is a
     * key-value pair.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Location of the Deployment.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The resource id of the game server deployment, eg:
     * 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
     * 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a GameServerDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerDeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GameServerDeploymentArgs | GameServerDeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GameServerDeploymentState | undefined;
            inputs["deploymentId"] = state ? state.deploymentId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as GameServerDeploymentArgs | undefined;
            if (!args || args.deploymentId === undefined) {
                throw new Error("Missing required property 'deploymentId'");
            }
            inputs["deploymentId"] = args ? args.deploymentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GameServerDeployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GameServerDeployment resources.
 */
export interface GameServerDeploymentState {
    /**
     * A unique id for the deployment.
     */
    readonly deploymentId?: pulumi.Input<string>;
    /**
     * Human readable description of the game server deployment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The labels associated with this game server deployment. Each label is a
     * key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the Deployment.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource id of the game server deployment, eg:
     * 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}'. For example,
     * 'projects/my-project/locations/{location}/gameServerDeployments/my-deployment'.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GameServerDeployment resource.
 */
export interface GameServerDeploymentArgs {
    /**
     * A unique id for the deployment.
     */
    readonly deploymentId: pulumi.Input<string>;
    /**
     * Human readable description of the game server deployment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The labels associated with this game server deployment. Each label is a
     * key-value pair.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the Deployment.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
