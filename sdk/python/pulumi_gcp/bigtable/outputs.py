# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GCPolicyMaxAge',
    'GCPolicyMaxVersion',
    'InstanceCluster',
    'InstanceIamBindingCondition',
    'InstanceIamMemberCondition',
    'TableColumnFamily',
]

@pulumi.output_type
class GCPolicyMaxAge(dict):
    days: float = pulumi.output_property("days")
    """
    Number of days before applying GC policy.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GCPolicyMaxVersion(dict):
    number: float = pulumi.output_property("number")
    """
    Number of version before applying the GC policy.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceCluster(dict):
    cluster_id: str = pulumi.output_property("clusterId")
    """
    The ID of the Cloud Bigtable cluster.
    """
    num_nodes: Optional[float] = pulumi.output_property("numNodes")
    """
    The number of nodes in your Cloud Bigtable cluster.
    Required, with a minimum of `1` for a `PRODUCTION` instance. Must be left unset
    for a `DEVELOPMENT` instance.
    """
    storage_type: Optional[str] = pulumi.output_property("storageType")
    """
    The storage type to use. One of `"SSD"` or
    `"HDD"`. Defaults to `"SSD"`.
    """
    zone: str = pulumi.output_property("zone")
    """
    The zone to create the Cloud Bigtable cluster in. Each
    cluster must have a different zone in the same region. Zones that support
    Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceIamBindingCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class InstanceIamMemberCondition(dict):
    description: Optional[str] = pulumi.output_property("description")
    expression: str = pulumi.output_property("expression")
    title: str = pulumi.output_property("title")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TableColumnFamily(dict):
    family: str = pulumi.output_property("family")
    """
    The name of the column family.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


