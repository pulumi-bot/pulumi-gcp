# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RegionTargetHttpsProxy(pulumi.CustomResource):
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
    region: pulumi.Output[str]
    """
    The Region in which the created target https proxy should reside.
    If it is not provided, the provider region is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    ssl_certificates: pulumi.Output[list]
    """
    A list of RegionSslCertificate resources that are used to authenticate
    connections between users and the load balancer. Currently, exactly
    one SSL certificate must be specified.
    """
    url_map: pulumi.Output[str]
    """
    A reference to the RegionUrlMap resource that defines the mapping from URL
    to the RegionBackendService.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, project=None, region=None, ssl_certificates=None, url_map=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents a RegionTargetHttpsProxy resource, which is used by one or more
        forwarding rules to route incoming HTTPS requests to a URL map.

        To get more information about RegionTargetHttpsProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/regionTargetHttpsProxies)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)

        ## Example Usage

        ### Region Target Https Proxy Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_region_ssl_certificate = gcp.compute.RegionSslCertificate("defaultRegionSslCertificate",
            region="us-central1",
            private_key=(lambda path: open(path).read())("path/to/private.key"),
            certificate=(lambda path: open(path).read())("path/to/certificate.crt"))
        default_region_health_check = gcp.compute.RegionHealthCheck("defaultRegionHealthCheck",
            region="us-central1",
            http_health_check={
                "port": 80,
            })
        default_region_backend_service = gcp.compute.RegionBackendService("defaultRegionBackendService",
            region="us-central1",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_region_health_check.id])
        default_region_url_map = gcp.compute.RegionUrlMap("defaultRegionUrlMap",
            region="us-central1",
            description="a description",
            default_service=default_region_backend_service.id,
            host_rule=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matcher=[{
                "name": "allpaths",
                "default_service": default_region_backend_service.id,
                "path_rule": [{
                    "paths": ["/*"],
                    "service": default_region_backend_service.id,
                }],
            }])
        default_region_target_https_proxy = gcp.compute.RegionTargetHttpsProxy("defaultRegionTargetHttpsProxy",
            region="us-central1",
            url_map=default_region_url_map.id,
            ssl_certificates=[default_region_ssl_certificate.id])
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
        :param pulumi.Input[str] region: The Region in which the created target https proxy should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[list] ssl_certificates: A list of RegionSslCertificate resources that are used to authenticate
               connections between users and the load balancer. Currently, exactly
               one SSL certificate must be specified.
        :param pulumi.Input[str] url_map: A reference to the RegionUrlMap resource that defines the mapping from URL
               to the RegionBackendService.
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
            __props__['region'] = region
            if ssl_certificates is None:
                raise TypeError("Missing required property 'ssl_certificates'")
            __props__['ssl_certificates'] = ssl_certificates
            if url_map is None:
                raise TypeError("Missing required property 'url_map'")
            __props__['url_map'] = url_map
            __props__['creation_timestamp'] = None
            __props__['proxy_id'] = None
            __props__['self_link'] = None
        super(RegionTargetHttpsProxy, __self__).__init__(
            'gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, creation_timestamp=None, description=None, name=None, project=None, proxy_id=None, region=None, self_link=None, ssl_certificates=None, url_map=None):
        """
        Get an existing RegionTargetHttpsProxy resource's state with the given name, id, and optional extra
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
        :param pulumi.Input[str] region: The Region in which the created target https proxy should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[list] ssl_certificates: A list of RegionSslCertificate resources that are used to authenticate
               connections between users and the load balancer. Currently, exactly
               one SSL certificate must be specified.
        :param pulumi.Input[str] url_map: A reference to the RegionUrlMap resource that defines the mapping from URL
               to the RegionBackendService.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["name"] = name
        __props__["project"] = project
        __props__["proxy_id"] = proxy_id
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["ssl_certificates"] = ssl_certificates
        __props__["url_map"] = url_map
        return RegionTargetHttpsProxy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

