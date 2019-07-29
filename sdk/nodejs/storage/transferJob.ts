// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new Transfer Job in Google Cloud Storage Transfer.
 * 
 * To get more information about Google Cloud Storage Transfer, see:
 * 
 * * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
 * * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
 * * How-to Guides
 *     * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
 * 
 * ## Example Usage
 * 
 * Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultTransferProjectServieAccount = pulumi.output(gcp.storage.getTransferProjectServieAccount({
 *     project: var_project,
 * }));
 * const s3_backup_bucketBucket = new gcp.storage.Bucket("s3-backup-bucket", {
 *     project: var_project,
 *     storageClass: "NEARLINE",
 * });
 * const s3_backup_bucketBucketIAMMember = new gcp.storage.BucketIAMMember("s3-backup-bucket", {
 *     bucket: s3_backup_bucketBucket.name,
 *     member: pulumi.interpolate`serviceAccount:${defaultTransferProjectServieAccount.email}`,
 *     role: "roles/storage.admin",
 * }, {dependsOn: [s3_backup_bucketBucket]});
 * const s3_bucket_nightly_backup = new gcp.storage.TransferJob("s3-bucket-nightly-backup", {
 *     description: "Nightly backup of S3 bucket",
 *     project: var_project,
 *     schedule: {
 *         scheduleEndDate: {
 *             day: 15,
 *             month: 1,
 *             year: 2019,
 *         },
 *         scheduleStartDate: {
 *             day: 1,
 *             month: 10,
 *             year: 2018,
 *         },
 *         startTimeOfDay: {
 *             hours: 23,
 *             minutes: 30,
 *             nanos: 0,
 *             seconds: 0,
 *         },
 *     },
 *     transferSpec: {
 *         awsS3DataSource: {
 *             awsAccessKey: {
 *                 accessKeyId: var_aws_access_key,
 *                 secretAccessKey: var_aws_secret_key,
 *             },
 *             bucketName: var_aws_s3_bucket,
 *         },
 *         gcsDataSink: {
 *             bucketName: s3_backup_bucketBucket.name,
 *         },
 *         objectConditions: {
 *             excludePrefixes: ["requests.gz"],
 *             maxTimeElapsedSinceLastModification: "600s",
 *         },
 *         transferOptions: {
 *             deleteObjectsUniqueInSink: false,
 *         },
 *     },
 * }, {dependsOn: [s3_backup_bucketBucketIAMMember]});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_transfer_job.html.markdown.
 */
export class TransferJob extends pulumi.CustomResource {
    /**
     * Get an existing TransferJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransferJobState, opts?: pulumi.CustomResourceOptions): TransferJob {
        return new TransferJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:storage/transferJob:TransferJob';

    /**
     * Returns true if the given object is an instance of TransferJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransferJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransferJob.__pulumiType;
    }

    /**
     * When the Transfer Job was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * When the Transfer Job was deleted.
     */
    public /*out*/ readonly deletionTime!: pulumi.Output<string>;
    /**
     * Unique description to identify the Transfer Job.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * When the Transfer Job was last modified.
     */
    public /*out*/ readonly lastModificationTime!: pulumi.Output<string>;
    /**
     * The name of the Transfer Job.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
     */
    public readonly schedule!: pulumi.Output<{ scheduleEndDate?: { day: number, month: number, year: number }, scheduleStartDate: { day: number, month: number, year: number }, startTimeOfDay?: { hours: number, minutes: number, nanos: number, seconds: number } }>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Transfer specification. Structure documented below.
     */
    public readonly transferSpec!: pulumi.Output<{ awsS3DataSource?: { awsAccessKey: { accessKeyId: string, secretAccessKey: string }, bucketName: string }, gcsDataSink?: { bucketName: string }, gcsDataSource?: { bucketName: string }, httpDataSource?: { listUrl: string }, objectConditions?: { excludePrefixes?: string[], includePrefixes?: string[], maxTimeElapsedSinceLastModification?: string, minTimeElapsedSinceLastModification?: string }, transferOptions?: { deleteObjectsFromSourceAfterTransfer?: boolean, deleteObjectsUniqueInSink?: boolean, overwriteObjectsAlreadyExistingInSink?: boolean } }>;

