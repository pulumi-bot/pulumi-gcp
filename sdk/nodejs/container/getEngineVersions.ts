// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides access to available Google Kubernetes Engine versions in a zone or region for a given project.
 * 
 * > If you are using the `google_container_engine_versions` datasource with a
 * regional cluster, ensure that you have provided a region as the `location` to
 * the datasource. A region can have a different set of supported versions than
 * its component zones, and not all zones in a region are guaranteed to
 * support the same version.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const central1b = pulumi.output(gcp.container.getEngineVersions({
 *     location: "us-central1-b",
 *     versionPrefix: "1.12.",
 * }));
 * const foo = new gcp.container.Cluster("foo", {
 *     initialNodeCount: 1,
 *     location: "us-central1-b",
 *     masterAuth: {
 *         password: "adoy.rm",
 *         username: "mr.yoda",
 *     },
 *     nodeVersion: central1b.latestNodeVersion,
 * });
 * ```
 */
export function getEngineVersions(args?: GetEngineVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetEngineVersionsResult> {
    args = args || {};
    return pulumi.runtime.invoke("gcp:container/getEngineVersions:getEngineVersions", {
        "location": args.location,
        "project": args.project,
        "region": args.region,
        "versionPrefix": args.versionPrefix,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getEngineVersions.
 */
export interface GetEngineVersionsArgs {
    /**
     * The location (region or zone) to list versions for.
     * Must exactly match the location the cluster will be deployed in, or listed
     * versions may not be available. If `location`, `region`, and `zone` are not
     * specified, the provider-level zone must be set and is used instead.
     */
    readonly location?: string;
    /**
     * ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
     * Defaults to the project that the provider is authenticated with.
     */
    readonly project?: string;
    readonly region?: string;
    /**
     * If provided, Terraform will only return versions
     * that match the string prefix. For example, `1.11.` will match all `1.11` series
     * releases. Since this is just a string match, it's recommended that you append a
     * `.` after minor versions to ensure that prefixes such as `1.1` don't match
     * versions like `1.12.5-gke.10` accidentally. See [the docs on versioning schema](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#versioning_scheme)
     * for full details on how version strings are formatted.
     */
    readonly versionPrefix?: string;
    readonly zone?: string;
}

/**
 * A collection of values returned by getEngineVersions.
 */
export interface GetEngineVersionsResult {
    /**
     * Version of Kubernetes the service deploys by default.
     */
    readonly defaultClusterVersion: string;
    /**
     * The latest version available in the given zone for use with master instances.
     */
    readonly latestMasterVersion: string;
    /**
     * The latest version available in the given zone for use with node instances.
     */
    readonly latestNodeVersion: string;
    readonly location?: string;
    readonly project?: string;
    readonly region?: string;
    /**
     * A list of versions available in the given zone for use with master instances.
     */
    readonly validMasterVersions: string[];
    /**
     * A list of versions available in the given zone for use with node instances.
     */
    readonly validNodeVersions: string[];
    readonly versionPrefix?: string;
    readonly zone?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
