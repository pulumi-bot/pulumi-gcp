// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    public static class GetBillingAccount
    {
        /// <summary>
        /// Use this data source to get information about a Google Billing Account.
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var acct = Output.Create(Gcp.Organizations.GetBillingAccount.InvokeAsync(new Gcp.Organizations.GetBillingAccountArgs
        ///         {
        ///             DisplayName = "My Billing Account",
        ///             Open = true,
        ///         }));
        ///         var myProject = new Gcp.Organizations.Project("myProject", new Gcp.Organizations.ProjectArgs
        ///         {
        ///             ProjectId = "your-project-id",
        ///             OrgId = "1234567",
        ///             BillingAccount = acct.Apply(acct =&gt; acct.Id),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Task<GetBillingAccountResult> InvokeAsync(GetBillingAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBillingAccountResult>("gcp:organizations/getBillingAccount:getBillingAccount", args ?? new GetBillingAccountArgs(), options.WithVersion());

        public static Output<GetBillingAccountResult> Invoke(GetBillingAccountOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetBillingAccountOutputArgs();
            return Pulumi.Output.All(
                args.BillingAccount.Box(),
                args.DisplayName.Box(),
                args.Open.Box()
            ).Apply(a => {
                    var args = new GetBillingAccountArgs();
                    a[0].Set(args, nameof(args.BillingAccount));
                    a[1].Set(args, nameof(args.DisplayName));
                    a[2].Set(args, nameof(args.Open));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetBillingAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        /// </summary>
        [Input("billingAccount")]
        public string? BillingAccount { get; set; }

        /// <summary>
        /// The display name of the billing account.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// `true` if the billing account is open, `false` if the billing account is closed.
        /// </summary>
        [Input("open")]
        public bool? Open { get; set; }

        public GetBillingAccountArgs()
        {
        }
    }

    public sealed class GetBillingAccountOutputArgs
    {
        /// <summary>
        /// The name of the billing account in the form `{billing_account_id}` or `billingAccounts/{billing_account_id}`.
        /// </summary>
        [Input("billingAccount")]
        public Input<string>? BillingAccount { get; set; }

        /// <summary>
        /// The display name of the billing account.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// `true` if the billing account is open, `false` if the billing account is closed.
        /// </summary>
        [Input("open")]
        public Input<bool>? Open { get; set; }

        public GetBillingAccountOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBillingAccountResult
    {
        public readonly string? BillingAccount;
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource name of the billing account in the form `billingAccounts/{billing_account_id}`.
        /// </summary>
        public readonly string Name;
        public readonly bool Open;
        /// <summary>
        /// The IDs of any projects associated with the billing account.
        /// </summary>
        public readonly ImmutableArray<string> ProjectIds;

        [OutputConstructor]
        private GetBillingAccountResult(
            string? billingAccount,

            string displayName,

            string id,

            string name,

            bool open,

            ImmutableArray<string> projectIds)
        {
            BillingAccount = billingAccount;
            DisplayName = displayName;
            Id = id;
            Name = name;
            Open = open;
            ProjectIds = projectIds;
        }
    }
}
