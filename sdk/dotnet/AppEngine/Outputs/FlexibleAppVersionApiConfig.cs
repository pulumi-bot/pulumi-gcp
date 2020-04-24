// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Outputs
{

    [OutputType]
    public sealed class FlexibleAppVersionApiConfig
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Action to take when users access resources that require authentication. Defaults to "AUTH_FAIL_ACTION_REDIRECT".
        /// </summary>
        public readonly string? AuthFailAction;
        /// <summary>
        /// -
        /// (Optional)
        /// Level of login required to access this resource. Defaults to "LOGIN_OPTIONAL".
        /// </summary>
        public readonly string? Login;
        /// <summary>
        /// -
        /// (Required)
        /// Path to the script from the application root directory.
        /// </summary>
        public readonly string Script;
        /// <summary>
        /// -
        /// (Optional)
        /// Security (HTTPS) enforcement for this URL.
        /// </summary>
        public readonly string? SecurityLevel;
        /// <summary>
        /// -
        /// (Optional)
        /// URL to serve the endpoint at.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private FlexibleAppVersionApiConfig(
            string? authFailAction,

            string? login,

            string script,

            string? securityLevel,

            string? url)
        {
            AuthFailAction = authFailAction;
            Login = login;
            Script = script;
            SecurityLevel = securityLevel;
            Url = url;
        }
    }
}
