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
    /// Three different resources help you manage your IAM policy for Cloud Storage Bucket. Each of these resources serves a different use case:
    /// 
    /// * `gcp.storage.BucketIAMPolicy`: Authoritative. Sets the IAM policy for the bucket and replaces any existing policy already attached.
    /// * `gcp.storage.BucketIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the bucket are preserved.
    /// * `gcp.storage.BucketIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the bucket are preserved.
    /// 
    /// &gt; **Note:** `gcp.storage.BucketIAMPolicy` **cannot** be used in conjunction with `gcp.storage.BucketIAMBinding` and `gcp.storage.BucketIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.storage.BucketIAMBinding` resources **can be** used in conjunction with `gcp.storage.BucketIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_storage\_bucket\_iam\_policy
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Role = "roles/storage.admin",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.Storage.BucketIAMPolicy("policy", new Gcp.Storage.BucketIAMPolicyArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var admin = Output.Create(Gcp.Organizations.GetIAMPolicy.InvokeAsync(new Gcp.Organizations.GetIAMPolicyArgs
    ///         {
    ///             Bindings = 
    ///             {
    ///                 new Gcp.Organizations.Inputs.GetIAMPolicyBindingArgs
    ///                 {
    ///                     Role = "roles/storage.admin",
    ///                     Members = 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     },
    ///                     Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionArgs
    ///                     {
    ///                         Title = "expires_after_2019_12_31",
    ///                         Description = "Expiring at midnight of 2019-12-31",
    ///                         Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.Storage.BucketIAMPolicy("policy", new Gcp.Storage.BucketIAMPolicyArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## google\_storage\_bucket\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.Storage.BucketIAMBinding("binding", new Gcp.Storage.BucketIAMBindingArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             Role = "roles/storage.admin",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.Storage.BucketIAMBinding("binding", new Gcp.Storage.BucketIAMBindingArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             Role = "roles/storage.admin",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Condition = new Gcp.Storage.Inputs.BucketIAMBindingConditionArgs
    ///             {
    ///                 Title = "expires_after_2019_12_31",
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## google\_storage\_bucket\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.Storage.BucketIAMMember("member", new Gcp.Storage.BucketIAMMemberArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             Role = "roles/storage.admin",
    ///             Member = "user:jane@example.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// With IAM Conditions:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.Storage.BucketIAMMember("member", new Gcp.Storage.BucketIAMMemberArgs
    ///         {
    ///             Bucket = google_storage_bucket.Default.Name,
    ///             Role = "roles/storage.admin",
    ///             Member = "user:jane@example.com",
    ///             Condition = new Gcp.Storage.Inputs.BucketIAMMemberConditionArgs
    ///             {
    ///                 Title = "expires_after_2019_12_31",
    ///                 Description = "Expiring at midnight of 2019-12-31",
    ///                 Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
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
