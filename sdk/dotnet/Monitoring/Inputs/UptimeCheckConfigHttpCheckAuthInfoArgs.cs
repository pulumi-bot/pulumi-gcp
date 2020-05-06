// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Inputs
{

    public sealed class UptimeCheckConfigHttpCheckAuthInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password to authenticate.  **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username to authenticate.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UptimeCheckConfigHttpCheckAuthInfoArgs()
        {
        }
    }
}
