# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['TargetHttpsProxy']


class TargetHttpsProxy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 quic_override: Optional[pulumi.Input[str]] = None,
                 ssl_certificates: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 ssl_policy: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a TargetHttpsProxy resource, which is used by one or more
        global forwarding rule to route incoming HTTPS requests to a URL map.

        To get more information about TargetHttpsProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] quic_override: Specifies the QUIC override policy for this resource. This determines
               whether the load balancer will attempt to negotiate QUIC with clients
               or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
               specified, uses the QUIC policy with no user overrides, which is
               equivalent to DISABLE.
               Default value is `NONE`.
               Possible values are `NONE`, `ENABLE`, and `DISABLE`.
        :param pulumi.Input[List[pulumi.Input[str]]] ssl_certificates: A list of SslCertificate resources that are used to authenticate
               connections between users and the load balancer. At least one SSL
               certificate must be specified.
        :param pulumi.Input[str] ssl_policy: A reference to the SslPolicy resource that will be associated with
               the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
               resource will not have any SSL policy configured.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
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
            __props__['name'] = name
            __props__['project'] = project
            __props__['quic_override'] = quic_override
            if ssl_certificates is None:
                raise TypeError("Missing required property 'ssl_certificates'")
            __props__['ssl_certificates'] = ssl_certificates
            __props__['ssl_policy'] = ssl_policy
            if url_map is None:
                raise TypeError("Missing required property 'url_map'")
            __props__['url_map'] = url_map
            __props__['creation_timestamp'] = None
            __props__['proxy_id'] = None
            __props__['self_link'] = None
        super(TargetHttpsProxy, __self__).__init__(
            'gcp:compute/targetHttpsProxy:TargetHttpsProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            proxy_id: Optional[pulumi.Input[int]] = None,
            quic_override: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            ssl_certificates: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            ssl_policy: Optional[pulumi.Input[str]] = None,
            url_map: Optional[pulumi.Input[str]] = None) -> 'TargetHttpsProxy':
        """
        Get an existing TargetHttpsProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] proxy_id: The unique identifier for the resource.
        :param pulumi.Input[str] quic_override: Specifies the QUIC override policy for this resource. This determines
               whether the load balancer will attempt to negotiate QUIC with clients
               or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
               specified, uses the QUIC policy with no user overrides, which is
               equivalent to DISABLE.
               Default value is `NONE`.
               Possible values are `NONE`, `ENABLE`, and `DISABLE`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[List[pulumi.Input[str]]] ssl_certificates: A list of SslCertificate resources that are used to authenticate
               connections between users and the load balancer. At least one SSL
               certificate must be specified.
        :param pulumi.Input[str] ssl_policy: A reference to the SslPolicy resource that will be associated with
               the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
               resource will not have any SSL policy configured.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["proxy_id"] = proxy_id
        __props__["quic_override"] = quic_override
        __props__["self_link"] = self_link
        __props__["ssl_certificates"] = ssl_certificates
        __props__["ssl_policy"] = ssl_policy
        __props__["url_map"] = url_map
        return TargetHttpsProxy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> pulumi.Output[int]:
        """
        The unique identifier for the resource.
        """
        return pulumi.get(self, "proxy_id")

    @property
    @pulumi.getter(name="quicOverride")
    def quic_override(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the QUIC override policy for this resource. This determines
        whether the load balancer will attempt to negotiate QUIC with clients
        or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
        specified, uses the QUIC policy with no user overrides, which is
        equivalent to DISABLE.
        Default value is `NONE`.
        Possible values are `NONE`, `ENABLE`, and `DISABLE`.
        """
        return pulumi.get(self, "quic_override")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> pulumi.Output[List[str]]:
        """
        A list of SslCertificate resources that are used to authenticate
        connections between users and the load balancer. At least one SSL
        certificate must be specified.
        """
        return pulumi.get(self, "ssl_certificates")

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> pulumi.Output[Optional[str]]:
        """
        A reference to the SslPolicy resource that will be associated with
        the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
        resource will not have any SSL policy configured.
        """
        return pulumi.get(self, "ssl_policy")

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> pulumi.Output[str]:
        """
        A reference to the UrlMap resource that defines the mapping from URL
        to the BackendService.
        """
        return pulumi.get(self, "url_map")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

