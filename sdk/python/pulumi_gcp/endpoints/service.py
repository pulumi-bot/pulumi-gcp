# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Service(pulumi.CustomResource):
    apis: pulumi.Output[list]
    config_id: pulumi.Output[str]
    dns_address: pulumi.Output[str]
    endpoints: pulumi.Output[list]
    grpc_config: pulumi.Output[str]
    openapi_config: pulumi.Output[str]
    project: pulumi.Output[str]
    protoc_output: pulumi.Output[str]
    protoc_output_base64: pulumi.Output[str]
    service_name: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, grpc_config=None, openapi_config=None, project=None, protoc_output=None, protoc_output_base64=None, service_name=None, __name__=None, __opts__=None):
        """
        This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grpc_config
        :param pulumi.Input[str] openapi_config
        :param pulumi.Input[str] project
        :param pulumi.Input[str] protoc_output
        :param pulumi.Input[str] protoc_output_base64
        :param pulumi.Input[str] service_name
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

        __props__['grpc_config'] = grpc_config

        __props__['openapi_config'] = openapi_config

        __props__['project'] = project

        __props__['protoc_output'] = protoc_output

        __props__['protoc_output_base64'] = protoc_output_base64

        if not service_name:
            raise TypeError('Missing required property service_name')
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


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

