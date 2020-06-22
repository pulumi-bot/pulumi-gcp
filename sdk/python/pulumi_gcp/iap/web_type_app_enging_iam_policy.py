# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class WebTypeAppEngingIamPolicy(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the IAM policy.
    """
    policy_data: pulumi.Output[str]
    """
    The policy data generated by
    a `organizations.getIAMPolicy` data source.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, policy_data=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:

        * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
        * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
        * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.

        > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_iap\_web\_type\_app\_engine\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/iap.httpsResourceAccessor",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/iap.httpsResourceAccessor",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            },
        }])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```
        ## google\_iap\_web\_type\_app\_engine\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```
        ## google\_iap\_web\_type\_app\_engine\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if app_id is None:
                raise TypeError("Missing required property 'app_id'")
            __props__['app_id'] = app_id
            if policy_data is None:
                raise TypeError("Missing required property 'policy_data'")
            __props__['policy_data'] = policy_data
            __props__['project'] = project
            __props__['etag'] = None
        super(WebTypeAppEngingIamPolicy, __self__).__init__(
            'gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, etag=None, policy_data=None, project=None):
        """
        Get an existing WebTypeAppEngingIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["etag"] = etag
        __props__["policy_data"] = policy_data
        __props__["project"] = project
        return WebTypeAppEngingIamPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
