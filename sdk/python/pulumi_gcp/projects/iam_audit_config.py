# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['IAMAuditConfig']


class IAMAuditConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_log_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAMAuditConfigAuditLogConfigArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:

        * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
        * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
        * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
        * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.

        > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.

        > **Note:** It is not possible to grant the `roles/owner` role using any of these resources due to this being disallowed by the underlying `projects.setIamPolicy` API method. See the method [documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy) for full details. It is, however, possible to remove all owners from the project by passing in an empty `members = []` list to the `projects.IAMBinding` resource. This is useful for removing the owner role from a project upon creation, however, precautions should be taken to avoid inadvertently locking oneself out of a project such as by granting additional roles to alternate entities.

        ## google\_project\_iam\_policy

        > **Be careful!** You can accidentally lock yourself out of your project
           using this resource. Deleting a `projects.IAMPolicy` removes access
           from anyone without organization-level access to the project. Proceed with caution.
           It's not recommended to use `projects.IAMPolicy` with your provider project
           to avoid locking yourself out, and it should generally only be used with projects
           fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
           applying the change.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/editor",
            "members": ["user:jane@example.com"],
        }])
        project = gcp.projects.IAMPolicy("project",
            project="your-project-id",
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "condition": {
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            "members": ["user:jane@example.com"],
            "role": "roles/editor",
        }])
        project = gcp.projects.IAMPolicy("project",
            policy_data=admin.policy_data,
            project="your-project-id")
        ```

        ## google\_project\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMBinding("project",
            members=["user:jane@example.com"],
            project="your-project-id",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMBinding("project",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            members=["user:jane@example.com"],
            project="your-project-id",
            role="roles/editor")
        ```

        ## google\_project\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMMember("project",
            member="user:jane@example.com",
            project="your-project-id",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMMember("project",
            condition={
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                "title": "expires_after_2019_12_31",
            },
            member="user:jane@example.com",
            project="your-project-id",
            role="roles/editor")
        ```

        ## google\_project\_iam\_audit\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMAuditConfig("project",
            audit_log_configs=[
                {
                    "logType": "ADMIN_READ",
                },
                {
                    "exemptedMembers": ["user:joebloggs@hashicorp.com"],
                    "logType": "DATA_READ",
                },
            ],
            project="your-project-id",
            service="allServices")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `project_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `project_id` and role, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `project_id`.

        ```sh
         $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project your-project-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMAuditConfig:IAMAuditConfig my_project "your-project-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAMAuditConfigAuditLogConfigArgs']]]] audit_log_configs: The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        :param pulumi.Input[str] service: Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
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

            if audit_log_configs is None:
                raise TypeError("Missing required property 'audit_log_configs'")
            __props__['audit_log_configs'] = audit_log_configs
            __props__['project'] = project
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['etag'] = None
        super(IAMAuditConfig, __self__).__init__(
            'gcp:projects/iAMAuditConfig:IAMAuditConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audit_log_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAMAuditConfigAuditLogConfigArgs']]]]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service: Optional[pulumi.Input[str]] = None) -> 'IAMAuditConfig':
        """
        Get an existing IAMAuditConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAMAuditConfigAuditLogConfigArgs']]]] audit_log_configs: The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        :param pulumi.Input[str] service: Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["audit_log_configs"] = audit_log_configs
        __props__["etag"] = etag
        __props__["project"] = project
        __props__["service"] = service
        return IAMAuditConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="auditLogConfigs")
    def audit_log_configs(self) -> pulumi.Output[Sequence['outputs.IAMAuditConfigAuditLogConfig']]:
        """
        The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
        """
        return pulumi.get(self, "audit_log_configs")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        Required for `projects.IAMPolicy` - you must explicitly set the project, and it
        will not be inferred from the provider.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_project\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
        """
        return pulumi.get(self, "service")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

