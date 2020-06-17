# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Subnetwork(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property when
    you create the resource. This field can be set only at resource
    creation time.
    """
    fingerprint: pulumi.Output[str]
    """
    Fingerprint of this resource. This field is used internally during updates of this resource.
    """
    gateway_address: pulumi.Output[str]
    """
    The gateway address for default routes to reach destination addresses outside this subnetwork.
    """
    ip_cidr_range: pulumi.Output[str]
    """
    The range of IP addresses belonging to this subnetwork secondary
    range. Provide this property when you create the subnetwork.
    Ranges must be unique and non-overlapping with all primary and
    secondary IP ranges within a network. Only IPv4 is supported.
    """
    log_config: pulumi.Output[dict]
    """
    Denotes the logging options for the subnetwork flow logs. If logging is enabled
    logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
    subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.

      * `aggregationInterval` (`str`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
        Toggles the aggregation interval for collecting flow logs. Increasing the
        interval time will reduce the amount of generated flow logs for long
        lasting connections. Default is an interval of 5 seconds per connection.
        Possible values are INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN,
        INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN
      * `flowSampling` (`float`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
        The value of the field must be in [0, 1]. Set the sampling rate of VPC
        flow logs within the subnetwork where 1.0 means all collected logs are
        reported and 0.0 means no logs are reported. Default is 0.5 which means
        half of all collected logs are reported.
      * `metadata` (`str`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
        Configures whether metadata fields should be added to the reported VPC
        flow logs.
    """
    name: pulumi.Output[str]
    """
    The name of the resource, provided by the client when initially
    creating the resource. The name must be 1-63 characters long, and
    comply with RFC1035. Specifically, the name must be 1-63 characters
    long and match the regular expression `a-z?` which
    means the first character must be a lowercase letter, and all
    following characters must be a dash, lowercase letter, or digit,
    except the last character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    The network this subnet belongs to.
    Only networks that are in the distributed mode can have subnetworks.
    """
    private_ip_google_access: pulumi.Output[bool]
    """
    When enabled, VMs in this subnetwork without external IP addresses can
    access Google APIs and services by using Private Google Access.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    purpose: pulumi.Output[str]
    """
    The purpose of the resource. This field can be either PRIVATE
    or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
    INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
    reserved for Internal HTTP(S) Load Balancing. If unspecified, the
    purpose defaults to PRIVATE.
    If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
    """
    region: pulumi.Output[str]
    """
    The GCP region for this subnetwork.
    """
    role: pulumi.Output[str]
    """
    The role of subnetwork. Currently, this field is only used when
    purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
    or BACKUP. An ACTIVE subnetwork is one that is currently being used
    for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
    is ready to be promoted to ACTIVE or is currently draining.
    """
    secondary_ip_ranges: pulumi.Output[list]
    """
    An array of configurations for secondary IP ranges for VM instances
    contained in this subnetwork. The primary IP of such VM must belong
    to the primary ipCidrRange of the subnetwork. The alias IPs may belong
    to either primary or secondary ranges. Structure is documented below.

      * `ip_cidr_range` (`str`) - The range of IP addresses belonging to this subnetwork secondary
        range. Provide this property when you create the subnetwork.
        Ranges must be unique and non-overlapping with all primary and
        secondary IP ranges within a network. Only IPv4 is supported.
      * `rangeName` (`str`) - The name associated with this subnetwork secondary range, used
        when adding an alias IP range to a VM instance. The name must
        be 1-63 characters long, and comply with RFC1035. The name
        must be unique within the subnetwork.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, ip_cidr_range=None, log_config=None, name=None, network=None, private_ip_google_access=None, project=None, purpose=None, region=None, role=None, secondary_ip_ranges=None, __props__=None, __name__=None, __opts__=None):
        """
        A VPC network is a virtual version of the traditional physical networks
        that exist within and between physical data centers. A VPC network
        provides connectivity for your Compute Engine virtual machine (VM)
        instances, Container Engine containers, App Engine Flex services, and
        other network-related resources.

        Each GCP project contains one or more VPC networks. Each VPC network is a
        global entity spanning all GCP regions. This global VPC network allows VM
        instances and other resources to communicate with each other via internal,
        private IP addresses.

        Each VPC network is subdivided into subnets, and each subnet is contained
        within a single region. You can have more than one subnet in a region for
        a given VPC network. Each subnet has a contiguous private RFC1918 IP
        space. You create instances, containers, and the like in these subnets.
        When you create an instance, you must create it in a subnet, and the
        instance draws its internal IP address from that subnet.

        Virtual machine (VM) instances in a VPC network can communicate with
        instances in all other subnets of the same VPC network, regardless of
        region, using their RFC1918 private IP addresses. You can isolate portions
        of the network, even entire subnets, using firewall rules.


        To get more information about Subnetwork, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
        * How-to Guides
            * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
            * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)

        ## Example Usage

        ### Subnetwork Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        custom_test = gcp.compute.Network("custom-test", auto_create_subnetworks=False)
        network_with_private_secondary_ip_ranges = gcp.compute.Subnetwork("network-with-private-secondary-ip-ranges",
            ip_cidr_range="10.2.0.0/16",
            region="us-central1",
            network=custom_test.id,
            secondary_ip_range=[{
                "rangeName": "tf-test-secondary-range-update1",
                "ip_cidr_range": "192.168.10.0/24",
            }])
        ```

        ### Subnetwork Logging Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        custom_test = gcp.compute.Network("custom-test", auto_create_subnetworks=False)
        subnet_with_logging = gcp.compute.Subnetwork("subnet-with-logging",
            ip_cidr_range="10.2.0.0/16",
            region="us-central1",
            network=custom_test.id,
            log_config={
                "aggregationInterval": "INTERVAL_10_MIN",
                "flowSampling": 0.5,
                "metadata": "INCLUDE_ALL_METADATA",
            })
        ```

        ### Subnetwork Internal L7lb

        ```python
        import pulumi
        import pulumi_gcp as gcp

        custom_test = gcp.compute.Network("custom-test", auto_create_subnetworks=False)
        network_for_l7lb = gcp.compute.Subnetwork("network-for-l7lb",
            ip_cidr_range="10.0.0.0/22",
            region="us-central1",
            purpose="INTERNAL_HTTPS_LOAD_BALANCER",
            role="ACTIVE",
            network=custom_test.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource. This field can be set only at resource
               creation time.
        :param pulumi.Input[str] ip_cidr_range: The range of IP addresses belonging to this subnetwork secondary
               range. Provide this property when you create the subnetwork.
               Ranges must be unique and non-overlapping with all primary and
               secondary IP ranges within a network. Only IPv4 is supported.
        :param pulumi.Input[dict] log_config: Denotes the logging options for the subnetwork flow logs. If logging is enabled
               logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
               subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially
               creating the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The network this subnet belongs to.
               Only networks that are in the distributed mode can have subnetworks.
        :param pulumi.Input[bool] private_ip_google_access: When enabled, VMs in this subnetwork without external IP addresses can
               access Google APIs and services by using Private Google Access.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. This field can be either PRIVATE
               or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
               INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
               reserved for Internal HTTP(S) Load Balancing. If unspecified, the
               purpose defaults to PRIVATE.
               If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
        :param pulumi.Input[str] role: The role of subnetwork. Currently, this field is only used when
               purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
               or BACKUP. An ACTIVE subnetwork is one that is currently being used
               for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
               is ready to be promoted to ACTIVE or is currently draining.
        :param pulumi.Input[list] secondary_ip_ranges: An array of configurations for secondary IP ranges for VM instances
               contained in this subnetwork. The primary IP of such VM must belong
               to the primary ipCidrRange of the subnetwork. The alias IPs may belong
               to either primary or secondary ranges. Structure is documented below.

        The **log_config** object supports the following:

          * `aggregationInterval` (`pulumi.Input[str]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            Toggles the aggregation interval for collecting flow logs. Increasing the
            interval time will reduce the amount of generated flow logs for long
            lasting connections. Default is an interval of 5 seconds per connection.
            Possible values are INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN,
            INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN
          * `flowSampling` (`pulumi.Input[float]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            The value of the field must be in [0, 1]. Set the sampling rate of VPC
            flow logs within the subnetwork where 1.0 means all collected logs are
            reported and 0.0 means no logs are reported. Default is 0.5 which means
            half of all collected logs are reported.
          * `metadata` (`pulumi.Input[str]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            Configures whether metadata fields should be added to the reported VPC
            flow logs.

        The **secondary_ip_ranges** object supports the following:

          * `ip_cidr_range` (`pulumi.Input[str]`) - The range of IP addresses belonging to this subnetwork secondary
            range. Provide this property when you create the subnetwork.
            Ranges must be unique and non-overlapping with all primary and
            secondary IP ranges within a network. Only IPv4 is supported.
          * `rangeName` (`pulumi.Input[str]`) - The name associated with this subnetwork secondary range, used
            when adding an alias IP range to a VM instance. The name must
            be 1-63 characters long, and comply with RFC1035. The name
            must be unique within the subnetwork.
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

            __props__['description'] = description
            if ip_cidr_range is None:
                raise TypeError("Missing required property 'ip_cidr_range'")
            __props__['ip_cidr_range'] = ip_cidr_range
            __props__['log_config'] = log_config
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['private_ip_google_access'] = private_ip_google_access
            __props__['project'] = project
            __props__['purpose'] = purpose
            __props__['region'] = region
            __props__['role'] = role
            __props__['secondary_ip_ranges'] = secondary_ip_ranges
            __props__['creation_timestamp'] = None
            __props__['fingerprint'] = None
            __props__['gateway_address'] = None
            __props__['self_link'] = None
        super(Subnetwork, __self__).__init__(
            'gcp:compute/subnetwork:Subnetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, fingerprint=None, gateway_address=None, ip_cidr_range=None, log_config=None, name=None, network=None, private_ip_google_access=None, project=None, purpose=None, region=None, role=None, secondary_ip_ranges=None, self_link=None):
        """
        Get an existing Subnetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource. This field can be set only at resource
               creation time.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. This field is used internally during updates of this resource.
        :param pulumi.Input[str] gateway_address: The gateway address for default routes to reach destination addresses outside this subnetwork.
        :param pulumi.Input[str] ip_cidr_range: The range of IP addresses belonging to this subnetwork secondary
               range. Provide this property when you create the subnetwork.
               Ranges must be unique and non-overlapping with all primary and
               secondary IP ranges within a network. Only IPv4 is supported.
        :param pulumi.Input[dict] log_config: Denotes the logging options for the subnetwork flow logs. If logging is enabled
               logs will be exported to Stackdriver. This field cannot be set if the `purpose` of this
               subnetwork is `INTERNAL_HTTPS_LOAD_BALANCER`  Structure is documented below.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially
               creating the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The network this subnet belongs to.
               Only networks that are in the distributed mode can have subnetworks.
        :param pulumi.Input[bool] private_ip_google_access: When enabled, VMs in this subnetwork without external IP addresses can
               access Google APIs and services by using Private Google Access.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. This field can be either PRIVATE
               or INTERNAL_HTTPS_LOAD_BALANCER. A subnetwork with purpose set to
               INTERNAL_HTTPS_LOAD_BALANCER is a user-created subnetwork that is
               reserved for Internal HTTP(S) Load Balancing. If unspecified, the
               purpose defaults to PRIVATE.
               If set to INTERNAL_HTTPS_LOAD_BALANCER you must also set the role.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
        :param pulumi.Input[str] role: The role of subnetwork. Currently, this field is only used when
               purpose = INTERNAL_HTTPS_LOAD_BALANCER. The value can be set to ACTIVE
               or BACKUP. An ACTIVE subnetwork is one that is currently being used
               for Internal HTTP(S) Load Balancing. A BACKUP subnetwork is one that
               is ready to be promoted to ACTIVE or is currently draining.
        :param pulumi.Input[list] secondary_ip_ranges: An array of configurations for secondary IP ranges for VM instances
               contained in this subnetwork. The primary IP of such VM must belong
               to the primary ipCidrRange of the subnetwork. The alias IPs may belong
               to either primary or secondary ranges. Structure is documented below.
        :param pulumi.Input[str] self_link: The URI of the created resource.

        The **log_config** object supports the following:

          * `aggregationInterval` (`pulumi.Input[str]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            Toggles the aggregation interval for collecting flow logs. Increasing the
            interval time will reduce the amount of generated flow logs for long
            lasting connections. Default is an interval of 5 seconds per connection.
            Possible values are INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN,
            INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN
          * `flowSampling` (`pulumi.Input[float]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            The value of the field must be in [0, 1]. Set the sampling rate of VPC
            flow logs within the subnetwork where 1.0 means all collected logs are
            reported and 0.0 means no logs are reported. Default is 0.5 which means
            half of all collected logs are reported.
          * `metadata` (`pulumi.Input[str]`) - Can only be specified if VPC flow logging for this subnetwork is enabled.
            Configures whether metadata fields should be added to the reported VPC
            flow logs.

        The **secondary_ip_ranges** object supports the following:

          * `ip_cidr_range` (`pulumi.Input[str]`) - The range of IP addresses belonging to this subnetwork secondary
            range. Provide this property when you create the subnetwork.
            Ranges must be unique and non-overlapping with all primary and
            secondary IP ranges within a network. Only IPv4 is supported.
          * `rangeName` (`pulumi.Input[str]`) - The name associated with this subnetwork secondary range, used
            when adding an alias IP range to a VM instance. The name must
            be 1-63 characters long, and comply with RFC1035. The name
            must be unique within the subnetwork.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["fingerprint"] = fingerprint
        __props__["gateway_address"] = gateway_address
        __props__["ip_cidr_range"] = ip_cidr_range
        __props__["log_config"] = log_config
        __props__["name"] = name
        __props__["network"] = network
        __props__["private_ip_google_access"] = private_ip_google_access
        __props__["project"] = project
        __props__["purpose"] = purpose
        __props__["region"] = region
        __props__["role"] = role
        __props__["secondary_ip_ranges"] = secondary_ip_ranges
        __props__["self_link"] = self_link
        return Subnetwork(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
