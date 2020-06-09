# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class HttpHealthCheck(pulumi.CustomResource):
    check_interval_sec: pulumi.Output[float]
    """
    How often (in seconds) to send a health check. The default value is 5
    seconds.
    """
    creation_timestamp: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    An optional description of this resource. Provide this property when
    you create the resource.
    """
    healthy_threshold: pulumi.Output[float]
    """
    A so-far unhealthy instance will be marked healthy after this many
    consecutive successes. The default value is 2.
    """
    host: pulumi.Output[str]
    """
    The value of the host header in the HTTP health check request. If
    left empty (default value), the public IP on behalf of which this
    health check is performed will be used.
    """
    name: pulumi.Output[str]
    """
    Name of the resource. Provided by the client when the resource is
    created. The name must be 1-63 characters long, and comply with
    RFC1035.  Specifically, the name must be 1-63 characters long and
    match the regular expression `a-z?` which means
    the first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the
    last character, which cannot be a dash.
    """
    port: pulumi.Output[float]
    """
    The TCP port number for the HTTP health check request.
    The default value is 80.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    request_path: pulumi.Output[str]
    """
    The request path of the HTTP health check request.
    The default value is /.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    timeout_sec: pulumi.Output[float]
    """
    How long (in seconds) to wait before claiming failure.
    The default value is 5 seconds.  It is invalid for timeoutSec to have
    greater value than checkIntervalSec.
    """
    unhealthy_threshold: pulumi.Output[float]
    """
    A so-far healthy instance will be marked unhealthy after this many
    consecutive failures. The default value is 2.
    """
    def __init__(__self__, resource_name, opts=None, check_interval_sec=None, description=None, healthy_threshold=None, host=None, name=None, port=None, project=None, request_path=None, timeout_sec=None, unhealthy_threshold=None, __props__=None, __name__=None, __opts__=None):
        """
        An HttpHealthCheck resource. This resource defines a template for how
        individual VMs should be checked for health, via HTTP.


        > **Note:** compute.HttpHealthCheck is a legacy health check.
        The newer [compute.HealthCheck](https://www.terraform.io/docs/providers/google/r/compute_health_check.html)
        should be preferred for all uses except
        [Network Load Balancers](https://cloud.google.com/compute/docs/load-balancing/network/)
        which still require the legacy version.


        To get more information about HttpHealthCheck, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/httpHealthChecks)
        * How-to Guides
            * [Adding Health Checks](https://cloud.google.com/compute/docs/load-balancing/health-checks#legacy_health_checks)

        ## Example Usage

        ### Http Health Check Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.HttpHealthCheck("default",
            check_interval_sec=1,
            request_path="/health_check",
            timeout_sec=1)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] check_interval_sec: How often (in seconds) to send a health check. The default value is 5
               seconds.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[float] healthy_threshold: A so-far unhealthy instance will be marked healthy after this many
               consecutive successes. The default value is 2.
        :param pulumi.Input[str] host: The value of the host header in the HTTP health check request. If
               left empty (default value), the public IP on behalf of which this
               health check is performed will be used.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[float] port: The TCP port number for the HTTP health check request.
               The default value is 80.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] request_path: The request path of the HTTP health check request.
               The default value is /.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['check_interval_sec'] = check_interval_sec
            __props__['description'] = description
            __props__['healthy_threshold'] = healthy_threshold
            __props__['host'] = host
            __props__['name'] = name
            __props__['port'] = port
            __props__['project'] = project
            __props__['request_path'] = request_path
            __props__['timeout_sec'] = timeout_sec
            __props__['unhealthy_threshold'] = unhealthy_threshold
            __props__['creation_timestamp'] = None
            __props__['self_link'] = None
        super(HttpHealthCheck, __self__).__init__(
            'gcp:compute/httpHealthCheck:HttpHealthCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, check_interval_sec=None, creation_timestamp=None, description=None, healthy_threshold=None, host=None, name=None, port=None, project=None, request_path=None, self_link=None, timeout_sec=None, unhealthy_threshold=None):
        """
        Get an existing HttpHealthCheck resource's state with the given name, id, and optional extra
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
        :param pulumi.Input[str] host: The value of the host header in the HTTP health check request. If
               left empty (default value), the public IP on behalf of which this
               health check is performed will be used.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
        :param pulumi.Input[float] port: The TCP port number for the HTTP health check request.
               The default value is 80.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] request_path: The request path of the HTTP health check request.
               The default value is /.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[float] timeout_sec: How long (in seconds) to wait before claiming failure.
               The default value is 5 seconds.  It is invalid for timeoutSec to have
               greater value than checkIntervalSec.
        :param pulumi.Input[float] unhealthy_threshold: A so-far healthy instance will be marked unhealthy after this many
               consecutive failures. The default value is 2.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["check_interval_sec"] = check_interval_sec
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["healthy_threshold"] = healthy_threshold
        __props__["host"] = host
        __props__["name"] = name
        __props__["port"] = port
        __props__["project"] = project
        __props__["request_path"] = request_path
        __props__["self_link"] = self_link
        __props__["timeout_sec"] = timeout_sec
        __props__["unhealthy_threshold"] = unhealthy_threshold
        return HttpHealthCheck(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

