// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iap
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
    /// 
    /// * `gcp.iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
    /// * `gcp.iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
    /// * `gcp.iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
    /// 
    /// &gt; **Note:** `gcp.iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `gcp.iap.WebBackendServiceIamBinding` and `gcp.iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `gcp.iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_iap\_web\_backend\_service\_iam\_policy
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
    ///             Binding = 
    ///             {
    ///                 
    ///                 {
    ///                     { "role", "roles/iap.httpsResourceAccessor" },
    ///                     { "members", 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     } },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.Iap.WebBackendServiceIamPolicy("policy", new Gcp.Iap.WebBackendServiceIamPolicyArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
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
    ///             Binding = 
    ///             {
    ///                 
    ///                 {
    ///                     { "role", "roles/iap.httpsResourceAccessor" },
    ///                     { "members", 
    ///                     {
    ///                         "user:jane@example.com",
    ///                     } },
    ///                     { "condition", 
    ///                     {
    ///                         { "title", "expires_after_2019_12_31" },
    ///                         { "description", "Expiring at midnight of 2019-12-31" },
    ///                         { "expression", "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")" },
    ///                     } },
    ///                 },
    ///             },
    ///         }));
    ///         var policy = new Gcp.Iap.WebBackendServiceIamPolicy("policy", new Gcp.Iap.WebBackendServiceIamPolicyArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## google\_iap\_web\_backend\_service\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.Iap.WebBackendServiceIamBinding("binding", new Gcp.Iap.WebBackendServiceIamBindingArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
    ///             Role = "roles/iap.httpsResourceAccessor",
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
    ///         var binding = new Gcp.Iap.WebBackendServiceIamBinding("binding", new Gcp.Iap.WebBackendServiceIamBindingArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
    ///             Role = "roles/iap.httpsResourceAccessor",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Condition = new Gcp.Iap.Inputs.WebBackendServiceIamBindingConditionArgs
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
    /// ## google\_iap\_web\_backend\_service\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.Iap.WebBackendServiceIamMember("member", new Gcp.Iap.WebBackendServiceIamMemberArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
    ///             Role = "roles/iap.httpsResourceAccessor",
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
    ///         var member = new Gcp.Iap.WebBackendServiceIamMember("member", new Gcp.Iap.WebBackendServiceIamMemberArgs
    ///         {
    ///             Project = google_compute_backend_service.Default.Project,
    ///             WebBackendService = google_compute_backend_service.Default.Name,
    ///             Role = "roles/iap.httpsResourceAccessor",
    ///             Member = "user:jane@example.com",
    ///             Condition = new Gcp.Iap.Inputs.WebBackendServiceIamMemberConditionArgs
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
    public partial class WebBackendServiceIamPolicy : Pulumi.CustomResource
    {
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
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("webBackendService")]
        public Output<string> WebBackendService { get; private set; } = null!;


        /// <summary>
        /// Create a WebBackendServiceIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebBackendServiceIamPolicy(string name, WebBackendServiceIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, args ?? new WebBackendServiceIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebBackendServiceIamPolicy(string name, Input<string> id, WebBackendServiceIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebBackendServiceIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebBackendServiceIamPolicy Get(string name, Input<string> id, WebBackendServiceIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new WebBackendServiceIamPolicy(name, id, state, options);
        }
    }

    public sealed class WebBackendServiceIamPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("webBackendService", required: true)]
        public Input<string> WebBackendService { get; set; } = null!;

        public WebBackendServiceIamPolicyArgs()
        {
        }
    }

    public sealed class WebBackendServiceIamPolicyState : Pulumi.ResourceArgs
    {
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

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("webBackendService")]
        public Input<string>? WebBackendService { get; set; }

        public WebBackendServiceIamPolicyState()
        {
        }
    }
}
