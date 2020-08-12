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
    /// Allows creation and management of a single binding within IAM policy for
    /// an existing Google Cloud Platform Billing Account.
    /// 
    /// &gt; **Note:** This resource __must not__ be used in conjunction with
    ///    `gcp.billing.AccountIamMember` for the __same role__ or they will fight over
    ///    what your policy should be.
    /// 
    /// &gt; **Note:** On create, this resource will overwrite members of any existing roles.
    ///     Use `pulumi import` and inspect the output to ensure
    ///     your existing members are preserved.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.Billing.AccountIamBinding("binding", new Gcp.Billing.AccountIamBindingArgs
    ///         {
    ///             BillingAccountId = "00AA00-000AAA-00AA0A",
    ///             Members = 
    ///             {
    ///                 "user:alice@gmail.com",
    ///             },
    ///             Role = "roles/billing.viewer",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AccountIamBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Output("billingAccountId")]
        public Output<string> BillingAccountId { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.AccountIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the billing account's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AccountIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountIamBinding(string name, AccountIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamBinding:AccountIamBinding", name, args ?? new AccountIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountIamBinding(string name, Input<string> id, AccountIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamBinding:AccountIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountIamBinding Get(string name, Input<string> id, AccountIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountIamBinding(name, id, state, options);
        }
    }

    public sealed class AccountIamBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.AccountIamBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AccountIamBindingArgs()
        {
        }
    }

    public sealed class AccountIamBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId")]
        public Input<string>? BillingAccountId { get; set; }

        [Input("condition")]
        public Input<Inputs.AccountIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the billing account's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public AccountIamBindingState()
        {
        }
    }
}
