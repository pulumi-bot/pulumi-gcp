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
        /// This data source allows you to use data encrypted with Google Cloud KMS
        /// within your resource definitions.
        /// 
        /// For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
        /// 
        /// &gt; **NOTE**: Using this data provider will allow you to conceal secret data within your
        /// resource definitions, but it does not take care of protecting that data in the
        /// logging output, plan output, or state output.  Please take care to secure your secret
        /// data outside of resource definitions.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_secret.html.markdown.
        /// </summary>
        public static Task<GetKMSSecretResult> GetKMSSecret(GetKMSSecretArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSSecretResult>("gcp:kms/getKMSSecret:getKMSSecret", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetKMSSecretArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ciphertext to be decrypted, encoded in base64
        /// </summary>
        [Input("ciphertext", required: true)]
        public string Ciphertext { get; set; } = null!;

        /// <summary>
        /// The id of the CryptoKey that will be used to
        /// decrypt the provided ciphertext. This is represented by the format
        /// `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public string CryptoKey { get; set; } = null!;

        public GetKMSSecretArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetKMSSecretResult
    {
        public readonly string Ciphertext;
        public readonly string CryptoKey;
        /// <summary>
        /// Contains the result of decrypting the provided ciphertext.
        /// </summary>
        public readonly string Plaintext;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetKMSSecretResult(
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
