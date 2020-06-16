// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ServicePerimeter describes a set of GCP resources which can freely import
 * and export data amongst themselves, but not export outside of the
 * ServicePerimeter. If a request with a source within this ServicePerimeter
 * has a target outside of the ServicePerimeter, the request will be blocked.
 * Otherwise the request is allowed. There are two types of Service Perimeter
 * - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
 *   GCP project can only belong to a single regular Service Perimeter. Service
 *   Perimeter Bridges can contain only GCP projects as members, a single GCP
 *   project may belong to multiple Service Perimeter Bridges.
 *
 * To get more information about ServicePerimeter, see:
 *
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
 * * How-to Guides
 *     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * ### Access Context Manager Service Perimeter Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const access_policy = new gcp.accesscontextmanager.AccessPolicy("access-policy", {
 *     parent: "organizations/123456789",
 *     title: "my policy",
 * });
 * const service_perimeter = new gcp.accesscontextmanager.ServicePerimeter("service-perimeter", {
 *     parent: pulumi.interpolate`accessPolicies/${access_policy.name}`,
 *     status: {
 *         restrictedServices: ["storage.googleapis.com"],
 *     },
 *     title: "restrict_storage",
 * });
 * const access_level = new gcp.accesscontextmanager.AccessLevel("access-level", {
 *     basic: {
 *         conditions: [{
 *             devicePolicy: {
 *                 osConstraints: [{
 *                     osType: "DESKTOP_CHROME_OS",
 *                 }],
 *                 requireScreenLock: false,
 *             },
 *             regions: [
 *                 "CH",
 *                 "IT",
 *                 "US",
 *             ],
 *         }],
 *     },
 *     parent: pulumi.interpolate`accessPolicies/${access_policy.name}`,
 *     title: "chromeos_no_lock",
 * });
 * ```
 * {{% /example %}}
 * {{% example %}}
 * ### Access Context Manager Service Perimeter Dry Run
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const access_policy = new gcp.accesscontextmanager.AccessPolicy("access-policy", {
 *     parent: "organizations/123456789",
 *     title: "my policy",
 * });
 * const service_perimeter = new gcp.accesscontextmanager.ServicePerimeter("service-perimeter", {
 *     parent: pulumi.interpolate`accessPolicies/${access_policy.name}`,
 *     // Service 'storage.googleapis.com' will be in dry-run mode.
 *     spec: {
 *         restrictedServices: ["storage.googleapis.com"],
 *     },
 *     // Service 'bigquery.googleapis.com' will be restricted.
 *     status: {
 *         restrictedServices: ["bigquery.googleapis.com"],
 *     },
 *     title: "restrict_bigquery_dryrun_storage",
 *     useExplicitDryRunSpec: true,
 * });
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export class ServicePerimeter extends pulumi.CustomResource {
    /**
     * Get an existing ServicePerimeter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePerimeterState, opts?: pulumi.CustomResourceOptions): ServicePerimeter {
        return new ServicePerimeter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:accesscontextmanager/servicePerimeter:ServicePerimeter';

    /**
     * Returns true if the given object is an instance of ServicePerimeter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePerimeter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePerimeter.__pulumiType;
    }

    /**
     * Time the AccessPolicy was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the ServicePerimeter and its use. Does not affect
     * behavior.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Resource name for the ServicePerimeter. The shortName component must
     * begin with a letter and only include alphanumeric and '_'.
     * Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Specifies the type of the Perimeter. There are two types: regular and
     * bridge. Regular Service Perimeter contains resources, access levels,
     * and restricted services. Every resource can be in at most
     * ONE regular Service Perimeter.
     * In addition to being in a regular service perimeter, a resource can also
     * be in zero or more perimeter bridges. A perimeter bridge only contains
     * resources. Cross project operations are permitted if all effected
     * resources share some perimeter (whether bridge or regular). Perimeter
     * Bridge does not contain access levels or services: those are governed
     * entirely by the regular perimeter that resource is in.
     * Perimeter Bridges are typically useful when building more complex
     * topologies with many independent perimeters that need to share some data
     * with a common perimeter, but should not be able to share data among
     * themselves.
     */
    public readonly perimeterType!: pulumi.Output<string | undefined>;
    /**
     * Proposed (or dry run) ServicePerimeter configuration.
     * This configuration allows to specify and test ServicePerimeter configuration
     * without enforcing actual access restrictions. Only allowed to be set when
     * the `useExplicitDryRunSpec` flag is set.  Structure is documented below.
     */
    public readonly spec!: pulumi.Output<outputs.accesscontextmanager.ServicePerimeterSpec | undefined>;
    /**
     * ServicePerimeter configuration. Specifies sets of resources,
     * restricted services and access levels that determine
     * perimeter content and boundaries.  Structure is documented below.
     */
    public readonly status!: pulumi.Output<outputs.accesscontextmanager.ServicePerimeterStatus | undefined>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * Time the AccessPolicy was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
     * for all Service Perimeters, and that spec is identical to the status for those
     * Service Perimeters. When this flag is set, it inhibits the generation of the
     * implicit spec, thereby allowing the user to explicitly provide a
     * configuration ("spec") to use in a dry-run version of the Service Perimeter.
     * This allows the user to test changes to the enforced config ("status") without
     * actually enforcing them. This testing is done through analyzing the differences
     * between currently enforced and suggested restrictions. useExplicitDryRunSpec must
     * bet set to True if any of the fields in the spec are set to non-default values.
     */
    public readonly useExplicitDryRunSpec!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ServicePerimeter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePerimeterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePerimeterArgs | ServicePerimeterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServicePerimeterState | undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parent"] = state ? state.parent : undefined;
            inputs["perimeterType"] = state ? state.perimeterType : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["updateTime"] = state ? state.updateTime : undefined;
            inputs["useExplicitDryRunSpec"] = state ? state.useExplicitDryRunSpec : undefined;
        } else {
            const args = argsOrState as ServicePerimeterArgs | undefined;
            if (!args || args.parent === undefined) {
                throw new Error("Missing required property 'parent'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parent"] = args ? args.parent : undefined;
            inputs["perimeterType"] = args ? args.perimeterType : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["useExplicitDryRunSpec"] = args ? args.useExplicitDryRunSpec : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServicePerimeter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePerimeter resources.
 */
export interface ServicePerimeterState {
    /**
     * Time the AccessPolicy was created in UTC.
     */
    readonly createTime?: pulumi.Input<string>;
    /**
     * Description of the ServicePerimeter and its use. Does not affect
     * behavior.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource name for the ServicePerimeter. The shortName component must
     * begin with a letter and only include alphanumeric and '_'.
     * Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    readonly parent?: pulumi.Input<string>;
    /**
     * Specifies the type of the Perimeter. There are two types: regular and
     * bridge. Regular Service Perimeter contains resources, access levels,
     * and restricted services. Every resource can be in at most
     * ONE regular Service Perimeter.
     * In addition to being in a regular service perimeter, a resource can also
     * be in zero or more perimeter bridges. A perimeter bridge only contains
     * resources. Cross project operations are permitted if all effected
     * resources share some perimeter (whether bridge or regular). Perimeter
     * Bridge does not contain access levels or services: those are governed
     * entirely by the regular perimeter that resource is in.
     * Perimeter Bridges are typically useful when building more complex
     * topologies with many independent perimeters that need to share some data
     * with a common perimeter, but should not be able to share data among
     * themselves.
     */
    readonly perimeterType?: pulumi.Input<string>;
    /**
     * Proposed (or dry run) ServicePerimeter configuration.
     * This configuration allows to specify and test ServicePerimeter configuration
     * without enforcing actual access restrictions. Only allowed to be set when
     * the `useExplicitDryRunSpec` flag is set.  Structure is documented below.
     */
    readonly spec?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterSpec>;
    /**
     * ServicePerimeter configuration. Specifies sets of resources,
     * restricted services and access levels that determine
     * perimeter content and boundaries.  Structure is documented below.
     */
    readonly status?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterStatus>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    readonly title?: pulumi.Input<string>;
    /**
     * Time the AccessPolicy was updated in UTC.
     */
    readonly updateTime?: pulumi.Input<string>;
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
     * for all Service Perimeters, and that spec is identical to the status for those
     * Service Perimeters. When this flag is set, it inhibits the generation of the
     * implicit spec, thereby allowing the user to explicitly provide a
     * configuration ("spec") to use in a dry-run version of the Service Perimeter.
     * This allows the user to test changes to the enforced config ("status") without
     * actually enforcing them. This testing is done through analyzing the differences
     * between currently enforced and suggested restrictions. useExplicitDryRunSpec must
     * bet set to True if any of the fields in the spec are set to non-default values.
     */
    readonly useExplicitDryRunSpec?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ServicePerimeter resource.
 */
export interface ServicePerimeterArgs {
    /**
     * Description of the ServicePerimeter and its use. Does not affect
     * behavior.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Resource name for the ServicePerimeter. The shortName component must
     * begin with a letter and only include alphanumeric and '_'.
     * Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     */
    readonly parent: pulumi.Input<string>;
    /**
     * Specifies the type of the Perimeter. There are two types: regular and
     * bridge. Regular Service Perimeter contains resources, access levels,
     * and restricted services. Every resource can be in at most
     * ONE regular Service Perimeter.
     * In addition to being in a regular service perimeter, a resource can also
     * be in zero or more perimeter bridges. A perimeter bridge only contains
     * resources. Cross project operations are permitted if all effected
     * resources share some perimeter (whether bridge or regular). Perimeter
     * Bridge does not contain access levels or services: those are governed
     * entirely by the regular perimeter that resource is in.
     * Perimeter Bridges are typically useful when building more complex
     * topologies with many independent perimeters that need to share some data
     * with a common perimeter, but should not be able to share data among
     * themselves.
     */
    readonly perimeterType?: pulumi.Input<string>;
    /**
     * Proposed (or dry run) ServicePerimeter configuration.
     * This configuration allows to specify and test ServicePerimeter configuration
     * without enforcing actual access restrictions. Only allowed to be set when
     * the `useExplicitDryRunSpec` flag is set.  Structure is documented below.
     */
    readonly spec?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterSpec>;
    /**
     * ServicePerimeter configuration. Specifies sets of resources,
     * restricted services and access levels that determine
     * perimeter content and boundaries.  Structure is documented below.
     */
    readonly status?: pulumi.Input<inputs.accesscontextmanager.ServicePerimeterStatus>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    readonly title: pulumi.Input<string>;
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
     * for all Service Perimeters, and that spec is identical to the status for those
     * Service Perimeters. When this flag is set, it inhibits the generation of the
     * implicit spec, thereby allowing the user to explicitly provide a
     * configuration ("spec") to use in a dry-run version of the Service Perimeter.
     * This allows the user to test changes to the enforced config ("status") without
     * actually enforcing them. This testing is done through analyzing the differences
     * between currently enforced and suggested restrictions. useExplicitDryRunSpec must
     * bet set to True if any of the fields in the spec are set to non-default values.
     */
    readonly useExplicitDryRunSpec?: pulumi.Input<boolean>;
}
