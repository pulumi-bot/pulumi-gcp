# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetWorkloadIdentityPoolResult',
    'AwaitableGetWorkloadIdentityPoolResult',
    'get_workload_identity_pool',
]

@pulumi.output_type
class GetWorkloadIdentityPoolResult:
    """
    A collection of values returned by getWorkloadIdentityPool.
    """
    def __init__(__self__, description=None, disabled=None, display_name=None, id=None, name=None, project=None, state=None, workload_identity_pool_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if workload_identity_pool_id and not isinstance(workload_identity_pool_id, str):
            raise TypeError("Expected argument 'workload_identity_pool_id' to be a str")
        pulumi.set(__self__, "workload_identity_pool_id", workload_identity_pool_id)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

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
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> str:
        return pulumi.get(self, "workload_identity_pool_id")


class AwaitableGetWorkloadIdentityPoolResult(GetWorkloadIdentityPoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkloadIdentityPoolResult(
            description=self.description,
            disabled=self.disabled,
            display_name=self.display_name,
            id=self.id,
            name=self.name,
            project=self.project,
            state=self.state,
            workload_identity_pool_id=self.workload_identity_pool_id)


def get_workload_identity_pool(project: Optional[str] = None,
                               workload_identity_pool_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkloadIdentityPoolResult:
    """
    Get a IAM workload identity pool from Google Cloud by its id.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    foo = gcp.iam.get_workload_identity_pool(workload_identity_pool_id="foo-pool")
    ```


    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str workload_identity_pool_id: The id of the pool which is the
           final component of the resource name.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['workloadIdentityPoolId'] = workload_identity_pool_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:iam/getWorkloadIdentityPool:getWorkloadIdentityPool', __args__, opts=opts, typ=GetWorkloadIdentityPoolResult).value

    return AwaitableGetWorkloadIdentityPoolResult(
        description=__ret__.description,
        disabled=__ret__.disabled,
        display_name=__ret__.display_name,
        id=__ret__.id,
        name=__ret__.name,
        project=__ret__.project,
        state=__ret__.state,
        workload_identity_pool_id=__ret__.workload_identity_pool_id)


@_utilities.lift_output_func(get_workload_identity_pool)
def get_workload_identity_pool_apply(project: Optional[pulumi.Input[str]] = None,
                                     workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkloadIdentityPoolResult]:
    ...
