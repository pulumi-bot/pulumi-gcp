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
    'MembershipAuthority',
    'MembershipEndpoint',
    'MembershipEndpointGkeCluster',
]

@pulumi.output_type
class MembershipAuthority(dict):
    def __init__(__self__, *,
                 issuer: str):
        """
        :param str issuer: A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
               with length <2000 characters.
        """
        pulumi.set(__self__, "issuer", issuer)

    @property
    @pulumi.getter
    def issuer(self) -> str:
        """
        A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and // be a valid
        with length <2000 characters.
        """
        return pulumi.get(self, "issuer")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MembershipEndpoint(dict):
    def __init__(__self__, *,
                 gke_cluster: Optional['outputs.MembershipEndpointGkeCluster'] = None):
        """
        :param 'MembershipEndpointGkeClusterArgs' gke_cluster: If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
               Structure is documented below.
        """
        if gke_cluster is not None:
            pulumi.set(__self__, "gke_cluster", gke_cluster)

    @property
    @pulumi.getter(name="gkeCluster")
    def gke_cluster(self) -> Optional['outputs.MembershipEndpointGkeCluster']:
        """
        If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
        Structure is documented below.
        """
        return pulumi.get(self, "gke_cluster")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MembershipEndpointGkeCluster(dict):
    def __init__(__self__, *,
                 resource_link: str):
        pulumi.set(__self__, "resource_link", resource_link)

    @property
    @pulumi.getter(name="resourceLink")
    def resource_link(self) -> str:
        return pulumi.get(self, "resource_link")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


