// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
    /// * `gcp.healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
    /// * `gcp.healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.Hl7StoreIamBinding` and `gcp.healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `hl7_v2_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `hl7_v2_store_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `hl7_v2_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember hl7_v2_store_iam your-project-id/location-name/dataset-name/hl7-v2-store-name
    /// ```
    /// </summary>
    public partial class Hl7StoreIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.Hl7StoreIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("hl7V2StoreId")]
        public Output<string> Hl7V2StoreId { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7StoreIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7StoreIamMember(string name, Hl7StoreIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember", name, args ?? new Hl7StoreIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7StoreIamMember(string name, Input<string> id, Hl7StoreIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamMember:Hl7StoreIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hl7StoreIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7StoreIamMember Get(string name, Input<string> id, Hl7StoreIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new Hl7StoreIamMember(name, id, state, options);
        }
    }

    public sealed class Hl7StoreIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.Hl7StoreIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId", required: true)]
        public Input<string> Hl7V2StoreId { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public Hl7StoreIamMemberArgs()
        {
        }
    }

    public sealed class Hl7StoreIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.Hl7StoreIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId")]
        public Input<string>? Hl7V2StoreId { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public Hl7StoreIamMemberState()
        {
        }
    }
}
