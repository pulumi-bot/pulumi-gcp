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

__all__ = ['URLMap']


class URLMap(pulumi.CustomResource):
    creation_timestamp: pulumi.Output[str] = pulumi.property("creationTimestamp")
    """
    Creation timestamp in RFC3339 text format.
    """

    default_route_action: pulumi.Output[Optional['outputs.URLMapDefaultRouteAction']] = pulumi.property("defaultRouteAction")
    """
    defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
    advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
    to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
    Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
    Only one of defaultRouteAction or defaultUrlRedirect must be set.  Structure is documented below.
    """

    default_service: pulumi.Output[Optional[str]] = pulumi.property("defaultService")
    """
    The backend service or backend bucket to use when none of the given paths match.
    """

    default_url_redirect: pulumi.Output[Optional['outputs.URLMapDefaultUrlRedirect']] = pulumi.property("defaultUrlRedirect")
    """
    When none of the specified hostRules match, the request is redirected to a URL specified
    by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
    defaultRouteAction must not be set.  Structure is documented below.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    Description of this test case.
    """

    fingerprint: pulumi.Output[str] = pulumi.property("fingerprint")
    """
    Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
    """

    header_action: pulumi.Output[Optional['outputs.URLMapHeaderAction']] = pulumi.property("headerAction")
    """
    Specifies changes to request and response headers that need to take effect for
    the selected backendService.
    headerAction specified here take effect before headerAction in the enclosing
    HttpRouteRule, PathMatcher and UrlMap.  Structure is documented below.
    """

    host_rules: pulumi.Output[Optional[List['outputs.URLMapHostRule']]] = pulumi.property("hostRules")
    """
    The list of HostRules to use against the URL.  Structure is documented below.
    """

    map_id: pulumi.Output[float] = pulumi.property("mapId")
    """
    The unique identifier for the resource.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the query parameter to match. The query parameter must exist in the
    request, in the absence of which the request match fails.
    """

    path_matchers: pulumi.Output[Optional[List['outputs.URLMapPathMatcher']]] = pulumi.property("pathMatchers")
    """
    The name of the PathMatcher to use to match the path portion of the URL if the
    hostRule matches the URL's host portion.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """

    self_link: pulumi.Output[str] = pulumi.property("selfLink")
    """
    The URI of the created resource.
    """

    tests: pulumi.Output[Optional[List['outputs.URLMapTest']]] = pulumi.property("tests")
    """
    The list of expected URL mapping tests. Request to update this UrlMap will
    succeed only if all of the test cases pass. You can specify a maximum of 100
    tests per UrlMap.  Structure is documented below.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_route_action: Optional[pulumi.Input[pulumi.InputType['URLMapDefaultRouteActionArgs']]] = None,
                 default_service: Optional[pulumi.Input[str]] = None,
                 default_url_redirect: Optional[pulumi.Input[pulumi.InputType['URLMapDefaultUrlRedirectArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 header_action: Optional[pulumi.Input[pulumi.InputType['URLMapHeaderActionArgs']]] = None,
                 host_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapHostRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path_matchers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapPathMatcherArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tests: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapTestArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        UrlMaps are used to route requests to a backend service based on rules
        that you define for the host and path of an incoming URL.

        To get more information about UrlMap, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps)

        ## Example Usage
        ### Url Map Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HttpHealthCheck("default",
            request_path="/",
            check_interval_sec=1,
            timeout_sec=1)
        login = gcp.compute.BackendService("login",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id])
        home = gcp.compute.BackendService("home",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id])
        static_bucket = gcp.storage.Bucket("staticBucket", location="US")
        static_backend_bucket = gcp.compute.BackendBucket("staticBackendBucket",
            bucket_name=static_bucket.name,
            enable_cdn=True)
        urlmap = gcp.compute.URLMap("urlmap",
            description="a description",
            default_service=home.id,
            host_rules=[
                {
                    "hosts": ["mysite.com"],
                    "pathMatcher": "mysite",
                },
                {
                    "hosts": ["myothersite.com"],
                    "pathMatcher": "otherpaths",
                },
            ],
            path_matchers=[
                {
                    "name": "mysite",
                    "default_service": home.id,
                    "pathRules": [
                        {
                            "paths": ["/home"],
                            "service": home.id,
                        },
                        {
                            "paths": ["/login"],
                            "service": login.id,
                        },
                        {
                            "paths": ["/static"],
                            "service": static_backend_bucket.id,
                        },
                    ],
                },
                {
                    "name": "otherpaths",
                    "default_service": home.id,
                },
            ],
            tests=[{
                "service": home.id,
                "host": "hi.com",
                "path": "/home",
            }])
        ```
        ### Url Map Traffic Director Route

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HealthCheck("default", http_health_check={
            "port": 80,
        })
        home = gcp.compute.BackendService("home",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id],
            load_balancing_scheme="INTERNAL_SELF_MANAGED")
        urlmap = gcp.compute.URLMap("urlmap",
            description="a description",
            default_service=home.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": home.id,
                "routeRules": [{
                    "priority": 1,
                    "header_action": {
                        "requestHeadersToRemoves": ["RemoveMe2"],
                        "requestHeadersToAdds": [{
                            "headerName": "AddSomethingElse",
                            "headerValue": "MyOtherValue",
                            "replace": True,
                        }],
                        "responseHeadersToRemoves": ["RemoveMe3"],
                        "responseHeadersToAdds": [{
                            "headerName": "AddMe",
                            "headerValue": "MyValue",
                            "replace": False,
                        }],
                    },
                    "matchRules": [{
                        "fullPathMatch": "a full path",
                        "headerMatches": [{
                            "headerName": "someheader",
                            "exactMatch": "match this exactly",
                            "invertMatch": True,
                        }],
                        "ignoreCase": True,
                        "metadata_filters": [{
                            "filterMatchCriteria": "MATCH_ANY",
                            "filterLabels": [{
                                "name": "PLANET",
                                "value": "MARS",
                            }],
                        }],
                        "queryParameterMatches": [{
                            "name": "a query parameter",
                            "presentMatch": True,
                        }],
                    }],
                    "urlRedirect": {
                        "hostRedirect": "A host",
                        "httpsRedirect": False,
                        "pathRedirect": "some/path",
                        "redirectResponseCode": "TEMPORARY_REDIRECT",
                        "stripQuery": True,
                    },
                }],
            }],
            tests=[{
                "service": home.id,
                "host": "hi.com",
                "path": "/home",
            }])
        ```
        ### Url Map Traffic Director Route Partial

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HealthCheck("default", http_health_check={
            "port": 80,
        })
        home = gcp.compute.BackendService("home",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id],
            load_balancing_scheme="INTERNAL_SELF_MANAGED")
        urlmap = gcp.compute.URLMap("urlmap",
            description="a description",
            default_service=home.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": home.id,
                "routeRules": [{
                    "priority": 1,
                    "matchRules": [{
                        "prefixMatch": "/someprefix",
                        "headerMatches": [{
                            "headerName": "someheader",
                            "exactMatch": "match this exactly",
                            "invertMatch": True,
                        }],
                    }],
                    "urlRedirect": {
                        "pathRedirect": "some/path",
                        "redirectResponseCode": "TEMPORARY_REDIRECT",
                    },
                }],
            }],
            tests=[{
                "service": home.id,
                "host": "hi.com",
                "path": "/home",
            }])
        ```
        ### Url Map Traffic Director Path

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HealthCheck("default", http_health_check={
            "port": 80,
        })
        home = gcp.compute.BackendService("home",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id],
            load_balancing_scheme="INTERNAL_SELF_MANAGED")
        urlmap = gcp.compute.URLMap("urlmap",
            description="a description",
            default_service=home.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": home.id,
                "pathRules": [{
                    "paths": ["/home"],
                    "routeAction": {
                        "corsPolicy": {
                            "allowCredentials": True,
                            "allowHeaders": ["Allowed content"],
                            "allowMethods": ["GET"],
                            "allowOriginRegexes": ["abc.*"],
                            "allowOrigins": ["Allowed origin"],
                            "exposeHeaders": ["Exposed header"],
                            "maxAge": 30,
                            "disabled": False,
                        },
                        "faultInjectionPolicy": {
                            "abort": {
                                "httpStatus": 234,
                                "percentage": 5.6,
                            },
                            "delay": {
                                "fixedDelay": {
                                    "seconds": "0",
                                    "nanos": 50000,
                                },
                                "percentage": 7.8,
                            },
                        },
                        "requestMirrorPolicy": {
                            "backend_service": home.id,
                        },
                        "retryPolicy": {
                            "numRetries": 4,
                            "perTryTimeout": {
                                "seconds": "30",
                            },
                            "retryConditions": [
                                "5xx",
                                "deadline-exceeded",
                            ],
                        },
                        "timeout": {
                            "seconds": "20",
                            "nanos": 750000000,
                        },
                        "urlRewrite": {
                            "hostRewrite": "A replacement header",
                            "pathPrefixRewrite": "A replacement path",
                        },
                        "weightedBackendServices": [{
                            "backend_service": home.id,
                            "weight": 400,
                            "header_action": {
                                "requestHeadersToRemoves": ["RemoveMe"],
                                "requestHeadersToAdds": [{
                                    "headerName": "AddMe",
                                    "headerValue": "MyValue",
                                    "replace": True,
                                }],
                                "responseHeadersToRemoves": ["RemoveMe"],
                                "responseHeadersToAdds": [{
                                    "headerName": "AddMe",
                                    "headerValue": "MyValue",
                                    "replace": False,
                                }],
                            },
                        }],
                    },
                }],
            }],
            tests=[{
                "service": home.id,
                "host": "hi.com",
                "path": "/home",
            }])
        ```
        ### Url Map Traffic Director Path Partial

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HealthCheck("default", http_health_check={
            "port": 80,
        })
        home = gcp.compute.BackendService("home",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default.id],
            load_balancing_scheme="INTERNAL_SELF_MANAGED")
        urlmap = gcp.compute.URLMap("urlmap",
            description="a description",
            default_service=home.id,
            host_rules=[{
                "hosts": ["mysite.com"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": home.id,
                "pathRules": [{
                    "paths": ["/home"],
                    "routeAction": {
                        "corsPolicy": {
                            "allowCredentials": True,
                            "allowHeaders": ["Allowed content"],
                            "allowMethods": ["GET"],
                            "allowOriginRegexes": ["abc.*"],
                            "allowOrigins": ["Allowed origin"],
                            "exposeHeaders": ["Exposed header"],
                            "maxAge": 30,
                            "disabled": False,
                        },
                        "weightedBackendServices": [{
                            "backend_service": home.id,
                            "weight": 400,
                            "header_action": {
                                "requestHeadersToRemoves": ["RemoveMe"],
                                "requestHeadersToAdds": [{
                                    "headerName": "AddMe",
                                    "headerValue": "MyValue",
                                    "replace": True,
                                }],
                                "responseHeadersToRemoves": ["RemoveMe"],
                                "responseHeadersToAdds": [{
                                    "headerName": "AddMe",
                                    "headerValue": "MyValue",
                                    "replace": False,
                                }],
                            },
                        }],
                    },
                }],
            }],
            tests=[{
                "service": home.id,
                "host": "hi.com",
                "path": "/home",
            }])
        ```
        ### Url Map Header Based Routing

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
        service_a = gcp.compute.BackendService("service-a",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id])
        service_b = gcp.compute.BackendService("service-b",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id])
        urlmap = gcp.compute.URLMap("urlmap",
            description="header-based routing example",
            default_service=default_backend_service.id,
            host_rules=[{
                "hosts": ["*"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": default_backend_service.id,
                "routeRules": [
                    {
                        "priority": 1,
                        "service": service_a.id,
                        "matchRules": [{
                            "prefixMatch": "/",
                            "ignoreCase": True,
                            "headerMatches": [{
                                "headerName": "abtest",
                                "exactMatch": "a",
                            }],
                        }],
                    },
                    {
                        "priority": 2,
                        "service": service_b.id,
                        "matchRules": [{
                            "ignoreCase": True,
                            "prefixMatch": "/",
                            "headerMatches": [{
                                "headerName": "abtest",
                                "exactMatch": "b",
                            }],
                        }],
                    },
                ],
            }])
        ```
        ### Url Map Parameter Based Routing

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
        service_a = gcp.compute.BackendService("service-a",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id])
        service_b = gcp.compute.BackendService("service-b",
            port_name="http",
            protocol="HTTP",
            timeout_sec=10,
            health_checks=[default_http_health_check.id])
        urlmap = gcp.compute.URLMap("urlmap",
            description="parameter-based routing example",
            default_service=default_backend_service.id,
            host_rules=[{
                "hosts": ["*"],
                "pathMatcher": "allpaths",
            }],
            path_matchers=[{
                "name": "allpaths",
                "default_service": default_backend_service.id,
                "routeRules": [
                    {
                        "priority": 1,
                        "service": service_a.id,
                        "matchRules": [{
                            "prefixMatch": "/",
                            "ignoreCase": True,
                            "queryParameterMatches": [{
                                "name": "abtest",
                                "exactMatch": "a",
                            }],
                        }],
                    },
                    {
                        "priority": 2,
                        "service": service_b.id,
                        "matchRules": [{
                            "ignoreCase": True,
                            "prefixMatch": "/",
                            "queryParameterMatches": [{
                                "name": "abtest",
                                "exactMatch": "b",
                            }],
                        }],
                    },
                ],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['URLMapDefaultRouteActionArgs']] default_route_action: defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
               advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
               to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
               Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
               Only one of defaultRouteAction or defaultUrlRedirect must be set.  Structure is documented below.
        :param pulumi.Input[str] default_service: The backend service or backend bucket to use when none of the given paths match.
        :param pulumi.Input[pulumi.InputType['URLMapDefaultUrlRedirectArgs']] default_url_redirect: When none of the specified hostRules match, the request is redirected to a URL specified
               by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
               defaultRouteAction must not be set.  Structure is documented below.
        :param pulumi.Input[str] description: Description of this test case.
        :param pulumi.Input[pulumi.InputType['URLMapHeaderActionArgs']] header_action: Specifies changes to request and response headers that need to take effect for
               the selected backendService.
               headerAction specified here take effect before headerAction in the enclosing
               HttpRouteRule, PathMatcher and UrlMap.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapHostRuleArgs']]]] host_rules: The list of HostRules to use against the URL.  Structure is documented below.
        :param pulumi.Input[str] name: The name of the query parameter to match. The query parameter must exist in the
               request, in the absence of which the request match fails.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapPathMatcherArgs']]]] path_matchers: The name of the PathMatcher to use to match the path portion of the URL if the
               hostRule matches the URL's host portion.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapTestArgs']]]] tests: The list of expected URL mapping tests. Request to update this UrlMap will
               succeed only if all of the test cases pass. You can specify a maximum of 100
               tests per UrlMap.  Structure is documented below.
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

            __props__['default_route_action'] = default_route_action
            __props__['default_service'] = default_service
            __props__['default_url_redirect'] = default_url_redirect
            __props__['description'] = description
            __props__['header_action'] = header_action
            __props__['host_rules'] = host_rules
            __props__['name'] = name
            __props__['path_matchers'] = path_matchers
            __props__['project'] = project
            __props__['tests'] = tests
            __props__['creation_timestamp'] = None
            __props__['fingerprint'] = None
            __props__['map_id'] = None
            __props__['self_link'] = None
        super(URLMap, __self__).__init__(
            'gcp:compute/uRLMap:URLMap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            default_route_action: Optional[pulumi.Input[pulumi.InputType['URLMapDefaultRouteActionArgs']]] = None,
            default_service: Optional[pulumi.Input[str]] = None,
            default_url_redirect: Optional[pulumi.Input[pulumi.InputType['URLMapDefaultUrlRedirectArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            header_action: Optional[pulumi.Input[pulumi.InputType['URLMapHeaderActionArgs']]] = None,
            host_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapHostRuleArgs']]]]] = None,
            map_id: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path_matchers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapPathMatcherArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            tests: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapTestArgs']]]]] = None) -> 'URLMap':
        """
        Get an existing URLMap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[pulumi.InputType['URLMapDefaultRouteActionArgs']] default_route_action: defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs
               advanced routing actions like URL rewrites, header transformations, etc. prior to forwarding the request
               to the selected backend. If defaultRouteAction specifies any weightedBackendServices, defaultService must not be set.
               Conversely if defaultService is set, defaultRouteAction cannot contain any weightedBackendServices.
               Only one of defaultRouteAction or defaultUrlRedirect must be set.  Structure is documented below.
        :param pulumi.Input[str] default_service: The backend service or backend bucket to use when none of the given paths match.
        :param pulumi.Input[pulumi.InputType['URLMapDefaultUrlRedirectArgs']] default_url_redirect: When none of the specified hostRules match, the request is redirected to a URL specified
               by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
               defaultRouteAction must not be set.  Structure is documented below.
        :param pulumi.Input[str] description: Description of this test case.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
        :param pulumi.Input[pulumi.InputType['URLMapHeaderActionArgs']] header_action: Specifies changes to request and response headers that need to take effect for
               the selected backendService.
               headerAction specified here take effect before headerAction in the enclosing
               HttpRouteRule, PathMatcher and UrlMap.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapHostRuleArgs']]]] host_rules: The list of HostRules to use against the URL.  Structure is documented below.
        :param pulumi.Input[float] map_id: The unique identifier for the resource.
        :param pulumi.Input[str] name: The name of the query parameter to match. The query parameter must exist in the
               request, in the absence of which the request match fails.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapPathMatcherArgs']]]] path_matchers: The name of the PathMatcher to use to match the path portion of the URL if the
               hostRule matches the URL's host portion.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['URLMapTestArgs']]]] tests: The list of expected URL mapping tests. Request to update this UrlMap will
               succeed only if all of the test cases pass. You can specify a maximum of 100
               tests per UrlMap.  Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["default_route_action"] = default_route_action
        __props__["default_service"] = default_service
        __props__["default_url_redirect"] = default_url_redirect
        __props__["description"] = description
        __props__["fingerprint"] = fingerprint
        __props__["header_action"] = header_action
        __props__["host_rules"] = host_rules
        __props__["map_id"] = map_id
        __props__["name"] = name
        __props__["path_matchers"] = path_matchers
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["tests"] = tests
        return URLMap(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

