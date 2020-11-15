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
    /// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
    /// 
    /// * `gcp.kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
    /// * `gcp.kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
    /// * `gcp.kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
    /// 
    /// &gt; **Note:** `gcp.kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `gcp.kms.KeyRingIAMBinding` and `gcp.kms.KeyRingIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.kms.KeyRingIAMBinding` resources **can be** used in conjunction with `gcp.kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `key_ring_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `key_ring_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `key_ring_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding key_ring_iam your-project-id/location-name/key-ring-name
    /// ```
    /// </summary>
    public partial class KeyRingIAMBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.KeyRingIAMBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the key ring's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("keyRingId")]
        public Output<string> KeyRingId { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a KeyRingIAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRingIAMBinding(string name, KeyRingIAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, args ?? new KeyRingIAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyRingIAMBinding(string name, Input<string> id, KeyRingIAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeyRingIAMBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRingIAMBinding Get(string name, Input<string> id, KeyRingIAMBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new KeyRingIAMBinding(name, id, state, options);
        }
    }

    public sealed class KeyRingIAMBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.KeyRingIAMBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public KeyRingIAMBindingArgs()
        {
        }
    }

    public sealed class KeyRingIAMBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.KeyRingIAMBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the key ring's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId")]
        public Input<string>? KeyRingId { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public KeyRingIAMBindingState()
        {
        }
    }
}
