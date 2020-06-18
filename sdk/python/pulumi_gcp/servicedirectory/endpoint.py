# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Endpoint(pulumi.CustomResource):
    address: pulumi.Output[str]
    """
    IPv4 or IPv6 address of the endpoint.
    """
    endpoint_id: pulumi.Output[str]
    """
    The Resource ID must be 1-63 characters long, including digits,
    lowercase letters or the hyphen character.
    """
    metadata: pulumi.Output[dict]
    """
    Metadata for the endpoint. This data can be consumed
    by service clients. The entire metadata dictionary may contain
    up to 512 characters, spread across all key-value pairs.
    Metadata that goes beyond any these limits will be rejected.
    """
    name: pulumi.Output[str]
    """
    The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
    """
    port: pulumi.Output[float]
    """
    Port that the endpoint is running on, must be in the
    range of [0, 65535]. If unspecified, the default is 0.
    """
    service: pulumi.Output[str]
    """
    The resource name of the service that this endpoint provides.
    """
    def __init__(__self__, resource_name, opts=None, address=None, endpoint_id=None, metadata=None, port=None, service=None, __props__=None, __name__=None, __opts__=None):
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
            location="us-central1")
        example_service = gcp.servicedirectory.Service("exampleService",
            service_id="example-service",
            namespace=example_namespace.id)
        example_endpoint = gcp.servicedirectory.Endpoint("exampleEndpoint",
            endpoint_id="example-endpoint",
            service=example_service.id,
            metadata={
                "stage": "prod",
                "region": "us-central1",
            },
            address="1.2.3.4",
            port=5353)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address of the endpoint.
        :param pulumi.Input[str] endpoint_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        :param pulumi.Input[dict] metadata: Metadata for the endpoint. This data can be consumed
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, address=None, endpoint_id=None, metadata=None, name=None, port=None, service=None):
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address of the endpoint.
        :param pulumi.Input[str] endpoint_id: The Resource ID must be 1-63 characters long, including digits,
               lowercase letters or the hyphen character.
        :param pulumi.Input[dict] metadata: Metadata for the endpoint. This data can be consumed
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