    /**
     * Create a TransferJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransferJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransferJobArgs | TransferJobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TransferJobState | undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["deletionTime"] = state ? state.deletionTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lastModificationTime"] = state ? state.lastModificationTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["transferSpec"] = state ? state.transferSpec : undefined;
        } else {
            const args = argsOrState as TransferJobArgs | undefined;
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.schedule === undefined) {
                throw new Error("Missing required property 'schedule'");
            }
            if (!args || args.transferSpec === undefined) {
                throw new Error("Missing required property 'transferSpec'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["transferSpec"] = args ? args.transferSpec : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["deletionTime"] = undefined /*out*/;
            inputs["lastModificationTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TransferJob.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransferJob resources.
 */
export interface TransferJobState {
    /**
     * When the Transfer Job was created.
     */
    readonly creationTime?: pulumi.Input<string>;
    /**
     * When the Transfer Job was deleted.
     */
    readonly deletionTime?: pulumi.Input<string>;
    /**
     * Unique description to identify the Transfer Job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When the Transfer Job was last modified.
     */
    readonly lastModificationTime?: pulumi.Input<string>;
    /**
     * The name of the Transfer Job.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
     */
    readonly schedule?: pulumi.Input<{ scheduleEndDate?: pulumi.Input<{ day: pulumi.Input<number>, month: pulumi.Input<number>, year: pulumi.Input<number> }>, scheduleStartDate: pulumi.Input<{ day: pulumi.Input<number>, month: pulumi.Input<number>, year: pulumi.Input<number> }>, startTimeOfDay?: pulumi.Input<{ hours: pulumi.Input<number>, minutes: pulumi.Input<number>, nanos: pulumi.Input<number>, seconds: pulumi.Input<number> }> }>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Transfer specification. Structure documented below.
     */
    readonly transferSpec?: pulumi.Input<{ awsS3DataSource?: pulumi.Input<{ awsAccessKey: pulumi.Input<{ accessKeyId: pulumi.Input<string>, secretAccessKey: pulumi.Input<string> }>, bucketName: pulumi.Input<string> }>, gcsDataSink?: pulumi.Input<{ bucketName: pulumi.Input<string> }>, gcsDataSource?: pulumi.Input<{ bucketName: pulumi.Input<string> }>, httpDataSource?: pulumi.Input<{ listUrl: pulumi.Input<string> }>, objectConditions?: pulumi.Input<{ excludePrefixes?: pulumi.Input<pulumi.Input<string>[]>, includePrefixes?: pulumi.Input<pulumi.Input<string>[]>, maxTimeElapsedSinceLastModification?: pulumi.Input<string>, minTimeElapsedSinceLastModification?: pulumi.Input<string> }>, transferOptions?: pulumi.Input<{ deleteObjectsFromSourceAfterTransfer?: pulumi.Input<boolean>, deleteObjectsUniqueInSink?: pulumi.Input<boolean>, overwriteObjectsAlreadyExistingInSink?: pulumi.Input<boolean> }> }>;
}

/**
 * The set of arguments for constructing a TransferJob resource.
 */
export interface TransferJobArgs {
    /**
     * Unique description to identify the Transfer Job.
     */
    readonly description: pulumi.Input<string>;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below.
     */
    readonly schedule: pulumi.Input<{ scheduleEndDate?: pulumi.Input<{ day: pulumi.Input<number>, month: pulumi.Input<number>, year: pulumi.Input<number> }>, scheduleStartDate: pulumi.Input<{ day: pulumi.Input<number>, month: pulumi.Input<number>, year: pulumi.Input<number> }>, startTimeOfDay?: pulumi.Input<{ hours: pulumi.Input<number>, minutes: pulumi.Input<number>, nanos: pulumi.Input<number>, seconds: pulumi.Input<number> }> }>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Transfer specification. Structure documented below.
     */
    readonly transferSpec: pulumi.Input<{ awsS3DataSource?: pulumi.Input<{ awsAccessKey: pulumi.Input<{ accessKeyId: pulumi.Input<string>, secretAccessKey: pulumi.Input<string> }>, bucketName: pulumi.Input<string> }>, gcsDataSink?: pulumi.Input<{ bucketName: pulumi.Input<string> }>, gcsDataSource?: pulumi.Input<{ bucketName: pulumi.Input<string> }>, httpDataSource?: pulumi.Input<{ listUrl: pulumi.Input<string> }>, objectConditions?: pulumi.Input<{ excludePrefixes?: pulumi.Input<pulumi.Input<string>[]>, includePrefixes?: pulumi.Input<pulumi.Input<string>[]>, maxTimeElapsedSinceLastModification?: pulumi.Input<string>, minTimeElapsedSinceLastModification?: pulumi.Input<string> }>, transferOptions?: pulumi.Input<{ deleteObjectsFromSourceAfterTransfer?: pulumi.Input<boolean>, deleteObjectsUniqueInSink?: pulumi.Input<boolean>, overwriteObjectsAlreadyExistingInSink?: pulumi.Input<boolean> }> }>;
}
