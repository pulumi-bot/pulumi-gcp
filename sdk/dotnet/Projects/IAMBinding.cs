// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    /// <summary>
    /// Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
    /// 
    /// * `gcp.projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
    /// * `gcp.projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
    /// * `gcp.projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
    /// * `gcp.projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.
    /// 
    /// &gt; **Note:** `gcp.projects.IAMPolicy` **cannot** be used in conjunction with `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.projects.IAMBinding` resources **can be** used in conjunction with `gcp.projects.IAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_project\_iam\_policy
    /// 
    /// &gt; **Be careful!** You can accidentally lock yourself out of your project
    ///    using this resource. Deleting a `gcp.projects.IAMPolicy` removes access
    ///    from anyone without organization-level access to the project. Proceed with caution.
    ///    It's not recommended to use `gcp.projects.IAMPolicy` with your provider project
    ///    to avoid locking yourself out, and it should generally only be used with projects
    ///    fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
    ///    applying the change.
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
    ///                     Role = "roles/editor",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var project = new Gcp.Projects.IAMPolicy("project", new Gcp.Projects.IAMPolicyArgs
    ///         {
    ///             Project = "your-project-id",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions):
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
    ///                     Role = "roles/editor",
    ///                 },
    ///             },
    ///         }));
    ///         var project = new Gcp.Projects.IAMPolicy("project", new Gcp.Projects.IAMPolicyArgs
    ///         {
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///             Project = "your-project-id",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_project\_iam\_binding
    /// 
    /// &gt; **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your project.
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = new Gcp.Projects.IAMBinding("project", new Gcp.Projects.IAMBindingArgs
    ///         {
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Project = "your-project-id",
    ///             Role = "roles/editor",
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
    ///         var project = new Gcp.Projects.IAMBinding("project", new Gcp.Projects.IAMBindingArgs
    ///         {
    ///             Condition = new Gcp.Projects.Inputs.IAMBindingConditionArgs
    ///             {
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 Title = "expires_after_2019_12_31",
    ///             },
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Project = "your-project-id",
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_project\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = new Gcp.Projects.IAMMember("project", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Member = "user:jane@example.com",
    ///             Project = "your-project-id",
    ///             Role = "roles/editor",
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
    ///         var project = new Gcp.Projects.IAMMember("project", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Condition = new Gcp.Projects.Inputs.IAMMemberConditionArgs
    ///             {
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 Title = "expires_after_2019_12_31",
    ///             },
    ///             Member = "user:jane@example.com",
    ///             Project = "your-project-id",
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_project\_iam\_audit\_config
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project = new Gcp.Projects.IAMAuditConfig("project", new Gcp.Projects.IAMAuditConfigArgs
    ///         {
    ///             AuditLogConfigs = 
    ///             {
    ///                 new Gcp.Projects.Inputs.IAMAuditConfigAuditLogConfigArgs
    ///                 {
    ///                     LogType = "ADMIN_READ",
    ///                 },
    ///                 new Gcp.Projects.Inputs.IAMAuditConfigAuditLogConfigArgs
    ///                 {
    ///                     ExemptedMembers = 
    ///                     {
    ///                         "user:joebloggs@hashicorp.com",
    ///                     },
    ///                     LogType = "DATA_READ",
    ///                 },
    ///             },
    ///             Project = "your-project-id",
    ///             Service = "allServices",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class IAMBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.IAMBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the project's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        /// Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
        /// will not be inferred from the provider.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a IAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMBinding(string name, IAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/iAMBinding:IAMBinding", name, args ?? new IAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMBinding(string name, Input<string> id, IAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/iAMBinding:IAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        /// Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
        /// will not be inferred from the provider.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

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
        /// The project ID. If not specified for `gcp.projects.IAMBinding`, `gcp.projects.IAMMember`, or `gcp.projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        /// Required for `gcp.projects.IAMPolicy` - you must explicitly set the project, and it
        /// will not be inferred from the provider.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public IAMBindingState()
        {
        }
    }
}
