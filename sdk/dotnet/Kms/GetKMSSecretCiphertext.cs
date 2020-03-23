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
        /// <summary>
        /// !&gt; **Warning:** This data source is deprecated. Use the `gcp.kms.SecretCiphertext` **resource** instead.
        /// 
        /// This data source allows you to encrypt data with Google Cloud KMS and use the
        /// ciphertext within your resource definitions.
        /// 
        /// For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
        /// 
        /// &gt; **NOTE**: Using this data source will allow you to conceal secret data within your
        /// resource definitions, but it does not take care of protecting that data in the
        /// logging output, plan output, or state output.  Please take care to secure your secret
        /// data outside of resource definitions.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_secret_ciphertext.html.markdown.
        /// </summary>
        [Obsolete("Use GetKMSSecretCiphertext.InvokeAsync() instead")]
        public static Task<GetKMSSecretCiphertextResult> GetKMSSecretCiphertext(GetKMSSecretCiphertextArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSSecretCiphertextResult>("gcp:kms/getKMSSecretCiphertext:getKMSSecretCiphertext", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetKMSSecretCiphertext
    {
        /// <summary>
        /// !&gt; **Warning:** This data source is deprecated. Use the `gcp.kms.SecretCiphertext` **resource** instead.
        /// 
        /// This data source allows you to encrypt data with Google Cloud KMS and use the
        /// ciphertext within your resource definitions.
        /// 
        /// For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
        /// 
        /// &gt; **NOTE**: Using this data source will allow you to conceal secret data within your
        /// resource definitions, but it does not take care of protecting that data in the
        /// logging output, plan output, or state output.  Please take care to secure your secret
        /// data outside of resource definitions.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_kms_secret_ciphertext.html.markdown.
        /// </summary>
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
