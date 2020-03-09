// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public partial class InterconnectAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the
        /// interconnect attachment
        /// </summary>
        [Output("adminEnabled")]
        public Output<bool?> AdminEnabled { get; private set; } = null!;

        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user
        /// can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the
        /// interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and
        /// DEDICATED, Defaults to BPS_10G
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
        /// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space
        /// (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29
        /// from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's
        /// edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        [Output("candidateSubnets")]
        public Output<ImmutableArray<string>> CandidateSubnets { get; private set; } = null!;

        /// <summary>
        /// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        [Output("cloudRouterIpAddress")]
        public Output<string> CloudRouterIpAddress { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect
        /// attachment.
        /// </summary>
        [Output("customerRouterIpAddress")]
        public Output<string> CustomerRouterIpAddress { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For
        /// improved reliability, customers should configure a pair of attachments with one per availability domain. The
        /// selected availability domain will be provided to the Partner via the pairing key so that the provisioned
        /// circuit will lie in the specified domain. If not specified, the value will default to
        /// AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        [Output("edgeAvailabilityDomain")]
        public Output<string?> EdgeAvailabilityDomain { get; private set; } = null!;

        /// <summary>
        /// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend
        /// connectivity issues.
        /// </summary>
        [Output("googleReferenceId")]
        public Output<string> GoogleReferenceId { get; private set; } = null!;

        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if
        /// type is DEDICATED, must not be set if type is PARTNER.
        /// </summary>
        [Output("interconnect")]
        public Output<string?> Interconnect { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment
        /// used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        /// </summary>
        [Output("pairingKey")]
        public Output<string> PairingKey { get; private set; } = null!;

        /// <summary>
        /// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be
        /// supplied by a layer 3 Partner if they configured BGP on behalf of the customer.
        /// </summary>
        [Output("partnerAsn")]
        public Output<string> PartnerAsn { get; private set; } = null!;

        /// <summary>
        /// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this
        /// is attached to is of type DEDICATED.
        /// </summary>
        [Output("privateInterconnectInfo")]
        public Output<Outputs.InterconnectAttachmentPrivateInterconnectInfo> PrivateInterconnectInfo { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Region where the regional interconnect attachment resides.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
        /// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the
        /// network &amp; region within which the Cloud Router is configured.
        /// </summary>
        [Output("router")]
        public Output<string> Router { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// [Output Only] The current state of this attachment's functionality.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be
        /// managed upstream.
        /// </summary>
        [Output("vlanTag8021q")]
        public Output<int> VlanTag8021q { get; private set; } = null!;


        /// <summary>
        /// Create a InterconnectAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InterconnectAttachment(string name, InterconnectAttachmentArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/interconnectAttachment:InterconnectAttachment", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private InterconnectAttachment(string name, Input<string> id, InterconnectAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/interconnectAttachment:InterconnectAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InterconnectAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InterconnectAttachment Get(string name, Input<string> id, InterconnectAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new InterconnectAttachment(name, id, state, options);
        }
    }

    public sealed class InterconnectAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the
        /// interconnect attachment
        /// </summary>
        [Input("adminEnabled")]
        public Input<bool>? AdminEnabled { get; set; }

        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user
        /// can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the
        /// interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and
        /// DEDICATED, Defaults to BPS_10G
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("candidateSubnets")]
        private InputList<string>? _candidateSubnets;

        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
        /// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space
        /// (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29
        /// from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's
        /// edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        public InputList<string> CandidateSubnets
        {
            get => _candidateSubnets ?? (_candidateSubnets = new InputList<string>());
            set => _candidateSubnets = value;
        }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For
        /// improved reliability, customers should configure a pair of attachments with one per availability domain. The
        /// selected availability domain will be provided to the Partner via the pairing key so that the provisioned
        /// circuit will lie in the specified domain. If not specified, the value will default to
        /// AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        [Input("edgeAvailabilityDomain")]
        public Input<string>? EdgeAvailabilityDomain { get; set; }

        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if
        /// type is DEDICATED, must not be set if type is PARTNER.
        /// </summary>
        [Input("interconnect")]
        public Input<string>? Interconnect { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the regional interconnect attachment resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
        /// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the
        /// network &amp; region within which the Cloud Router is configured.
        /// </summary>
        [Input("router", required: true)]
        public Input<string> Router { get; set; } = null!;

        /// <summary>
        /// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be
        /// managed upstream.
        /// </summary>
        [Input("vlanTag8021q")]
        public Input<int>? VlanTag8021q { get; set; }

        public InterconnectAttachmentArgs()
        {
        }
    }

    public sealed class InterconnectAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the
        /// interconnect attachment
        /// </summary>
        [Input("adminEnabled")]
        public Input<bool>? AdminEnabled { get; set; }

        /// <summary>
        /// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user
        /// can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the
        /// interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and
        /// DEDICATED, Defaults to BPS_10G
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        [Input("candidateSubnets")]
        private InputList<string>? _candidateSubnets;

        /// <summary>
        /// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
        /// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space
        /// (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29
        /// from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's
        /// edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        /// </summary>
        public InputList<string> CandidateSubnets
        {
            get => _candidateSubnets ?? (_candidateSubnets = new InputList<string>());
            set => _candidateSubnets = value;
        }

        /// <summary>
        /// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        /// </summary>
        [Input("cloudRouterIpAddress")]
        public Input<string>? CloudRouterIpAddress { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect
        /// attachment.
        /// </summary>
        [Input("customerRouterIpAddress")]
        public Input<string>? CustomerRouterIpAddress { get; set; }

        /// <summary>
        /// An optional description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For
        /// improved reliability, customers should configure a pair of attachments with one per availability domain. The
        /// selected availability domain will be provided to the Partner via the pairing key so that the provisioned
        /// circuit will lie in the specified domain. If not specified, the value will default to
        /// AVAILABILITY_DOMAIN_ANY.
        /// </summary>
        [Input("edgeAvailabilityDomain")]
        public Input<string>? EdgeAvailabilityDomain { get; set; }

        /// <summary>
        /// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend
        /// connectivity issues.
        /// </summary>
        [Input("googleReferenceId")]
        public Input<string>? GoogleReferenceId { get; set; }

        /// <summary>
        /// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if
        /// type is DEDICATED, must not be set if type is PARTNER.
        /// </summary>
        [Input("interconnect")]
        public Input<string>? Interconnect { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters
        /// long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
        /// expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all
        /// following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be
        /// a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment
        /// used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        /// </summary>
        [Input("pairingKey")]
        public Input<string>? PairingKey { get; set; }

        /// <summary>
        /// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be
        /// supplied by a layer 3 Partner if they configured BGP on behalf of the customer.
        /// </summary>
        [Input("partnerAsn")]
        public Input<string>? PartnerAsn { get; set; }

        /// <summary>
        /// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this
        /// is attached to is of type DEDICATED.
        /// </summary>
        [Input("privateInterconnectInfo")]
        public Input<Inputs.InterconnectAttachmentPrivateInterconnectInfoGetArgs>? PrivateInterconnectInfo { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Region where the regional interconnect attachment resides.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
        /// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the
        /// network &amp; region within which the Cloud Router is configured.
        /// </summary>
        [Input("router")]
        public Input<string>? Router { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// [Output Only] The current state of this attachment's functionality.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be
        /// managed upstream.
        /// </summary>
        [Input("vlanTag8021q")]
        public Input<int>? VlanTag8021q { get; set; }

        public InterconnectAttachmentState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class InterconnectAttachmentPrivateInterconnectInfoGetArgs : Pulumi.ResourceArgs
    {
        [Input("tag8021q")]
        public Input<int>? Tag8021q { get; set; }

        public InterconnectAttachmentPrivateInterconnectInfoGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class InterconnectAttachmentPrivateInterconnectInfo
    {
        public readonly int Tag8021q;

        [OutputConstructor]
        private InterconnectAttachmentPrivateInterconnectInfo(int tag8021q)
        {
            Tag8021q = tag8021q;
        }
    }
    }
}
