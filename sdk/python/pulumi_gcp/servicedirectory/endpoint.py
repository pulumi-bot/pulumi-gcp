# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Endpoint']


class Endpoint(pulumi.CustomResource):
    address: pulumi.Output[Optional[str]] = pulumi.output_property("address")
    """
    IPv4 or IPv6 address of the endpoint.
    """
    endpoint_id: pulumi.Output[str] = pulumi.output_property("endpointId")
    """
    The Resource ID must be 1-63 characters long, including digits,
    lowercase letters or the hyphen character.
    """
    metadata: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("metadata")
    """
    Metadata for the endpoint. This data can be consumed
    by service clients. The entire metadata dictionary may contain
    up to 512 characters, spread across all key-value pairs.
    Metadata that goes beyond any these limits will be rejected.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
    """
    port: pulumi.Output[Optional[float]] = pulumi.output_property("port")
    """
    Port that the endpoint is running on, must be in the
    range of [0, 65535]. If unspecified, the default is 0.
    """
    service: pulumi.Output[str] = pulumi.output_property("service")
    """
    The resource name of the service that this endpoint provides.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, address: Optional[pulumi.Input[str]] = None, endpoint_id: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, port: Optional[pulumi.Input[float]] = None, service: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        An individual endpoint that provides a service.

        To get more information about Endpoint, see:

        * [API documentation](https://cloud.google.com/service-directory/docs/reference/rest/v1beta1/projects.locations.namespaces.services.endpoints)
        * How-to Guides
            * [Configuring an endpoint](https://cloud.google.com/service-directory/docs/configuring-service-directory#configuring_an_endpoint)

        ## Example Usage
        ### Service Directory Endpoint Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_namespace = gcp.servicedirectory.Namespace("exampleNamespace",
            namespace_id="example-namespace",
            location="us-central1",
            opts=ResourceOptions(provider=google_beta))
        example_service = gcp.servicedirectory.Service("exampleService",
            service_id="example-service",
            namespace=example_namespace.id,
            opts=ResourceOptions(provider=google_beta))
        example_endpoint = gcp.servicedirectory.Endpoint("exampleEndpoint",
            endpoint_id="example-endpoint",
            service=example_service.id,
            metadata={
                "stage": "prod",
                "region": "us-central1",
            },
            address="1.2.3.4",
            port=5353,
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address of the endpoint.
        :param pulumi.Input[str] endpoint_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: Metadata for the endpoint. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 512 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[float] port: Port that the endpoint is running on, must be in the
               range of [0, 65535]. If unspecified, the default is 0.
        :param pulumi.Input[str] service: The resource name of the service that this endpoint provides.
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

            __props__['address'] = address
            if endpoint_id is None:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__['endpoint_id'] = endpoint_id
            __props__['metadata'] = metadata
            __props__['port'] = port
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['name'] = None
        super(Endpoint, __self__).__init__(
            'gcp:servicedirectory/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, address: Optional[pulumi.Input[str]] = None, endpoint_id: Optional[pulumi.Input[str]] = None, metadata: Optional[pulumi.Input[Dict[str, pulumi.Input[str]]]] = None, name: Optional[pulumi.Input[str]] = None, port: Optional[pulumi.Input[float]] = None, service: Optional[pulumi.Input[str]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address of the endpoint.
        :param pulumi.Input[str] endpoint_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] metadata: Metadata for the endpoint. This data can be consumed
               by service clients. The entire metadata dictionary may contain
               up to 512 characters, spread across all key-value pairs.
               Metadata that goes beyond any these limits will be rejected.
        :param pulumi.Input[str] name: The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
        :param pulumi.Input[float] port: Port that the endpoint is running on, must be in the
               range of [0, 65535]. If unspecified, the default is 0.
        :param pulumi.Input[str] service: The resource name of the service that this endpoint provides.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["endpoint_id"] = endpoint_id
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["port"] = port
        __props__["service"] = service
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

