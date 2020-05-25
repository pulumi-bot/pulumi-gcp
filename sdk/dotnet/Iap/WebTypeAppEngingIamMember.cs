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
    /// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:
    /// 
    /// * `gcp.iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
    /// * `gcp.iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
    /// * `gcp.iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.
    /// 
    /// &gt; **Note:** `gcp.iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `gcp.iap.WebTypeAppEngingIamBinding` and `gcp.iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `gcp.iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// 
    /// 
    /// ## google\_iap\_web\_type\_app\_engine\_iam\_policy
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
    ///         var policy = new Gcp.Iap.WebTypeAppEngingIamPolicy("policy", new Gcp.Iap.WebTypeAppEngingIamPolicyArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
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
    ///         var policy = new Gcp.Iap.WebTypeAppEngingIamPolicy("policy", new Gcp.Iap.WebTypeAppEngingIamPolicyArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
    ///             PolicyData = admin.Apply(admin =&gt; admin.PolicyData),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## google\_iap\_web\_type\_app\_engine\_iam\_binding
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var binding = new Gcp.Iap.WebTypeAppEngingIamBinding("binding", new Gcp.Iap.WebTypeAppEngingIamBindingArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
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
    ///         var binding = new Gcp.Iap.WebTypeAppEngingIamBinding("binding", new Gcp.Iap.WebTypeAppEngingIamBindingArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
    ///             Role = "roles/iap.httpsResourceAccessor",
    ///             Members = 
    ///             {
    ///                 "user:jane@example.com",
    ///             },
    ///             Condition = new Gcp.Iap.Inputs.WebTypeAppEngingIamBindingConditionArgs
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
    /// ## google\_iap\_web\_type\_app\_engine\_iam\_member
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var member = new Gcp.Iap.WebTypeAppEngingIamMember("member", new Gcp.Iap.WebTypeAppEngingIamMemberArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
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
    ///         var member = new Gcp.Iap.WebTypeAppEngingIamMember("member", new Gcp.Iap.WebTypeAppEngingIamMemberArgs
    ///         {
    ///             Project = google_app_engine_application.App.Project,
    ///             AppId = google_app_engine_application.App.App_id,
    ///             Role = "roles/iap.httpsResourceAccessor",
    ///             Member = "user:jane@example.com",
    ///             Condition = new Gcp.Iap.Inputs.WebTypeAppEngingIamMemberConditionArgs
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
    public partial class WebTypeAppEngingIamMember : Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.WebTypeAppEngingIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a WebTypeAppEngingIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebTypeAppEngingIamMember(string name, WebTypeAppEngingIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember", name, args ?? new WebTypeAppEngingIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebTypeAppEngingIamMember(string name, Input<string> id, WebTypeAppEngingIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iap/webTypeAppEngingIamMember:WebTypeAppEngingIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebTypeAppEngingIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebTypeAppEngingIamMember Get(string name, Input<string> id, WebTypeAppEngingIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new WebTypeAppEngingIamMember(name, id, state, options);
        }
    }

    public sealed class WebTypeAppEngingIamMemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        /// <summary>
        /// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.WebTypeAppEngingIamMemberConditionArgs>? Condition { get; set; }

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public WebTypeAppEngingIamMemberArgs()
        {
        }
    }

    public sealed class WebTypeAppEngingIamMemberState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.WebTypeAppEngingIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.iap.WebTypeAppEngingIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public WebTypeAppEngingIamMemberState()
        {
        }
    }
}
