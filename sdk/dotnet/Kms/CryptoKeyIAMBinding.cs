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
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var keyring = new Gcp.Kms.KeyRing("keyring", new Gcp.Kms.KeyRingArgs
    ///         {
    ///             Location = "global",
    ///         });
    ///         var key = new Gcp.Kms.CryptoKey("key", new Gcp.Kms.CryptoKeyArgs
    ///         {
    ///             KeyRing = keyring.Id,
    ///             RotationPeriod = "100000s",
    ///         });
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var cryptoKey = new Gcp.Kms.CryptoKeyIAMPolicy("cryptoKey", new Gcp.Kms.CryptoKeyIAMPolicyArgs
    ///         {
    ///             CryptoKeyId = key.Id,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionArgs
    ///                     {
    ///                         Description = "Expiring at midnight of 2019-12-31",
    ///                         Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                         Title = "expires_after_2019_12_31",
    ///                     },
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                     Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///                 },
    ///             },
    ///         }));
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cryptoKey = new Gcp.Kms.CryptoKeyIAMBinding("cryptoKey", new Gcp.Kms.CryptoKeyIAMBindingArgs
    ///         {
    ///             CryptoKeyId = google_kms_crypto_key.Key.Id,
    ///             Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cryptoKey = new Gcp.Kms.CryptoKeyIAMBinding("cryptoKey", new Gcp.Kms.CryptoKeyIAMBindingArgs
    ///         {
    ///             CryptoKeyId = google_kms_crypto_key.Key.Id,
    ///             Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Condition = new Gcp.Kms.Inputs.CryptoKeyIAMBindingConditionArgs
    ///             {
    ///                 Title = "expires_after_2019_12_31",
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cryptoKey = new Gcp.Kms.CryptoKeyIAMMember("cryptoKey", new Gcp.Kms.CryptoKeyIAMMemberArgs
    ///         {
    ///             CryptoKeyId = google_kms_crypto_key.Key.Id,
    ///             Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///             Member = "user:jane@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var cryptoKey = new Gcp.Kms.CryptoKeyIAMMember("cryptoKey", new Gcp.Kms.CryptoKeyIAMMemberArgs
    ///         {
    ///             CryptoKeyId = google_kms_crypto_key.Key.Id,
    ///             Role = "roles/cloudkms.cryptoKeyEncrypter",
    ///             Member = "user:jane@example.com",
    ///             Condition = new Gcp.Kms.Inputs.CryptoKeyIAMMemberConditionArgs
    ///             {
    ///                 Title = "expires_after_2019_12_31",
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; first the resource in question and then the role.
    /// 
    /// These bindings can be imported using the `crypto_key_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `crypto_key_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding crypto_key your-project-id/location-name/key-ring-name/key-name
    /// ```
    /// </summary>
    [GcpResourceType("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding")]
    public partial class CryptoKeyIAMBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.CryptoKeyIAMBindingCondition?> Condition { get; private set; } = null!;

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

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a CryptoKeyIAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CryptoKeyIAMBinding(string name, CryptoKeyIAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding", name, args ?? new CryptoKeyIAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CryptoKeyIAMBinding(string name, Input<string> id, CryptoKeyIAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyIAMBinding:CryptoKeyIAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CryptoKeyIAMBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CryptoKeyIAMBinding Get(string name, Input<string> id, CryptoKeyIAMBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new CryptoKeyIAMBinding(name, id, state, options);
        }
    }

    public sealed class CryptoKeyIAMBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.CryptoKeyIAMBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The crypto key ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        /// `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        /// the provider's project setting will be used as a fallback.
        /// </summary>
        [Input("cryptoKeyId", required: true)]
        public Input<string> CryptoKeyId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public CryptoKeyIAMBindingArgs()
        {
        }
    }

    public sealed class CryptoKeyIAMBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.CryptoKeyIAMBindingConditionGetArgs>? Condition { get; set; }

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

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public CryptoKeyIAMBindingState()
        {
        }
    }
}
