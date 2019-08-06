# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DatabaseIAMPolicy(pulumi.CustomResource):
    database: pulumi.Output[str]
    """
    The name of the Spanner database.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the database's IAM policy.
    """
    instance: pulumi.Output[str]
    """
    The name of the Spanner instance the database belongs to.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `google_iam_policy` data source.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, database=None, instance=None, policy_data=None, project=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
        
        * `google_spanner_database_iam_policy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
        
        > **Warning:** It's entirely possibly to lock yourself out of your database using `google_spanner_database_iam_policy`. Any permissions granted by default will be removed unless you include them in your config.
        
        * `google_spanner_database_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
        * `google_spanner_database_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
        
        > **Note:** `google_spanner_database_iam_policy` **cannot** be used in conjunction with `google_spanner_database_iam_binding` and `google_spanner_database_iam_member` or they will fight over what your policy should be.
        
        > **Note:** `google_spanner_database_iam_binding` resources **can be** used in conjunction with `google_spanner_database_iam_member` resources **only if** they do not grant privilege to the same role.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The name of the Spanner database.
        :param pulumi.Input[str] instance: The name of the Spanner instance the database belongs to.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `google_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/spanner_database_iam_policy.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__

        __props__ = dict()

        if database is None:
            raise TypeError("Missing required property 'database'")
        __props__['database'] = database
        if instance is None:
            raise TypeError("Missing required property 'instance'")
        __props__['instance'] = instance
        if policy_data is None:
            raise TypeError("Missing required property 'policy_data'")
        __props__['policy_data'] = policy_data
        __props__['project'] = project
        __props__['etag'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(DatabaseIAMPolicy, __self__).__init__(
            'gcp:spanner/databaseIAMPolicy:DatabaseIAMPolicy',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

