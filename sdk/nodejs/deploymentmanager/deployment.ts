// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A collection of resources that are deployed and managed together using
 * a configuration file
 *
 * > **Warning:** This resource is intended only to manage a Deployment resource,
 * and attempts to manage the Deployment's resources in the provider as well
 * will likely result in errors or unexpected behavior as the two tools
 * fight over ownership. We strongly discourage doing so unless you are an
 * experienced user of both tools.
 *
 * In addition, due to limitations of the API, the provider will treat
 * deployments in preview as recreate-only for any update operation other
 * than actually deploying an in-preview deployment (i.e. `preview=true` to
 * `preview=false`).
 *
 * ## Example Usage
 * ### Deployment Manager Deployment Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * from "fs";
 *
 * const deployment = new gcp.deploymentmanager.Deployment("deployment", {
 *     target: {
 *         config: {
 *             content: fs.readFileSync("path/to/config.yml"),
 *         },
 *     },
 *     labels: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentState, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:deploymentmanager/deployment:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     */
    public readonly createPolicy!: pulumi.Output<string | undefined>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     */
    public readonly deletePolicy!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier for deployment. Output only.
     */
    public /*out*/ readonly deploymentId!: pulumi.Output<string>;
    /**
     * Optional user-provided description of deployment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key-value pairs to apply to this labels.  Structure is documented below.
     */
    public readonly labels!: pulumi.Output<outputs.deploymentmanager.DeploymentLabel[] | undefined>;
    /**
     * Output only. URL of the manifest representing the last manifest that was successfully deployed.
     */
    public /*out*/ readonly manifest!: pulumi.Output<string>;
    /**
     * The name of the template to import, as declared in the YAML
     * configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If set to true, a deployment is created with "shell" resources
     * that are not actually instantiated. This allows you to preview a
     * deployment. It can be updated to false to actually deploy
     * with real resources.
     * ~>**NOTE**: Deployment Manager does not allow update
     * of a deployment in preview (unless updating to preview=false). Thus,
     * the provider will force-recreate deployments if either preview is updated
     * to true or if other fields are updated while preview is true.
     */
    public readonly preview!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. Server defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.  Structure is documented below.
     */
    public readonly target!: pulumi.Output<outputs.deploymentmanager.DeploymentTarget>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentArgs | DeploymentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeploymentState | undefined;
            inputs["createPolicy"] = state ? state.createPolicy : undefined;
            inputs["deletePolicy"] = state ? state.deletePolicy : undefined;
            inputs["deploymentId"] = state ? state.deploymentId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["manifest"] = state ? state.manifest : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["preview"] = state ? state.preview : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as DeploymentArgs | undefined;
            if (!args || args.target === undefined) {
                throw new Error("Missing required property 'target'");
            }
            inputs["createPolicy"] = args ? args.createPolicy : undefined;
            inputs["deletePolicy"] = args ? args.deletePolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["preview"] = args ? args.preview : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["target"] = args ? args.target : undefined;
            inputs["deploymentId"] = undefined /*out*/;
            inputs["manifest"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Deployment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Deployment resources.
 */
export interface DeploymentState {
    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     */
    readonly createPolicy?: pulumi.Input<string>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     */
    readonly deletePolicy?: pulumi.Input<string>;
    /**
     * Unique identifier for deployment. Output only.
     */
    readonly deploymentId?: pulumi.Input<string>;
    /**
     * Optional user-provided description of deployment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Key-value pairs to apply to this labels.  Structure is documented below.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.deploymentmanager.DeploymentLabel>[]>;
    /**
     * Output only. URL of the manifest representing the last manifest that was successfully deployed.
     */
    readonly manifest?: pulumi.Input<string>;
    /**
     * The name of the template to import, as declared in the YAML
     * configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If set to true, a deployment is created with "shell" resources
     * that are not actually instantiated. This allows you to preview a
     * deployment. It can be updated to false to actually deploy
     * with real resources.
     * ~>**NOTE**: Deployment Manager does not allow update
     * of a deployment in preview (unless updating to preview=false). Thus,
     * the provider will force-recreate deployments if either preview is updated
     * to true or if other fields are updated while preview is true.
     */
    readonly preview?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Output only. Server defined URL for the resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.  Structure is documented below.
     */
    readonly target?: pulumi.Input<inputs.deploymentmanager.DeploymentTarget>;
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     */
    readonly createPolicy?: pulumi.Input<string>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     */
    readonly deletePolicy?: pulumi.Input<string>;
    /**
     * Optional user-provided description of deployment.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Key-value pairs to apply to this labels.  Structure is documented below.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.deploymentmanager.DeploymentLabel>[]>;
    /**
     * The name of the template to import, as declared in the YAML
     * configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If set to true, a deployment is created with "shell" resources
     * that are not actually instantiated. This allows you to preview a
     * deployment. It can be updated to false to actually deploy
     * with real resources.
     * ~>**NOTE**: Deployment Manager does not allow update
     * of a deployment in preview (unless updating to preview=false). Thus,
     * the provider will force-recreate deployments if either preview is updated
     * to true or if other fields are updated while preview is true.
     */
    readonly preview?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.  Structure is documented below.
     */
    readonly target: pulumi.Input<inputs.deploymentmanager.DeploymentTarget>;
}
