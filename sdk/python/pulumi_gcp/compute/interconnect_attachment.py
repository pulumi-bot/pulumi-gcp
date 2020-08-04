# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class InterconnectAttachment(pulumi.CustomResource):
    admin_enabled: pulumi.Output[bool]
    """
    Whether the VLAN attachment is enabled or disabled.  When using
    PARTNER type this will Pre-Activate the interconnect attachment
    """
    bandwidth: pulumi.Output[str]
    """
    Provisioned bandwidth capacity for the interconnect attachment.
    For attachments of type DEDICATED, the user can set the bandwidth.
    For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
    Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
    Defaults to BPS_10G
    """
    candidate_subnets: pulumi.Output[list]
    """
    Up to 16 candidate prefixes that can be used to restrict the allocation
    of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
    All prefixes must be within link-local address space (169.254.0.0/16)
    and must be /29 or shorter (/28, /27, etc). Google will attempt to select
    an unused /29 from the supplied candidate prefix(es). The request will
    fail if all possible /29s are in use on Google's edge. If not supplied,
    Google will randomly select an unused /29 from all of link-local space.
    """
    cloud_router_ip_address: pulumi.Output[str]
    """
    IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    customer_router_ip_address: pulumi.Output[str]
    """
    IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    edge_availability_domain: pulumi.Output[str]
    """
    Desired availability domain for the attachment. Only available for type
    PARTNER, at creation time. For improved reliability, customers should
    configure a pair of attachments with one per availability domain. The
    selected availability domain will be provided to the Partner via the
    pairing key so that the provisioned circuit will lie in the specified
    domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
    """
    google_reference_id: pulumi.Output[str]
    """
    Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
    issues.
    """
    interconnect: pulumi.Output[str]
    """
    URL of the underlying Interconnect object that this attachment's
    traffic will traverse through. Required if type is DEDICATED, must not
    be set if type is PARTNER.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is created. The
    name must be 1-63 characters long, and comply with RFC1035. Specifically, the
    name must be 1-63 characters long and match the regular expression
    `a-z?` which means the first character must be a
    lowercase letter, and all following characters must be a dash, lowercase
    letter, or digit, except the last character, which cannot be a dash.
    """
    pairing_key: pulumi.Output[str]
    """
    [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
    initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
    """
    partner_asn: pulumi.Output[str]
    """
    [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
    layer 3 Partner if they configured BGP on behalf of the customer.
    """
    private_interconnect_info: pulumi.Output[dict]
    """
    Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
    to is of type DEDICATED.

      * `tag8021q` (`float`)
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Region where the regional interconnect attachment resides.
    """
    router: pulumi.Output[str]
    """
    URL of the cloud router to be used for dynamic routing. This router must be in
    the same region as this InterconnectAttachment. The InterconnectAttachment will
    automatically connect the Interconnect to the network & region within which the
    Cloud Router is configured.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    state: pulumi.Output[str]
    """
    [Output Only] The current state of this attachment's functionality.
    """
    type: pulumi.Output[str]
    """
    The type of InterconnectAttachment you wish to create. Defaults to
    DEDICATED.
    """
    vlan_tag8021q: pulumi.Output[float]
    """
    The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
    using PARTNER type this will be managed upstream.
    """
    def __init__(__self__, resource_name, opts=None, admin_enabled=None, bandwidth=None, candidate_subnets=None, description=None, edge_availability_domain=None, interconnect=None, name=None, project=None, region=None, router=None, type=None, vlan_tag8021q=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents an InterconnectAttachment (VLAN attachment) resource. For more
        information, see Creating VLAN Attachments.

        ## Example Usage
        ### Interconnect Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        foobar = gcp.compute.Router("foobar", network=google_compute_network["foobar"]["name"])
        on_prem = gcp.compute.InterconnectAttachment("onPrem",
            interconnect="my-interconnect-id",
            router=foobar.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_enabled: Whether the VLAN attachment is enabled or disabled.  When using
               PARTNER type this will Pre-Activate the interconnect attachment
        :param pulumi.Input[str] bandwidth: Provisioned bandwidth capacity for the interconnect attachment.
               For attachments of type DEDICATED, the user can set the bandwidth.
               For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
               Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
               Defaults to BPS_10G
        :param pulumi.Input[list] candidate_subnets: Up to 16 candidate prefixes that can be used to restrict the allocation
               of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
               All prefixes must be within link-local address space (169.254.0.0/16)
               and must be /29 or shorter (/28, /27, etc). Google will attempt to select
               an unused /29 from the supplied candidate prefix(es). The request will
               fail if all possible /29s are in use on Google's edge. If not supplied,
               Google will randomly select an unused /29 from all of link-local space.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] edge_availability_domain: Desired availability domain for the attachment. Only available for type
               PARTNER, at creation time. For improved reliability, customers should
               configure a pair of attachments with one per availability domain. The
               selected availability domain will be provided to the Partner via the
               pairing key so that the provisioned circuit will lie in the specified
               domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        :param pulumi.Input[str] interconnect: URL of the underlying Interconnect object that this attachment's
               traffic will traverse through. Required if type is DEDICATED, must not
               be set if type is PARTNER.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The
               name must be 1-63 characters long, and comply with RFC1035. Specifically, the
               name must be 1-63 characters long and match the regular expression
               `a-z?` which means the first character must be a
               lowercase letter, and all following characters must be a dash, lowercase
               letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the regional interconnect attachment resides.
        :param pulumi.Input[str] router: URL of the cloud router to be used for dynamic routing. This router must be in
               the same region as this InterconnectAttachment. The InterconnectAttachment will
               automatically connect the Interconnect to the network & region within which the
               Cloud Router is configured.
        :param pulumi.Input[str] type: The type of InterconnectAttachment you wish to create. Defaults to
               DEDICATED.
        :param pulumi.Input[float] vlan_tag8021q: The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
               using PARTNER type this will be managed upstream.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_enabled'] = admin_enabled
            __props__['bandwidth'] = bandwidth
            __props__['candidate_subnets'] = candidate_subnets
            __props__['description'] = description
            __props__['edge_availability_domain'] = edge_availability_domain
            __props__['interconnect'] = interconnect
            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            if router is None:
                raise TypeError("Missing required property 'router'")
            __props__['router'] = router
            __props__['type'] = type
            __props__['vlan_tag8021q'] = vlan_tag8021q
            __props__['cloud_router_ip_address'] = None
            __props__['creation_timestamp'] = None
            __props__['customer_router_ip_address'] = None
            __props__['google_reference_id'] = None
            __props__['pairing_key'] = None
            __props__['partner_asn'] = None
            __props__['private_interconnect_info'] = None
            __props__['self_link'] = None
            __props__['state'] = None
        super(InterconnectAttachment, __self__).__init__(
            'gcp:compute/interconnectAttachment:InterconnectAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_enabled=None, bandwidth=None, candidate_subnets=None, cloud_router_ip_address=None, creation_timestamp=None, customer_router_ip_address=None, description=None, edge_availability_domain=None, google_reference_id=None, interconnect=None, name=None, pairing_key=None, partner_asn=None, private_interconnect_info=None, project=None, region=None, router=None, self_link=None, state=None, type=None, vlan_tag8021q=None):
        """
        Get an existing InterconnectAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_enabled: Whether the VLAN attachment is enabled or disabled.  When using
               PARTNER type this will Pre-Activate the interconnect attachment
        :param pulumi.Input[str] bandwidth: Provisioned bandwidth capacity for the interconnect attachment.
               For attachments of type DEDICATED, the user can set the bandwidth.
               For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
               Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
               Defaults to BPS_10G
        :param pulumi.Input[list] candidate_subnets: Up to 16 candidate prefixes that can be used to restrict the allocation
               of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
               All prefixes must be within link-local address space (169.254.0.0/16)
               and must be /29 or shorter (/28, /27, etc). Google will attempt to select
               an unused /29 from the supplied candidate prefix(es). The request will
               fail if all possible /29s are in use on Google's edge. If not supplied,
               Google will randomly select an unused /29 from all of link-local space.
        :param pulumi.Input[str] cloud_router_ip_address: IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] customer_router_ip_address: IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] edge_availability_domain: Desired availability domain for the attachment. Only available for type
               PARTNER, at creation time. For improved reliability, customers should
               configure a pair of attachments with one per availability domain. The
               selected availability domain will be provided to the Partner via the
               pairing key so that the provisioned circuit will lie in the specified
               domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        :param pulumi.Input[str] google_reference_id: Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
               issues.
        :param pulumi.Input[str] interconnect: URL of the underlying Interconnect object that this attachment's
               traffic will traverse through. Required if type is DEDICATED, must not
               be set if type is PARTNER.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The
               name must be 1-63 characters long, and comply with RFC1035. Specifically, the
               name must be 1-63 characters long and match the regular expression
               `a-z?` which means the first character must be a
               lowercase letter, and all following characters must be a dash, lowercase
               letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] pairing_key: [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
               initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        :param pulumi.Input[str] partner_asn: [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
               layer 3 Partner if they configured BGP on behalf of the customer.
        :param pulumi.Input[dict] private_interconnect_info: Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
               to is of type DEDICATED.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the regional interconnect attachment resides.
        :param pulumi.Input[str] router: URL of the cloud router to be used for dynamic routing. This router must be in
               the same region as this InterconnectAttachment. The InterconnectAttachment will
               automatically connect the Interconnect to the network & region within which the
               Cloud Router is configured.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] state: [Output Only] The current state of this attachment's functionality.
        :param pulumi.Input[str] type: The type of InterconnectAttachment you wish to create. Defaults to
               DEDICATED.
        :param pulumi.Input[float] vlan_tag8021q: The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
               using PARTNER type this will be managed upstream.

        The **private_interconnect_info** object supports the following:

          * `tag8021q` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_enabled"] = admin_enabled
        __props__["bandwidth"] = bandwidth
        __props__["candidate_subnets"] = candidate_subnets
        __props__["cloud_router_ip_address"] = cloud_router_ip_address
        __props__["creation_timestamp"] = creation_timestamp
        __props__["customer_router_ip_address"] = customer_router_ip_address
        __props__["description"] = description
        __props__["edge_availability_domain"] = edge_availability_domain
        __props__["google_reference_id"] = google_reference_id
        __props__["interconnect"] = interconnect
        __props__["name"] = name
        __props__["pairing_key"] = pairing_key
        __props__["partner_asn"] = partner_asn
        __props__["private_interconnect_info"] = private_interconnect_info
        __props__["project"] = project
        __props__["region"] = region
        __props__["router"] = router
        __props__["self_link"] = self_link
        __props__["state"] = state
        __props__["type"] = type
        __props__["vlan_tag8021q"] = vlan_tag8021q
        return InterconnectAttachment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
