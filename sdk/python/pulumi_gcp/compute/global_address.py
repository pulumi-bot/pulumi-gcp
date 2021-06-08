# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GlobalAddressArgs', 'GlobalAddress']

@pulumi.input_type
class GlobalAddressArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GlobalAddress resource.
        :param pulumi.Input[str] address: The IP address or beginning of the address range represented by this
               resource. This can be supplied as an input to reserve a specific
               address or omitted to allow GCP to choose a valid one for you.
        :param pulumi.Input[str] address_type: The type of the address to reserve.
               * EXTERNAL indicates public/external single IP address.
               * INTERNAL indicates internal IP ranges belonging to some network.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL` and `INTERNAL`.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this address. The default value is `IPV4`.
               Possible values are `IPV4` and `IPV6`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network in which to reserve the IP range. The IP range
               must be in RFC1918 space. The network cannot be deleted if there are
               any reserved IP ranges referring to it.
               This should only be set when using an Internal address.
        :param pulumi.Input[int] prefix_length: The prefix length of the IP range. If not present, it means the
               address field is a single IP address.
               This field is not applicable to addresses with addressType=EXTERNAL,
               or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. Possible values include:
               * VPC_PEERING - for peer networks
               * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if address_type is not None:
            pulumi.set(__self__, "address_type", address_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address or beginning of the address range represented by this
        resource. This can be supplied as an input to reserve a specific
        address or omitted to allow GCP to choose a valid one for you.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the address to reserve.
        * EXTERNAL indicates public/external single IP address.
        * INTERNAL indicates internal IP ranges belonging to some network.
        Default value is `EXTERNAL`.
        Possible values are `EXTERNAL` and `INTERNAL`.
        """
        return pulumi.get(self, "address_type")

    @address_type.setter
    def address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Version that will be used by this address. The default value is `IPV4`.
        Possible values are `IPV4` and `IPV6`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to apply to this address.  A list of key->value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the network in which to reserve the IP range. The IP range
        must be in RFC1918 space. The network cannot be deleted if there are
        any reserved IP ranges referring to it.
        This should only be set when using an Internal address.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[int]]:
        """
        The prefix length of the IP range. If not present, it means the
        address field is a single IP address.
        This field is not applicable to addresses with addressType=EXTERNAL,
        or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix_length", value)

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

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        The purpose of the resource. Possible values include:
        * VPC_PEERING - for peer networks
        * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)


@pulumi.input_type
class _GlobalAddressState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 label_fingerprint: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GlobalAddress resources.
        :param pulumi.Input[str] address: The IP address or beginning of the address range represented by this
               resource. This can be supplied as an input to reserve a specific
               address or omitted to allow GCP to choose a valid one for you.
        :param pulumi.Input[str] address_type: The type of the address to reserve.
               * EXTERNAL indicates public/external single IP address.
               * INTERNAL indicates internal IP ranges belonging to some network.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL` and `INTERNAL`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this address. The default value is `IPV4`.
               Possible values are `IPV4` and `IPV6`.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network in which to reserve the IP range. The IP range
               must be in RFC1918 space. The network cannot be deleted if there are
               any reserved IP ranges referring to it.
               This should only be set when using an Internal address.
        :param pulumi.Input[int] prefix_length: The prefix length of the IP range. If not present, it means the
               address field is a single IP address.
               This field is not applicable to addresses with addressType=EXTERNAL,
               or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. Possible values include:
               * VPC_PEERING - for peer networks
               * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if address_type is not None:
            pulumi.set(__self__, "address_type", address_type)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if label_fingerprint is not None:
            pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address or beginning of the address range represented by this
        resource. This can be supplied as an input to reserve a specific
        address or omitted to allow GCP to choose a valid one for you.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the address to reserve.
        * EXTERNAL indicates public/external single IP address.
        * INTERNAL indicates internal IP ranges belonging to some network.
        Default value is `EXTERNAL`.
        Possible values are `EXTERNAL` and `INTERNAL`.
        """
        return pulumi.get(self, "address_type")

    @address_type.setter
    def address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_type", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Version that will be used by this address. The default value is `IPV4`.
        Possible values are `IPV4` and `IPV6`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The fingerprint used for optimistic locking of this resource. Used internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @label_fingerprint.setter
    def label_fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label_fingerprint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to apply to this address.  A list of key->value pairs.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the network in which to reserve the IP range. The IP range
        must be in RFC1918 space. The network cannot be deleted if there are
        any reserved IP ranges referring to it.
        This should only be set when using an Internal address.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[int]]:
        """
        The prefix length of the IP range. If not present, it means the
        address field is a single IP address.
        This field is not applicable to addresses with addressType=EXTERNAL,
        or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix_length", value)

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

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        The purpose of the resource. Possible values include:
        * VPC_PEERING - for peer networks
        * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


class GlobalAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a Global Address resource. Global addresses are used for
        HTTP(S) load balancing.

        To get more information about GlobalAddress, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
        * How-to Guides
            * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)

        ## Example Usage
        ### Global Address Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.GlobalAddress("default")
        ```

        ## Import

        GlobalAddress can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default projects/{{project}}/global/addresses/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address or beginning of the address range represented by this
               resource. This can be supplied as an input to reserve a specific
               address or omitted to allow GCP to choose a valid one for you.
        :param pulumi.Input[str] address_type: The type of the address to reserve.
               * EXTERNAL indicates public/external single IP address.
               * INTERNAL indicates internal IP ranges belonging to some network.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL` and `INTERNAL`.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this address. The default value is `IPV4`.
               Possible values are `IPV4` and `IPV6`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network in which to reserve the IP range. The IP range
               must be in RFC1918 space. The network cannot be deleted if there are
               any reserved IP ranges referring to it.
               This should only be set when using an Internal address.
        :param pulumi.Input[int] prefix_length: The prefix length of the IP range. If not present, it means the
               address field is a single IP address.
               This field is not applicable to addresses with addressType=EXTERNAL,
               or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. Possible values include:
               * VPC_PEERING - for peer networks
               * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: Optional[GlobalAddressArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a Global Address resource. Global addresses are used for
        HTTP(S) load balancing.

        To get more information about GlobalAddress, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/globalAddresses)
        * How-to Guides
            * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address)

        ## Example Usage
        ### Global Address Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.GlobalAddress("default")
        ```

        ## Import

        GlobalAddress can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default projects/{{project}}/global/addresses/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalAddress:GlobalAddress default {{name}}
        ```

        :param str resource_name_: The name of the resource.
        :param GlobalAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
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
            __props__ = GlobalAddressArgs.__new__(GlobalAddressArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["address_type"] = address_type
            __props__.__dict__["description"] = description
            __props__.__dict__["ip_version"] = ip_version
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["network"] = network
            __props__.__dict__["prefix_length"] = prefix_length
            __props__.__dict__["project"] = project
            __props__.__dict__["purpose"] = purpose
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["label_fingerprint"] = None
            __props__.__dict__["self_link"] = None
        super(GlobalAddress, __self__).__init__(
            'gcp:compute/globalAddress:GlobalAddress',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            address_type: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_version: Optional[pulumi.Input[str]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None,
            prefix_length: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'GlobalAddress':
        """
        Get an existing GlobalAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address or beginning of the address range represented by this
               resource. This can be supplied as an input to reserve a specific
               address or omitted to allow GCP to choose a valid one for you.
        :param pulumi.Input[str] address_type: The type of the address to reserve.
               * EXTERNAL indicates public/external single IP address.
               * INTERNAL indicates internal IP ranges belonging to some network.
               Default value is `EXTERNAL`.
               Possible values are `EXTERNAL` and `INTERNAL`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] ip_version: The IP Version that will be used by this address. The default value is `IPV4`.
               Possible values are `IPV4` and `IPV6`.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network: The URL of the network in which to reserve the IP range. The IP range
               must be in RFC1918 space. The network cannot be deleted if there are
               any reserved IP ranges referring to it.
               This should only be set when using an Internal address.
        :param pulumi.Input[int] prefix_length: The prefix length of the IP range. If not present, it means the
               address field is a single IP address.
               This field is not applicable to addresses with addressType=EXTERNAL,
               or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of the resource. Possible values include:
               * VPC_PEERING - for peer networks
               * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalAddressState.__new__(_GlobalAddressState)

        __props__.__dict__["address"] = address
        __props__.__dict__["address_type"] = address_type
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["label_fingerprint"] = label_fingerprint
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["network"] = network
        __props__.__dict__["prefix_length"] = prefix_length
        __props__.__dict__["project"] = project
        __props__.__dict__["purpose"] = purpose
        __props__.__dict__["self_link"] = self_link
        return GlobalAddress(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The IP address or beginning of the address range represented by this
        resource. This can be supplied as an input to reserve a specific
        address or omitted to allow GCP to choose a valid one for you.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the address to reserve.
        * EXTERNAL indicates public/external single IP address.
        * INTERNAL indicates internal IP ranges belonging to some network.
        Default value is `EXTERNAL`.
        Possible values are `EXTERNAL` and `INTERNAL`.
        """
        return pulumi.get(self, "address_type")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[Optional[str]]:
        """
        The IP Version that will be used by this address. The default value is `IPV4`.
        Possible values are `IPV4` and `IPV6`.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint used for optimistic locking of this resource. Used internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to apply to this address.  A list of key->value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[Optional[str]]:
        """
        The URL of the network in which to reserve the IP range. The IP range
        must be in RFC1918 space. The network cannot be deleted if there are
        any reserved IP ranges referring to it.
        This should only be set when using an Internal address.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[Optional[int]]:
        """
        The prefix length of the IP range. If not present, it means the
        address field is a single IP address.
        This field is not applicable to addresses with addressType=EXTERNAL,
        or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[Optional[str]]:
        """
        The purpose of the resource. Possible values include:
        * VPC_PEERING - for peer networks
        * PRIVATE_SERVICE_CONNECT - for Private Service Connect networks
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

