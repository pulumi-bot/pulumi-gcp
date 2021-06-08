# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceAttachmentArgs', 'InstanceAttachment']

@pulumi.input_type
class InstanceAttachmentArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceAttachment resource.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] instance_id: The Apigee instance associated with the Apigee environment,
               in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The Apigee instance associated with the Apigee environment,
        in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _InstanceAttachmentState:
    def __init__(__self__, *,
                 environment: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceAttachment resources.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] instance_id: The Apigee instance associated with the Apigee environment,
               in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        :param pulumi.Input[str] name: The name of the newly created attachment (output parameter).
        """
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Apigee instance associated with the Apigee environment,
        in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the newly created attachment (output parameter).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class InstanceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An `Instance attachment` in Apigee.

        To get more information about InstanceAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances.attachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage

        ## Import

        InstanceAttachment can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/attachments/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] instance_id: The Apigee instance associated with the Apigee environment,
               in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: InstanceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An `Instance attachment` in Apigee.

        To get more information about InstanceAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances.attachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage

        ## Import

        InstanceAttachment can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/attachments/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/instanceAttachment:InstanceAttachment default {{instance_id}}/{{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param InstanceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = InstanceAttachmentArgs.__new__(InstanceAttachmentArgs)

            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = None
        super(InstanceAttachment, __self__).__init__(
            'gcp:apigee/instanceAttachment:InstanceAttachment',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            environment: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'InstanceAttachment':
        """
        Get an existing InstanceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] environment: The resource ID of the environment.
        :param pulumi.Input[str] instance_id: The Apigee instance associated with the Apigee environment,
               in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        :param pulumi.Input[str] name: The name of the newly created attachment (output parameter).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceAttachmentState.__new__(_InstanceAttachmentState)

        __props__.__dict__["environment"] = environment
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        return InstanceAttachment(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The resource ID of the environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The Apigee instance associated with the Apigee environment,
        in the format `organisations/{{org_name}}/instances/{{instance_name}}`.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the newly created attachment (output parameter).
        """
        return pulumi.get(self, "name")

