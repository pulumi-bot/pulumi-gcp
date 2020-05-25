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
    /// An individual endpoint that provides a service.
    /// 
    /// To get more information about Endpoint, see:
    /// 
    /// * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services.endpoints)
    /// * How-to Guides
    ///     * [Configuring an endpoint](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_an_endpoint)
    /// 
    /// ## Example Usage - Service Directory Endpoint Basic
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
    ///         });
    ///         var exampleEndpoint = new Gcp.ServiceDirectory.Endpoint("exampleEndpoint", new Gcp.ServiceDirectory.EndpointArgs
    ///         {
    ///             EndpointId = "example-endpoint",
    ///             Service = exampleService.Id,
    ///             Metadata = 
    ///             {
    ///                 { "stage", "prod" },
    ///                 { "region", "us-central1" },
    ///             },
    ///             Address = "1.2.3.4",
    ///             Port = 5353,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// IPv4 or IPv6 address of the endpoint.
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// Metadata for the endpoint. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 512 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port that the endpoint is running on, must be in the
        /// range of [0, 65535]. If unspecified, the default is 0.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The resource name of the service that this endpoint provides.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 or IPv6 address of the endpoint.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata for the endpoint. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 512 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Port that the endpoint is running on, must be in the
        /// range of [0, 65535]. If unspecified, the default is 0.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The resource name of the service that this endpoint provides.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 or IPv6 address of the endpoint.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The Resource ID must be 1-63 characters long, including digits,
        /// lowercase letters or the hyphen character.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata for the endpoint. This data can be consumed
        /// by service clients. The entire metadata dictionary may contain
        /// up to 512 characters, spread across all key-value pairs.
        /// Metadata that goes beyond any these limits will be rejected.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port that the endpoint is running on, must be in the
        /// range of [0, 65535]. If unspecified, the default is 0.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The resource name of the service that this endpoint provides.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public EndpointState()
        {
        }
    }
}
