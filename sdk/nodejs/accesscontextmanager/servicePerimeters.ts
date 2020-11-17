// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Replace all existing Service Perimeters in an Access Policy with the Service Perimeters provided. This is done atomically.
 * This is a bulk edit of all Service Perimeters and may override existing Service Perimeters created by `gcp.accesscontextmanager.ServicePerimeter`,
 * thus causing a permadiff if used alongside `gcp.accesscontextmanager.ServicePerimeter` on the same parent.
 *
 * To get more information about ServicePerimeters, see:
 *
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
 * * How-to Guides
 *     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
 *
 * ## Example Usage
 *
 * ## Import
 *
 * ServicePerimeters can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}/servicePerimeters
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}
 * ```
 */
export class ServicePerimeters extends pulumi.CustomResource {
    /**
     * Get an existing ServicePerimeters resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePerimetersState, opts?: pulumi.CustomResourceOptions): ServicePerimeters {
        return new ServicePerimeters(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/servicePerimeters:ServicePerimeters';

    /**
     * Returns true if the given object is an instance of ServicePerimeters.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePerimeters {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePerimeters.__pulumiType;
    }

    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
     * Structure is documented below.
     */
    public readonly servicePerimeters!: pulumi.Output<outputs.accesscontextmanager.ServicePerimetersServicePerimeter[] | undefined>;

    /**
     * Create a ServicePerimeters resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePerimetersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePerimetersArgs | ServicePerimetersState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServicePerimetersState | undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["servicePerimeters"] = state ? state.servicePerimeters : undefined;
        } else {
            const args = argsOrState as ServicePerimetersArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            inputs["parent"] = args ? args.parent : undefined;
            inputs["servicePerimeters"] = args ? args.servicePerimeters : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServicePerimeters.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePerimeters resources.
 */
export interface ServicePerimetersState {
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
     * Structure is documented below.
     */
    readonly servicePerimeters?: pulumi.Input<pulumi.Input<inputs.accesscontextmanager.ServicePerimetersServicePerimeter>[]>;
}

/**
 * The set of arguments for constructing a ServicePerimeters resource.
 */
export interface ServicePerimetersArgs {
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    readonly parent: pulumi.Input<string>;
    /**
     * The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
     * Structure is documented below.
     */
    readonly servicePerimeters?: pulumi.Input<pulumi.Input<inputs.accesscontextmanager.ServicePerimetersServicePerimeter>[]>;
}
