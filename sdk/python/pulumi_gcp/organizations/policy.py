# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 constraint: pulumi.Input[str],
                 org_id: pulumi.Input[str],
                 boolean_policy: Optional[pulumi.Input['PolicyBooleanPolicyArgs']] = None,
                 list_policy: Optional[pulumi.Input['PolicyListPolicyArgs']] = None,
                 restore_policy: Optional[pulumi.Input['PolicyRestorePolicyArgs']] = None,
                 version: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] org_id: The numeric ID of the organization to set the policy for.
        :param pulumi.Input['PolicyBooleanPolicyArgs'] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input['PolicyListPolicyArgs'] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input['PolicyRestorePolicyArgs'] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[int] version: Version of the Policy. Default version is 0.
        """
        pulumi.set(__self__, "constraint", constraint)
        pulumi.set(__self__, "org_id", org_id)
        if boolean_policy is not None:
            pulumi.set(__self__, "boolean_policy", boolean_policy)
        if list_policy is not None:
            pulumi.set(__self__, "list_policy", list_policy)
        if restore_policy is not None:
            pulumi.set(__self__, "restore_policy", restore_policy)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def constraint(self) -> pulumi.Input[str]:
        """
        The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        """
        return pulumi.get(self, "constraint")

    @constraint.setter
    def constraint(self, value: pulumi.Input[str]):
        pulumi.set(self, "constraint", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        The numeric ID of the organization to set the policy for.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="booleanPolicy")
    def boolean_policy(self) -> Optional[pulumi.Input['PolicyBooleanPolicyArgs']]:
        """
        A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        """
        return pulumi.get(self, "boolean_policy")

    @boolean_policy.setter
    def boolean_policy(self, value: Optional[pulumi.Input['PolicyBooleanPolicyArgs']]):
        pulumi.set(self, "boolean_policy", value)

    @property
    @pulumi.getter(name="listPolicy")
    def list_policy(self) -> Optional[pulumi.Input['PolicyListPolicyArgs']]:
        """
        A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        """
        return pulumi.get(self, "list_policy")

    @list_policy.setter
    def list_policy(self, value: Optional[pulumi.Input['PolicyListPolicyArgs']]):
        pulumi.set(self, "list_policy", value)

    @property
    @pulumi.getter(name="restorePolicy")
    def restore_policy(self) -> Optional[pulumi.Input['PolicyRestorePolicyArgs']]:
        """
        A restore policy is a constraint to restore the default policy. Structure is documented below.
        """
        return pulumi.get(self, "restore_policy")

    @restore_policy.setter
    def restore_policy(self, value: Optional[pulumi.Input['PolicyRestorePolicyArgs']]):
        pulumi.set(self, "restore_policy", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[int]]:
        """
        Version of the Policy. Default version is 0.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "version", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boolean_policy: Optional[pulumi.Input[pulumi.InputType['PolicyBooleanPolicyArgs']]] = None,
                 constraint: Optional[pulumi.Input[str]] = None,
                 list_policy: Optional[pulumi.Input[pulumi.InputType['PolicyListPolicyArgs']]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 restore_policy: Optional[pulumi.Input[pulumi.InputType['PolicyRestorePolicyArgs']]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows management of Organization policies for a Google Organization. For more information see
        [the official
        documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
        [API](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/setOrgPolicy).

        ## Example Usage

        To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        serial_port_policy = gcp.organizations.Policy("serialPortPolicy",
            boolean_policy=gcp.organizations.PolicyBooleanPolicyArgs(
                enforced=True,
            ),
            constraint="compute.disableSerialPortAccess",
            org_id="123456789")
        ```

        To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            list_policy=gcp.organizations.PolicyListPolicyArgs(
                allow=gcp.organizations.PolicyListPolicyAllowArgs(
                    all=True,
                ),
            ),
            org_id="123456789")
        ```

        Or to deny some services, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            list_policy=gcp.organizations.PolicyListPolicyArgs(
                deny=gcp.organizations.PolicyListPolicyDenyArgs(
                    values=["cloudresourcemanager.googleapis.com"],
                ),
                suggested_value="compute.googleapis.com",
            ),
            org_id="123456789")
        ```

        To restore the default organization policy, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            org_id="123456789",
            restore_policy=gcp.organizations.PolicyRestorePolicyArgs(
                default=True,
            ))
        ```

        ## Import

        Organization Policies can be imported using the `org_id` and the `constraint`, e.g.

        ```sh
         $ pulumi import gcp:organizations/policy:Policy services_policy 123456789/constraints/serviceuser.services
        ```

         It is all right if the constraint contains a slash, as in the example above.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyBooleanPolicyArgs']] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[pulumi.InputType['PolicyListPolicyArgs']] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to set the policy for.
        :param pulumi.Input[pulumi.InputType['PolicyRestorePolicyArgs']] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[int] version: Version of the Policy. Default version is 0.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows management of Organization policies for a Google Organization. For more information see
        [the official
        documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
        [API](https://cloud.google.com/resource-manager/reference/rest/v1/organizations/setOrgPolicy).

        ## Example Usage

        To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        serial_port_policy = gcp.organizations.Policy("serialPortPolicy",
            boolean_policy=gcp.organizations.PolicyBooleanPolicyArgs(
                enforced=True,
            ),
            constraint="compute.disableSerialPortAccess",
            org_id="123456789")
        ```

        To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            list_policy=gcp.organizations.PolicyListPolicyArgs(
                allow=gcp.organizations.PolicyListPolicyAllowArgs(
                    all=True,
                ),
            ),
            org_id="123456789")
        ```

        Or to deny some services, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            list_policy=gcp.organizations.PolicyListPolicyArgs(
                deny=gcp.organizations.PolicyListPolicyDenyArgs(
                    values=["cloudresourcemanager.googleapis.com"],
                ),
                suggested_value="compute.googleapis.com",
            ),
            org_id="123456789")
        ```

        To restore the default organization policy, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.organizations.Policy("servicesPolicy",
            constraint="serviceuser.services",
            org_id="123456789",
            restore_policy=gcp.organizations.PolicyRestorePolicyArgs(
                default=True,
            ))
        ```

        ## Import

        Organization Policies can be imported using the `org_id` and the `constraint`, e.g.

        ```sh
         $ pulumi import gcp:organizations/policy:Policy services_policy 123456789/constraints/serviceuser.services
        ```

         It is all right if the constraint contains a slash, as in the example above.

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boolean_policy: Optional[pulumi.Input[pulumi.InputType['PolicyBooleanPolicyArgs']]] = None,
                 constraint: Optional[pulumi.Input[str]] = None,
                 list_policy: Optional[pulumi.Input[pulumi.InputType['PolicyListPolicyArgs']]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 restore_policy: Optional[pulumi.Input[pulumi.InputType['PolicyRestorePolicyArgs']]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['boolean_policy'] = boolean_policy
            if constraint is None and not opts.urn:
                raise TypeError("Missing required property 'constraint'")
            __props__['constraint'] = constraint
            __props__['list_policy'] = list_policy
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__['org_id'] = org_id
            __props__['restore_policy'] = restore_policy
            __props__['version'] = version
            __props__['etag'] = None
            __props__['update_time'] = None
        super(Policy, __self__).__init__(
            'gcp:organizations/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            boolean_policy: Optional[pulumi.Input[pulumi.InputType['PolicyBooleanPolicyArgs']]] = None,
            constraint: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            list_policy: Optional[pulumi.Input[pulumi.InputType['PolicyListPolicyArgs']]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            restore_policy: Optional[pulumi.Input[pulumi.InputType['PolicyRestorePolicyArgs']]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PolicyBooleanPolicyArgs']] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] etag: (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        :param pulumi.Input[pulumi.InputType['PolicyListPolicyArgs']] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[str] org_id: The numeric ID of the organization to set the policy for.
        :param pulumi.Input[pulumi.InputType['PolicyRestorePolicyArgs']] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[str] update_time: (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        :param pulumi.Input[int] version: Version of the Policy. Default version is 0.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["boolean_policy"] = boolean_policy
        __props__["constraint"] = constraint
        __props__["etag"] = etag
        __props__["list_policy"] = list_policy
        __props__["org_id"] = org_id
        __props__["restore_policy"] = restore_policy
        __props__["update_time"] = update_time
        __props__["version"] = version
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="booleanPolicy")
    def boolean_policy(self) -> pulumi.Output[Optional['outputs.PolicyBooleanPolicy']]:
        """
        A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        """
        return pulumi.get(self, "boolean_policy")

    @property
    @pulumi.getter
    def constraint(self) -> pulumi.Output[str]:
        """
        The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        """
        return pulumi.get(self, "constraint")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="listPolicy")
    def list_policy(self) -> pulumi.Output[Optional['outputs.PolicyListPolicy']]:
        """
        A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
        """
        return pulumi.get(self, "list_policy")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The numeric ID of the organization to set the policy for.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="restorePolicy")
    def restore_policy(self) -> pulumi.Output[Optional['outputs.PolicyRestorePolicy']]:
        """
        A restore policy is a constraint to restore the default policy. Structure is documented below.
        """
        return pulumi.get(self, "restore_policy")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Version of the Policy. Default version is 0.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

