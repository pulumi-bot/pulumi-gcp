# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMPolicy(pulumi.CustomResource):
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
    The project ID. If not specified for `google_project_iam_binding`
    or `google_project_iam_member`, uses the ID of the project configured with the provider.
    Required for `google_project_iam_policy` - you must explicitly set the project, and it
    will not be inferred from the provider.
    """
    def __init__(__self__, resource_name, opts=None, policy_data=None, project=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:
        
        * `google_project_iam_policy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
        * `google_project_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
        * `google_project_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
        
        > **Note:** `google_project_iam_policy` **cannot** be used in conjunction with `google_project_iam_binding` and `google_project_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_project_iam_binding` resources **can be** used in conjunction with `google_project_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The `google_iam_policy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified for `google_project_iam_binding`
               or `google_project_iam_member`, uses the ID of the project configured with the provider.
               Required for `google_project_iam_policy` - you must explicitly set the project, and it
               will not be inferred from the provider.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_iam_policy.html.markdown.
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

        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data
        if project is None:
            raise TypeError("Missing required property 'project'")
        __props__['project'] = project
        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(IAMPolicy, __self__).__init__(
            'gcp:projects/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

