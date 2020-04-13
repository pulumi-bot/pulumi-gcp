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
    /// Represents an Address resource.
    /// 
    /// Each virtual machine instance has an ephemeral internal IP address and,
    /// optionally, an external IP address. To communicate between instances on
    /// the same network, you can use an instance's internal IP address. To
    /// communicate with the Internet and instances outside of the same network,
    /// you must specify the instance's external IP address.
    /// 
    /// Internal IP addresses are ephemeral and only belong to an instance for
    /// the lifetime of the instance; if the instance is deleted and recreated,
    /// the instance is assigned a new internal IP address, either by Compute
    /// Engine or by you. External IP addresses can be either ephemeral or
    /// static.
    /// 
    /// 
    /// To get more information about Address, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
    /// * How-to Guides
    ///     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
    ///     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_address.html.markdown.
    /// </summary>
    public partial class Address : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The static external IP address represented by this resource. Only
        /// IPv4 is supported. An address may only be specified for INTERNAL
        /// address types. The IP address must be inside the specified subnetwork,
        /// if any.
        /// </summary>
        [Output("address")]
        public Output<string> IPAddress { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The type of address to reserve, either INTERNAL or EXTERNAL.
        /// If unspecified, defaults to EXTERNAL.
        /// </summary>
        [Output("addressType")]
        public Output<string?> AddressType { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Output("labelFingerprint")]
        public Output<string> LabelFingerprint { get; private set; } = null!;

        /// <summary>
        /// Labels to apply to this address. A list of key-&gt;value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?`
        /// which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The networking tier used for configuring this address. This field can
        /// take the following values: PREMIUM or STANDARD. If this field is not
        /// specified, it is assumed to be PREMIUM.
        /// </summary>
        [Output("networkTier")]
        public Output<string> NetworkTier { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The purpose of this resource, which can be one of the following values:
        /// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
        /// This should only be set when using an Internal address.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The Region in which the created address should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// -
        /// (Optional)
        /// The URL of the subnetwork in which to reserve the address. If an IP
        /// address is specified, it must be within the subnetwork's IP range.
        /// This field can only be used with INTERNAL type with
        /// GCE_ENDPOINT/DNS_RESOLVER purposes.
        /// </summary>
        [Output("subnetwork")]
        public Output<string> Subnetwork { get; private set; } = null!;

        /// <summary>
        /// The URLs of the resources that are using this address.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Address resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Address(string name, AddressArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:compute/address:Address", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Address(string name, Input<string> id, AddressState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/address:Address", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Address resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Address Get(string name, Input<string> id, AddressState? state = null, CustomResourceOptions? options = null)
        {
            return new Address(name, id, state, options);
        }
    }

    public sealed class AddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The static external IP address represented by this resource. Only
        /// IPv4 is supported. An address may only be specified for INTERNAL
        /// address types. The IP address must be inside the specified subnetwork,
        /// if any.
        /// </summary>
        [Input("address")]
        public Input<string>? IPAddress { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The type of address to reserve, either INTERNAL or EXTERNAL.
        /// If unspecified, defaults to EXTERNAL.
        /// </summary>
        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this address. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?`
        /// which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The networking tier used for configuring this address. This field can
        /// take the following values: PREMIUM or STANDARD. If this field is not
        /// specified, it is assumed to be PREMIUM.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The purpose of this resource, which can be one of the following values:
        /// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
        /// This should only be set when using an Internal address.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The Region in which the created address should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The URL of the subnetwork in which to reserve the address. If an IP
        /// address is specified, it must be within the subnetwork's IP range.
        /// This field can only be used with INTERNAL type with
        /// GCE_ENDPOINT/DNS_RESOLVER purposes.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        public AddressArgs()
        {
        }
    }

    public sealed class AddressState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// (Optional)
        /// The static external IP address represented by this resource. Only
        /// IPv4 is supported. An address may only be specified for INTERNAL
        /// address types. The IP address must be inside the specified subnetwork,
        /// if any.
        /// </summary>
        [Input("address")]
        public Input<string>? IPAddress { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The type of address to reserve, either INTERNAL or EXTERNAL.
        /// If unspecified, defaults to EXTERNAL.
        /// </summary>
        [Input("addressType")]
        public Input<string>? AddressType { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The fingerprint used for optimistic locking of this resource. Used internally during updates.
        /// </summary>
        [Input("labelFingerprint")]
        public Input<string>? LabelFingerprint { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels to apply to this address. A list of key-&gt;value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// -
        /// (Required)
        /// Name of the resource. The name must be 1-63 characters long, and
        /// comply with RFC1035. Specifically, the name must be 1-63 characters
        /// long and match the regular expression `a-z?`
        /// which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit,
        /// except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The networking tier used for configuring this address. This field can
        /// take the following values: PREMIUM or STANDARD. If this field is not
        /// specified, it is assumed to be PREMIUM.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The purpose of this resource, which can be one of the following values:
        /// - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
        /// This should only be set when using an Internal address.
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The Region in which the created address should reside.
        /// If it is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// -
        /// (Optional)
        /// The URL of the subnetwork in which to reserve the address. If an IP
        /// address is specified, it must be within the subnetwork's IP range.
        /// This field can only be used with INTERNAL type with
        /// GCE_ENDPOINT/DNS_RESOLVER purposes.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// The URLs of the resources that are using this address.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public AddressState()
        {
        }
    }
}
