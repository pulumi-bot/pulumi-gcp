// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class BackendServiceIap
    {
        /// <summary>
        /// -
        /// (Required)
        /// OAuth2 Client ID for IAP
        /// </summary>
        public readonly string Oauth2ClientId;
        /// <summary>
        /// -
        /// (Required)
        /// OAuth2 Client Secret for IAP
        /// </summary>
        public readonly string Oauth2ClientSecret;
        /// <summary>
        /// -
        /// OAuth2 Client Secret SHA-256 for IAP
        /// </summary>
        public readonly string? Oauth2ClientSecretSha256;

        [OutputConstructor]
        private BackendServiceIap(
            string oauth2ClientId,

            string oauth2ClientSecret,

            string? oauth2ClientSecretSha256)
        {
            Oauth2ClientId = oauth2ClientId;
            Oauth2ClientSecret = oauth2ClientSecret;
            Oauth2ClientSecretSha256 = oauth2ClientSecretSha256;
        }
    }
}
