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
    /// Authoritatively manages the access control list (ACL) for an object in a Google
    /// Cloud Storage (GCS) bucket. Removing a `gcp.storage.ObjectACL` sets the
    /// acl to the `private` [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl).
    /// 
    /// For more information see
    /// [the official documentation](https://cloud.google.com/storage/docs/access-control/lists) 
    /// and 
    /// [API](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls).
    /// 
    /// &gt; Want fine-grained control over object ACLs? Use `gcp.storage.ObjectAccessControl` to control individual
    /// role entity pairs.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_object_acl.html.markdown.
    /// </summary>
    public partial class ObjectACL : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket the object is stored in.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The name of the object to apply the acl to.
        /// </summary>
        [Output("object")]
        public Output<string> Object { get; private set; } = null!;

        /// <summary>
        /// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        /// </summary>
        [Output("predefinedAcl")]
        public Output<string?> PredefinedAcl { get; private set; } = null!;

        /// <summary>
        /// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
        /// Must be set if `predefined_acl` is not.
        /// </summary>
        [Output("roleEntities")]
        public Output<ImmutableArray<string>> RoleEntities { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectACL resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectACL(string name, ObjectACLArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/objectACL:ObjectACL", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ObjectACL(string name, Input<string> id, ObjectACLState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/objectACL:ObjectACL", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectACL resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectACL Get(string name, Input<string> id, ObjectACLState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectACL(name, id, state, options);
        }
    }

    public sealed class ObjectACLArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket the object is stored in.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The name of the object to apply the acl to.
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        /// <summary>
        /// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        /// </summary>
        [Input("predefinedAcl")]
        public Input<string>? PredefinedAcl { get; set; }

        [Input("roleEntities")]
        private InputList<string>? _roleEntities;

        /// <summary>
        /// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
        /// Must be set if `predefined_acl` is not.
        /// </summary>
        public InputList<string> RoleEntities
        {
            get => _roleEntities ?? (_roleEntities = new InputList<string>());
            set => _roleEntities = value;
        }

        public ObjectACLArgs()
        {
        }
    }

    public sealed class ObjectACLState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket the object is stored in.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The name of the object to apply the acl to.
        /// </summary>
        [Input("object")]
        public Input<string>? Object { get; set; }

        /// <summary>
        /// The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if `role_entity` is not.
        /// </summary>
        [Input("predefinedAcl")]
        public Input<string>? PredefinedAcl { get; set; }

        [Input("roleEntities")]
        private InputList<string>? _roleEntities;

        /// <summary>
        /// List of role/entity pairs in the form `ROLE:entity`. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details.
        /// Must be set if `predefined_acl` is not.
        /// </summary>
        public InputList<string> RoleEntities
        {
            get => _roleEntities ?? (_roleEntities = new InputList<string>());
            set => _roleEntities = value;
        }

        public ObjectACLState()
        {
        }
    }
}
