// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.RuntimeConfig
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Runtime Configurator Config. Each of these resources serves a different use case:
    /// 
    /// * `gcp.runtimeconfig.ConfigIamPolicy`: Authoritative. Sets the IAM policy for the config and replaces any existing policy already attached.
    /// * `gcp.runtimeconfig.ConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the config are preserved.
    /// * `gcp.runtimeconfig.ConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the config are preserved.
    /// 
    /// &gt; **Note:** `gcp.runtimeconfig.ConfigIamPolicy` **cannot** be used in conjunction with `gcp.runtimeconfig.ConfigIamBinding` and `gcp.runtimeconfig.ConfigIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.runtimeconfig.ConfigIamBinding` resources **can be** used in conjunction with `gcp.runtimeconfig.ConfigIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// 
    /// 
    /// ## google\_runtimeconfig\_config\_iam\_policy
    /// {{% example %}}
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
    ///             Binding = 
    ///             {
    ///                 
    ///                 {
    ///                     { "role", "roles/viewer" },
    ///                     { "members", 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     } },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.RuntimeConfig.ConfigIamPolicy("policy", new Gcp.RuntimeConfig.ConfigIamPolicyArgs
    ///         {
    ///             Project = google_runtimeconfig_config.Config.Project,
    ///             Config = google_runtimeconfig_config.Config.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// ## google\_runtimeconfig\_config\_iam\_binding
    /// {{% example %}}
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.RuntimeConfig.ConfigIamBinding("binding", new Gcp.RuntimeConfig.ConfigIamBindingArgs
    ///         {
    ///             Project = google_runtimeconfig_config.Config.Project,
    ///             Config = google_runtimeconfig_config.Config.Name,
    ///             Role = "roles/viewer",
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
    /// {{% /example %}}
    /// ## google\_runtimeconfig\_config\_iam\_member
    /// {{% example %}}
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.RuntimeConfig.ConfigIamMember("member", new Gcp.RuntimeConfig.ConfigIamMemberArgs
    ///         {
    ///             Project = google_runtimeconfig_config.Config.Project,
    ///             Config = google_runtimeconfig_config.Config.Name,
    ///             Role = "roles/viewer",
    ///             Member = "user:jane@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// </summary>
    public partial class ConfigIamPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigIamPolicy(string name, ConfigIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, args ?? new ConfigIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigIamPolicy(string name, Input<string> id, ConfigIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/configIamPolicy:ConfigIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigIamPolicy Get(string name, Input<string> id, ConfigIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigIamPolicy(name, id, state, options);
        }
    }

    public sealed class ConfigIamPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConfigIamPolicyArgs()
        {
        }
    }

    public sealed class ConfigIamPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ConfigIamPolicyState()
        {
        }
    }
}
