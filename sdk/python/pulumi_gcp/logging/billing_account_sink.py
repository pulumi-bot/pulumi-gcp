# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BillingAccountSink(pulumi.CustomResource):
    billing_account: pulumi.Output[str]
    """
    The billing account exported to the sink.
    """
    destination: pulumi.Output[str]
    """
    The destination of the sink (or, in other words, where logs are written to). Can be a
    Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
    The writer associated with the sink must have access to write to the above resource.
    """
    filter: pulumi.Output[str]
    """
    The filter to apply when exporting logs. Only log entries that match the filter are exported.
    See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
    write a filter.
    """
    name: pulumi.Output[str]
    """
    The name of the logging sink.
    """
    writer_identity: pulumi.Output[str]
    """
    The identity associated with this sink. This identity must be granted write access to the
    configured `destination`.
    """
    def __init__(__self__, resource_name, opts=None, billing_account=None, destination=None, filter=None, name=None, __name__=None, __opts__=None):
        """
        Manages a billing account logging sink. For more information see
        [the official documentation](https://cloud.google.com/logging/docs/) and
        [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        
        > **Note** You must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
        [granted on the billing account](https://cloud.google.com/billing/reference/rest/v1/billingAccounts/getIamPolicy) to
        the credentials used with Terraform. [IAM roles granted on a billing account](https://cloud.google.com/billing/docs/how-to/billing-access) are separate from the
        typical IAM roles granted on a project.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account: The billing account exported to the sink.
        :param pulumi.Input[str] destination: The destination of the sink (or, in other words, where logs are written to). Can be a
               Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples:
               The writer associated with the sink must have access to write to the above resource.
        :param pulumi.Input[str] filter: The filter to apply when exporting logs. Only log entries that match the filter are exported.
               See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
               write a filter.
        :param pulumi.Input[str] name: The name of the logging sink.
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

        if billing_account is None:
            raise TypeError("Missing required property 'billing_account'")
        __props__['billing_account'] = billing_account

        if destination is None:
            raise TypeError("Missing required property 'destination'")
        __props__['destination'] = destination

        __props__['filter'] = filter

        __props__['name'] = name

        __props__['writer_identity'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(BillingAccountSink, __self__).__init__(
            'gcp:logging/billingAccountSink:BillingAccountSink',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

