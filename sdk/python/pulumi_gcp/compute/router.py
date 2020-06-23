# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Router(pulumi.CustomResource):
    bgp: pulumi.Output[dict]
    """
    BGP information specific to this router.  Structure is documented below.

      * `advertise_mode` (`str`) - User-specified flag to indicate which mode to use for advertisement.
      * `advertised_groups` (`list`) - User-specified list of prefix groups to advertise in custom mode.
        This field can only be populated if advertiseMode is CUSTOM and
        is advertised to all peers of the router. These groups will be
        advertised in addition to any specified prefixes. Leave this field
        blank to advertise no custom groups.
        This enum field has the one valid value: ALL_SUBNETS
      * `advertised_ip_ranges` (`list`) - User-specified list of individual IP ranges to advertise in
        custom mode. This field can only be populated if advertiseMode
        is CUSTOM and is advertised to all peers of the router. These IP
        ranges will be advertised in addition to any specified groups.
        Leave this field blank to advertise no custom IP ranges.  Structure is documented below.
        * `description` (`str`) - User-specified description for the IP range.
        * `range` (`str`) - The IP range to advertise. The value must be a
          CIDR-formatted string.

      * `asn` (`float`) - Local BGP Autonomous System Number (ASN). Must be an RFC6996
        private ASN, either 16-bit or 32-bit. The value will be fixed for
        this router resource. All VPN tunnels that link to this router
        will have the same local ASN.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    User-specified description for the IP range.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. The name must be 1-63 characters long, and
    comply with RFC1035. Specifically, the name must be 1-63 characters
    long and match the regular expression `a-z?`
    which means the first character must be a lowercase letter, and all
    following characters must be a dash, lowercase letter, or digit,
    except the last character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    A reference to the network to which this router belongs.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Region where the router resides.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, bgp=None, description=None, name=None, network=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a Router resource.

        To get more information about Router, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
        * How-to Guides
            * [Google Cloud Router](https://cloud.google.com/router/docs/)

        ## Example Usage
        ### Router Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        foobar_network = gcp.compute.Network("foobarNetwork", auto_create_subnetworks=False)
        foobar_router = gcp.compute.Router("foobarRouter",
            network=foobar_network.name,
            bgp={
                "asn": 64514,
                "advertise_mode": "CUSTOM",
                "advertised_groups": ["ALL_SUBNETS"],
                "advertised_ip_ranges": [
                    {
                        "range": "1.2.3.4",
                    },
                    {
                        "range": "6.7.0.0/16",
                    },
                ],
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bgp: BGP information specific to this router.  Structure is documented below.
        :param pulumi.Input[str] description: User-specified description for the IP range.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?`
               which means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: A reference to the network to which this router belongs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the router resides.

        The **bgp** object supports the following:

          * `advertise_mode` (`pulumi.Input[str]`) - User-specified flag to indicate which mode to use for advertisement.
          * `advertised_groups` (`pulumi.Input[list]`) - User-specified list of prefix groups to advertise in custom mode.
            This field can only be populated if advertiseMode is CUSTOM and
            is advertised to all peers of the router. These groups will be
            advertised in addition to any specified prefixes. Leave this field
            blank to advertise no custom groups.
            This enum field has the one valid value: ALL_SUBNETS
          * `advertised_ip_ranges` (`pulumi.Input[list]`) - User-specified list of individual IP ranges to advertise in
            custom mode. This field can only be populated if advertiseMode
            is CUSTOM and is advertised to all peers of the router. These IP
            ranges will be advertised in addition to any specified groups.
            Leave this field blank to advertise no custom IP ranges.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - User-specified description for the IP range.
            * `range` (`pulumi.Input[str]`) - The IP range to advertise. The value must be a
              CIDR-formatted string.

          * `asn` (`pulumi.Input[float]`) - Local BGP Autonomous System Number (ASN). Must be an RFC6996
            private ASN, either 16-bit or 32-bit. The value will be fixed for
            this router resource. All VPN tunnels that link to this router
            will have the same local ASN.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['bgp'] = bgp
            __props__['description'] = description
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['project'] = project
            __props__['region'] = region
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(Router, __self__).__init__(
            'gcp:compute/router:Router',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, bgp=None, creation_timestamp=None, description=None, name=None, network=None, project=None, region=None, self_link=None):
        """
        Get an existing Router resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bgp: BGP information specific to this router.  Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: User-specified description for the IP range.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?`
               which means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: A reference to the network to which this router belongs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where the router resides.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        The **bgp** object supports the following:

          * `advertise_mode` (`pulumi.Input[str]`) - User-specified flag to indicate which mode to use for advertisement.
          * `advertised_groups` (`pulumi.Input[list]`) - User-specified list of prefix groups to advertise in custom mode.
            This field can only be populated if advertiseMode is CUSTOM and
            is advertised to all peers of the router. These groups will be
            advertised in addition to any specified prefixes. Leave this field
            blank to advertise no custom groups.
            This enum field has the one valid value: ALL_SUBNETS
          * `advertised_ip_ranges` (`pulumi.Input[list]`) - User-specified list of individual IP ranges to advertise in
            custom mode. This field can only be populated if advertiseMode
            is CUSTOM and is advertised to all peers of the router. These IP
            ranges will be advertised in addition to any specified groups.
            Leave this field blank to advertise no custom IP ranges.  Structure is documented below.
            * `description` (`pulumi.Input[str]`) - User-specified description for the IP range.
            * `range` (`pulumi.Input[str]`) - The IP range to advertise. The value must be a
              CIDR-formatted string.

          * `asn` (`pulumi.Input[float]`) - Local BGP Autonomous System Number (ASN). Must be an RFC6996
            private ASN, either 16-bit or 32-bit. The value will be fixed for
            this router resource. All VPN tunnels that link to this router
            will have the same local ASN.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bgp"] = bgp
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["network"] = network
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        return Router(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
