# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'EnvironmentConfigArgs',
    'EnvironmentConfigNodeConfigArgs',
    'EnvironmentConfigNodeConfigIpAllocationPolicyArgs',
    'EnvironmentConfigPrivateEnvironmentConfigArgs',
    'EnvironmentConfigSoftwareConfigArgs',
    'EnvironmentConfigWebServerNetworkAccessControlArgs',
    'EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs',
]

@pulumi.input_type
class EnvironmentConfigArgs:
    def __init__(__self__, *,
                 airflow_uri: Optional[pulumi.Input[str]] = None,
                 dag_gcs_prefix: Optional[pulumi.Input[str]] = None,
                 gke_cluster: Optional[pulumi.Input[str]] = None,
                 node_config: Optional[pulumi.Input['EnvironmentConfigNodeConfigArgs']] = None,
                 node_count: Optional[pulumi.Input[float]] = None,
                 private_environment_config: Optional[pulumi.Input['EnvironmentConfigPrivateEnvironmentConfigArgs']] = None,
                 software_config: Optional[pulumi.Input['EnvironmentConfigSoftwareConfigArgs']] = None,
                 web_server_network_access_control: Optional[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlArgs']] = None):
        """
        :param pulumi.Input['EnvironmentConfigNodeConfigArgs'] node_config: The configuration used for the Kubernetes Engine cluster.  Structure is documented below.
        :param pulumi.Input[float] node_count: The number of nodes in the Kubernetes Engine cluster that
               will be used to run this environment.
        :param pulumi.Input['EnvironmentConfigPrivateEnvironmentConfigArgs'] private_environment_config: The configuration used for the Private IP Cloud Composer environment. Structure is documented below.
        :param pulumi.Input['EnvironmentConfigSoftwareConfigArgs'] software_config: The configuration settings for software inside the environment.  Structure is documented below.
        :param pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlArgs'] web_server_network_access_control: The network-level access control policy for the Airflow web server. If unspecified, no network-level access restrictions will be applied.
        """
        if airflow_uri is not None:
            pulumi.set(__self__, "airflow_uri", airflow_uri)
        if dag_gcs_prefix is not None:
            pulumi.set(__self__, "dag_gcs_prefix", dag_gcs_prefix)
        if gke_cluster is not None:
            pulumi.set(__self__, "gke_cluster", gke_cluster)
        if node_config is not None:
            pulumi.set(__self__, "node_config", node_config)
        if node_count is not None:
            pulumi.set(__self__, "node_count", node_count)
        if private_environment_config is not None:
            pulumi.set(__self__, "private_environment_config", private_environment_config)
        if software_config is not None:
            pulumi.set(__self__, "software_config", software_config)
        if web_server_network_access_control is not None:
            pulumi.set(__self__, "web_server_network_access_control", web_server_network_access_control)

    @property
    @pulumi.getter(name="airflowUri")
    def airflow_uri(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "airflow_uri")

    @airflow_uri.setter
    def airflow_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "airflow_uri", value)

    @property
    @pulumi.getter(name="dagGcsPrefix")
    def dag_gcs_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dag_gcs_prefix")

    @dag_gcs_prefix.setter
    def dag_gcs_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dag_gcs_prefix", value)

    @property
    @pulumi.getter(name="gkeCluster")
    def gke_cluster(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gke_cluster")

    @gke_cluster.setter
    def gke_cluster(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gke_cluster", value)

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> Optional[pulumi.Input['EnvironmentConfigNodeConfigArgs']]:
        """
        The configuration used for the Kubernetes Engine cluster.  Structure is documented below.
        """
        return pulumi.get(self, "node_config")

    @node_config.setter
    def node_config(self, value: Optional[pulumi.Input['EnvironmentConfigNodeConfigArgs']]):
        pulumi.set(self, "node_config", value)

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> Optional[pulumi.Input[float]]:
        """
        The number of nodes in the Kubernetes Engine cluster that
        will be used to run this environment.
        """
        return pulumi.get(self, "node_count")

    @node_count.setter
    def node_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "node_count", value)

    @property
    @pulumi.getter(name="privateEnvironmentConfig")
    def private_environment_config(self) -> Optional[pulumi.Input['EnvironmentConfigPrivateEnvironmentConfigArgs']]:
        """
        The configuration used for the Private IP Cloud Composer environment. Structure is documented below.
        """
        return pulumi.get(self, "private_environment_config")

    @private_environment_config.setter
    def private_environment_config(self, value: Optional[pulumi.Input['EnvironmentConfigPrivateEnvironmentConfigArgs']]):
        pulumi.set(self, "private_environment_config", value)

    @property
    @pulumi.getter(name="softwareConfig")
    def software_config(self) -> Optional[pulumi.Input['EnvironmentConfigSoftwareConfigArgs']]:
        """
        The configuration settings for software inside the environment.  Structure is documented below.
        """
        return pulumi.get(self, "software_config")

    @software_config.setter
    def software_config(self, value: Optional[pulumi.Input['EnvironmentConfigSoftwareConfigArgs']]):
        pulumi.set(self, "software_config", value)

    @property
    @pulumi.getter(name="webServerNetworkAccessControl")
    def web_server_network_access_control(self) -> Optional[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlArgs']]:
        """
        The network-level access control policy for the Airflow web server. If unspecified, no network-level access restrictions will be applied.
        """
        return pulumi.get(self, "web_server_network_access_control")

    @web_server_network_access_control.setter
    def web_server_network_access_control(self, value: Optional[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlArgs']]):
        pulumi.set(self, "web_server_network_access_control", value)


@pulumi.input_type
class EnvironmentConfigNodeConfigArgs:
    def __init__(__self__, *,
                 zone: pulumi.Input[str],
                 disk_size_gb: Optional[pulumi.Input[float]] = None,
                 ip_allocation_policy: Optional[pulumi.Input['EnvironmentConfigNodeConfigIpAllocationPolicyArgs']] = None,
                 machine_type: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 oauth_scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 service_account: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] zone: The Compute Engine zone in which to deploy the VMs running the
               Apache Airflow software, specified as the zone name or
               relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project
               and region.
        :param pulumi.Input[float] disk_size_gb: The disk size in GB used for node VMs. Minimum size is 20GB.
               If unspecified, defaults to 100GB. Cannot be updated.
        :param pulumi.Input['EnvironmentConfigNodeConfigIpAllocationPolicyArgs'] ip_allocation_policy: Configuration for controlling how IPs are allocated in the GKE cluster.
               Structure is documented below.
               Cannot be updated.
        :param pulumi.Input[str] machine_type: The Compute Engine machine type used for cluster instances,
               specified as a name or relative resource name. For example:
               "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and
               region/zone.
        :param pulumi.Input[str] network: The Compute Engine network to be used for machine
               communications, specified as a self-link, relative resource name
               (e.g. "projects/{project}/global/networks/{network}"), by name.
        :param pulumi.Input[List[pulumi.Input[str]]] oauth_scopes: The set of Google API scopes to be made available on all node
               VMs. Cannot be updated. If empty, defaults to
               `["https://www.googleapis.com/auth/cloud-platform"]`
        :param pulumi.Input[str] service_account: The Google Cloud Platform Service Account to be used by the
               node VMs. If a service account is not specified, the "default"
               Compute Engine service account is used. Cannot be updated. If given,
               note that the service account must have `roles/composer.worker`
               for any GCP resources created under the Cloud Composer Environment.
        :param pulumi.Input[str] subnetwork: The Compute Engine subnetwork to be used for machine
               communications, , specified as a self-link, relative resource name (e.g.
               "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided,
               network must also be provided and the subnetwork must belong to the enclosing environment's project and region.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: The list of instance tags applied to all node VMs. Tags are
               used to identify valid sources or targets for network
               firewalls. Each tag within the list must comply with RFC1035.
               Cannot be updated.
        """
        pulumi.set(__self__, "zone", zone)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if ip_allocation_policy is not None:
            pulumi.set(__self__, "ip_allocation_policy", ip_allocation_policy)
        if machine_type is not None:
            pulumi.set(__self__, "machine_type", machine_type)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if oauth_scopes is not None:
            pulumi.set(__self__, "oauth_scopes", oauth_scopes)
        if service_account is not None:
            pulumi.set(__self__, "service_account", service_account)
        if subnetwork is not None:
            pulumi.set(__self__, "subnetwork", subnetwork)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The Compute Engine zone in which to deploy the VMs running the
        Apache Airflow software, specified as the zone name or
        relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project
        and region.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter(name="diskSizeGb")
    def disk_size_gb(self) -> Optional[pulumi.Input[float]]:
        """
        The disk size in GB used for node VMs. Minimum size is 20GB.
        If unspecified, defaults to 100GB. Cannot be updated.
        """
        return pulumi.get(self, "disk_size_gb")

    @disk_size_gb.setter
    def disk_size_gb(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "disk_size_gb", value)

    @property
    @pulumi.getter(name="ipAllocationPolicy")
    def ip_allocation_policy(self) -> Optional[pulumi.Input['EnvironmentConfigNodeConfigIpAllocationPolicyArgs']]:
        """
        Configuration for controlling how IPs are allocated in the GKE cluster.
        Structure is documented below.
        Cannot be updated.
        """
        return pulumi.get(self, "ip_allocation_policy")

    @ip_allocation_policy.setter
    def ip_allocation_policy(self, value: Optional[pulumi.Input['EnvironmentConfigNodeConfigIpAllocationPolicyArgs']]):
        pulumi.set(self, "ip_allocation_policy", value)

    @property
    @pulumi.getter(name="machineType")
    def machine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The Compute Engine machine type used for cluster instances,
        specified as a name or relative resource name. For example:
        "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and
        region/zone.
        """
        return pulumi.get(self, "machine_type")

    @machine_type.setter
    def machine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_type", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The Compute Engine network to be used for machine
        communications, specified as a self-link, relative resource name
        (e.g. "projects/{project}/global/networks/{network}"), by name.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter(name="oauthScopes")
    def oauth_scopes(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The set of Google API scopes to be made available on all node
        VMs. Cannot be updated. If empty, defaults to
        `["https://www.googleapis.com/auth/cloud-platform"]`
        """
        return pulumi.get(self, "oauth_scopes")

    @oauth_scopes.setter
    def oauth_scopes(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "oauth_scopes", value)

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> Optional[pulumi.Input[str]]:
        """
        The Google Cloud Platform Service Account to be used by the
        node VMs. If a service account is not specified, the "default"
        Compute Engine service account is used. Cannot be updated. If given,
        note that the service account must have `roles/composer.worker`
        for any GCP resources created under the Cloud Composer Environment.
        """
        return pulumi.get(self, "service_account")

    @service_account.setter
    def service_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account", value)

    @property
    @pulumi.getter
    def subnetwork(self) -> Optional[pulumi.Input[str]]:
        """
        The Compute Engine subnetwork to be used for machine
        communications, , specified as a self-link, relative resource name (e.g.
        "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided,
        network must also be provided and the subnetwork must belong to the enclosing environment's project and region.
        """
        return pulumi.get(self, "subnetwork")

    @subnetwork.setter
    def subnetwork(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnetwork", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of instance tags applied to all node VMs. Tags are
        used to identify valid sources or targets for network
        firewalls. Each tag within the list must comply with RFC1035.
        Cannot be updated.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class EnvironmentConfigNodeConfigIpAllocationPolicyArgs:
    def __init__(__self__, *,
                 use_ip_aliases: pulumi.Input[bool],
                 cluster_ipv4_cidr_block: Optional[pulumi.Input[str]] = None,
                 cluster_secondary_range_name: Optional[pulumi.Input[str]] = None,
                 services_ipv4_cidr_block: Optional[pulumi.Input[str]] = None,
                 services_secondary_range_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] use_ip_aliases: Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created.
               Defaults to true if the `ip_allocation_block` is present in config.
        :param pulumi.Input[str] cluster_ipv4_cidr_block: The IP address range used to allocate IP addresses to pods in the cluster.
               Set to blank to have GKE choose a range with the default size.
               Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
               Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
               (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
               Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        :param pulumi.Input[str] cluster_secondary_range_name: The name of the cluster's secondary range used to allocate IP addresses to pods.
               Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
               This field is applicable only when `use_ip_aliases` is true.
        :param pulumi.Input[str] services_ipv4_cidr_block: The IP address range used to allocate IP addresses in this cluster.
               Set to blank to have GKE choose a range with the default size.
               Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
               Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
               (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
               Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        :param pulumi.Input[str] services_secondary_range_name: The name of the services' secondary range used to allocate IP addresses to the cluster.
               Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
               This field is applicable only when `use_ip_aliases` is true.
        """
        pulumi.set(__self__, "use_ip_aliases", use_ip_aliases)
        if cluster_ipv4_cidr_block is not None:
            pulumi.set(__self__, "cluster_ipv4_cidr_block", cluster_ipv4_cidr_block)
        if cluster_secondary_range_name is not None:
            pulumi.set(__self__, "cluster_secondary_range_name", cluster_secondary_range_name)
        if services_ipv4_cidr_block is not None:
            pulumi.set(__self__, "services_ipv4_cidr_block", services_ipv4_cidr_block)
        if services_secondary_range_name is not None:
            pulumi.set(__self__, "services_secondary_range_name", services_secondary_range_name)

    @property
    @pulumi.getter(name="useIpAliases")
    def use_ip_aliases(self) -> pulumi.Input[bool]:
        """
        Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created.
        Defaults to true if the `ip_allocation_block` is present in config.
        """
        return pulumi.get(self, "use_ip_aliases")

    @use_ip_aliases.setter
    def use_ip_aliases(self, value: pulumi.Input[bool]):
        pulumi.set(self, "use_ip_aliases", value)

    @property
    @pulumi.getter(name="clusterIpv4CidrBlock")
    def cluster_ipv4_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address range used to allocate IP addresses to pods in the cluster.
        Set to blank to have GKE choose a range with the default size.
        Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
        Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
        (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
        Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        """
        return pulumi.get(self, "cluster_ipv4_cidr_block")

    @cluster_ipv4_cidr_block.setter
    def cluster_ipv4_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_ipv4_cidr_block", value)

    @property
    @pulumi.getter(name="clusterSecondaryRangeName")
    def cluster_secondary_range_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster's secondary range used to allocate IP addresses to pods.
        Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        This field is applicable only when `use_ip_aliases` is true.
        """
        return pulumi.get(self, "cluster_secondary_range_name")

    @cluster_secondary_range_name.setter
    def cluster_secondary_range_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_secondary_range_name", value)

    @property
    @pulumi.getter(name="servicesIpv4CidrBlock")
    def services_ipv4_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address range used to allocate IP addresses in this cluster.
        Set to blank to have GKE choose a range with the default size.
        Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
        Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
        (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
        Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        """
        return pulumi.get(self, "services_ipv4_cidr_block")

    @services_ipv4_cidr_block.setter
    def services_ipv4_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "services_ipv4_cidr_block", value)

    @property
    @pulumi.getter(name="servicesSecondaryRangeName")
    def services_secondary_range_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the services' secondary range used to allocate IP addresses to the cluster.
        Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        This field is applicable only when `use_ip_aliases` is true.
        """
        return pulumi.get(self, "services_secondary_range_name")

    @services_secondary_range_name.setter
    def services_secondary_range_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "services_secondary_range_name", value)


@pulumi.input_type
class EnvironmentConfigPrivateEnvironmentConfigArgs:
    def __init__(__self__, *,
                 cloud_sql_ipv4_cidr_block: Optional[pulumi.Input[str]] = None,
                 enable_private_endpoint: Optional[pulumi.Input[bool]] = None,
                 master_ipv4_cidr_block: Optional[pulumi.Input[str]] = None,
                 web_server_ipv4_cidr_block: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] cloud_sql_ipv4_cidr_block: The CIDR block from which IP range in tenant project will be reserved for Cloud SQL. Needs to be disjoint from `web_server_ipv4_cidr_block`
        :param pulumi.Input[bool] enable_private_endpoint: -
               If true, access to the public endpoint of the GKE cluster is denied.
        :param pulumi.Input[str] master_ipv4_cidr_block: The IP range in CIDR notation to use for the hosted master network. This range is used
               for assigning internal IP addresses to the cluster master or set of masters and to the
               internal load balancer virtual IP. This range must not overlap with any other ranges
               in use within the cluster's network.
               If left blank, the default value of '172.16.0.0/28' is used.
        :param pulumi.Input[str] web_server_ipv4_cidr_block: The CIDR block from which IP range for web server will be reserved. Needs to be disjoint from `master_ipv4_cidr_block` and `cloud_sql_ipv4_cidr_block`.
        """
        if cloud_sql_ipv4_cidr_block is not None:
            pulumi.set(__self__, "cloud_sql_ipv4_cidr_block", cloud_sql_ipv4_cidr_block)
        if enable_private_endpoint is not None:
            pulumi.set(__self__, "enable_private_endpoint", enable_private_endpoint)
        if master_ipv4_cidr_block is not None:
            pulumi.set(__self__, "master_ipv4_cidr_block", master_ipv4_cidr_block)
        if web_server_ipv4_cidr_block is not None:
            pulumi.set(__self__, "web_server_ipv4_cidr_block", web_server_ipv4_cidr_block)

    @property
    @pulumi.getter(name="cloudSqlIpv4CidrBlock")
    def cloud_sql_ipv4_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block from which IP range in tenant project will be reserved for Cloud SQL. Needs to be disjoint from `web_server_ipv4_cidr_block`
        """
        return pulumi.get(self, "cloud_sql_ipv4_cidr_block")

    @cloud_sql_ipv4_cidr_block.setter
    def cloud_sql_ipv4_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_sql_ipv4_cidr_block", value)

    @property
    @pulumi.getter(name="enablePrivateEndpoint")
    def enable_private_endpoint(self) -> Optional[pulumi.Input[bool]]:
        """
        -
        If true, access to the public endpoint of the GKE cluster is denied.
        """
        return pulumi.get(self, "enable_private_endpoint")

    @enable_private_endpoint.setter
    def enable_private_endpoint(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_private_endpoint", value)

    @property
    @pulumi.getter(name="masterIpv4CidrBlock")
    def master_ipv4_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The IP range in CIDR notation to use for the hosted master network. This range is used
        for assigning internal IP addresses to the cluster master or set of masters and to the
        internal load balancer virtual IP. This range must not overlap with any other ranges
        in use within the cluster's network.
        If left blank, the default value of '172.16.0.0/28' is used.
        """
        return pulumi.get(self, "master_ipv4_cidr_block")

    @master_ipv4_cidr_block.setter
    def master_ipv4_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_ipv4_cidr_block", value)

    @property
    @pulumi.getter(name="webServerIpv4CidrBlock")
    def web_server_ipv4_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR block from which IP range for web server will be reserved. Needs to be disjoint from `master_ipv4_cidr_block` and `cloud_sql_ipv4_cidr_block`.
        """
        return pulumi.get(self, "web_server_ipv4_cidr_block")

    @web_server_ipv4_cidr_block.setter
    def web_server_ipv4_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_server_ipv4_cidr_block", value)


@pulumi.input_type
class EnvironmentConfigSoftwareConfigArgs:
    def __init__(__self__, *,
                 airflow_config_overrides: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 env_variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 image_version: Optional[pulumi.Input[str]] = None,
                 pypi_packages: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 python_version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] airflow_config_overrides: -
               (Optional) Apache Airflow configuration properties to override. Property keys contain the section and property names,
               separated by a hyphen, for example "core-dags_are_paused_at_creation".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] env_variables: Additional environment variables to provide to the Apache Airflow scheduler, worker, and webserver processes.
               Environment variable names must match the regular expression `[a-zA-Z_][a-zA-Z0-9_]*`.
               They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression
               `AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+`), and they cannot match any of the following reserved names:
               ```python
               import pulumi
               ```
        :param pulumi.Input[str] image_version: -
               The version of the software running in the environment. This encapsulates both the version of Cloud Composer
               functionality and the version of Apache Airflow. It must match the regular expression
               `composer-[0-9]+\.[0-9]+(\.[0-9]+)?-airflow-[0-9]+\.[0-9]+(\.[0-9]+.*)?`.
               The Cloud Composer portion of the version is a semantic version.
               The portion of the image version following 'airflow-' is an official Apache Airflow repository release name.
               See [documentation](https://cloud.google.com/composer/docs/reference/rest/v1beta1/projects.locations.environments#softwareconfig)
               for allowed release names.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pypi_packages: Custom Python Package Index (PyPI) packages to be installed
               in the environment. Keys refer to the lowercase package name (e.g. "numpy"). Values are the lowercase extras and
               version specifier (e.g. "==1.12.0", "[devel,gcp_api]", "[devel]>=1.8.2, <1.9.2"). To specify a package without
               pinning it to a version specifier, use the empty string as the value.
        :param pulumi.Input[str] python_version: -
               The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes.
               Can be set to '2' or '3'. If not specified, the default is '2'. Cannot be updated.
        """
        if airflow_config_overrides is not None:
            pulumi.set(__self__, "airflow_config_overrides", airflow_config_overrides)
        if env_variables is not None:
            pulumi.set(__self__, "env_variables", env_variables)
        if image_version is not None:
            pulumi.set(__self__, "image_version", image_version)
        if pypi_packages is not None:
            pulumi.set(__self__, "pypi_packages", pypi_packages)
        if python_version is not None:
            pulumi.set(__self__, "python_version", python_version)

    @property
    @pulumi.getter(name="airflowConfigOverrides")
    def airflow_config_overrides(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        -
        (Optional) Apache Airflow configuration properties to override. Property keys contain the section and property names,
        separated by a hyphen, for example "core-dags_are_paused_at_creation".
        """
        return pulumi.get(self, "airflow_config_overrides")

    @airflow_config_overrides.setter
    def airflow_config_overrides(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "airflow_config_overrides", value)

    @property
    @pulumi.getter(name="envVariables")
    def env_variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Additional environment variables to provide to the Apache Airflow scheduler, worker, and webserver processes.
        Environment variable names must match the regular expression `[a-zA-Z_][a-zA-Z0-9_]*`.
        They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression
        `AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+`), and they cannot match any of the following reserved names:
        ```python
        import pulumi
        ```
        """
        return pulumi.get(self, "env_variables")

    @env_variables.setter
    def env_variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "env_variables", value)

    @property
    @pulumi.getter(name="imageVersion")
    def image_version(self) -> Optional[pulumi.Input[str]]:
        """
        -
        The version of the software running in the environment. This encapsulates both the version of Cloud Composer
        functionality and the version of Apache Airflow. It must match the regular expression
        `composer-[0-9]+\.[0-9]+(\.[0-9]+)?-airflow-[0-9]+\.[0-9]+(\.[0-9]+.*)?`.
        The Cloud Composer portion of the version is a semantic version.
        The portion of the image version following 'airflow-' is an official Apache Airflow repository release name.
        See [documentation](https://cloud.google.com/composer/docs/reference/rest/v1beta1/projects.locations.environments#softwareconfig)
        for allowed release names.
        """
        return pulumi.get(self, "image_version")

    @image_version.setter
    def image_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_version", value)

    @property
    @pulumi.getter(name="pypiPackages")
    def pypi_packages(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Custom Python Package Index (PyPI) packages to be installed
        in the environment. Keys refer to the lowercase package name (e.g. "numpy"). Values are the lowercase extras and
        version specifier (e.g. "==1.12.0", "[devel,gcp_api]", "[devel]>=1.8.2, <1.9.2"). To specify a package without
        pinning it to a version specifier, use the empty string as the value.
        """
        return pulumi.get(self, "pypi_packages")

    @pypi_packages.setter
    def pypi_packages(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "pypi_packages", value)

    @property
    @pulumi.getter(name="pythonVersion")
    def python_version(self) -> Optional[pulumi.Input[str]]:
        """
        -
        The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes.
        Can be set to '2' or '3'. If not specified, the default is '2'. Cannot be updated.
        """
        return pulumi.get(self, "python_version")

    @python_version.setter
    def python_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "python_version", value)


@pulumi.input_type
class EnvironmentConfigWebServerNetworkAccessControlArgs:
    def __init__(__self__, *,
                 allowed_ip_ranges: Optional[pulumi.Input[List[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs']]]] = None):
        """
        :param pulumi.Input[List[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs']]] allowed_ip_ranges: -
               A collection of allowed IP ranges with descriptions. Structure is documented below.
        """
        if allowed_ip_ranges is not None:
            pulumi.set(__self__, "allowed_ip_ranges", allowed_ip_ranges)

    @property
    @pulumi.getter(name="allowedIpRanges")
    def allowed_ip_ranges(self) -> Optional[pulumi.Input[List[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs']]]]:
        """
        -
        A collection of allowed IP ranges with descriptions. Structure is documented below.
        """
        return pulumi.get(self, "allowed_ip_ranges")

    @allowed_ip_ranges.setter
    def allowed_ip_ranges(self, value: Optional[pulumi.Input[List[pulumi.Input['EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs']]]]):
        pulumi.set(self, "allowed_ip_ranges", value)


@pulumi.input_type
class EnvironmentConfigWebServerNetworkAccessControlAllowedIpRangeArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] value: IP address or range, defined using CIDR notation, of requests that this rule applies to.
               Examples: `192.168.1.1` or `192.168.0.0/16` or `2001:db8::/32` or `2001:0db8:0000:0042:0000:8a2e:0370:7334`.
               IP range prefixes should be properly truncated. For example,
               `1.2.3.4/24` should be truncated to `1.2.3.0/24`. Similarly, for IPv6, `2001:db8::1/32` should be truncated to `2001:db8::/32`.
        :param pulumi.Input[str] description: A description of this ip range.
        """
        pulumi.set(__self__, "value", value)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        IP address or range, defined using CIDR notation, of requests that this rule applies to.
        Examples: `192.168.1.1` or `192.168.0.0/16` or `2001:db8::/32` or `2001:0db8:0000:0042:0000:8a2e:0370:7334`.
        IP range prefixes should be properly truncated. For example,
        `1.2.3.4/24` should be truncated to `1.2.3.0/24`. Similarly, for IPv6, `2001:db8::1/32` should be truncated to `2001:db8::/32`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of this ip range.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


