# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    alternative_name_server_config: pulumi.Output[Optional['outputs.PolicyAlternativeNameServerConfig']] = pulumi.output_property("alternativeNameServerConfig")
    """
    Sets an alternative name server for the associated networks.
    When specified, all DNS queries are forwarded to a name server that you choose.
    Names such as .internal are not available when an alternative name server is specified.  Structure is documented below.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    A textual description field. Defaults to 'Managed by Pulumi'.
    """
    enable_inbound_forwarding: pulumi.Output[Optional[bool]] = pulumi.output_property("enableInboundForwarding")
    """
    Allows networks bound to this policy to receive DNS queries sent
    by VMs or applications over VPN connections. When enabled, a
    virtual IP address will be allocated from each of the sub-networks
    that are bound to this policy.
    """
    enable_logging: pulumi.Output[Optional[bool]] = pulumi.output_property("enableLogging")
    """
    Controls whether logging is enabled for the networks bound to this policy.
    Defaults to no logging if not set.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    User assigned name for this policy.
    """
    networks: pulumi.Output[Optional[List['outputs.PolicyNetwork']]] = pulumi.output_property("networks")
    """
    List of network names specifying networks to which this policy is applied.  Structure is documented below.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, alternative_name_server_config: Optional[pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']]] = None, description: Optional[pulumi.Input[str]] = None, enable_inbound_forwarding: Optional[pulumi.Input[bool]] = None, enable_logging: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, networks: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]]] = None, project: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A policy is a collection of DNS rules applied to one or more Virtual
        Private Cloud resources.

        To get more information about Policy, see:

        * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
        * How-to Guides
            * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)

        ## Example Usage
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
                "targetNameServers": [
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
                    "networkUrl": network_1.id,
                },
                {
                    "networkUrl": network_2.id,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']] alternative_name_server_config: Sets an alternative name server for the associated networks.
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
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]] networks: List of network names specifying networks to which this policy is applied.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, alternative_name_server_config: Optional[pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']]] = None, description: Optional[pulumi.Input[str]] = None, enable_inbound_forwarding: Optional[pulumi.Input[bool]] = None, enable_logging: Optional[pulumi.Input[bool]] = None, name: Optional[pulumi.Input[str]] = None, networks: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]]] = None, project: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyAlternativeNameServerConfigArgs']] alternative_name_server_config: Sets an alternative name server for the associated networks.
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
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyNetworkArgs']]]] networks: List of network names specifying networks to which this policy is applied.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

