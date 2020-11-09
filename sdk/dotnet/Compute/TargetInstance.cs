// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents a TargetInstance resource which defines an endpoint instance
    /// that terminates traffic of certain protocols. In particular, they are used
    /// in Protocol Forwarding, where forwarding rules can send packets to a
    /// non-NAT'ed target instance. Each target instance contains a single
    /// virtual machine instance that receives and handles traffic from the
    /// corresponding forwarding rules.
    /// 
    /// To get more information about TargetInstance, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetInstances)
    /// * How-to Guides
    ///     * [Using Protocol Forwarding](https://cloud.google.com/compute/docs/protocol-forwarding)
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// TargetInstance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetInstance:TargetInstance default projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetInstance:TargetInstance default {{project}}/{{zone}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetInstance:TargetInstance default {{zone}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/targetInstance:TargetInstance default {{name}}
    /// ```
    /// </summary>
    public partial class TargetInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Compute instance VM handling traffic for this target instance.
        /// Accepts the instance self-link, relative path
        /// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
        /// name is given, the zone will default to the given zone or
        /// the provider-default zone and the project will default to the
        /// provider-level project.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// NAT option controlling how IPs are NAT'ed to the instance.
        /// Currently only NO_NAT (default value) is supported.
        /// Default value is `NO_NAT`.
        /// Possible values are `NO_NAT`.
        /// </summary>
        [Output("natPolicy")]
        public Output<string?> NatPolicy { get; private set; } = null!;

        /// <summary>
        /// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// URL of the zone where the target instance resides.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a TargetInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TargetInstance(string name, TargetInstanceArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/targetInstance:TargetInstance", name, args ?? new TargetInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TargetInstance(string name, Input<string> id, TargetInstanceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/targetInstance:TargetInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TargetInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TargetInstance Get(string name, Input<string> id, TargetInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new TargetInstance(name, id, state, options);
        }
    }

    public sealed class TargetInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Compute instance VM handling traffic for this target instance.
        /// Accepts the instance self-link, relative path
        /// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
        /// name is given, the zone will default to the given zone or
        /// the provider-default zone and the project will default to the
        /// provider-level project.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NAT option controlling how IPs are NAT'ed to the instance.
        /// Currently only NO_NAT (default value) is supported.
        /// Default value is `NO_NAT`.
        /// Possible values are `NO_NAT`.
        /// </summary>
        [Input("natPolicy")]
        public Input<string>? NatPolicy { get; set; }

        /// <summary>
        /// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// URL of the zone where the target instance resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public TargetInstanceArgs()
        {
        }
    }

    public sealed class TargetInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Compute instance VM handling traffic for this target instance.
        /// Accepts the instance self-link, relative path
        /// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
        /// name is given, the zone will default to the given zone or
        /// the provider-default zone and the project will default to the
        /// provider-level project.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NAT option controlling how IPs are NAT'ed to the instance.
        /// Currently only NO_NAT (default value) is supported.
        /// Default value is `NO_NAT`.
        /// Possible values are `NO_NAT`.
        /// </summary>
        [Input("natPolicy")]
        public Input<string>? NatPolicy { get; set; }

        /// <summary>
        /// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// URL of the zone where the target instance resides.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public TargetInstanceState()
        {
        }
    }
}
