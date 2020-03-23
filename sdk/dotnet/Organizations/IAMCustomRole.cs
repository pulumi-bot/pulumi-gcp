// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Organizations
{
    /// <summary>
    /// Allows management of a customized Cloud IAM organization role. For more information see
    /// [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
    /// and
    /// [API](https://cloud.google.com/iam/reference/rest/v1/organizations.roles).
    /// 
    /// &gt; **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
    ///  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
    ///  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
    ///  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
    ///  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
    ///  by this provider, and new roles cannot share that name.
    ///  
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_organization_iam_custom_role.html.markdown.
    /// </summary>
    public partial class IAMCustomRole : Pulumi.CustomResource
    {
        /// <summary>
        /// (Optional) The current deleted state of the role.
        /// </summary>
        [Output("deleted")]
        public Output<bool> Deleted { get; private set; } = null!;

        /// <summary>
        /// A human-readable description for the role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The numeric ID of the organization in which you want to create a custom role.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The role id to use for this role.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// The current launch stage of the role.
        /// Defaults to `GA`.
        /// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        /// </summary>
        [Output("stage")]
        public Output<string?> Stage { get; private set; } = null!;

        /// <summary>
        /// A human-readable title for the role.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a IAMCustomRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMCustomRole(string name, IAMCustomRoleArgs args, CustomResourceOptions? options = null)
            : base("gcp:organizations/iAMCustomRole:IAMCustomRole", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IAMCustomRole(string name, Input<string> id, IAMCustomRoleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:organizations/iAMCustomRole:IAMCustomRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMCustomRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMCustomRole Get(string name, Input<string> id, IAMCustomRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMCustomRole(name, id, state, options);
        }
    }

    public sealed class IAMCustomRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description for the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The numeric ID of the organization in which you want to create a custom role.
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        [Input("permissions", required: true)]
        private InputList<string>? _permissions;

        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The role id to use for this role.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// The current launch stage of the role.
        /// Defaults to `GA`.
        /// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        /// </summary>
        [Input("stage")]
        public Input<string>? Stage { get; set; }

        /// <summary>
        /// A human-readable title for the role.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public IAMCustomRoleArgs()
        {
        }
    }

    public sealed class IAMCustomRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) The current deleted state of the role.
        /// </summary>
        [Input("deleted")]
        public Input<bool>? Deleted { get; set; }

        /// <summary>
        /// A human-readable description for the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The numeric ID of the organization in which you want to create a custom role.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The role id to use for this role.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// The current launch stage of the role.
        /// Defaults to `GA`.
        /// List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        /// </summary>
        [Input("stage")]
        public Input<string>? Stage { get; set; }

        /// <summary>
        /// A human-readable title for the role.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public IAMCustomRoleState()
        {
        }
    }
}
