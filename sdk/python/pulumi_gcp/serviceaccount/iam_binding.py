# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class IAMBinding(pulumi.CustomResource):
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
    (Computed) The etag of the service account IAM policy.
    """
    members: pulumi.Output[list]
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    service_account_id: pulumi.Output[str]
    """
    The fully-qualified name of the service account to apply policy to.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, members=None, role=None, service_account_id=None, __props__=None, __name__=None, __opts__=None):
        """
        When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource **to configure permissions for who can edit the service account**. To configure permissions for a service account to act as an identity that can manage other GCP resources, use the google_project_iam set of resources.

        Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:

        * `serviceAccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
        * `serviceAccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
        * `serviceAccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.

        > **Note:** `serviceAccount.IAMPolicy` **cannot** be used in conjunction with `serviceAccount.IAMBinding` and `serviceAccount.IAMMember` or they will fight over what your policy should be.

        > **Note:** `serviceAccount.IAMBinding` resources **can be** used in conjunction with `serviceAccount.IAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\_service\_account\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/iam.serviceAccountUser",
            "members": ["user:jane@example.com"],
        }])
        sa = gcp.service_account.Account("sa",
            account_id="my-service-account",
            display_name="A service account that only Jane can interact with")
        admin_account_iam = gcp.service_account.IAMPolicy("admin-account-iam",
            service_account_id=sa.name,
            policy_data=admin.policy_data)
        ```

        ## google\_service\_account\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        sa = gcp.service_account.Account("sa",
            account_id="my-service-account",
            display_name="A service account that only Jane can use")
        admin_account_iam = gcp.service_account.IAMBinding("admin-account-iam",
            service_account_id=sa.name,
            role="roles/iam.serviceAccountUser",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        sa = gcp.service_account.Account("sa",
            account_id="my-service-account",
            display_name="A service account that only Jane can use")
        admin_account_iam = gcp.service_account.IAMBinding("admin-account-iam",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            members=["user:jane@example.com"],
            role="roles/iam.serviceAccountUser",
            service_account_id=sa.name)
        ```

        ## google\_service\_account\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.get_default_service_account()
        sa = gcp.service_account.Account("sa",
            account_id="my-service-account",
            display_name="A service account that Jane can use")
        admin_account_iam = gcp.service_account.IAMMember("admin-account-iam",
            service_account_id=sa.name,
            role="roles/iam.serviceAccountUser",
            member="user:jane@example.com")
        # Allow SA service account use the default GCE account
        gce_default_account_iam = gcp.service_account.IAMMember("gce-default-account-iam",
            service_account_id=default.name,
            role="roles/iam.serviceAccountUser",
            member=sa.email.apply(lambda email: f"serviceAccount:{email}"))
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        sa = gcp.service_account.Account("sa",
            account_id="my-service-account",
            display_name="A service account that Jane can use")
        admin_account_iam = gcp.service_account.IAMMember("admin-account-iam",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            member="user:jane@example.com",
            role="roles/iam.serviceAccountUser",
            service_account_id=sa.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.

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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['condition'] = condition
            if members is None:
                raise TypeError("Missing required property 'members'")
            __props__['members'] = members
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if service_account_id is None:
                raise TypeError("Missing required property 'service_account_id'")
            __props__['service_account_id'] = service_account_id
            __props__['etag'] = None
        super(IAMBinding, __self__).__init__(
            'gcp:serviceAccount/iAMBinding:IAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, etag=None, members=None, role=None, service_account_id=None):
        """
        Get an existing IAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the service account IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `serviceAccount.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`) - An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
          * `expression` (`pulumi.Input[str]`) - Textual representation of an expression in Common Expression Language syntax.
          * `title` (`pulumi.Input[str]`) - A title for the expression, i.e. a short string describing its purpose.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["members"] = members
        __props__["role"] = role
        __props__["service_account_id"] = service_account_id
        return IAMBinding(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
