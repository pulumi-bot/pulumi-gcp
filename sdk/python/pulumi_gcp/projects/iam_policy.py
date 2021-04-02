# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IAMPolicyArgs', 'IAMPolicy']

@pulumi.input_type
class IAMPolicyArgs:
    def __init__(__self__, *,
                 policy_data: pulumi.Input[str],
                 project: pulumi.Input[str]):
        """
        The set of arguments for constructing a IAMPolicy resource.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        """
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The `organizations.getIAMPolicy` data source that represents
        the IAM policy that will be applied to the project. The policy will be
        merged with any existing policy applied to the project.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        Required for `projects.IAMPolicy` - you must explicitly set the project, and it
        will not be inferred from the provider.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _IAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The `organizations.getIAMPolicy` data source that represents
        the IAM policy that will be applied to the project. The policy will be
        merged with any existing policy applied to the project.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        Required for `projects.IAMPolicy` - you must explicitly set the project, and it
        will not be inferred from the provider.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class IAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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

        > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.

        > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.

        > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
           IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.

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

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        project = gcp.projects.IAMPolicy("project",
            project="your-project-id",
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/compute.admin",
        )])
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
            condition=gcp.projects.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            project="your-project-id",
            role="roles/container.admin")
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
            condition=gcp.projects.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            project="your-project-id",
            role="roles/firebase.admin")
        ```

        ## google\_project\_iam\_audit\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMAuditConfig("project",
            audit_log_configs=[
                gcp.projects.IAMAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.projects.IAMAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            project="your-project-id",
            service="allServices")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `project_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `project_id` and role, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `project_id`.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project your-project-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:

        * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
        * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
        * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
        * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.

        > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.

        > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.

        > **Note:** The underlying API method `projects.setIamPolicy` has a lot of constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
           IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning 400 error code so please review these if you encounter errors with this resource.

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

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        project = gcp.projects.IAMPolicy("project",
            project="your-project-id",
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/compute.admin",
        )])
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
            condition=gcp.projects.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            project="your-project-id",
            role="roles/container.admin")
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
            condition=gcp.projects.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            project="your-project-id",
            role="roles/firebase.admin")
        ```

        ## google\_project\_iam\_audit\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.IAMAuditConfig("project",
            audit_log_configs=[
                gcp.projects.IAMAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.projects.IAMAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            project="your-project-id",
            service="allServices")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `project_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `project_id` and role, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `project_id`.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project your-project-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMPolicy:IAMPolicy my_project "your-project-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param IAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = IAMPolicyArgs.__new__(IAMPolicyArgs)

            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__['policy_data'] = policy_data
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__['project'] = project
            __props__.__dict__['etag'] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:projects/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'IAMPolicy':
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] policy_data: The `organizations.getIAMPolicy` data source that represents
               the IAM policy that will be applied to the project. The policy will be
               merged with any existing policy applied to the project.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMPolicyState.__new__(_IAMPolicyState)

        __props__.__dict__['etag'] = etag
        __props__.__dict__['policy_data'] = policy_data
        __props__.__dict__['project'] = project
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The `organizations.getIAMPolicy` data source that represents
        the IAM policy that will be applied to the project. The policy will be
        merged with any existing policy applied to the project.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
        Required for `projects.IAMPolicy` - you must explicitly set the project, and it
        will not be inferred from the provider.
        """
        return pulumi.get(self, "project")

