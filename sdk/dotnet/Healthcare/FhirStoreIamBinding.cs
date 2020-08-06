// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
    /// * `gcp.healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
    /// * `gcp.healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.FhirStoreIamBinding` and `gcp.healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_healthcare\_fhir\_store\_iam\_policy
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
    ///         var fhirStore = new Gcp.Healthcare.FhirStoreIamPolicy("fhirStore", new Gcp.Healthcare.FhirStoreIamPolicyArgs
    ///         {
    ///             FhirStoreId = "your-fhir-store-id",
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## google\_healthcare\_fhir\_store\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fhirStore = new Gcp.Healthcare.FhirStoreIamBinding("fhirStore", new Gcp.Healthcare.FhirStoreIamBindingArgs
    ///         {
    ///             FhirStoreId = "your-fhir-store-id",
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
    /// ## google\_healthcare\_fhir\_store\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fhirStore = new Gcp.Healthcare.FhirStoreIamMember("fhirStore", new Gcp.Healthcare.FhirStoreIamMemberArgs
    ///         {
    ///             FhirStoreId = "your-fhir-store-id",
    ///             Member = "user:jane@example.com",
    ///             Role = "roles/editor",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class FhirStoreIamBinding : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.FhirStoreIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the FHIR store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("fhirStoreId")]
        public Output<string> FhirStoreId { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a FhirStoreIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FhirStoreIamBinding(string name, FhirStoreIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, args ?? new FhirStoreIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FhirStoreIamBinding(string name, Input<string> id, FhirStoreIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStoreIamBinding:FhirStoreIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FhirStoreIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FhirStoreIamBinding Get(string name, Input<string> id, FhirStoreIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new FhirStoreIamBinding(name, id, state, options);
        }
    }

    public sealed class FhirStoreIamBindingArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.FhirStoreIamBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("fhirStoreId", required: true)]
        public Input<string> FhirStoreId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public FhirStoreIamBindingArgs()
        {
        }
    }

    public sealed class FhirStoreIamBindingState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.FhirStoreIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the FHIR store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("fhirStoreId")]
        public Input<string>? FhirStoreId { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public FhirStoreIamBindingState()
        {
        }
    }
}
