# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMMember(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the service account IAM policy.
    """
    member: pulumi.Output[str]
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_service_account_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    service_account_id: pulumi.Output[str]
    """
    The fully-qualified name of the service account to apply policy to.
    """
    def __init__(__self__, resource_name, opts=None, member=None, role=None, service_account_id=None, __name__=None, __opts__=None):
        """
        When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the google_project_iam set of resources.
        
        Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
        
        * `google_service_account_iam_policy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
        * `google_service_account_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
        * `google_service_account_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
        
        > **Note:** `google_service_account_iam_policy` **cannot** be used in conjunction with `google_service_account_iam_binding` and `google_service_account_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_service_account_iam_binding` resources **can be** used in conjunction with `google_service_account_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_service_account_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account_iam_member.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if member is None:
            raise TypeError("Missing required property 'member'")
        __props__['member'] = member
        if role is None:
            raise TypeError("Missing required property 'role'")
        __props__['role'] = role
        if service_account_id is None:
            raise TypeError("Missing required property 'service_account_id'")
        __props__['service_account_id'] = service_account_id
        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(IAMMember, __self__).__init__(
            'gcp:serviceAccount/iAMMember:IAMMember',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

