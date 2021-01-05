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
    /// Contains the data that describes an Identity Aware Proxy owned client.
    /// 
    /// &gt; **Note:** Only internal org clients can be created via declarative tools. Other types of clients must be
    /// manually created via the GCP console. This restriction is due to the existing APIs and not lack of support
    /// in this tool.
    /// 
    /// To get more information about Client, see:
    /// 
    /// * [API documentation](https://cloud.google.com/iap/docs/reference/rest/v1/projects.brands.identityAwareProxyClients)
    /// * How-to Guides
    ///     * [Setting up IAP Client](https://cloud.google.com/iap/docs/authentication-howto)
    /// 
    /// &gt; **Warning:** All arguments including `secret` will be stored in the raw
    /// state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Client can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iap/client:Client default {{brand}}/identityAwareProxyClients/{{client_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iap/client:Client default {{brand}}/{{client_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:iap/client:Client")]
    public partial class Client : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier of the brand to which this client
        /// is attached to. The format is
        /// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
        /// </summary>
        [Output("brand")]
        public Output<string> Brand { get; private set; } = null!;

        /// <summary>
        /// Output only. Unique identifier of the OAuth client.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Output only. Client secret of the OAuth client.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs args, CustomResourceOptions? options = null)
            : base("gcp:iap/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iap/client:Client", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the brand to which this client
        /// is attached to. The format is
        /// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
        /// </summary>
        [Input("brand", required: true)]
        public Input<string> Brand { get; set; } = null!;

        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        public ClientArgs()
        {
        }
    }

    public sealed class ClientState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the brand to which this client
        /// is attached to. The format is
        /// `projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}`.
        /// </summary>
        [Input("brand")]
        public Input<string>? Brand { get; set; }

        /// <summary>
        /// Output only. Unique identifier of the OAuth client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// Human-friendly name given to the OAuth client.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Output only. Client secret of the OAuth client.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ClientState()
        {
        }
    }
}
