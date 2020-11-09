// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Billing
{
    /// <summary>
    /// Allows creation and management of a single member for a single binding within
    /// the IAM policy for an existing Google Cloud Platform Billing Account.
    /// 
    /// &gt; **Note:** This resource __must not__ be used in conjunction with
    ///    `gcp.billing.AccountIamBinding` for the __same role__ or they will fight over
    ///    what your policy should be.
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `billing_account_id`, role, and member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:billing/accountIamMember:AccountIamMember binding "your-billing-account-id roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM member with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class AccountIamMember : Pulumi.CustomResource
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Output("billingAccountId")]
        public Output<string> BillingAccountId { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.AccountIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the billing account's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AccountIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountIamMember(string name, AccountIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamMember:AccountIamMember", name, args ?? new AccountIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountIamMember(string name, Input<string> id, AccountIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamMember:AccountIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountIamMember Get(string name, Input<string> id, AccountIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountIamMember(name, id, state, options);
        }
    }

    public sealed class AccountIamMemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.AccountIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AccountIamMemberArgs()
        {
        }
    }

    public sealed class AccountIamMemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId")]
        public Input<string>? BillingAccountId { get; set; }

        [Input("condition")]
        public Input<Inputs.AccountIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the billing account's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public AccountIamMemberState()
        {
        }
    }
}
