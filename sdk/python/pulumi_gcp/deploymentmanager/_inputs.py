# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'DeploymentLabelArgs',
    'DeploymentTargetArgs',
    'DeploymentTargetConfigArgs',
    'DeploymentTargetImportArgs',
]

@pulumi.input_type
class DeploymentLabelArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] key: Key for label.
        :param pulumi.Input[str] value: Value of label.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key for label.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value of label.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DeploymentTargetArgs:
    def __init__(__self__, *,
                 config: pulumi.Input['DeploymentTargetConfigArgs'],
                 imports: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentTargetImportArgs']]]] = None):
        """
        :param pulumi.Input['DeploymentTargetConfigArgs'] config: The root configuration file to use for this deployment.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['DeploymentTargetImportArgs']]] imports: Specifies import files for this configuration. This can be
               used to import templates or other files. For example, you might
               import a text file in order to use the file in a template.
               Structure is documented below.
        """
        pulumi.set(__self__, "config", config)
        if imports is not None:
            pulumi.set(__self__, "imports", imports)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['DeploymentTargetConfigArgs']:
        """
        The root configuration file to use for this deployment.
        Structure is documented below.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['DeploymentTargetConfigArgs']):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def imports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentTargetImportArgs']]]]:
        """
        Specifies import files for this configuration. This can be
        used to import templates or other files. For example, you might
        import a text file in order to use the file in a template.
        Structure is documented below.
        """
        return pulumi.get(self, "imports")

    @imports.setter
    def imports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DeploymentTargetImportArgs']]]]):
        pulumi.set(self, "imports", value)


@pulumi.input_type
class DeploymentTargetConfigArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str]):
        """
        :param pulumi.Input[str] content: The full contents of the template that you want to import.
        """
        pulumi.set(__self__, "content", content)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The full contents of the template that you want to import.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)


@pulumi.input_type
class DeploymentTargetImportArgs:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] content: The full contents of the template that you want to import.
        :param pulumi.Input[str] name: The name of the template to import, as declared in the YAML
               configuration.
        """
        if content is not None:
            pulumi.set(__self__, "content", content)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The full contents of the template that you want to import.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template to import, as declared in the YAML
        configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


