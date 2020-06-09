# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PacketMirroring(pulumi.CustomResource):
    collector_ilb: pulumi.Output[dict]
    """
    The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
    that will be used as collector for mirrored traffic. The
    specified forwarding rule must have is_mirroring_collector
    set to true.  Structure is documented below.

      * `url` (`str`) - The URL of the instances where this rule should be active.
    """
    description: pulumi.Output[str]
    """
    A human-readable description of the rule.
    """
    filter: pulumi.Output[dict]
    """
    A filter for mirrored traffic.  If unset, all traffic is mirrored.  Structure is documented below.

      * `cidrRanges` (`list`) - IP CIDR ranges that apply as a filter on the source (ingress) or
        destination (egress) IP in the IP header. Only IPv4 is supported.
      * `ipProtocols` (`list`) - Protocols that apply as a filter on mirrored traffic.
    """
    mirrored_resources: pulumi.Output[dict]
    """
    A means of specifying which resources to mirror.  Structure is documented below.

      * `instances` (`list`) - All the listed instances will be mirrored.  Specify at most 50.  Structure is documented below.
        * `url` (`str`) - The URL of the instances where this rule should be active.

      * `subnetworks` (`list`) - All instances in one of these subnetworks will be mirrored.  Structure is documented below.
        * `url` (`str`) - The URL of the instances where this rule should be active.

      * `tags` (`list`) - All instances with these tags will be mirrored.
    """
    name: pulumi.Output[str]
    """
    The name of the packet mirroring rule
    """
    network: pulumi.Output[dict]
    """
    Specifies the mirrored VPC network. Only packets in this network
    will be mirrored. All mirrored VMs should have a NIC in the given
    network. All mirrored subnetworks should belong to the given network.  Structure is documented below.

      * `url` (`str`) - The URL of the instances where this rule should be active.
    """
    priority: pulumi.Output[float]
    """
    Since only one rule can be active at a time, priority is
    used to break ties in the case of two rules that apply to
    the same instances.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The Region in which the created address should reside.
    If it is not provided, the provider region is used.
    """
    def __init__(__self__, resource_name, opts=None, collector_ilb=None, description=None, filter=None, mirrored_resources=None, name=None, network=None, priority=None, project=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Packet Mirroring mirrors traffic to and from particular VM instances.
        You can use the collected traffic to help you detect security threats
        and monitor application performance.

        To get more information about PacketMirroring, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/packetMirroring)
        * How-to Guides
            * [Using Packet Mirroring](https://cloud.google.com/vpc/docs/using-packet-mirroring#creating)

        ## Example Usage

        ### Compute Packet Mirroring Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.Network("defaultNetwork")
        mirror = gcp.compute.Instance("mirror",
            machine_type="n1-standard-1",
            boot_disk={
                "initialize_params": {
                    "image": "debian-cloud/debian-9",
                },
            },
            network_interface=[{
                "network": default_network.id,
                "access_config": [{}],
            }])
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            network=default_network.id,
            ip_cidr_range="10.2.0.0/16")
        default_health_check = gcp.compute.HealthCheck("defaultHealthCheck",
            check_interval_sec=1,
            timeout_sec=1,
            tcp_health_check={
                "port": "80",
            })
        default_region_backend_service = gcp.compute.RegionBackendService("defaultRegionBackendService", health_checks=[default_health_check.id])
        default_forwarding_rule = gcp.compute.ForwardingRule("defaultForwardingRule",
            is_mirroring_collector=True,
            ip_protocol="TCP",
            load_balancing_scheme="INTERNAL",
            backend_service=default_region_backend_service.id,
            all_ports=True,
            network=default_network.id,
            subnetwork=default_subnetwork.id,
            network_tier="PREMIUM")
        foobar = gcp.compute.PacketMirroring("foobar",
            description="bar",
            network={
                "url": default_network.id,
            },
            collector_ilb={
                "url": default_forwarding_rule.id,
            },
            mirrored_resources={
                "tags": ["foo"],
                "instances": [{
                    "url": mirror.id,
                }],
            },
            filter={
                "ipProtocols": ["tcp"],
                "cidrRanges": ["0.0.0.0/0"],
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] collector_ilb: The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
               that will be used as collector for mirrored traffic. The
               specified forwarding rule must have is_mirroring_collector
               set to true.  Structure is documented below.
        :param pulumi.Input[str] description: A human-readable description of the rule.
        :param pulumi.Input[dict] filter: A filter for mirrored traffic.  If unset, all traffic is mirrored.  Structure is documented below.
        :param pulumi.Input[dict] mirrored_resources: A means of specifying which resources to mirror.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the packet mirroring rule
        :param pulumi.Input[dict] network: Specifies the mirrored VPC network. Only packets in this network
               will be mirrored. All mirrored VMs should have a NIC in the given
               network. All mirrored subnetworks should belong to the given network.  Structure is documented below.
        :param pulumi.Input[float] priority: Since only one rule can be active at a time, priority is
               used to break ties in the case of two rules that apply to
               the same instances.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created address should reside.
               If it is not provided, the provider region is used.

        The **collector_ilb** object supports the following:

          * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

        The **filter** object supports the following:

          * `cidrRanges` (`pulumi.Input[list]`) - IP CIDR ranges that apply as a filter on the source (ingress) or
            destination (egress) IP in the IP header. Only IPv4 is supported.
          * `ipProtocols` (`pulumi.Input[list]`) - Protocols that apply as a filter on mirrored traffic.

        The **mirrored_resources** object supports the following:

          * `instances` (`pulumi.Input[list]`) - All the listed instances will be mirrored.  Specify at most 50.  Structure is documented below.
            * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

          * `subnetworks` (`pulumi.Input[list]`) - All instances in one of these subnetworks will be mirrored.  Structure is documented below.
            * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

          * `tags` (`pulumi.Input[list]`) - All instances with these tags will be mirrored.

        The **network** object supports the following:

          * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.
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

            if collector_ilb is None:
                raise TypeError("Missing required property 'collector_ilb'")
            __props__['collector_ilb'] = collector_ilb
            __props__['description'] = description
            __props__['filter'] = filter
            if mirrored_resources is None:
                raise TypeError("Missing required property 'mirrored_resources'")
            __props__['mirrored_resources'] = mirrored_resources
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['priority'] = priority
            __props__['project'] = project
            __props__['region'] = region
        super(PacketMirroring, __self__).__init__(
            'gcp:compute/packetMirroring:PacketMirroring',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, collector_ilb=None, description=None, filter=None, mirrored_resources=None, name=None, network=None, priority=None, project=None, region=None):
        """
        Get an existing PacketMirroring resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] collector_ilb: The Forwarding Rule resource (of type load_balancing_scheme=INTERNAL)
               that will be used as collector for mirrored traffic. The
               specified forwarding rule must have is_mirroring_collector
               set to true.  Structure is documented below.
        :param pulumi.Input[str] description: A human-readable description of the rule.
        :param pulumi.Input[dict] filter: A filter for mirrored traffic.  If unset, all traffic is mirrored.  Structure is documented below.
        :param pulumi.Input[dict] mirrored_resources: A means of specifying which resources to mirror.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the packet mirroring rule
        :param pulumi.Input[dict] network: Specifies the mirrored VPC network. Only packets in this network
               will be mirrored. All mirrored VMs should have a NIC in the given
               network. All mirrored subnetworks should belong to the given network.  Structure is documented below.
        :param pulumi.Input[float] priority: Since only one rule can be active at a time, priority is
               used to break ties in the case of two rules that apply to
               the same instances.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The Region in which the created address should reside.
               If it is not provided, the provider region is used.

        The **collector_ilb** object supports the following:

          * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

        The **filter** object supports the following:

          * `cidrRanges` (`pulumi.Input[list]`) - IP CIDR ranges that apply as a filter on the source (ingress) or
            destination (egress) IP in the IP header. Only IPv4 is supported.
          * `ipProtocols` (`pulumi.Input[list]`) - Protocols that apply as a filter on mirrored traffic.

        The **mirrored_resources** object supports the following:

          * `instances` (`pulumi.Input[list]`) - All the listed instances will be mirrored.  Specify at most 50.  Structure is documented below.
            * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

          * `subnetworks` (`pulumi.Input[list]`) - All instances in one of these subnetworks will be mirrored.  Structure is documented below.
            * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.

          * `tags` (`pulumi.Input[list]`) - All instances with these tags will be mirrored.

        The **network** object supports the following:

          * `url` (`pulumi.Input[str]`) - The URL of the instances where this rule should be active.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collector_ilb"] = collector_ilb
        __props__["description"] = description
        __props__["filter"] = filter
        __props__["mirrored_resources"] = mirrored_resources
        __props__["name"] = name
        __props__["network"] = network
        __props__["priority"] = priority
        __props__["project"] = project
        __props__["region"] = region
        return PacketMirroring(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

