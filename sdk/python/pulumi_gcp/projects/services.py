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
    """
    Whether or not to disable APIs on project
    when destroyed. Defaults to true. **Note**: When `disable_on_destroy` is
    true and the project is changed, this provider will force disable API services
    managed by this provider for the previous project.
    """
    project: pulumi.Output[str]
    """
    The project ID.
    Changing this forces this provider to attempt to disable all previously managed
    API services in the previous project.
    """
    services: pulumi.Output[list]
    """
    The list of services that are enabled. Supports
    update.
    """
    def __init__(__self__, resource_name, opts=None, disable_on_destroy=None, project=None, services=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[bool] disable_on_destroy: Whether or not to disable APIs on project
               when destroyed. Defaults to true. **Note**: When `disable_on_destroy` is
               true and the project is changed, this provider will force disable API services
               managed by this provider for the previous project.
        :param pulumi.Input[str] project: The project ID.
               Changing this forces this provider to attempt to disable all previously managed
               API services in the previous project.
        :param pulumi.Input[list] services: The list of services that are enabled. Supports
               update.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_services.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['disable_on_destroy'] = disable_on_destroy
            __props__['project'] = project
            if services is None:
                raise TypeError("Missing required property 'services'")
            __props__['services'] = services
        super(Services, __self__).__init__(
            'gcp:projects/services:Services',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disable_on_destroy=None, project=None, services=None):
        """
        Get an existing Services resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_on_destroy: Whether or not to disable APIs on project
               when destroyed. Defaults to true. **Note**: When `disable_on_destroy` is
               true and the project is changed, this provider will force disable API services
               managed by this provider for the previous project.
        :param pulumi.Input[str] project: The project ID.
               Changing this forces this provider to attempt to disable all previously managed
               API services in the previous project.
        :param pulumi.Input[list] services: The list of services that are enabled. Supports
               update.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_services.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["disable_on_destroy"] = disable_on_destroy
        __props__["project"] = project
        __props__["services"] = services
        return Services(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

