# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Service(pulumi.CustomResource):
    disable_dependent_services: pulumi.Output[bool]
    """
    If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
    If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
    """
    disable_on_destroy: pulumi.Output[bool]
    """
    If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
    """
    project: pulumi.Output[str]
    """
    The project ID. If not provided, the provider project is used.
    """
    service: pulumi.Output[str]
    """
    The service to enable.
    """
    def __init__(__self__, resource_name, opts=None, disable_dependent_services=None, disable_on_destroy=None, project=None, service=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows management of a single API service for an existing Google Cloud Platform project. 

        For a list of services available, visit the
        [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.Service("project",
            disable_dependent_services=True,
            project="your-project-id",
            service="iam.googleapis.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
               If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project is used.
        :param pulumi.Input[str] service: The service to enable.
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

            __props__['disable_dependent_services'] = disable_dependent_services
            __props__['disable_on_destroy'] = disable_on_destroy
            __props__['project'] = project
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
        super(Service, __self__).__init__(
            'gcp:projects/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disable_dependent_services=None, disable_on_destroy=None, project=None, service=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
               If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project is used.
        :param pulumi.Input[str] service: The service to enable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disable_dependent_services"] = disable_dependent_services
        __props__["disable_on_destroy"] = disable_on_destroy
        __props__["project"] = project
        __props__["service"] = service
        return Service(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
