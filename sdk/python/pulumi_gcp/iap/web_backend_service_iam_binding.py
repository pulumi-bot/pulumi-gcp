# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['WebBackendServiceIamBindingArgs', 'WebBackendServiceIamBinding']

@pulumi.input_type
class WebBackendServiceIamBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 web_backend_service: pulumi.Input[str],
                 condition: Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebBackendServiceIamBinding resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] web_backend_service: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input['WebBackendServiceIamBindingConditionArgs'] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "web_backend_service", web_backend_service)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="webBackendService")
    def web_backend_service(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "web_backend_service")

    @web_backend_service.setter
    def web_backend_service(self, value: pulumi.Input[str]):
        pulumi.set(self, "web_backend_service", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _WebBackendServiceIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 web_backend_service: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebBackendServiceIamBinding resources.
        :param pulumi.Input['WebBackendServiceIamBindingConditionArgs'] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] web_backend_service: Used to find the parent resource to bind the IAM policy to
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if web_backend_service is not None:
            pulumi.set(__self__, "web_backend_service", web_backend_service)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['WebBackendServiceIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="webBackendService")
    def web_backend_service(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "web_backend_service")

    @web_backend_service.setter
    def web_backend_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_backend_service", value)


class WebBackendServiceIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['WebBackendServiceIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 web_backend_service: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:

        * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
        * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
        * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.

        > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_iap\_web\_backend\_service\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iap.WebBackendServiceIamPolicy("policy",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ),
        )])
        policy = gcp.iap.WebBackendServiceIamPolicy("policy",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\_iap\_web\_backend\_service\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebBackendServiceIamBinding("binding",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebBackendServiceIamBinding("binding",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.iap.WebBackendServiceIamBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```
        ## google\_iap\_web\_backend\_service\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebBackendServiceIamMember("member",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebBackendServiceIamMember("member",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com",
            condition=gcp.iap.WebBackendServiceIamMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WebBackendServiceIamBindingConditionArgs']] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] web_backend_service: Used to find the parent resource to bind the IAM policy to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: WebBackendServiceIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:

        * `iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
        * `iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
        * `iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.

        > **Note:** `iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `iap.WebBackendServiceIamBinding` and `iap.WebBackendServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_iap\_web\_backend\_service\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iap.WebBackendServiceIamPolicy("policy",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ),
        )])
        policy = gcp.iap.WebBackendServiceIamPolicy("policy",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\_iap\_web\_backend\_service\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebBackendServiceIamBinding("binding",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebBackendServiceIamBinding("binding",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.iap.WebBackendServiceIamBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```
        ## google\_iap\_web\_backend\_service\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebBackendServiceIamMember("member",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebBackendServiceIamMember("member",
            project=google_compute_backend_service["default"]["project"],
            web_backend_service=google_compute_backend_service["default"]["name"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com",
            condition=gcp.iap.WebBackendServiceIamMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor "projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name_: The name of the resource.
        :param WebBackendServiceIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebBackendServiceIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['WebBackendServiceIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 web_backend_service: Optional[pulumi.Input[str]] = None,
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
            __props__ = WebBackendServiceIamBindingArgs.__new__(WebBackendServiceIamBindingArgs)

            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if web_backend_service is None and not opts.urn:
                raise TypeError("Missing required property 'web_backend_service'")
            __props__.__dict__["web_backend_service"] = web_backend_service
            __props__.__dict__["etag"] = None
        super(WebBackendServiceIamBinding, __self__).__init__(
            'gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['WebBackendServiceIamBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            web_backend_service: Optional[pulumi.Input[str]] = None) -> 'WebBackendServiceIamBinding':
        """
        Get an existing WebBackendServiceIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WebBackendServiceIamBindingConditionArgs']] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] web_backend_service: Used to find the parent resource to bind the IAM policy to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebBackendServiceIamBindingState.__new__(_WebBackendServiceIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        __props__.__dict__["web_backend_service"] = web_backend_service
        return WebBackendServiceIamBinding(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.WebBackendServiceIamBindingCondition']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="webBackendService")
    def web_backend_service(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "web_backend_service")

