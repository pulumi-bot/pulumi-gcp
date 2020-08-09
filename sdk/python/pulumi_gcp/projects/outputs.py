# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'IAMAuditConfigAuditLogConfig',
    'IAMBindingCondition',
    'IAMMemberCondition',
    'OrganizationPolicyBooleanPolicy',
    'OrganizationPolicyListPolicy',
    'OrganizationPolicyListPolicyAllow',
    'OrganizationPolicyListPolicyDeny',
    'OrganizationPolicyRestorePolicy',
    'GetOrganizationPolicyBooleanPolicy',
    'GetOrganizationPolicyListPolicy',
    'GetOrganizationPolicyListPolicyAllow',
    'GetOrganizationPolicyListPolicyDeny',
    'GetOrganizationPolicyRestorePolicy',
    'GetProjectProject',
]

@pulumi.output_type
class IAMAuditConfigAuditLogConfig(dict):
    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[List[str]]:
        """
        Identities that do not cause logging for this type of permission.  The format is the same as that for `members`.
        """
        ...

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        """
        Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IAMBindingCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IAMMemberCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyBooleanPolicy(dict):
    @property
    @pulumi.getter
    def enforced(self) -> bool:
        """
        If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyListPolicy(dict):
    @property
    @pulumi.getter
    def allow(self) -> Optional['outputs.OrganizationPolicyListPolicyAllow']:
        """
        or `deny` - (Optional) One or the other must be set.
        """
        ...

    @property
    @pulumi.getter
    def deny(self) -> Optional['outputs.OrganizationPolicyListPolicyDeny']:
        ...

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> Optional[bool]:
        """
        If set to true, the values from the effective Policy of the parent resource
        are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        """
        ...

    @property
    @pulumi.getter(name="suggestedValue")
    def suggested_value(self) -> Optional[str]:
        """
        The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyListPolicyAllow(dict):
    @property
    @pulumi.getter
    def all(self) -> Optional[bool]:
        """
        The policy allows or denies all values.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> Optional[List[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyListPolicyDeny(dict):
    @property
    @pulumi.getter
    def all(self) -> Optional[bool]:
        """
        The policy allows or denies all values.
        """
        ...

    @property
    @pulumi.getter
    def values(self) -> Optional[List[str]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OrganizationPolicyRestorePolicy(dict):
    @property
    @pulumi.getter
    def default(self) -> bool:
        """
        May only be set to true. If set, then the default Policy is restored.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetOrganizationPolicyBooleanPolicy(dict):
    @property
    @pulumi.getter
    def enforced(self) -> bool:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetOrganizationPolicyListPolicy(dict):
    @property
    @pulumi.getter
    def allows(self) -> List['outputs.GetOrganizationPolicyListPolicyAllow']:
        ...

    @property
    @pulumi.getter
    def denies(self) -> List['outputs.GetOrganizationPolicyListPolicyDeny']:
        ...

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> bool:
        ...

    @property
    @pulumi.getter(name="suggestedValue")
    def suggested_value(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetOrganizationPolicyListPolicyAllow(dict):
    @property
    @pulumi.getter
    def all(self) -> bool:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetOrganizationPolicyListPolicyDeny(dict):
    @property
    @pulumi.getter
    def all(self) -> bool:
        ...

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetOrganizationPolicyRestorePolicy(dict):
    @property
    @pulumi.getter
    def default(self) -> bool:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetProjectProject(dict):
    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project id of the project.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


