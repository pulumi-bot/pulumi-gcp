# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ManagedZone']


class ManagedZone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_name: Optional[pulumi.Input[str]] = None,
                 dnssec_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneDnssecConfigArgs']]] = None,
                 forwarding_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneForwardingConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peering_config: Optional[pulumi.Input[pulumi.InputType['ManagedZonePeeringConfigArgs']]] = None,
                 private_visibility_config: Optional[pulumi.Input[pulumi.InputType['ManagedZonePrivateVisibilityConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reverse_lookup: Optional[pulumi.Input[bool]] = None,
                 service_directory_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneServiceDirectoryConfigArgs']]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A zone is a subtree of the DNS namespace under one administrative
        responsibility. A ManagedZone is a resource that represents a DNS zone
        hosted by the Cloud DNS service.

        To get more information about ManagedZone, see:

        * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
        * How-to Guides
            * [Managing Zones](https://cloud.google.com/dns/zones/)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] dns_name: The DNS name of this managed zone, for instance "example.com.".
        :param pulumi.Input[pulumi.InputType['ManagedZoneDnssecConfigArgs']] dnssec_config: DNSSEC configuration
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ManagedZoneForwardingConfigArgs']] forwarding_config: The presence for this field indicates that outbound forwarding is enabled
               for this zone. The value of this field contains the set of destinations
               to forward to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this ManagedZone.
        :param pulumi.Input[str] name: User assigned name for this resource.
               Must be unique within the project.
        :param pulumi.Input[pulumi.InputType['ManagedZonePeeringConfigArgs']] peering_config: The presence of this field indicates that DNS Peering is enabled for this
               zone. The value of this field contains the network to peer with.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ManagedZonePrivateVisibilityConfigArgs']] private_visibility_config: For privately visible zones, the set of Virtual Private Cloud
               resources that the zone is visible from.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reverse_lookup: Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
               lookup queries using automatically configured records for VPC resources. This only applies
               to networks listed under `private_visibility_config`.
        :param pulumi.Input[pulumi.InputType['ManagedZoneServiceDirectoryConfigArgs']] service_directory_config: The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        :param pulumi.Input[str] visibility: The zone's visibility: public zones are exposed to the Internet,
               while private zones are visible only to Virtual Private Cloud resources.
               Default value is `public`.
               Possible values are `private` and `public`.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            if dns_name is None:
                raise TypeError("Missing required property 'dns_name'")
            __props__['dns_name'] = dns_name
            __props__['dnssec_config'] = dnssec_config
            __props__['forwarding_config'] = forwarding_config
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['peering_config'] = peering_config
            __props__['private_visibility_config'] = private_visibility_config
            __props__['project'] = project
            __props__['reverse_lookup'] = reverse_lookup
            __props__['service_directory_config'] = service_directory_config
            __props__['visibility'] = visibility
            __props__['name_servers'] = None
        super(ManagedZone, __self__).__init__(
            'gcp:dns/managedZone:ManagedZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            dnssec_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneDnssecConfigArgs']]] = None,
            forwarding_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneForwardingConfigArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_servers: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            peering_config: Optional[pulumi.Input[pulumi.InputType['ManagedZonePeeringConfigArgs']]] = None,
            private_visibility_config: Optional[pulumi.Input[pulumi.InputType['ManagedZonePrivateVisibilityConfigArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            reverse_lookup: Optional[pulumi.Input[bool]] = None,
            service_directory_config: Optional[pulumi.Input[pulumi.InputType['ManagedZoneServiceDirectoryConfigArgs']]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'ManagedZone':
        """
        Get an existing ManagedZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description field. Defaults to 'Managed by Pulumi'.
        :param pulumi.Input[str] dns_name: The DNS name of this managed zone, for instance "example.com.".
        :param pulumi.Input[pulumi.InputType['ManagedZoneDnssecConfigArgs']] dnssec_config: DNSSEC configuration
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ManagedZoneForwardingConfigArgs']] forwarding_config: The presence for this field indicates that outbound forwarding is enabled
               for this zone. The value of this field contains the set of destinations
               to forward to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this ManagedZone.
        :param pulumi.Input[str] name: User assigned name for this resource.
               Must be unique within the project.
        :param pulumi.Input[List[pulumi.Input[str]]] name_servers: Delegate your managed_zone to these virtual name servers; defined by the server
        :param pulumi.Input[pulumi.InputType['ManagedZonePeeringConfigArgs']] peering_config: The presence of this field indicates that DNS Peering is enabled for this
               zone. The value of this field contains the network to peer with.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['ManagedZonePrivateVisibilityConfigArgs']] private_visibility_config: For privately visible zones, the set of Virtual Private Cloud
               resources that the zone is visible from.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reverse_lookup: Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
               lookup queries using automatically configured records for VPC resources. This only applies
               to networks listed under `private_visibility_config`.
        :param pulumi.Input[pulumi.InputType['ManagedZoneServiceDirectoryConfigArgs']] service_directory_config: The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        :param pulumi.Input[str] visibility: The zone's visibility: public zones are exposed to the Internet,
               while private zones are visible only to Virtual Private Cloud resources.
               Default value is `public`.
               Possible values are `private` and `public`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["dns_name"] = dns_name
        __props__["dnssec_config"] = dnssec_config
        __props__["forwarding_config"] = forwarding_config
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["name_servers"] = name_servers
        __props__["peering_config"] = peering_config
        __props__["private_visibility_config"] = private_visibility_config
        __props__["project"] = project
        __props__["reverse_lookup"] = reverse_lookup
        __props__["service_directory_config"] = service_directory_config
        __props__["visibility"] = visibility
        return ManagedZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A textual description field. Defaults to 'Managed by Pulumi'.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        """
        The DNS name of this managed zone, for instance "example.com.".
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="dnssecConfig")
    def dnssec_config(self) -> Optional['outputs.ManagedZoneDnssecConfig']:
        """
        DNSSEC configuration
        Structure is documented below.
        """
        return pulumi.get(self, "dnssec_config")

    @property
    @pulumi.getter(name="forwardingConfig")
    def forwarding_config(self) -> Optional['outputs.ManagedZoneForwardingConfig']:
        """
        The presence for this field indicates that outbound forwarding is enabled
        for this zone. The value of this field contains the set of destinations
        to forward to.
        Structure is documented below.
        """
        return pulumi.get(self, "forwarding_config")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        A set of key/value label pairs to assign to this ManagedZone.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        User assigned name for this resource.
        Must be unique within the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> List[str]:
        """
        Delegate your managed_zone to these virtual name servers; defined by the server
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter(name="peeringConfig")
    def peering_config(self) -> Optional['outputs.ManagedZonePeeringConfig']:
        """
        The presence of this field indicates that DNS Peering is enabled for this
        zone. The value of this field contains the network to peer with.
        Structure is documented below.
        """
        return pulumi.get(self, "peering_config")

    @property
    @pulumi.getter(name="privateVisibilityConfig")
    def private_visibility_config(self) -> Optional['outputs.ManagedZonePrivateVisibilityConfig']:
        """
        For privately visible zones, the set of Virtual Private Cloud
        resources that the zone is visible from.
        Structure is documented below.
        """
        return pulumi.get(self, "private_visibility_config")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="reverseLookup")
    def reverse_lookup(self) -> Optional[bool]:
        """
        Specifies if this is a managed reverse lookup zone. If true, Cloud DNS will resolve reverse
        lookup queries using automatically configured records for VPC resources. This only applies
        to networks listed under `private_visibility_config`.
        """
        return pulumi.get(self, "reverse_lookup")

    @property
    @pulumi.getter(name="serviceDirectoryConfig")
    def service_directory_config(self) -> Optional['outputs.ManagedZoneServiceDirectoryConfig']:
        """
        The presence of this field indicates that this zone is backed by Service Directory. The value of this field contains information related to the namespace associated with the zone.  Structure is documented below.
        """
        return pulumi.get(self, "service_directory_config")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        """
        The zone's visibility: public zones are exposed to the Internet,
        while private zones are visible only to Virtual Private Cloud resources.
        Default value is `public`.
        Possible values are `private` and `public`.
        """
        return pulumi.get(self, "visibility")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

