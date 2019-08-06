# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AccountIamPolicy(pulumi.CustomResource):
    billing_account_id: pulumi.Output[str]
    """
    The billing account id.
    """
    etag: pulumi.Output[str]
    policy_data: pulumi.Output[str]
    """
    The `google_iam_policy` data source that represents
    the IAM policy that will be applied to the billing account. This policy overrides any existing
    policy applied to the billing account.
    """
    def __init__(__self__, resource_name, opts=None, billing_account_id=None, policy_data=None, __name__=None, __opts__=None):
        """
        Allows management of the entire IAM policy for an existing Google Cloud Platform Billing Account.
        
        > **Warning:** Billing accounts have a default user that can be **overwritten**
        by use of this resource. The safest alternative is to use multiple `billing.AccountIamBinding`
           resources. If you do use this resource, the best way to be sure that you are
           not making dangerous changes is to start by importing your existing policy,
           and examining the diff very closely.
        
        > **Note:** This resource __must not__ be used in conjunction with
           `billing.AccountIamMember` or `billing.AccountIamBinding`
           or they will fight over what your policy should be.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
        :param pulumi.Input[str] policy_data: The `google_iam_policy` data source that represents
               the IAM policy that will be applied to the billing account. This policy overrides any existing
               policy applied to the billing account.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/billing_account_iam_policy.html.markdown.
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

        if billing_account_id is None:
            raise TypeError("Missing required property 'billing_account_id'")
        __props__['billing_account_id'] = billing_account_id

        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data

        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(AccountIamPolicy, __self__).__init__(
            'gcp:billing/accountIamPolicy:AccountIamPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

