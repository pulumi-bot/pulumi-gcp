// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine
{
    /// <summary>
    /// A domain serving an App Engine application.
    /// 
    /// 
    /// To get more information about DomainMapping, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.domainMappings)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/mapping-custom-domains)
    /// 
    /// ## Example Usage - App Engine Domain Mapping Basic
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var domainMapping = new Gcp.AppEngine.DomainMapping("domainMapping", new Gcp.AppEngine.DomainMappingArgs
    ///         {
    ///             DomainName = "verified-domain.com",
    ///             SslSettings = new Gcp.AppEngine.Inputs.DomainMappingSslSettingsArgs
    ///             {
    ///                 SslManagementType = "AUTOMATIC",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DomainMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// Relative name of the domain serving the application. Example: example.com.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the domain creation should override any existing mappings for this domain.
        /// By default, overrides are rejected.
        /// </summary>
        [Output("overrideStrategy")]
        public Output<string?> OverrideStrategy { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
        /// configuration in order to serve the application via this domain mapping.
        /// </summary>
        [Output("resourceRecords")]
        public Output<ImmutableArray<Outputs.DomainMappingResourceRecord>> ResourceRecords { get; private set; } = null!;

        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.  Structure is documented below.
        /// </summary>
        [Output("sslSettings")]
        public Output<Outputs.DomainMappingSslSettings?> SslSettings { get; private set; } = null!;


        /// <summary>
        /// Create a DomainMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainMapping(string name, DomainMappingArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/domainMapping:DomainMapping", name, args ?? new DomainMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainMapping(string name, Input<string> id, DomainMappingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/domainMapping:DomainMapping", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainMapping Get(string name, Input<string> id, DomainMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainMapping(name, id, state, options);
        }
    }

    public sealed class DomainMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Relative name of the domain serving the application. Example: example.com.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Whether the domain creation should override any existing mappings for this domain.
        /// By default, overrides are rejected.
        /// </summary>
        [Input("overrideStrategy")]
        public Input<string>? OverrideStrategy { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.  Structure is documented below.
        /// </summary>
        [Input("sslSettings")]
        public Input<Inputs.DomainMappingSslSettingsArgs>? SslSettings { get; set; }

        public DomainMappingArgs()
        {
        }
    }

    public sealed class DomainMappingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Relative name of the domain serving the application. Example: example.com.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the domain creation should override any existing mappings for this domain.
        /// By default, overrides are rejected.
        /// </summary>
        [Input("overrideStrategy")]
        public Input<string>? OverrideStrategy { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("resourceRecords")]
        private InputList<Inputs.DomainMappingResourceRecordGetArgs>? _resourceRecords;

        /// <summary>
        /// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
        /// configuration in order to serve the application via this domain mapping.
        /// </summary>
        public InputList<Inputs.DomainMappingResourceRecordGetArgs> ResourceRecords
        {
            get => _resourceRecords ?? (_resourceRecords = new InputList<Inputs.DomainMappingResourceRecordGetArgs>());
            set => _resourceRecords = value;
        }

        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.  Structure is documented below.
        /// </summary>
        [Input("sslSettings")]
        public Input<Inputs.DomainMappingSslSettingsGetArgs>? SslSettings { get; set; }

        public DomainMappingState()
        {
        }
    }
}
