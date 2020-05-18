// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows management of a customized Cloud IAM project role. For more information see
 * [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
 * and
 * [API](https://cloud.google.com/iam/reference/rest/v1/projects.roles).
 *
 * > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
 *  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
 *  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
 *  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
 *  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
 *  by the provider, and new roles cannot share that name.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myCustomRole = new gcp.projects.IAMCustomRole("my-custom-role", {
 *     description: "A description",
 *     permissions: [
 *         "iam.roles.list",
 *         "iam.roles.create",
 *         "iam.roles.delete",
 *     ],
 *     roleId: "myCustomRole",
 *     title: "My Custom Role",
 * });
 * ```
 */
export class IAMCustomRole extends pulumi.CustomResource {
    /**
     * Get an existing IAMCustomRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAMCustomRoleState, opts?: pulumi.CustomResourceOptions): IAMCustomRole {
        return new IAMCustomRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/iAMCustomRole:IAMCustomRole';

    /**
     * Returns true if the given object is an instance of IAMCustomRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IAMCustomRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IAMCustomRole.__pulumiType;
    }

    /**
     * (Optional) The current deleted state of the role.
     */
    public /*out*/ readonly deleted!: pulumi.Output<boolean>;
    /**
     * A human-readable description for the role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
     */
    public readonly permissions!: pulumi.Output<string[]>;
    /**
     * The project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The camel case role id to use for this role. Cannot contain `-` characters.
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * The current launch stage of the role.
     * Defaults to `GA`.
     * List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
     */
    public readonly stage!: pulumi.Output<string | undefined>;
    /**
     * A human-readable title for the role.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a IAMCustomRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAMCustomRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAMCustomRoleArgs | IAMCustomRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAMCustomRoleState | undefined;
            inputs["deleted"] = state ? state.deleted : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["roleId"] = state ? state.roleId : undefined;
            inputs["stage"] = state ? state.stage : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as IAMCustomRoleArgs | undefined;
            if (!args || args.permissions === undefined) {
                throw new Error("Missing required property 'permissions'");
            }
            if (!args || args.roleId === undefined) {
                throw new Error("Missing required property 'roleId'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["roleId"] = args ? args.roleId : undefined;
            inputs["stage"] = args ? args.stage : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["deleted"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IAMCustomRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IAMCustomRole resources.
 */
export interface IAMCustomRoleState {
    /**
     * (Optional) The current deleted state of the role.
     */
    readonly deleted?: pulumi.Input<boolean>;
    /**
     * A human-readable description for the role.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
     */
    readonly permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The camel case role id to use for this role. Cannot contain `-` characters.
     */
    readonly roleId?: pulumi.Input<string>;
    /**
     * The current launch stage of the role.
     * Defaults to `GA`.
     * List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
     */
    readonly stage?: pulumi.Input<string>;
    /**
     * A human-readable title for the role.
     */
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IAMCustomRole resource.
 */
export interface IAMCustomRoleArgs {
    /**
     * A human-readable description for the role.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
     */
    readonly permissions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The project that the service account will be created in.
     * Defaults to the provider project configuration.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The camel case role id to use for this role. Cannot contain `-` characters.
     */
    readonly roleId: pulumi.Input<string>;
    /**
     * The current launch stage of the role.
     * Defaults to `GA`.
     * List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
     */
    readonly stage?: pulumi.Input<string>;
    /**
     * A human-readable title for the role.
     */
    readonly title: pulumi.Input<string>;
}
