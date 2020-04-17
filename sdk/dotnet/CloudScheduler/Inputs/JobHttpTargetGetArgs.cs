// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudScheduler.Inputs
{

    public sealed class JobHttpTargetGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// HTTP request body.
        /// A request body is allowed only if the HTTP method is POST or PUT.
        /// It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// -
        /// (Optional)
        /// HTTP request headers.
        /// This map contains the header field names and values.
        /// Headers can be set when the job is created.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// -
        /// (Optional)
        /// Which HTTP method to use for the request.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Contains information needed for generating an OAuth token.
        /// This type of authorization should be used when sending requests to a GCP endpoint.  Structure is documented below.
        /// </summary>
        [Input("oauthToken")]
        public Input<Inputs.JobHttpTargetOauthTokenGetArgs>? OauthToken { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// Contains information needed for generating an OpenID Connect token.
        /// This type of authorization should be used when sending requests to third party endpoints or Cloud Run.  Structure is documented below.
        /// </summary>
        [Input("oidcToken")]
        public Input<Inputs.JobHttpTargetOidcTokenGetArgs>? OidcToken { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The full URI path that the request will be sent to.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public JobHttpTargetGetArgs()
        {
        }
    }
}
