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
    /// A policy that can be attached to a resource to specify or schedule actions on that resource.
    /// 
    /// ## Example Usage
    /// ### Resource Policy Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Gcp.Compute.ResourcePolicy("foo", new Gcp.Compute.ResourcePolicyArgs
    ///         {
    ///             Region = "us-central1",
    ///             SnapshotSchedulePolicy = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyArgs
    ///             {
    ///                 Schedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs
    ///                 {
    ///                     DailySchedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleArgs
    ///                     {
    ///                         DaysInCycle = 1,
    ///                         StartTime = "04:00",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Resource Policy Full
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bar = new Gcp.Compute.ResourcePolicy("bar", new Gcp.Compute.ResourcePolicyArgs
    ///         {
    ///             Region = "us-central1",
    ///             SnapshotSchedulePolicy = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyArgs
    ///             {
    ///                 RetentionPolicy = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyRetentionPolicyArgs
    ///                 {
    ///                     MaxRetentionDays = 10,
    ///                     OnSourceDiskDelete = "KEEP_AUTO_SNAPSHOTS",
    ///                 },
    ///                 Schedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleArgs
    ///                 {
    ///                     HourlySchedule = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleArgs
    ///                     {
    ///                         HoursInCycle = 20,
    ///                         StartTime = "23:00",
    ///                     },
    ///                 },
    ///                 SnapshotProperties = new Gcp.Compute.Inputs.ResourcePolicySnapshotSchedulePolicySnapshotPropertiesArgs
    ///                 {
    ///                     GuestFlush = true,
    ///                     Labels = 
    ///                     {
    ///                         { "myLabel", "value" },
    ///                     },
    ///                     StorageLocations = "us",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Resource Policy Placement Policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var baz = new Gcp.Compute.ResourcePolicy("baz", new Gcp.Compute.ResourcePolicyArgs
    ///         {
    ///             GroupPlacementPolicy = new Gcp.Compute.Inputs.ResourcePolicyGroupPlacementPolicyArgs
    ///             {
    ///                 Collocation = "COLLOCATED",
    ///                 VmCount = 2,
    ///             },
    ///             Region = "us-central1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ResourcePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Output("groupPlacementPolicy")]
        public Output<Outputs.ResourcePolicyGroupPlacementPolicy?> GroupPlacementPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the resource, provided by the client when initially creating
        /// the resource. The resource name must be 1-63 characters long, and comply
        /// with RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z`? which means the
        /// first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character,
        /// which cannot be a dash.
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
        /// Region where resource policy resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Output("snapshotSchedulePolicy")]
        public Output<Outputs.ResourcePolicySnapshotSchedulePolicy?> SnapshotSchedulePolicy { get; private set; } = null!;


        /// <summary>
        /// Create a ResourcePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourcePolicy(string name, ResourcePolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/resourcePolicy:ResourcePolicy", name, args ?? new ResourcePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourcePolicy(string name, Input<string> id, ResourcePolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/resourcePolicy:ResourcePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourcePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourcePolicy Get(string name, Input<string> id, ResourcePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourcePolicy(name, id, state, options);
        }
    }

    public sealed class ResourcePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Input("groupPlacementPolicy")]
        public Input<Inputs.ResourcePolicyGroupPlacementPolicyArgs>? GroupPlacementPolicy { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating
        /// the resource. The resource name must be 1-63 characters long, and comply
        /// with RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z`? which means the
        /// first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character,
        /// which cannot be a dash.
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
        /// Region where resource policy resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Input("snapshotSchedulePolicy")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyArgs>? SnapshotSchedulePolicy { get; set; }

        public ResourcePolicyArgs()
        {
        }
    }

    public sealed class ResourcePolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Input("groupPlacementPolicy")]
        public Input<Inputs.ResourcePolicyGroupPlacementPolicyGetArgs>? GroupPlacementPolicy { get; set; }

        /// <summary>
        /// The name of the resource, provided by the client when initially creating
        /// the resource. The resource name must be 1-63 characters long, and comply
        /// with RFC1035. Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z`? which means the
        /// first character must be a lowercase letter, and all following characters
        /// must be a dash, lowercase letter, or digit, except the last character,
        /// which cannot be a dash.
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
        /// Region where resource policy resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// Policy for creating snapshots of persistent disks.  Structure is documented below.
        /// </summary>
        [Input("snapshotSchedulePolicy")]
        public Input<Inputs.ResourcePolicySnapshotSchedulePolicyGetArgs>? SnapshotSchedulePolicy { get; set; }

        public ResourcePolicyState()
        {
        }
    }
}
