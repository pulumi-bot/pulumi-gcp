// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudScheduler.Outputs
{

    [OutputType]
    public sealed class JobAppEngineHttpTarget
    {
        /// <summary>
        /// -
        /// (Optional)
        /// App Engine Routing setting for the job.  Structure is documented below.
        /// </summary>
        public readonly Outputs.JobAppEngineHttpTargetAppEngineRouting? AppEngineRouting;
        /// <summary>
        /// -
        /// (Optional)
        /// HTTP request body.
        /// A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        /// It is an error to set body on a job with an incompatible HttpMethod.
        /// </summary>
        public readonly string? Body;
        /// <summary>
        /// -
        /// (Optional)
        /// This map contains the header field names and values.
        /// Repeated headers are not supported, but a header value can contain commas.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Headers;
        /// <summary>
        /// -
        /// (Optional)
        /// Which HTTP method to use for the request.
        /// </summary>
        public readonly string? HttpMethod;
        /// <summary>
        /// -
        /// (Required)
        /// The relative URI.
        /// The relative URL must begin with "/" and must be a valid HTTP relative URL.
        /// It can contain a path, query string arguments, and \# fragments.
        /// If the relative URL is empty, then the root path "/" will be used.
        /// No spaces are allowed, and the maximum length allowed is 2083 characters
        /// </summary>
        public readonly string RelativeUri;

        [OutputConstructor]
        private JobAppEngineHttpTarget(
            Outputs.JobAppEngineHttpTargetAppEngineRouting? appEngineRouting,

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
}
