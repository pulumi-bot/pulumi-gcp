// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of a single API service for an existing Google Cloud Platform project. 
 * 
 * For a list of services available, visit the
 * [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
 * 
 * > **Note:** This resource _must not_ be used in conjunction with
 *    `google_project_services` or they will fight over which services should be enabled.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const project = new gcp.projects.Service("project", {
 *     disableDependentServices: true,
 *     project: "your-project-id",
 *     service: "iam.googleapis.com",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_service.html.markdown.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
     * If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
     */
    public readonly disableDependentServices!: pulumi.Output<boolean | undefined>;
    public readonly disableOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The project ID. If not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The service to enable.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["disableDependentServices"] = state ? state.disableDependentServices : undefined;
            inputs["disableOnDestroy"] = state ? state.disableOnDestroy : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["disableDependentServices"] = args ? args.disableDependentServices : undefined;
            inputs["disableOnDestroy"] = args ? args.disableOnDestroy : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
     * If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
     */
    readonly disableDependentServices?: pulumi.Input<boolean>;
    readonly disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID. If not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The service to enable.
     */
    readonly service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
     * If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
     */
    readonly disableDependentServices?: pulumi.Input<boolean>;
    readonly disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID. If not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The service to enable.
     */
    readonly service: pulumi.Input<string>;
}
