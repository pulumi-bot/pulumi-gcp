// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Access Approval enables you to require your explicit approval whenever Google support and engineering need to access your customer content.
 *
 * To get more information about OrganizationSettings, see:
 *
 * * [API documentation](https://cloud.google.com/access-approval/docs/reference/rest/v1/organizations)
 *
 * ## Example Usage
 * ### Organization Access Approval Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const organizationAccessApproval = new gcp.organizations.AccessApprovalSettings("organization_access_approval", {
 *     enrolledServices: [
 *         {
 *             cloudProduct: "appengine.googleapis.com",
 *         },
 *         {
 *             cloudProduct: "dataflow.googleapis.com",
 *             enrollmentLevel: "BLOCK_ALL",
 *         },
 *     ],
 *     notificationEmails: [
 *         "testuser@example.com",
 *         "example.user@example.com",
 *     ],
 *     organizationId: "123456789",
 * });
 * ```
 *
 * ## Import
 *
 * OrganizationSettings can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:organizations/accessApprovalSettings:AccessApprovalSettings default organizations/{{organization_id}}/accessApprovalSettings
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:organizations/accessApprovalSettings:AccessApprovalSettings default {{organization_id}}
 * ```
 */
export class AccessApprovalSettings extends pulumi.CustomResource {
    /**
     * Get an existing AccessApprovalSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessApprovalSettingsState, opts?: pulumi.CustomResourceOptions): AccessApprovalSettings {
        return new AccessApprovalSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/accessApprovalSettings:AccessApprovalSettings';

    /**
     * Returns true if the given object is an instance of AccessApprovalSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessApprovalSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessApprovalSettings.__pulumiType;
    }

    /**
     * This field will always be unset for the organization since organizations do not have ancestors.
     */
    public /*out*/ readonly enrolledAncestor!: pulumi.Output<boolean>;
    /**
     * A list of Google Cloud Services for which the given resource has Access Approval enrolled.
     * Access requests for the resource given by name against any of these services contained here will be required
     * to have explicit approval. Enrollment can be done for individual services.
     * A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
     * Structure is documented below.
     */
    public readonly enrolledServices!: pulumi.Output<outputs.organizations.AccessApprovalSettingsEnrolledService[]>;
    /**
     * The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A list of email addresses to which notifications relating to approval requests should be sent.
     * Notifications relating to a resource will be sent to all emails in the settings of ancestor
     * resources of that resource. A maximum of 50 email addresses are allowed.
     */
    public readonly notificationEmails!: pulumi.Output<string[]>;
    /**
     * ID of the organization of the access approval settings.
     */
    public readonly organizationId!: pulumi.Output<string>;

    /**
     * Create a AccessApprovalSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessApprovalSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessApprovalSettingsArgs | AccessApprovalSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessApprovalSettingsState | undefined;
            inputs["enrolledAncestor"] = state ? state.enrolledAncestor : undefined;
            inputs["enrolledServices"] = state ? state.enrolledServices : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationEmails"] = state ? state.notificationEmails : undefined;
            inputs["organizationId"] = state ? state.organizationId : undefined;
        } else {
            const args = argsOrState as AccessApprovalSettingsArgs | undefined;
            if ((!args || args.enrolledServices === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enrolledServices'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            inputs["enrolledServices"] = args ? args.enrolledServices : undefined;
            inputs["notificationEmails"] = args ? args.notificationEmails : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["enrolledAncestor"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AccessApprovalSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessApprovalSettings resources.
 */
export interface AccessApprovalSettingsState {
    /**
     * This field will always be unset for the organization since organizations do not have ancestors.
     */
    enrolledAncestor?: pulumi.Input<boolean>;
    /**
     * A list of Google Cloud Services for which the given resource has Access Approval enrolled.
     * Access requests for the resource given by name against any of these services contained here will be required
     * to have explicit approval. Enrollment can be done for individual services.
     * A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
     * Structure is documented below.
     */
    enrolledServices?: pulumi.Input<pulumi.Input<inputs.organizations.AccessApprovalSettingsEnrolledService>[]>;
    /**
     * The resource name of the settings. Format is "organizations/{organization_id}/accessApprovalSettings"
     */
    name?: pulumi.Input<string>;
    /**
     * A list of email addresses to which notifications relating to approval requests should be sent.
     * Notifications relating to a resource will be sent to all emails in the settings of ancestor
     * resources of that resource. A maximum of 50 email addresses are allowed.
     */
    notificationEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the organization of the access approval settings.
     */
    organizationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccessApprovalSettings resource.
 */
export interface AccessApprovalSettingsArgs {
    /**
     * A list of Google Cloud Services for which the given resource has Access Approval enrolled.
     * Access requests for the resource given by name against any of these services contained here will be required
     * to have explicit approval. Enrollment can be done for individual services.
     * A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
     * Structure is documented below.
     */
    enrolledServices: pulumi.Input<pulumi.Input<inputs.organizations.AccessApprovalSettingsEnrolledService>[]>;
    /**
     * A list of email addresses to which notifications relating to approval requests should be sent.
     * Notifications relating to a resource will be sent to all emails in the settings of ancestor
     * resources of that resource. A maximum of 50 email addresses are allowed.
     */
    notificationEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the organization of the access approval settings.
     */
    organizationId: pulumi.Input<string>;
}
