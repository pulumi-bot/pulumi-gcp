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
    /// &gt; **Note:** It is not possible to grant the `roles/owner` role using any of these resources due to this being disallowed by the underlying `projects.setIamPolicy` API method. See the method [documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy) for full details. It is, however, possible to remove all owners from the project by passing in an empty `members = []` list to the `gcp.projects.IAMBinding` resource. This is useful for removing the owner role from a project upon creation, however, precautions should be taken to avoid inadvertently locking oneself out of a project such as by granting additional roles to alternate entities.
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `project_id`, role, and member e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `project_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `project_id`.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/iAMMember:IAMMember my_project your-project-id
    /// ```
    /// 
    ///  IAM audit config imports use the identifier of the resource in question and the service, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:projects/iAMMember:IAMMember my_project "your-project-id foo.googleapis.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class IAMMember : Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.IAMMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the project's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

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
        /// Create a IAMMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMMember(string name, IAMMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/iAMMember:IAMMember", name, args ?? new IAMMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMMember(string name, Input<string> id, IAMMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/iAMMember:IAMMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMMember Get(string name, Input<string> id, IAMMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMMember(name, id, state, options);
        }
    }

    public sealed class IAMMemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

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

        public IAMMemberArgs()
        {
        }
    }

    public sealed class IAMMemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the project's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

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

        public IAMMemberState()
        {
        }
    }
}
