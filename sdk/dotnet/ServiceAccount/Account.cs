// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceAccount
{
    /// <summary>
    /// Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
    /// 
    /// &gt; Creation of service accounts is eventually consistent, and that can lead to
    /// errors when you try to apply ACLs to service accounts immediately after
    /// creation.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account.html.markdown.
    /// </summary>
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// The account id that is used to generate the service
        /// account email address and a stable unique id. It is unique within a project,
        /// must be 6-30 characters long, and match the regular expression `a-z`
        /// to comply with RFC1035. Changing this forces a new service account to be created.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// A text description of the service account.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name for the service account.
        /// Can be updated without creating a new resource.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The e-mail address of the service account. This value
        /// should be referenced from any `gcp.organizations.getIAMPolicy` data sources
        /// that would grant the service account privileges.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified name of the service account.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that the service account will be created in.
        /// Defaults to the provider project configuration.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The unique id of the service account.
        /// </summary>
        [Output("uniqueId")]
        public Output<string> UniqueId { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("gcp:serviceAccount/account:Account", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("gcp:serviceAccount/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id that is used to generate the service
        /// account email address and a stable unique id. It is unique within a project,
        /// must be 6-30 characters long, and match the regular expression `a-z`
        /// to comply with RFC1035. Changing this forces a new service account to be created.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// A text description of the service account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the service account.
        /// Can be updated without creating a new resource.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of the project that the service account will be created in.
        /// Defaults to the provider project configuration.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public AccountArgs()
        {
        }
    }

    public sealed class AccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account id that is used to generate the service
        /// account email address and a stable unique id. It is unique within a project,
        /// must be 6-30 characters long, and match the regular expression `a-z`
        /// to comply with RFC1035. Changing this forces a new service account to be created.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// A text description of the service account.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the service account.
        /// Can be updated without creating a new resource.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The e-mail address of the service account. This value
        /// should be referenced from any `gcp.organizations.getIAMPolicy` data sources
        /// that would grant the service account privileges.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The fully-qualified name of the service account.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project that the service account will be created in.
        /// Defaults to the provider project configuration.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The unique id of the service account.
        /// </summary>
        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        public AccountState()
        {
        }
    }
}
