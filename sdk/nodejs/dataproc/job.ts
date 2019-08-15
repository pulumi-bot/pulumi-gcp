// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a job resource within a Dataproc cluster within GCE. For more information see
 * [the official dataproc documentation](https://cloud.google.com/dataproc/).
 * 
 * !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
 * 
 * ## Example usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const mycluster = new gcp.dataproc.Cluster("mycluster", {
 *     region: "us-central1",
 * });
 * // Submit an example pyspark job to a dataproc cluster
 * const pyspark = new gcp.dataproc.Job("pyspark", {
 *     forceDelete: true,
 *     placement: {
 *         clusterName: mycluster.name,
 *     },
 *     pysparkConfig: {
 *         mainPythonFileUri: "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py",
 *         properties: {
 *             "spark.logConf": "true",
 *         },
 *     },
 *     region: mycluster.region,
 * });
 * // Submit an example spark job to a dataproc cluster
 * const spark = new gcp.dataproc.Job("spark", {
 *     forceDelete: true,
 *     placement: {
 *         clusterName: mycluster.name,
 *     },
 *     region: mycluster.region,
 *     sparkConfig: {
 *         args: ["1000"],
 *         jarFileUris: ["file:///usr/lib/spark/examples/jars/spark-examples.jar"],
 *         loggingConfig: {
 *             driverLogLevels: {
 *                 root: "INFO",
 *             },
 *         },
 *         mainClass: "org.apache.spark.examples.SparkPi",
 *         properties: {
 *             "spark.logConf": "true",
 *         },
 *     },
 * });
 * 
 * export const pysparkStatus = pyspark.status.state;
 * // Check out current state of the jobs
 * export const sparkStatus = spark.status.state;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dataproc_job.html.markdown.
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    public /*out*/ readonly driverControlsFilesUri!: pulumi.Output<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    public /*out*/ readonly driverOutputResourceUri!: pulumi.Output<string>;
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    public readonly hadoopConfig!: pulumi.Output<outputApi.dataproc.JobHadoopConfig | undefined>;
    public readonly hiveConfig!: pulumi.Output<outputApi.dataproc.JobHiveConfig | undefined>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly pigConfig!: pulumi.Output<outputApi.dataproc.JobPigConfig | undefined>;
    public readonly placement!: pulumi.Output<outputApi.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly pysparkConfig!: pulumi.Output<outputApi.dataproc.JobPysparkConfig | undefined>;
    public readonly reference!: pulumi.Output<outputApi.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * Optional. Job scheduling configuration.
     */
    public readonly scheduling!: pulumi.Output<outputApi.dataproc.JobScheduling | undefined>;
    public readonly sparkConfig!: pulumi.Output<outputApi.dataproc.JobSparkConfig | undefined>;
    public readonly sparksqlConfig!: pulumi.Output<outputApi.dataproc.JobSparksqlConfig | undefined>;
    public /*out*/ readonly status!: pulumi.Output<outputApi.dataproc.JobStatus>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["driverControlsFilesUri"] = state ? state.driverControlsFilesUri : undefined;
            inputs["driverOutputResourceUri"] = state ? state.driverOutputResourceUri : undefined;
            inputs["forceDelete"] = state ? state.forceDelete : undefined;
            inputs["hadoopConfig"] = state ? state.hadoopConfig : undefined;
            inputs["hiveConfig"] = state ? state.hiveConfig : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["pigConfig"] = state ? state.pigConfig : undefined;
            inputs["placement"] = state ? state.placement : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pysparkConfig"] = state ? state.pysparkConfig : undefined;
            inputs["reference"] = state ? state.reference : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["scheduling"] = state ? state.scheduling : undefined;
            inputs["sparkConfig"] = state ? state.sparkConfig : undefined;
            inputs["sparksqlConfig"] = state ? state.sparksqlConfig : undefined;
            inputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if (!args || args.placement === undefined) {
                throw new Error("Missing required property 'placement'");
            }
            inputs["forceDelete"] = args ? args.forceDelete : undefined;
            inputs["hadoopConfig"] = args ? args.hadoopConfig : undefined;
            inputs["hiveConfig"] = args ? args.hiveConfig : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["pigConfig"] = args ? args.pigConfig : undefined;
            inputs["placement"] = args ? args.placement : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pysparkConfig"] = args ? args.pysparkConfig : undefined;
            inputs["reference"] = args ? args.reference : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["sparkConfig"] = args ? args.sparkConfig : undefined;
            inputs["sparksqlConfig"] = args ? args.sparksqlConfig : undefined;
            inputs["driverControlsFilesUri"] = undefined /*out*/;
            inputs["driverOutputResourceUri"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
     */
    readonly driverControlsFilesUri?: pulumi.Input<string>;
    /**
     * A URI pointing to the location of the stdout of the job's driver program.
     */
    readonly driverOutputResourceUri?: pulumi.Input<string>;
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    readonly hadoopConfig?: pulumi.Input<inputApi.dataproc.JobHadoopConfig>;
    readonly hiveConfig?: pulumi.Input<inputApi.dataproc.JobHiveConfig>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly pigConfig?: pulumi.Input<inputApi.dataproc.JobPigConfig>;
    readonly placement?: pulumi.Input<inputApi.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pysparkConfig?: pulumi.Input<inputApi.dataproc.JobPysparkConfig>;
    readonly reference?: pulumi.Input<inputApi.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    readonly scheduling?: pulumi.Input<inputApi.dataproc.JobScheduling>;
    readonly sparkConfig?: pulumi.Input<inputApi.dataproc.JobSparkConfig>;
    readonly sparksqlConfig?: pulumi.Input<inputApi.dataproc.JobSparksqlConfig>;
    readonly status?: pulumi.Input<inputApi.dataproc.JobStatus>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * By default, you can only delete inactive jobs within
     * Dataproc. Setting this to true, and calling destroy, will ensure that the
     * job is first cancelled before issuing the delete.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    readonly hadoopConfig?: pulumi.Input<inputApi.dataproc.JobHadoopConfig>;
    readonly hiveConfig?: pulumi.Input<inputApi.dataproc.JobHiveConfig>;
    /**
     * The list of labels (key/value pairs) to add to the job.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly pigConfig?: pulumi.Input<inputApi.dataproc.JobPigConfig>;
    readonly placement: pulumi.Input<inputApi.dataproc.JobPlacement>;
    /**
     * The project in which the `cluster` can be found and jobs
     * subsequently run against. If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly pysparkConfig?: pulumi.Input<inputApi.dataproc.JobPysparkConfig>;
    readonly reference?: pulumi.Input<inputApi.dataproc.JobReference>;
    /**
     * The Cloud Dataproc region. This essentially determines which clusters are available
     * for this job to be submitted to. If not specified, defaults to `global`.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Optional. Job scheduling configuration.
     */
    readonly scheduling?: pulumi.Input<inputApi.dataproc.JobScheduling>;
    readonly sparkConfig?: pulumi.Input<inputApi.dataproc.JobSparkConfig>;
    readonly sparksqlConfig?: pulumi.Input<inputApi.dataproc.JobSparksqlConfig>;
}
