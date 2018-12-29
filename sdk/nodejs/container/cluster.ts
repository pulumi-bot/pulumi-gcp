// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    public readonly additionalZones: pulumi.Output<string[]>;
    public readonly addonsConfig: pulumi.Output<{ horizontalPodAutoscaling: { disabled?: boolean }, httpLoadBalancing: { disabled?: boolean }, kubernetesDashboard: { disabled?: boolean }, networkPolicyConfig: { disabled?: boolean } }>;
    public readonly clusterIpv4Cidr: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly enableBinaryAuthorization: pulumi.Output<boolean | undefined>;
    public readonly enableKubernetesAlpha: pulumi.Output<boolean | undefined>;
    public readonly enableLegacyAbac: pulumi.Output<boolean | undefined>;
    public readonly enableTpu: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly endpoint: pulumi.Output<string>;
    public readonly initialNodeCount: pulumi.Output<number | undefined>;
    public /*out*/ readonly instanceGroupUrls: pulumi.Output<string[]>;
    public readonly ipAllocationPolicy: pulumi.Output<{ clusterIpv4CidrBlock: string, clusterSecondaryRangeName: string, createSubnetwork?: boolean, servicesIpv4CidrBlock: string, servicesSecondaryRangeName: string, subnetworkName?: string } | undefined>;
    public readonly loggingService: pulumi.Output<string>;
    public readonly maintenancePolicy: pulumi.Output<{ dailyMaintenanceWindow: { duration: string, startTime: string } } | undefined>;
    public readonly masterAuth: pulumi.Output<{ clientCertificate: string, clientCertificateConfig?: { issueClientCertificate: boolean }, clientKey: string, clusterCaCertificate: string, password: string, username: string }>;
    public readonly masterAuthorizedNetworksConfig: pulumi.Output<{ cidrBlocks: { cidrBlock: string, displayName?: string }[] } | undefined>;
    public readonly masterIpv4CidrBlock: pulumi.Output<string | undefined>;
    public /*out*/ readonly masterVersion: pulumi.Output<string>;
    public readonly minMasterVersion: pulumi.Output<string | undefined>;
    public readonly monitoringService: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly network: pulumi.Output<string | undefined>;
    public readonly networkPolicy: pulumi.Output<{ enabled?: boolean, provider?: string }>;
    public readonly nodeConfig: pulumi.Output<{ diskSizeGb: number, diskType: string, guestAccelerators: { count: number, type: string }[], imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, minCpuPlatform?: string, oauthScopes: string[], preemptible?: boolean, serviceAccount: string, tags?: string[], taints?: { effect: string, key: string, value: string }[], workloadMetadataConfig?: { nodeMetadata: string } }>;
    public readonly nodePools: pulumi.Output<{ autoscaling?: { maxNodeCount: number, minNodeCount: number }, initialNodeCount: number, instanceGroupUrls: string[], management: { autoRepair?: boolean, autoUpgrade?: boolean }, maxPodsPerNode: number, name: string, namePrefix: string, nodeConfig: { diskSizeGb: number, diskType: string, guestAccelerators: { count: number, type: string }[], imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata?: {[key: string]: string}, minCpuPlatform?: string, oauthScopes: string[], preemptible?: boolean, serviceAccount: string, tags?: string[], taints?: { effect: string, key: string, value: string }[], workloadMetadataConfig?: { nodeMetadata: string } }, nodeCount: number, version: string }[]>;
    public readonly nodeVersion: pulumi.Output<string>;
    public readonly podSecurityPolicyConfig: pulumi.Output<{ enabled: boolean } | undefined>;
    public readonly privateCluster: pulumi.Output<boolean | undefined>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly removeDefaultNodePool: pulumi.Output<boolean | undefined>;
    public readonly resourceLabels: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly subnetwork: pulumi.Output<string>;
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ClusterState = argsOrState as ClusterState | undefined;
            inputs["additionalZones"] = state ? state.additionalZones : undefined;
            inputs["addonsConfig"] = state ? state.addonsConfig : undefined;
            inputs["clusterIpv4Cidr"] = state ? state.clusterIpv4Cidr : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enableBinaryAuthorization"] = state ? state.enableBinaryAuthorization : undefined;
            inputs["enableKubernetesAlpha"] = state ? state.enableKubernetesAlpha : undefined;
            inputs["enableLegacyAbac"] = state ? state.enableLegacyAbac : undefined;
            inputs["enableTpu"] = state ? state.enableTpu : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["initialNodeCount"] = state ? state.initialNodeCount : undefined;
            inputs["instanceGroupUrls"] = state ? state.instanceGroupUrls : undefined;
            inputs["ipAllocationPolicy"] = state ? state.ipAllocationPolicy : undefined;
            inputs["loggingService"] = state ? state.loggingService : undefined;
            inputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            inputs["masterAuth"] = state ? state.masterAuth : undefined;
            inputs["masterAuthorizedNetworksConfig"] = state ? state.masterAuthorizedNetworksConfig : undefined;
            inputs["masterIpv4CidrBlock"] = state ? state.masterIpv4CidrBlock : undefined;
            inputs["masterVersion"] = state ? state.masterVersion : undefined;
            inputs["minMasterVersion"] = state ? state.minMasterVersion : undefined;
            inputs["monitoringService"] = state ? state.monitoringService : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["networkPolicy"] = state ? state.networkPolicy : undefined;
            inputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            inputs["nodePools"] = state ? state.nodePools : undefined;
            inputs["nodeVersion"] = state ? state.nodeVersion : undefined;
            inputs["podSecurityPolicyConfig"] = state ? state.podSecurityPolicyConfig : undefined;
            inputs["privateCluster"] = state ? state.privateCluster : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["removeDefaultNodePool"] = state ? state.removeDefaultNodePool : undefined;
            inputs["resourceLabels"] = state ? state.resourceLabels : undefined;
            inputs["subnetwork"] = state ? state.subnetwork : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["additionalZones"] = args ? args.additionalZones : undefined;
            inputs["addonsConfig"] = args ? args.addonsConfig : undefined;
            inputs["clusterIpv4Cidr"] = args ? args.clusterIpv4Cidr : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enableBinaryAuthorization"] = args ? args.enableBinaryAuthorization : undefined;
            inputs["enableKubernetesAlpha"] = args ? args.enableKubernetesAlpha : undefined;
            inputs["enableLegacyAbac"] = args ? args.enableLegacyAbac : undefined;
            inputs["enableTpu"] = args ? args.enableTpu : undefined;
            inputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            inputs["ipAllocationPolicy"] = args ? args.ipAllocationPolicy : undefined;
            inputs["loggingService"] = args ? args.loggingService : undefined;
            inputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            inputs["masterAuth"] = args ? args.masterAuth : undefined;
            inputs["masterAuthorizedNetworksConfig"] = args ? args.masterAuthorizedNetworksConfig : undefined;
            inputs["masterIpv4CidrBlock"] = args ? args.masterIpv4CidrBlock : undefined;
            inputs["minMasterVersion"] = args ? args.minMasterVersion : undefined;
            inputs["monitoringService"] = args ? args.monitoringService : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["networkPolicy"] = args ? args.networkPolicy : undefined;
            inputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            inputs["nodePools"] = args ? args.nodePools : undefined;
            inputs["nodeVersion"] = args ? args.nodeVersion : undefined;
            inputs["podSecurityPolicyConfig"] = args ? args.podSecurityPolicyConfig : undefined;
            inputs["privateCluster"] = args ? args.privateCluster : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["removeDefaultNodePool"] = args ? args.removeDefaultNodePool : undefined;
            inputs["resourceLabels"] = args ? args.resourceLabels : undefined;
            inputs["subnetwork"] = args ? args.subnetwork : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["endpoint"] = undefined /*out*/;
            inputs["instanceGroupUrls"] = undefined /*out*/;
            inputs["masterVersion"] = undefined /*out*/;
        }
        super("gcp:container/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    readonly additionalZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly addonsConfig?: pulumi.Input<{ horizontalPodAutoscaling?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, httpLoadBalancing?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, kubernetesDashboard?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, networkPolicyConfig?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }> }>;
    readonly clusterIpv4Cidr?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly enableBinaryAuthorization?: pulumi.Input<boolean>;
    readonly enableKubernetesAlpha?: pulumi.Input<boolean>;
    readonly enableLegacyAbac?: pulumi.Input<boolean>;
    readonly enableTpu?: pulumi.Input<boolean>;
    readonly endpoint?: pulumi.Input<string>;
    readonly initialNodeCount?: pulumi.Input<number>;
    readonly instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>;
    readonly ipAllocationPolicy?: pulumi.Input<{ clusterIpv4CidrBlock?: pulumi.Input<string>, clusterSecondaryRangeName?: pulumi.Input<string>, createSubnetwork?: pulumi.Input<boolean>, servicesIpv4CidrBlock?: pulumi.Input<string>, servicesSecondaryRangeName?: pulumi.Input<string>, subnetworkName?: pulumi.Input<string> }>;
    readonly loggingService?: pulumi.Input<string>;
    readonly maintenancePolicy?: pulumi.Input<{ dailyMaintenanceWindow: pulumi.Input<{ duration?: pulumi.Input<string>, startTime: pulumi.Input<string> }> }>;
    readonly masterAuth?: pulumi.Input<{ clientCertificate?: pulumi.Input<string>, clientCertificateConfig?: pulumi.Input<{ issueClientCertificate: pulumi.Input<boolean> }>, clientKey?: pulumi.Input<string>, clusterCaCertificate?: pulumi.Input<string>, password: pulumi.Input<string>, username: pulumi.Input<string> }>;
    readonly masterAuthorizedNetworksConfig?: pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<{ cidrBlock: pulumi.Input<string>, displayName?: pulumi.Input<string> }>[]> }>;
    readonly masterIpv4CidrBlock?: pulumi.Input<string>;
    readonly masterVersion?: pulumi.Input<string>;
    readonly minMasterVersion?: pulumi.Input<string>;
    readonly monitoringService?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkPolicy?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, provider?: pulumi.Input<string> }>;
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    readonly nodePools?: pulumi.Input<pulumi.Input<{ autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>, initialNodeCount?: pulumi.Input<number>, instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>, management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>, maxPodsPerNode?: pulumi.Input<number>, name?: pulumi.Input<string>, namePrefix?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>, nodeCount?: pulumi.Input<number>, version?: pulumi.Input<string> }>[]>;
    readonly nodeVersion?: pulumi.Input<string>;
    readonly podSecurityPolicyConfig?: pulumi.Input<{ enabled: pulumi.Input<boolean> }>;
    readonly privateCluster?: pulumi.Input<boolean>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly removeDefaultNodePool?: pulumi.Input<boolean>;
    readonly resourceLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    readonly additionalZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly addonsConfig?: pulumi.Input<{ horizontalPodAutoscaling?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, httpLoadBalancing?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, kubernetesDashboard?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }>, networkPolicyConfig?: pulumi.Input<{ disabled?: pulumi.Input<boolean> }> }>;
    readonly clusterIpv4Cidr?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly enableBinaryAuthorization?: pulumi.Input<boolean>;
    readonly enableKubernetesAlpha?: pulumi.Input<boolean>;
    readonly enableLegacyAbac?: pulumi.Input<boolean>;
    readonly enableTpu?: pulumi.Input<boolean>;
    readonly initialNodeCount?: pulumi.Input<number>;
    readonly ipAllocationPolicy?: pulumi.Input<{ clusterIpv4CidrBlock?: pulumi.Input<string>, clusterSecondaryRangeName?: pulumi.Input<string>, createSubnetwork?: pulumi.Input<boolean>, servicesIpv4CidrBlock?: pulumi.Input<string>, servicesSecondaryRangeName?: pulumi.Input<string>, subnetworkName?: pulumi.Input<string> }>;
    readonly loggingService?: pulumi.Input<string>;
    readonly maintenancePolicy?: pulumi.Input<{ dailyMaintenanceWindow: pulumi.Input<{ duration?: pulumi.Input<string>, startTime: pulumi.Input<string> }> }>;
    readonly masterAuth?: pulumi.Input<{ clientCertificate?: pulumi.Input<string>, clientCertificateConfig?: pulumi.Input<{ issueClientCertificate: pulumi.Input<boolean> }>, clientKey?: pulumi.Input<string>, clusterCaCertificate?: pulumi.Input<string>, password: pulumi.Input<string>, username: pulumi.Input<string> }>;
    readonly masterAuthorizedNetworksConfig?: pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<{ cidrBlock: pulumi.Input<string>, displayName?: pulumi.Input<string> }>[]> }>;
    readonly masterIpv4CidrBlock?: pulumi.Input<string>;
    readonly minMasterVersion?: pulumi.Input<string>;
    readonly monitoringService?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly network?: pulumi.Input<string>;
    readonly networkPolicy?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, provider?: pulumi.Input<string> }>;
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    readonly nodePools?: pulumi.Input<pulumi.Input<{ autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>, initialNodeCount?: pulumi.Input<number>, instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>, management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>, maxPodsPerNode?: pulumi.Input<number>, name?: pulumi.Input<string>, namePrefix?: pulumi.Input<string>, nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>, nodeCount?: pulumi.Input<number>, version?: pulumi.Input<string> }>[]>;
    readonly nodeVersion?: pulumi.Input<string>;
    readonly podSecurityPolicyConfig?: pulumi.Input<{ enabled: pulumi.Input<boolean> }>;
    readonly privateCluster?: pulumi.Input<boolean>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly removeDefaultNodePool?: pulumi.Input<boolean>;
    readonly resourceLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly subnetwork?: pulumi.Input<string>;
    readonly zone?: pulumi.Input<string>;
}
