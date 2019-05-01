// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get info about a GKE cluster from its name and location.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const myCluster = pulumi.output(gcp.container.getCluster({
 *     location: "us-east1-a",
 *     name: "my-cluster",
 * }));
 * 
 * export const clusterPassword = myCluster.apply(myCluster => myCluster.masterAuths[0].password);
 * export const clusterUsername = myCluster.apply(myCluster => myCluster.masterAuths[0].username);
 * export const endpoint = myCluster.apply(myCluster => myCluster.endpoint);
 * export const instanceGroupUrls = myCluster.apply(myCluster => myCluster.instanceGroupUrls);
 * export const nodeConfig = myCluster.apply(myCluster => myCluster.nodeConfigs);
 * export const nodePools = myCluster.apply(myCluster => myCluster.nodePools);
 * ```
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
    readonly addonsConfigs: { cloudrunConfigs: { disabled: boolean }[], horizontalPodAutoscalings: { disabled: boolean }[], httpLoadBalancings: { disabled: boolean }[], istioConfigs: { auth: string, disabled: boolean }[], kubernetesDashboards: { disabled: boolean }[], networkPolicyConfigs: { disabled: boolean }[] }[];
    readonly clusterAutoscalings: { enabled: boolean, resourceLimits: { maximum: number, minimum: number, resourceType: string }[] }[];
    readonly clusterIpv4Cidr: string;
    readonly defaultMaxPodsPerNode: number;
    readonly description: string;
    readonly enableBinaryAuthorization: boolean;
    readonly enableKubernetesAlpha: boolean;
    readonly enableLegacyAbac: boolean;
    readonly enableTpu: boolean;
    readonly endpoint: string;
    readonly initialNodeCount: number;
    readonly instanceGroupUrls: string[];
    readonly ipAllocationPolicies: { clusterIpv4CidrBlock: string, clusterSecondaryRangeName: string, createSubnetwork: boolean, nodeIpv4CidrBlock: string, servicesIpv4CidrBlock: string, servicesSecondaryRangeName: string, subnetworkName: string, useIpAliases: boolean }[];
    readonly location?: string;
    readonly loggingService: string;
    readonly maintenancePolicies: { dailyMaintenanceWindows: { duration: string, startTime: string }[] }[];
    readonly masterAuths: { clientCertificate: string, clientCertificateConfigs: { issueClientCertificate: boolean }[], clientKey: string, clusterCaCertificate: string, password: string, username: string }[];
    readonly masterAuthorizedNetworksConfigs: { cidrBlocks: { cidrBlock: string, displayName: string }[] }[];
    readonly masterVersion: string;
    readonly minMasterVersion: string;
    readonly monitoringService: string;
    readonly name: string;
    readonly network: string;
    readonly networkPolicies: { enabled: boolean, provider: string }[];
    readonly nodeConfigs: { diskSizeGb: number, diskType: string, guestAccelerators: { count: number, type: string }[], imageType: string, labels: {[key: string]: string}, localSsdCount: number, machineType: string, metadata: {[key: string]: string}, minCpuPlatform: string, oauthScopes: string[], preemptible: boolean, serviceAccount: string, tags: string[], taints: { effect: string, key: string, value: string }[], workloadMetadataConfigs: { nodeMetadata: string }[] }[];
    readonly nodeLocations: string[];
    readonly nodePools: { autoscalings: { maxNodeCount: number, minNodeCount: number }[], initialNodeCount: number, instanceGroupUrls: string[], managements: { autoRepair: boolean, autoUpgrade: boolean }[], maxPodsPerNode: number, name: string, namePrefix: string, nodeConfigs: { diskSizeGb: number, diskType: string, guestAccelerators: { count: number, type: string }[], imageType: string, labels: {[key: string]: string}, localSsdCount: number, machineType: string, metadata: {[key: string]: string}, minCpuPlatform: string, oauthScopes: string[], preemptible: boolean, serviceAccount: string, tags: string[], taints: { effect: string, key: string, value: string }[], workloadMetadataConfigs: { nodeMetadata: string }[] }[], nodeCount: number, version: string }[];
    readonly nodeVersion: string;
    readonly podSecurityPolicyConfigs: { enabled: boolean }[];
    readonly privateClusterConfigs: { enablePrivateEndpoint: boolean, enablePrivateNodes: boolean, masterIpv4CidrBlock: string, privateEndpoint: string, publicEndpoint: string }[];
    readonly project?: string;
    readonly region?: string;
    readonly removeDefaultNodePool: boolean;
    readonly resourceLabels: {[key: string]: string};
    readonly subnetwork: string;
    readonly tpuIpv4CidrBlock: string;
    readonly zone?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
