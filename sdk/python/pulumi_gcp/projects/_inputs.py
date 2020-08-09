# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'IAMAuditConfigAuditLogConfigArgs',
    'IAMBindingConditionArgs',
    'IAMMemberConditionArgs',
    'OrganizationPolicyBooleanPolicyArgs',
    'OrganizationPolicyListPolicyArgs',
    'OrganizationPolicyListPolicyAllowArgs',
    'OrganizationPolicyListPolicyDenyArgs',
    'OrganizationPolicyRestorePolicyArgs',
]

@pulumi.input_type
class IAMAuditConfigAuditLogConfigArgs:
    def __init__(__self__, *,
                 log_type: pulumi.Input[str],
                 exempted_members: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] log_type: Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        :param pulumi.Input[List[pulumi.Input[str]]] exempted_members: Identities that do not cause logging for this type of permission.  The format is the same as that for `members`.
        """
        pulumi.set(__self__, "logType", log_type)
        pulumi.set(__self__, "exemptedMembers", exempted_members)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Input[str]:
        """
        Permission type for which logging is to be configured.  Must be one of `DATA_READ`, `DATA_WRITE`, or `ADMIN_READ`.
        """
        ...

    @log_type.setter
    def log_type(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Identities that do not cause logging for this type of permission.  The format is the same as that for `members`.
        """
        ...

    @exempted_members.setter
    def exempted_members(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class IAMBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class IAMMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        ...

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        A title for the expression, i.e. a short string describing its purpose.
        """
        ...

    @title.setter
    def title(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        ...

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class OrganizationPolicyBooleanPolicyArgs:
    def __init__(__self__, *,
                 enforced: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] enforced: If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        pulumi.set(__self__, "enforced", enforced)

    @property
    @pulumi.getter
    def enforced(self) -> pulumi.Input[bool]:
        """
        If true, then the Policy is enforced. If false, then any configuration is acceptable.
        """
        ...

    @enforced.setter
    def enforced(self, value: pulumi.Input[bool]):
        ...


@pulumi.input_type
class OrganizationPolicyListPolicyArgs:
    def __init__(__self__, *,
                 allow: Optional[pulumi.Input['OrganizationPolicyListPolicyAllowArgs']] = None,
                 deny: Optional[pulumi.Input['OrganizationPolicyListPolicyDenyArgs']] = None,
                 inherit_from_parent: Optional[pulumi.Input[bool]] = None,
                 suggested_value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input['OrganizationPolicyListPolicyAllowArgs'] allow: or `deny` - (Optional) One or the other must be set.
        :param pulumi.Input[bool] inherit_from_parent: If set to true, the values from the effective Policy of the parent resource
               are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        :param pulumi.Input[str] suggested_value: The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        pulumi.set(__self__, "allow", allow)
        pulumi.set(__self__, "deny", deny)
        pulumi.set(__self__, "inheritFromParent", inherit_from_parent)
        pulumi.set(__self__, "suggestedValue", suggested_value)

    @property
    @pulumi.getter
    def allow(self) -> Optional[pulumi.Input['OrganizationPolicyListPolicyAllowArgs']]:
        """
        or `deny` - (Optional) One or the other must be set.
        """
        ...

    @allow.setter
    def allow(self, value: Optional[pulumi.Input['OrganizationPolicyListPolicyAllowArgs']]):
        ...

    @property
    @pulumi.getter
    def deny(self) -> Optional[pulumi.Input['OrganizationPolicyListPolicyDenyArgs']]:
        ...

    @deny.setter
    def deny(self, value: Optional[pulumi.Input['OrganizationPolicyListPolicyDenyArgs']]):
        ...

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the values from the effective Policy of the parent resource
        are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
        """
        ...

    @inherit_from_parent.setter
    def inherit_from_parent(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter(name="suggestedValue")
    def suggested_value(self) -> Optional[pulumi.Input[str]]:
        """
        The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
        """
        ...

    @suggested_value.setter
    def suggested_value(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class OrganizationPolicyListPolicyAllowArgs:
    def __init__(__self__, *,
                 all: Optional[pulumi.Input[bool]] = None,
                 values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[bool] all: The policy allows or denies all values.
        :param pulumi.Input[List[pulumi.Input[str]]] values: The policy can define specific values that are allowed or denied.
        """
        pulumi.set(__self__, "all", all)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> Optional[pulumi.Input[bool]]:
        """
        The policy allows or denies all values.
        """
        ...

    @all.setter
    def all(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        ...

    @values.setter
    def values(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class OrganizationPolicyListPolicyDenyArgs:
    def __init__(__self__, *,
                 all: Optional[pulumi.Input[bool]] = None,
                 values: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[bool] all: The policy allows or denies all values.
        :param pulumi.Input[List[pulumi.Input[str]]] values: The policy can define specific values that are allowed or denied.
        """
        pulumi.set(__self__, "all", all)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def all(self) -> Optional[pulumi.Input[bool]]:
        """
        The policy allows or denies all values.
        """
        ...

    @all.setter
    def all(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The policy can define specific values that are allowed or denied.
        """
        ...

    @values.setter
    def values(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class OrganizationPolicyRestorePolicyArgs:
    def __init__(__self__, *,
                 default: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] default: May only be set to true. If set, then the default Policy is restored.
        """
        pulumi.set(__self__, "default", default)

    @property
    @pulumi.getter
    def default(self) -> pulumi.Input[bool]:
        """
        May only be set to true. If set, then the default Policy is restored.
        """
        ...

    @default.setter
    def default(self, value: pulumi.Input[bool]):
        ...


