# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Node(pulumi.CustomResource):
    accelerator_type: pulumi.Output[str]
    """
    The type of hardware accelerators associated with this node.
    """
    cidr_block: pulumi.Output[str]
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
    description: pulumi.Output[str]
    """
    The user-supplied description of the TPU. Maximum of 512 characters.
    """
    labels: pulumi.Output[dict]
    """
    Resource labels to represent user provided metadata.
    """
    name: pulumi.Output[str]
    """
    The immutable name of the TPU.
    """
    network: pulumi.Output[str]
    """
    The name of a network to peer the TPU node to. It must be a
    preexisting Compute Engine network inside of the project on which
    this API has been activated. If none is provided, "default" will be
    used.
    """
    network_endpoints: pulumi.Output[list]
    """
    The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
    node first reach out to the first (index 0) entry.

      * `ip_address` (`str`)
      * `port` (`float`)
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    scheduling_config: pulumi.Output[dict]
    """
    Sets the scheduling options for this TPU instance.  Structure is documented below.

      * `preemptible` (`bool`) - Defines whether the TPU instance is preemptible.
    """
    service_account: pulumi.Output[str]
    """
    The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
    Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
    """
    tensorflow_version: pulumi.Output[str]
    """
    The version of Tensorflow running in the Node.
    """
    zone: pulumi.Output[str]
    """
    The GCP location for the TPU.
    """
    def __init__(__self__, resource_name, opts=None, accelerator_type=None, cidr_block=None, description=None, labels=None, name=None, network=None, project=None, scheduling_config=None, tensorflow_version=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        A Cloud TPU instance.


        To get more information about Node, see:

        * [API documentation](https://cloud.google.com/tpu/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/tpu/docs/)

        ## Example Usage

        ### TPU Node Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        available = gcp.tpu.get_tensorflow_versions()
        tpu = gcp.tpu.Node("tpu",
            zone="us-central1-b",
            accelerator_type="v3-8",
            tensorflow_version=available.versions[0],
            cidr_block="10.2.0.0/29")
        ```

        ### TPU Node Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        available = gcp.tpu.get_tensorflow_versions()
        tpu = gcp.tpu.Node("tpu",
            zone="us-central1-b",
            accelerator_type="v3-8",
            cidr_block="10.3.0.0/29",
            tensorflow_version=available.versions[0],
            description="Google Provider test TPU",
            network="default",
            labels={
                "foo": "bar",
            },
            scheduling_config={
                "preemptible": True,
            })
        ```


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
        :param pulumi.Input[dict] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: The immutable name of the TPU.
        :param pulumi.Input[str] network: The name of a network to peer the TPU node to. It must be a
               preexisting Compute Engine network inside of the project on which
               this API has been activated. If none is provided, "default" will be
               used.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] scheduling_config: Sets the scheduling options for this TPU instance.  Structure is documented below.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[str] zone: The GCP location for the TPU.

        The **scheduling_config** object supports the following:

          * `preemptible` (`pulumi.Input[bool]`) - Defines whether the TPU instance is preemptible.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, accelerator_type=None, cidr_block=None, description=None, labels=None, name=None, network=None, network_endpoints=None, project=None, scheduling_config=None, service_account=None, tensorflow_version=None, zone=None):
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[dict] labels: Resource labels to represent user provided metadata.
        :param pulumi.Input[str] name: The immutable name of the TPU.
        :param pulumi.Input[str] network: The name of a network to peer the TPU node to. It must be a
               preexisting Compute Engine network inside of the project on which
               this API has been activated. If none is provided, "default" will be
               used.
        :param pulumi.Input[list] network_endpoints: The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the
               node first reach out to the first (index 0) entry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] scheduling_config: Sets the scheduling options for this TPU instance.  Structure is documented below.
        :param pulumi.Input[str] service_account: The service account used to run the tensor flow services within the node. To share resources, including Google Cloud
               Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        :param pulumi.Input[str] tensorflow_version: The version of Tensorflow running in the Node.
        :param pulumi.Input[str] zone: The GCP location for the TPU.

        The **network_endpoints** object supports the following:

          * `ip_address` (`pulumi.Input[str]`)
          * `port` (`pulumi.Input[float]`)

        The **scheduling_config** object supports the following:

          * `preemptible` (`pulumi.Input[bool]`) - Defines whether the TPU instance is preemptible.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

