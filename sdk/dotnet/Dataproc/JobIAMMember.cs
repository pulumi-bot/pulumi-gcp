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
    /// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
    /// * `gcp.dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
    /// * `gcp.dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
    /// 
    /// &gt; **Note:** `gcp.dataproc.JobIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.JobIAMBinding` and `gcp.dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `gcp.dataproc.JobIAMPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.dataproc.JobIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
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
    ///         var editor = new Gcp.Dataproc.JobIAMPolicy("editor", new Gcp.Dataproc.JobIAMPolicyArgs
    ///         {
    ///             Project = "your-project",
    ///             Region = "your-region",
    ///             JobId = "your-dataproc-job",
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
    ///         var editor = new Gcp.Dataproc.JobIAMBinding("editor", new Gcp.Dataproc.JobIAMBindingArgs
    ///         {
    ///             JobId = "your-dataproc-job",
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
    ///         var editor = new Gcp.Dataproc.JobIAMMember("editor", new Gcp.Dataproc.JobIAMMemberArgs
    ///         {
    ///             JobId = "your-dataproc-job",
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
    public partial class JobIAMMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.JobIAMMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the jobs's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a JobIAMMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobIAMMember(string name, JobIAMMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/jobIAMMember:JobIAMMember", name, args ?? new JobIAMMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobIAMMember(string name, Input<string> id, JobIAMMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/jobIAMMember:JobIAMMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing JobIAMMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobIAMMember Get(string name, Input<string> id, JobIAMMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new JobIAMMember(name, id, state, options);
        }
    }

    public sealed class JobIAMMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.JobIAMMemberConditionArgs>? Condition { get; set; }

        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public JobIAMMemberArgs()
        {
        }
    }

    public sealed class JobIAMMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.JobIAMMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the jobs's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataproc.JobIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public JobIAMMemberState()
        {
        }
    }
}
