// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase
{
    public partial class ProjectLocation : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Required)
        /// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
        /// resource locations.
        /// </summary>
        [Output("locationId")]
        public Output<string> LocationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectLocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectLocation(string name, ProjectLocationArgs args, CustomResourceOptions? options = null)
            : base("gcp:firebase/projectLocation:ProjectLocation", name, args ?? new ProjectLocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectLocation(string name, Input<string> id, ProjectLocationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firebase/projectLocation:ProjectLocation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectLocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectLocation Get(string name, Input<string> id, ProjectLocationState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectLocation(name, id, state, options);
        }
    }

    public sealed class ProjectLocationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
        /// resource locations.
        /// </summary>
        [Input("locationId", required: true)]
        public Input<string> LocationId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectLocationArgs()
        {
        }
    }

    public sealed class ProjectLocationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Required)
        /// The ID of the default GCP resource location for the Project. The location must be one of the available GCP
        /// resource locations.
        /// </summary>
        [Input("locationId")]
        public Input<string>? LocationId { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectLocationState()
        {
        }
    }
}
