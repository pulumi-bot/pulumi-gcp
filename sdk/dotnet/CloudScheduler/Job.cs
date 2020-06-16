// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudScheduler
{
    /// <summary>
    /// A scheduled job that can publish a pubsub message or a http request
    /// every X interval of time, using crontab format string.
    /// 
    /// To use Cloud Scheduler your project must contain an App Engine app
    /// that is located in one of the supported regions. If your project
    /// does not have an App Engine app, you must create one.
    /// 
    /// To get more information about Job, see:
    /// 
    /// * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/scheduler/)
    /// 
    /// {{% examples %}}
    /// ## Example Usage
    /// {{% example %}}
    /// ### Scheduler Job Http
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var job = new Gcp.CloudScheduler.Job("job", new Gcp.CloudScheduler.JobArgs
    ///         {
    ///             AttemptDeadline = "320s",
    ///             Description = "test http job",
    ///             HttpTarget = new Gcp.CloudScheduler.Inputs.JobHttpTargetArgs
    ///             {
    ///                 HttpMethod = "POST",
    ///                 Uri = "https://example.com/ping",
    ///             },
    ///             RetryConfig = new Gcp.CloudScheduler.Inputs.JobRetryConfigArgs
    ///             {
    ///                 RetryCount = 1,
    ///             },
    ///             Schedule = "*/8 * * * *",
    ///             TimeZone = "America/New_York",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// {{% example %}}
    /// ### Scheduler Job App Engine
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var job = new Gcp.CloudScheduler.Job("job", new Gcp.CloudScheduler.JobArgs
    ///         {
    ///             AppEngineHttpTarget = new Gcp.CloudScheduler.Inputs.JobAppEngineHttpTargetArgs
    ///             {
    ///                 AppEngineRouting = new Gcp.CloudScheduler.Inputs.JobAppEngineHttpTargetAppEngineRoutingArgs
    ///                 {
    ///                     Instance = "my-instance-001",
    ///                     Service = "web",
    ///                     Version = "prod",
    ///                 },
    ///                 HttpMethod = "POST",
    ///                 RelativeUri = "/ping",
    ///             },
    ///             AttemptDeadline = "320s",
    ///             Description = "test app engine job",
    ///             RetryConfig = new Gcp.CloudScheduler.Inputs.JobRetryConfigArgs
    ///             {
    ///                 MaxDoublings = 2,
    ///                 MaxRetryDuration = "10s",
    ///                 MinBackoffDuration = "1s",
    ///                 RetryCount = 3,
    ///             },
    ///             Schedule = "*/4 * * * *",
    ///             TimeZone = "Europe/London",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// {{% /example %}}
    /// {{% /examples %}}
    /// </summary>
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// App Engine HTTP target.
        /// If the job providers a App Engine HTTP target the cron will
        /// send a request to the service instance  Structure is documented below.
        /// </summary>
        [Output("appEngineHttpTarget")]
        public Output<Outputs.JobAppEngineHttpTarget?> AppEngineHttpTarget { get; private set; } = null!;

        /// <summary>
        /// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
        /// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
        /// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
        /// The allowed duration for this deadline is:
        /// * For HTTP targets, between 15 seconds and 30 minutes.
        /// * For App Engine HTTP targets, between 15 seconds and 24 hours.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
        /// </summary>
        [Output("attemptDeadline")]
        public Output<string?> AttemptDeadline { get; private set; } = null!;

        /// <summary>
        /// A human-readable description for the job.
        /// This string must not contain more than 500 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// HTTP target.
        /// If the job providers a http_target the cron will
        /// send a request to the targeted url  Structure is documented below.
        /// </summary>
        [Output("httpTarget")]
        public Output<Outputs.JobHttpTarget?> HttpTarget { get; private set; } = null!;

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Pub/Sub target
        /// If the job providers a Pub/Sub target the cron will publish
        /// a message to the provided topic  Structure is documented below.
        /// </summary>
        [Output("pubsubTarget")]
        public Output<Outputs.JobPubsubTarget?> PubsubTarget { get; private set; } = null!;

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// By default, if a job does not complete successfully,
        /// meaning that an acknowledgement is not received from the handler,
        /// then it will be retried with exponential backoff according to the settings  Structure is documented below.
        /// </summary>
        [Output("retryConfig")]
        public Output<Outputs.JobRetryConfig?> RetryConfig { get; private set; } = null!;

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule.
        /// The value of this field must be a time zone name from the tz database.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:cloudscheduler/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudscheduler/job:Job", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App Engine HTTP target.
        /// If the job providers a App Engine HTTP target the cron will
        /// send a request to the service instance  Structure is documented below.
        /// </summary>
        [Input("appEngineHttpTarget")]
        public Input<Inputs.JobAppEngineHttpTargetArgs>? AppEngineHttpTarget { get; set; }

        /// <summary>
        /// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
        /// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
        /// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
        /// The allowed duration for this deadline is:
        /// * For HTTP targets, between 15 seconds and 30 minutes.
        /// * For App Engine HTTP targets, between 15 seconds and 24 hours.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
        /// </summary>
        [Input("attemptDeadline")]
        public Input<string>? AttemptDeadline { get; set; }

        /// <summary>
        /// A human-readable description for the job.
        /// This string must not contain more than 500 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTTP target.
        /// If the job providers a http_target the cron will
        /// send a request to the targeted url  Structure is documented below.
        /// </summary>
        [Input("httpTarget")]
        public Input<Inputs.JobHttpTargetArgs>? HttpTarget { get; set; }

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Pub/Sub target
        /// If the job providers a Pub/Sub target the cron will publish
        /// a message to the provided topic  Structure is documented below.
        /// </summary>
        [Input("pubsubTarget")]
        public Input<Inputs.JobPubsubTargetArgs>? PubsubTarget { get; set; }

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// By default, if a job does not complete successfully,
        /// meaning that an acknowledgement is not received from the handler,
        /// then it will be retried with exponential backoff according to the settings  Structure is documented below.
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.JobRetryConfigArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule.
        /// The value of this field must be a time zone name from the tz database.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public JobArgs()
        {
        }
    }

    public sealed class JobState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// App Engine HTTP target.
        /// If the job providers a App Engine HTTP target the cron will
        /// send a request to the service instance  Structure is documented below.
        /// </summary>
        [Input("appEngineHttpTarget")]
        public Input<Inputs.JobAppEngineHttpTargetGetArgs>? AppEngineHttpTarget { get; set; }

        /// <summary>
        /// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
        /// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
        /// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
        /// The allowed duration for this deadline is:
        /// * For HTTP targets, between 15 seconds and 30 minutes.
        /// * For App Engine HTTP targets, between 15 seconds and 24 hours.
        /// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
        /// </summary>
        [Input("attemptDeadline")]
        public Input<string>? AttemptDeadline { get; set; }

        /// <summary>
        /// A human-readable description for the job.
        /// This string must not contain more than 500 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTTP target.
        /// If the job providers a http_target the cron will
        /// send a request to the targeted url  Structure is documented below.
        /// </summary>
        [Input("httpTarget")]
        public Input<Inputs.JobHttpTargetGetArgs>? HttpTarget { get; set; }

        /// <summary>
        /// The name of the job.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Pub/Sub target
        /// If the job providers a Pub/Sub target the cron will publish
        /// a message to the provided topic  Structure is documented below.
        /// </summary>
        [Input("pubsubTarget")]
        public Input<Inputs.JobPubsubTargetGetArgs>? PubsubTarget { get; set; }

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// By default, if a job does not complete successfully,
        /// meaning that an acknowledgement is not received from the handler,
        /// then it will be retried with exponential backoff according to the settings  Structure is documented below.
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.JobRetryConfigGetArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule.
        /// The value of this field must be a time zone name from the tz database.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public JobState()
        {
        }
    }
}
