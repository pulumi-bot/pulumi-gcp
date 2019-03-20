# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BackendService(pulumi.CustomResource):
    backends: pulumi.Output[list]
    """
    The list of backends that serve this BackendService. Structure is documented below.
    """
    cdn_policy: pulumi.Output[dict]
    """
    Cloud CDN configuration for this BackendService. Structure is documented below.
    """
    connection_draining_timeout_sec: pulumi.Output[float]
    """
    Time for which instance will be drained (not accept new connections,
    but still work to finish started ones). Defaults to `300`.
    """
    custom_request_headers: pulumi.Output[list]
    """
    Headers that the
    HTTP/S load balancer should add to proxied requests. See [guide](https://cloud.google.com/compute/docs/load-balancing/http/backend-service#user-defined-request-headers) for details.
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    description: pulumi.Output[str]
    """
    The textual description for the backend service.
    """
    enable_cdn: pulumi.Output[bool]
    """
    Whether or not to enable the Cloud CDN on the backend service.
    """
    fingerprint: pulumi.Output[str]
    """
    The fingerprint of the backend service.
    """
    health_checks: pulumi.Output[str]
    """
    Specifies a list of HTTP/HTTPS health checks
    for checking the health of the backend service. Currently at most one health
    check can be specified, and a health check is required.
    """
    iap: pulumi.Output[dict]
    """
    Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
    """
    name: pulumi.Output[str]
    """
    The name of the backend service.
    """
    port_name: pulumi.Output[str]
    """
    The name of a service that has been added to an
    instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    protocol: pulumi.Output[str]
    """
    The protocol for incoming requests. Defaults to
    `HTTP`.
    """
    security_policy: pulumi.Output[str]
    """
    Name or URI of a
    [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    session_affinity: pulumi.Output[str]
    """
    How to distribute load. Options are `NONE` (no
    affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
    `GENERATED_COOKIE` (distribute load using a generated session cookie).
    """
    timeout_sec: pulumi.Output[float]
    """
    The number of secs to wait for a backend to respond
    to a request before considering the request failed. Defaults to `30`.
    """
    def __init__(__self__, resource_name, opts=None, backends=None, cdn_policy=None, connection_draining_timeout_sec=None, custom_request_headers=None, description=None, enable_cdn=None, health_checks=None, iap=None, name=None, port_name=None, project=None, protocol=None, security_policy=None, session_affinity=None, timeout_sec=None, __name__=None, __opts__=None):
        """
        A Backend Service defines a group of virtual machines that will serve traffic for load balancing. For more information
        see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
        and the [API](https://cloud.google.com/compute/docs/reference/latest/backendServices).
        
        For internal load balancing, use a [google_compute_region_backend_service](https://www.terraform.io/docs/providers/google/r/compute_region_backend_service.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] backends: The list of backends that serve this BackendService. Structure is documented below.
        :param pulumi.Input[dict] cdn_policy: Cloud CDN configuration for this BackendService. Structure is documented below.
        :param pulumi.Input[float] connection_draining_timeout_sec: Time for which instance will be drained (not accept new connections,
               but still work to finish started ones). Defaults to `300`.
        :param pulumi.Input[list] custom_request_headers: Headers that the
               HTTP/S load balancer should add to proxied requests. See [guide](https://cloud.google.com/compute/docs/load-balancing/http/backend-service#user-defined-request-headers) for details.
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[str] description: The textual description for the backend service.
        :param pulumi.Input[bool] enable_cdn: Whether or not to enable the Cloud CDN on the backend service.
        :param pulumi.Input[str] health_checks: Specifies a list of HTTP/HTTPS health checks
               for checking the health of the backend service. Currently at most one health
               check can be specified, and a health check is required.
        :param pulumi.Input[dict] iap: Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
        :param pulumi.Input[str] name: The name of the backend service.
        :param pulumi.Input[str] port_name: The name of a service that has been added to an
               instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] protocol: The protocol for incoming requests. Defaults to
               `HTTP`.
        :param pulumi.Input[str] security_policy: Name or URI of a
               [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
        :param pulumi.Input[str] session_affinity: How to distribute load. Options are `NONE` (no
               affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
               `GENERATED_COOKIE` (distribute load using a generated session cookie).
        :param pulumi.Input[float] timeout_sec: The number of secs to wait for a backend to respond
               to a request before considering the request failed. Defaults to `30`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['backends'] = backends

        __props__['cdn_policy'] = cdn_policy

        __props__['connection_draining_timeout_sec'] = connection_draining_timeout_sec

        __props__['custom_request_headers'] = custom_request_headers

        __props__['description'] = description

        __props__['enable_cdn'] = enable_cdn

        if health_checks is None:
            raise TypeError('Missing required property health_checks')
        __props__['health_checks'] = health_checks

        __props__['iap'] = iap

        __props__['name'] = name

        __props__['port_name'] = port_name

        __props__['project'] = project

        __props__['protocol'] = protocol

        __props__['security_policy'] = security_policy

        __props__['session_affinity'] = session_affinity

        __props__['timeout_sec'] = timeout_sec

        __props__['fingerprint'] = None
        __props__['self_link'] = None

        super(BackendService, __self__).__init__(
            'gcp:compute/backendService:BackendService',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

