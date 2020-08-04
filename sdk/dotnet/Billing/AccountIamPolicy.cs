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
    /// Allows management of the entire IAM policy for an existing Google Cloud Platform Billing Account.
    /// 
    /// &gt; **Warning:** Billing accounts have a default user that can be **overwritten**
    /// by use of this resource. The safest alternative is to use multiple `gcp.billing.AccountIamBinding`
    ///    resources. If you do use this resource, the best way to be sure that you are
    ///    not making dangerous changes is to start by importing your existing policy,
    ///    and examining the diff very closely.
    /// 
    /// &gt; **Note:** This resource __must not__ be used in conjunction with
    ///    `gcp.billing.AccountIamMember` or `gcp.billing.AccountIamBinding`
    ///    or they will fight over what your policy should be.
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
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Role = "roles/billing.viewer",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.Billing.AccountIamPolicy("policy", new Gcp.Billing.AccountIamPolicyArgs
    ///         {
    ///             BillingAccountId = "00AA00-000AAA-00AA0A",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AccountIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Output("billingAccountId")]
        public Output<string> BillingAccountId { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the billing account. This policy overrides any existing
        /// policy applied to the billing account.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a AccountIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountIamPolicy(string name, AccountIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamPolicy:AccountIamPolicy", name, args ?? new AccountIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountIamPolicy(string name, Input<string> id, AccountIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:billing/accountIamPolicy:AccountIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountIamPolicy Get(string name, Input<string> id, AccountIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountIamPolicy(name, id, state, options);
        }
    }

    public sealed class AccountIamPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the billing account. This policy overrides any existing
        /// policy applied to the billing account.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public AccountIamPolicyArgs()
        {
        }
    }

    public sealed class AccountIamPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The billing account id.
        /// </summary>
        [Input("billingAccountId")]
        public Input<string>? BillingAccountId { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the billing account. This policy overrides any existing
        /// policy applied to the billing account.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public AccountIamPolicyState()
        {
        }
    }
}
