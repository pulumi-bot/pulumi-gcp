// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.storage.getTransferProjectServieAccount({
 *     project: _var.project,
 * });
 * const s3_backup_bucketBucket = new gcp.storage.Bucket("s3-backup-bucketBucket", {
 *     storageClass: "NEARLINE",
 *     project: _var.project,
 * });
 * const s3_backup_bucketBucketIAMMember = new gcp.storage.BucketIAMMember("s3-backup-bucketBucketIAMMember", {
 *     bucket: s3_backup_bucketBucket.name,
 *     role: "roles/storage.admin",
 *     member: _default.then(_default => `serviceAccount:${_default.email}`),
 * });
 * const s3BucketNightlyBackup = new gcp.storage.TransferJob("s3-bucket-nightly-backup", {
 *     description: "Nightly backup of S3 bucket",
 *     project: _var.project,
 *     transfer_spec: {
 *         object_conditions: {
 *             maxTimeElapsedSinceLastModification: "600s",
 *             excludePrefixes: ["requests.gz"],
 *         },
 *         transfer_options: {
 *             deleteObjectsUniqueInSink: false,
 *         },
 *         aws_s3_data_source: {
 *             bucketName: _var.aws_s3_bucket,
 *             aws_access_key: {
 *                 accessKeyId: _var.aws_access_key,
 *                 secretAccessKey: _var.aws_secret_key,
 *             },
 *         },
 *         gcs_data_sink: {
 *             bucketName: s3_backup_bucketBucket.name,
 *         },
 *     },
 *     schedule: {
 *         schedule_start_date: {
 *             year: 2018,
 *             month: 10,
 *             day: 1,
 *         },
 *         schedule_end_date: {
 *             year: 2019,
 *             month: 1,
 *             day: 15,
 *         },
 *         start_time_of_day: {
 *             hours: 23,
 *             minutes: 30,
 *             seconds: 0,
 *             nanos: 0,
 *         },
 *     },
 * });
 * ```
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
    public readonly schedule!: pulumi.Output<outputs.storage.TransferJobSchedule>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Transfer specification. Structure documented below.
     */
    public readonly transferSpec!: pulumi.Output<outputs.storage.TransferJobTransferSpec>;

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
    readonly schedule?: pulumi.Input<inputs.storage.TransferJobSchedule>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Transfer specification. Structure documented below.
     */
    readonly transferSpec?: pulumi.Input<inputs.storage.TransferJobTransferSpec>;
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
    readonly schedule: pulumi.Input<inputs.storage.TransferJobSchedule>;
    /**
     * Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Transfer specification. Structure documented below.
     */
    readonly transferSpec: pulumi.Input<inputs.storage.TransferJobTransferSpec>;
}
