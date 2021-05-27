// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetSubnetwork
    {
        /// <summary>
        /// Get a subnetwork within GCE from its name and region.
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
        ///         var my_subnetwork = Output.Create(Gcp.Compute.GetSubnetwork.InvokeAsync(new Gcp.Compute.GetSubnetworkArgs
        ///         {
        ///             Name = "default-us-east1",
        ///             Region = "us-east1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSubnetworkResult> InvokeAsync(GetSubnetworkArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubnetworkResult>("gcp:compute/getSubnetwork:getSubnetwork", args ?? new GetSubnetworkArgs(), options.WithVersion());

        public static Output<GetSubnetworkResult> Invoke(GetSubnetworkOutputArgs? args = null, InvokeOptions? options = null)
        {
            args = args ?? new GetSubnetworkOutputArgs();
            return Pulumi.Output.All(
                args.Name.Box(),
                args.Project.Box(),
                args.Region.Box(),
                args.SelfLink.Box()
            ).Apply(a => {
                    var args = new GetSubnetworkArgs();
                    a[0].Set(args, nameof(args.Name));
                    a[1].Set(args, nameof(args.Project));
                    a[2].Set(args, nameof(args.Region));
                    a[3].Set(args, nameof(args.SelfLink));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetSubnetworkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the subnetwork. One of `name` or `self_link`
        /// must be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region this subnetwork has been created in. If
        /// unspecified, this defaults to the region configured in the provider.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The self link of the subnetwork. If `self_link` is
        /// specified, `name`, `project`, and `region` are ignored.
        /// </summary>
        [Input("selfLink")]
        public string? SelfLink { get; set; }

        public GetSubnetworkArgs()
        {
        }
    }

    public sealed class GetSubnetworkOutputArgs
    {
        /// <summary>
        /// The name of the subnetwork. One of `name` or `self_link`
        /// must be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region this subnetwork has been created in. If
        /// unspecified, this defaults to the region configured in the provider.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The self link of the subnetwork. If `self_link` is
        /// specified, `name`, `project`, and `region` are ignored.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public GetSubnetworkOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubnetworkResult
    {
        /// <summary>
        /// Description of this subnetwork.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IP address of the gateway.
        /// </summary>
        public readonly string GatewayAddress;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The range of IP addresses belonging to this subnetwork
        /// secondary range.
        /// </summary>
        public readonly string IpCidrRange;
        public readonly string? Name;
        /// <summary>
        /// The network name or resource link to the parent
        /// network of this subnetwork.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Whether the VMs in this subnet
        /// can access Google services without assigned external IP
        /// addresses.
        /// </summary>
        public readonly bool PrivateIpGoogleAccess;
        public readonly string Project;
        public readonly string Region;
        /// <summary>
        /// An array of configurations for secondary IP ranges for
        /// VM instances contained in this subnetwork. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubnetworkSecondaryIpRangeResult> SecondaryIpRanges;
        public readonly string SelfLink;

        [OutputConstructor]
        private GetSubnetworkResult(
            string description,

            string gatewayAddress,

            string id,

            string ipCidrRange,

            string? name,

            string network,

            bool privateIpGoogleAccess,

            string project,

            string region,

            ImmutableArray<Outputs.GetSubnetworkSecondaryIpRangeResult> secondaryIpRanges,

            string selfLink)
        {
            Description = description;
            GatewayAddress = gatewayAddress;
            Id = id;
            IpCidrRange = ipCidrRange;
            Name = name;
            Network = network;
            PrivateIpGoogleAccess = privateIpGoogleAccess;
            Project = project;
            Region = region;
            SecondaryIpRanges = secondaryIpRanges;
            SelfLink = selfLink;
        }
    }
}
