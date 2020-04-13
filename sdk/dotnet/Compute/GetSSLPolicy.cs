// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Gets an SSL Policy within GCE from its name, for use with Target HTTPS and Target SSL Proxies.
        ///     For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_ssl_policy.html.markdown.
        /// </summary>
        [Obsolete("Use GetSSLPolicy.InvokeAsync() instead")]
        public static Task<GetSSLPolicyResult> GetSSLPolicy(GetSSLPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSSLPolicyResult>("gcp:compute/getSSLPolicy:getSSLPolicy", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSSLPolicy
    {
        /// <summary>
        /// Gets an SSL Policy within GCE from its name, for use with Target HTTPS and Target SSL Proxies.
        ///     For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies).
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_ssl_policy.html.markdown.
        /// </summary>
        public static Task<GetSSLPolicyResult> InvokeAsync(GetSSLPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSSLPolicyResult>("gcp:compute/getSSLPolicy:getSSLPolicy", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSSLPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the SSL Policy.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetSSLPolicyArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSSLPolicyResult
    {
        public readonly string CreationTimestamp;
        /// <summary>
        /// If the `profile` is `CUSTOM`, these are the custom encryption
        /// ciphers supported by the profile. If the `profile` is *not* `CUSTOM`, this
        /// attribute will be empty.
        /// </summary>
        public readonly ImmutableArray<string> CustomFeatures;
        /// <summary>
        /// Description of this SSL Policy.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The set of enabled encryption ciphers as a result of the policy config
        /// </summary>
        public readonly ImmutableArray<string> EnabledFeatures;
        /// <summary>
        /// Fingerprint of this resource.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The minimum supported TLS version of this policy.
        /// </summary>
        public readonly string MinTlsVersion;
        public readonly string Name;
        /// <summary>
        /// The Google-curated or custom profile used by this policy.
        /// </summary>
        public readonly string Profile;
        public readonly string? Project;
        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSSLPolicyResult(
            string creationTimestamp,
            ImmutableArray<string> customFeatures,
            string description,
            ImmutableArray<string> enabledFeatures,
            string fingerprint,
            string minTlsVersion,
            string name,
            string profile,
            string? project,
            string selfLink,
            string id)
        {
            CreationTimestamp = creationTimestamp;
            CustomFeatures = customFeatures;
            Description = description;
            EnabledFeatures = enabledFeatures;
            Fingerprint = fingerprint;
            MinTlsVersion = minTlsVersion;
            Name = name;
            Profile = profile;
            Project = project;
            SelfLink = selfLink;
            Id = id;
        }
    }
}
