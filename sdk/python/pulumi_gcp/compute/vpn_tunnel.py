# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VPNTunnel(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    detailed_status: pulumi.Output[str]
    """
    Detailed status message for the VPN tunnel.
    """
    ike_version: pulumi.Output[float]
    """
    IKE protocol version to use when establishing the VPN tunnel with
    peer VPN gateway.
    Acceptable IKE versions are 1 or 2. Default version is 2.
    """
    label_fingerprint: pulumi.Output[str]
    """
    The fingerprint used for optimistic locking of this resource. Used internally during updates.
    """
    labels: pulumi.Output[dict]
    """
    Labels to apply to this VpnTunnel.
    """
    local_traffic_selectors: pulumi.Output[list]
    """
    Local traffic selector to use when establishing the VPN tunnel with
    peer VPN gateway. The value should be a CIDR formatted string,
    for example `192.168.0.0/16`. The ranges should be disjoint.
    Only IPv4 is supported.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. The name must be 1-63 characters long, and
    comply with RFC1035. Specifically, the name must be 1-63
    characters long and match the regular expression
    `a-z?` which means the first character
    must be a lowercase letter, and all following characters must
    be a dash, lowercase letter, or digit,
    except the last character, which cannot be a dash.
    """
    peer_external_gateway: pulumi.Output[str]
    """
    URL of the peer side external VPN gateway to which this VPN tunnel is connected.
    """
    peer_external_gateway_interface: pulumi.Output[float]
    """
    The interface ID of the external VPN gateway to which this VPN tunnel is connected.
    """
    peer_gcp_gateway: pulumi.Output[str]
    """
    URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
    If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
    ID in the peer GCP VPN gateway.
    This field must reference a `compute.HaVpnGateway` resource.
    """
    peer_ip: pulumi.Output[str]
    """
    IP address of the peer VPN gateway. Only IPv4 is supported.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
    """
    remote_traffic_selectors: pulumi.Output[list]
    """
    Remote traffic selector to use when establishing the VPN tunnel with
    peer VPN gateway. The value should be a CIDR formatted string,
    for example `192.168.0.0/16`. The ranges should be disjoint.
    Only IPv4 is supported.
    """
    router: pulumi.Output[str]
    """
    URL of router resource to be used for dynamic routing.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    shared_secret: pulumi.Output[str]
    """
    Shared secret used to set the secure session between the Cloud VPN
    gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
    """
    shared_secret_hash: pulumi.Output[str]
    """
    Hash of the shared secret.
    """
    target_vpn_gateway: pulumi.Output[str]
    """
    URL of the Target VPN gateway with which this VPN tunnel is
    associated.
    """
    tunnel_id: pulumi.Output[str]
    """
    The unique identifier for the resource. This identifier is defined by the server.
    """
    vpn_gateway: pulumi.Output[str]
    """
    URL of the VPN gateway with which this VPN tunnel is associated.
    This must be used if a High Availability VPN gateway resource is created.
    This field must reference a `compute.HaVpnGateway` resource.
    """
    vpn_gateway_interface: pulumi.Output[float]
    """
    The interface ID of the VPN gateway with which this VPN tunnel is associated.
    """
    def __init__(__self__, resource_name, opts=None, description=None, ike_version=None, labels=None, local_traffic_selectors=None, name=None, peer_external_gateway=None, peer_external_gateway_interface=None, peer_gcp_gateway=None, peer_ip=None, project=None, region=None, remote_traffic_selectors=None, router=None, shared_secret=None, target_vpn_gateway=None, vpn_gateway=None, vpn_gateway_interface=None, __props__=None, __name__=None, __opts__=None):
        """
        VPN tunnel resource.


        To get more information about VpnTunnel, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
        * How-to Guides
            * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
            * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)

        > **Warning:** All arguments including `shared_secret` will be stored in the raw
        state as plain-text.

        ## Example Usage

        ### Vpn Tunnel Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network1 = gcp.compute.Network("network1")
        target_gateway = gcp.compute.VPNGateway("targetGateway", network=network1.id)
        vpn_static_ip = gcp.compute.Address("vpnStaticIp")
        fr_esp = gcp.compute.ForwardingRule("frEsp",
            ip_protocol="ESP",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        fr_udp500 = gcp.compute.ForwardingRule("frUdp500",
            ip_protocol="UDP",
            port_range="500",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        fr_udp4500 = gcp.compute.ForwardingRule("frUdp4500",
            ip_protocol="UDP",
            port_range="4500",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        tunnel1 = gcp.compute.VPNTunnel("tunnel1",
            peer_ip="15.0.0.120",
            shared_secret="a secret message",
            target_vpn_gateway=target_gateway.id)
        route1 = gcp.compute.Route("route1",
            network=network1.name,
            dest_range="15.0.0.0/24",
            priority=1000,
            next_hop_vpn_tunnel=tunnel1.id)
        ```

        ### Vpn Tunnel Beta

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network1 = gcp.compute.Network("network1")
        target_gateway = gcp.compute.VPNGateway("targetGateway", network=network1.id)
        vpn_static_ip = gcp.compute.Address("vpnStaticIp")
        fr_esp = gcp.compute.ForwardingRule("frEsp",
            ip_protocol="ESP",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        fr_udp500 = gcp.compute.ForwardingRule("frUdp500",
            ip_protocol="UDP",
            port_range="500",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        fr_udp4500 = gcp.compute.ForwardingRule("frUdp4500",
            ip_protocol="UDP",
            port_range="4500",
            ip_address=vpn_static_ip.address,
            target=target_gateway.id)
        tunnel1 = gcp.compute.VPNTunnel("tunnel1",
            peer_ip="15.0.0.120",
            shared_secret="a secret message",
            target_vpn_gateway=target_gateway.id,
            labels={
                "foo": "bar",
            })
        route1 = gcp.compute.Route("route1",
            network=network1.name,
            dest_range="15.0.0.0/24",
            priority=1000,
            next_hop_vpn_tunnel=tunnel1.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[float] ike_version: IKE protocol version to use when establishing the VPN tunnel with
               peer VPN gateway.
               Acceptable IKE versions are 1 or 2. Default version is 2.
        :param pulumi.Input[dict] labels: Labels to apply to this VpnTunnel.
        :param pulumi.Input[list] local_traffic_selectors: Local traffic selector to use when establishing the VPN tunnel with
               peer VPN gateway. The value should be a CIDR formatted string,
               for example `192.168.0.0/16`. The ranges should be disjoint.
               Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63
               characters long and match the regular expression
               `a-z?` which means the first character
               must be a lowercase letter, and all following characters must
               be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] peer_external_gateway: URL of the peer side external VPN gateway to which this VPN tunnel is connected.
        :param pulumi.Input[float] peer_external_gateway_interface: The interface ID of the external VPN gateway to which this VPN tunnel is connected.
        :param pulumi.Input[str] peer_gcp_gateway: URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
               If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
               ID in the peer GCP VPN gateway.
               This field must reference a `compute.HaVpnGateway` resource.
        :param pulumi.Input[str] peer_ip: IP address of the peer VPN gateway. Only IPv4 is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
        :param pulumi.Input[list] remote_traffic_selectors: Remote traffic selector to use when establishing the VPN tunnel with
               peer VPN gateway. The value should be a CIDR formatted string,
               for example `192.168.0.0/16`. The ranges should be disjoint.
               Only IPv4 is supported.
        :param pulumi.Input[str] router: URL of router resource to be used for dynamic routing.
        :param pulumi.Input[str] shared_secret: Shared secret used to set the secure session between the Cloud VPN
               gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] target_vpn_gateway: URL of the Target VPN gateway with which this VPN tunnel is
               associated.
        :param pulumi.Input[str] vpn_gateway: URL of the VPN gateway with which this VPN tunnel is associated.
               This must be used if a High Availability VPN gateway resource is created.
               This field must reference a `compute.HaVpnGateway` resource.
        :param pulumi.Input[float] vpn_gateway_interface: The interface ID of the VPN gateway with which this VPN tunnel is associated.
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
            __props__['ike_version'] = ike_version
            __props__['labels'] = labels
            __props__['local_traffic_selectors'] = local_traffic_selectors
            __props__['name'] = name
            __props__['peer_external_gateway'] = peer_external_gateway
            __props__['peer_external_gateway_interface'] = peer_external_gateway_interface
            __props__['peer_gcp_gateway'] = peer_gcp_gateway
            __props__['peer_ip'] = peer_ip
            __props__['project'] = project
            __props__['region'] = region
            __props__['remote_traffic_selectors'] = remote_traffic_selectors
            __props__['router'] = router
            if shared_secret is None:
                raise TypeError("Missing required property 'shared_secret'")
            __props__['shared_secret'] = shared_secret
            __props__['target_vpn_gateway'] = target_vpn_gateway
            __props__['vpn_gateway'] = vpn_gateway
            __props__['vpn_gateway_interface'] = vpn_gateway_interface
            __props__['creation_timestamp'] = None
            __props__['detailed_status'] = None
            __props__['label_fingerprint'] = None
            __props__['self_link'] = None
            __props__['shared_secret_hash'] = None
            __props__['tunnel_id'] = None
        super(VPNTunnel, __self__).__init__(
            'gcp:compute/vPNTunnel:VPNTunnel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, detailed_status=None, ike_version=None, label_fingerprint=None, labels=None, local_traffic_selectors=None, name=None, peer_external_gateway=None, peer_external_gateway_interface=None, peer_gcp_gateway=None, peer_ip=None, project=None, region=None, remote_traffic_selectors=None, router=None, self_link=None, shared_secret=None, shared_secret_hash=None, target_vpn_gateway=None, tunnel_id=None, vpn_gateway=None, vpn_gateway_interface=None):
        """
        Get an existing VPNTunnel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] detailed_status: Detailed status message for the VPN tunnel.
        :param pulumi.Input[float] ike_version: IKE protocol version to use when establishing the VPN tunnel with
               peer VPN gateway.
               Acceptable IKE versions are 1 or 2. Default version is 2.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[dict] labels: Labels to apply to this VpnTunnel.
        :param pulumi.Input[list] local_traffic_selectors: Local traffic selector to use when establishing the VPN tunnel with
               peer VPN gateway. The value should be a CIDR formatted string,
               for example `192.168.0.0/16`. The ranges should be disjoint.
               Only IPv4 is supported.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63
               characters long and match the regular expression
               `a-z?` which means the first character
               must be a lowercase letter, and all following characters must
               be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] peer_external_gateway: URL of the peer side external VPN gateway to which this VPN tunnel is connected.
        :param pulumi.Input[float] peer_external_gateway_interface: The interface ID of the external VPN gateway to which this VPN tunnel is connected.
        :param pulumi.Input[str] peer_gcp_gateway: URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
               If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
               ID in the peer GCP VPN gateway.
               This field must reference a `compute.HaVpnGateway` resource.
        :param pulumi.Input[str] peer_ip: IP address of the peer VPN gateway. Only IPv4 is supported.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
        :param pulumi.Input[list] remote_traffic_selectors: Remote traffic selector to use when establishing the VPN tunnel with
               peer VPN gateway. The value should be a CIDR formatted string,
               for example `192.168.0.0/16`. The ranges should be disjoint.
               Only IPv4 is supported.
        :param pulumi.Input[str] router: URL of router resource to be used for dynamic routing.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] shared_secret: Shared secret used to set the secure session between the Cloud VPN
               gateway and the peer VPN gateway.  **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] shared_secret_hash: Hash of the shared secret.
        :param pulumi.Input[str] target_vpn_gateway: URL of the Target VPN gateway with which this VPN tunnel is
               associated.
        :param pulumi.Input[str] tunnel_id: The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] vpn_gateway: URL of the VPN gateway with which this VPN tunnel is associated.
               This must be used if a High Availability VPN gateway resource is created.
               This field must reference a `compute.HaVpnGateway` resource.
        :param pulumi.Input[float] vpn_gateway_interface: The interface ID of the VPN gateway with which this VPN tunnel is associated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["detailed_status"] = detailed_status
        __props__["ike_version"] = ike_version
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["local_traffic_selectors"] = local_traffic_selectors
        __props__["name"] = name
        __props__["peer_external_gateway"] = peer_external_gateway
        __props__["peer_external_gateway_interface"] = peer_external_gateway_interface
        __props__["peer_gcp_gateway"] = peer_gcp_gateway
        __props__["peer_ip"] = peer_ip
        __props__["project"] = project
        __props__["region"] = region
        __props__["remote_traffic_selectors"] = remote_traffic_selectors
        __props__["router"] = router
        __props__["self_link"] = self_link
        __props__["shared_secret"] = shared_secret
        __props__["shared_secret_hash"] = shared_secret_hash
        __props__["target_vpn_gateway"] = target_vpn_gateway
        __props__["tunnel_id"] = tunnel_id
        __props__["vpn_gateway"] = vpn_gateway
        __props__["vpn_gateway_interface"] = vpn_gateway_interface
        return VPNTunnel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

