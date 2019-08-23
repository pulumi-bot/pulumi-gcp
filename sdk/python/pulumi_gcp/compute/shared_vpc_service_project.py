# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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
    def __init__(__self__, resource_name, opts=None, host_project=None, service_project=None, __props__=None, __name__=None, __opts__=None):
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

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_shared_vpc_service_project.html.markdown.
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

            if host_project is None:
                raise TypeError("Missing required property 'host_project'")
            __props__['host_project'] = host_project
            if service_project is None:
                raise TypeError("Missing required property 'service_project'")
            __props__['service_project'] = service_project
        super(SharedVPCServiceProject, __self__).__init__(
            'gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, host_project=None, service_project=None):
        """
        Get an existing SharedVPCServiceProject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] host_project: The ID of a host project to associate.
        :param pulumi.Input[str] service_project: The ID of the project that will serve as a Shared VPC service project.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_shared_vpc_service_project.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["host_project"] = host_project
        __props__["service_project"] = service_project
        return SharedVPCServiceProject(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

