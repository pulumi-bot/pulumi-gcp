// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Allows management of Organization policies for a Google Folder. For more information see
 * [the official
 * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
 * [API](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy).
 *
 * ## Example Usage
 *
 * To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const serialPortPolicy = new gcp.folder.OrganizationPolicy("serial_port_policy", {
 *     booleanPolicy: {
 *         enforced: true,
 *     },
 *     constraint: "compute.disableSerialPortAccess",
 *     folder: "folders/123456789",
 * });
 * ```
 *
 * To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const servicesPolicy = new gcp.folder.OrganizationPolicy("services_policy", {
 *     constraint: "serviceuser.services",
 *     folder: "folders/123456789",
 *     listPolicy: {
 *         allow: {
 *             all: true,
 *         },
 *     },
 * });
 * ```
 *
 * Or to deny some services, use the following instead:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const servicesPolicy = new gcp.folder.OrganizationPolicy("services_policy", {
 *     constraint: "serviceuser.services",
 *     folder: "folders/123456789",
 *     listPolicy: {
 *         deny: {
 *             values: ["cloudresourcemanager.googleapis.com"],
 *         },
 *         suggestedValue: "compute.googleapis.com",
 *     },
 * });
 * ```
 *
 * To restore the default folder organization policy, use the following instead:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const servicesPolicy = new gcp.folder.OrganizationPolicy("services_policy", {
 *     constraint: "serviceuser.services",
 *     folder: "folders/123456789",
 *     restorePolicy: {
 *         default: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Folder organization policies can be imported using any of the follow formats
 *
 * ```sh
 *  $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folders/folder-1234/constraints/serviceuser.services
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folder-1234/serviceuser.services
 * ```
 */
export class OrganizationPolicy extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationPolicyState, opts?: pulumi.CustomResourceOptions): OrganizationPolicy {
        return new OrganizationPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:folder/organizationPolicy:OrganizationPolicy';

    /**
     * Returns true if the given object is an instance of OrganizationPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationPolicy.__pulumiType;
    }

    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below.
     */
    public readonly booleanPolicy!: pulumi.Output<outputs.folder.OrganizationPolicyBooleanPolicy | undefined>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    public readonly constraint!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    public readonly folder!: pulumi.Output<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It
     * can also be used to allow or deny all values. Structure is documented below.
     */
    public readonly listPolicy!: pulumi.Output<outputs.folder.OrganizationPolicyListPolicy | undefined>;
    /**
     * A restore policy is a constraint to restore the default policy. Structure is documented below.
     */
    public readonly restorePolicy!: pulumi.Output<outputs.folder.OrganizationPolicyRestorePolicy | undefined>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Version of the Policy. Default version is 0.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a OrganizationPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationPolicyArgs | OrganizationPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrganizationPolicyState | undefined;
            inputs["booleanPolicy"] = state ? state.booleanPolicy : undefined;
            inputs["constraint"] = state ? state.constraint : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["listPolicy"] = state ? state.listPolicy : undefined;
            inputs["restorePolicy"] = state ? state.restorePolicy : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as OrganizationPolicyArgs | undefined;
            if ((!args || args.constraint === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'constraint'");
            }
            if ((!args || args.folder === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'folder'");
            }
            inputs["booleanPolicy"] = args ? args.booleanPolicy : undefined;
            inputs["constraint"] = args ? args.constraint : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["listPolicy"] = args ? args.listPolicy : undefined;
            inputs["restorePolicy"] = args ? args.restorePolicy : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrganizationPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationPolicy resources.
 */
export interface OrganizationPolicyState {
    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below.
     */
    readonly booleanPolicy?: pulumi.Input<inputs.folder.OrganizationPolicyBooleanPolicy>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It
     * can also be used to allow or deny all values. Structure is documented below.
     */
    readonly listPolicy?: pulumi.Input<inputs.folder.OrganizationPolicyListPolicy>;
    /**
     * A restore policy is a constraint to restore the default policy. Structure is documented below.
     */
    readonly restorePolicy?: pulumi.Input<inputs.folder.OrganizationPolicyRestorePolicy>;
    /**
     * (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * Version of the Policy. Default version is 0.
     */
    readonly version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrganizationPolicy resource.
 */
export interface OrganizationPolicyArgs {
    /**
     * A boolean policy is a constraint that is either enforced or not. Structure is documented below.
     */
    readonly booleanPolicy?: pulumi.Input<inputs.folder.OrganizationPolicyBooleanPolicy>;
    /**
     * The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint: pulumi.Input<string>;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    readonly folder: pulumi.Input<string>;
    /**
     * A policy that can define specific values that are allowed or denied for the given constraint. It
     * can also be used to allow or deny all values. Structure is documented below.
     */
    readonly listPolicy?: pulumi.Input<inputs.folder.OrganizationPolicyListPolicy>;
    /**
     * A restore policy is a constraint to restore the default policy. Structure is documented below.
     */
    readonly restorePolicy?: pulumi.Input<inputs.folder.OrganizationPolicyRestorePolicy>;
    /**
     * Version of the Policy. Default version is 0.
     */
    readonly version?: pulumi.Input<number>;
}
