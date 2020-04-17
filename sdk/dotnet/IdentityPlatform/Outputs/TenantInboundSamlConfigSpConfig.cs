// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform.Outputs
{

    [OutputType]
    public sealed class TenantInboundSamlConfigSpConfig
    {
        /// <summary>
        /// -
        /// (Required)
        /// Callback URI where responses from IDP are handled. Must start with `https://`.
        /// </summary>
        public readonly string CallbackUri;
        /// <summary>
        /// -
        /// The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.TenantInboundSamlConfigSpConfigSpCertificate> SpCertificates;
        /// <summary>
        /// -
        /// (Required)
        /// Unique identifier for all SAML entities.
        /// </summary>
        public readonly string SpEntityId;

        [OutputConstructor]
        private TenantInboundSamlConfigSpConfig(
            string callbackUri,

            ImmutableArray<Outputs.TenantInboundSamlConfigSpConfigSpCertificate> spCertificates,

            string spEntityId)
        {
            CallbackUri = callbackUri;
            SpCertificates = spCertificates;
            SpEntityId = spEntityId;
        }
    }
}
