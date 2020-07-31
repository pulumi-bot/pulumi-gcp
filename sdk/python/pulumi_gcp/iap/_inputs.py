# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AppEngineServiceIamBindingConditionArgs',
    'AppEngineServiceIamMemberConditionArgs',
    'AppEngineVersionIamBindingConditionArgs',
    'AppEngineVersionIamMemberConditionArgs',
    'TunnelInstanceIAMBindingConditionArgs',
    'TunnelInstanceIAMMemberConditionArgs',
    'WebBackendServiceIamBindingConditionArgs',
    'WebBackendServiceIamMemberConditionArgs',
    'WebIamBindingConditionArgs',
    'WebIamMemberConditionArgs',
    'WebTypeAppEngingIamBindingConditionArgs',
    'WebTypeAppEngingIamMemberConditionArgs',
    'WebTypeComputeIamBindingConditionArgs',
    'WebTypeComputeIamMemberConditionArgs',
]

@pulumi.input_type
class AppEngineServiceIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class AppEngineServiceIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class AppEngineVersionIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class AppEngineVersionIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class TunnelInstanceIAMBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class TunnelInstanceIAMMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebBackendServiceIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebBackendServiceIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebTypeAppEngingIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebTypeAppEngingIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebTypeComputeIamBindingConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

@pulumi.input_type
class WebTypeComputeIamMemberConditionArgs:
    expression: pulumi.Input[str] = pulumi.input_property("expression")
    """
    Textual representation of an expression in Common Expression Language syntax.
    """
    title: pulumi.Input[str] = pulumi.input_property("title")
    """
    A title for the expression, i.e. a short string describing its purpose.
    """
    description: Optional[pulumi.Input[str]] = pulumi.input_property("description")
    """
    An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, expression: pulumi.Input[str], title: pulumi.Input[str], description: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] title: A title for the expression, i.e. a short string describing its purpose.
        :param pulumi.Input[str] description: An optional description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        __self__.expression = expression
        __self__.title = title
        __self__.description = description

