# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Policy(pulumi.CustomResource):
    alternative_name_server_config: pulumi.Output[dict]
    """
    Sets an alternative name server for the associated networks.
    When specified, all DNS queries are forwarded to a name server that you choose.
    Names such as .internal are not available when an alternative name server is specified.  Structure is documented below.

      * `targetNameServers` (`list`) - Sets an alternative name server for the associated networks. When specified,
        all DNS queries are forwarded to a name server that you choose. Names such as .internal
        are not available when an alternative name server is specified.  Structure is documented below.
        * `ipv4Address` (`str`) - IPv4 address to forward to.
    """
    description: pulumi.Output[str]
    """
    A textual description field. Defaults to 'Managed by Pulumi'.
    """
    enable_inbound_forwarding: pulumi.Output[bool]
    """
    Allows networks bound to this policy to receive DNS queries sent
    by VMs or applications over VPN connections. When enabled, a
    virtual IP address will be allocated from each of the sub-networks
    that are bound to this policy.
    """
    enable_logging: pulumi.Output[bool]
    """
    Controls whether logging is enabled for the networks bound to this policy.
    Defaults to no logging if not set.
    """
    name: pulumi.Output[str]
    """
    User assigned name for this policy.
    """
    networks: pulumi.Output[list]
    """
    List of network names specifying networks to which this policy is applied.  Structure is documented below.

      * `networkUrl` (`str`) - The fully qualified URL of the VPC network to bind to.
        This should be formatted like
        `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, alternative_name_server_config=None, description=None, enable_inbound_forwarding=None, enable_logging=None, name=None, networks=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        A policy is a collection of DNS rules applied to one or more Virtual
        Private Cloud resources.

        To get more information about Policy, see:

        * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
        * How-to Guides
            * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)

        {{% examples %}}
        ## Example Usage
        {{% example %}}
        ### Dns Policy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network_1 = gcp.compute.Network("network-1", auto_create_subnetworks=False)
        network_2 = gcp.compute.Network("network-2", auto_create_subnetworks=False)
        example_policy = gcp.dns.Policy("example-policy",
            enable_inbound_forwarding=True,
            enable_logging=True,
            alternative_name_server_config={
                "target_name_servers": [
                    {
                        "ipv4Address": "172.16.1.10",
                    },
                    {
                        "ipv4Address": "172.16.1.20",
                    },
                ],
            },
            networks=[
                {
                    "networkUrl": network_1.self_link,
                },
                {
                    "networkUrl": network_2.self_link,
                },
            ])
        ```
        {{% /example %}}
        {{% /examples %}}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] alternative_name_server_config: Sets an alternative name server for the associated networks.
               When specified, all DNS queries are forwarded to a name server that you choose.
               Names such as .internal are not available when an alternative name server is specified.  Structure is documented below.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[bool] enable_inbound_forwarding: Allows networks bound to this policy to receive DNS queries sent
               by VMs or applications over VPN connections. When enabled, a
               virtual IP address will be allocated from each of the sub-networks
               that are bound to this policy.
        :param pulumi.Input[bool] enable_logging: Controls whether logging is enabled for the networks bound to this policy.
               Defaults to no logging if not set.
        :param pulumi.Input[str] name: User assigned name for this policy.
        :param pulumi.Input[list] networks: List of network names specifying networks to which this policy is applied.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **alternative_name_server_config** object supports the following:

          * `targetNameServers` (`pulumi.Input[list]`) - Sets an alternative name server for the associated networks. When specified,
            all DNS queries are forwarded to a name server that you choose. Names such as .internal
            are not available when an alternative name server is specified.  Structure is documented below.
            * `ipv4Address` (`pulumi.Input[str]`) - IPv4 address to forward to.

        The **networks** object supports the following:

          * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to bind to.
            This should be formatted like
            `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
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

            __props__['alternative_name_server_config'] = alternative_name_server_config
            __props__['description'] = description
            __props__['enable_inbound_forwarding'] = enable_inbound_forwarding
            __props__['enable_logging'] = enable_logging
            __props__['name'] = name
            __props__['networks'] = networks
            __props__['project'] = project
        super(Policy, __self__).__init__(
            'gcp:dns/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alternative_name_server_config=None, description=None, enable_inbound_forwarding=None, enable_logging=None, name=None, networks=None, project=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] alternative_name_server_config: Sets an alternative name server for the associated networks.
               When specified, all DNS queries are forwarded to a name server that you choose.
               Names such as .internal are not available when an alternative name server is specified.  Structure is documented below.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[bool] enable_inbound_forwarding: Allows networks bound to this policy to receive DNS queries sent
               by VMs or applications over VPN connections. When enabled, a
               virtual IP address will be allocated from each of the sub-networks
               that are bound to this policy.
        :param pulumi.Input[bool] enable_logging: Controls whether logging is enabled for the networks bound to this policy.
               Defaults to no logging if not set.
        :param pulumi.Input[str] name: User assigned name for this policy.
        :param pulumi.Input[list] networks: List of network names specifying networks to which this policy is applied.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **alternative_name_server_config** object supports the following:

          * `targetNameServers` (`pulumi.Input[list]`) - Sets an alternative name server for the associated networks. When specified,
            all DNS queries are forwarded to a name server that you choose. Names such as .internal
            are not available when an alternative name server is specified.  Structure is documented below.
            * `ipv4Address` (`pulumi.Input[str]`) - IPv4 address to forward to.

        The **networks** object supports the following:

          * `networkUrl` (`pulumi.Input[str]`) - The fully qualified URL of the VPC network to bind to.
            This should be formatted like
            `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alternative_name_server_config"] = alternative_name_server_config
        __props__["description"] = description
        __props__["enable_inbound_forwarding"] = enable_inbound_forwarding
        __props__["enable_logging"] = enable_logging
        __props__["name"] = name
        __props__["networks"] = networks
        __props__["project"] = project
        return Policy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
