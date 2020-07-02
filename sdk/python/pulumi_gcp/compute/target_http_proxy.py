# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TargetHttpProxy(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035. Specifically, the name must be 1-63 characters long and match
    the regular expression `a-z?` which means the
    first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last
    character, which cannot be a dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    proxy_id: pulumi.Output[float]
    """
    The unique identifier for the resource.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    url_map: pulumi.Output[str]
    """
    A reference to the UrlMap resource that defines the mapping from URL
    to the BackendService.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, project=None, url_map=None, __props__=None, __name__=None, __opts__=None):
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
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
        ```
        ### Target Http Proxy Https Redirect

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_url_map = gcp.compute.URLMap("defaultURLMap", default_url_redirect={
            "httpsRedirect": True,
            "stripQuery": False,
        })
        default_target_http_proxy = gcp.compute.TargetHttpProxy("defaultTargetHttpProxy", url_map=default_url_map.id)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['name'] = name
            __props__['project'] = project
            if url_map is None:
                raise TypeError("Missing required property 'url_map'")
            __props__['url_map'] = url_map
            __props__['creation_timestamp'] = None
            __props__['proxy_id'] = None
            __props__['self_link'] = None
        super(TargetHttpProxy, __self__).__init__(
            'gcp:compute/targetHttpProxy:TargetHttpProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, name=None, project=None, proxy_id=None, self_link=None, url_map=None):
        """
        Get an existing TargetHttpProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[float] proxy_id: The unique identifier for the resource.
        :param pulumi.Input[str] self_link: The URI of the created resource.
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
        __props__["self_link"] = self_link
        __props__["url_map"] = url_map
        return TargetHttpProxy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
