// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access a Network Endpoint Group's attributes.
        /// 
        /// The NEG may be found by providing either a `self_link`, or a `name` and a `zone`.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_compute_network_endpoint_group.html.markdown.
        /// </summary>
        [Obsolete("Use GetNetworkEndpointGroup.InvokeAsync() instead")]
        public static Task<GetNetworkEndpointGroupResult> GetNetworkEndpointGroup(GetNetworkEndpointGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkEndpointGroupResult>("gcp:compute/getNetworkEndpointGroup:getNetworkEndpointGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetNetworkEndpointGroup
    {
        /// <summary>
        /// Use this data source to access a Network Endpoint Group's attributes.
        /// 
        /// The NEG may be found by providing either a `self_link`, or a `name` and a `zone`.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_compute_network_endpoint_group.html.markdown.
        /// </summary>
        public static Task<GetNetworkEndpointGroupResult> InvokeAsync(GetNetworkEndpointGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkEndpointGroupResult>("gcp:compute/getNetworkEndpointGroup:getNetworkEndpointGroup", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetNetworkEndpointGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Network Endpoint Group name.
        /// Provide either this or a `self_link`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The Network Endpoint Group self\_link.
        /// </summary>
        [Input("selfLink")]
        public string? SelfLink { get; set; }

        /// <summary>
        /// The Network Endpoint Group availability zone.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetNetworkEndpointGroupArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetNetworkEndpointGroupResult
    {
        /// <summary>
        /// The NEG default port.
        /// </summary>
        public readonly int DefaultPort;
        /// <summary>
        /// The NEG description.
        /// </summary>
        public readonly string Description;
        public readonly string? Name;
        /// <summary>
        /// The network to which all network endpoints in the NEG belong.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Type of network endpoints in this network endpoint group.
        /// </summary>
        public readonly string NetworkEndpointType;
        public readonly string Project;
        public readonly string? SelfLink;
        /// <summary>
        /// Number of network endpoints in the network endpoint group.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        public readonly string Subnetwork;
        public readonly string? Zone;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNetworkEndpointGroupResult(
            int defaultPort,
            string description,
            string? name,
            string network,
            string networkEndpointType,
            string project,
            string? selfLink,
            int size,
            string subnetwork,
            string? zone,
            string id)
        {
            DefaultPort = defaultPort;
            Description = description;
            Name = name;
            Network = network;
            NetworkEndpointType = networkEndpointType;
            Project = project;
            SelfLink = selfLink;
            Size = size;
            Subnetwork = subnetwork;
            Zone = zone;
            Id = id;
        }
    }
}
