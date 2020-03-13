// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static partial class Invokes
    {
        [Obsolete("Use GetKMSSecretCiphertext.InvokeAsync() instead")]
        public static Task<GetKMSSecretCiphertextResult> GetKMSSecretCiphertext(GetKMSSecretCiphertextArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSSecretCiphertextResult>("gcp:kms/getKMSSecretCiphertext:getKMSSecretCiphertext", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetKMSSecretCiphertext
    {
        public static Task<GetKMSSecretCiphertextResult> InvokeAsync(GetKMSSecretCiphertextArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSSecretCiphertextResult>("gcp:kms/getKMSSecretCiphertext:getKMSSecretCiphertext", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetKMSSecretCiphertextArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the CryptoKey that will be used to
        /// encrypt the provided plaintext. This is represented by the format
        /// `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public string CryptoKey { get; set; } = null!;

        /// <summary>
        /// The plaintext to be encrypted
        /// </summary>
        [Input("plaintext", required: true)]
        public string Plaintext { get; set; } = null!;

        public GetKMSSecretCiphertextArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKMSSecretCiphertextResult
    {
        /// <summary>
        /// Contains the result of encrypting the provided plaintext, encoded in base64.
        /// </summary>
        public readonly string Ciphertext;
        public readonly string CryptoKey;
        public readonly string Plaintext;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKMSSecretCiphertextResult(
            string ciphertext,
            string cryptoKey,
            string plaintext,
            string id)
        {
            Ciphertext = ciphertext;
            CryptoKey = cryptoKey;
            Plaintext = plaintext;
            Id = id;
        }
    }
}
