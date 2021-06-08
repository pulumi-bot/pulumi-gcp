# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrganizationSecurityPolicyArgs', 'OrganizationSecurityPolicy']

@pulumi.input_type
class OrganizationSecurityPolicyArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 parent: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationSecurityPolicy resource.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "parent", parent)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        A textual name of the security policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        Format: organizations/{organization_id} or folders/{folder_id}
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A textual description for the organization security policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type indicates the intended use of the security policy.
        For organization security policies, the only supported type
        is "FIREWALL".
        Default value is `FIREWALL`.
        Possible values are `FIREWALL`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _OrganizationSecurityPolicyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationSecurityPolicy resources.
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. This field is used internally during updates of this resource.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] policy_id: The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A textual description for the organization security policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A textual name of the security policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. This field is used internally during updates of this resource.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        Format: organizations/{organization_id} or folders/{folder_id}
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type indicates the intended use of the security policy.
        For organization security policies, the only supported type
        is "FIREWALL".
        Default value is `FIREWALL`.
        Possible values are `FIREWALL`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class OrganizationSecurityPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Organization security policies are used to control incoming/outgoing traffic.

        To get more information about OrganizationSecurityPolicy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies)
        * How-to Guides
            * [Creating a firewall policy](https://cloud.google.com/vpc/docs/using-firewall-policies#create-policy)

        ## Example Usage

        ## Import

        OrganizationSecurityPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default locations/global/securityPolicies/{{policy_id}}
        ```

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default {{policy_id}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: OrganizationSecurityPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Organization security policies are used to control incoming/outgoing traffic.

        To get more information about OrganizationSecurityPolicy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies)
        * How-to Guides
            * [Creating a firewall policy](https://cloud.google.com/vpc/docs/using-firewall-policies#create-policy)

        ## Example Usage

        ## Import

        OrganizationSecurityPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default locations/global/securityPolicies/{{policy_id}}
        ```

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy default {{policy_id}}
        ```

        :param str resource_name_: The name of the resource.
        :param OrganizationSecurityPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSecurityPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationSecurityPolicyArgs.__new__(OrganizationSecurityPolicyArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["type"] = type
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["policy_id"] = None
        super(OrganizationSecurityPolicy, __self__).__init__(
            'gcp:compute/organizationSecurityPolicy:OrganizationSecurityPolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'OrganizationSecurityPolicy':
        """
        Get an existing OrganizationSecurityPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A textual description for the organization security policy.
        :param pulumi.Input[str] display_name: A textual name of the security policy.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. This field is used internally during updates of this resource.
        :param pulumi.Input[str] parent: The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
               Format: organizations/{organization_id} or folders/{folder_id}
        :param pulumi.Input[str] policy_id: The unique identifier for the resource. This identifier is defined by the server.
        :param pulumi.Input[str] type: The type indicates the intended use of the security policy.
               For organization security policies, the only supported type
               is "FIREWALL".
               Default value is `FIREWALL`.
               Possible values are `FIREWALL`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationSecurityPolicyState.__new__(_OrganizationSecurityPolicyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["parent"] = parent
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["type"] = type
        return OrganizationSecurityPolicy(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A textual description for the organization security policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A textual name of the security policy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. This field is used internally during updates of this resource.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The parent of this OrganizationSecurityPolicy in the Cloud Resource Hierarchy.
        Format: organizations/{organization_id} or folders/{folder_id}
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type indicates the intended use of the security policy.
        For organization security policies, the only supported type
        is "FIREWALL".
        Default value is `FIREWALL`.
        Possible values are `FIREWALL`.
        """
        return pulumi.get(self, "type")

