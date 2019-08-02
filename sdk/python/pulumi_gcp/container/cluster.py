# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    additional_zones: pulumi.Output[list]
    """
    The list of zones in which the cluster's nodes
    should be located. These must be in the same region as the cluster zone for
    zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
    the number of nodes specified in `initial_node_count` is created in
    all specified zones as well as the primary zone. If specified for a regional
    cluster, nodes will only be created in these zones. `additional_zones` has been
    deprecated in favour of `node_locations`.
    """
    addons_config: pulumi.Output[dict]
    """
    The configuration for addons supported by GKE.
    Structure is documented below.
    """
    authenticator_groups_config: pulumi.Output[dict]
    """
    ) Configuration for the
    [Google Groups for GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#groups-setup-gsuite) feature.
    Structure is documented below.
    """
    cluster_autoscaling: pulumi.Output[dict]
    """
    )
    Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to
    automatically adjust the size of the cluster and create/delete node pools based
    on the current needs of the cluster's workload. See the
    [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning)
    for more details. Structure is documented below.
    """
    cluster_ipv4_cidr: pulumi.Output[str]
    """
    The IP address range of the kubernetes pods in
    this cluster. Default is an automatically assigned CIDR.
    """
    database_encryption: pulumi.Output[dict]
    """
    ).
    Structure is documented below.
    """
    default_max_pods_per_node: pulumi.Output[float]
    """
    ) The default maximum number of pods per node in this cluster.
    Note that this does not work on node pools which are "route-based" - that is, node
    pools belonging to clusters that do not have IP Aliasing enabled.
    See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
    for more information.
    """
    description: pulumi.Output[str]
    """
    Description of the cluster.
    """
    enable_binary_authorization: pulumi.Output[bool]
    """
    ) Enable Binary Authorization for this cluster.
    If enabled, all container images will be validated by Google Binary Authorization.
    """
    enable_intranode_visibility: pulumi.Output[bool]
    """
    )
    Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.
    """
    enable_kubernetes_alpha: pulumi.Output[bool]
    """
    Whether to enable Kubernetes Alpha features for
    this cluster. Note that when this option is enabled, the cluster cannot be upgraded
    and will be automatically deleted after 30 days.
    """
    enable_legacy_abac: pulumi.Output[bool]
    """
    Whether the ABAC authorizer is enabled for this cluster.
    When enabled, identities in the system, including service accounts, nodes, and controllers,
    will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
    Defaults to `false`
    """
    enable_tpu: pulumi.Output[bool]
    """
    ) Whether to enable Cloud TPU resources in this cluster.
    See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
    """
    endpoint: pulumi.Output[str]
    """
    The IP address of this cluster's Kubernetes master.
    """
    initial_node_count: pulumi.Output[float]
    """
    The number of nodes to create in this
    cluster's default node pool. Must be set if `node_pool` is not set. If
    you're using `google_container_node_pool` objects with no default node pool,
    you'll need to set this to a value of at least `1`, alongside setting
    `remove_default_node_pool` to `true`.
    """
    instance_group_urls: pulumi.Output[list]
    """
    List of instance group URLs which have been assigned
    to the cluster.
    """
    ip_allocation_policy: pulumi.Output[dict]
    """
    Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
    This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
    Structure is documented below. This field is marked to use [Attribute as Block](https://www.terraform.io/docs/configuration/attr-as-blocks.html)
    in order to support explicit removal with `ip_allocation_policy = []`.
    """
    location: pulumi.Output[str]
    """
    The location (region or zone) in which the cluster
    master will be created, as well as the default node location. If you specify a
    zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
    single cluster master. If you specify a region (such as `us-west1`), the
    cluster will be a regional cluster with multiple masters spread across zones in
    the region, and with default node locations in those zones as well.
    """
    logging_service: pulumi.Output[str]
    """
    The logging service that the cluster should
    write logs to. Available options include `logging.googleapis.com`,
    `logging.googleapis.com/kubernetes`, and `none`. Defaults to `logging.googleapis.com`
    """
    maintenance_policy: pulumi.Output[dict]
    """
    The maintenance policy to use for the cluster. Structure is
    documented below.
    """
    master_auth: pulumi.Output[dict]
    """
    The authentication information for accessing the
    Kubernetes master. Some values in this block are only returned by the API if
    your service account has permission to get credentials for your GKE cluster. If
    you see an unexpected diff removing a username/password or unsetting your client
    cert, ensure you have the `container.clusters.getCredentials` permission.
    Structure is documented below.
    """
    master_authorized_networks_config: pulumi.Output[dict]
    """
    The desired configuration options
    for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
    external access (except the cluster node IPs, which GKE automatically whitelists).
    """
    master_version: pulumi.Output[str]
    """
    The current version of the master in the cluster. This may
    be different than the `min_master_version` set in the config if the master
    has been updated by GKE.
    """
    min_master_version: pulumi.Output[str]
    monitoring_service: pulumi.Output[str]
    """
    The monitoring service that the cluster
    should write metrics to.
    Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
    VM metrics will be collected by Google Compute Engine regardless of this setting
    Available options include
    `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes`, and `none`.
    Defaults to `monitoring.googleapis.com`
    """
    name: pulumi.Output[str]
    """
    The name of the cluster, unique within the project and
    location.
    """
    network: pulumi.Output[str]
    """
    The name or self_link of the Google Compute Engine
    network to which the cluster is connected. For Shared VPC, set this to the self link of the
    shared network.
    """
    network_policy: pulumi.Output[dict]
    """
    Configuration options for the
    [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
    feature. Structure is documented below.
    """
    node_config: pulumi.Output[dict]
    node_locations: pulumi.Output[list]
    """
    The list of zones in which the cluster's nodes
    should be located. These must be in the same region as the cluster zone for
    zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
    the number of nodes specified in `initial_node_count` is created in
    all specified zones as well as the primary zone. If specified for a regional
    cluster, nodes will be created in only these zones.
    """
    node_pools: pulumi.Output[list]
    """
    List of node pools associated with this cluster.
    See google_container_node_pool for schema.
    **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
    cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
    to say "these are the _only_ node pools associated with this cluster", use the
    google_container_node_pool resource instead of this property.
    """
    node_version: pulumi.Output[str]
    pod_security_policy_config: pulumi.Output[dict]
    """
    ) Configuration for the
    [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
    Structure is documented below.
    """
    private_cluster_config: pulumi.Output[dict]
    """
    A set of options for creating
    a private cluster. Structure is documented below.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    remove_default_node_pool: pulumi.Output[bool]
    """
    If `true`, deletes the default node
    pool upon cluster creation. If you're using `google_container_node_pool`
    resources with no default node pool, this should be set to `true`, alongside
    setting `initial_node_count` to at least `1`.
    """
    resource_labels: pulumi.Output[dict]
    """
    The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
    """
    resource_usage_export_config: pulumi.Output[dict]
    """
    ) Configuration for the
    [ResourceUsageExportConfig](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering) feature.
    Structure is documented below.
    """
    services_ipv4_cidr: pulumi.Output[str]
    """
    The IP address range of the Kubernetes services in this
    cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
    notation (e.g. `1.2.3.4/29`). Service addresses are typically put in the last
    `/16` from the container CIDR.
    """
    subnetwork: pulumi.Output[str]
    """
    The name or self_link of the Google Compute Engine subnetwork in
    which the cluster's instances are launched.
    """
    tpu_ipv4_cidr_block: pulumi.Output[str]
    vertical_pod_autoscaling: pulumi.Output[dict]
    """
    )
    Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
    Structure is documented below.
    """
    workload_identity_config: pulumi.Output[dict]
    """
    )
    Workload Identity allows Kubernetes service accounts to act as a user-managed
    [Google IAM Service Account](https://cloud.google.com/iam/docs/service-accounts#user-managed_service_accounts).
    """
    zone: pulumi.Output[str]
    """
    The zone that the cluster master and nodes
    should be created in. If specified, this cluster will be a zonal cluster. `zone`
    has been deprecated in favour of `location`.
    """
    def __init__(__self__, resource_name, opts=None, additional_zones=None, addons_config=None, authenticator_groups_config=None, cluster_autoscaling=None, cluster_ipv4_cidr=None, database_encryption=None, default_max_pods_per_node=None, description=None, enable_binary_authorization=None, enable_intranode_visibility=None, enable_kubernetes_alpha=None, enable_legacy_abac=None, enable_tpu=None, initial_node_count=None, ip_allocation_policy=None, location=None, logging_service=None, maintenance_policy=None, master_auth=None, master_authorized_networks_config=None, min_master_version=None, monitoring_service=None, name=None, network=None, network_policy=None, node_config=None, node_locations=None, node_pools=None, node_version=None, pod_security_policy_config=None, private_cluster_config=None, project=None, region=None, remove_default_node_pool=None, resource_labels=None, resource_usage_export_config=None, subnetwork=None, vertical_pod_autoscaling=None, workload_identity_config=None, zone=None, __name__=None, __opts__=None):
        """
        Manages a Google Kubernetes Engine (GKE) cluster. For more information see
        [the official documentation](https://cloud.google.com/container-engine/docs/clusters)
        and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters).
        
        > **Note:** All arguments and attributes, including basic auth username and
        passwords as well as certificate outputs will be stored in the raw state as
        plaintext. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_zones: The list of zones in which the cluster's nodes
               should be located. These must be in the same region as the cluster zone for
               zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
               the number of nodes specified in `initial_node_count` is created in
               all specified zones as well as the primary zone. If specified for a regional
               cluster, nodes will only be created in these zones. `additional_zones` has been
               deprecated in favour of `node_locations`.
        :param pulumi.Input[dict] addons_config: The configuration for addons supported by GKE.
               Structure is documented below.
        :param pulumi.Input[dict] authenticator_groups_config: ) Configuration for the
               [Google Groups for GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#groups-setup-gsuite) feature.
               Structure is documented below.
        :param pulumi.Input[dict] cluster_autoscaling: )
               Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to
               automatically adjust the size of the cluster and create/delete node pools based
               on the current needs of the cluster's workload. See the
               [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning)
               for more details. Structure is documented below.
        :param pulumi.Input[str] cluster_ipv4_cidr: The IP address range of the kubernetes pods in
               this cluster. Default is an automatically assigned CIDR.
        :param pulumi.Input[dict] database_encryption: ).
               Structure is documented below.
        :param pulumi.Input[float] default_max_pods_per_node: ) The default maximum number of pods per node in this cluster.
               Note that this does not work on node pools which are "route-based" - that is, node
               pools belonging to clusters that do not have IP Aliasing enabled.
               See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
               for more information.
        :param pulumi.Input[str] description: Description of the cluster.
        :param pulumi.Input[bool] enable_binary_authorization: ) Enable Binary Authorization for this cluster.
               If enabled, all container images will be validated by Google Binary Authorization.
        :param pulumi.Input[bool] enable_intranode_visibility: )
               Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network.
        :param pulumi.Input[bool] enable_kubernetes_alpha: Whether to enable Kubernetes Alpha features for
               this cluster. Note that when this option is enabled, the cluster cannot be upgraded
               and will be automatically deleted after 30 days.
        :param pulumi.Input[bool] enable_legacy_abac: Whether the ABAC authorizer is enabled for this cluster.
               When enabled, identities in the system, including service accounts, nodes, and controllers,
               will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
               Defaults to `false`
        :param pulumi.Input[bool] enable_tpu: ) Whether to enable Cloud TPU resources in this cluster.
               See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).
        :param pulumi.Input[float] initial_node_count: The number of nodes to create in this
               cluster's default node pool. Must be set if `node_pool` is not set. If
               you're using `google_container_node_pool` objects with no default node pool,
               you'll need to set this to a value of at least `1`, alongside setting
               `remove_default_node_pool` to `true`.
        :param pulumi.Input[dict] ip_allocation_policy: Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported.
               This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases)
               Structure is documented below. This field is marked to use [Attribute as Block](https://www.terraform.io/docs/configuration/attr-as-blocks.html)
               in order to support explicit removal with `ip_allocation_policy = []`.
        :param pulumi.Input[str] location: The location (region or zone) in which the cluster
               master will be created, as well as the default node location. If you specify a
               zone (such as `us-central1-a`), the cluster will be a zonal cluster with a
               single cluster master. If you specify a region (such as `us-west1`), the
               cluster will be a regional cluster with multiple masters spread across zones in
               the region, and with default node locations in those zones as well.
        :param pulumi.Input[str] logging_service: The logging service that the cluster should
               write logs to. Available options include `logging.googleapis.com`,
               `logging.googleapis.com/kubernetes`, and `none`. Defaults to `logging.googleapis.com`
        :param pulumi.Input[dict] maintenance_policy: The maintenance policy to use for the cluster. Structure is
               documented below.
        :param pulumi.Input[dict] master_auth: The authentication information for accessing the
               Kubernetes master. Some values in this block are only returned by the API if
               your service account has permission to get credentials for your GKE cluster. If
               you see an unexpected diff removing a username/password or unsetting your client
               cert, ensure you have the `container.clusters.getCredentials` permission.
               Structure is documented below.
        :param pulumi.Input[dict] master_authorized_networks_config: The desired configuration options
               for master authorized networks. Omit the nested `cidr_blocks` attribute to disallow
               external access (except the cluster node IPs, which GKE automatically whitelists).
        :param pulumi.Input[str] monitoring_service: The monitoring service that the cluster
               should write metrics to.
               Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API.
               VM metrics will be collected by Google Compute Engine regardless of this setting
               Available options include
               `monitoring.googleapis.com`, `monitoring.googleapis.com/kubernetes`, and `none`.
               Defaults to `monitoring.googleapis.com`
        :param pulumi.Input[str] name: The name of the cluster, unique within the project and
               location.
        :param pulumi.Input[str] network: The name or self_link of the Google Compute Engine
               network to which the cluster is connected. For Shared VPC, set this to the self link of the
               shared network.
        :param pulumi.Input[dict] network_policy: Configuration options for the
               [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/)
               feature. Structure is documented below.
        :param pulumi.Input[list] node_locations: The list of zones in which the cluster's nodes
               should be located. These must be in the same region as the cluster zone for
               zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster,
               the number of nodes specified in `initial_node_count` is created in
               all specified zones as well as the primary zone. If specified for a regional
               cluster, nodes will be created in only these zones.
        :param pulumi.Input[list] node_pools: List of node pools associated with this cluster.
               See google_container_node_pool for schema.
               **Warning:** node pools defined inside a cluster can't be changed (or added/removed) after
               cluster creation without deleting and recreating the entire cluster. Unless you absolutely need the ability
               to say "these are the _only_ node pools associated with this cluster", use the
               google_container_node_pool resource instead of this property.
        :param pulumi.Input[dict] pod_security_policy_config: ) Configuration for the
               [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature.
               Structure is documented below.
        :param pulumi.Input[dict] private_cluster_config: A set of options for creating
               a private cluster. Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[bool] remove_default_node_pool: If `true`, deletes the default node
               pool upon cluster creation. If you're using `google_container_node_pool`
               resources with no default node pool, this should be set to `true`, alongside
               setting `initial_node_count` to at least `1`.
        :param pulumi.Input[dict] resource_labels: The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
        :param pulumi.Input[dict] resource_usage_export_config: ) Configuration for the
               [ResourceUsageExportConfig](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering) feature.
               Structure is documented below.
        :param pulumi.Input[str] subnetwork: The name or self_link of the Google Compute Engine subnetwork in
               which the cluster's instances are launched.
        :param pulumi.Input[dict] vertical_pod_autoscaling: )
               Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it.
               Structure is documented below.
        :param pulumi.Input[dict] workload_identity_config: )
               Workload Identity allows Kubernetes service accounts to act as a user-managed
               [Google IAM Service Account](https://cloud.google.com/iam/docs/service-accounts#user-managed_service_accounts).
        :param pulumi.Input[str] zone: The zone that the cluster master and nodes
               should be created in. If specified, this cluster will be a zonal cluster. `zone`
               has been deprecated in favour of `location`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/container_cluster.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['additional_zones'] = additional_zones

        __props__['addons_config'] = addons_config

        __props__['authenticator_groups_config'] = authenticator_groups_config

        __props__['cluster_autoscaling'] = cluster_autoscaling

        __props__['cluster_ipv4_cidr'] = cluster_ipv4_cidr

        __props__['database_encryption'] = database_encryption

        __props__['default_max_pods_per_node'] = default_max_pods_per_node

        __props__['description'] = description

        __props__['enable_binary_authorization'] = enable_binary_authorization

        __props__['enable_intranode_visibility'] = enable_intranode_visibility

        __props__['enable_kubernetes_alpha'] = enable_kubernetes_alpha

        __props__['enable_legacy_abac'] = enable_legacy_abac

        __props__['enable_tpu'] = enable_tpu

        __props__['initial_node_count'] = initial_node_count

        __props__['ip_allocation_policy'] = ip_allocation_policy

        __props__['location'] = location

        __props__['logging_service'] = logging_service

        __props__['maintenance_policy'] = maintenance_policy

        __props__['master_auth'] = master_auth

        __props__['master_authorized_networks_config'] = master_authorized_networks_config

        __props__['min_master_version'] = min_master_version

        __props__['monitoring_service'] = monitoring_service

        __props__['name'] = name

        __props__['network'] = network

        __props__['network_policy'] = network_policy

        __props__['node_config'] = node_config

        __props__['node_locations'] = node_locations

        __props__['node_pools'] = node_pools

        __props__['node_version'] = node_version

        __props__['pod_security_policy_config'] = pod_security_policy_config

        __props__['private_cluster_config'] = private_cluster_config

        __props__['project'] = project

        __props__['region'] = region

        __props__['remove_default_node_pool'] = remove_default_node_pool

        __props__['resource_labels'] = resource_labels

        __props__['resource_usage_export_config'] = resource_usage_export_config

        __props__['subnetwork'] = subnetwork

        __props__['vertical_pod_autoscaling'] = vertical_pod_autoscaling

        __props__['workload_identity_config'] = workload_identity_config

        __props__['zone'] = zone

        __props__['endpoint'] = None
        __props__['instance_group_urls'] = None
        __props__['master_version'] = None
        __props__['services_ipv4_cidr'] = None
        __props__['tpu_ipv4_cidr_block'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Cluster, __self__).__init__(
            'gcp:container/cluster:Cluster',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

