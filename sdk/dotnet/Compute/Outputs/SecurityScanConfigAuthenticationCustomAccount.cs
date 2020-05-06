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
    public sealed class SecurityScanConfigAuthenticationCustomAccount
    {
        /// <summary>
        /// The login form URL of the website.
        /// </summary>
        public readonly string LoginUrl;
        /// <summary>
        /// The password of the custom account. The credential is stored encrypted
        /// in GCP.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The user name of the custom account.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private SecurityScanConfigAuthenticationCustomAccount(
            string loginUrl,

            string password,

            string username)
        {
            LoginUrl = loginUrl;
            Password = password;
            Username = username;
        }
    }
}
