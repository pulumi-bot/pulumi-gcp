# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'IAMBindingCondition',
    'IAMMemberCondition',
    'IamAuditConfigAuditLogConfig',
    'PolicyBooleanPolicy',
    'PolicyListPolicy',
    'PolicyListPolicyAllow',
    'PolicyListPolicyDeny',
    'PolicyRestorePolicy',
    'GetIAMPolicyAuditConfig',
    'GetIAMPolicyAuditConfigAuditLogConfig',
    'GetIAMPolicyBinding',
    'GetIAMPolicyBindingCondition',
]

@pulumi.output_type
class IAMBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IAMMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamAuditConfigAuditLogConfig(dict):
    exempted_members: Optional[List[str]] = pulumi.output_property("exemptedMembers")
    """
    Identities that do not cause logging for this type of permission.
    Each entry can have one of the following values:
    * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
    * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
    * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
    * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
    """
    log_type: str = pulumi.output_property("logType")
    """
    Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyBooleanPolicy(dict):
    enforced: bool = pulumi.output_property("enforced")
    """
    If true, then the Policy is enforced. If false, then any configuration is acceptable.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicy(dict):
    allow: Optional['outputs.PolicyListPolicyAllow'] = pulumi.output_property("allow")
    """
    or `deny` - (Optional) One or the other must be set.
    """
    deny: Optional['outputs.PolicyListPolicyDeny'] = pulumi.output_property("deny")
    inherit_from_parent: Optional[bool] = pulumi.output_property("inheritFromParent")
    """
    If set to true, the values from the effective Policy of the parent resource
    are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
    """
    suggested_value: Optional[str] = pulumi.output_property("suggestedValue")
    """
    The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicyAllow(dict):
    all: Optional[bool] = pulumi.output_property("all")
    """
    The policy allows or denies all values.
    """
    values: Optional[List[str]] = pulumi.output_property("values")
    """
    The policy can define specific values that are allowed or denied.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyListPolicyDeny(dict):
    all: Optional[bool] = pulumi.output_property("all")
    """
    The policy allows or denies all values.
    """
    values: Optional[List[str]] = pulumi.output_property("values")
    """
    The policy can define specific values that are allowed or denied.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyRestorePolicy(dict):
    default: bool = pulumi.output_property("default")
    """
    May only be set to true. If set, then the default Policy is restored.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetIAMPolicyAuditConfig(dict):
    audit_log_configs: List['outputs.GetIAMPolicyAuditConfigAuditLogConfig'] = pulumi.output_property("auditLogConfigs")
    """
    A nested block that defines the operations you'd like to log.
    """
    service: str = pulumi.output_property("service")
    """
    Defines a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetIAMPolicyAuditConfigAuditLogConfig(dict):
    exempted_members: Optional[List[str]] = pulumi.output_property("exemptedMembers")
    """
    Specifies the identities that are exempt from these types of logging operations. Follows the same format of the `members` array for `binding`.
    """
    log_type: str = pulumi.output_property("logType")
    """
    Defines the logging level. `DATA_READ`, `DATA_WRITE` and `ADMIN_READ` capture different types of events. See [the audit configuration documentation](https://cloud.google.com/resource-manager/reference/rest/Shared.Types/AuditConfig) for more details.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetIAMPolicyBinding(dict):
    condition: Optional['outputs.GetIAMPolicyBindingCondition'] = pulumi.output_property("condition")
    """
    An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding. Structure is documented below.
    """
    members: List[str] = pulumi.output_property("members")
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
    role: str = pulumi.output_property("role")
    """
    The role/permission that will be granted to the members.
    See the [IAM Roles](https://cloud.google.com/compute/docs/access/iam) documentation for a complete list of roles.
    Note that custom roles must be of the format `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetIAMPolicyBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """
    expression: str = pulumi.output_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: str = pulumi.output_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


