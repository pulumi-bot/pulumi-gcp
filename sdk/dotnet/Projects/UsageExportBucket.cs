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
    /// Allows creation and management of a Google Cloud Platform project.
    /// 
    /// Projects created with this resource must be associated with an Organization.
    /// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
    /// 
    /// The service account used to run this provider when creating a `gcp.organizations.Project`
    /// resource must have `roles/resourcemanager.projectCreator`. See the
    /// [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
    /// doc for more information.
    /// 
    /// Note that prior to 0.8.5, `gcp.organizations.Project` functioned like a data source,
    /// meaning any project referenced by it had to be created and managed outside
    /// this provider. As of 0.8.5, `gcp.organizations.Project` functions like any other
    /// resource, with this provider creating and managing the project. To replicate the old
    /// behavior, either:
    /// 
    /// * Use the project ID directly in whatever is referencing the project, using the
    ///   [gcp.projects.IAMPolicy](https://www.terraform.io/docs/providers/google/r/google_project_iam.html)
    ///   to replace the old `policy_data` property.
    /// * Use the [import](https://www.terraform.io/docs/import/usage.html) functionality
    ///   to import your pre-existing project into this provider, where it can be referenced and
    ///   used just like always, keeping in mind that this provider will attempt to undo any changes
    ///   made outside this provider.
    /// 
    /// &gt; It's important to note that any project resources that were added to your config
    /// prior to 0.8.5 will continue to function as they always have, and will not be managed by
    /// this provider. Only newly added projects are affected.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_project.html.markdown.
    /// </summary>
    public partial class UsageExportBucket : Pulumi.CustomResource
    {
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a UsageExportBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsageExportBucket(string name, UsageExportBucketArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/usageExportBucket:UsageExportBucket", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private UsageExportBucket(string name, Input<string> id, UsageExportBucketState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/usageExportBucket:UsageExportBucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsageExportBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsageExportBucket Get(string name, Input<string> id, UsageExportBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new UsageExportBucket(name, id, state, options);
        }
    }

    public sealed class UsageExportBucketArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public UsageExportBucketArgs()
        {
        }
    }

    public sealed class UsageExportBucketState : Pulumi.ResourceArgs
    {
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public UsageExportBucketState()
        {
        }
    }
}
