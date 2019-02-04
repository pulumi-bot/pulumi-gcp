# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMPolicy(pulumi.CustomResource):
    authoritative: pulumi.Output[bool]
    """
    (Optional, only for `google_project_iam_policy`)
    A boolean value indicating if this policy
    should overwrite any existing IAM policy on the project. When set to true,
    **any policies not in your config file will be removed**. This can **lock
    you out** of your project until an Organization Administrator grants you
    access again, so please exercise caution. If this argument is `true` and you
    want to delete the resource, you must set the `disable_project` argument to
    `true`, acknowledging that the project will be inaccessible to anyone but the
    Organization Admins, as it will no longer have an IAM policy. Rather than using
    this, you should use `google_project_iam_binding` and
    `google_project_iam_member`.
    """
    disable_project: pulumi.Output[bool]
    """
    (Optional, only for `google_project_iam_policy`)
    A boolean value that must be set to `true`
    if you want to delete a `google_project_iam_policy` that is authoritative.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the project's IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The `google_iam_policy` data source that represents
    the IAM policy that will be applied to the project. The policy will be
    merged with any existing policy applied to the project.
    """
    project: pulumi.Output[str]
    """
    The project ID. If not specified, uses the
    ID of the project configured with the provider.
    """
    restore_policy: pulumi.Output[str]
    """
    (DEPRECATED) (Computed, only for `google_project_iam_policy`)
    The IAM policy that will be restored when a
    non-authoritative policy resource is deleted.
    """
    def __init__(__self__, resource_name, opts=None, authoritative=None, disable_project=None, policy_data=None, project=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
        
        * `google_project_iam_policy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
        * `google_project_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
        * `google_project_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
        
        > **Note:** `google_project_iam_policy` **cannot** be used in conjunction with `google_project_iam_binding` and `google_project_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_project_iam_binding` resources **can be** used in conjunction with `google_project_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authoritative: (Optional, only for `google_project_iam_policy`)
               A boolean value indicating if this policy
               should overwrite any existing IAM policy on the project. When set to true,
               **any policies not in your config file will be removed**. This can **lock
               you out** of your project until an Organization Administrator grants you
               access again, so please exercise caution. If this argument is `true` and you
               want to delete the resource, you must set the `disable_project` argument to
               `true`, acknowledging that the project will be inaccessible to anyone but the
               Organization Admins, as it will no longer have an IAM policy. Rather than using
               this, you should use `google_project_iam_binding` and
               `google_project_iam_member`.
        :param pulumi.Input[bool] disable_project: (Optional, only for `google_project_iam_policy`)
               A boolean value that must be set to `true`
               if you want to delete a `google_project_iam_policy` that is authoritative.
        :param pulumi.Input[str] policy_data: The `google_iam_policy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified, uses the
               ID of the project configured with the provider.
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

        __props__['authoritative'] = authoritative

        __props__['disable_project'] = disable_project

        if not policy_data:
            raise TypeError('Missing required property policy_data')
        __props__['policy_data'] = policy_data

        __props__['project'] = project

        __props__['etag'] = None
        __props__['restore_policy'] = None

        super(IAMPolicy, __self__).__init__(
            'gcp:projects/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

