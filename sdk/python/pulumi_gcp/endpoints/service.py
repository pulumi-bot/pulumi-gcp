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

__all__ = ['Service']


class Service(pulumi.CustomResource):
    apis: pulumi.Output[List['outputs.ServiceApi']] = pulumi.property("apis")
    """
    A list of API objects.
    """

    config_id: pulumi.Output[str] = pulumi.property("configId")
    """
    The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
    to compute engine instances as a tag.
    """

    dns_address: pulumi.Output[str] = pulumi.property("dnsAddress")
    """
    The address at which the service can be found - usually the same as the service name.
    """

    endpoints: pulumi.Output[List['outputs.ServiceEndpoint']] = pulumi.property("endpoints")
    """
    A list of Endpoint objects.
    """

    grpc_config: pulumi.Output[Optional[str]] = pulumi.property("grpcConfig")
    """
    The full text of the Service Config YAML file (Example located here). If provided, must also provide
    protoc_output_base64. open_api config must not be provided.
    """

    openapi_config: pulumi.Output[Optional[str]] = pulumi.property("openapiConfig")
    """
    The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
    protoc_output_base64 must be specified.
    """

    project: pulumi.Output[str] = pulumi.property("project")
    """
    The project ID that the service belongs to. If not provided, provider project is used.
    """

    protoc_output_base64: pulumi.Output[Optional[str]] = pulumi.property("protocOutputBase64")
    """
    The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
    base64-encoded.
    """

    service_name: pulumi.Output[str] = pulumi.property("serviceName")
    """
    The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grpc_config: Optional[pulumi.Input[str]] = None,
                 openapi_config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protoc_output_base64: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
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

            __props__['grpc_config'] = grpc_config
            __props__['openapi_config'] = openapi_config
            __props__['project'] = project
            __props__['protoc_output_base64'] = protoc_output_base64
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['apis'] = None
            __props__['config_id'] = None
            __props__['dns_address'] = None
            __props__['endpoints'] = None
        super(Service, __self__).__init__(
            'gcp:endpoints/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            apis: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ServiceApiArgs']]]]] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            dns_address: Optional[pulumi.Input[str]] = None,
            endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ServiceEndpointArgs']]]]] = None,
            grpc_config: Optional[pulumi.Input[str]] = None,
            openapi_config: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            protoc_output_base64: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ServiceApiArgs']]]] apis: A list of API objects.
        :param pulumi.Input[str] config_id: The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
               to compute engine instances as a tag.
        :param pulumi.Input[str] dns_address: The address at which the service can be found - usually the same as the service name.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ServiceEndpointArgs']]]] endpoints: A list of Endpoint objects.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apis"] = apis
        __props__["config_id"] = config_id
        __props__["dns_address"] = dns_address
        __props__["endpoints"] = endpoints
        __props__["grpc_config"] = grpc_config
        __props__["openapi_config"] = openapi_config
        __props__["project"] = project
        __props__["protoc_output_base64"] = protoc_output_base64
        __props__["service_name"] = service_name
        return Service(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

