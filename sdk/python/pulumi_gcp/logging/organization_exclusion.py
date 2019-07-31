# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OrganizationExclusion(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A human-readable description.
    """
    disabled: pulumi.Output[bool]
    """
    Whether this exclusion rule should be disabled or not. This defaults to
    false.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when excluding logs. Only log entries that match the filter are excluded.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
    write a filter.
    """
    name: pulumi.Output[str]
    """
    The name of the logging exclusion.
    """
    org_id: pulumi.Output[str]
    """
    The organization to create the exclusion in.
    """
    def __init__(__self__, resource_name, opts=None, description=None, disabled=None, filter=None, name=None, org_id=None, __name__=None, __opts__=None):
        """
        Create a OrganizationExclusion resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        :param pulumi.Input[str] org_id: The organization to create the exclusion in.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/logging_organization_exclusion.html.markdown.
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

        __props__['description'] = description

        __props__['disabled'] = disabled

        if filter is None:
            raise TypeError("Missing required property 'filter'")
        __props__['filter'] = filter

        __props__['name'] = name

        if org_id is None:
            raise TypeError("Missing required property 'org_id'")
        __props__['org_id'] = org_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(OrganizationExclusion, __self__).__init__(
            'gcp:logging/organizationExclusion:OrganizationExclusion',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

