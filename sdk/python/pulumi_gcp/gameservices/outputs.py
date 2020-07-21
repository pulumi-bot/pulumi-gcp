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

@pulumi.output_type
class GameServerClusterConnectionInfo(dict):
    gke_cluster_reference: 'outputs.GameServerClusterConnectionInfoGkeClusterReference' = pulumi.output_property("gkeClusterReference")
    """
    Reference of the GKE cluster where the game servers are installed.  Structure is documented below.
    """
    namespace: str = pulumi.output_property("namespace")
    """
    Namespace designated on the game server cluster where the game server
    instances will be created. The namespace existence will be validated
    during creation.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerClusterConnectionInfoGkeClusterReference(dict):
    cluster: str = pulumi.output_property("cluster")
    """
    The full or partial name of a GKE cluster, using one of the following
    forms:
    * `projects/{project_id}/locations/{location}/clusters/{cluster_id}`
    * `locations/{location}/clusters/{cluster_id}`
    * `{cluster_id}`
    If project and location are not specified, the project and location of the
    GameServerCluster resource are used to generate the full name of the
    GKE cluster.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerConfigFleetConfig(dict):
    fleet_spec: str = pulumi.output_property("fleetSpec")
    """
    The fleet spec, which is sent to Agones to configure fleet.
    The spec can be passed as inline json but it is recommended to use a file reference
    instead. File references can contain the json or yaml format of the fleet spec. Eg:
    * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
    * fleet_spec = file("fleet_configs.json")
    The format of the spec can be found :
    `https://agones.dev/site/docs/reference/fleet/`.
    """
    name: Optional[str] = pulumi.output_property("name")
    """
    The name of the ScalingConfig
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerConfigScalingConfig(dict):
    fleet_autoscaler_spec: str = pulumi.output_property("fleetAutoscalerSpec")
    """
    Fleet autoscaler spec, which is sent to Agones.
    Example spec can be found :
    https://agones.dev/site/docs/reference/fleetautoscaler/
    """
    name: str = pulumi.output_property("name")
    """
    The name of the ScalingConfig
    """
    schedules: Optional[List['outputs.GameServerConfigScalingConfigSchedule']] = pulumi.output_property("schedules")
    """
    The schedules to which this scaling config applies.  Structure is documented below.
    """
    selectors: Optional[List['outputs.GameServerConfigScalingConfigSelector']] = pulumi.output_property("selectors")
    """
    Labels used to identify the clusters to which this scaling config
    applies. A cluster is subject to this scaling config if its labels match
    any of the selector entries.  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerConfigScalingConfigSchedule(dict):
    cron_job_duration: Optional[str] = pulumi.output_property("cronJobDuration")
    """
    The duration for the cron job event. The duration of the event is effective
    after the cron job's start time.
    A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
    """
    cron_spec: Optional[str] = pulumi.output_property("cronSpec")
    """
    The cron definition of the scheduled event. See
    https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
    defined by the realm.
    """
    end_time: Optional[str] = pulumi.output_property("endTime")
    """
    The end time of the event.
    A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
    """
    start_time: Optional[str] = pulumi.output_property("startTime")
    """
    The start time of the event.
    A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerConfigScalingConfigSelector(dict):
    labels: Optional[Dict[str, str]] = pulumi.output_property("labels")
    """
    Set of labels to group by.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerDeploymentRolloutGameServerConfigOverride(dict):
    config_version: Optional[str] = pulumi.output_property("configVersion")
    """
    Version of the configuration.
    """
    realms_selector: Optional['outputs.GameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector'] = pulumi.output_property("realmsSelector")
    """
    Selection by realms.  Structure is documented below.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector(dict):
    realms: Optional[List[str]] = pulumi.output_property("realms")
    """
    List of realms to match against.
    """

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetGameServerDeploymentRolloutGameServerConfigOverride(dict):
    config_version: str = pulumi.output_property("configVersion")
    realms_selectors: List['outputs.GetGameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector'] = pulumi.output_property("realmsSelectors")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetGameServerDeploymentRolloutGameServerConfigOverrideRealmsSelector(dict):
    realms: List[str] = pulumi.output_property("realms")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


