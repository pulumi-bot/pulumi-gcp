# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs

@pulumi.output_type
class AppEngineServiceIamBindingCondition(dict):
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


@pulumi.output_type
class AppEngineServiceIamMemberCondition(dict):
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


@pulumi.output_type
class AppEngineVersionIamBindingCondition(dict):
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


@pulumi.output_type
class AppEngineVersionIamMemberCondition(dict):
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


@pulumi.output_type
class TunnelInstanceIAMBindingCondition(dict):
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


@pulumi.output_type
class TunnelInstanceIAMMemberCondition(dict):
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


@pulumi.output_type
class WebBackendServiceIamBindingCondition(dict):
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


@pulumi.output_type
class WebBackendServiceIamMemberCondition(dict):
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


@pulumi.output_type
class WebIamBindingCondition(dict):
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


@pulumi.output_type
class WebIamMemberCondition(dict):
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


@pulumi.output_type
class WebTypeAppEngingIamBindingCondition(dict):
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


@pulumi.output_type
class WebTypeAppEngingIamMemberCondition(dict):
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


@pulumi.output_type
class WebTypeComputeIamBindingCondition(dict):
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


@pulumi.output_type
class WebTypeComputeIamMemberCondition(dict):
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


