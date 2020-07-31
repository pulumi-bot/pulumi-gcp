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

__all__ = ['GlobalForwardingRule']


class GlobalForwardingRule(pulumi.CustomResource):
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    An optional description of this resource. Provide this property when
    you create the resource.
    """
    ip_address: pulumi.Output[str] = pulumi.output_property("ipAddress")
    """
    The IP address that this forwarding rule is serving on behalf of.
    Addresses are restricted based on the forwarding rule's load balancing
    scheme (EXTERNAL or INTERNAL) and scope (global or regional).
    When the load balancing scheme is EXTERNAL, for global forwarding
    rules, the address must be a global IP, and for regional forwarding
    rules, the address must live in the same region as the forwarding
    rule. If this field is empty, an ephemeral IPv4 address from the same
    scope (global or regional) will be assigned. A regional forwarding
    rule supports IPv4 only. A global forwarding rule supports either IPv4
    or IPv6.
    When the load balancing scheme is INTERNAL, this can only be an RFC
    1918 IP address belonging to the network/subnet configured for the
    forwarding rule. By default, if this field is empty, an ephemeral
    internal IP address will be automatically allocated from the IP range
    of the subnet or network configured for this forwarding rule.
    An address must be specified by a literal IP address. > **NOTE**: While
    the API allows you to specify various resource paths for an address resource
    instead, this provider requires this to specifically be an IP address to
    avoid needing to fetching the IP address from resource paths on refresh
    or unnecessary diffs.
    """
    ip_protocol: pulumi.Output[str] = pulumi.output_property("ipProtocol")
    """
    The IP protocol to which this rule applies. When the load balancing scheme is
    INTERNAL_SELF_MANAGED, only TCP is valid.
    """
    ip_version: pulumi.Output[Optional[str]] = pulumi.output_property("ipVersion")
    """
    The IP Version that will be used by this global forwarding rule.
    """
    label_fingerprint: pulumi.Output[str] = pulumi.output_property("labelFingerprint")
    """
    The fingerprint used for optimistic locking of this resource. Used internally during updates.
    """
    labels: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("labels")
    """
    Labels to apply to this forwarding rule.  A list of key->value pairs.
    """
    load_balancing_scheme: pulumi.Output[Optional[str]] = pulumi.output_property("loadBalancingScheme")
    """
    This signifies what the GlobalForwardingRule will be used for.
    The value of INTERNAL_SELF_MANAGED means that this will be used for
    Internal Global HTTP(S) LB. The value of EXTERNAL means that this
    will be used for External Global Load Balancing (HTTP(S) LB,
    External TCP/UDP LB, SSL Proxy)
    NOTE: Currently global forwarding rules cannot be used for INTERNAL
    load balancing.
    """
    metadata_filters: pulumi.Output[Optional[List['outputs.GlobalForwardingRuleMetadataFilter']]] = pulumi.output_property("metadataFilters")
    """
    Opaque filter criteria used by Loadbalancer to restrict routing
    configuration to a limited set xDS compliant clients. In their xDS
    requests to Loadbalancer, xDS clients present node metadata. If a
    match takes place, the relevant routing configuration is made available
    to those proxies.
    For each metadataFilter in this list, if its filterMatchCriteria is set
    to MATCH_ANY, at least one of the filterLabels must match the
    corresponding label provided in the metadata. If its filterMatchCriteria
    is set to MATCH_ALL, then all of its filterLabels must match with
    corresponding labels in the provided metadata.
    metadataFilters specified here can be overridden by those specified in
    the UrlMap that this ForwardingRule references.
    metadataFilters only applies to Loadbalancers that have their
    loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Name of the metadata label. The length must be between
    1 and 1024 characters, inclusive.
    """
    network: pulumi.Output[str] = pulumi.output_property("network")
    """
    This field is not used for external load balancing.
    For INTERNAL_SELF_MANAGED load balancing, this field
    identifies the network that the load balanced IP should belong to
    for this global forwarding rule. If this field is not specified,
    the default network will be used.
    """
    port_range: pulumi.Output[Optional[str]] = pulumi.output_property("portRange")
    """
    This field is used along with the target field for TargetHttpProxy,
    TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
    TargetPool, TargetInstance.
    Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
    addressed to ports in the specified range will be forwarded to target.
    Forwarding rules with the same [IPAddress, IPProtocol] pair must have
    disjoint port ranges.
    Some types of forwarding target have constraints on the acceptable
    ports:
    * TargetHttpProxy: 80, 8080
    * TargetHttpsProxy: 443
    * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
    1883, 5222
    * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
    1883, 5222
    * TargetVpnGateway: 500, 4500
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str] = pulumi.output_property("selfLink")
    """
    The URI of the created resource.
    """
    target: pulumi.Output[str] = pulumi.output_property("target")
    """
    The URL of the target resource to receive the matched traffic.
    The forwarded traffic must be of a type appropriate to the target object.
    For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
    are valid.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, ip_address: Optional[pulumi.Input[str]] = None, ip_protocol: Optional[pulumi.Input[str]] = None, ip_version: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, load_balancing_scheme: Optional[pulumi.Input[str]] = None, metadata_filters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GlobalForwardingRuleMetadataFilterArgs']]]]] = None, name: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, port_range: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, target: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Represents a GlobalForwardingRule resource. Global forwarding rules are
        used to forward traffic to the correct load balancer for HTTP load
        balancing. Global forwarding rules can only be used for HTTP load
        balancing.

        For more information, see
        https://cloud.google.com/compute/docs/load-balancing/http/

        ## Example Usage
        ### Global Forwarding Rule Http

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_http_health_check = gcp.compute.HttpHealthCheck("defaultHttpHealthCheck",
            request_path="/",
            check_interval_sec=1,
            timeout_sec=1)
        default_backend_service = gcp.compute.BackendService("defaultBackendService",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id])
        default_url_map = gcp.compute.URLMap("defaultURLMap",
            description="a description",
            default_service=default_backend_service.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": default_backend_service.id,
                "pathRules": [{
                    "paths": ["/*"],
                    "service": default_backend_service.id,
                }],
            }])
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy",
            description="a description",
            url_map=default_url_map.id)
        default_global_forwarding_rule = gcp.compute.GlobalForwardingRule("defaultGlobalForwardingRule",
            target=default_target_http_proxy.id,
            port_range="80")
        ```
        ### Global Forwarding Rule Internal

        ```python
        import pulumi
        import pulumi_gcp as gcp

        debian_image = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        instance_template = gcp.compute.InstanceTemplate("instanceTemplate",
            machine_type="n1-standard-1",
            network_interfaces=[{
                "network": "default",
            }],
            disks=[{
                "sourceImage": debian_image.self_link,
                "autoDelete": True,
                "boot": True,
            }],
            opts=ResourceOptions(provider=google_beta))
        igm = gcp.compute.InstanceGroupManager("igm",
            versions=[{
                "instanceTemplate": instance_template.id,
                "name": "primary",
            }],
            base_instance_name="internal-glb",
            zone="us-central1-f",
            target_size=1,
            opts=ResourceOptions(provider=google_beta))
        default_health_check = gcp.compute.HealthCheck("defaultHealthCheck",
            check_interval_sec=1,
            timeout_sec=1,
            tcp_health_check={
                "port": "80",
            },
            opts=ResourceOptions(provider=google_beta))
        default_backend_service = gcp.compute.BackendService("defaultBackendService",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            load_balancing_scheme="INTERNAL_SELF_MANAGED",
            backends=[{
                "group": igm.instance_group,
                "balancingMode": "RATE",
                "capacityScaler": 0.4,
                "maxRatePerInstance": 50,
            }],
            health_checks=[default_health_check.id],
            opts=ResourceOptions(provider=google_beta))
        default_url_map = gcp.compute.URLMap("defaultURLMap",
            description="a description",
            default_service=default_backend_service.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": default_backend_service.id,
                "pathRules": [{
                    "paths": ["/*"],
                    "service": default_backend_service.id,
                }],
            }],
            opts=ResourceOptions(provider=google_beta))
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy",
            description="a description",
            url_map=default_url_map.id,
            opts=ResourceOptions(provider=google_beta))
        default_global_forwarding_rule = gcp.compute.GlobalForwardingRule("defaultGlobalForwardingRule",
            target=default_target_http_proxy.id,
            port_range="80",
            load_balancing_scheme="INTERNAL_SELF_MANAGED",
            ip_address="0.0.0.0",
            metadata_filters=[{
                "filterMatchCriteria": "MATCH_ANY",
                "filterLabels": [{
                    "name": "PLANET",
                    "value": "MARS",
                }],
            }],
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] ip_address: The IP address that this forwarding rule is serving on behalf of.
               Addresses are restricted based on the forwarding rule's load balancing
               scheme (EXTERNAL or INTERNAL) and scope (global or regional).
               When the load balancing scheme is EXTERNAL, for global forwarding
               rules, the address must be a global IP, and for regional forwarding
               rules, the address must live in the same region as the forwarding
               rule. If this field is empty, an ephemeral IPv4 address from the same
               scope (global or regional) will be assigned. A regional forwarding
               rule supports IPv4 only. A global forwarding rule supports either IPv4
               or IPv6.
               When the load balancing scheme is INTERNAL, this can only be an RFC
               1918 IP address belonging to the network/subnet configured for the
               forwarding rule. By default, if this field is empty, an ephemeral
               internal IP address will be automatically allocated from the IP range
               of the subnet or network configured for this forwarding rule.
               An address must be specified by a literal IP address. > **NOTE**: While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies. When the load balancing scheme is
               INTERNAL_SELF_MANAGED, only TCP is valid.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this global forwarding rule.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the GlobalForwardingRule will be used for.
               The value of INTERNAL_SELF_MANAGED means that this will be used for
               Internal Global HTTP(S) LB. The value of EXTERNAL means that this
               will be used for External Global Load Balancing (HTTP(S) LB,
               External TCP/UDP LB, SSL Proxy)
               NOTE: Currently global forwarding rules cannot be used for INTERNAL
               load balancing.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GlobalForwardingRuleMetadataFilterArgs']]]] metadata_filters: Opaque filter criteria used by Loadbalancer to restrict routing
               configuration to a limited set xDS compliant clients. In their xDS
               requests to Loadbalancer, xDS clients present node metadata. If a
               match takes place, the relevant routing configuration is made available
               to those proxies.
               For each metadataFilter in this list, if its filterMatchCriteria is set
               to MATCH_ANY, at least one of the filterLabels must match the
               corresponding label provided in the metadata. If its filterMatchCriteria
               is set to MATCH_ALL, then all of its filterLabels must match with
               corresponding labels in the provided metadata.
               metadataFilters specified here can be overridden by those specified in
               the UrlMap that this ForwardingRule references.
               metadataFilters only applies to Loadbalancers that have their
               loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the metadata label. The length must be between
               1 and 1024 characters, inclusive.
        :param pulumi.Input[str] network: This field is not used for external load balancing.
               For INTERNAL_SELF_MANAGED load balancing, this field
               identifies the network that the load balanced IP should belong to
               for this global forwarding rule. If this field is not specified,
               the default network will be used.
        :param pulumi.Input[str] port_range: This field is used along with the target field for TargetHttpProxy,
               TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
               TargetPool, TargetInstance.
               Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
               addressed to ports in the specified range will be forwarded to target.
               Forwarding rules with the same [IPAddress, IPProtocol] pair must have
               disjoint port ranges.
               Some types of forwarding target have constraints on the acceptable
               ports:
               * TargetHttpProxy: 80, 8080
               * TargetHttpsProxy: 443
               * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetVpnGateway: 500, 4500
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] target: The URL of the target resource to receive the matched traffic.
               The forwarded traffic must be of a type appropriate to the target object.
               For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
               are valid.
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

            __props__['description'] = description
            __props__['ip_address'] = ip_address
            __props__['ip_protocol'] = ip_protocol
            __props__['ip_version'] = ip_version
            __props__['labels'] = labels
            __props__['load_balancing_scheme'] = load_balancing_scheme
            __props__['metadata_filters'] = metadata_filters
            __props__['name'] = name
            __props__['network'] = network
            __props__['port_range'] = port_range
            __props__['project'] = project
            if target is None:
                raise TypeError("Missing required property 'target'")
            __props__['target'] = target
            __props__['label_fingerprint'] = None
            __props__['self_link'] = None
        super(GlobalForwardingRule, __self__).__init__(
            'gcp:compute/globalForwardingRule:GlobalForwardingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, description: Optional[pulumi.Input[str]] = None, ip_address: Optional[pulumi.Input[str]] = None, ip_protocol: Optional[pulumi.Input[str]] = None, ip_version: Optional[pulumi.Input[str]] = None, label_fingerprint: Optional[pulumi.Input[str]] = None, labels: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, load_balancing_scheme: Optional[pulumi.Input[str]] = None, metadata_filters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GlobalForwardingRuleMetadataFilterArgs']]]]] = None, name: Optional[pulumi.Input[str]] = None, network: Optional[pulumi.Input[str]] = None, port_range: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, self_link: Optional[pulumi.Input[str]] = None, target: Optional[pulumi.Input[str]] = None) -> 'GlobalForwardingRule':
        """
        Get an existing GlobalForwardingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] ip_address: The IP address that this forwarding rule is serving on behalf of.
               Addresses are restricted based on the forwarding rule's load balancing
               scheme (EXTERNAL or INTERNAL) and scope (global or regional).
               When the load balancing scheme is EXTERNAL, for global forwarding
               rules, the address must be a global IP, and for regional forwarding
               rules, the address must live in the same region as the forwarding
               rule. If this field is empty, an ephemeral IPv4 address from the same
               scope (global or regional) will be assigned. A regional forwarding
               rule supports IPv4 only. A global forwarding rule supports either IPv4
               or IPv6.
               When the load balancing scheme is INTERNAL, this can only be an RFC
               1918 IP address belonging to the network/subnet configured for the
               forwarding rule. By default, if this field is empty, an ephemeral
               internal IP address will be automatically allocated from the IP range
               of the subnet or network configured for this forwarding rule.
               An address must be specified by a literal IP address. > **NOTE**: While
               the API allows you to specify various resource paths for an address resource
               instead, this provider requires this to specifically be an IP address to
               avoid needing to fetching the IP address from resource paths on refresh
               or unnecessary diffs.
        :param pulumi.Input[str] ip_protocol: The IP protocol to which this rule applies. When the load balancing scheme is
               INTERNAL_SELF_MANAGED, only TCP is valid.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this global forwarding rule.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Labels to apply to this forwarding rule.  A list of key->value pairs.
        :param pulumi.Input[str] load_balancing_scheme: This signifies what the GlobalForwardingRule will be used for.
               The value of INTERNAL_SELF_MANAGED means that this will be used for
               Internal Global HTTP(S) LB. The value of EXTERNAL means that this
               will be used for External Global Load Balancing (HTTP(S) LB,
               External TCP/UDP LB, SSL Proxy)
               NOTE: Currently global forwarding rules cannot be used for INTERNAL
               load balancing.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GlobalForwardingRuleMetadataFilterArgs']]]] metadata_filters: Opaque filter criteria used by Loadbalancer to restrict routing
               configuration to a limited set xDS compliant clients. In their xDS
               requests to Loadbalancer, xDS clients present node metadata. If a
               match takes place, the relevant routing configuration is made available
               to those proxies.
               For each metadataFilter in this list, if its filterMatchCriteria is set
               to MATCH_ANY, at least one of the filterLabels must match the
               corresponding label provided in the metadata. If its filterMatchCriteria
               is set to MATCH_ALL, then all of its filterLabels must match with
               corresponding labels in the provided metadata.
               metadataFilters specified here can be overridden by those specified in
               the UrlMap that this ForwardingRule references.
               metadataFilters only applies to Loadbalancers that have their
               loadBalancingScheme set to INTERNAL_SELF_MANAGED.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the metadata label. The length must be between
               1 and 1024 characters, inclusive.
        :param pulumi.Input[str] network: This field is not used for external load balancing.
               For INTERNAL_SELF_MANAGED load balancing, this field
               identifies the network that the load balanced IP should belong to
               for this global forwarding rule. If this field is not specified,
               the default network will be used.
        :param pulumi.Input[str] port_range: This field is used along with the target field for TargetHttpProxy,
               TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
               TargetPool, TargetInstance.
               Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
               addressed to ports in the specified range will be forwarded to target.
               Forwarding rules with the same [IPAddress, IPProtocol] pair must have
               disjoint port ranges.
               Some types of forwarding target have constraints on the acceptable
               ports:
               * TargetHttpProxy: 80, 8080
               * TargetHttpsProxy: 443
               * TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
               1883, 5222
               * TargetVpnGateway: 500, 4500
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] target: The URL of the target resource to receive the matched traffic.
               The forwarded traffic must be of a type appropriate to the target object.
               For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
               are valid.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["ip_address"] = ip_address
        __props__["ip_protocol"] = ip_protocol
        __props__["ip_version"] = ip_version
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["load_balancing_scheme"] = load_balancing_scheme
        __props__["metadata_filters"] = metadata_filters
        __props__["name"] = name
        __props__["network"] = network
        __props__["port_range"] = port_range
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["target"] = target
        return GlobalForwardingRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

