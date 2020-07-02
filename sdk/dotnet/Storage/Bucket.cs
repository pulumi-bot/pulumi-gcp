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
    /// Creates a new bucket in Google cloud storage service (GCS).
    /// Once a bucket has been created, its location can't be changed.
    /// [ACLs](https://cloud.google.com/storage/docs/access-control/lists) can be applied
    /// using the [`gcp.storage.BucketACL`](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html) resource.
    /// 
    /// For more information see
    /// [the official documentation](https://cloud.google.com/storage/docs/overview)
    /// and
    /// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
    /// 
    /// **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
    /// determined which will require enabling the compute api.
    /// 
    /// ## Example Usage
    /// ### Creating A Private Bucket In Standard Storage, In The EU Region. Bucket Configured As Static Website And CORS Configurations
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var static_site = new Gcp.Storage.Bucket("static-site", new Gcp.Storage.BucketArgs
    ///         {
    ///             BucketPolicyOnly = true,
    ///             Cors = 
    ///             {
    ///                 new Gcp.Storage.Inputs.BucketCorArgs
    ///                 {
    ///                     MaxAgeSeconds = 3600,
    ///                     Methods = 
    ///                     {
    ///                         "GET",
    ///                         "HEAD",
    ///                         "PUT",
    ///                         "POST",
    ///                         "DELETE",
    ///                     },
    ///                     Origins = 
    ///                     {
    ///                         "http://image-store.com",
    ///                     },
    ///                     ResponseHeaders = 
    ///                     {
    ///                         "*",
    ///                     },
    ///                 },
    ///             },
    ///             ForceDestroy = true,
    ///             Location = "EU",
    ///             Website = new Gcp.Storage.Inputs.BucketWebsiteArgs
    ///             {
    ///                 MainPageSuffix = "index.html",
    ///                 NotFoundPage = "404.html",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Life Cycle Settings For Storage Bucket Objects
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var auto_expire = new Gcp.Storage.Bucket("auto-expire", new Gcp.Storage.BucketArgs
    ///         {
    ///             ForceDestroy = true,
    ///             LifecycleRules = 
    ///             {
    ///                 new Gcp.Storage.Inputs.BucketLifecycleRuleArgs
    ///                 {
    ///                     Action = new Gcp.Storage.Inputs.BucketLifecycleRuleActionArgs
    ///                     {
    ///                         Type = "Delete",
    ///                     },
    ///                     Condition = new Gcp.Storage.Inputs.BucketLifecycleRuleConditionArgs
    ///                     {
    ///                         Age = 3,
    ///                     },
    ///                 },
    ///             },
    ///             Location = "US",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Bucket : Pulumi.CustomResource
    {
        /// <summary>
        /// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        /// </summary>
        [Output("bucketPolicyOnly")]
        public Output<bool> BucketPolicyOnly { get; private set; } = null!;

        /// <summary>
        /// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        [Output("cors")]
        public Output<ImmutableArray<Outputs.BucketCor>> Cors { get; private set; } = null!;

        [Output("defaultEventBasedHold")]
        public Output<bool?> DefaultEventBasedHold { get; private set; } = null!;

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.BucketEncryption?> Encryption { get; private set; } = null!;

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, the provider will fail that run.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the bucket.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        [Output("lifecycleRules")]
        public Output<ImmutableArray<Outputs.BucketLifecycleRule>> LifecycleRules { get; private set; } = null!;

        /// <summary>
        /// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The bucket's [Access &amp; Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        /// </summary>
        [Output("logging")]
        public Output<Outputs.BucketLogging?> Logging { get; private set; } = null!;

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        /// </summary>
        [Output("requesterPays")]
        public Output<bool?> RequesterPays { get; private set; } = null!;

        /// <summary>
        /// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.BucketRetentionPolicy?> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        [Output("storageClass")]
        public Output<string?> StorageClass { get; private set; } = null!;

        /// <summary>
        /// The base URL of the bucket, in the format `gs://&lt;bucket-name&gt;`.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        /// </summary>
        [Output("versioning")]
        public Output<Outputs.BucketVersioning?> Versioning { get; private set; } = null!;

        /// <summary>
        /// Configuration if the bucket acts as a website. Structure is documented below.
        /// </summary>
        [Output("website")]
        public Output<Outputs.BucketWebsite?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bucket(string name, BucketArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:storage/bucket:Bucket", name, args ?? new BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Bucket(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/bucket:Bucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Bucket Get(string name, Input<string> id, BucketState? state = null, CustomResourceOptions? options = null)
        {
            return new Bucket(name, id, state, options);
        }
    }

    public sealed class BucketArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        /// </summary>
        [Input("bucketPolicyOnly")]
        public Input<bool>? BucketPolicyOnly { get; set; }

        [Input("cors")]
        private InputList<Inputs.BucketCorArgs>? _cors;

        /// <summary>
        /// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketCorArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.BucketCorArgs>());
            set => _cors = value;
        }

        [Input("defaultEventBasedHold")]
        public Input<bool>? DefaultEventBasedHold { get; set; }

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.BucketEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, the provider will fail that run.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the bucket.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleArgs>? _lifecycleRules;

        /// <summary>
        /// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRuleArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The bucket's [Access &amp; Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.BucketLoggingArgs>? Logging { get; set; }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        /// </summary>
        [Input("requesterPays")]
        public Input<bool>? RequesterPays { get; set; }

        /// <summary>
        /// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.BucketRetentionPolicyArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        /// <summary>
        /// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningArgs>? Versioning { get; set; }

        /// <summary>
        /// Configuration if the bucket acts as a website. Structure is documented below.
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteArgs>? Website { get; set; }

        public BucketArgs()
        {
        }
    }

    public sealed class BucketState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket.
        /// </summary>
        [Input("bucketPolicyOnly")]
        public Input<bool>? BucketPolicyOnly { get; set; }

        [Input("cors")]
        private InputList<Inputs.BucketCorGetArgs>? _cors;

        /// <summary>
        /// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketCorGetArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.BucketCorGetArgs>());
            set => _cors = value;
        }

        [Input("defaultEventBasedHold")]
        public Input<bool>? DefaultEventBasedHold { get; set; }

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.BucketEncryptionGetArgs>? Encryption { get; set; }

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, the provider will fail that run.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the bucket.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("lifecycleRules")]
        private InputList<Inputs.BucketLifecycleRuleGetArgs>? _lifecycleRules;

        /// <summary>
        /// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRuleGetArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRuleGetArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The bucket's [Access &amp; Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.
        /// </summary>
        [Input("logging")]
        public Input<Inputs.BucketLoggingGetArgs>? Logging { get; set; }

        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.
        /// </summary>
        [Input("requesterPays")]
        public Input<bool>? RequesterPays { get; set; }

        /// <summary>
        /// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.BucketRetentionPolicyGetArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        /// <summary>
        /// The base URL of the bucket, in the format `gs://&lt;bucket-name&gt;`.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.BucketVersioningGetArgs>? Versioning { get; set; }

        /// <summary>
        /// Configuration if the bucket acts as a website. Structure is documented below.
        /// </summary>
        [Input("website")]
        public Input<Inputs.BucketWebsiteGetArgs>? Website { get; set; }

        public BucketState()
        {
        }
    }
}
