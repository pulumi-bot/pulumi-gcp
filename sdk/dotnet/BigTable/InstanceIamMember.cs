// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable
{
    /// <summary>
    /// Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
    /// * `gcp.bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
    /// * `gcp.bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
    /// 
    /// &gt; **Note:** `gcp.bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.InstanceIamBinding` and `gcp.bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `gcp.bigtable.InstanceIamPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.bigtable.InstanceIamBinding` resources **can be** used in conjunction with `gcp.bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## Import
    /// 
    /// Instance IAM resources can be imported using the project, instance name, role and/or member.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance}"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance} roles/editor"
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/instanceIamMember:InstanceIamMember editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    public partial class InstanceIamMember : Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.InstanceIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the instances's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name or relative resource id of the instance to manage IAM policies for.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The project in which the instance belongs. If it
        /// is not provided, a default will be supplied.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceIamMember(string name, InstanceIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigtable/instanceIamMember:InstanceIamMember", name, args ?? new InstanceIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceIamMember(string name, Input<string> id, InstanceIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/instanceIamMember:InstanceIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceIamMember Get(string name, Input<string> id, InstanceIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceIamMember(name, id, state, options);
        }
    }

    public sealed class InstanceIamMemberArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.InstanceIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The name or relative resource id of the instance to manage IAM policies for.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The project in which the instance belongs. If it
        /// is not provided, a default will be supplied.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public InstanceIamMemberArgs()
        {
        }
    }

    public sealed class InstanceIamMemberState : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.InstanceIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the instances's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name or relative resource id of the instance to manage IAM policies for.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The project in which the instance belongs. If it
        /// is not provided, a default will be supplied.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public InstanceIamMemberState()
        {
        }
    }
}
