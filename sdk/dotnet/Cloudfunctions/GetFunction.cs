// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctions
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get information about a Google Cloud Function. For more information see
        /// the [official documentation](https://cloud.google.com/functions/docs/)
        /// and [API](https://cloud.google.com/functions/docs/apis).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/cloudfunctions_function.html.markdown.
        /// </summary>
        public static Task<GetFunctionResult> GetFunction(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("gcp:cloudfunctions/getFunction:getFunction", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a Cloud Function.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetFunctionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Available memory (in MB) to the function.
        /// </summary>
        public readonly int AvailableMemoryMb;
        /// <summary>
        /// Description of the function.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
        /// </summary>
        public readonly string EntryPoint;
        public readonly ImmutableDictionary<string, object> EnvironmentVariables;
        /// <summary>
        /// A source that fires events in response to a condition in another service. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionEventTriggersResult> EventTriggers;
        /// <summary>
        /// If function is triggered by HTTP, trigger URL is set here.
        /// </summary>
        public readonly string HttpsTriggerUrl;
        /// <summary>
        /// A map of labels applied to this function.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        public readonly int MaxInstances;
        /// <summary>
        /// The name of the Cloud Function.
        /// </summary>
        public readonly string Name;
        public readonly string? Project;
        public readonly string? Region;
        /// <summary>
        /// The runtime in which the function is running.
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// The service account email to be assumed by the cloud function.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// The GCS bucket containing the zip archive which contains the function.
        /// </summary>
        public readonly string SourceArchiveBucket;
        /// <summary>
        /// The source archive object (file) in archive bucket.
        /// </summary>
        public readonly string SourceArchiveObject;
        public readonly ImmutableArray<Outputs.GetFunctionSourceRepositoriesResult> SourceRepositories;
        /// <summary>
        /// Function execution timeout (in seconds).
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// If function is triggered by HTTP, this boolean is set.
        /// </summary>
        public readonly bool TriggerHttp;
        public readonly string VpcConnector;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFunctionResult(
            int availableMemoryMb,
            string description,
            string entryPoint,
            ImmutableDictionary<string, object> environmentVariables,
            ImmutableArray<Outputs.GetFunctionEventTriggersResult> eventTriggers,
            string httpsTriggerUrl,
            ImmutableDictionary<string, object> labels,
            int maxInstances,
            string name,
            string? project,
            string? region,
            string runtime,
            string serviceAccountEmail,
            string sourceArchiveBucket,
            string sourceArchiveObject,
            ImmutableArray<Outputs.GetFunctionSourceRepositoriesResult> sourceRepositories,
            int timeout,
            bool triggerHttp,
            string vpcConnector,
            string id)
        {
            AvailableMemoryMb = availableMemoryMb;
            Description = description;
            EntryPoint = entryPoint;
            EnvironmentVariables = environmentVariables;
            EventTriggers = eventTriggers;
            HttpsTriggerUrl = httpsTriggerUrl;
            Labels = labels;
            MaxInstances = maxInstances;
            Name = name;
            Project = project;
            Region = region;
            Runtime = runtime;
            ServiceAccountEmail = serviceAccountEmail;
            SourceArchiveBucket = sourceArchiveBucket;
            SourceArchiveObject = sourceArchiveObject;
            SourceRepositories = sourceRepositories;
            Timeout = timeout;
            TriggerHttp = triggerHttp;
            VpcConnector = vpcConnector;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetFunctionEventTriggersFailurePoliciesResult
    {
        /// <summary>
        /// Whether the function should be retried on failure.
        /// </summary>
        public readonly bool Retry;

        [OutputConstructor]
        private GetFunctionEventTriggersFailurePoliciesResult(bool retry)
        {
            Retry = retry;
        }
    }

    [OutputType]
    public sealed class GetFunctionEventTriggersResult
    {
        /// <summary>
        /// The type of event to observe. For example: `"google.storage.object.finalize"`.
        /// See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/)
        /// for a full reference of accepted triggers.
        /// </summary>
        public readonly string EventType;
        /// <summary>
        /// Policy for failed executions. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<GetFunctionEventTriggersFailurePoliciesResult> FailurePolicies;
        /// <summary>
        /// The name of the resource whose events are being observed, for example, `"myBucket"`
        /// </summary>
        public readonly string Resource;

        [OutputConstructor]
        private GetFunctionEventTriggersResult(
            string eventType,
            ImmutableArray<GetFunctionEventTriggersFailurePoliciesResult> failurePolicies,
            string resource)
        {
            EventType = eventType;
            FailurePolicies = failurePolicies;
            Resource = resource;
        }
    }

    [OutputType]
    public sealed class GetFunctionSourceRepositoriesResult
    {
        public readonly string DeployedUrl;
        public readonly string Url;

        [OutputConstructor]
        private GetFunctionSourceRepositoriesResult(
            string deployedUrl,
            string url)
        {
            DeployedUrl = deployedUrl;
            Url = url;
        }
    }
    }
}
