// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ResourceManager
{
    /// <summary>
    /// A Lien represents an encumbrance on the actions that can be performed on a resource.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Lien can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:resourcemanager/lien:Lien default {{parent}}/{{name}}
    /// ```
    /// </summary>
    public partial class Lien : Pulumi.CustomResource
    {
        /// <summary>
        /// Time of creation
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A system-generated unique identifier for this Lien.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A stable, user-visible/meaningful string identifying the origin
        /// of the Lien, intended to be inspected programmatically. Maximum length of
        /// 200 characters.
        /// </summary>
        [Output("origin")]
        public Output<string> Origin { get; private set; } = null!;

        /// <summary>
        /// A reference to the resource this Lien is attached to.
        /// The server will validate the parent against those for which Liens are supported.
        /// Since a variety of objects can have Liens against them, you must provide the type
        /// prefix (e.g. "projects/my-project-name").
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Concise user-visible strings indicating why an action cannot be performed
        /// on a resource. Maximum length of 200 characters.
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// The types of operations which should be blocked as a result of this Lien.
        /// Each value should correspond to an IAM permission. The server will validate
        /// the permissions against those for which Liens are supported.  An empty
        /// list is meaningless and will be rejected.
        /// e.g. ['resourcemanager.projects.delete']
        /// </summary>
        [Output("restrictions")]
        public Output<ImmutableArray<string>> Restrictions { get; private set; } = null!;


        /// <summary>
        /// Create a Lien resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lien(string name, LienArgs args, CustomResourceOptions? options = null)
            : base("gcp:resourcemanager/lien:Lien", name, args ?? new LienArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lien(string name, Input<string> id, LienState? state = null, CustomResourceOptions? options = null)
            : base("gcp:resourcemanager/lien:Lien", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Lien resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lien Get(string name, Input<string> id, LienState? state = null, CustomResourceOptions? options = null)
        {
            return new Lien(name, id, state, options);
        }
    }

    public sealed class LienArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A stable, user-visible/meaningful string identifying the origin
        /// of the Lien, intended to be inspected programmatically. Maximum length of
        /// 200 characters.
        /// </summary>
        [Input("origin", required: true)]
        public Input<string> Origin { get; set; } = null!;

        /// <summary>
        /// A reference to the resource this Lien is attached to.
        /// The server will validate the parent against those for which Liens are supported.
        /// Since a variety of objects can have Liens against them, you must provide the type
        /// prefix (e.g. "projects/my-project-name").
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Concise user-visible strings indicating why an action cannot be performed
        /// on a resource. Maximum length of 200 characters.
        /// </summary>
        [Input("reason", required: true)]
        public Input<string> Reason { get; set; } = null!;

        [Input("restrictions", required: true)]
        private InputList<string>? _restrictions;

        /// <summary>
        /// The types of operations which should be blocked as a result of this Lien.
        /// Each value should correspond to an IAM permission. The server will validate
        /// the permissions against those for which Liens are supported.  An empty
        /// list is meaningless and will be rejected.
        /// e.g. ['resourcemanager.projects.delete']
        /// </summary>
        public InputList<string> Restrictions
        {
            get => _restrictions ?? (_restrictions = new InputList<string>());
            set => _restrictions = value;
        }

        public LienArgs()
        {
        }
    }

    public sealed class LienState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time of creation
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A system-generated unique identifier for this Lien.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A stable, user-visible/meaningful string identifying the origin
        /// of the Lien, intended to be inspected programmatically. Maximum length of
        /// 200 characters.
        /// </summary>
        [Input("origin")]
        public Input<string>? Origin { get; set; }

        /// <summary>
        /// A reference to the resource this Lien is attached to.
        /// The server will validate the parent against those for which Liens are supported.
        /// Since a variety of objects can have Liens against them, you must provide the type
        /// prefix (e.g. "projects/my-project-name").
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Concise user-visible strings indicating why an action cannot be performed
        /// on a resource. Maximum length of 200 characters.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        [Input("restrictions")]
        private InputList<string>? _restrictions;

        /// <summary>
        /// The types of operations which should be blocked as a result of this Lien.
        /// Each value should correspond to an IAM permission. The server will validate
        /// the permissions against those for which Liens are supported.  An empty
        /// list is meaningless and will be rejected.
        /// e.g. ['resourcemanager.projects.delete']
        /// </summary>
        public InputList<string> Restrictions
        {
            get => _restrictions ?? (_restrictions = new InputList<string>());
            set => _restrictions = value;
        }

        public LienState()
        {
        }
    }
}
