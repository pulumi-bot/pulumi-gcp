// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SecurityScanConfigAuthenticationGoogleAccountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password of the custom account. The credential is stored encrypted
        /// in GCP.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The user name of the custom account.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SecurityScanConfigAuthenticationGoogleAccountGetArgs()
        {
        }
    }
}
