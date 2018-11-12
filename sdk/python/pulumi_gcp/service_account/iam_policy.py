# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMPolicy(pulumi.CustomResource):
    """
    When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource to configure permissions for who can edit the service account. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the google_project_iam set of resources.
    
    Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
    
    * `google_service_account_iam_policy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
    * `google_service_account_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
    * `google_service_account_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
    
    ~> **Note:** `google_service_account_iam_policy` **cannot** be used in conjunction with `google_service_account_iam_binding` and `google_service_account_iam_member` or they will fight over what your policy should be.
    
    ~> **Note:** `google_service_account_iam_binding` resources **can be** used in conjunction with `google_service_account_iam_member` resources **only if** they do not grant privilege to the same role.
    """
    def __init__(__self__, __name__, __opts__=None, policy_data=None, service_account_id=None):
        """Create a IAMPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not policy_data:
            raise TypeError('Missing required property policy_data')
        __props__['policy_data'] = policy_data

        if not service_account_id:
            raise TypeError('Missing required property service_account_id')
        __props__['service_account_id'] = service_account_id

        __props__['etag'] = None

        super(IAMPolicy, __self__).__init__(
            'gcp:serviceAccount/iAMPolicy:IAMPolicy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

