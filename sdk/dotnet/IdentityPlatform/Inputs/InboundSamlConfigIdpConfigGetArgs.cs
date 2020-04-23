// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform.Inputs
{

    public sealed class InboundSamlConfigIdpConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("idpCertificates", required: true)]
        private InputList<Inputs.InboundSamlConfigIdpConfigIdpCertificateGetArgs>? _idpCertificates;

        /// <summary>
        /// -
        /// (Required)
        /// The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.InboundSamlConfigIdpConfigIdpCertificateGetArgs> IdpCertificates
        {
            get => _idpCertificates ?? (_idpCertificates = new InputList<Inputs.InboundSamlConfigIdpConfigIdpCertificateGetArgs>());
            set => _idpCertificates = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Unique identifier for all SAML entities
        /// </summary>
        [Input("idpEntityId", required: true)]
        public Input<string> IdpEntityId { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// Indicates if outbounding SAMLRequest should be signed.
        /// </summary>
        [Input("signRequest")]
        public Input<bool>? SignRequest { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// URL to send Authentication request to.
        /// </summary>
        [Input("ssoUrl", required: true)]
        public Input<string> SsoUrl { get; set; } = null!;

        public InboundSamlConfigIdpConfigGetArgs()
        {
        }
    }
}
