// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Folder
{
    [GcpResourceType("gcp:folder/iAMPolicy:IAMPolicy")]
    public partial class IAMPolicy : Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a IAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMPolicy(string name, IAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, args ?? new IAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMPolicy(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMPolicy Get(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMPolicy(name, id, state, options);
        }
    }

    public sealed class IAMPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public IAMPolicyArgs()
        {
        }
    }

    public sealed class IAMPolicyState : Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("folder")]
        public Input<string>? Folder { get; set; }

        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public IAMPolicyState()
        {
        }
    }
}
