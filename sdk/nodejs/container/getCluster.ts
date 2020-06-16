// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get info about a GKE cluster from its name and location.
 *
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myCluster = gcp.container.getCluster({
 *     name: "my-cluster",
 *     location: "us-east1-a",
 * });
 * export const clusterUsername = myCluster.then(myCluster => myCluster.masterAuths[0].username);
 * export const clusterPassword = myCluster.then(myCluster => myCluster.masterAuths[0].password);
 * export const endpoint = myCluster.then(myCluster => myCluster.endpoint);
 * export const instanceGroupUrls = myCluster.then(myCluster => myCluster.instanceGroupUrls);
 * export const nodeConfig = myCluster.then(myCluster => myCluster.nodeConfigs);
 * export const nodePools = myCluster.then(myCluster => myCluster.nodePools);
 * ```
 * {{% /example %}}
 * {{% /examples %}}
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:container/getCluster:getCluster", {
        "location": args.location,
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getCluster.
 */
export interface GetClusterArgs {
    /**
     * The location (zone or region) this cluster has been
     * created in. One of `location`, `region`, `zone`, or a provider-level `zone` must
     * be specified.
     */
    readonly location?: string;
    /**
     * The name of the cluster.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The region this cluster has been created in. Deprecated
     * in favour of `location`.
     */
    readonly region?: string;
    /**
     * The zone this cluster has been created in. Deprecated in
     * favour of `location`.
     */
    readonly zone?: string;
}

/**
 * A collection of values returned by getCluster.
 */
export interface GetClusterResult {
    readonly additionalZones: string[];
    readonly addonsConfigs: outputs.container.GetClusterAddonsConfig[];
    readonly authenticatorGroupsConfigs: outputs.container.GetClusterAuthenticatorGroupsConfig[];
    readonly clusterAutoscalings: outputs.container.GetClusterClusterAutoscaling[];
    readonly clusterIpv4Cidr: string;
    readonly clusterTelemetries: outputs.container.GetClusterClusterTelemetry[];
    readonly databaseEncryptions: outputs.container.GetClusterDatabaseEncryption[];
    readonly defaultMaxPodsPerNode: number;
    readonly description: string;
    readonly enableBinaryAuthorization: boolean;
    readonly enableIntranodeVisibility: boolean;
    readonly enableKubernetesAlpha: boolean;
    readonly enableLegacyAbac: boolean;
    readonly enableShieldedNodes: boolean;
    readonly enableTpu: boolean;
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly initialNodeCount: number;
    readonly instanceGroupUrls: string[];
    readonly ipAllocationPolicies: outputs.container.GetClusterIpAllocationPolicy[];
    readonly labelFingerprint: string;
    readonly location?: string;
    readonly loggingService: string;
    readonly maintenancePolicies: outputs.container.GetClusterMaintenancePolicy[];
    readonly masterAuthorizedNetworksConfigs: outputs.container.GetClusterMasterAuthorizedNetworksConfig[];
    readonly masterAuths: outputs.container.GetClusterMasterAuth[];
    readonly masterVersion: string;
    readonly minMasterVersion: string;
    readonly monitoringService: string;
    readonly name: string;
    readonly network: string;
    readonly networkPolicies: outputs.container.GetClusterNetworkPolicy[];
    readonly nodeConfigs: outputs.container.GetClusterNodeConfig[];
    readonly nodeLocations: string[];
    readonly nodePools: outputs.container.GetClusterNodePool[];
    readonly nodeVersion: string;
    readonly operation: string;
    readonly podSecurityPolicyConfigs: outputs.container.GetClusterPodSecurityPolicyConfig[];
    readonly privateClusterConfigs: outputs.container.GetClusterPrivateClusterConfig[];
    readonly project?: string;
    readonly region?: string;
    readonly releaseChannels: outputs.container.GetClusterReleaseChannel[];
    readonly removeDefaultNodePool: boolean;
    readonly resourceLabels: {[key: string]: string};
    readonly resourceUsageExportConfigs: outputs.container.GetClusterResourceUsageExportConfig[];
    readonly servicesIpv4Cidr: string;
    readonly subnetwork: string;
    readonly tpuIpv4CidrBlock: string;
    readonly verticalPodAutoscalings: outputs.container.GetClusterVerticalPodAutoscaling[];
    readonly workloadIdentityConfigs: outputs.container.GetClusterWorkloadIdentityConfig[];
    readonly zone?: string;
}
