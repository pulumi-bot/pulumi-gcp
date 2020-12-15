// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// Creates a new Transfer Job in Google Cloud Storage Transfer.
    /// 
    /// To get more information about Google Cloud Storage Transfer, see:
    /// 
    /// * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
    /// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
    /// * How-to Guides
    ///     * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
    /// 
    /// ## Example Usage
    /// 
    /// Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = Output.Create(Gcp.Storage.GetTransferProjectServieAccount.InvokeAsync(new Gcp.Storage.GetTransferProjectServieAccountArgs
    ///         {
    ///             Project = @var.Project,
    ///         }));
    ///         var s3_backup_bucketBucket = new Gcp.Storage.Bucket("s3-backup-bucketBucket", new Gcp.Storage.BucketArgs
    ///         {
    ///             StorageClass = "NEARLINE",
    ///             Project = @var.Project,
    ///         });
    ///         var s3_backup_bucketBucketIAMMember = new Gcp.Storage.BucketIAMMember("s3-backup-bucketBucketIAMMember", new Gcp.Storage.BucketIAMMemberArgs
    ///         {
    ///             Bucket = s3_backup_bucketBucket.Name,
    ///             Role = "roles/storage.admin",
    ///             Member = @default.Apply(@default =&gt; $"serviceAccount:{_default.Email}"),
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 s3_backup_bucketBucket,
    ///             },
    ///         });
    ///         var s3_bucket_nightly_backup = new Gcp.Storage.TransferJob("s3-bucket-nightly-backup", new Gcp.Storage.TransferJobArgs
    ///         {
    ///             Description = "Nightly backup of S3 bucket",
    ///             Project = @var.Project,
    ///             TransferSpec = new Gcp.Storage.Inputs.TransferJobTransferSpecArgs
    ///             {
    ///                 ObjectConditions = new Gcp.Storage.Inputs.TransferJobTransferSpecObjectConditionsArgs
    ///                 {
    ///                     MaxTimeElapsedSinceLastModification = "600s",
    ///                     ExcludePrefixes = 
    ///                     {
    ///                         "requests.gz",
    ///                     },
    ///                 },
    ///                 TransferOptions = new Gcp.Storage.Inputs.TransferJobTransferSpecTransferOptionsArgs
    ///                 {
    ///                     DeleteObjectsUniqueInSink = false,
    ///                 },
    ///                 AwsS3DataSource = new Gcp.Storage.Inputs.TransferJobTransferSpecAwsS3DataSourceArgs
    ///                 {
    ///                     BucketName = @var.Aws_s3_bucket,
    ///                     AwsAccessKey = new Gcp.Storage.Inputs.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs
    ///                     {
    ///                         AccessKeyId = @var.Aws_access_key,
    ///                         SecretAccessKey = @var.Aws_secret_key,
    ///                     },
    ///                 },
    ///                 GcsDataSink = new Gcp.Storage.Inputs.TransferJobTransferSpecGcsDataSinkArgs
    ///                 {
    ///                     BucketName = s3_backup_bucketBucket.Name,
    ///                 },
    ///             },
    ///             Schedule = new Gcp.Storage.Inputs.TransferJobScheduleArgs
    ///             {
    ///                 ScheduleStartDate = new Gcp.Storage.Inputs.TransferJobScheduleScheduleStartDateArgs
    ///                 {
    ///                     Year = 2018,
    ///                     Month = 10,
    ///                     Day = 1,
    ///                 },
    ///                 ScheduleEndDate = new Gcp.Storage.Inputs.TransferJobScheduleScheduleEndDateArgs
    ///                 {
    ///                     Year = 2019,
    ///                     Month = 1,
    ///                     Day = 15,
    ///                 },
    ///                 StartTimeOfDay = new Gcp.Storage.Inputs.TransferJobScheduleStartTimeOfDayArgs
    ///                 {
    ///                     Hours = 23,
    ///                     Minutes = 30,
    ///                     Seconds = 0,
    ///                     Nanos = 0,
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 s3_backup_bucketBucketIAMMember,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Storage buckets can be imported using the Transfer Job's `project` and `name` without the `transferJob/` prefix, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/transferJob:TransferJob nightly-backup-transfer-job my-project-1asd32/8422144862922355674
    /// ```
    /// </summary>
    [GcpResourceType("gcp:storage/transferJob:TransferJob")]
    public partial class TransferJob : Pulumi.CustomResource
    {
        /// <summary>
        /// When the Transfer Job was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// When the Transfer Job was deleted.
        /// </summary>
        [Output("deletionTime")]
        public Output<string> DeletionTime { get; private set; } = null!;

        /// <summary>
        /// Unique description to identify the Transfer Job.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// When the Transfer Job was last modified.
        /// </summary>
        [Output("lastModificationTime")]
        public Output<string> LastModificationTime { get; private set; } = null!;

        /// <summary>
        /// The name of the Transfer Job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.TransferJobSchedule> Schedule { get; private set; } = null!;

        /// <summary>
        /// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Transfer specification. Structure documented below.
        /// </summary>
        [Output("transferSpec")]
        public Output<Outputs.TransferJobTransferSpec> TransferSpec { get; private set; } = null!;


        /// <summary>
        /// Create a TransferJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransferJob(string name, TransferJobArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/transferJob:TransferJob", name, args ?? new TransferJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransferJob(string name, Input<string> id, TransferJobState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/transferJob:TransferJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TransferJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransferJob Get(string name, Input<string> id, TransferJobState? state = null, CustomResourceOptions? options = null)
        {
            return new TransferJob(name, id, state, options);
        }
    }

    public sealed class TransferJobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique description to identify the Transfer Job.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.TransferJobScheduleArgs> Schedule { get; set; } = null!;

        /// <summary>
        /// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Transfer specification. Structure documented below.
        /// </summary>
        [Input("transferSpec", required: true)]
        public Input<Inputs.TransferJobTransferSpecArgs> TransferSpec { get; set; } = null!;

        public TransferJobArgs()
        {
        }
    }

    public sealed class TransferJobState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When the Transfer Job was created.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// When the Transfer Job was deleted.
        /// </summary>
        [Input("deletionTime")]
        public Input<string>? DeletionTime { get; set; }

        /// <summary>
        /// Unique description to identify the Transfer Job.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When the Transfer Job was last modified.
        /// </summary>
        [Input("lastModificationTime")]
        public Input<string>? LastModificationTime { get; set; }

        /// <summary>
        /// The name of the Transfer Job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.TransferJobScheduleGetArgs>? Schedule { get; set; }

        /// <summary>
        /// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Transfer specification. Structure documented below.
        /// </summary>
        [Input("transferSpec")]
        public Input<Inputs.TransferJobTransferSpecGetArgs>? TransferSpec { get; set; }

        public TransferJobState()
        {
        }
    }
}
