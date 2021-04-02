# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['OrganizationSecurityPolicyAssociationArgs', 'OrganizationSecurityPolicyAssociation']

@pulumi.input_type
class OrganizationSecurityPolicyAssociationArgs:
    def __init__(__self__, *,
                 attachment_id: pulumi.Input[str],
                 policy_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationSecurityPolicyAssociation resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        """
        pulumi.set(__self__, "attachment_id", attachment_id)
        pulumi.set(__self__, "policy_id", policy_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Input[str]:
        """
        The resource that the security policy is attached to.
        """
        return pulumi.get(self, "attachment_id")

    @attachment_id.setter
    def attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "attachment_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The security policy ID of the association.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class OrganizationSecurityPolicyAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An association for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyAssociation, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
        * How-to Guides
            * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)

        ## Example Usage

        ## Import

        OrganizationSecurityPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/association/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationSecurityPolicyAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An association for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyAssociation, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
        * How-to Guides
            * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)

        ## Example Usage

        ## Import

        OrganizationSecurityPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/association/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationSecurityPolicyAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSecurityPolicyAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
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

            if attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'attachment_id'")
            __props__['attachment_id'] = attachment_id
            __props__['name'] = name
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__['policy_id'] = policy_id
            __props__['display_name'] = None
        super(OrganizationSecurityPolicyAssociation, __self__).__init__(
            'gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachment_id: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None) -> 'OrganizationSecurityPolicyAssociation':
        """
        Get an existing OrganizationSecurityPolicyAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] display_name: The display name of the security policy of the association.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attachment_id"] = attachment_id
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["policy_id"] = policy_id
        return OrganizationSecurityPolicyAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Output[str]:
        """
        The resource that the security policy is attached to.
        """
        return pulumi.get(self, "attachment_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the security policy of the association.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The security policy ID of the association.
        """
        return pulumi.get(self, "policy_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

