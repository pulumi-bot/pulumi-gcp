// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A game server config resource. Configs are global and immutable.
 *
 * To get more information about GameServerConfig, see:
 *
 * * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/game-servers/docs)
 *
 * ## Example Usage - Game Service Config Basic
 * {{% example %}}
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultGameServerDeployment = new gcp.gameservices.GameServerDeployment("defaultGameServerDeployment", {
 *     deploymentId: "tf-test-deployment",
 *     description: "a deployment description",
 * });
 * const defaultGameServerConfig = new gcp.gameservices.GameServerConfig("defaultGameServerConfig", {
 *     configId: "tf-test-config",
 *     deploymentId: defaultGameServerDeployment.deploymentId,
 *     description: "a config description",
 *     fleet_configs: [{
 *         name: "something-unique",
 *         fleetSpec: JSON.stringify({
 *             replicas: 1,
 *             scheduling: "Packed",
 *             template: {
 *                 metadata: {
 *                     name: "tf-test-game-server-template",
 *                 },
 *                 spec: {
 *                     template: {
 *                         spec: {
 *                             containers: [{
 *                                 name: "simple-udp-server",
 *                                 image: "gcr.io/agones-images/udp-server:0.14",
 *                             }],
 *                         },
 *                     },
 *                 },
 *             },
 *         }),
 *     }],
 *     scaling_configs: [{
 *         name: "scaling-config-name",
 *         fleetAutoscalerSpec: JSON.stringify({
 *             policy: {
 *                 type: "Webhook",
 *                 webhook: {
 *                     service: {
 *                         name: "autoscaler-webhook-service",
 *                         namespace: "default",
 *                         path: "scale",
 *                     },
 *                 },
 *             },
 *         }),
 *         selectors: [{
 *             labels: {
 *                 one: "two",
 *             },
 *         }],
 *         schedules: [{
 *             cronJobDuration: "3.500s",
 *             cronSpec: "0 0 * * 0",
 *         }],
 *     }],
 * });
 * ```
 *
 * {{% /example %}}
 */
export class GameServerConfig extends pulumi.CustomResource {
    /**
     * Get an existing GameServerConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GameServerConfigState, opts?: pulumi.CustomResourceOptions): GameServerConfig {
        return new GameServerConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gameservices/gameServerConfig:GameServerConfig';

    /**
     * Returns true if the given object is an instance of GameServerConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GameServerConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GameServerConfig.__pulumiType;
    }

    /**
     * A unique id for the deployment config.
     */
    public readonly configId!: pulumi.Output<string>;
    /**
     * A unique id for the deployment.
     */
    public readonly deploymentId!: pulumi.Output<string>;
    /**
     * The description of the game server config.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The fleet config contains list of fleet specs. In the Single Cloud, there
     * will be only one.  Structure is documented below.
     */
    public readonly fleetConfigs!: pulumi.Output<outputs.gameservices.GameServerConfigFleetConfig[]>;
    /**
     * Set of labels to group by.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Location of the Deployment.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the ScalingConfig
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. This contains the autoscaling settings.  Structure is documented below.
     */
    public readonly scalingConfigs!: pulumi.Output<outputs.gameservices.GameServerConfigScalingConfig[] | undefined>;

    /**
     * Create a GameServerConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GameServerConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GameServerConfigArgs | GameServerConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GameServerConfigState | undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["deploymentId"] = state ? state.deploymentId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fleetConfigs"] = state ? state.fleetConfigs : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["scalingConfigs"] = state ? state.scalingConfigs : undefined;
        } else {
            const args = argsOrState as GameServerConfigArgs | undefined;
            if (!args || args.configId === undefined) {
                throw new Error("Missing required property 'configId'");
            }
            if (!args || args.deploymentId === undefined) {
                throw new Error("Missing required property 'deploymentId'");
            }
            if (!args || args.fleetConfigs === undefined) {
                throw new Error("Missing required property 'fleetConfigs'");
            }
            inputs["configId"] = args ? args.configId : undefined;
            inputs["deploymentId"] = args ? args.deploymentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fleetConfigs"] = args ? args.fleetConfigs : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["scalingConfigs"] = args ? args.scalingConfigs : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GameServerConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GameServerConfig resources.
 */
export interface GameServerConfigState {
    /**
     * A unique id for the deployment config.
     */
    readonly configId?: pulumi.Input<string>;
    /**
     * A unique id for the deployment.
     */
    readonly deploymentId?: pulumi.Input<string>;
    /**
     * The description of the game server config.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fleet config contains list of fleet specs. In the Single Cloud, there
     * will be only one.  Structure is documented below.
     */
    readonly fleetConfigs?: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerConfigFleetConfig>[]>;
    /**
     * Set of labels to group by.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Location of the Deployment.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the ScalingConfig
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Optional. This contains the autoscaling settings.  Structure is documented below.
     */
    readonly scalingConfigs?: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerConfigScalingConfig>[]>;
}

/**
 * The set of arguments for constructing a GameServerConfig resource.
 */
export interface GameServerConfigArgs {
    /**
     * A unique id for the deployment config.
     */
    readonly configId: pulumi.Input<string>;
    /**
     * A unique id for the deployment.
     */
    readonly deploymentId: pulumi.Input<string>;
    /**
     * The description of the game server config.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fleet config contains list of fleet specs. In the Single Cloud, there
     * will be only one.  Structure is documented below.
     */
    readonly fleetConfigs: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerConfigFleetConfig>[]>;
    /**
     * Set of labels to group by.
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
    /**
     * Optional. This contains the autoscaling settings.  Structure is documented below.
     */
    readonly scalingConfigs?: pulumi.Input<pulumi.Input<inputs.gameservices.GameServerConfigScalingConfig>[]>;
}
