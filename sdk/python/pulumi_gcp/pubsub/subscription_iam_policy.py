# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SubscriptionIAMPolicy(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the subscription's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `google_iam_policy` data source.
    """
    project: pulumi.Output[str]
    """
    The project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    subscription: pulumi.Output[str]
    """
    The subscription name or id to bind to attach IAM policy to.
    """
    def __init__(__self__, resource_name, opts=None, policy_data=None, project=None, subscription=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:
        
        * `google_pubsub_subscription_iam_policy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
        * `google_pubsub_subscription_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
        * `google_pubsub_subscription_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.
        
        > **Note:** `google_pubsub_subscription_iam_policy` **cannot** be used in conjunction with `google_pubsub_subscription_iam_binding` and `google_pubsub_subscription_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_pubsub_subscription_iam_binding` resources **can be** used in conjunction with `google_pubsub_subscription_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `google_iam_policy` data source.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] subscription: The subscription name or id to bind to attach IAM policy to.
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

        if not policy_data:
            raise TypeError('Missing required property policy_data')
        __props__['policy_data'] = policy_data

        __props__['project'] = project

        if not subscription:
            raise TypeError('Missing required property subscription')
        __props__['subscription'] = subscription

        __props__['etag'] = None

        super(SubscriptionIAMPolicy, __self__).__init__(
            'gcp:pubsub/subscriptionIAMPolicy:SubscriptionIAMPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

