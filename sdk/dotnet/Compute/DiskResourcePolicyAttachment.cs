// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Adds existing resource policies to a disk. You can only add one policy
    /// which will be applied to this disk for scheduling snapshot creation.
    /// 
    /// &gt; **Note:** This resource does not support regional disks (`gcp.compute.RegionDisk`).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_disk_resource_policy_attachment.html.markdown.
    /// </summary>
    public partial class DiskResourcePolicyAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the disk in which the resource policies are attached to.
        /// </summary>
        [Output("disk")]
        public Output<string> Disk { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a DiskResourcePolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskResourcePolicyAttachment(string name, DiskResourcePolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DiskResourcePolicyAttachment(string name, Input<string> id, DiskResourcePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskResourcePolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskResourcePolicyAttachment Get(string name, Input<string> id, DiskResourcePolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskResourcePolicyAttachment(name, id, state, options);
        }
    }

    public sealed class DiskResourcePolicyAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the disk in which the resource policies are attached to.
        /// </summary>
        [Input("disk", required: true)]
        public Input<string> Disk { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DiskResourcePolicyAttachmentArgs()
        {
        }
    }

    public sealed class DiskResourcePolicyAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The name of the disk in which the resource policies are attached to.
        /// </summary>
        [Input("disk")]
        public Input<string>? Disk { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// The resource policy to be attached to the disk for scheduling snapshot
        /// creation. Do not specify the self link.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// A reference to the zone where the disk resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DiskResourcePolicyAttachmentState()
        {
        }
    }
}
