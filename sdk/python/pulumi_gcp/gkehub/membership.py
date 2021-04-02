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

__all__ = ['MembershipArgs', 'Membership']

@pulumi.input_type
class MembershipArgs:
    def __init__(__self__, *,
                 membership_id: pulumi.Input[str],
                 authority: Optional[pulumi.Input['MembershipAuthorityArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input['MembershipEndpointArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Membership resource.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input['MembershipAuthorityArgs'] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console.
        :param pulumi.Input['MembershipEndpointArgs'] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "membership_id", membership_id)
        if authority is not None:
            pulumi.set(__self__, "authority", authority)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Input[str]:
        """
        The client-provided identifier of the membership.
        """
        return pulumi.get(self, "membership_id")

    @membership_id.setter
    def membership_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "membership_id", value)

    @property
    @pulumi.getter
    def authority(self) -> Optional[pulumi.Input['MembershipAuthorityArgs']]:
        """
        Authority encodes how Google will recognize identities from this Membership.
        See the workload identity documentation for more details:
        https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        Structure is documented below.
        """
        return pulumi.get(self, "authority")

    @authority.setter
    def authority(self, value: Optional[pulumi.Input['MembershipAuthorityArgs']]):
        pulumi.set(self, "authority", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input['MembershipEndpointArgs']]:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input['MembershipEndpointArgs']]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to apply to this membership.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Membership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Membership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gkehub/membership:Membership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console.
        :param pulumi.Input[pulumi.InputType['MembershipEndpointArgs']] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Membership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gkehub/membership:Membership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param MembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 membership_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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

            __props__['authority'] = authority
            __props__['description'] = description
            __props__['endpoint'] = endpoint
            __props__['labels'] = labels
            if membership_id is None and not opts.urn:
                raise TypeError("Missing required property 'membership_id'")
            __props__['membership_id'] = membership_id
            __props__['project'] = project
            __props__['name'] = None
        super(Membership, __self__).__init__(
            'gcp:gkehub/membership:Membership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authority: Optional[pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[pulumi.InputType['MembershipEndpointArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            membership_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Membership':
        """
        Get an existing Membership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['MembershipAuthorityArgs']] authority: Authority encodes how Google will recognize identities from this Membership.
               See the workload identity documentation for more details:
               https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
               Structure is documented below.
        :param pulumi.Input[str] description: The name of this entity type to be displayed on the console.
        :param pulumi.Input[pulumi.InputType['MembershipEndpointArgs']] endpoint: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this membership.
        :param pulumi.Input[str] membership_id: The client-provided identifier of the membership.
        :param pulumi.Input[str] name: The unique identifier of the membership.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authority"] = authority
        __props__["description"] = description
        __props__["endpoint"] = endpoint
        __props__["labels"] = labels
        __props__["membership_id"] = membership_id
        __props__["name"] = name
        __props__["project"] = project
        return Membership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authority(self) -> pulumi.Output[Optional['outputs.MembershipAuthority']]:
        """
        Authority encodes how Google will recognize identities from this Membership.
        See the workload identity documentation for more details:
        https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
        Structure is documented below.
        """
        return pulumi.get(self, "authority")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional['outputs.MembershipEndpoint']]:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to apply to this membership.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> pulumi.Output[str]:
        """
        The client-provided identifier of the membership.
        """
        return pulumi.get(self, "membership_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the membership.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

