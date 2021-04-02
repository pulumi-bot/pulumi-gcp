# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVariableResult',
    'AwaitableGetVariableResult',
    'get_variable',
]

@pulumi.output_type
class GetVariableResult:
    """
    A collection of values returned by getVariable.
    """
    def __init__(__self__, id=None, name=None, parent=None, project=None, text=None, update_time=None, value=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent and not isinstance(parent, str):
            raise TypeError("Expected argument 'parent' to be a str")
        pulumi.set(__self__, "parent", parent)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if text and not isinstance(text, str):
            raise TypeError("Expected argument 'text' to be a str")
        pulumi.set(__self__, "text", text)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> str:
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def text(self) -> str:
        return pulumi.get(self, "text")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


class AwaitableGetVariableResult(GetVariableResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVariableResult(
            id=self.id,
            name=self.name,
            parent=self.parent,
            project=self.project,
            text=self.text,
            update_time=self.update_time,
            value=self.value)


def get_variable(name: Optional[str] = None,
                 parent: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVariableResult:
    """
    To get more information about RuntimeConfigs, see:

    * [API documentation](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/v1beta1/projects.configs)
    * How-to Guides
        * [Runtime Configurator Fundamentals](https://cloud.google.com/deployment-manager/runtime-configurator/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    run_service = gcp.runtimeconfig.get_variable(name="prod-variables/hostname",
        parent="my-service")
    ```


    :param str name: The name of the Runtime Configurator configuration.
    :param str parent: The name of the RuntimeConfig resource containing this variable.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['parent'] = parent
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:runtimeconfig/getVariable:getVariable', __args__, opts=opts, typ=GetVariableResult).value

    return AwaitableGetVariableResult(
        id=__ret__.id,
        name=__ret__.name,
        parent=__ret__.parent,
        project=__ret__.project,
        text=__ret__.text,
        update_time=__ret__.update_time,
        value=__ret__.value)
