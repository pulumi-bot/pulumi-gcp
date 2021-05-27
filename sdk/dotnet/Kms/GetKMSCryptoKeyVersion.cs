// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static class GetKMSCryptoKeyVersion
    {
        /// <summary>
        /// Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
        /// [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
        /// and
        /// [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).
        /// 
        /// A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var myKeyRing = Output.Create(Gcp.Kms.GetKMSKeyRing.InvokeAsync(new Gcp.Kms.GetKMSKeyRingArgs
        ///         {
        ///             Name = "my-key-ring",
        ///             Location = "us-central1",
        ///         }));
        ///         var myCryptoKey = myKeyRing.Apply(myKeyRing =&gt; Output.Create(Gcp.Kms.GetKMSCryptoKey.InvokeAsync(new Gcp.Kms.GetKMSCryptoKeyArgs
        ///         {
        ///             Name = "my-crypto-key",
        ///             KeyRing = myKeyRing.SelfLink,
        ///         })));
        ///         var myCryptoKeyVersion = Output.Create(Gcp.Kms.GetKMSCryptoKeyVersion.InvokeAsync(new Gcp.Kms.GetKMSCryptoKeyVersionArgs
        ///         {
        ///             CryptoKey = data.Google_kms_key.My_key.Self_link,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKMSCryptoKeyVersionResult> InvokeAsync(GetKMSCryptoKeyVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKMSCryptoKeyVersionResult>("gcp:kms/getKMSCryptoKeyVersion:getKMSCryptoKeyVersion", args ?? new GetKMSCryptoKeyVersionArgs(), options.WithVersion());

        public static Output<GetKMSCryptoKeyVersionResult> Apply(GetKMSCryptoKeyVersionApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.CryptoKey.Box(),
                args.Version.Box()
            ).Apply(a => {
                    var args = new GetKMSCryptoKeyVersionArgs();
                    a[0].Set(args, nameof(args.CryptoKey));
                    a[1].Set(args, nameof(args.Version));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetKMSCryptoKeyVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The `self_link` of the Google Cloud Platform CryptoKey to which the key version belongs.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public string CryptoKey { get; set; } = null!;

        /// <summary>
        /// The version number for this CryptoKeyVersion. Defaults to `1`.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetKMSCryptoKeyVersionArgs()
        {
        }
    }

    public sealed class GetKMSCryptoKeyVersionApplyArgs
    {
        /// <summary>
        /// The `self_link` of the Google Cloud Platform CryptoKey to which the key version belongs.
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        /// <summary>
        /// The version number for this CryptoKeyVersion. Defaults to `1`.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetKMSCryptoKeyVersionApplyArgs()
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
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.
        /// </summary>
        public readonly string ProtectionLevel;
        /// <summary>
        /// If the enclosing CryptoKey has purpose `ASYMMETRIC_SIGN` or `ASYMMETRIC_DECRYPT`, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKMSCryptoKeyVersionPublicKeyResult> PublicKeys;
        /// <summary>
        /// The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.
        /// </summary>
        public readonly string State;
        public readonly int? Version;

        [OutputConstructor]
        private GetKMSCryptoKeyVersionResult(
            string algorithm,

            string cryptoKey,

            string id,

            string protectionLevel,

            ImmutableArray<Outputs.GetKMSCryptoKeyVersionPublicKeyResult> publicKeys,

            string state,

            int? version)
        {
            Algorithm = algorithm;
            CryptoKey = cryptoKey;
            Id = id;
            ProtectionLevel = protectionLevel;
            PublicKeys = publicKeys;
            State = state;
            Version = version;
        }
    }
}
