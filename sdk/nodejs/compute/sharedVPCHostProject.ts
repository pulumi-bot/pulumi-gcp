// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Enables the Google Compute Engine
 * [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
 * feature for a project, assigning it as a Shared VPC host project.
 *
 * For more information, see,
 * [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
 * where the Shared VPC feature is referred to by its former name "XPN".
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // A host project provides network resources to associated service projects.
 * const host = new gcp.compute.SharedVPCHostProject("host", {project: "host-project-id"});
 * // A service project gains access to network resources provided by its
 * // associated host project.
 * const service1 = new gcp.compute.SharedVPCServiceProject("service1", {
 *     hostProject: host.project,
 *     serviceProject: "service-project-id-1",
 * });
 * const service2 = new gcp.compute.SharedVPCServiceProject("service2", {
 *     hostProject: host.project,
 *     serviceProject: "service-project-id-2",
 * });
 * ```
 *
 * ## Import
 *
 * Google Compute Engine Shared VPC host project feature can be imported using the `project`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/sharedVPCHostProject:SharedVPCHostProject host host-project-id
 * ```
 */
export class SharedVPCHostProject extends pulumi.CustomResource {
    /**
     * Get an existing SharedVPCHostProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SharedVPCHostProjectState, opts?: pulumi.CustomResourceOptions): SharedVPCHostProject {
        return new SharedVPCHostProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/sharedVPCHostProject:SharedVPCHostProject';

    /**
     * Returns true if the given object is an instance of SharedVPCHostProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SharedVPCHostProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SharedVPCHostProject.__pulumiType;
    }

    /**
     * The ID of the project that will serve as a Shared VPC host project
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a SharedVPCHostProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SharedVPCHostProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SharedVPCHostProjectArgs | SharedVPCHostProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SharedVPCHostProjectState | undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as SharedVPCHostProjectArgs | undefined;
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            inputs["project"] = args ? args.project : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SharedVPCHostProject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SharedVPCHostProject resources.
 */
export interface SharedVPCHostProjectState {
    /**
     * The ID of the project that will serve as a Shared VPC host project
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SharedVPCHostProject resource.
 */
export interface SharedVPCHostProjectArgs {
    /**
     * The ID of the project that will serve as a Shared VPC host project
     */
    readonly project: pulumi.Input<string>;
}
