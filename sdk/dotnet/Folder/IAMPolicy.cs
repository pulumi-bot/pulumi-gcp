// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Folder
{
    /// <summary>
    /// Allows creation and management of the IAM policy for an existing Google Cloud
    /// Platform folder.
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
    ///         var department1 = new Gcp.Organizations.Folder("department1", new Gcp.Organizations.FolderArgs
    ///         {
    ///             DisplayName = "Department 1",
    ///             Parent = "organizations/1234567",
    ///         });
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
    ///         var folderAdminPolicy = new Gcp.Folder.IAMPolicy("folderAdminPolicy", new Gcp.Folder.IAMPolicyArgs
    ///         {
    ///             Folder = department1.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class IAMPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. This policy overrides any existing
        /// policy applied to the folder.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a IAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMPolicy(string name, IAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, args ?? new IAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMPolicy(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMPolicy Get(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMPolicy(name, id, state, options);
        }
    }

    public sealed class IAMPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. This policy overrides any existing
        /// policy applied to the folder.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public IAMPolicyArgs()
        {
        }
    }

    public sealed class IAMPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the folder's IAM policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. This policy overrides any existing
        /// policy applied to the folder.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public IAMPolicyState()
        {
        }
    }
}
