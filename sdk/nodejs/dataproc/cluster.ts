// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Cloud Dataproc cluster resource within GCP. For more information see
 * [the official dataproc documentation](https://cloud.google.com/dataproc/).
 * 
 * 
 * !> **Warning:** Due to limitations of the API, all arguments except
 * `labels`,`cluster_config.worker_config.num_instances` and `cluster_config.preemptible_worker_config.num_instances` are non-updateable. Changing others will cause recreation of the
 * whole cluster!
 * 
 * ## Example Usage - Basic
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const simplecluster = new gcp.dataproc.Cluster("simplecluster", {
 *     region: "us-central1",
 * });
 * ```
 * 
 * ## Example Usage - Advanced
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const mycluster = new gcp.dataproc.Cluster("mycluster", {
 *     clusterConfig: {
 *         gceClusterConfig: {
 *             //network = "${google_compute_network.dataproc_network.name}"
 *             tags: [
 *                 "foo",
 *                 "bar",
 *             ],
 *         },
 *         // You can define multiple initializationAction blocks
 *         initializationActions: [{
 *             script: "gs://dataproc-initialization-actions/stackdriver/stackdriver.sh",
 *             timeoutSec: 500,
 *         }],
 *         masterConfig: {
 *             diskConfig: {
 *                 bootDiskSizeGb: 15,
 *                 bootDiskType: "pd-ssd",
 *             },
 *             machineType: "n1-standard-1",
 *             numInstances: 1,
 *         },
 *         preemptibleWorkerConfig: {
 *             numInstances: 0,
 *         },
 *         // Override or set some custom properties
 *         softwareConfig: {
 *             imageVersion: "1.3.7-deb9",
 *             overrideProperties: {
 *                 "dataproc:dataproc.allow.zero.workers": "true",
 *             },
 *         },
 *         stagingBucket: "dataproc-staging-bucket",
 *         workerConfig: {
 *             diskConfig: {
 *                 bootDiskSizeGb: 15,
 *                 numLocalSsds: 1,
 *             },
 *             machineType: "n1-standard-1",
 *             minCpuPlatform: "Intel Skylake",
 *             numInstances: 2,
 *         },
 *     },
 *     labels: {
 *         foo: "bar",
 *     },
 *     region: "us-central1",
 * });
 * ```
 * 
 * ## Example Usage - Using a GPU accelerator
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const acceleratedCluster = new gcp.dataproc.Cluster("acceleratedCluster", {
 *     clusterConfig: {
 *         gceClusterConfig: {
 *             zone: "us-central1-a",
 *         },
 *         masterConfig: {
 *             accelerators: [{
 *                 acceleratorCount: 1,
 *                 acceleratorType: "nvidia-tesla-k80",
 *             }],
 *         },
 *     },
 *     region: "us-central1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_cluster.html.markdown.
 */
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

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    public readonly clusterConfig!: pulumi.Output<{ bucket: string, encryptionConfig?: { kmsKeyName: string }, gceClusterConfig: { internalIpOnly?: boolean, metadata?: {[key: string]: string}, network: string, serviceAccount?: string, serviceAccountScopes: string[], subnetwork?: string, tags?: string[], zone: string }, initializationActions?: { script: string, timeoutSec?: number }[], masterConfig: { accelerators?: { acceleratorCount: number, acceleratorType: string }[], diskConfig: { bootDiskSizeGb: number, bootDiskType?: string, numLocalSsds: number }, imageUri: string, instanceNames: string[], machineType: string, minCpuPlatform: string, numInstances: number }, preemptibleWorkerConfig: { diskConfig: { bootDiskSizeGb: number, bootDiskType?: string, numLocalSsds: number }, instanceNames: string[], numInstances: number }, softwareConfig: { imageVersion: string, optionalComponents?: string[], overrideProperties?: {[key: string]: string}, properties: {[key: string]: any} }, stagingBucket?: string, workerConfig: { accelerators?: { acceleratorCount: number, acceleratorType: string }[], diskConfig: { bootDiskSizeGb: number, bootDiskType?: string, numLocalSsds: number }, imageUri: string, instanceNames: string[], machineType: string, minCpuPlatform: string, numInstances: number } }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    public readonly region!: pulumi.Output<string | undefined>;

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
            const state = argsOrState as ClusterState | undefined;
            inputs["clusterConfig"] = state ? state.clusterConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<{ bucket?: pulumi.Input<string>, encryptionConfig?: pulumi.Input<{ kmsKeyName: pulumi.Input<string> }>, gceClusterConfig?: pulumi.Input<{ internalIpOnly?: pulumi.Input<boolean>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, network?: pulumi.Input<string>, serviceAccount?: pulumi.Input<string>, serviceAccountScopes?: pulumi.Input<pulumi.Input<string>[]>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone?: pulumi.Input<string> }>, initializationActions?: pulumi.Input<pulumi.Input<{ script: pulumi.Input<string>, timeoutSec?: pulumi.Input<number> }>[]>, masterConfig?: pulumi.Input<{ accelerators?: pulumi.Input<pulumi.Input<{ acceleratorCount: pulumi.Input<number>, acceleratorType: pulumi.Input<string> }>[]>, diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, imageUri?: pulumi.Input<string>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, minCpuPlatform?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }>, preemptibleWorkerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, numInstances?: pulumi.Input<number> }>, softwareConfig?: pulumi.Input<{ imageVersion?: pulumi.Input<string>, optionalComponents?: pulumi.Input<pulumi.Input<string>[]>, overrideProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, properties?: pulumi.Input<{[key: string]: any}> }>, stagingBucket?: pulumi.Input<string>, workerConfig?: pulumi.Input<{ accelerators?: pulumi.Input<pulumi.Input<{ acceleratorCount: pulumi.Input<number>, acceleratorType: pulumi.Input<string> }>[]>, diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, imageUri?: pulumi.Input<string>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, minCpuPlatform?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }> }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Allows you to configure various aspects of the cluster.
     * Structure defined below.
     */
    readonly clusterConfig?: pulumi.Input<{ bucket?: pulumi.Input<string>, encryptionConfig?: pulumi.Input<{ kmsKeyName: pulumi.Input<string> }>, gceClusterConfig?: pulumi.Input<{ internalIpOnly?: pulumi.Input<boolean>, metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, network?: pulumi.Input<string>, serviceAccount?: pulumi.Input<string>, serviceAccountScopes?: pulumi.Input<pulumi.Input<string>[]>, subnetwork?: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]>, zone?: pulumi.Input<string> }>, initializationActions?: pulumi.Input<pulumi.Input<{ script: pulumi.Input<string>, timeoutSec?: pulumi.Input<number> }>[]>, masterConfig?: pulumi.Input<{ accelerators?: pulumi.Input<pulumi.Input<{ acceleratorCount: pulumi.Input<number>, acceleratorType: pulumi.Input<string> }>[]>, diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, imageUri?: pulumi.Input<string>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, minCpuPlatform?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }>, preemptibleWorkerConfig?: pulumi.Input<{ diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, numInstances?: pulumi.Input<number> }>, softwareConfig?: pulumi.Input<{ imageVersion?: pulumi.Input<string>, optionalComponents?: pulumi.Input<pulumi.Input<string>[]>, overrideProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, properties?: pulumi.Input<{[key: string]: any}> }>, stagingBucket?: pulumi.Input<string>, workerConfig?: pulumi.Input<{ accelerators?: pulumi.Input<pulumi.Input<{ acceleratorCount: pulumi.Input<number>, acceleratorType: pulumi.Input<string> }>[]>, diskConfig?: pulumi.Input<{ bootDiskSizeGb?: pulumi.Input<number>, bootDiskType?: pulumi.Input<string>, numLocalSsds?: pulumi.Input<number> }>, imageUri?: pulumi.Input<string>, instanceNames?: pulumi.Input<pulumi.Input<string>[]>, machineType?: pulumi.Input<string>, minCpuPlatform?: pulumi.Input<string>, numInstances?: pulumi.Input<number> }> }>;
    /**
     * The list of labels (key/value pairs) to be applied to
     * instances in the cluster. GCP generates some itself including `goog-dataproc-cluster-name`
     * which is the name of the cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the cluster, unique within the project and
     * zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the `cluster` will exist. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region in which the cluster and associated nodes will be created in.
     * Defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
}
