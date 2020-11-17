// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Organization security policies are used to control incoming/outgoing traffic.
    /// 
    /// To get more information about OrganizationSecurityPolicy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies)
    /// * How-to Guides
    ///     * [Creating a firewall policy](https://cloud.google.com/vpc/docs/using-firewall-policies#create-policy)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// OrganizationSecurityPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default locations/global/securityPolicies/{{policy_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default {{policy_id}}
    /// ```
    /// </summary>
    public partial class OrganizationSecurityPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A textual description for the organization security policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A textual name of the security policy.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        /// Format: organizations/{organization_id} or folders/{folder_id}
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The type indicates the intended use of the security policy.
        /// For organization security policies, the only supported type
        /// is "FIREWALL".
        /// Default value is `FIREWALL`.
        /// Possible values are `FIREWALL`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSecurityPolicy(string name, OrganizationSecurityPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy", name, args ?? new OrganizationSecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSecurityPolicy(string name, Input<string> id, OrganizationSecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSecurityPolicy Get(string name, Input<string> id, OrganizationSecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationSecurityPolicy(name, id, state, options);
        }
    }

    public sealed class OrganizationSecurityPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A textual description for the organization security policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A textual name of the security policy.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        /// Format: organizations/{organization_id} or folders/{folder_id}
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// The type indicates the intended use of the security policy.
        /// For organization security policies, the only supported type
        /// is "FIREWALL".
        /// Default value is `FIREWALL`.
        /// Possible values are `FIREWALL`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OrganizationSecurityPolicyArgs()
        {
        }
    }

    public sealed class OrganizationSecurityPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A textual description for the organization security policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A textual name of the security policy.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Fingerprint of this resource. This field is used internally during updates of this resource.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        /// Format: organizations/{organization_id} or folders/{folder_id}
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// The unique identifier for the resource. This identifier is defined by the server.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The type indicates the intended use of the security policy.
        /// For organization security policies, the only supported type
        /// is "FIREWALL".
        /// Default value is `FIREWALL`.
        /// Possible values are `FIREWALL`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public OrganizationSecurityPolicyState()
        {
        }
    }
}
