// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    /// Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
    /// 
    /// 
    /// &gt; **NOTE**: Using this resource will allow you to conceal secret data within your
    /// resource definitions, but it does not take care of protecting that data in the
    /// logging output, plan output, or state output.  Please take care to secure your secret
    /// data outside of resource definitions.
    /// 
    /// 
    /// To get more information about SecretCiphertext, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys/encrypt)
    /// * How-to Guides
    ///     * [Encrypting and decrypting data with a symmetric key](https://cloud.google.com/kms/docs/encrypt-decrypt)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/kms_secret_ciphertext.html.markdown.
    /// </summary>
    public partial class SecretCiphertext : Pulumi.CustomResource
    {
        /// <summary>
        /// Contains the result of encrypting the provided plaintext, encoded in base64.
        /// </summary>
        [Output("ciphertext")]
        public Output<string> Ciphertext { get; private set; } = null!;

        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
        /// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
        /// </summary>
        [Output("cryptoKey")]
        public Output<string> CryptoKey { get; private set; } = null!;

        /// <summary>
        /// The plaintext to be encrypted.
        /// </summary>
        [Output("plaintext")]
        public Output<string> Plaintext { get; private set; } = null!;


        /// <summary>
        /// Create a SecretCiphertext resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretCiphertext(string name, SecretCiphertextArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/secretCiphertext:SecretCiphertext", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SecretCiphertext(string name, Input<string> id, SecretCiphertextState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/secretCiphertext:SecretCiphertext", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretCiphertext resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretCiphertext Get(string name, Input<string> id, SecretCiphertextState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretCiphertext(name, id, state, options);
        }
    }

    public sealed class SecretCiphertextArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
        /// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        /// <summary>
        /// The plaintext to be encrypted.
        /// </summary>
        [Input("plaintext", required: true)]
        public Input<string> Plaintext { get; set; } = null!;

        public SecretCiphertextArgs()
        {
        }
    }

    public sealed class SecretCiphertextState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains the result of encrypting the provided plaintext, encoded in base64.
        /// </summary>
        [Input("ciphertext")]
        public Input<string>? Ciphertext { get; set; }

        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext. Format:
        /// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}''
        /// </summary>
        [Input("cryptoKey")]
        public Input<string>? CryptoKey { get; set; }

        /// <summary>
        /// The plaintext to be encrypted.
        /// </summary>
        [Input("plaintext")]
        public Input<string>? Plaintext { get; set; }

        public SecretCiphertextState()
        {
        }
    }
}
