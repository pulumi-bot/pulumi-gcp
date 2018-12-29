# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class BillingAccountSink(pulumi.CustomResource):
    """
    Manages a billing account logging sink. For more information see
    [the official documentation](https://cloud.google.com/logging/docs/) and
    [Exporting Logs in the API](https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
    
    > **Note** You must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    [granted on the billing account](https://cloud.google.com/billing/reference/rest/v1/billingAccounts/getIamPolicy) to
    the credentials used with Terraform. [IAM roles granted on a billing account](https://cloud.google.com/billing/docs/how-to/billing-access) are separate from the
    typical IAM roles granted on a project.
    """
    def __init__(__self__, __name__, __opts__=None, billing_account=None, destination=None, filter=None, name=None):
        """Create a BillingAccountSink resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not billing_account:
            raise TypeError('Missing required property billing_account')
        __props__['billing_account'] = billing_account

        if not destination:
            raise TypeError('Missing required property destination')
        __props__['destination'] = destination

        __props__['filter'] = filter

        __props__['name'] = name

        __props__['writer_identity'] = None

        super(BillingAccountSink, __self__).__init__(
            'gcp:logging/billingAccountSink:BillingAccountSink',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

