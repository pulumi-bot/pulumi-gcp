# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetIAMPolicyResult',
    'AwaitableGetIAMPolicyResult',
    'get_iam_policy',
]



@pulumi.output_type
class GetIAMPolicyResult:
    """
    A collection of values returned by getIAMPolicy.
    """
    def __init__(__self__, audit_configs=None, bindings=None, id=None, policy_data=None):
        if audit_configs and not isinstance(audit_configs, list):
            raise TypeError("Expected argument 'audit_configs' to be a list")
        pulumi.set(__self__, "audit_configs", audit_configs)
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        pulumi.set(__self__, "bindings", bindings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="auditConfigs")
    def audit_configs(self) -> Optional[List['outputs.GetIAMPolicyAuditConfigResult']]:
        return pulumi.get(self, "audit_configs")

    @property
    @pulumi.getter
    def bindings(self) -> Optional[List['outputs.GetIAMPolicyBindingResult']]:
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        The above bindings serialized in a format suitable for
        referencing from a resource that supports IAM.
        """
        return pulumi.get(self, "policy_data")



class AwaitableGetIAMPolicyResult(GetIAMPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIAMPolicyResult(
            audit_configs=self.audit_configs,
            bindings=self.bindings,
            id=self.id,
            policy_data=self.policy_data)


def get_iam_policy(audit_configs: Optional[List[pulumi.InputType['GetIAMPolicyAuditConfigArgs']]] = None,
                   bindings: Optional[List[pulumi.InputType['GetIAMPolicyBindingArgs']]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIAMPolicyResult:
    """
    Generates an IAM policy document that may be referenced by and applied to
    other Google Cloud Platform resources, such as the `organizations.Project` resource.

    **Note:** Several restrictions apply when setting IAM policies through this API.
    See the [setIamPolicy docs](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy)
    for a list of these restrictions.

    This data source is used to define IAM policies to apply to other resources.
    Currently, defining a policy through a datasource and referencing that policy
    from another resource is the only way to apply an IAM policy to a resource.


    :param List[pulumi.InputType['GetIAMPolicyAuditConfigArgs']] audit_configs: A nested configuration block that defines logging additional configuration for your project.
    :param List[pulumi.InputType['GetIAMPolicyBindingArgs']] bindings: A nested configuration block (described below)
           defining a binding to be included in the policy document. Multiple
           `binding` arguments are supported.
    """
    __args__ = dict()
    __args__['auditConfigs'] = audit_configs
    __args__['bindings'] = bindings
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getIAMPolicy:getIAMPolicy', __args__, opts=opts, typ=GetIAMPolicyResult).value

    return AwaitableGetIAMPolicyResult(
        audit_configs=__ret__.audit_configs,
        bindings=__ret__.bindings,
        id=__ret__.id,
        policy_data=__ret__.policy_data)
