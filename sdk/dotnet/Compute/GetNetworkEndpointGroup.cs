// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetNetworkEndpointGroup
    {
        /// <summary>
        /// Use this data source to access a Network Endpoint Group's attributes.
        /// 
        /// The NEG may be found by providing either a `self_link`, or a `name` and a `zone`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var neg1 = Output.Create(Gcp.Compute.GetNetworkEndpointGroup.InvokeAsync(new Gcp.Compute.GetNetworkEndpointGroupArgs
        ///         {
        ///             Name = "k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///             Zone = "us-central1-a",
        ///         }));
        ///         var neg2 = Output.Create(Gcp.Compute.GetNetworkEndpointGroup.InvokeAsync(new Gcp.Compute.GetNetworkEndpointGroupArgs
        ///         {
        ///             SelfLink = "https://www.googleapis.com/compute/v1/projects/myproject/zones/us-central1-a/networkEndpointGroups/k8s1-abcdef01-myns-mysvc-8080-4b6bac43",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkEndpointGroupResult> InvokeAsync(GetNetworkEndpointGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkEndpointGroupResult>("gcp:compute/getNetworkEndpointGroup:getNetworkEndpointGroup", args ?? new GetNetworkEndpointGroupArgs(), options.WithVersion());

        public static Output<GetNetworkEndpointGroupResult> Invoke(GetNetworkEndpointGroupOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetNetworkEndpointGroupOutputArgs();
            return Pulumi.Output.All(
                args.Name.Box(),
                args.Project.Box(),
                args.SelfLink.Box(),
                args.Zone.Box()
            ).Apply(a => {
                    var args = new GetNetworkEndpointGroupArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.Project));
                    a[2].Set(args, nameof(args.SelfLink));
                    a[3].Set(args, nameof(args.Zone));
                    return InvokeAsync(args, options);
            });
        }
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
        /// The ID of the project to list versions in.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

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

    public sealed class GetNetworkEndpointGroupOutputArgs
    {
        /// <summary>
        /// The Network Endpoint Group name.
        /// Provide either this or a `self_link`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project to list versions in.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The Network Endpoint Group self\_link.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The Network Endpoint Group availability zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetNetworkEndpointGroupOutputArgs()
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
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The network to which all network endpoints in the NEG belong.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Type of network endpoints in this network endpoint group.
        /// </summary>
        public readonly string NetworkEndpointType;
        public readonly string? Project;
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

        [OutputConstructor]
        private GetNetworkEndpointGroupResult(
            int defaultPort,

            string description,

            string id,

            string? name,

            string network,

            string networkEndpointType,

            string? project,

            string? selfLink,

            int size,

            string subnetwork,

            string? zone)
        {
            DefaultPort = defaultPort;
            Description = description;
            Id = id;
            Name = name;
            Network = network;
            NetworkEndpointType = networkEndpointType;
            Project = project;
            SelfLink = selfLink;
            Size = size;
            Subnetwork = subnetwork;
            Zone = zone;
        }
    }
}
