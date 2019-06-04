// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
 * the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
 * and [the API reference](https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters.nodePools).
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:container/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * Configuration required by cluster autoscaler to adjust
     * the size of the node pool to the current cluster usage. Structure is documented below.
     */
    public readonly autoscaling!: pulumi.Output<{ maxNodeCount: number, minNodeCount: number } | undefined>;
    /**
     * The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * The initial node count for the pool. Changing this will force
     * recreation of the resource.
     */
    public readonly initialNodeCount!: pulumi.Output<number>;
    public /*out*/ readonly instanceGroupUrls!: pulumi.Output<string[]>;
    /**
     * The location (region or zone) in which the cluster
     * resides.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Node management configuration, wherein auto-repair and
     * auto-upgrade is configured. Structure is documented below.
     */
    public readonly management!: pulumi.Output<{ autoRepair?: boolean, autoUpgrade?: boolean }>;
    /**
     * ) The maximum number of pods per node in this node pool.
     * Note that this does not work on node pools which are "route-based" - that is, node
     * pools belonging to clusters that do not have IP Aliasing enabled.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
     * for more information.
     */
    public readonly maxPodsPerNode!: pulumi.Output<number>;
    /**
     * The name of the node pool. If left blank, Terraform will
     * auto-generate a unique name.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The node configuration of the pool. See
     * google_container_cluster for schema.
     */
    public readonly nodeConfig!: pulumi.Output<{ diskSizeGb: number, diskType: string, guestAccelerators: { count: number, type: string }[], imageType: string, labels?: {[key: string]: string}, localSsdCount: number, machineType: string, metadata: {[key: string]: string}, minCpuPlatform?: string, oauthScopes: string[], preemptible?: boolean, serviceAccount: string, tags?: string[], taints?: { effect: string, key: string, value: string }[], workloadMetadataConfig?: { nodeMetadata: string } }>;
    /**
     * The number of nodes per instance group. This field can be used to
     * update the number of nodes per instance group but should not be used alongside `autoscaling`.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * The ID of the project in which to create the node pool. If blank,
     * the provider-configured project will be used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the cluster resides (for
     * regional clusters). `zone` has been deprecated in favor of `location`.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The Kubernetes version for the nodes in this pool. Note that if this field
     * and `auto_upgrade` are both specified, they will fight each other for what the node version should
     * be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
     * recommended that you specify explicit versions as Terraform will see spurious diffs
     * when fuzzy versions are used. See the `google_container_engine_versions` data source's
     * `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * The zone in which the cluster resides. `zone`
     * has been deprecated in favor of `location`.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            inputs["autoscaling"] = state ? state.autoscaling : undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["initialNodeCount"] = state ? state.initialNodeCount : undefined;
            inputs["instanceGroupUrls"] = state ? state.instanceGroupUrls : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["management"] = state ? state.management : undefined;
            inputs["maxPodsPerNode"] = state ? state.maxPodsPerNode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if (!args || args.cluster === undefined) {
                throw new Error("Missing required property 'cluster'");
            }
            inputs["autoscaling"] = args ? args.autoscaling : undefined;
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["initialNodeCount"] = args ? args.initialNodeCount : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["management"] = args ? args.management : undefined;
            inputs["maxPodsPerNode"] = args ? args.maxPodsPerNode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["instanceGroupUrls"] = undefined /*out*/;
        }
        super(NodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * Configuration required by cluster autoscaler to adjust
     * the size of the node pool to the current cluster usage. Structure is documented below.
     */
    readonly autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>;
    /**
     * The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
     */
    readonly cluster?: pulumi.Input<string>;
    /**
     * The initial node count for the pool. Changing this will force
     * recreation of the resource.
     */
    readonly initialNodeCount?: pulumi.Input<number>;
    readonly instanceGroupUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The location (region or zone) in which the cluster
     * resides.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Node management configuration, wherein auto-repair and
     * auto-upgrade is configured. Structure is documented below.
     */
    readonly management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>;
    /**
     * ) The maximum number of pods per node in this node pool.
     * Note that this does not work on node pools which are "route-based" - that is, node
     * pools belonging to clusters that do not have IP Aliasing enabled.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
     * for more information.
     */
    readonly maxPodsPerNode?: pulumi.Input<number>;
    /**
     * The name of the node pool. If left blank, Terraform will
     * auto-generate a unique name.
     */
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The node configuration of the pool. See
     * google_container_cluster for schema.
     */
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    /**
     * The number of nodes per instance group. This field can be used to
     * update the number of nodes per instance group but should not be used alongside `autoscaling`.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The ID of the project in which to create the node pool. If blank,
     * the provider-configured project will be used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster resides (for
     * regional clusters). `zone` has been deprecated in favor of `location`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The Kubernetes version for the nodes in this pool. Note that if this field
     * and `auto_upgrade` are both specified, they will fight each other for what the node version should
     * be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
     * recommended that you specify explicit versions as Terraform will see spurious diffs
     * when fuzzy versions are used. See the `google_container_engine_versions` data source's
     * `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The zone in which the cluster resides. `zone`
     * has been deprecated in favor of `location`.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Configuration required by cluster autoscaler to adjust
     * the size of the node pool to the current cluster usage. Structure is documented below.
     */
    readonly autoscaling?: pulumi.Input<{ maxNodeCount: pulumi.Input<number>, minNodeCount: pulumi.Input<number> }>;
    /**
     * The cluster to create the node pool for.  Cluster must be present in `zone` provided for zonal clusters.
     */
    readonly cluster: pulumi.Input<string>;
    /**
     * The initial node count for the pool. Changing this will force
     * recreation of the resource.
     */
    readonly initialNodeCount?: pulumi.Input<number>;
    /**
     * The location (region or zone) in which the cluster
     * resides.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Node management configuration, wherein auto-repair and
     * auto-upgrade is configured. Structure is documented below.
     */
    readonly management?: pulumi.Input<{ autoRepair?: pulumi.Input<boolean>, autoUpgrade?: pulumi.Input<boolean> }>;
    /**
     * ) The maximum number of pods per node in this node pool.
     * Note that this does not work on node pools which are "route-based" - that is, node
     * pools belonging to clusters that do not have IP Aliasing enabled.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
     * for more information.
     */
    readonly maxPodsPerNode?: pulumi.Input<number>;
    /**
     * The name of the node pool. If left blank, Terraform will
     * auto-generate a unique name.
     */
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The node configuration of the pool. See
     * google_container_cluster for schema.
     */
    readonly nodeConfig?: pulumi.Input<{ diskSizeGb?: pulumi.Input<number>, diskType?: pulumi.Input<string>, guestAccelerators?: pulumi.Input<pulumi.Input<{ count: pulumi.Input<number>, type: pulumi.Input<string> }>[]>, imageType?: pulumi.Input<string>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, localSsdCount?: pulumi.Input<number>, machineType?: pulumi.Input<string>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, minCpuPlatform?: pulumi.Input<string>, oauthScopes?: pulumi.Input<pulumi.Input<string>[]>, preemptible?: pulumi.Input<boolean>, serviceAccount?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, taints?: pulumi.Input<pulumi.Input<{ effect: pulumi.Input<string>, key: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, workloadMetadataConfig?: pulumi.Input<{ nodeMetadata: pulumi.Input<string> }> }>;
    /**
     * The number of nodes per instance group. This field can be used to
     * update the number of nodes per instance group but should not be used alongside `autoscaling`.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The ID of the project in which to create the node pool. If blank,
     * the provider-configured project will be used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster resides (for
     * regional clusters). `zone` has been deprecated in favor of `location`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The Kubernetes version for the nodes in this pool. Note that if this field
     * and `auto_upgrade` are both specified, they will fight each other for what the node version should
     * be, so setting both is highly discouraged. While a fuzzy version can be specified, it's
     * recommended that you specify explicit versions as Terraform will see spurious diffs
     * when fuzzy versions are used. See the `google_container_engine_versions` data source's
     * `version_prefix` field to approximate fuzzy versions in a Terraform-compatible way.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The zone in which the cluster resides. `zone`
     * has been deprecated in favor of `location`.
     */
    readonly zone?: pulumi.Input<string>;
}
