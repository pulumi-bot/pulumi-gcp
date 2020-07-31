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

__all__ = ['HealthCheck']


class HealthCheck(pulumi.CustomResource):
    check_interval_sec: pulumi.Output[Optional[float]] = pulumi.output_property("checkIntervalSec")
    """
    How often (in seconds) to send a health check. The default value is 5
    seconds.
    """
    creation_timestamp: pulumi.Output[str] = pulumi.output_property("creationTimestamp")
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    An optional description of this resource. Provide this property when
    you create the resource.
    """
    healthy_threshold: pulumi.Output[Optional[float]] = pulumi.output_property("healthyThreshold")
    """
    A so-far unhealthy instance will be marked healthy after this many
    consecutive successes. The default value is 2.
    """
    http2_health_check: pulumi.Output[Optional['outputs.HealthCheckHttp2HealthCheck']] = pulumi.output_property("http2HealthCheck")
    """
    A nested object resource  Structure is documented below.
    """
    http_health_check: pulumi.Output[Optional['outputs.HealthCheckHttpHealthCheck']] = pulumi.output_property("httpHealthCheck")
    """
    A nested object resource  Structure is documented below.
    """
    https_health_check: pulumi.Output[Optional['outputs.HealthCheckHttpsHealthCheck']] = pulumi.output_property("httpsHealthCheck")
    """
    A nested object resource  Structure is documented below.
    """
    log_config: pulumi.Output[Optional['outputs.HealthCheckLogConfig']] = pulumi.output_property("logConfig")
    """
    Configure logging on this health check.  Structure is documented below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035.  Specifically, the name must be 1-63 characters long and
    match the regular expression `a-z?` which means
    the first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the
    last character, which cannot be a dash.
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
    ssl_health_check: pulumi.Output[Optional['outputs.HealthCheckSslHealthCheck']] = pulumi.output_property("sslHealthCheck")
    """
    A nested object resource  Structure is documented below.
    """
    tcp_health_check: pulumi.Output[Optional['outputs.HealthCheckTcpHealthCheck']] = pulumi.output_property("tcpHealthCheck")
    """
    A nested object resource  Structure is documented below.
    """
    timeout_sec: pulumi.Output[Optional[float]] = pulumi.output_property("timeoutSec")
    """
    How long (in seconds) to wait before claiming failure.
    The default value is 5 seconds.  It is invalid for timeoutSec to have
    greater value than checkIntervalSec.
    """
    type: pulumi.Output[str] = pulumi.output_property("type")
    """
    The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
    """
    unhealthy_threshold: pulumi.Output[Optional[float]] = pulumi.output_property("unhealthyThreshold")
    """
    A so-far healthy instance will be marked unhealthy after this many
    consecutive failures. The default value is 2.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, check_interval_sec: Optional[pulumi.Input[float]] = None, description: Optional[pulumi.Input[str]] = None, healthy_threshold: Optional[pulumi.Input[float]] = None, http2_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttp2HealthCheckArgs']]] = None, http_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttpHealthCheckArgs']]] = None, https_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttpsHealthCheckArgs']]] = None, log_config: Optional[pulumi.Input[pulumi.InputType['HealthCheckLogConfigArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, ssl_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckSslHealthCheckArgs']]] = None, tcp_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckTcpHealthCheckArgs']]] = None, timeout_sec: Optional[pulumi.Input[float]] = None, unhealthy_threshold: Optional[pulumi.Input[float]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        Health Checks determine whether instances are responsive and able to do work.
        They are an important part of a comprehensive load balancing configuration,
        as they enable monitoring instances behind load balancers.

        Health Checks poll instances at a specified interval. Instances that
        do not respond successfully to some number of probes in a row are marked
        as unhealthy. No new connections are sent to unhealthy instances,
        though existing connections will continue. The health check will
        continue to poll unhealthy instances. If an instance later responds
        successfully to some number of consecutive probes, it is marked
        healthy again and can receive new connections.

        To get more information about HealthCheck, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/healthChecks)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/health-checks)

        ## Example Usage
        ### Health Check Tcp

        ```python
        import pulumi
        import pulumi_gcp as gcp

        tcp_health_check = gcp.compute.HealthCheck("tcp-health-check",
            check_interval_sec=1,
            tcp_health_check={
                "port": "80",
            },
            timeout_sec=1)
        ```
        ### Health Check Tcp Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        tcp_health_check = gcp.compute.HealthCheck("tcp-health-check",
            check_interval_sec=1,
            description="Health check via tcp",
            healthy_threshold=4,
            tcp_health_check={
                "port_name": "health-check-port",
                "portSpecification": "USE_NAMED_PORT",
                "proxy_header": "NONE",
                "request": "ARE YOU HEALTHY?",
                "response": "I AM HEALTHY",
            },
            timeout_sec=1,
            unhealthy_threshold=5)
        ```
        ### Health Check Ssl

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ssl_health_check = gcp.compute.HealthCheck("ssl-health-check",
            check_interval_sec=1,
            ssl_health_check={
                "port": "443",
            },
            timeout_sec=1)
        ```
        ### Health Check Ssl Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ssl_health_check = gcp.compute.HealthCheck("ssl-health-check",
            check_interval_sec=1,
            description="Health check via ssl",
            healthy_threshold=4,
            ssl_health_check={
                "port_name": "health-check-port",
                "portSpecification": "USE_NAMED_PORT",
                "proxy_header": "NONE",
                "request": "ARE YOU HEALTHY?",
                "response": "I AM HEALTHY",
            },
            timeout_sec=1,
            unhealthy_threshold=5)
        ```
        ### Health Check Http

        ```python
        import pulumi
        import pulumi_gcp as gcp

        http_health_check = gcp.compute.HealthCheck("http-health-check",
            check_interval_sec=1,
            http_health_check={
                "port": 80,
            },
            timeout_sec=1)
        ```
        ### Health Check Http Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        http_health_check = gcp.compute.HealthCheck("http-health-check",
            check_interval_sec=1,
            description="Health check via http",
            healthy_threshold=4,
            http_health_check={
                "host": "1.2.3.4",
                "port_name": "health-check-port",
                "portSpecification": "USE_NAMED_PORT",
                "proxy_header": "NONE",
                "request_path": "/mypath",
                "response": "I AM HEALTHY",
            },
            timeout_sec=1,
            unhealthy_threshold=5)
        ```
        ### Health Check Https

        ```python
        import pulumi
        import pulumi_gcp as gcp

        https_health_check = gcp.compute.HealthCheck("https-health-check",
            check_interval_sec=1,
            https_health_check={
                "port": "443",
            },
            timeout_sec=1)
        ```
        ### Health Check Https Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        https_health_check = gcp.compute.HealthCheck("https-health-check",
            check_interval_sec=1,
            description="Health check via https",
            healthy_threshold=4,
            https_health_check={
                "host": "1.2.3.4",
                "port_name": "health-check-port",
                "portSpecification": "USE_NAMED_PORT",
                "proxy_header": "NONE",
                "request_path": "/mypath",
                "response": "I AM HEALTHY",
            },
            timeout_sec=1,
            unhealthy_threshold=5)
        ```
        ### Health Check Http2

        ```python
        import pulumi
        import pulumi_gcp as gcp

        http2_health_check = gcp.compute.HealthCheck("http2-health-check",
            check_interval_sec=1,
            http2_health_check={
                "port": "443",
            },
            timeout_sec=1)
        ```
        ### Health Check Http2 Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        http2_health_check = gcp.compute.HealthCheck("http2-health-check",
            check_interval_sec=1,
            description="Health check via http2",
            healthy_threshold=4,
            http2_health_check={
                "host": "1.2.3.4",
                "port_name": "health-check-port",
                "portSpecification": "USE_NAMED_PORT",
                "proxy_header": "NONE",
                "request_path": "/mypath",
                "response": "I AM HEALTHY",
            },
            timeout_sec=1,
            unhealthy_threshold=5)
        ```
        ### Health Check With Logging

        ```python
        import pulumi
        import pulumi_gcp as gcp

        health_check_with_logging = gcp.compute.HealthCheck("health-check-with-logging",
            timeout_sec=1,
            check_interval_sec=1,
            tcp_health_check={
                "port": "22",
            },
            log_config={
                "enable": True,
            },
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] check_interval_sec: How often (in seconds) to send a health check. The default value is 5
               seconds.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[float] healthy_threshold: A so-far unhealthy instance will be marked healthy after this many
               consecutive successes. The default value is 2.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttp2HealthCheckArgs']] http2_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttpHealthCheckArgs']] http_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttpsHealthCheckArgs']] https_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckLogConfigArgs']] log_config: Configure logging on this health check.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['HealthCheckSslHealthCheckArgs']] ssl_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckTcpHealthCheckArgs']] tcp_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[float] timeout_sec: How long (in seconds) to wait before claiming failure.
               The default value is 5 seconds.  It is invalid for timeoutSec to have
               greater value than checkIntervalSec.
        :param pulumi.Input[float] unhealthy_threshold: A so-far healthy instance will be marked unhealthy after this many
               consecutive failures. The default value is 2.
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

            __props__['check_interval_sec'] = check_interval_sec
            __props__['description'] = description
            __props__['healthy_threshold'] = healthy_threshold
            __props__['http2_health_check'] = http2_health_check
            __props__['http_health_check'] = http_health_check
            __props__['https_health_check'] = https_health_check
            __props__['log_config'] = log_config
            __props__['name'] = name
            __props__['project'] = project
            __props__['ssl_health_check'] = ssl_health_check
            __props__['tcp_health_check'] = tcp_health_check
            __props__['timeout_sec'] = timeout_sec
            __props__['unhealthy_threshold'] = unhealthy_threshold
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
            __props__['type'] = None
        super(HealthCheck, __self__).__init__(
            'gcp:compute/healthCheck:HealthCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, check_interval_sec: Optional[pulumi.Input[float]] = None, creation_timestamp: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, healthy_threshold: Optional[pulumi.Input[float]] = None, http2_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttp2HealthCheckArgs']]] = None, http_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttpHealthCheckArgs']]] = None, https_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckHttpsHealthCheckArgs']]] = None, log_config: Optional[pulumi.Input[pulumi.InputType['HealthCheckLogConfigArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, self_link: Optional[pulumi.Input[str]] = None, ssl_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckSslHealthCheckArgs']]] = None, tcp_health_check: Optional[pulumi.Input[pulumi.InputType['HealthCheckTcpHealthCheckArgs']]] = None, timeout_sec: Optional[pulumi.Input[float]] = None, type: Optional[pulumi.Input[str]] = None, unhealthy_threshold: Optional[pulumi.Input[float]] = None) -> 'HealthCheck':
        """
        Get an existing HealthCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] check_interval_sec: How often (in seconds) to send a health check. The default value is 5
               seconds.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[float] healthy_threshold: A so-far unhealthy instance will be marked healthy after this many
               consecutive successes. The default value is 2.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttp2HealthCheckArgs']] http2_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttpHealthCheckArgs']] http_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckHttpsHealthCheckArgs']] https_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckLogConfigArgs']] log_config: Configure logging on this health check.  Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['HealthCheckSslHealthCheckArgs']] ssl_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[pulumi.InputType['HealthCheckTcpHealthCheckArgs']] tcp_health_check: A nested object resource  Structure is documented below.
        :param pulumi.Input[float] timeout_sec: How long (in seconds) to wait before claiming failure.
               The default value is 5 seconds.  It is invalid for timeoutSec to have
               greater value than checkIntervalSec.
        :param pulumi.Input[str] type: The type of the health check. One of HTTP, HTTPS, TCP, or SSL.
        :param pulumi.Input[float] unhealthy_threshold: A so-far healthy instance will be marked unhealthy after this many
               consecutive failures. The default value is 2.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["check_interval_sec"] = check_interval_sec
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["healthy_threshold"] = healthy_threshold
        __props__["http2_health_check"] = http2_health_check
        __props__["http_health_check"] = http_health_check
        __props__["https_health_check"] = https_health_check
        __props__["log_config"] = log_config
        __props__["name"] = name
        __props__["project"] = project
        __props__["self_link"] = self_link
        __props__["ssl_health_check"] = ssl_health_check
        __props__["tcp_health_check"] = tcp_health_check
        __props__["timeout_sec"] = timeout_sec
        __props__["type"] = type
        __props__["unhealthy_threshold"] = unhealthy_threshold
        return HealthCheck(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

