# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMBinding(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the organization's IAM policy.
    """
    members: pulumi.Output[list]
    """
    A list of users that the role should apply to.
    """
    org_id: pulumi.Output[str]
    """
    The numeric ID of the organization in which you want to create a custom role.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, members=None, org_id=None, role=None, __name__=None, __opts__=None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud Platform Organization.
        
        > **Note:** This resource __must not__ be used in conjunction with
           `google_organization_iam_member` for the __same role__ or they will fight over
           what your policy should be.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] members: A list of users that the role should apply to.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
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

        if members is None:
            raise TypeError('Missing required property members')
        __props__['members'] = members

        if org_id is None:
            raise TypeError('Missing required property org_id')
        __props__['org_id'] = org_id

        if role is None:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(IAMBinding, __self__).__init__(
            'gcp:organizations/iAMBinding:IAMBinding',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

