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
    /// WorkloadIdentityPoolProvider can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}/providers/{{workload_identity_pool_provider_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{project}}/{{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider default {{workload_identity_pool_id}}/{{workload_identity_pool_provider_id}}
    /// ```
    /// </summary>
    public partial class WorkloadIdentityPoolProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// * `assertion`: JSON representing the authentication credential issued by the provider.
        /// * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
        /// * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
        /// The maximum length of the attribute condition expression is 4096 characters. If
        /// unspecified, all valid authentication credential are accepted.
        /// The following example shows how to only allow credentials with a mapped `google.groups`
        /// value of `admins`:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        [Output("attributeCondition")]
        public Output<string?> AttributeCondition { get; private set; } = null!;

        /// <summary>
        /// Maps attributes from authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value
        /// in IAM bindings. This is also the subject that appears in Cloud Logging logs.
        /// Cannot exceed 127 characters.
        /// * `google.groups`: Groups the external identity belongs to. You can grant groups
        /// access to resources using an IAM `principalSet` binding; access applies to all
        /// members of the group.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
        /// where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
        /// define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
        /// is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a
        /// workload to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For AWS providers, the following rules apply:
        /// - If no attribute mapping is defined, the following default mapping applies:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// - If any custom attribute mappings are defined, they must include a mapping to the
        /// `google.subject` attribute.
        /// For OIDC providers, the following rules apply:
        /// - Custom attribute mappings must be defined, and must include a mapping to the
        /// `google.subject` attribute. For example, the following maps the `sub` claim of the
        /// incoming credential to the `subject` attribute on a Google token.
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        [Output("attributeMapping")]
        public Output<ImmutableDictionary<string, string>?> AttributeMapping { get; private set; } = null!;

        /// <summary>
        /// An Amazon Web Services identity provider. Not compatible with the property oidc.
        /// Structure is documented below.
        /// </summary>
        [Output("aws")]
        public Output<Outputs.WorkloadIdentityPoolProviderAws?> Aws { get; private set; } = null!;

        /// <summary>
        /// A description for the provider. Cannot exceed 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// A display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The resource name of the provider as
        /// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
        /// Structure is documented below.
        /// </summary>
        [Output("oidc")]
        public Output<Outputs.WorkloadIdentityPoolProviderOidc?> Oidc { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
        /// validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
        /// deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
        /// You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The ID used for the pool, which is the final component of the pool resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Output("workloadIdentityPoolId")]
        public Output<string> WorkloadIdentityPoolId { get; private set; } = null!;

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name. This
        /// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Output("workloadIdentityPoolProviderId")]
        public Output<string> WorkloadIdentityPoolProviderId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkloadIdentityPoolProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadIdentityPoolProvider(string name, WorkloadIdentityPoolProviderArgs args, CustomResourceOptions? options = null)
            : base("gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider", name, args ?? new WorkloadIdentityPoolProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadIdentityPoolProvider(string name, Input<string> id, WorkloadIdentityPoolProviderState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkloadIdentityPoolProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadIdentityPoolProvider Get(string name, Input<string> id, WorkloadIdentityPoolProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkloadIdentityPoolProvider(name, id, state, options);
        }
    }

    public sealed class WorkloadIdentityPoolProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// * `assertion`: JSON representing the authentication credential issued by the provider.
        /// * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
        /// * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
        /// The maximum length of the attribute condition expression is 4096 characters. If
        /// unspecified, all valid authentication credential are accepted.
        /// The following example shows how to only allow credentials with a mapped `google.groups`
        /// value of `admins`:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        [Input("attributeCondition")]
        public Input<string>? AttributeCondition { get; set; }

        [Input("attributeMapping")]
        private InputMap<string>? _attributeMapping;

        /// <summary>
        /// Maps attributes from authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value
        /// in IAM bindings. This is also the subject that appears in Cloud Logging logs.
        /// Cannot exceed 127 characters.
        /// * `google.groups`: Groups the external identity belongs to. You can grant groups
        /// access to resources using an IAM `principalSet` binding; access applies to all
        /// members of the group.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
        /// where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
        /// define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
        /// is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a
        /// workload to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For AWS providers, the following rules apply:
        /// - If no attribute mapping is defined, the following default mapping applies:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// - If any custom attribute mappings are defined, they must include a mapping to the
        /// `google.subject` attribute.
        /// For OIDC providers, the following rules apply:
        /// - Custom attribute mappings must be defined, and must include a mapping to the
        /// `google.subject` attribute. For example, the following maps the `sub` claim of the
        /// incoming credential to the `subject` attribute on a Google token.
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public InputMap<string> AttributeMapping
        {
            get => _attributeMapping ?? (_attributeMapping = new InputMap<string>());
            set => _attributeMapping = value;
        }

        /// <summary>
        /// An Amazon Web Services identity provider. Not compatible with the property oidc.
        /// Structure is documented below.
        /// </summary>
        [Input("aws")]
        public Input<Inputs.WorkloadIdentityPoolProviderAwsArgs>? Aws { get; set; }

        /// <summary>
        /// A description for the provider. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
        /// Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.WorkloadIdentityPoolProviderOidcArgs>? Oidc { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The ID used for the pool, which is the final component of the pool resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolId", required: true)]
        public Input<string> WorkloadIdentityPoolId { get; set; } = null!;

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name. This
        /// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolProviderId", required: true)]
        public Input<string> WorkloadIdentityPoolProviderId { get; set; } = null!;

        public WorkloadIdentityPoolProviderArgs()
        {
        }
    }

    public sealed class WorkloadIdentityPoolProviderState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [A Common Expression Language](https://opensource.google/projects/cel) expression, in
        /// plain text, to restrict what otherwise valid authentication credentials issued by the
        /// provider should not be accepted.
        /// The expression must output a boolean representing whether to allow the federation.
        /// The following keywords may be referenced in the expressions:
        /// * `assertion`: JSON representing the authentication credential issued by the provider.
        /// * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.
        /// * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.
        /// The maximum length of the attribute condition expression is 4096 characters. If
        /// unspecified, all valid authentication credential are accepted.
        /// The following example shows how to only allow credentials with a mapped `google.groups`
        /// value of `admins`:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        [Input("attributeCondition")]
        public Input<string>? AttributeCondition { get; set; }

        [Input("attributeMapping")]
        private InputMap<string>? _attributeMapping;

        /// <summary>
        /// Maps attributes from authentication credentials issued by an external identity provider
        /// to Google Cloud attributes, such as `subject` and `segment`.
        /// Each key must be a string specifying the Google Cloud IAM attribute to map to.
        /// The following keys are supported:
        /// * `google.subject`: The principal IAM is authenticating. You can reference this value
        /// in IAM bindings. This is also the subject that appears in Cloud Logging logs.
        /// Cannot exceed 127 characters.
        /// * `google.groups`: Groups the external identity belongs to. You can grant groups
        /// access to resources using an IAM `principalSet` binding; access applies to all
        /// members of the group.
        /// You can also provide custom attributes by specifying `attribute.{custom_attribute}`,
        /// where `{custom_attribute}` is the name of the custom attribute to be mapped. You can
        /// define a maximum of 50 custom attributes. The maximum length of a mapped attribute key
        /// is 100 characters, and the key may only contain the characters [a-z0-9_].
        /// You can reference these attributes in IAM policies to define fine-grained access for a
        /// workload to Google Cloud resources. For example:
        /// * `google.subject`:
        /// `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}`
        /// * `google.groups`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}`
        /// * `attribute.{custom_attribute}`:
        /// `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}`
        /// Each value must be a [Common Expression Language](https://opensource.google/projects/cel)
        /// function that maps an identity provider credential to the normalized attribute specified
        /// by the corresponding map key.
        /// You can use the `assertion` keyword in the expression to access a JSON representation of
        /// the authentication credential issued by the provider.
        /// The maximum length of an attribute mapping expression is 2048 characters. When evaluated,
        /// the total size of all mapped attributes must not exceed 8KB.
        /// For AWS providers, the following rules apply:
        /// - If no attribute mapping is defined, the following default mapping applies:
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// - If any custom attribute mappings are defined, they must include a mapping to the
        /// `google.subject` attribute.
        /// For OIDC providers, the following rules apply:
        /// - Custom attribute mappings must be defined, and must include a mapping to the
        /// `google.subject` attribute. For example, the following maps the `sub` claim of the
        /// incoming credential to the `subject` attribute on a Google token.
        /// ```csharp
        /// using Pulumi;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public InputMap<string> AttributeMapping
        {
            get => _attributeMapping ?? (_attributeMapping = new InputMap<string>());
            set => _attributeMapping = value;
        }

        /// <summary>
        /// An Amazon Web Services identity provider. Not compatible with the property oidc.
        /// Structure is documented below.
        /// </summary>
        [Input("aws")]
        public Input<Inputs.WorkloadIdentityPoolProviderAwsGetArgs>? Aws { get; set; }

        /// <summary>
        /// A description for the provider. Cannot exceed 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the provider is disabled. You cannot use a disabled provider to exchange tokens.
        /// However, existing tokens still grant access.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// A display name for the provider. Cannot exceed 32 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The resource name of the provider as
        /// 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}/providers/{workload_identity_pool_provider_id}'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An OpenId Connect 1.0 identity provider. Not compatible with the property aws.
        /// Structure is documented below.
        /// </summary>
        [Input("oidc")]
        public Input<Inputs.WorkloadIdentityPoolProviderOidcGetArgs>? Oidc { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The state of the provider. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The provider is active, and may be used to
        /// validate authentication credentials. * DELETED: The provider is soft-deleted. Soft-deleted providers are permanently
        /// deleted after approximately 30 days. You can restore a soft-deleted provider using UndeleteWorkloadIdentityPoolProvider.
        /// You cannot reuse the ID of a soft-deleted provider until it is permanently deleted.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The ID used for the pool, which is the final component of the pool resource name. This
        /// value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolId")]
        public Input<string>? WorkloadIdentityPoolId { get; set; }

        /// <summary>
        /// The ID for the provider, which becomes the final component of the resource name. This
        /// value must be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        /// `gcp-` is reserved for use by Google, and may not be specified.
        /// </summary>
        [Input("workloadIdentityPoolProviderId")]
        public Input<string>? WorkloadIdentityPoolProviderId { get; set; }

        public WorkloadIdentityPoolProviderState()
        {
        }
    }
}
