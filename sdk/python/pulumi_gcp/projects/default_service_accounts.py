# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['DefaultServiceAccounts']


class DefaultServiceAccounts(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 restore_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows management of Google Cloud Platform project default service accounts.

        When certain service APIs are enabled, Google Cloud Platform automatically creates service accounts to help get started, but
        this is not recommended for production environments as per [Google's documentation](https://cloud.google.com/iam/docs/service-accounts#default).
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
        > This resource works on a best-effort basis, as no API formally describes the default service accounts. If the default service accounts change their name or additional service accounts are added, this resource will need to be updated.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.projects.DefaultServiceAccounts("myProject",
            action="DELETE",
            project="my-project-id")
        ```

        To enable the default service accounts on the resource destroy:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.projects.DefaultServiceAccounts("myProject",
            action="DISABLE",
            project="my-project-id",
            restore_policy="REVERT")
        ```

        ## Import

        This resource does not support import

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
        :param pulumi.Input[str] project: The project ID where service accounts are created.
        :param pulumi.Input[str] restore_policy: The action to be performed in the default service accounts on the resource destroy. Valid values are `NONE` and `REVERT`. If set to `REVERT` it will attempt to restore all default SAs but in the `DEPRIVILEGE` action.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['restore_policy'] = restore_policy
            __props__['service_accounts'] = None
        super(DefaultServiceAccounts, __self__).__init__(
            'gcp:projects/defaultServiceAccounts:DefaultServiceAccounts',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            restore_policy: Optional[pulumi.Input[str]] = None,
            service_accounts: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'DefaultServiceAccounts':
        """
        Get an existing DefaultServiceAccounts resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
        :param pulumi.Input[str] project: The project ID where service accounts are created.
        :param pulumi.Input[str] restore_policy: The action to be performed in the default service accounts on the resource destroy. Valid values are `NONE` and `REVERT`. If set to `REVERT` it will attempt to restore all default SAs but in the `DEPRIVILEGE` action.
        :param pulumi.Input[Mapping[str, Any]] service_accounts: The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["project"] = project
        __props__["restore_policy"] = restore_policy
        __props__["service_accounts"] = service_accounts
        return DefaultServiceAccounts(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID where service accounts are created.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="restorePolicy")
    def restore_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The action to be performed in the default service accounts on the resource destroy. Valid values are `NONE` and `REVERT`. If set to `REVERT` it will attempt to restore all default SAs but in the `DEPRIVILEGE` action.
        """
        return pulumi.get(self, "restore_policy")

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
        """
        return pulumi.get(self, "service_accounts")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

