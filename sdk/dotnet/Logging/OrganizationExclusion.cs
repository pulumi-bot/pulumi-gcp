// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// Manages an organization-level logging exclusion. For more information see
    /// [the official documentation](https://cloud.google.com/logging/docs/) and
    /// [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
    /// 
    /// Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    /// granted to the credentials used with this provider.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_organization_exclusion.html.markdown.
    /// </summary>
    public partial class OrganizationExclusion : Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization to create the exclusion in.
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationExclusion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationExclusion(string name, OrganizationExclusionArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/organizationExclusion:OrganizationExclusion", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private OrganizationExclusion(string name, Input<string> id, OrganizationExclusionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/organizationExclusion:OrganizationExclusion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationExclusion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationExclusion Get(string name, Input<string> id, OrganizationExclusionState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationExclusion(name, id, state, options);
        }
    }

    public sealed class OrganizationExclusionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to create the exclusion in.
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public OrganizationExclusionArgs()
        {
        }
    }

    public sealed class OrganizationExclusionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization to create the exclusion in.
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public OrganizationExclusionState()
        {
        }
    }
}
