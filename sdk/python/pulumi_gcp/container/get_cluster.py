# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, additional_zones=None, addons_configs=None, cluster_autoscalings=None, cluster_ipv4_cidr=None, description=None, enable_binary_authorization=None, enable_kubernetes_alpha=None, enable_legacy_abac=None, enable_tpu=None, endpoint=None, initial_node_count=None, instance_group_urls=None, ip_allocation_policies=None, logging_service=None, maintenance_policies=None, master_auths=None, master_authorized_networks_configs=None, master_ipv4_cidr_block=None, master_version=None, min_master_version=None, monitoring_service=None, network=None, network_policies=None, node_configs=None, node_pools=None, node_version=None, pod_security_policy_configs=None, private_cluster=None, private_cluster_configs=None, remove_default_node_pool=None, resource_labels=None, subnetwork=None, id=None):
        if additional_zones and not isinstance(additional_zones, list):
            raise TypeError('Expected argument additional_zones to be a list')
        __self__.additional_zones = additional_zones
        if addons_configs and not isinstance(addons_configs, list):
            raise TypeError('Expected argument addons_configs to be a list')
        __self__.addons_configs = addons_configs
        if cluster_autoscalings and not isinstance(cluster_autoscalings, list):
            raise TypeError('Expected argument cluster_autoscalings to be a list')
        __self__.cluster_autoscalings = cluster_autoscalings
        if cluster_ipv4_cidr and not isinstance(cluster_ipv4_cidr, str):
            raise TypeError('Expected argument cluster_ipv4_cidr to be a str')
        __self__.cluster_ipv4_cidr = cluster_ipv4_cidr
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        if enable_binary_authorization and not isinstance(enable_binary_authorization, bool):
            raise TypeError('Expected argument enable_binary_authorization to be a bool')
        __self__.enable_binary_authorization = enable_binary_authorization
        if enable_kubernetes_alpha and not isinstance(enable_kubernetes_alpha, bool):
            raise TypeError('Expected argument enable_kubernetes_alpha to be a bool')
        __self__.enable_kubernetes_alpha = enable_kubernetes_alpha
        if enable_legacy_abac and not isinstance(enable_legacy_abac, bool):
            raise TypeError('Expected argument enable_legacy_abac to be a bool')
        __self__.enable_legacy_abac = enable_legacy_abac
        if enable_tpu and not isinstance(enable_tpu, bool):
            raise TypeError('Expected argument enable_tpu to be a bool')
        __self__.enable_tpu = enable_tpu
        if endpoint and not isinstance(endpoint, str):
            raise TypeError('Expected argument endpoint to be a str')
        __self__.endpoint = endpoint
        if initial_node_count and not isinstance(initial_node_count, float):
            raise TypeError('Expected argument initial_node_count to be a float')
        __self__.initial_node_count = initial_node_count
        if instance_group_urls and not isinstance(instance_group_urls, list):
            raise TypeError('Expected argument instance_group_urls to be a list')
        __self__.instance_group_urls = instance_group_urls
        if ip_allocation_policies and not isinstance(ip_allocation_policies, list):
            raise TypeError('Expected argument ip_allocation_policies to be a list')
        __self__.ip_allocation_policies = ip_allocation_policies
        if logging_service and not isinstance(logging_service, str):
            raise TypeError('Expected argument logging_service to be a str')
        __self__.logging_service = logging_service
        if maintenance_policies and not isinstance(maintenance_policies, list):
            raise TypeError('Expected argument maintenance_policies to be a list')
        __self__.maintenance_policies = maintenance_policies
        if master_auths and not isinstance(master_auths, list):
            raise TypeError('Expected argument master_auths to be a list')
        __self__.master_auths = master_auths
        if master_authorized_networks_configs and not isinstance(master_authorized_networks_configs, list):
            raise TypeError('Expected argument master_authorized_networks_configs to be a list')
        __self__.master_authorized_networks_configs = master_authorized_networks_configs
        if master_ipv4_cidr_block and not isinstance(master_ipv4_cidr_block, str):
            raise TypeError('Expected argument master_ipv4_cidr_block to be a str')
        __self__.master_ipv4_cidr_block = master_ipv4_cidr_block
        if master_version and not isinstance(master_version, str):
            raise TypeError('Expected argument master_version to be a str')
        __self__.master_version = master_version
        if min_master_version and not isinstance(min_master_version, str):
            raise TypeError('Expected argument min_master_version to be a str')
        __self__.min_master_version = min_master_version
        if monitoring_service and not isinstance(monitoring_service, str):
            raise TypeError('Expected argument monitoring_service to be a str')
        __self__.monitoring_service = monitoring_service
        if network and not isinstance(network, str):
            raise TypeError('Expected argument network to be a str')
        __self__.network = network
        if network_policies and not isinstance(network_policies, list):
            raise TypeError('Expected argument network_policies to be a list')
        __self__.network_policies = network_policies
        if node_configs and not isinstance(node_configs, list):
            raise TypeError('Expected argument node_configs to be a list')
        __self__.node_configs = node_configs
        if node_pools and not isinstance(node_pools, list):
            raise TypeError('Expected argument node_pools to be a list')
        __self__.node_pools = node_pools
        if node_version and not isinstance(node_version, str):
            raise TypeError('Expected argument node_version to be a str')
        __self__.node_version = node_version
        if pod_security_policy_configs and not isinstance(pod_security_policy_configs, list):
            raise TypeError('Expected argument pod_security_policy_configs to be a list')
        __self__.pod_security_policy_configs = pod_security_policy_configs
        if private_cluster and not isinstance(private_cluster, bool):
            raise TypeError('Expected argument private_cluster to be a bool')
        __self__.private_cluster = private_cluster
        if private_cluster_configs and not isinstance(private_cluster_configs, list):
            raise TypeError('Expected argument private_cluster_configs to be a list')
        __self__.private_cluster_configs = private_cluster_configs
        if remove_default_node_pool and not isinstance(remove_default_node_pool, bool):
            raise TypeError('Expected argument remove_default_node_pool to be a bool')
        __self__.remove_default_node_pool = remove_default_node_pool
        if resource_labels and not isinstance(resource_labels, dict):
            raise TypeError('Expected argument resource_labels to be a dict')
        __self__.resource_labels = resource_labels
        if subnetwork and not isinstance(subnetwork, str):
            raise TypeError('Expected argument subnetwork to be a str')
        __self__.subnetwork = subnetwork
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_cluster(name=None,project=None,region=None,zone=None,opts=None):
    """
    Get info about a cluster within GKE from its name and zone.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['zone'] = zone
    __ret__ = await pulumi.runtime.invoke('gcp:container/getCluster:getCluster', __args__, opts=opts)

    return GetClusterResult(
        additional_zones=__ret__.get('additionalZones'),
        addons_configs=__ret__.get('addonsConfigs'),
        cluster_autoscalings=__ret__.get('clusterAutoscalings'),
        cluster_ipv4_cidr=__ret__.get('clusterIpv4Cidr'),
        description=__ret__.get('description'),
        enable_binary_authorization=__ret__.get('enableBinaryAuthorization'),
        enable_kubernetes_alpha=__ret__.get('enableKubernetesAlpha'),
        enable_legacy_abac=__ret__.get('enableLegacyAbac'),
        enable_tpu=__ret__.get('enableTpu'),
        endpoint=__ret__.get('endpoint'),
        initial_node_count=__ret__.get('initialNodeCount'),
        instance_group_urls=__ret__.get('instanceGroupUrls'),
        ip_allocation_policies=__ret__.get('ipAllocationPolicies'),
        logging_service=__ret__.get('loggingService'),
        maintenance_policies=__ret__.get('maintenancePolicies'),
        master_auths=__ret__.get('masterAuths'),
        master_authorized_networks_configs=__ret__.get('masterAuthorizedNetworksConfigs'),
        master_ipv4_cidr_block=__ret__.get('masterIpv4CidrBlock'),
        master_version=__ret__.get('masterVersion'),
        min_master_version=__ret__.get('minMasterVersion'),
        monitoring_service=__ret__.get('monitoringService'),
        network=__ret__.get('network'),
        network_policies=__ret__.get('networkPolicies'),
        node_configs=__ret__.get('nodeConfigs'),
        node_pools=__ret__.get('nodePools'),
        node_version=__ret__.get('nodeVersion'),
        pod_security_policy_configs=__ret__.get('podSecurityPolicyConfigs'),
        private_cluster=__ret__.get('privateCluster'),
        private_cluster_configs=__ret__.get('privateClusterConfigs'),
        remove_default_node_pool=__ret__.get('removeDefaultNodePool'),
        resource_labels=__ret__.get('resourceLabels'),
        subnetwork=__ret__.get('subnetwork'),
        id=__ret__.get('id'))
