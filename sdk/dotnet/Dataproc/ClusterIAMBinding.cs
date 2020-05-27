// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc
{
    /// <summary>
    /// Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
    /// * `gcp.dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
    /// * `gcp.dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.
    /// 
    /// &gt; **Note:** `gcp.dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.ClusterIAMBinding` and `gcp.dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `gcp.dataproc.ClusterIAMPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_pubsub\_subscription\_iam\_policy
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
    ///                     { "role", "roles/editor" },
    ///                     { "members", 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     } },
    ///                 },
    ///             },
    ///         }));
    ///         var editor = new Gcp.Dataproc.ClusterIAMPolicy("editor", new Gcp.Dataproc.ClusterIAMPolicyArgs
    ///         {
    ///             Project = "your-project",
    ///             Region = "your-region",
    ///             Cluster = "your-dataproc-cluster",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// ## google\_pubsub\_subscription\_iam\_binding
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
    ///         var editor = new Gcp.Dataproc.ClusterIAMBinding("editor", new Gcp.Dataproc.ClusterIAMBindingArgs
    ///         {
    ///             Cluster = "your-dataproc-cluster",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// ## google\_pubsub\_subscription\_iam\_member
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
    ///         var editor = new Gcp.Dataproc.ClusterIAMMember("editor", new Gcp.Dataproc.ClusterIAMMemberArgs
    ///         {
    ///             Cluster = "your-dataproc-cluster",
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// </summary>
    public partial class ClusterIAMBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// The name or relative resource id of the cluster to manage IAM policies for.
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.ClusterIAMBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the clusters's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The project in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterIAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterIAMBinding(string name, ClusterIAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, args ?? new ClusterIAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterIAMBinding(string name, Input<string> id, ClusterIAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/clusterIAMBinding:ClusterIAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClusterIAMBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterIAMBinding Get(string name, Input<string> id, ClusterIAMBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterIAMBinding(name, id, state, options);
        }
    }

    public sealed class ClusterIAMBindingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or relative resource id of the cluster to manage IAM policies for.
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.ClusterIAMBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The project in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public ClusterIAMBindingArgs()
        {
        }
    }

    public sealed class ClusterIAMBindingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or relative resource id of the cluster to manage IAM policies for.
        /// </summary>
        [Input("cluster")]
        public Input<string>? Cluster { get; set; }

        [Input("condition")]
        public Input<Inputs.ClusterIAMBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the clusters's IAM policy.
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
        /// The project in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the cluster belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.ClusterIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public ClusterIAMBindingState()
        {
        }
    }
}
