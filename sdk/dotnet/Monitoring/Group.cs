// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// The description of a dynamic collection of monitored resources. Each group
    /// has a filter that is matched against monitored resources and their
    /// associated metadata. If a group's filter matches an available monitored
    /// resource, then that resource is a member of that group.
    /// 
    /// 
    /// To get more information about Group, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.groups)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/groups/)
    /// </summary>
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// A user-assigned name for this group, used only for display
        /// purposes.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The filter used to determine which monitored resources
        /// belong to this group.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// If true, the members of this group are considered to be a
        /// cluster. The system can perform additional analysis on
        /// groups that are clusters.
        /// </summary>
        [Output("isCluster")]
        public Output<bool?> IsCluster { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the group's parent, if it has one. The format is
        /// "projects/{project_id_or_number}/groups/{group_id}". For
        /// groups with no parent, parentName is the empty string, "".
        /// </summary>
        [Output("parentName")]
        public Output<string?> ParentName { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-assigned name for this group, used only for display
        /// purposes.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The filter used to determine which monitored resources
        /// belong to this group.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// If true, the members of this group are considered to be a
        /// cluster. The system can perform additional analysis on
        /// groups that are clusters.
        /// </summary>
        [Input("isCluster")]
        public Input<bool>? IsCluster { get; set; }

        /// <summary>
        /// The name of the group's parent, if it has one. The format is
        /// "projects/{project_id_or_number}/groups/{group_id}". For
        /// groups with no parent, parentName is the empty string, "".
        /// </summary>
        [Input("parentName")]
        public Input<string>? ParentName { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user-assigned name for this group, used only for display
        /// purposes.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The filter used to determine which monitored resources
        /// belong to this group.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// If true, the members of this group are considered to be a
        /// cluster. The system can perform additional analysis on
        /// groups that are clusters.
        /// </summary>
        [Input("isCluster")]
        public Input<bool>? IsCluster { get; set; }

        /// <summary>
        /// A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the group's parent, if it has one. The format is
        /// "projects/{project_id_or_number}/groups/{group_id}". For
        /// groups with no parent, parentName is the empty string, "".
        /// </summary>
        [Input("parentName")]
        public Input<string>? ParentName { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GroupState()
        {
        }
    }
}
