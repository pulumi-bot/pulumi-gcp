# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Services(pulumi.CustomResource):
    disable_on_destroy: pulumi.Output[bool]
    project: pulumi.Output[str]
    """
    The project ID.
    Changing this forces Terraform to attempt to disable all previously managed
    API services in the previous project.
    """
    services: pulumi.Output[list]
    """
    The list of services that are enabled. Supports
    update.
    """
    def __init__(__self__, resource_name, opts=None, disable_on_destroy=None, project=None, services=None, __name__=None, __opts__=None):
        """
        Allows management of enabled API services for an existing Google Cloud
        Platform project. Services in an existing project that are not defined
        in the config will be removed.
        
        For a list of services available, visit the
        [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
        
        > **Note:** This resource attempts to be the authoritative source on *all* enabled APIs, which often
        	leads to conflicts when certain actions enable other APIs. If you do not need to ensure that
        	*exclusively* a particular set of APIs are enabled, you should most likely use the
        	google_project_service resource, one resource per API.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_on_destroy
        :param pulumi.Input[str] project: The project ID.
               Changing this forces Terraform to attempt to disable all previously managed
               API services in the previous project.
        :param pulumi.Input[list] services: The list of services that are enabled. Supports
               update.
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

        __props__['disable_on_destroy'] = disable_on_destroy

        __props__['project'] = project

        if services is None:
            raise TypeError('Missing required property services')
        __props__['services'] = services

        super(Services, __self__).__init__(
            'gcp:projects/services:Services',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

