// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceDirectory
{
    /// <summary>
    /// An individual service. A service contains a name and optional metadata.
    /// 
    /// To get more information about Service, see:
    /// 
    /// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services)
    /// * How-to Guides
    ///     * [Configuring a service](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_a_service)
    /// 
    /// ## Example Usage - Service Directory Service Basic
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
    ///         var exampleNamespace = new Gcp.ServiceDirectory.Namespace("exampleNamespace", new Gcp.ServiceDirectory.NamespaceArgs
    ///         {
    ///             NamespaceId = "example-namespace",
    ///             Location = "us-central1",
    ///         });
    ///         var exampleService = new Gcp.ServiceDirectory.Service("exampleService", new Gcp.ServiceDirectory.ServiceArgs
    ///         {
    ///             ServiceId = "example-service",
    ///             @namespace = exampleNamespace.Id,
    ///             Metadata = 
    ///             {
    ///                 { "stage", "prod" },
    ///                 { "region", "us-central1" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// Metadata for the service. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 2000 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource name of the namespace this service will belong to.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/service:Service", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata for the service. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 2000 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The resource name of the namespace this service will belong to.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata for the service. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 2000 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The resource name for the service in the format 'projects/*/locations/*/namespaces/*/services/*'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource name of the namespace this service will belong to.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public ServiceState()
        {
        }
    }
}
