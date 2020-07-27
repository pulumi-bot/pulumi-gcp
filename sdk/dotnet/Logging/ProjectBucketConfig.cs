// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// Manages a project-level logging bucket config. For more information see
    /// [the official logging documentation](https://cloud.google.com/logging/docs/) and
    /// [Storing Logs](https://cloud.google.com/logging/docs/storage).
    /// 
    /// &gt; **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.Organizations.Project("default", new Gcp.Organizations.ProjectArgs
    ///         {
    ///             ProjectId = "your-project-id",
    ///             OrgId = "123456789",
    ///         });
    ///         var basic = new Gcp.Logging.ProjectBucketConfig("basic", new Gcp.Logging.ProjectBucketConfigArgs
    ///         {
    ///             Project = @default.Name,
    ///             Location = "global",
    ///             RetentionDays = 30,
    ///             BucketId = "_Default",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ProjectBucketConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Output("bucketId")]
        public Output<string> BucketId { get; private set; } = null!;

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        /// </summary>
        [Output("lifecycleState")]
        public Output<string> LifecycleState { get; private set; } = null!;

        /// <summary>
        /// The location of the bucket. The supported locations are: "global" "us-central1"
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectBucketConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectBucketConfig(string name, ProjectBucketConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, args ?? new ProjectBucketConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectBucketConfig(string name, Input<string> id, ProjectBucketConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectBucketConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectBucketConfig Get(string name, Input<string> id, ProjectBucketConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectBucketConfig(name, id, state, options);
        }
    }

    public sealed class ProjectBucketConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The location of the bucket. The supported locations are: "global" "us-central1"
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public ProjectBucketConfigArgs()
        {
        }
    }

    public sealed class ProjectBucketConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId")]
        public Input<string>? BucketId { get; set; }

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        /// </summary>
        [Input("lifecycleState")]
        public Input<string>? LifecycleState { get; set; }

        /// <summary>
        /// The location of the bucket. The supported locations are: "global" "us-central1"
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public ProjectBucketConfigState()
        {
        }
    }
}
