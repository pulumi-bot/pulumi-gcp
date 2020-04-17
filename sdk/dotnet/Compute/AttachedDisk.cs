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
    /// Persistent disks can be attached to a compute instance using the `attached_disk`
    /// section within the compute instance configuration.
    /// However there may be situations where managing the attached disks via the compute
    /// instance config isn't preferable or possible, such as attaching dynamic
    /// numbers of disks using the `count` variable.
    /// 
    /// 
    /// To get more information about attaching disks, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
    /// * How-to Guides
    ///     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
    /// 
    /// **Note:** When using `gcp.compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attached_disk"]` on the `gcp.compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
    /// </summary>
    public partial class AttachedDisk : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a unique device name of your choice that is
        /// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        /// system running within the instance. This name can be used to
        /// reference the device for mounting, resizing, and so on, from within
        /// the instance.
        /// </summary>
        [Output("deviceName")]
        public Output<string> DeviceName { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the disk that will be attached.
        /// </summary>
        [Output("disk")]
        public Output<string> Disk { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the compute instance that the disk will be attached to.
        /// If the `self_link` is provided then `zone` and `project` are extracted from the
        /// self link. If only the name is used then `zone` and `project` must be defined
        /// as properties on the resource or provider.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The mode in which to attach this disk, either READ_WRITE or
        /// READ_ONLY. If not specified, the default is to attach the disk in
        /// READ_WRITE mode.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The project that the referenced compute instance is a part of. If `instance` is referenced by its
        /// `self_link` the project defined in the link will take precedence.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The zone that the referenced compute instance is located within. If `instance` is referenced by its
        /// `self_link` the zone defined in the link will take precedence.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a AttachedDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttachedDisk(string name, AttachedDiskArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/attachedDisk:AttachedDisk", name, args ?? new AttachedDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttachedDisk(string name, Input<string> id, AttachedDiskState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/attachedDisk:AttachedDisk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AttachedDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttachedDisk Get(string name, Input<string> id, AttachedDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new AttachedDisk(name, id, state, options);
        }
    }

    public sealed class AttachedDiskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a unique device name of your choice that is
        /// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        /// system running within the instance. This name can be used to
        /// reference the device for mounting, resizing, and so on, from within
        /// the instance.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the disk that will be attached.
        /// </summary>
        [Input("disk", required: true)]
        public Input<string> Disk { get; set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the compute instance that the disk will be attached to.
        /// If the `self_link` is provided then `zone` and `project` are extracted from the
        /// self link. If only the name is used then `zone` and `project` must be defined
        /// as properties on the resource or provider.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The mode in which to attach this disk, either READ_WRITE or
        /// READ_ONLY. If not specified, the default is to attach the disk in
        /// READ_WRITE mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The project that the referenced compute instance is a part of. If `instance` is referenced by its
        /// `self_link` the project defined in the link will take precedence.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The zone that the referenced compute instance is located within. If `instance` is referenced by its
        /// `self_link` the zone defined in the link will take precedence.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AttachedDiskArgs()
        {
        }
    }

    public sealed class AttachedDiskState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// Specifies a unique device name of your choice that is
        /// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        /// system running within the instance. This name can be used to
        /// reference the device for mounting, resizing, and so on, from within
        /// the instance.
        /// </summary>
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the disk that will be attached.
        /// </summary>
        [Input("disk")]
        public Input<string>? Disk { get; set; }

        /// <summary>
        /// -
        /// (Required)
        /// `name` or `self_link` of the compute instance that the disk will be attached to.
        /// If the `self_link` is provided then `zone` and `project` are extracted from the
        /// self link. If only the name is used then `zone` and `project` must be defined
        /// as properties on the resource or provider.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The mode in which to attach this disk, either READ_WRITE or
        /// READ_ONLY. If not specified, the default is to attach the disk in
        /// READ_WRITE mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The project that the referenced compute instance is a part of. If `instance` is referenced by its
        /// `self_link` the project defined in the link will take precedence.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The zone that the referenced compute instance is located within. If `instance` is referenced by its
        /// `self_link` the zone defined in the link will take precedence.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AttachedDiskState()
        {
        }
    }
}
