// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of enabled API services for an existing Google Cloud
 * Platform project. Services in an existing project that are not defined
 * in the config will be removed.
 * 
 * For a list of services available, visit the
 * [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
 * 
 * > **Note:** This resource attempts to be the authoritative source on *all* enabled APIs, which often
 * 	leads to conflicts when certain actions enable other APIs. If you do not need to ensure that
 * 	*exclusively* a particular set of APIs are enabled, you should most likely use the
 * 	google_project_service resource, one resource per API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const project = new gcp.projects.Services("project", {
 *     project: "your-project-id",
 *     services: [
 *         "iam.googleapis.com",
 *         "cloudresourcemanager.googleapis.com",
 *     ],
 * });
 * ```
 */
export class Services extends pulumi.CustomResource {
    /**
     * Get an existing Services resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicesState, opts?: pulumi.CustomResourceOptions): Services {
        return new Services(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'gcp:projects/services:Services';

    /**
     * Returns true if the given object is an instance of Services.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Services {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === Services.__pulumiType;
    }

    public readonly disableOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The project ID.
     * Changing this forces Terraform to attempt to disable all previously managed
     * API services in the previous project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The list of services that are enabled. Supports
     * update.
     */
    public readonly services!: pulumi.Output<string[]>;

    /**
     * Create a Services resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicesArgs | ServicesState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServicesState | undefined;
            inputs["disableOnDestroy"] = state ? state.disableOnDestroy : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["services"] = state ? state.services : undefined;
        } else {
            const args = argsOrState as ServicesArgs | undefined;
            if (!args || args.services === undefined) {
                throw new Error("Missing required property 'services'");
            }
            inputs["disableOnDestroy"] = args ? args.disableOnDestroy : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["services"] = args ? args.services : undefined;
        }
        super(Services.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Services resources.
 */
export interface ServicesState {
    readonly disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID.
     * Changing this forces Terraform to attempt to disable all previously managed
     * API services in the previous project.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The list of services that are enabled. Supports
     * update.
     */
    readonly services?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Services resource.
 */
export interface ServicesArgs {
    readonly disableOnDestroy?: pulumi.Input<boolean>;
    /**
     * The project ID.
     * Changing this forces Terraform to attempt to disable all previously managed
     * API services in the previous project.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The list of services that are enabled. Supports
     * update.
     */
    readonly services: pulumi.Input<pulumi.Input<string>[]>;
}
