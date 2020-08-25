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

__all__ = ['Node']


class Node(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accelerator_type: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scheduling_config: Optional[pulumi.Input[pulumi.InputType['NodeSchedulingConfigArgs']]] = None,
                 tensorflow_version: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Cloud TPU instance.

        To get more information about Node, see:

        * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/tpu/docs/)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_type: The type of hardware accelerators associated with this node.
        :param pulumi.Input[str] cidr_block: The CIDR block that the TPU node will use when selecting an IP
               address. This CIDR block must be a /29 block; the Compute Engine
               networks API forbids a smaller block, and using a larger block would
               be wasteful (a node can only consume one IP address).
               Errors will occur if the CIDR block has already been used for a
               currently existing TPU node, the CIDR block conflicts with any
               subnetworks in the user's provided network, or the provided network
               is peered with another network that is using that CIDR block.
        :param pulumi.Input[str] description: The user-supplied description of the TPU. Maximum of 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: The immutable name of the TPU.
        :param pulumi.Input[str] network: The name of a network to peer the TPU node to. It must be a
               preexisting Compute Engine network inside of the project on which
               this API has been activated. If none is provided, "default" will be
               used.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['NodeSchedulingConfigArgs']] scheduling_config: Sets the scheduling options for this TPU instance.
               Structure is documented below.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[str] zone: The GCP location for the TPU.
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

            if accelerator_type is None:
                raise TypeError("Missing required property 'accelerator_type'")
            __props__['accelerator_type'] = accelerator_type
            if cidr_block is None:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['network'] = network
            __props__['project'] = project
            __props__['scheduling_config'] = scheduling_config
            if tensorflow_version is None:
                raise TypeError("Missing required property 'tensorflow_version'")
            __props__['tensorflow_version'] = tensorflow_version
            if zone is None:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
            __props__['network_endpoints'] = None
            __props__['service_account'] = None
        super(Node, __self__).__init__(
            'gcp:tpu/node:Node',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accelerator_type: Optional[pulumi.Input[str]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            network_endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['NodeNetworkEndpointArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            scheduling_config: Optional[pulumi.Input[pulumi.InputType['NodeSchedulingConfigArgs']]] = None,
            service_account: Optional[pulumi.Input[str]] = None,
            tensorflow_version: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Node':
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accelerator_type: The type of hardware accelerators associated with this node.
        :param pulumi.Input[str] cidr_block: The CIDR block that the TPU node will use when selecting an IP
               address. This CIDR block must be a /29 block; the Compute Engine
               networks API forbids a smaller block, and using a larger block would
               be wasteful (a node can only consume one IP address).
               Errors will occur if the CIDR block has already been used for a
               currently existing TPU node, the CIDR block conflicts with any
               subnetworks in the user's provided network, or the provided network
               is peered with another network that is using that CIDR block.
        :param pulumi.Input[str] description: The user-supplied description of the TPU. Maximum of 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: The immutable name of the TPU.
        :param pulumi.Input[str] network: The name of a network to peer the TPU node to. It must be a
               preexisting Compute Engine network inside of the project on which
               this API has been activated. If none is provided, "default" will be
               used.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['NodeNetworkEndpointArgs']]]] network_endpoints: The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
               node first reach out to the first (index 0) entry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['NodeSchedulingConfigArgs']] scheduling_config: Sets the scheduling options for this TPU instance.
               Structure is documented below.
        :param pulumi.Input[str] service_account: The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
               Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[str] zone: The GCP location for the TPU.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accelerator_type"] = accelerator_type
        __props__["cidr_block"] = cidr_block
        __props__["description"] = description
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["network"] = network
        __props__["network_endpoints"] = network_endpoints
        __props__["project"] = project
        __props__["scheduling_config"] = scheduling_config
        __props__["service_account"] = service_account
        __props__["tensorflow_version"] = tensorflow_version
        __props__["zone"] = zone
        return Node(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceleratorType")
    def accelerator_type(self) -> str:
        """
        The type of hardware accelerators associated with this node.
        """
        return pulumi.get(self, "accelerator_type")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> str:
        """
        The CIDR block that the TPU node will use when selecting an IP
        address. This CIDR block must be a /29 block; the Compute Engine
        networks API forbids a smaller block, and using a larger block would
        be wasteful (a node can only consume one IP address).
        Errors will occur if the CIDR block has already been used for a
        currently existing TPU node, the CIDR block conflicts with any
        subnetworks in the user's provided network, or the provided network
        is peered with another network that is using that CIDR block.
        """
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The user-supplied description of the TPU. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Resource labels to represent user provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The immutable name of the TPU.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The name of a network to peer the TPU node to. It must be a
        preexisting Compute Engine network inside of the project on which
        this API has been activated. If none is provided, "default" will be
        used.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="networkEndpoints")
    def network_endpoints(self) -> List['outputs.NodeNetworkEndpoint']:
        """
        The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
        node first reach out to the first (index 0) entry.
        """
        return pulumi.get(self, "network_endpoints")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="schedulingConfig")
    def scheduling_config(self) -> Optional['outputs.NodeSchedulingConfig']:
        """
        Sets the scheduling options for this TPU instance.
        Structure is documented below.
        """
        return pulumi.get(self, "scheduling_config")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
        Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="tensorflowVersion")
    def tensorflow_version(self) -> str:
        """
        The version of Tensorflow running in the Node.
        """
        return pulumi.get(self, "tensorflow_version")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The GCP location for the TPU.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

