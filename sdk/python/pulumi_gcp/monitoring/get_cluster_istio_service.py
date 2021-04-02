# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetClusterIstioServiceResult',
    'AwaitableGetClusterIstioServiceResult',
    'get_cluster_istio_service',
]

@pulumi.output_type
class GetClusterIstioServiceResult:
    """
    A collection of values returned by getClusterIstioService.
    """
    def __init__(__self__, cluster_name=None, display_name=None, id=None, location=None, name=None, project=None, service_id=None, service_name=None, service_namespace=None, telemetries=None):
        if cluster_name and not isinstance(cluster_name, str):
            raise TypeError("Expected argument 'cluster_name' to be a str")
        pulumi.set(__self__, "cluster_name", cluster_name)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if service_namespace and not isinstance(service_namespace, str):
            raise TypeError("Expected argument 'service_namespace' to be a str")
        pulumi.set(__self__, "service_namespace", service_namespace)
        if telemetries and not isinstance(telemetries, list):
            raise TypeError("Expected argument 'telemetries' to be a list")
        pulumi.set(__self__, "telemetries", telemetries)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> str:
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceNamespace")
    def service_namespace(self) -> str:
        return pulumi.get(self, "service_namespace")

    @property
    @pulumi.getter
    def telemetries(self) -> Sequence['outputs.GetClusterIstioServiceTelemetryResult']:
        return pulumi.get(self, "telemetries")


class AwaitableGetClusterIstioServiceResult(GetClusterIstioServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterIstioServiceResult(
            cluster_name=self.cluster_name,
            display_name=self.display_name,
            id=self.id,
            location=self.location,
            name=self.name,
            project=self.project,
            service_id=self.service_id,
            service_name=self.service_name,
            service_namespace=self.service_namespace,
            telemetries=self.telemetries)


def get_cluster_istio_service(cluster_name: Optional[str] = None,
                              location: Optional[str] = None,
                              project: Optional[str] = None,
                              service_name: Optional[str] = None,
                              service_namespace: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterIstioServiceResult:
    """
    A Monitoring Service is the root resource under which operational aspects of a
    generic service are accessible. A service is some discrete, autonomous, and
    network-accessible unit, designed to solve an individual concern

    An Cluster Istio monitoring service is automatically created by GCP to monitor
    Cluster Istio services.

    To get more information about Service, see:

    * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
    * How-to Guides
        * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

    ## Example Usage
    ### Monitoring Cluster Istio Service

    ```python
    import pulumi
    import pulumi_gcp as gcp

    default = gcp.monitoring.get_cluster_istio_service(cluster_name="west",
        location="us-west2-a",
        service_name="istio-policy",
        service_namespace="istio-system")
    ```


    :param str cluster_name: The name of the Kubernetes cluster in which this Istio service 
           is defined. Corresponds to the clusterName resource label in k8s_cluster resources.
    :param str location: The location of the Kubernetes cluster in which this Istio service 
           is defined. Corresponds to the location resource label in k8s_cluster resources.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    :param str service_name: The name of the Istio service underlying this service.
           Corresponds to the destination_service_name metric label in Istio metrics.
    :param str service_namespace: The namespace of the Istio service underlying this service.
           Corresponds to the destination_service_namespace metric label in Istio metrics.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['location'] = location
    __args__['project'] = project
    __args__['serviceName'] = service_name
    __args__['serviceNamespace'] = service_namespace
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:monitoring/getClusterIstioService:getClusterIstioService', __args__, opts=opts, typ=GetClusterIstioServiceResult).value

    return AwaitableGetClusterIstioServiceResult(
        cluster_name=__ret__.cluster_name,
        display_name=__ret__.display_name,
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        project=__ret__.project,
        service_id=__ret__.service_id,
        service_name=__ret__.service_name,
        service_namespace=__ret__.service_namespace,
        telemetries=__ret__.telemetries)
