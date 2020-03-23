# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Firewall(pulumi.CustomResource):
    allows: pulumi.Output[list]
    """
    The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
    permitted connection.

      * `ports` (`list`)
      * `protocol` (`str`)
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    denies: pulumi.Output[list]
    """
    The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
    denied connection.

      * `ports` (`list`)
      * `protocol` (`str`)
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property when you create the resource.
    """
    destination_ranges: pulumi.Output[list]
    """
    If destination ranges are specified, the firewall will apply only to traffic that has destination IP address in these
    ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
    """
    direction: pulumi.Output[str]
    """
    Direction of traffic to which this firewall applies; default is INGRESS. Note: For INGRESS traffic, it is NOT supported
    to specify destinationRanges; For EGRESS traffic, it is NOT supported to specify sourceRanges OR sourceTags.
    """
    disabled: pulumi.Output[bool]
    """
    Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with. When set to true,
    the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall
    rule will be enabled.
    """
    enable_logging: pulumi.Output[bool]
    """
    This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be
    exported to Stackdriver.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
    comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
    '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
    must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
    """
    network: pulumi.Output[str]
    """
    The name or self_link of the network to attach this firewall to.
    """
    priority: pulumi.Output[float]
    """
    Priority for this rule. This is an integer between 0 and 65535, both inclusive. When not specified, the value assumed is
    1000. Relative priorities determine precedence of conflicting rules. Lower value of priority implies higher precedence
    (eg, a rule with priority 0 has higher precedence than a rule with priority 1). DENY rules take precedence over ALLOW
    rules having equal priority.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    source_ranges: pulumi.Output[list]
    """
    If source ranges are specified, the firewall will apply only to traffic that has source IP address in these ranges.
    These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both properties
    are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP that belongs
    to a tag listed in the sourceTags property. The connection does not need to match both properties for the firewall to
    apply. Only IPv4 is supported.
    """
    source_service_accounts: pulumi.Output[list]
    """
    If source service accounts are specified, the firewall will apply only to traffic originating from an instance with a
    service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP
    address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same
    time as sourceServiceAccounts. If both are set, the firewall will apply to traffic that has source IP address within
    sourceRanges OR the source IP belongs to an instance with service account listed in sourceServiceAccount. The connection
    does not need to match both properties for the firewall to apply. sourceServiceAccounts cannot be used at the same time
    as sourceTags or targetTags.
    """
    source_tags: pulumi.Output[list]
    """
    If source tags are specified, the firewall will apply only to traffic with source IP that belongs to a tag listed in
    source tags. Source tags cannot be used to control traffic to an instance's external IP address. Because tags are
    associated with an instance, not an IP address. One or both of sourceRanges and sourceTags may be set. If both
    properties are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP
    that belongs to a tag listed in the sourceTags property. The connection does not need to match both properties for the
    firewall to apply.
    """
    target_service_accounts: pulumi.Output[list]
    """
    A list of service accounts indicating sets of instances located in the network that may make network connections as
    specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither
    targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
    """
    target_tags: pulumi.Output[list]
    """
    A list of instance tags indicating sets of instances located in the network that may make network connections as
    specified in allowed[]. If no targetTags are specified, the firewall rule applies to all instances on the specified
    network.
    """
    def __init__(__self__, resource_name, opts=None, allows=None, denies=None, description=None, destination_ranges=None, direction=None, disabled=None, enable_logging=None, name=None, network=None, priority=None, project=None, source_ranges=None, source_service_accounts=None, source_tags=None, target_service_accounts=None, target_tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Each network has its own firewall controlling access to and from the
        instances.

        All traffic to instances, even from other instances, is blocked by the
        firewall unless firewall rules are created to allow it.

        The default network has automatically created firewall rules that are
        shown in default firewall rules. No manually created network has
        automatically created firewall rules except for a default "allow" rule for
        outgoing traffic and a default "deny" for incoming traffic. For all
        networks except the default network, you must create any firewall rules
        you need.


        To get more information about Firewall, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/firewalls)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vpc/docs/firewalls)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_firewall.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allows: The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
               permitted connection.
        :param pulumi.Input[list] denies: The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
               denied connection.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[list] destination_ranges: If destination ranges are specified, the firewall will apply only to traffic that has destination IP address in these
               ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
        :param pulumi.Input[str] direction: Direction of traffic to which this firewall applies; default is INGRESS. Note: For INGRESS traffic, it is NOT supported
               to specify destinationRanges; For EGRESS traffic, it is NOT supported to specify sourceRanges OR sourceTags.
        :param pulumi.Input[bool] disabled: Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with. When set to true,
               the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall
               rule will be enabled.
        :param pulumi.Input[bool] enable_logging: This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be
               exported to Stackdriver.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
               '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The name or self_link of the network to attach this firewall to.
        :param pulumi.Input[float] priority: Priority for this rule. This is an integer between 0 and 65535, both inclusive. When not specified, the value assumed is
               1000. Relative priorities determine precedence of conflicting rules. Lower value of priority implies higher precedence
               (eg, a rule with priority 0 has higher precedence than a rule with priority 1). DENY rules take precedence over ALLOW
               rules having equal priority.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] source_ranges: If source ranges are specified, the firewall will apply only to traffic that has source IP address in these ranges.
               These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both properties
               are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP that belongs
               to a tag listed in the sourceTags property. The connection does not need to match both properties for the firewall to
               apply. Only IPv4 is supported.
        :param pulumi.Input[list] source_service_accounts: If source service accounts are specified, the firewall will apply only to traffic originating from an instance with a
               service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP
               address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same
               time as sourceServiceAccounts. If both are set, the firewall will apply to traffic that has source IP address within
               sourceRanges OR the source IP belongs to an instance with service account listed in sourceServiceAccount. The connection
               does not need to match both properties for the firewall to apply. sourceServiceAccounts cannot be used at the same time
               as sourceTags or targetTags.
        :param pulumi.Input[list] source_tags: If source tags are specified, the firewall will apply only to traffic with source IP that belongs to a tag listed in
               source tags. Source tags cannot be used to control traffic to an instance's external IP address. Because tags are
               associated with an instance, not an IP address. One or both of sourceRanges and sourceTags may be set. If both
               properties are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP
               that belongs to a tag listed in the sourceTags property. The connection does not need to match both properties for the
               firewall to apply.
        :param pulumi.Input[list] target_service_accounts: A list of service accounts indicating sets of instances located in the network that may make network connections as
               specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither
               targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
        :param pulumi.Input[list] target_tags: A list of instance tags indicating sets of instances located in the network that may make network connections as
               specified in allowed[]. If no targetTags are specified, the firewall rule applies to all instances on the specified
               network.

        The **allows** object supports the following:

          * `ports` (`pulumi.Input[list]`)
          * `protocol` (`pulumi.Input[str]`)

        The **denies** object supports the following:

          * `ports` (`pulumi.Input[list]`)
          * `protocol` (`pulumi.Input[str]`)
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

            __props__['allows'] = allows
            __props__['denies'] = denies
            __props__['description'] = description
            __props__['destination_ranges'] = destination_ranges
            __props__['direction'] = direction
            __props__['disabled'] = disabled
            __props__['enable_logging'] = enable_logging
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
            __props__['priority'] = priority
            __props__['project'] = project
            __props__['source_ranges'] = source_ranges
            __props__['source_service_accounts'] = source_service_accounts
            __props__['source_tags'] = source_tags
            __props__['target_service_accounts'] = target_service_accounts
            __props__['target_tags'] = target_tags
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(Firewall, __self__).__init__(
            'gcp:compute/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allows=None, creation_timestamp=None, denies=None, description=None, destination_ranges=None, direction=None, disabled=None, enable_logging=None, name=None, network=None, priority=None, project=None, self_link=None, source_ranges=None, source_service_accounts=None, source_tags=None, target_service_accounts=None, target_tags=None):
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allows: The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
               permitted connection.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[list] denies: The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a
               denied connection.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when you create the resource.
        :param pulumi.Input[list] destination_ranges: If destination ranges are specified, the firewall will apply only to traffic that has destination IP address in these
               ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.
        :param pulumi.Input[str] direction: Direction of traffic to which this firewall applies; default is INGRESS. Note: For INGRESS traffic, it is NOT supported
               to specify destinationRanges; For EGRESS traffic, it is NOT supported to specify sourceRanges OR sourceTags.
        :param pulumi.Input[bool] disabled: Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with. When set to true,
               the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall
               rule will be enabled.
        :param pulumi.Input[bool] enable_logging: This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be
               exported to Stackdriver.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
               '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
               must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] network: The name or self_link of the network to attach this firewall to.
        :param pulumi.Input[float] priority: Priority for this rule. This is an integer between 0 and 65535, both inclusive. When not specified, the value assumed is
               1000. Relative priorities determine precedence of conflicting rules. Lower value of priority implies higher precedence
               (eg, a rule with priority 0 has higher precedence than a rule with priority 1). DENY rules take precedence over ALLOW
               rules having equal priority.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[list] source_ranges: If source ranges are specified, the firewall will apply only to traffic that has source IP address in these ranges.
               These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both properties
               are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP that belongs
               to a tag listed in the sourceTags property. The connection does not need to match both properties for the firewall to
               apply. Only IPv4 is supported.
        :param pulumi.Input[list] source_service_accounts: If source service accounts are specified, the firewall will apply only to traffic originating from an instance with a
               service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP
               address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same
               time as sourceServiceAccounts. If both are set, the firewall will apply to traffic that has source IP address within
               sourceRanges OR the source IP belongs to an instance with service account listed in sourceServiceAccount. The connection
               does not need to match both properties for the firewall to apply. sourceServiceAccounts cannot be used at the same time
               as sourceTags or targetTags.
        :param pulumi.Input[list] source_tags: If source tags are specified, the firewall will apply only to traffic with source IP that belongs to a tag listed in
               source tags. Source tags cannot be used to control traffic to an instance's external IP address. Because tags are
               associated with an instance, not an IP address. One or both of sourceRanges and sourceTags may be set. If both
               properties are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP
               that belongs to a tag listed in the sourceTags property. The connection does not need to match both properties for the
               firewall to apply.
        :param pulumi.Input[list] target_service_accounts: A list of service accounts indicating sets of instances located in the network that may make network connections as
               specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither
               targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.
        :param pulumi.Input[list] target_tags: A list of instance tags indicating sets of instances located in the network that may make network connections as
               specified in allowed[]. If no targetTags are specified, the firewall rule applies to all instances on the specified
               network.

        The **allows** object supports the following:

          * `ports` (`pulumi.Input[list]`)
          * `protocol` (`pulumi.Input[str]`)

        The **denies** object supports the following:

          * `ports` (`pulumi.Input[list]`)
          * `protocol` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allows"] = allows
        __props__["creation_timestamp"] = creation_timestamp
        __props__["denies"] = denies
        __props__["description"] = description
        __props__["destination_ranges"] = destination_ranges
        __props__["direction"] = direction
        __props__["disabled"] = disabled
        __props__["enable_logging"] = enable_logging
        __props__["name"] = name
        __props__["network"] = network
        __props__["priority"] = priority
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["source_ranges"] = source_ranges
        __props__["source_service_accounts"] = source_service_accounts
        __props__["source_tags"] = source_tags
        __props__["target_service_accounts"] = target_service_accounts
        __props__["target_tags"] = target_tags
        return Firewall(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

