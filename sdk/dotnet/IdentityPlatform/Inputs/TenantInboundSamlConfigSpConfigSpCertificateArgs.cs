// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform.Inputs
{

    public sealed class TenantInboundSamlConfigSpConfigSpCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// The x509 certificate
        /// </summary>
        [Input("x509Certificate")]
        public Input<string>? X509Certificate { get; set; }

        public TenantInboundSamlConfigSpConfigSpCertificateArgs()
        {
        }
    }
}
