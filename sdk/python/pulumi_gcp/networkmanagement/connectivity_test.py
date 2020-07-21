# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from ._inputs import *
from . import outputs


class ConnectivityTest(pulumi.CustomResource):
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    The user-supplied description of the Connectivity Test.
    Maximum of 512 characters.
    """
    destination: pulumi.Output['outputs.ConnectivityTestDestination'] = pulumi.output_property("destination")
    """
    Required. Destination specification of the Connectivity Test.
    You can use a combination of destination IP address, Compute
    Engine VM instance, or VPC network to uniquely identify the
    destination location.
    Even if the destination IP address is not unique, the source IP
    location is unique. Usually, the analysis can infer the destination
    endpoint from route information.
    If the destination you specify is a VM instance and the instance has
    multiple network interfaces, then you must also specify either a
    destination IP address or VPC network to identify the destination
    interface.
    A reachability analysis proceeds even if the destination location
    is ambiguous. However, the result can include endpoints that you
    don't intend to test.  Structure is documented below.
    """
    labels: pulumi.Output[Optional[Dict[str, str]]] = pulumi.output_property("labels")
    """
    Resource labels to represent user-provided metadata.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    Unique name for the connectivity test.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    protocol: pulumi.Output[Optional[str]] = pulumi.output_property("protocol")
    """
    IP Protocol of the test. When not provided, "TCP" is assumed.
    """
    related_projects: pulumi.Output[Optional[List[str]]] = pulumi.output_property("relatedProjects")
    """
    Other projects that may be relevant for reachability analysis.
    This is applicable to scenarios where a test can cross project
    boundaries.
    """
    source: pulumi.Output['outputs.ConnectivityTestSource'] = pulumi.output_property("source")
    """
    Required. Source specification of the Connectivity Test.
    You can use a combination of source IP address, virtual machine
    (VM) instance, or Compute Engine network to uniquely identify the
    source location.
    Examples: If the source IP address is an internal IP address within
    a Google Cloud Virtual Private Cloud (VPC) network, then you must
    also specify the VPC network. Otherwise, specify the VM instance,
    which already contains its internal IP address and VPC network
    information.
    If the source of the test is within an on-premises network, then
    you must provide the destination VPC network.
    If the source endpoint is a Compute Engine VM instance with multiple
    network interfaces, the instance itself is not sufficient to
    identify the endpoint. So, you must also specify the source IP
    address or VPC network.
    A reachability analysis proceeds even if the source location is
    ambiguous. However, the test result may include endpoints that
    you don't intend to test.  Structure is documented below.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, description=None, destination=None, labels=None, name=None, project=None, protocol=None, related_projects=None, source=None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A connectivity test are a static analysis of your resource configurations
        that enables you to evaluate connectivity to and from Google Cloud
        resources in your Virtual Private Cloud (VPC) network.

        To get more information about ConnectivityTest, see:

        * [API documentation](https://cloud.google.com/network-intelligence-center/docs/connectivity-tests/reference/networkmanagement/rest/v1/projects.locations.global.connectivityTests)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/network-intelligence-center/docs)

        ## Example Usage
        ### Network Management Connectivity Test Instances

        ```python
        import pulumi
        import pulumi_gcp as gcp

        vpc = gcp.compute.Network("vpc")
        debian9 = gcp.compute.get_image(family="debian-9",
            project="debian-cloud")
        source = gcp.compute.Instance("source",
            machine_type="n1-standard-1",
            boot_disk={
                "initializeParams": {
                    "image": debian9.self_link,
                },
            },
            network_interfaces=[{
                "network": vpc.id,
                "accessConfigs": [{}],
            }])
        destination = gcp.compute.Instance("destination",
            machine_type="n1-standard-1",
            boot_disk={
                "initializeParams": {
                    "image": debian9.self_link,
                },
            },
            network_interfaces=[{
                "network": vpc.id,
                "accessConfigs": [{}],
            }])
        instance_test = gcp.networkmanagement.ConnectivityTest("instance-test",
            source={
                "instance": source.id,
            },
            destination={
                "instance": destination.id,
            },
            protocol="TCP")
        ```
        ### Network Management Connectivity Test Addresses

        ```python
        import pulumi
        import pulumi_gcp as gcp

        vpc = gcp.compute.Network("vpc")
        subnet = gcp.compute.Subnetwork("subnet",
            ip_cidr_range="10.0.0.0/16",
            region="us-central1",
            network=vpc.id)
        source_addr = gcp.compute.Address("source-addr",
            subnetwork=subnet.id,
            address_type="INTERNAL",
            address="10.0.42.42",
            region="us-central1")
        dest_addr = gcp.compute.Address("dest-addr",
            subnetwork=subnet.id,
            address_type="INTERNAL",
            address="10.0.43.43",
            region="us-central1")
        address_test = gcp.networkmanagement.ConnectivityTest("address-test",
            source={
                "ip_address": source_addr.address,
                "project_id": source_addr.project,
                "network": vpc.id,
                "networkType": "GCP_NETWORK",
            },
            destination={
                "ip_address": dest_addr.address,
                "project_id": dest_addr.project,
                "network": vpc.id,
            },
            protocol="UDP")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The user-supplied description of the Connectivity Test.
               Maximum of 512 characters.
        :param pulumi.Input['ConnectivityTestDestinationArgs'] destination: Required. Destination specification of the Connectivity Test.
               You can use a combination of destination IP address, Compute
               Engine VM instance, or VPC network to uniquely identify the
               destination location.
               Even if the destination IP address is not unique, the source IP
               location is unique. Usually, the analysis can infer the destination
               endpoint from route information.
               If the destination you specify is a VM instance and the instance has
               multiple network interfaces, then you must also specify either a
               destination IP address or VPC network to identify the destination
               interface.
               A reachability analysis proceeds even if the destination location
               is ambiguous. However, the result can include endpoints that you
               don't intend to test.  Structure is documented below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: Unique name for the connectivity test.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] protocol: IP Protocol of the test. When not provided, "TCP" is assumed.
        :param pulumi.Input[List[pulumi.Input[str]]] related_projects: Other projects that may be relevant for reachability analysis.
               This is applicable to scenarios where a test can cross project
               boundaries.
        :param pulumi.Input['ConnectivityTestSourceArgs'] source: Required. Source specification of the Connectivity Test.
               You can use a combination of source IP address, virtual machine
               (VM) instance, or Compute Engine network to uniquely identify the
               source location.
               Examples: If the source IP address is an internal IP address within
               a Google Cloud Virtual Private Cloud (VPC) network, then you must
               also specify the VPC network. Otherwise, specify the VM instance,
               which already contains its internal IP address and VPC network
               information.
               If the source of the test is within an on-premises network, then
               you must provide the destination VPC network.
               If the source endpoint is a Compute Engine VM instance with multiple
               network interfaces, the instance itself is not sufficient to
               identify the endpoint. So, you must also specify the source IP
               address or VPC network.
               A reachability analysis proceeds even if the source location is
               ambiguous. However, the test result may include endpoints that
               you don't intend to test.  Structure is documented below.
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

            __props__['description'] = description
            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['project'] = project
            __props__['protocol'] = protocol
            __props__['related_projects'] = related_projects
            if source is None:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
        super(ConnectivityTest, __self__).__init__(
            'gcp:networkmanagement/connectivityTest:ConnectivityTest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, destination=None, labels=None, name=None, project=None, protocol=None, related_projects=None, source=None):
        """
        Get an existing ConnectivityTest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The user-supplied description of the Connectivity Test.
               Maximum of 512 characters.
        :param pulumi.Input['ConnectivityTestDestinationArgs'] destination: Required. Destination specification of the Connectivity Test.
               You can use a combination of destination IP address, Compute
               Engine VM instance, or VPC network to uniquely identify the
               destination location.
               Even if the destination IP address is not unique, the source IP
               location is unique. Usually, the analysis can infer the destination
               endpoint from route information.
               If the destination you specify is a VM instance and the instance has
               multiple network interfaces, then you must also specify either a
               destination IP address or VPC network to identify the destination
               interface.
               A reachability analysis proceeds even if the destination location
               is ambiguous. However, the result can include endpoints that you
               don't intend to test.  Structure is documented below.
        :param pulumi.Input[Dict[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: Unique name for the connectivity test.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] protocol: IP Protocol of the test. When not provided, "TCP" is assumed.
        :param pulumi.Input[List[pulumi.Input[str]]] related_projects: Other projects that may be relevant for reachability analysis.
               This is applicable to scenarios where a test can cross project
               boundaries.
        :param pulumi.Input['ConnectivityTestSourceArgs'] source: Required. Source specification of the Connectivity Test.
               You can use a combination of source IP address, virtual machine
               (VM) instance, or Compute Engine network to uniquely identify the
               source location.
               Examples: If the source IP address is an internal IP address within
               a Google Cloud Virtual Private Cloud (VPC) network, then you must
               also specify the VPC network. Otherwise, specify the VM instance,
               which already contains its internal IP address and VPC network
               information.
               If the source of the test is within an on-premises network, then
               you must provide the destination VPC network.
               If the source endpoint is a Compute Engine VM instance with multiple
               network interfaces, the instance itself is not sufficient to
               identify the endpoint. So, you must also specify the source IP
               address or VPC network.
               A reachability analysis proceeds even if the source location is
               ambiguous. However, the test result may include endpoints that
               you don't intend to test.  Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["destination"] = destination
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["project"] = project
        __props__["protocol"] = protocol
        __props__["related_projects"] = related_projects
        __props__["source"] = source
        return ConnectivityTest(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

