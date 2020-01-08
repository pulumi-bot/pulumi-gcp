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
    /// using the [`gcp.storage.BucketACL` resource](https://www.terraform.io/docs/providers/google/r/storage_bucket_acl.html).
    /// 
    /// For more information see
    /// [the official documentation](https://cloud.google.com/storage/docs/overview)
    /// and
    /// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
    /// 
    /// **Note**: If the project id is not set on the resource or in the provider block it will be dynamically
    /// determined which will require enabling the compute api.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_bucket.html.markdown.
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
        public Output<ImmutableArray<Outputs.BucketCors>> Cors { get; private set; } = null!;

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.BucketEncryption?> Encryption { get; private set; } = null!;

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, this provider will fail that run.
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
        public Output<ImmutableArray<Outputs.BucketLifecycleRules>> LifecycleRules { get; private set; } = null!;

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
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
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
            : base("gcp:storage/bucket:Bucket", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        private InputList<Inputs.BucketCorsArgs>? _cors;

        /// <summary>
        /// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketCorsArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.BucketCorsArgs>());
            set => _cors = value;
        }

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.BucketEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, this provider will fail that run.
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
        private InputList<Inputs.BucketLifecycleRulesArgs>? _lifecycleRules;

        /// <summary>
        /// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRulesArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRulesArgs>());
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
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
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
        private InputList<Inputs.BucketCorsGetArgs>? _cors;

        /// <summary>
        /// The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketCorsGetArgs> Cors
        {
            get => _cors ?? (_cors = new InputList<Inputs.BucketCorsGetArgs>());
            set => _cors = value;
        }

        /// <summary>
        /// The bucket's encryption configuration.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.BucketEncryptionGetArgs>? Encryption { get; set; }

        /// <summary>
        /// When deleting a bucket, this
        /// boolean option will delete all contained objects. If you try to delete a
        /// bucket that contains objects, this provider will fail that run.
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
        private InputList<Inputs.BucketLifecycleRulesGetArgs>? _lifecycleRules;

        /// <summary>
        /// The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.
        /// </summary>
        public InputList<Inputs.BucketLifecycleRulesGetArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.BucketLifecycleRulesGetArgs>());
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
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
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

    namespace Inputs
    {

    public sealed class BucketCorsArgs : Pulumi.ResourceArgs
    {
        [Input("maxAgeSeconds")]
        public Input<int>? MaxAgeSeconds { get; set; }

        [Input("methods")]
        private InputList<string>? _methods;
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        [Input("origins")]
        private InputList<string>? _origins;
        public InputList<string> Origins
        {
            get => _origins ?? (_origins = new InputList<string>());
            set => _origins = value;
        }

        [Input("responseHeaders")]
        private InputList<string>? _responseHeaders;
        public InputList<string> ResponseHeaders
        {
            get => _responseHeaders ?? (_responseHeaders = new InputList<string>());
            set => _responseHeaders = value;
        }

        public BucketCorsArgs()
        {
        }
    }

    public sealed class BucketCorsGetArgs : Pulumi.ResourceArgs
    {
        [Input("maxAgeSeconds")]
        public Input<int>? MaxAgeSeconds { get; set; }

        [Input("methods")]
        private InputList<string>? _methods;
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        [Input("origins")]
        private InputList<string>? _origins;
        public InputList<string> Origins
        {
            get => _origins ?? (_origins = new InputList<string>());
            set => _origins = value;
        }

        [Input("responseHeaders")]
        private InputList<string>? _responseHeaders;
        public InputList<string> ResponseHeaders
        {
            get => _responseHeaders ?? (_responseHeaders = new InputList<string>());
            set => _responseHeaders = value;
        }

        public BucketCorsGetArgs()
        {
        }
    }

    public sealed class BucketEncryptionArgs : Pulumi.ResourceArgs
    {
        [Input("defaultKmsKeyName", required: true)]
        public Input<string> DefaultKmsKeyName { get; set; } = null!;

        public BucketEncryptionArgs()
        {
        }
    }

    public sealed class BucketEncryptionGetArgs : Pulumi.ResourceArgs
    {
        [Input("defaultKmsKeyName", required: true)]
        public Input<string> DefaultKmsKeyName { get; set; } = null!;

        public BucketEncryptionGetArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BucketLifecycleRulesActionArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesActionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BucketLifecycleRulesActionGetArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<BucketLifecycleRulesActionArgs> Action { get; set; } = null!;

        [Input("condition", required: true)]
        public Input<BucketLifecycleRulesConditionArgs> Condition { get; set; } = null!;

        public BucketLifecycleRulesArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesConditionArgs : Pulumi.ResourceArgs
    {
        [Input("age")]
        public Input<int>? Age { get; set; }

        [Input("createdBefore")]
        public Input<string>? CreatedBefore { get; set; }

        [Input("matchesStorageClasses")]
        private InputList<string>? _matchesStorageClasses;
        public InputList<string> MatchesStorageClasses
        {
            get => _matchesStorageClasses ?? (_matchesStorageClasses = new InputList<string>());
            set => _matchesStorageClasses = value;
        }

        [Input("numNewerVersions")]
        public Input<int>? NumNewerVersions { get; set; }

        [Input("withState")]
        public Input<string>? WithState { get; set; }

        public BucketLifecycleRulesConditionArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesConditionGetArgs : Pulumi.ResourceArgs
    {
        [Input("age")]
        public Input<int>? Age { get; set; }

        [Input("createdBefore")]
        public Input<string>? CreatedBefore { get; set; }

        [Input("matchesStorageClasses")]
        private InputList<string>? _matchesStorageClasses;
        public InputList<string> MatchesStorageClasses
        {
            get => _matchesStorageClasses ?? (_matchesStorageClasses = new InputList<string>());
            set => _matchesStorageClasses = value;
        }

        [Input("numNewerVersions")]
        public Input<int>? NumNewerVersions { get; set; }

        [Input("withState")]
        public Input<string>? WithState { get; set; }

        public BucketLifecycleRulesConditionGetArgs()
        {
        }
    }

    public sealed class BucketLifecycleRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<BucketLifecycleRulesActionGetArgs> Action { get; set; } = null!;

        [Input("condition", required: true)]
        public Input<BucketLifecycleRulesConditionGetArgs> Condition { get; set; } = null!;

        public BucketLifecycleRulesGetArgs()
        {
        }
    }

    public sealed class BucketLoggingArgs : Pulumi.ResourceArgs
    {
        [Input("logBucket", required: true)]
        public Input<string> LogBucket { get; set; } = null!;

        [Input("logObjectPrefix")]
        public Input<string>? LogObjectPrefix { get; set; }

        public BucketLoggingArgs()
        {
        }
    }

    public sealed class BucketLoggingGetArgs : Pulumi.ResourceArgs
    {
        [Input("logBucket", required: true)]
        public Input<string> LogBucket { get; set; } = null!;

        [Input("logObjectPrefix")]
        public Input<string>? LogObjectPrefix { get; set; }

        public BucketLoggingGetArgs()
        {
        }
    }

    public sealed class BucketRetentionPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("isLocked")]
        public Input<bool>? IsLocked { get; set; }

        [Input("retentionPeriod", required: true)]
        public Input<int> RetentionPeriod { get; set; } = null!;

        public BucketRetentionPolicyArgs()
        {
        }
    }

    public sealed class BucketRetentionPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("isLocked")]
        public Input<bool>? IsLocked { get; set; }

        [Input("retentionPeriod", required: true)]
        public Input<int> RetentionPeriod { get; set; } = null!;

        public BucketRetentionPolicyGetArgs()
        {
        }
    }

    public sealed class BucketVersioningArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BucketVersioningArgs()
        {
        }
    }

    public sealed class BucketVersioningGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BucketVersioningGetArgs()
        {
        }
    }

    public sealed class BucketWebsiteArgs : Pulumi.ResourceArgs
    {
        [Input("mainPageSuffix")]
        public Input<string>? MainPageSuffix { get; set; }

        [Input("notFoundPage")]
        public Input<string>? NotFoundPage { get; set; }

        public BucketWebsiteArgs()
        {
        }
    }

    public sealed class BucketWebsiteGetArgs : Pulumi.ResourceArgs
    {
        [Input("mainPageSuffix")]
        public Input<string>? MainPageSuffix { get; set; }

        [Input("notFoundPage")]
        public Input<string>? NotFoundPage { get; set; }

        public BucketWebsiteGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class BucketCors
    {
        public readonly int? MaxAgeSeconds;
        public readonly ImmutableArray<string> Methods;
        public readonly ImmutableArray<string> Origins;
        public readonly ImmutableArray<string> ResponseHeaders;

        [OutputConstructor]
        private BucketCors(
            int? maxAgeSeconds,
            ImmutableArray<string> methods,
            ImmutableArray<string> origins,
            ImmutableArray<string> responseHeaders)
        {
            MaxAgeSeconds = maxAgeSeconds;
            Methods = methods;
            Origins = origins;
            ResponseHeaders = responseHeaders;
        }
    }

    [OutputType]
    public sealed class BucketEncryption
    {
        public readonly string DefaultKmsKeyName;

        [OutputConstructor]
        private BucketEncryption(string defaultKmsKeyName)
        {
            DefaultKmsKeyName = defaultKmsKeyName;
        }
    }

    [OutputType]
    public sealed class BucketLifecycleRules
    {
        public readonly BucketLifecycleRulesAction Action;
        public readonly BucketLifecycleRulesCondition Condition;

        [OutputConstructor]
        private BucketLifecycleRules(
            BucketLifecycleRulesAction action,
            BucketLifecycleRulesCondition condition)
        {
            Action = action;
            Condition = condition;
        }
    }

    [OutputType]
    public sealed class BucketLifecycleRulesAction
    {
        /// <summary>
        /// The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: `STANDARD`, `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`.
        /// </summary>
        public readonly string? StorageClass;
        public readonly string Type;

        [OutputConstructor]
        private BucketLifecycleRulesAction(
            string? storageClass,
            string type)
        {
            StorageClass = storageClass;
            Type = type;
        }
    }

    [OutputType]
    public sealed class BucketLifecycleRulesCondition
    {
        public readonly int? Age;
        public readonly string? CreatedBefore;
        public readonly ImmutableArray<string> MatchesStorageClasses;
        public readonly int? NumNewerVersions;
        public readonly string WithState;

        [OutputConstructor]
        private BucketLifecycleRulesCondition(
            int? age,
            string? createdBefore,
            ImmutableArray<string> matchesStorageClasses,
            int? numNewerVersions,
            string withState)
        {
            Age = age;
            CreatedBefore = createdBefore;
            MatchesStorageClasses = matchesStorageClasses;
            NumNewerVersions = numNewerVersions;
            WithState = withState;
        }
    }

    [OutputType]
    public sealed class BucketLogging
    {
        public readonly string LogBucket;
        public readonly string LogObjectPrefix;

        [OutputConstructor]
        private BucketLogging(
            string logBucket,
            string logObjectPrefix)
        {
            LogBucket = logBucket;
            LogObjectPrefix = logObjectPrefix;
        }
    }

    [OutputType]
    public sealed class BucketRetentionPolicy
    {
        public readonly bool? IsLocked;
        public readonly int RetentionPeriod;

        [OutputConstructor]
        private BucketRetentionPolicy(
            bool? isLocked,
            int retentionPeriod)
        {
            IsLocked = isLocked;
            RetentionPeriod = retentionPeriod;
        }
    }

    [OutputType]
    public sealed class BucketVersioning
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private BucketVersioning(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class BucketWebsite
    {
        public readonly string? MainPageSuffix;
        public readonly string? NotFoundPage;

        [OutputConstructor]
        private BucketWebsite(
            string? mainPageSuffix,
            string? notFoundPage)
        {
            MainPageSuffix = mainPageSuffix;
            NotFoundPage = notFoundPage;
        }
    }
    }
}
