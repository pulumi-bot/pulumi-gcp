// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static class GetKMSSecret
    {
        /// <summary>
        /// This data source allows you to use data encrypted with Google Cloud KMS
        /// within your resource definitions.
        /// 
        /// For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
        /// 
        /// &gt; **NOTE:** Using this data provider will allow you to conceal secret data within your
        /// resource definitions, but it does not take care of protecting that data in the
        /// logging output, plan output, or state output.  Please take care to secure your secret
        /// data outside of resource definitions.
        /// </summary>
        public static Task<GetKMSSecretResult> InvokeAsync(GetKMSSecretArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSSecretResult>("gcp:kms/getKMSSecret:getKMSSecret", args ?? new GetKMSSecretArgs(), options.WithVersion());

        public static Output<GetKMSSecretResult> Invoke(GetKMSSecretOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.AdditionalAuthenticatedData.Box(),
                args.Ciphertext.Box(),
                args.CryptoKey.Box()
            ).Apply(a => {
                    var args = new GetKMSSecretArgs();
                    a[0].Set(args, nameof(args.AdditionalAuthenticatedData));
                    a[1].Set(args, nameof(args.Ciphertext));
                    a[2].Set(args, nameof(args.CryptoKey));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetKMSSecretArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [additional authenticated data](https://cloud.google.com/kms/docs/additional-authenticated-data) used for integrity checks during encryption and decryption.
        /// </summary>
        [Input("additionalAuthenticatedData")]
        public string? AdditionalAuthenticatedData { get; set; }

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

    public sealed class GetKMSSecretOutputArgs
    {
        /// <summary>
        /// The [additional authenticated data](https://cloud.google.com/kms/docs/additional-authenticated-data) used for integrity checks during encryption and decryption.
        /// </summary>
        [Input("additionalAuthenticatedData")]
        public Input<string>? AdditionalAuthenticatedData { get; set; }

        /// <summary>
        /// The ciphertext to be decrypted, encoded in base64
        /// </summary>
        [Input("ciphertext", required: true)]
        public Input<string> Ciphertext { get; set; } = null!;

        /// <summary>
        /// The id of the CryptoKey that will be used to
        /// decrypt the provided ciphertext. This is represented by the format
        /// `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        public GetKMSSecretOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKMSSecretResult
    {
        public readonly string? AdditionalAuthenticatedData;
        public readonly string Ciphertext;
        public readonly string CryptoKey;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Contains the result of decrypting the provided ciphertext.
        /// </summary>
        public readonly string Plaintext;

        [OutputConstructor]
        private GetKMSSecretResult(
            string? additionalAuthenticatedData,

            string ciphertext,

            string cryptoKey,

            string id,

            string plaintext)
        {
            AdditionalAuthenticatedData = additionalAuthenticatedData;
            Ciphertext = ciphertext;
            CryptoKey = cryptoKey;
            Id = id;
            Plaintext = plaintext;
        }
    }
}
