// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// ## Import
    /// 
    /// PolicyTag can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/policyTag:PolicyTag default {{name}}
    /// ```
    /// </summary>
    public partial class PolicyTag : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource names of child policy tags of this policy tag.
        /// </summary>
        [Output("childPolicyTags")]
        public Output<ImmutableArray<string>> ChildPolicyTags { get; private set; } = null!;

        /// <summary>
        /// Description of this policy tag. It must: contain only unicode characters, tabs,
        /// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
        /// encoded in UTF-8. If not set, defaults to an empty description.
        /// If not set, defaults to an empty description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User defined name of this policy tag. It must: be unique within the parent
        /// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
        /// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource name of this policy tag, whose format is:
        /// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource name of this policy tag's parent policy tag.
        /// If empty, it means this policy tag is a top level policy tag.
        /// If not set, defaults to an empty string.
        /// </summary>
        [Output("parentPolicyTag")]
        public Output<string?> ParentPolicyTag { get; private set; } = null!;

        /// <summary>
        /// Taxonomy the policy tag is associated with
        /// </summary>
        [Output("taxonomy")]
        public Output<string> Taxonomy { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyTag resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyTag(string name, PolicyTagArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/policyTag:PolicyTag", name, args ?? new PolicyTagArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyTag(string name, Input<string> id, PolicyTagState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/policyTag:PolicyTag", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyTag resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyTag Get(string name, Input<string> id, PolicyTagState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyTag(name, id, state, options);
        }
    }

    public sealed class PolicyTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of this policy tag. It must: contain only unicode characters, tabs,
        /// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
        /// encoded in UTF-8. If not set, defaults to an empty description.
        /// If not set, defaults to an empty description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User defined name of this policy tag. It must: be unique within the parent
        /// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
        /// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Resource name of this policy tag's parent policy tag.
        /// If empty, it means this policy tag is a top level policy tag.
        /// If not set, defaults to an empty string.
        /// </summary>
        [Input("parentPolicyTag")]
        public Input<string>? ParentPolicyTag { get; set; }

        /// <summary>
        /// Taxonomy the policy tag is associated with
        /// </summary>
        [Input("taxonomy", required: true)]
        public Input<string> Taxonomy { get; set; } = null!;

        public PolicyTagArgs()
        {
        }
    }

    public sealed class PolicyTagState : Pulumi.ResourceArgs
    {
        [Input("childPolicyTags")]
        private InputList<string>? _childPolicyTags;

        /// <summary>
        /// Resource names of child policy tags of this policy tag.
        /// </summary>
        public InputList<string> ChildPolicyTags
        {
            get => _childPolicyTags ?? (_childPolicyTags = new InputList<string>());
            set => _childPolicyTags = value;
        }

        /// <summary>
        /// Description of this policy tag. It must: contain only unicode characters, tabs,
        /// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
        /// encoded in UTF-8. If not set, defaults to an empty description.
        /// If not set, defaults to an empty description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User defined name of this policy tag. It must: be unique within the parent
        /// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
        /// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Resource name of this policy tag, whose format is:
        /// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource name of this policy tag's parent policy tag.
        /// If empty, it means this policy tag is a top level policy tag.
        /// If not set, defaults to an empty string.
        /// </summary>
        [Input("parentPolicyTag")]
        public Input<string>? ParentPolicyTag { get; set; }

        /// <summary>
        /// Taxonomy the policy tag is associated with
        /// </summary>
        [Input("taxonomy")]
        public Input<string>? Taxonomy { get; set; }

        public PolicyTagState()
        {
        }
    }
}
