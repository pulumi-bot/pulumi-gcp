// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* b/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Storage bucket IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor "b/{{bucket}} roles/storage.objectViewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor "b/{{bucket}} roles/storage.objectViewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/bucketIAMPolicy:BucketIAMPolicy editor b/{{bucket}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:storage/bucketIAMPolicy:BucketIAMPolicy")]
    public partial class BucketIAMPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a BucketIAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketIAMPolicy(string name, BucketIAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, args ?? new BucketIAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketIAMPolicy(string name, Input<string> id, BucketIAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/bucketIAMPolicy:BucketIAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketIAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketIAMPolicy Get(string name, Input<string> id, BucketIAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketIAMPolicy(name, id, state, options);
        }
    }

    public sealed class BucketIAMPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public BucketIAMPolicyArgs()
        {
        }
    }

    public sealed class BucketIAMPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public BucketIAMPolicyState()
        {
        }
    }
}
