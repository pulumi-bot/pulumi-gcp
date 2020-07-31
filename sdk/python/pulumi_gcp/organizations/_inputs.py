# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'IAMBindingConditionArgs',
    'IAMMemberConditionArgs',
    'IamAuditConfigAuditLogConfigArgs',
    'PolicyBooleanPolicyArgs',
    'PolicyListPolicyArgs',
    'PolicyListPolicyAllowArgs',
    'PolicyListPolicyDenyArgs',
    'PolicyRestorePolicyArgs',
    'GetIAMPolicyAuditConfigArgs',
    'GetIAMPolicyAuditConfigAuditLogConfigArgs',
    'GetIAMPolicyBindingArgs',
    'GetIAMPolicyBindingConditionArgs',
]

@pulumi.input_type
class IAMBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class IAMMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    title: pulumi.Input[str] = pulumi.input_property("title")
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class IamAuditConfigAuditLogConfigArgs:
    log_type: pulumi.Input[str] = pulumi.input_property("logType")
    """
    Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
    """
    exempted_members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("exemptedMembers")
    """
    Identities that do not cause logging for this type of permission.
    Each entry can have one of the following values:
    * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
    * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
    * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
    * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, log_type: pulumi.Input[str], exempted_members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[str] log_type: Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        :param pulumi.Input[List[pulumi.Input[str]]] exempted_members: Identities that do not cause logging for this type of permission.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        __self__.log_type = log_type
        __self__.exempted_members = exempted_members

@pulumi.input_type
class PolicyBooleanPolicyArgs:
    enforced: pulumi.Input[bool] = pulumi.input_property("enforced")
    """
    If true, then the Policy is enforced. If false, then any configuration is acceptable.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, enforced: pulumi.Input[bool]) -> None:
        """
        :param pulumi.Input[bool] enforced: If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        __self__.enforced = enforced

@pulumi.input_type
class PolicyListPolicyArgs:
    allow: Optional[pulumi.Input['PolicyListPolicyAllowArgs']] = pulumi.input_property("allow")
    """
    or `deny` - (Optional) One or the other must be set.
    """
    deny: Optional[pulumi.Input['PolicyListPolicyDenyArgs']] = pulumi.input_property("deny")
    inherit_from_parent: Optional[pulumi.Input[bool]] = pulumi.input_property("inheritFromParent")
    """
    If set to true, the values from the effective Policy of the parent resource
    are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
    """
    suggested_value: Optional[pulumi.Input[str]] = pulumi.input_property("suggestedValue")
    """
    The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, allow: Optional[pulumi.Input['PolicyListPolicyAllowArgs']] = None, deny: Optional[pulumi.Input['PolicyListPolicyDenyArgs']] = None, inherit_from_parent: Optional[pulumi.Input[bool]] = None, suggested_value: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input['PolicyListPolicyAllowArgs'] allow: or `deny` - (Optional) One or the other must be set.
        :param pulumi.Input[bool] inherit_from_parent: If set to true, the values from the effective Policy of the parent resource
               are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        :param pulumi.Input[str] suggested_value: The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        __self__.allow = allow
        __self__.deny = deny
        __self__.inherit_from_parent = inherit_from_parent
        __self__.suggested_value = suggested_value

@pulumi.input_type
class PolicyListPolicyAllowArgs:
    all: Optional[pulumi.Input[bool]] = pulumi.input_property("all")
    """
    The policy allows or denies all values.
    """
    values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("values")
    """
    The policy can define specific values that are allowed or denied.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, all: Optional[pulumi.Input[bool]] = None, values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[bool] all: The policy allows or denies all values.
        :param pulumi.Input[List[pulumi.Input[str]]] values: The policy can define specific values that are allowed or denied.
        """
        __self__.all = all
        __self__.values = values

@pulumi.input_type
class PolicyListPolicyDenyArgs:
    all: Optional[pulumi.Input[bool]] = pulumi.input_property("all")
    """
    The policy allows or denies all values.
    """
    values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("values")
    """
    The policy can define specific values that are allowed or denied.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, all: Optional[pulumi.Input[bool]] = None, values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[bool] all: The policy allows or denies all values.
        :param pulumi.Input[List[pulumi.Input[str]]] values: The policy can define specific values that are allowed or denied.
        """
        __self__.all = all
        __self__.values = values

@pulumi.input_type
class PolicyRestorePolicyArgs:
    default: pulumi.Input[bool] = pulumi.input_property("default")
    """
    May only be set to true. If set, then the default Policy is restored.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, default: pulumi.Input[bool]) -> None:
        """
        :param pulumi.Input[bool] default: May only be set to true. If set, then the default Policy is restored.
        """
        __self__.default = default

@pulumi.input_type
class GetIAMPolicyAuditConfigArgs:
    audit_log_configs: List['GetIAMPolicyAuditConfigAuditLogConfigArgs'] = pulumi.input_property("auditLogConfigs")
    """
    A nested block that defines the operations you'd like to log.
    """
    service: str = pulumi.input_property("service")
    """
    Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, audit_log_configs: List['GetIAMPolicyAuditConfigAuditLogConfigArgs'], service: str) -> None:
        """
        :param List['GetIAMPolicyAuditConfigAuditLogConfigArgs'] audit_log_configs: A nested block that defines the operations you'd like to log.
        :param str service: Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        __self__.audit_log_configs = audit_log_configs
        __self__.service = service

@pulumi.input_type
class GetIAMPolicyAuditConfigAuditLogConfigArgs:
    log_type: str = pulumi.input_property("logType")
    """
    Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
    """
    exempted_members: Optional[List[str]] = pulumi.input_property("exemptedMembers")
    """
    Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, log_type: str, exempted_members: Optional[List[str]] = None) -> None:
        """
        :param str log_type: Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
        :param List[str] exempted_members: Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
        """
        __self__.log_type = log_type
        __self__.exempted_members = exempted_members

@pulumi.input_type
class GetIAMPolicyBindingArgs:
    members: List[str] = pulumi.input_property("members")
    """
    An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
    Each entry can have one of the following values:
    * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
    * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
    * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
    * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
    * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
    * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
    """
    role: str = pulumi.input_property("role")
    """
    The role/permission that will be granted to the members.
    See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
    Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    condition: Optional['GetIAMPolicyBindingConditionArgs'] = pulumi.input_property("condition")
    """
    An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, members: List[str], role: str, condition: Optional['GetIAMPolicyBindingConditionArgs'] = None) -> None:
        """
        :param List[str] members: An array of identities that will be granted the privilege in the `role`. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account. It **can't** be used with the `organizations.Project` resource.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account. It **can't** be used with the `organizations.Project` resource.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param str role: The role/permission that will be granted to the members.
               See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
               Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param 'GetIAMPolicyBindingConditionArgs' condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
        """
        __self__.members = members
        __self__.role = role
        __self__.condition = condition

@pulumi.input_type
class GetIAMPolicyBindingConditionArgs:
    expression: str = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[str] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: str, title: str, description: Optional[str] = None) -> None:
        """
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str title: A title for the expression, i.e. a short string describing its purpose.
        :param str description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

