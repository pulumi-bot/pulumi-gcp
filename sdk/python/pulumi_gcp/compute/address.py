# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Address']


class Address(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_tier: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an Address resource.

        Each virtual machine instance has an ephemeral internal IP address and,
        optionally, an external IP address. To communicate between instances on
        the same network, you can use an instance's internal IP address. To
        communicate with the Internet and instances outside of the same network,
        you must specify the instance's external IP address.

        Internal IP addresses are ephemeral and only belong to an instance for
        the lifetime of the instance; if the instance is deleted and recreated,
        the instance is assigned a new internal IP address, either by Compute
        Engine or by you. External IP addresses can be either ephemeral or
        static.

        To get more information about Address, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
        * How-to Guides
            * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
            * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)

        ## Example Usage
        ### Address Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ip_address = gcp.compute.Address("ipAddress")
        ```
        ### Address With Subnetwork

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.Network("defaultNetwork")
        default_subnetwork = gcp.compute.Subnetwork("defaultSubnetwork",
            ip_cidr_range="10.0.0.0/16",
            region="us-central1",
            network=default_network.id)
        internal_with_subnet_and_address = gcp.compute.Address("internalWithSubnetAndAddress",
            subnetwork=default_subnetwork.id,
            address_type="INTERNAL",
            address="10.0.42.42",
            region="us-central1")
        ```
        ### Address With Gce Endpoint

        ```python
        import pulumi
        import pulumi_gcp as gcp

        internal_with_gce_endpoint = gcp.compute.Address("internalWithGceEndpoint",
            address_type="INTERNAL",
            purpose="GCE_ENDPOINT")
        ```
        ### Instance With Ip

        ```python
        import pulumi
        import pulumi_gcp as gcp

        static = gcp.compute.Address("static")
        debian_image = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        instance_with_ip = gcp.compute.Instance("instanceWithIp",
            machine_type="f1-micro",
            zone="us-central1-a",
            boot_disk={
                "initializeParams": {
                    "image": debian_image.self_link,
                },
            },
            network_interfaces=[{
                "network": "default",
                "accessConfigs": [{
                    "natIp": static.address,
                }],
            }])
        ```

        ## Import

        Address can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/address:Address default projects/{{project}}/regions/{{region}}/addresses/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/address:Address default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/address:Address default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/address:Address default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The static external IP address represented by this resource. Only
               IPv4 is supported. An address may only be specified for INTERNAL
               address types. The IP address must be inside the specified subnetwork,
               if any.
        :param pulumi.Input[str] address_type: The type of address to reserve.
               Default value is `EXTERNAL`.
               Possible values are `INTERNAL` and `EXTERNAL`.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?`
               which means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network_tier: The networking tier used for configuring this address. If this field is not
               specified, it is assumed to be PREMIUM.
               Possible values are `PREMIUM` and `STANDARD`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of this resource, which can be one of the following values:
               * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
               * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
               This should only be set when using an Internal address.
               Possible values are `GCE_ENDPOINT` and `SHARED_LOADBALANCER_VIP`.
        :param pulumi.Input[str] region: The Region in which the created address should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] subnetwork: The URL of the subnetwork in which to reserve the address. If an IP
               address is specified, it must be within the subnetwork's IP range.
               This field can only be used with INTERNAL type with
               GCE_ENDPOINT/DNS_RESOLVER purposes.
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

            __props__['address'] = address
            __props__['address_type'] = address_type
            __props__['description'] = description
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['network_tier'] = network_tier
            __props__['project'] = project
            __props__['purpose'] = purpose
            __props__['region'] = region
            __props__['subnetwork'] = subnetwork
            __props__['creation_timestamp'] = None
            __props__['label_fingerprint'] = None
            __props__['self_link'] = None
            __props__['users'] = None
        super(Address, __self__).__init__(
            'gcp:compute/address:Address',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            address_type: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_tier: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            subnetwork: Optional[pulumi.Input[str]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Address':
        """
        Get an existing Address resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The static external IP address represented by this resource. Only
               IPv4 is supported. An address may only be specified for INTERNAL
               address types. The IP address must be inside the specified subnetwork,
               if any.
        :param pulumi.Input[str] address_type: The type of address to reserve.
               Default value is `EXTERNAL`.
               Possible values are `INTERNAL` and `EXTERNAL`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this address.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. The name must be 1-63 characters long, and
               comply with RFC1035. Specifically, the name must be 1-63 characters
               long and match the regular expression `a-z?`
               which means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
        :param pulumi.Input[str] network_tier: The networking tier used for configuring this address. If this field is not
               specified, it is assumed to be PREMIUM.
               Possible values are `PREMIUM` and `STANDARD`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] purpose: The purpose of this resource, which can be one of the following values:
               * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
               * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
               This should only be set when using an Internal address.
               Possible values are `GCE_ENDPOINT` and `SHARED_LOADBALANCER_VIP`.
        :param pulumi.Input[str] region: The Region in which the created address should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] subnetwork: The URL of the subnetwork in which to reserve the address. If an IP
               address is specified, it must be within the subnetwork's IP range.
               This field can only be used with INTERNAL type with
               GCE_ENDPOINT/DNS_RESOLVER purposes.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: The URLs of the resources that are using this address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["address_type"] = address_type
        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["network_tier"] = network_tier
        __props__["project"] = project
        __props__["purpose"] = purpose
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["subnetwork"] = subnetwork
        __props__["users"] = users
        return Address(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The static external IP address represented by this resource. Only
        IPv4 is supported. An address may only be specified for INTERNAL
        address types. The IP address must be inside the specified subnetwork,
        if any.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressType")
    def address_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of address to reserve.
        Default value is `EXTERNAL`.
        Possible values are `INTERNAL` and `EXTERNAL`.
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
        Name of the resource. The name must be 1-63 characters long, and
        comply with RFC1035. Specifically, the name must be 1-63 characters
        long and match the regular expression `a-z?`
        which means the first character must be a lowercase letter, and all
        following characters must be a dash, lowercase letter, or digit,
        except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkTier")
    def network_tier(self) -> pulumi.Output[str]:
        """
        The networking tier used for configuring this address. If this field is not
        specified, it is assumed to be PREMIUM.
        Possible values are `PREMIUM` and `STANDARD`.
        """
        return pulumi.get(self, "network_tier")

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
    def purpose(self) -> pulumi.Output[str]:
        """
        The purpose of this resource, which can be one of the following values:
        * GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources.
        * SHARED_LOADBALANCER_VIP for an address that can be used by multiple internal load balancers.
        This should only be set when using an Internal address.
        Possible values are `GCE_ENDPOINT` and `SHARED_LOADBALANCER_VIP`.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The Region in which the created address should reside.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Output[str]:
        """
        The URL of the subnetwork in which to reserve the address. If an IP
        address is specified, it must be within the subnetwork's IP range.
        This field can only be used with INTERNAL type with
        GCE_ENDPOINT/DNS_RESOLVER purposes.
        """
        return pulumi.get(self, "subnetwork")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[str]]:
        """
        The URLs of the resources that are using this address.
        """
        return pulumi.get(self, "users")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

