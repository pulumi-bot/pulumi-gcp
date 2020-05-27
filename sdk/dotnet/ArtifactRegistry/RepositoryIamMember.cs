// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ArtifactRegistry
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Artifact Registry Repository. Each of these resources serves a different use case:
    /// 
    /// * `gcp.artifactregistry.RepositoryIamPolicy`: Authoritative. Sets the IAM policy for the repository and replaces any existing policy already attached.
    /// * `gcp.artifactregistry.RepositoryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the repository are preserved.
    /// * `gcp.artifactregistry.RepositoryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the repository are preserved.
    /// 
    /// &gt; **Note:** `gcp.artifactregistry.RepositoryIamPolicy` **cannot** be used in conjunction with `gcp.artifactregistry.RepositoryIamBinding` and `gcp.artifactregistry.RepositoryIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.artifactregistry.RepositoryIamBinding` resources **can be** used in conjunction with `gcp.artifactregistry.RepositoryIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// 
    /// ## google\_artifact\_registry\_repository\_iam\_policy
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
    ///         var policy = new Gcp.ArtifactRegistry.RepositoryIamPolicy("policy", new Gcp.ArtifactRegistry.RepositoryIamPolicyArgs
    ///         {
    ///             Project = google_artifact_registry_repository.My-repo.Project,
    ///             Location = google_artifact_registry_repository.My-repo.Location,
    ///             Repository = google_artifact_registry_repository.My-repo.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// {{% /example %}}
    /// ## google\_artifact\_registry\_repository\_iam\_binding
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
    ///         var binding = new Gcp.ArtifactRegistry.RepositoryIamBinding("binding", new Gcp.ArtifactRegistry.RepositoryIamBindingArgs
    ///         {
    ///             Project = google_artifact_registry_repository.My-repo.Project,
    ///             Location = google_artifact_registry_repository.My-repo.Location,
    ///             Repository = google_artifact_registry_repository.My-repo.Name,
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
    /// ## google\_artifact\_registry\_repository\_iam\_member
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
    ///         var member = new Gcp.ArtifactRegistry.RepositoryIamMember("member", new Gcp.ArtifactRegistry.RepositoryIamMemberArgs
    ///         {
    ///             Project = google_artifact_registry_repository.My-repo.Project,
    ///             Location = google_artifact_registry_repository.My-repo.Location,
    ///             Repository = google_artifact_registry_repository.My-repo.Name,
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
    public partial class RepositoryIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.RepositoryIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the location this repository is located in.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryIamMember(string name, RepositoryIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/repositoryIamMember:RepositoryIamMember", name, args ?? new RepositoryIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryIamMember(string name, Input<string> id, RepositoryIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:artifactregistry/repositoryIamMember:RepositoryIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryIamMember Get(string name, Input<string> id, RepositoryIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryIamMember(name, id, state, options);
        }
    }

    public sealed class RepositoryIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.RepositoryIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The name of the location this repository is located in.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public RepositoryIamMemberArgs()
        {
        }
    }

    public sealed class RepositoryIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.RepositoryIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name of the location this repository is located in.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.artifactregistry.RepositoryIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public RepositoryIamMemberState()
        {
        }
    }
}
