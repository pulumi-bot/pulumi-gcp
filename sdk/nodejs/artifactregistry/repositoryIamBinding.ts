// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Artifact Registry Repository. Each of these resources serves a different use case:
 *
 * * `gcp.artifactregistry.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
 * * `gcp.artifactregistry.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
 * * `gcp.artifactregistry.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
 *
 * > **Note:** `gcp.artifactregistry.RepositoryIamPolicy` **cannot** be used in conjunction with `gcp.artifactregistry.RepositoryIamBinding` and `gcp.artifactregistry.RepositoryIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.artifactregistry.RepositoryIamBinding` resources **can be** used in conjunction with `gcp.artifactregistry.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_artifact\_registry\_repository\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.artifactregistry.RepositoryIamPolicy("policy", {
 *     project: google_artifact_registry_repository["my-repo"].project,
 *     location: google_artifact_registry_repository["my-repo"].location,
 *     repository: google_artifact_registry_repository["my-repo"].name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_artifact\_registry\_repository\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.artifactregistry.RepositoryIamBinding("binding", {
 *     project: google_artifact_registry_repository["my-repo"].project,
 *     location: google_artifact_registry_repository["my-repo"].location,
 *     repository: google_artifact_registry_repository["my-repo"].name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_artifact\_registry\_repository\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.artifactregistry.RepositoryIamMember("member", {
 *     project: google_artifact_registry_repository["my-repo"].project,
 *     location: google_artifact_registry_repository["my-repo"].location,
 *     repository: google_artifact_registry_repository["my-repo"].name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/repositories/{{repository}} * {{project}}/{{location}}/{{repository}} * {{location}}/{{repository}} * {{repository}} Any variables not passed in the import command will be taken from the provider configuration. Artifact Registry repository IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor "projects/{{project}}/locations/{{location}}/repositories/{{repository}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding editor projects/{{project}}/locations/{{location}}/repositories/{{repository}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class RepositoryIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryIamBindingState, opts?: pulumi.CustomResourceOptions): RepositoryIamBinding {
        return new RepositoryIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:artifactregistry/repositoryIamBinding:RepositoryIamBinding';

    /**
     * Returns true if the given object is an instance of RepositoryIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.artifactregistry.RepositoryIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the location this repository is located in.
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly location!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a RepositoryIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryIamBindingArgs | RepositoryIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryIamBindingState | undefined;
            inputs["condition"] = state ? state.condition : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as RepositoryIamBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["condition"] = args ? args.condition : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RepositoryIamBinding.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryIamBinding resources.
 */
export interface RepositoryIamBindingState {
    readonly condition?: pulumi.Input<inputs.artifactregistry.RepositoryIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The name of the location this repository is located in.
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly location?: pulumi.Input<string>;
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryIamBinding resource.
 */
export interface RepositoryIamBindingArgs {
    readonly condition?: pulumi.Input<inputs.artifactregistry.RepositoryIamBindingCondition>;
    /**
     * The name of the location this repository is located in.
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly location?: pulumi.Input<string>;
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    readonly repository: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    readonly role: pulumi.Input<string>;
}
