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
    /// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the google_project_iam set of resources.
    /// 
    /// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
    /// 
    /// * `gcp.serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
    /// * `gcp.serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
    /// * `gcp.serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
    /// 
    /// &gt; **Note:** `gcp.serviceAccount.IAMPolicy` **cannot** be used in conjunction with `gcp.serviceAccount.IAMBinding` and `gcp.serviceAccount.IAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.serviceAccount.IAMBinding` resources **can be** used in conjunction with `gcp.serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_service\_account\_iam\_policy
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
    ///                     Role = "roles/iam.serviceAccountUser",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var sa = new Gcp.ServiceAccount.Account("sa", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-service-account",
    ///             DisplayName = "A service account that only Jane can interact with",
    ///         });
    ///         var admin_account_iam = new Gcp.ServiceAccount.IAMPolicy("admin-account-iam", new Gcp.ServiceAccount.IAMPolicyArgs
    ///         {
    ///             ServiceAccountId = sa.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_service\_account\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sa = new Gcp.ServiceAccount.Account("sa", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-service-account",
    ///             DisplayName = "A service account that only Jane can use",
    ///         });
    ///         var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new Gcp.ServiceAccount.IAMBindingArgs
    ///         {
    ///             ServiceAccountId = sa.Name,
    ///             Role = "roles/iam.serviceAccountUser",
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
    ///         var sa = new Gcp.ServiceAccount.Account("sa", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-service-account",
    ///             DisplayName = "A service account that only Jane can use",
    ///         });
    ///         var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new Gcp.ServiceAccount.IAMBindingArgs
    ///         {
    ///             Condition = new Gcp.ServiceAccount.Inputs.IAMBindingConditionArgs
    ///             {
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 Title = "expires_after_2019_12_31",
    ///             },
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Role = "roles/iam.serviceAccountUser",
    ///             ServiceAccountId = sa.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_service\_account\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = Output.Create(Gcp.Compute.GetDefaultServiceAccount.InvokeAsync());
    ///         var sa = new Gcp.ServiceAccount.Account("sa", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-service-account",
    ///             DisplayName = "A service account that Jane can use",
    ///         });
    ///         var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new Gcp.ServiceAccount.IAMMemberArgs
    ///         {
    ///             ServiceAccountId = sa.Name,
    ///             Role = "roles/iam.serviceAccountUser",
    ///             Member = "user:jane@example.com",
    ///         });
    ///         // Allow SA service account use the default GCE account
    ///         var gce_default_account_iam = new Gcp.ServiceAccount.IAMMember("gce-default-account-iam", new Gcp.ServiceAccount.IAMMemberArgs
    ///         {
    ///             ServiceAccountId = @default.Apply(@default =&gt; @default.Name),
    ///             Role = "roles/iam.serviceAccountUser",
    ///             Member = sa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
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
    ///         var sa = new Gcp.ServiceAccount.Account("sa", new Gcp.ServiceAccount.AccountArgs
    ///         {
    ///             AccountId = "my-service-account",
    ///             DisplayName = "A service account that Jane can use",
    ///         });
    ///         var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new Gcp.ServiceAccount.IAMMemberArgs
    ///         {
    ///             Condition = new Gcp.ServiceAccount.Inputs.IAMMemberConditionArgs
    ///             {
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 Title = "expires_after_2019_12_31",
    ///             },
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/iam.serviceAccountUser",
    ///             ServiceAccountId = sa.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).
    /// 
    /// ```sh
    ///  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions
    /// 
    /// ```sh
    ///  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser expires_after_2019_12_31"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:serviceAccount/iAMBinding:IAMBinding admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
    /// ```
    /// </summary>
    [GcpResourceType("gcp:serviceAccount/iAMBinding:IAMBinding")]
    public partial class IAMBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.IAMBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the service account IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a IAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMBinding(string name, IAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:serviceAccount/iAMBinding:IAMBinding", name, args ?? new IAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMBinding(string name, Input<string> id, IAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:serviceAccount/iAMBinding:IAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMBinding Get(string name, Input<string> id, IAMBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMBinding(name, id, state, options);
        }
    }

    public sealed class IAMBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public IAMBindingArgs()
        {
        }
    }

    public sealed class IAMBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the service account IAM policy.
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
        /// The role that should be applied. Only one
        /// `gcp.serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public IAMBindingState()
        {
        }
    }
}
