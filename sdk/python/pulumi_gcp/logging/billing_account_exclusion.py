# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BillingAccountExclusion(pulumi.CustomResource):
    billing_account: pulumi.Output[str]
    """
    The billing account to create the exclusion for.
    """
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
    def __init__(__self__, __name__, __opts__=None, billing_account=None, description=None, disabled=None, filter=None, name=None):
        """
        Manages a billing account logging exclusion. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
        
        Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        granted to the credentials used with Terraform.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] billing_account: The billing account to create the exclusion for.
        :param pulumi.Input[str] description: A human-readable description.
        :param pulumi.Input[bool] disabled: Whether this exclusion rule should be disabled or not. This defaults to
               false.
        :param pulumi.Input[str] filter: The filter to apply when excluding logs. Only log entries that match the filter are excluded.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging exclusion.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if billing_account is None:
            raise TypeError('Missing required property billing_account')
        __props__['billing_account'] = billing_account

        __props__['description'] = description

        __props__['disabled'] = disabled

        if filter is None:
            raise TypeError('Missing required property filter')
        __props__['filter'] = filter

        __props__['name'] = name

        super(BillingAccountExclusion, __self__).__init__(
            'gcp:logging/billingAccountExclusion:BillingAccountExclusion',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

