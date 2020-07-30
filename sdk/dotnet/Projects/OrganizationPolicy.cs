// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Projects
{
    /// <summary>
    /// Allows management of Organization policies for a Google Project. For more information see
    /// [the official
    /// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
    /// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setOrgPolicy).
    /// 
    /// ## Example Usage
    /// 
    /// To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var serialPortPolicy = new Gcp.Projects.OrganizationPolicy("serialPortPolicy", new Gcp.Projects.OrganizationPolicyArgs
    ///         {
    ///             BooleanPolicy = new Gcp.Projects.Inputs.OrganizationPolicyBooleanPolicyArgs
    ///             {
    ///                 Enforced = true,
    ///             },
    ///             Constraint = "compute.disableSerialPortAccess",
    ///             Project = "your-project-id",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var servicesPolicy = new Gcp.Projects.OrganizationPolicy("servicesPolicy", new Gcp.Projects.OrganizationPolicyArgs
    ///         {
    ///             Constraint = "serviceuser.services",
    ///             ListPolicy = new Gcp.Projects.Inputs.OrganizationPolicyListPolicyArgs
    ///             {
    ///                 Allow = new Gcp.Projects.Inputs.OrganizationPolicyListPolicyAllowArgs
    ///                 {
    ///                     All = true,
    ///                 },
    ///             },
    ///             Project = "your-project-id",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Or to deny some services, use the following instead:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var servicesPolicy = new Gcp.Projects.OrganizationPolicy("servicesPolicy", new Gcp.Projects.OrganizationPolicyArgs
    ///         {
    ///             Constraint = "serviceuser.services",
    ///             ListPolicy = new Gcp.Projects.Inputs.OrganizationPolicyListPolicyArgs
    ///             {
    ///                 Deny = new Gcp.Projects.Inputs.OrganizationPolicyListPolicyDenyArgs
    ///                 {
    ///                     Values = 
    ///                     {
    ///                         "cloudresourcemanager.googleapis.com",
    ///                     },
    ///                 },
    ///                 SuggestedValue = "compute.googleapis.com",
    ///             },
    ///             Project = "your-project-id",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// To restore the default project organization policy, use the following instead:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var servicesPolicy = new Gcp.Projects.OrganizationPolicy("servicesPolicy", new Gcp.Projects.OrganizationPolicyArgs
    ///         {
    ///             Constraint = "serviceuser.services",
    ///             Project = "your-project-id",
    ///             RestorePolicy = new Gcp.Projects.Inputs.OrganizationPolicyRestorePolicyArgs
    ///             {
    ///                 Default = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class OrganizationPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        /// </summary>
        [Output("booleanPolicy")]
        public Output<Outputs.OrganizationPolicyBooleanPolicy?> BooleanPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        /// </summary>
        [Output("constraint")]
        public Output<string> Constraint { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        /// </summary>
        [Output("listPolicy")]
        public Output<Outputs.OrganizationPolicyListPolicy?> ListPolicy { get; private set; } = null!;

        /// <summary>
        /// The project id of the project to set the policy for.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A restore policy is a constraint to restore the default policy. Structure is documented below.
        /// </summary>
        [Output("restorePolicy")]
        public Output<Outputs.OrganizationPolicyRestorePolicy?> RestorePolicy { get; private set; } = null!;

        /// <summary>
        /// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Version of the Policy. Default version is 0.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationPolicy(string name, OrganizationPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:projects/organizationPolicy:OrganizationPolicy", name, args ?? new OrganizationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationPolicy(string name, Input<string> id, OrganizationPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:projects/organizationPolicy:OrganizationPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationPolicy Get(string name, Input<string> id, OrganizationPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationPolicy(name, id, state, options);
        }
    }

    public sealed class OrganizationPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        /// </summary>
        [Input("booleanPolicy")]
        public Input<Inputs.OrganizationPolicyBooleanPolicyArgs>? BooleanPolicy { get; set; }

        /// <summary>
        /// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        /// </summary>
        [Input("constraint", required: true)]
        public Input<string> Constraint { get; set; } = null!;

        /// <summary>
        /// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        /// </summary>
        [Input("listPolicy")]
        public Input<Inputs.OrganizationPolicyListPolicyArgs>? ListPolicy { get; set; }

        /// <summary>
        /// The project id of the project to set the policy for.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// A restore policy is a constraint to restore the default policy. Structure is documented below.
        /// </summary>
        [Input("restorePolicy")]
        public Input<Inputs.OrganizationPolicyRestorePolicyArgs>? RestorePolicy { get; set; }

        /// <summary>
        /// Version of the Policy. Default version is 0.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public OrganizationPolicyArgs()
        {
        }
    }

    public sealed class OrganizationPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        /// </summary>
        [Input("booleanPolicy")]
        public Input<Inputs.OrganizationPolicyBooleanPolicyGetArgs>? BooleanPolicy { get; set; }

        /// <summary>
        /// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        /// </summary>
        [Input("constraint")]
        public Input<string>? Constraint { get; set; }

        /// <summary>
        /// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        /// </summary>
        [Input("listPolicy")]
        public Input<Inputs.OrganizationPolicyListPolicyGetArgs>? ListPolicy { get; set; }

        /// <summary>
        /// The project id of the project to set the policy for.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A restore policy is a constraint to restore the default policy. Structure is documented below.
        /// </summary>
        [Input("restorePolicy")]
        public Input<Inputs.OrganizationPolicyRestorePolicyGetArgs>? RestorePolicy { get; set; }

        /// <summary>
        /// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Version of the Policy. Default version is 0.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public OrganizationPolicyState()
        {
        }
    }
}
