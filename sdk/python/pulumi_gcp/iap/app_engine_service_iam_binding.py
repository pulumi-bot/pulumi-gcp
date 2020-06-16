# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AppEngineServiceIamBinding(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
    """
    condition: pulumi.Output[dict]
    """
    An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
    Structure is documented below.

      * `description` (`str`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
      * `expression` (`str`) - Textual representation of an expression in Common Expression Language syntax.
      * `title` (`str`) - A title for the expression, i.e. a short string describing its purpose.
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the IAM policy.
    """
    members: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    service: pulumi.Output[str]
    """
    Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, condition=None, members=None, project=None, role=None, service=None, __props__=None, __name__=None, __opts__=None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineService. Each of these resources serves a different use case:

        * `iap.AppEngineServiceIamPolicy`: Authoritative. Sets the IAM policy for the appengineservice and replaces any existing policy already attached.
        * `iap.AppEngineServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineservice are preserved.
        * `iap.AppEngineServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineservice are preserved.

        > **Note:** `iap.AppEngineServiceIamPolicy` **cannot** be used in conjunction with `iap.AppEngineServiceIamBinding` and `iap.AppEngineServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.AppEngineServiceIamBinding` resources **can be** used in conjunction with `iap.AppEngineServiceIamMember` resources **only if** they do not grant privilege to the same role.



        ## google\_iap\_app\_engine\_service\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(binding=[{
            "role": "roles/iap.httpsResourceAccessor",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.iap.AppEngineServiceIamPolicy("policy",
            project=google_app_engine_standard_app_version["version"]["project"],
            app_id=google_app_engine_standard_app_version["version"]["project"],
            service=google_app_engine_standard_app_version["version"]["service"],
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
        policy = gcp.iap.AppEngineServiceIamPolicy("policy",
            project=google_app_engine_standard_app_version["version"]["project"],
            app_id=google_app_engine_standard_app_version["version"]["project"],
            service=google_app_engine_standard_app_version["version"]["service"],
            policy_data=admin.policy_data)
        ```
        ## google\_iap\_app\_engine\_service\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.AppEngineServiceIamBinding("binding",
            app_id=google_app_engine_standard_app_version["version"]["project"],
            members=["user:jane@example.com"],
            project=google_app_engine_standard_app_version["version"]["project"],
            role="roles/iap.httpsResourceAccessor",
            service=google_app_engine_standard_app_version["version"]["service"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.AppEngineServiceIamBinding("binding",
            app_id=google_app_engine_standard_app_version["version"]["project"],
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            members=["user:jane@example.com"],
            project=google_app_engine_standard_app_version["version"]["project"],
            role="roles/iap.httpsResourceAccessor",
            service=google_app_engine_standard_app_version["version"]["service"])
        ```
        ## google\_iap\_app\_engine\_service\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.AppEngineServiceIamMember("member",
            app_id=google_app_engine_standard_app_version["version"]["project"],
            member="user:jane@example.com",
            project=google_app_engine_standard_app_version["version"]["project"],
            role="roles/iap.httpsResourceAccessor",
            service=google_app_engine_standard_app_version["version"]["service"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.AppEngineServiceIamMember("member",
            app_id=google_app_engine_standard_app_version["version"]["project"],
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            member="user:jane@example.com",
            project=google_app_engine_standard_app_version["version"]["project"],
            role="roles/iap.httpsResourceAccessor",
            service=google_app_engine_standard_app_version["version"]["service"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] service: Service id of the App Engine application Used to find the parent resource to bind the IAM policy to

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
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
            __props__['condition'] = condition
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            __props__['project'] = project
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['etag'] = None
        super(AppEngineServiceIamBinding, __self__).__init__(
            'gcp:iap/appEngineServiceIamBinding:AppEngineServiceIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, condition=None, etag=None, members=None, project=None, role=None, service=None):
        """
        Get an existing AppEngineServiceIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.AppEngineServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] service: Service id of the App Engine application Used to find the parent resource to bind the IAM policy to

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["members"] = members
        __props__["project"] = project
        __props__["role"] = role
        __props__["service"] = service
        return AppEngineServiceIamBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
