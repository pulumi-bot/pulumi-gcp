// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
        /// and
        /// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).
        /// 
        /// A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_crypto_key_version.html.markdown.
        /// </summary>
        public static Task<GetKMSCryptoKeyVersionResult> GetKMSCryptoKeyVersion(GetKMSCryptoKeyVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSCryptoKeyVersionResult>("gcp:kms/getKMSCryptoKeyVersion:getKMSCryptoKeyVersion", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetKMSCryptoKeyVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The `self_link` of the Google Cloud Platform CryptoKey to which the key version belongs.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        [Input("publicKey")]
        public Input<Inputs.GetKMSCryptoKeyVersionPublicKeyArgs>? PublicKey { get; set; }

        /// <summary>
        /// The version number for this CryptoKeyVersion. Defaults to `1`.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetKMSCryptoKeyVersionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKMSCryptoKeyVersionResult
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        public readonly string Algorithm;
        public readonly string CryptoKey;
        /// <summary>
        /// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.
        /// </summary>
        public readonly string ProtectionLevel;
        /// <summary>
        /// If the enclosing CryptoKey has purpose `ASYMMETRIC_SIGN` or `ASYMMETRIC_DECRYPT`, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below.
        /// </summary>
        public readonly Outputs.GetKMSCryptoKeyVersionPublicKeyResult? PublicKey;
        /// <summary>
        /// The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.
        /// </summary>
        public readonly string State;
        public readonly int? Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKMSCryptoKeyVersionResult(
            string algorithm,
            string cryptoKey,
            string protectionLevel,
            Outputs.GetKMSCryptoKeyVersionPublicKeyResult? publicKey,
            string state,
            int? version,
            string id)
        {
            Algorithm = algorithm;
            CryptoKey = cryptoKey;
            ProtectionLevel = protectionLevel;
            PublicKey = publicKey;
            State = state;
            Version = version;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetKMSCryptoKeyVersionPublicKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.
        /// </summary>
        [Input("pem")]
        public Input<string>? Pem { get; set; }

        public GetKMSCryptoKeyVersionPublicKeyArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetKMSCryptoKeyVersionPublicKeyResult
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// The public key, encoded in PEM format. For more information, see the RFC 7468 sections for General Considerations and Textual Encoding of Subject Public Key Info.
        /// </summary>
        public readonly string Pem;

        [OutputConstructor]
        private GetKMSCryptoKeyVersionPublicKeyResult(
            string algorithm,
            string pem)
        {
            Algorithm = algorithm;
            Pem = pem;
        }
    }
    }
}
