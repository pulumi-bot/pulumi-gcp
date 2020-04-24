// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class ManagedSslCertificateManagedArgs : Pulumi.ResourceArgs
    {
        [Input("domains", required: true)]
        private InputList<string>? _domains;

        /// <summary>
        /// Domains for which a managed SSL certificate will be valid.  Currently,
        /// there can be up to 100 domains in this list.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        public ManagedSslCertificateManagedArgs()
        {
        }
    }
}
