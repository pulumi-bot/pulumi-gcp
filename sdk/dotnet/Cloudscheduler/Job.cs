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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloud_scheduler_job.html.markdown.
    /// </summary>
    public partial class Job : Pulumi.CustomResource
    {
        /// <summary>
        /// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the
        /// service instance
        /// </summary>
        [Output("appEngineHttpTarget")]
        public Output<Outputs.JobAppEngineHttpTarget?> AppEngineHttpTarget { get; private set; } = null!;

        /// <summary>
        /// A human-readable description for the job. This string must not contain more than 500 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
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
        /// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
        /// </summary>
        [Output("pubsubTarget")]
        public Output<Outputs.JobPubsubTarget?> PubsubTarget { get; private set; } = null!;

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from
        /// the handler, then it will be retried with exponential backoff according to the settings
        /// </summary>
        [Output("retryConfig")]
        public Output<Outputs.JobRetryConfig?> RetryConfig { get; private set; } = null!;

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone
        /// name from the tz database.
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
            : base("gcp:cloudscheduler/job:Job", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        /// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the
        /// service instance
        /// </summary>
        [Input("appEngineHttpTarget")]
        public Input<Inputs.JobAppEngineHttpTargetArgs>? AppEngineHttpTarget { get; set; }

        /// <summary>
        /// A human-readable description for the job. This string must not contain more than 500 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
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
        /// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
        /// </summary>
        [Input("pubsubTarget")]
        public Input<Inputs.JobPubsubTargetArgs>? PubsubTarget { get; set; }

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from
        /// the handler, then it will be retried with exponential backoff according to the settings
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.JobRetryConfigArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone
        /// name from the tz database.
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
        /// App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the
        /// service instance
        /// </summary>
        [Input("appEngineHttpTarget")]
        public Input<Inputs.JobAppEngineHttpTargetGetArgs>? AppEngineHttpTarget { get; set; }

        /// <summary>
        /// A human-readable description for the job. This string must not contain more than 500 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// HTTP target. If the job providers a http_target the cron will send a request to the targeted url
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
        /// Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic
        /// </summary>
        [Input("pubsubTarget")]
        public Input<Inputs.JobPubsubTargetGetArgs>? PubsubTarget { get; set; }

        /// <summary>
        /// Region where the scheduler job resides
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// By default, if a job does not complete successfully, meaning that an acknowledgement is not received from
        /// the handler, then it will be retried with exponential backoff according to the settings
        /// </summary>
        [Input("retryConfig")]
        public Input<Inputs.JobRetryConfigGetArgs>? RetryConfig { get; set; }

        /// <summary>
        /// Describes the schedule on which the job will be executed.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone
        /// name from the tz database.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public JobState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class JobAppEngineHttpTargetAppEngineRoutingArgs : Pulumi.ResourceArgs
    {
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("service")]
        public Input<string>? Service { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public JobAppEngineHttpTargetAppEngineRoutingArgs()
        {
        }
    }

    public sealed class JobAppEngineHttpTargetAppEngineRoutingGetArgs : Pulumi.ResourceArgs
    {
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("service")]
        public Input<string>? Service { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public JobAppEngineHttpTargetAppEngineRoutingGetArgs()
        {
        }
    }

    public sealed class JobAppEngineHttpTargetArgs : Pulumi.ResourceArgs
    {
        [Input("appEngineRouting")]
        public Input<JobAppEngineHttpTargetAppEngineRoutingArgs>? AppEngineRouting { get; set; }

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("relativeUri", required: true)]
        public Input<string> RelativeUri { get; set; } = null!;

        public JobAppEngineHttpTargetArgs()
        {
        }
    }

    public sealed class JobAppEngineHttpTargetGetArgs : Pulumi.ResourceArgs
    {
        [Input("appEngineRouting")]
        public Input<JobAppEngineHttpTargetAppEngineRoutingGetArgs>? AppEngineRouting { get; set; }

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("relativeUri", required: true)]
        public Input<string> RelativeUri { get; set; } = null!;

        public JobAppEngineHttpTargetGetArgs()
        {
        }
    }

    public sealed class JobHttpTargetArgs : Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("oauthToken")]
        public Input<JobHttpTargetOauthTokenArgs>? OauthToken { get; set; }

        [Input("oidcToken")]
        public Input<JobHttpTargetOidcTokenArgs>? OidcToken { get; set; }

        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public JobHttpTargetArgs()
        {
        }
    }

    public sealed class JobHttpTargetGetArgs : Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("oauthToken")]
        public Input<JobHttpTargetOauthTokenGetArgs>? OauthToken { get; set; }

        [Input("oidcToken")]
        public Input<JobHttpTargetOidcTokenGetArgs>? OidcToken { get; set; }

        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public JobHttpTargetGetArgs()
        {
        }
    }

    public sealed class JobHttpTargetOauthTokenArgs : Pulumi.ResourceArgs
    {
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        public JobHttpTargetOauthTokenArgs()
        {
        }
    }

    public sealed class JobHttpTargetOauthTokenGetArgs : Pulumi.ResourceArgs
    {
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        public JobHttpTargetOauthTokenGetArgs()
        {
        }
    }

    public sealed class JobHttpTargetOidcTokenArgs : Pulumi.ResourceArgs
    {
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        public JobHttpTargetOidcTokenArgs()
        {
        }
    }

    public sealed class JobHttpTargetOidcTokenGetArgs : Pulumi.ResourceArgs
    {
        [Input("audience")]
        public Input<string>? Audience { get; set; }

        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        public JobHttpTargetOidcTokenGetArgs()
        {
        }
    }

    public sealed class JobPubsubTargetArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public JobPubsubTargetArgs()
        {
        }
    }

    public sealed class JobPubsubTargetGetArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("data")]
        public Input<string>? Data { get; set; }

        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        public JobPubsubTargetGetArgs()
        {
        }
    }

    public sealed class JobRetryConfigArgs : Pulumi.ResourceArgs
    {
        [Input("maxBackoffDuration")]
        public Input<string>? MaxBackoffDuration { get; set; }

        [Input("maxDoublings")]
        public Input<int>? MaxDoublings { get; set; }

        [Input("maxRetryDuration")]
        public Input<string>? MaxRetryDuration { get; set; }

        [Input("minBackoffDuration")]
        public Input<string>? MinBackoffDuration { get; set; }

        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        public JobRetryConfigArgs()
        {
        }
    }

    public sealed class JobRetryConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("maxBackoffDuration")]
        public Input<string>? MaxBackoffDuration { get; set; }

        [Input("maxDoublings")]
        public Input<int>? MaxDoublings { get; set; }

        [Input("maxRetryDuration")]
        public Input<string>? MaxRetryDuration { get; set; }

        [Input("minBackoffDuration")]
        public Input<string>? MinBackoffDuration { get; set; }

        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        public JobRetryConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class JobAppEngineHttpTarget
    {
        public readonly JobAppEngineHttpTargetAppEngineRouting? AppEngineRouting;
        public readonly string? Body;
        public readonly ImmutableDictionary<string, string>? Headers;
        public readonly string? HttpMethod;
        public readonly string RelativeUri;

        [OutputConstructor]
        private JobAppEngineHttpTarget(
            JobAppEngineHttpTargetAppEngineRouting? appEngineRouting,
            string? body,
            ImmutableDictionary<string, string>? headers,
            string? httpMethod,
            string relativeUri)
        {
            AppEngineRouting = appEngineRouting;
            Body = body;
            Headers = headers;
            HttpMethod = httpMethod;
            RelativeUri = relativeUri;
        }
    }

    [OutputType]
    public sealed class JobAppEngineHttpTargetAppEngineRouting
    {
        public readonly string? Instance;
        public readonly string? Service;
        public readonly string? Version;

        [OutputConstructor]
        private JobAppEngineHttpTargetAppEngineRouting(
            string? instance,
            string? service,
            string? version)
        {
            Instance = instance;
            Service = service;
            Version = version;
        }
    }

    [OutputType]
    public sealed class JobHttpTarget
    {
        public readonly string? Body;
        public readonly ImmutableDictionary<string, string>? Headers;
        public readonly string? HttpMethod;
        public readonly JobHttpTargetOauthToken? OauthToken;
        public readonly JobHttpTargetOidcToken? OidcToken;
        public readonly string Uri;

        [OutputConstructor]
        private JobHttpTarget(
            string? body,
            ImmutableDictionary<string, string>? headers,
            string? httpMethod,
            JobHttpTargetOauthToken? oauthToken,
            JobHttpTargetOidcToken? oidcToken,
            string uri)
        {
            Body = body;
            Headers = headers;
            HttpMethod = httpMethod;
            OauthToken = oauthToken;
            OidcToken = oidcToken;
            Uri = uri;
        }
    }

    [OutputType]
    public sealed class JobHttpTargetOauthToken
    {
        public readonly string? Scope;
        public readonly string ServiceAccountEmail;

        [OutputConstructor]
        private JobHttpTargetOauthToken(
            string? scope,
            string serviceAccountEmail)
        {
            Scope = scope;
            ServiceAccountEmail = serviceAccountEmail;
        }
    }

    [OutputType]
    public sealed class JobHttpTargetOidcToken
    {
        public readonly string? Audience;
        public readonly string ServiceAccountEmail;

        [OutputConstructor]
        private JobHttpTargetOidcToken(
            string? audience,
            string serviceAccountEmail)
        {
            Audience = audience;
            ServiceAccountEmail = serviceAccountEmail;
        }
    }

    [OutputType]
    public sealed class JobPubsubTarget
    {
        public readonly ImmutableDictionary<string, string>? Attributes;
        public readonly string? Data;
        public readonly string TopicName;

        [OutputConstructor]
        private JobPubsubTarget(
            ImmutableDictionary<string, string>? attributes,
            string? data,
            string topicName)
        {
            Attributes = attributes;
            Data = data;
            TopicName = topicName;
        }
    }

    [OutputType]
    public sealed class JobRetryConfig
    {
        public readonly string? MaxBackoffDuration;
        public readonly int? MaxDoublings;
        public readonly string? MaxRetryDuration;
        public readonly string? MinBackoffDuration;
        public readonly int? RetryCount;

        [OutputConstructor]
        private JobRetryConfig(
            string? maxBackoffDuration,
            int? maxDoublings,
            string? maxRetryDuration,
            string? minBackoffDuration,
            int? retryCount)
        {
            MaxBackoffDuration = maxBackoffDuration;
            MaxDoublings = maxDoublings;
            MaxRetryDuration = maxRetryDuration;
            MinBackoffDuration = minBackoffDuration;
            RetryCount = retryCount;
        }
    }
    }
}
