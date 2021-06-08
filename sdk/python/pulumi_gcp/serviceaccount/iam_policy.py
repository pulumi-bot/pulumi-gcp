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
                 service_account_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a IAMPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.
        """
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "service_account_id", service_account_id)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Input[str]:
        """
        The fully-qualified name of the service account to apply policy to.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_id", value)


@pulumi.input_type
class _IAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the service account IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the service account IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The fully-qualified name of the service account to apply policy to.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)


class IAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the google_project_iam set of resources.

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

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iam.serviceAccountUser",
            members=["user:jane@example.com"],
        )])
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
            condition=gcp.service.account.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
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
            condition=gcp.service.account.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            role="roles/iam.serviceAccountUser",
            service_account_id=sa.name)
        ```

        ## Import

        Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser"
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: IAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the google_project_iam set of resources.

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

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iam.serviceAccountUser",
            members=["user:jane@example.com"],
        )])
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
            condition=gcp.service.account.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
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
            condition=gcp.service.account.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            role="roles/iam.serviceAccountUser",
            service_account_id=sa.name)
        ```

        ## Import

        Service account IAM resources can be imported using the project, service account email, role, member identity, and condition (beta).

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email}
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser"
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. With conditions

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
        ```

        ```sh
         $ pulumi import gcp:serviceAccount/iAMPolicy:IAMPolicy admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
        ```

        :param str resource_name_: The name of the resource.
        :param IAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
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
            __props__.__dict__["policy_data"] = policy_data
            if service_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_id'")
            __props__.__dict__["service_account_id"] = service_account_id
            __props__.__dict__["etag"] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:serviceAccount/iAMPolicy:IAMPolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None) -> 'IAMPolicy':
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the service account IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] service_account_id: The fully-qualified name of the service account to apply policy to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMPolicyState.__new__(_IAMPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["service_account_id"] = service_account_id
        return IAMPolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the service account IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[str]:
        """
        The fully-qualified name of the service account to apply policy to.
        """
        return pulumi.get(self, "service_account_id")

