// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A container for `services`. Namespaces allow administrators to group services
 * together and define permissions for a collection of services.
 *
 * To get more information about Namespace, see:
 *
 * * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces)
 * * How-to Guides
 *     * [Configuring a namespace](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_namespace)
 *
 * ## Example Usage - Service Directory Namespace Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.servicedirectory.Namespace("example", {
 *     namespaceId: "example-namespace",
 *     location: "us-central1",
 *     labels: {
 *         key: "value",
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_directory_namespace.html.markdown.
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:servicedirectory/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * Resource labels associated with this Namespace. No more than 64 user
     * labels can be associated with a given resource. Label keys and values can
     * be no longer than 63 characters.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the Namespace.
     * A full list of valid locations can be found by running
     * `gcloud beta service-directory locations list`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name for the namespace in the format 'projects/*&#47;locations/*&#47;namespaces/*'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.namespaceId === undefined) {
                throw new Error("Missing required property 'namespaceId'");
            }
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * Resource labels associated with this Namespace. No more than 64 user
     * labels can be associated with a given resource. Label keys and values can
     * be no longer than 63 characters.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the Namespace.
     * A full list of valid locations can be found by running
     * `gcloud beta service-directory locations list`.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The resource name for the namespace in the format 'projects/*&#47;locations/*&#47;namespaces/*'.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    readonly namespaceId?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * Resource labels associated with this Namespace. No more than 64 user
     * labels can be associated with a given resource. Label keys and values can
     * be no longer than 63 characters.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the Namespace.
     * A full list of valid locations can be found by running
     * `gcloud beta service-directory locations list`.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The Resource ID must be 1-63 characters long, including digits,
     * lowercase letters or the hyphen character.
     */
    readonly namespaceId: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
