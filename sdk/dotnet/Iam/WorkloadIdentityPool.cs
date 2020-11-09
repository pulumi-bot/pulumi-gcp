// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam
{
    /// <summary>
    /// ## Import
    /// 
    /// WorkloadIdentityPool can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{project}}/{{workload_identity_pool_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{workload_identity_pool_id}}
    /// ```
    /// </summary>
    public partial class WorkloadIdentityPool : Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the pool. Cannot exceed 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        /// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        /// access again.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// A display name for the pool. Cannot exceed 32 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The resource name of the pool as 'projects/&lt;projectnumber&gt;/locations/global/workloadIdentityPools/&lt;id&gt;'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The state of the pool. STATE_UNSPECIFIED: State unspecified. ACTIVE: The pool is active, and may be used in Google Cloud
        /// policies. DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30 days. You
        /// can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until
        /// it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing tokens to
        /// access resources. If the pool is undeleted, existing tokens grant access again.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Output("workloadIdentityPoolId")]
        public Output<string> WorkloadIdentityPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkloadIdentityPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadIdentityPool(string name, WorkloadIdentityPoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:iam/workloadIdentityPool:WorkloadIdentityPool", name, args ?? new WorkloadIdentityPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadIdentityPool(string name, Input<string> id, WorkloadIdentityPoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iam/workloadIdentityPool:WorkloadIdentityPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkloadIdentityPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadIdentityPool Get(string name, Input<string> id, WorkloadIdentityPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkloadIdentityPool(name, id, state, options);
        }
    }

    public sealed class WorkloadIdentityPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the pool. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        /// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        /// access again.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A display name for the pool. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolId", required: true)]
        public Input<string> WorkloadIdentityPoolId { get; set; } = null!;

        public WorkloadIdentityPoolArgs()
        {
        }
    }

    public sealed class WorkloadIdentityPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the pool. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        /// existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        /// access again.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A display name for the pool. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource name of the pool as 'projects/&lt;projectnumber&gt;/locations/global/workloadIdentityPools/&lt;id&gt;'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The state of the pool. STATE_UNSPECIFIED: State unspecified. ACTIVE: The pool is active, and may be used in Google Cloud
        /// policies. DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30 days. You
        /// can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted pool until
        /// it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing tokens to
        /// access resources. If the pool is undeleted, existing tokens grant access again.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The ID to use for the pool, which becomes the final component of the resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolId")]
        public Input<string>? WorkloadIdentityPoolId { get; set; }

        public WorkloadIdentityPoolState()
        {
        }
    }
}
