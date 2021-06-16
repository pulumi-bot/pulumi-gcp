// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority.Inputs
{

    public sealed class CertificateCertificateDescriptionPublicKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. A public key. When this is specified in a request, the padding and encoding can be any of the options described by the respective 'KeyType' value. When this is generated by the service, it will always be an RFC 5280 SubjectPublicKeyInfo structure containing an algorithm identifier and a key. A base64-encoded string.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Types of public keys that are supported. At a minimum, we support RSA and ECDSA, for the key sizes or curves listed: https://cloud.google.com/kms/docs/algorithms#asymmetric_signing_algorithms
        /// Possible values are `KEY_TYPE_UNSPECIFIED`, `PEM_RSA_KEY`, and `PEM_EC_KEY`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public CertificateCertificateDescriptionPublicKeyArgs()
        {
        }
    }
}
