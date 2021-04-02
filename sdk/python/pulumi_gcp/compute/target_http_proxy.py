# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TargetHttpProxyArgs', 'TargetHttpProxy']

@pulumi.input_type
class TargetHttpProxyArgs:
    def __init__(__self__, *,
                 url_map: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy_bind: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TargetHttpProxy resource.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
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
        :param pulumi.Input[bool] proxy_bind: This field only applies when the forwarding rule that references
               this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        """
        pulumi.set(__self__, "url_map", url_map)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if proxy_bind is not None:
            pulumi.set(__self__, "proxy_bind", proxy_bind)

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> pulumi.Input[str]:
        """
        A reference to the UrlMap resource that defines the mapping from URL
        to the BackendService.
        """
        return pulumi.get(self, "url_map")

    @url_map.setter
    def url_map(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_map", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="proxyBind")
    def proxy_bind(self) -> Optional[pulumi.Input[bool]]:
        """
        This field only applies when the forwarding rule that references
        this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "proxy_bind")

    @proxy_bind.setter
    def proxy_bind(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "proxy_bind", value)


@pulumi.input_type
class _TargetHttpProxyState:
    def __init__(__self__, *,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy_bind: Optional[pulumi.Input[bool]] = None,
                 proxy_id: Optional[pulumi.Input[int]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TargetHttpProxy resources.
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
        :param pulumi.Input[bool] proxy_bind: This field only applies when the forwarding rule that references
               this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        :param pulumi.Input[int] proxy_id: The unique identifier for the resource.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
        """
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if proxy_bind is not None:
            pulumi.set(__self__, "proxy_bind", proxy_bind)
        if proxy_id is not None:
            pulumi.set(__self__, "proxy_id", proxy_id)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if url_map is not None:
            pulumi.set(__self__, "url_map", url_map)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="proxyBind")
    def proxy_bind(self) -> Optional[pulumi.Input[bool]]:
        """
        This field only applies when the forwarding rule that references
        this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "proxy_bind")

    @proxy_bind.setter
    def proxy_bind(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "proxy_bind", value)

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> Optional[pulumi.Input[int]]:
        """
        The unique identifier for the resource.
        """
        return pulumi.get(self, "proxy_id")

    @proxy_id.setter
    def proxy_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "proxy_id", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> Optional[pulumi.Input[str]]:
        """
        A reference to the UrlMap resource that defines the mapping from URL
        to the BackendService.
        """
        return pulumi.get(self, "url_map")

    @url_map.setter
    def url_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_map", value)


class TargetHttpProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy_bind: Optional[pulumi.Input[bool]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a TargetHttpProxy resource, which is used by one or more global
        forwarding rule to route incoming HTTP requests to a URL map.

        To get more information about TargetHttpProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)

        ## Example Usage
        ### Target Http Proxy Basic

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
            default_service=default_backend_service.id,
            host_rules=[gcp.compute.URLMapHostRuleArgs(
                hosts=["mysite.com"],
                path_matcher="allpaths",
            )],
            path_matchers=[gcp.compute.URLMapPathMatcherArgs(
                name="allpaths",
                default_service=default_backend_service.id,
                path_rules=[gcp.compute.URLMapPathMatcherPathRuleArgs(
                    paths=["/*"],
                    service=default_backend_service.id,
                )],
            )])
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
        ```
        ### Target Http Proxy Https Redirect

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_url_map = gcp.compute.URLMap("defaultURLMap", default_url_redirect=gcp.compute.URLMapDefaultUrlRedirectArgs(
            https_redirect=True,
            strip_query=False,
        ))
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
        ```

        ## Import

        TargetHttpProxy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default projects/{{project}}/global/targetHttpProxies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{name}}
        ```

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
        :param pulumi.Input[bool] proxy_bind: This field only applies when the forwarding rule that references
               this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TargetHttpProxyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a TargetHttpProxy resource, which is used by one or more global
        forwarding rule to route incoming HTTP requests to a URL map.

        To get more information about TargetHttpProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)

        ## Example Usage
        ### Target Http Proxy Basic

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
            default_service=default_backend_service.id,
            host_rules=[gcp.compute.URLMapHostRuleArgs(
                hosts=["mysite.com"],
                path_matcher="allpaths",
            )],
            path_matchers=[gcp.compute.URLMapPathMatcherArgs(
                name="allpaths",
                default_service=default_backend_service.id,
                path_rules=[gcp.compute.URLMapPathMatcherPathRuleArgs(
                    paths=["/*"],
                    service=default_backend_service.id,
                )],
            )])
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
        ```
        ### Target Http Proxy Https Redirect

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_url_map = gcp.compute.URLMap("defaultURLMap", default_url_redirect=gcp.compute.URLMapDefaultUrlRedirectArgs(
            https_redirect=True,
            strip_query=False,
        ))
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
        ```

        ## Import

        TargetHttpProxy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default projects/{{project}}/global/targetHttpProxies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetHttpProxy:TargetHttpProxy default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param TargetHttpProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetHttpProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 proxy_bind: Optional[pulumi.Input[bool]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = TargetHttpProxyArgs.__new__(TargetHttpProxyArgs)

            __props__.__dict__['description'] = description
            __props__.__dict__['name'] = name
            __props__.__dict__['project'] = project
            __props__.__dict__['proxy_bind'] = proxy_bind
            if url_map is None and not opts.urn:
                raise TypeError("Missing required property 'url_map'")
            __props__.__dict__['url_map'] = url_map
            __props__.__dict__['creation_timestamp'] = None
            __props__.__dict__['proxy_id'] = None
            __props__.__dict__['self_link'] = None
        super(TargetHttpProxy, __self__).__init__(
            'gcp:compute/targetHttpProxy:TargetHttpProxy',
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
            proxy_bind: Optional[pulumi.Input[bool]] = None,
            proxy_id: Optional[pulumi.Input[int]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            url_map: Optional[pulumi.Input[str]] = None) -> 'TargetHttpProxy':
        """
        Get an existing TargetHttpProxy resource's state with the given name, id, and optional extra
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
        :param pulumi.Input[bool] proxy_bind: This field only applies when the forwarding rule that references
               this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        :param pulumi.Input[int] proxy_id: The unique identifier for the resource.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] url_map: A reference to the UrlMap resource that defines the mapping from URL
               to the BackendService.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TargetHttpProxyState.__new__(_TargetHttpProxyState)

        __props__.__dict__['creation_timestamp'] = creation_timestamp
        __props__.__dict__['description'] = description
        __props__.__dict__['name'] = name
        __props__.__dict__['project'] = project
        __props__.__dict__['proxy_bind'] = proxy_bind
        __props__.__dict__['proxy_id'] = proxy_id
        __props__.__dict__['self_link'] = self_link
        __props__.__dict__['url_map'] = url_map
        return TargetHttpProxy(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="proxyBind")
    def proxy_bind(self) -> pulumi.Output[bool]:
        """
        This field only applies when the forwarding rule that references
        this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        """
        return pulumi.get(self, "proxy_bind")

    @property
    @pulumi.getter(name="proxyId")
    def proxy_id(self) -> pulumi.Output[int]:
        """
        The unique identifier for the resource.
        """
        return pulumi.get(self, "proxy_id")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> pulumi.Output[str]:
        """
        A reference to the UrlMap resource that defines the mapping from URL
        to the BackendService.
        """
        return pulumi.get(self, "url_map")

