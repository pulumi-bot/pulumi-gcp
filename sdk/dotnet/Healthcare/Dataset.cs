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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_dataset.html.markdown.
    /// </summary>
    public partial class Dataset : Pulumi.CustomResource
    {
        /// <summary>
        /// The location for the Dataset.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name for the Dataset.
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
        /// The fully qualified name of this dataset
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        /// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7
        /// messages) where no explicit timezone is specified.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dataset:Dataset", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dataset:Dataset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, state, options);
        }
    }

    public sealed class DatasetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location for the Dataset.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The resource name for the Dataset.
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
        /// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        /// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7
        /// messages) where no explicit timezone is specified.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public DatasetArgs()
        {
        }
    }

    public sealed class DatasetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location for the Dataset.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name for the Dataset.
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
        /// The fully qualified name of this dataset
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        /// "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7
        /// messages) where no explicit timezone is specified.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public DatasetState()
        {
        }
    }
}
