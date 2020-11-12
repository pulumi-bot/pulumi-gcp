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
    /// Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
    /// 
    /// * `gcp.kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
    /// * `gcp.kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
    /// * `gcp.kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
    /// 
    /// &gt; **Note:** `gcp.kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `gcp.kms.CryptoKeyIAMBinding` and `gcp.kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `gcp.kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// With IAM Conditions:
    /// 
    /// With IAM Conditions:
    /// 
    /// With IAM Conditions:
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
    /// 
    /// These bindings can be imported using the `crypto_key_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `crypto_key_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key your-project-id/location-name/key-ring-name/key-name
    /// ```
    /// </summary>
    public partial class CryptoKeyIAMPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The crypto key ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        /// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        /// the provider's project setting will be used as a fallback.
        /// </summary>
        [Output("cryptoKeyId")]
        public Output<string> CryptoKeyId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the project's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a CryptoKeyIAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CryptoKeyIAMPolicy(string name, CryptoKeyIAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy", name, args ?? new CryptoKeyIAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CryptoKeyIAMPolicy(string name, Input<string> id, CryptoKeyIAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CryptoKeyIAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CryptoKeyIAMPolicy Get(string name, Input<string> id, CryptoKeyIAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new CryptoKeyIAMPolicy(name, id, state, options);
        }
    }

    public sealed class CryptoKeyIAMPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The crypto key ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        /// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        /// the provider's project setting will be used as a fallback.
        /// </summary>
        [Input("cryptoKeyId", required: true)]
        public Input<string> CryptoKeyId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public CryptoKeyIAMPolicyArgs()
        {
        }
    }

    public sealed class CryptoKeyIAMPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The crypto key ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        /// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        /// the provider's project setting will be used as a fallback.
        /// </summary>
        [Input("cryptoKeyId")]
        public Input<string>? CryptoKeyId { get; set; }

        /// <summary>
        /// (Computed) The etag of the project's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public CryptoKeyIAMPolicyState()
        {
        }
    }
}
