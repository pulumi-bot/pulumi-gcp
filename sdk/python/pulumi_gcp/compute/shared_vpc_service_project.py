# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SharedVPCServiceProject(pulumi.CustomResource):
    host_project: pulumi.Output[str]
    """
    The ID of a host project to associate.
    """
    service_project: pulumi.Output[str]
    """
    The ID of the project that will serve as a Shared VPC service project.
    """
    def __init__(__self__, resource_name, opts=None, host_project=None, service_project=None, __name__=None, __opts__=None):
        """
        Enables the Google Compute Engine
        [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
        feature for a project, assigning it as a Shared VPC service project associated
        with a given host project.
        
        For more information, see,
        [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
        where the Shared VPC feature is referred to by its former name "XPN".
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host_project: The ID of a host project to associate.
        :param pulumi.Input[str] service_project: The ID of the project that will serve as a Shared VPC service project.
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

        if host_project is None:
            raise TypeError("Missing required property 'host_project'")
        __props__['host_project'] = host_project

        if service_project is None:
            raise TypeError("Missing required property 'service_project'")
        __props__['service_project'] = service_project

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SharedVPCServiceProject, __self__).__init__(
            'gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

