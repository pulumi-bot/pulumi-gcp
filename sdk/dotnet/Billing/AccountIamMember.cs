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
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_billing_account_iam_member.html.markdown.
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
            : base("gcp:billing/accountIamMember:AccountIamMember", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    namespace Inputs
    {

    public sealed class AccountIamMemberConditionArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AccountIamMemberConditionArgs()
        {
        }
    }

    public sealed class AccountIamMemberConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AccountIamMemberConditionGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AccountIamMemberCondition
    {
        public readonly string? Description;
        public readonly string Expression;
        public readonly string Title;

        [OutputConstructor]
        private AccountIamMemberCondition(
            string? description,
            string expression,
            string title)
        {
            Description = description;
            Expression = expression;
            Title = title;
        }
    }
    }
}
