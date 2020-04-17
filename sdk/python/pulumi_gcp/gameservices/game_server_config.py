# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GameServerConfig(pulumi.CustomResource):
    config_id: pulumi.Output[str]
    """
    -
    (Required)
    A unique id for the deployment config.
    """
    deployment_id: pulumi.Output[str]
    """
    -
    (Required)
    A unique id for the deployment.
    """
    description: pulumi.Output[str]
    """
    -
    (Optional)
    The description of the game server config.
    """
    fleet_configs: pulumi.Output[list]
    """
    -
    (Required)
    The fleet config contains list of fleet specs. In the Single Cloud, there
    will be only one.  Structure is documented below.

      * `fleetSpec` (`str`) - -
        (Required)
        The fleet spec, which is sent to Agones to configure fleet.
        The spec can be passed as inline json but it is recommended to use a file reference
        instead. File references can contain the json or yaml format of the fleet spec. Eg:
        * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
        * fleet_spec = file("fleet_configs.json")
        The format of the spec can be found :
        `https://agones.dev/site/docs/reference/fleet/`.
      * `name` (`str`) - -
        (Required)
        The name of the FleetConfig.
    """
    labels: pulumi.Output[dict]
    """
    -
    (Optional)
    The labels associated with this game server config. Each label is a
    key-value pair.
    """
    location: pulumi.Output[str]
    """
    -
    (Optional)
    Location of the Deployment.
    """
    name: pulumi.Output[str]
    """
    -
    (Required)
    The name of the FleetConfig.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    scaling_configs: pulumi.Output[list]
    """
    -
    (Optional)
    Optional. This contains the autoscaling settings.  Structure is documented below.

      * `fleetAutoscalerSpec` (`str`) - -
        (Required)
        Fleet autoscaler spec, which is sent to Agones.
        Example spec can be found :
        https://agones.dev/site/docs/reference/fleetautoscaler/
      * `name` (`str`) - -
        (Required)
        The name of the FleetConfig.
      * `schedules` (`list`) - -
        (Optional)
        The schedules to which this scaling config applies.  Structure is documented below.
        * `cronJobDuration` (`str`) - -
          (Optional)
          The duration for the cron job event. The duration of the event is effective
          after the cron job's start time.
          A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        * `cronSpec` (`str`) - -
          (Optional)
          The cron definition of the scheduled event. See
          https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
          defined by the realm.
        * `endTime` (`str`) - -
          (Optional)
          The end time of the event.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        * `startTime` (`str`) - -
          (Optional)
          The start time of the event.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

      * `selectors` (`list`) - -
        (Optional)
        Labels used to identify the clusters to which this scaling config
        applies. A cluster is subject to this scaling config if its labels match
        any of the selector entries.  Structure is documented below.
        * `labels` (`dict`) - -
          (Optional)
          The labels associated with this game server config. Each label is a
          key-value pair.
    """
    def __init__(__self__, resource_name, opts=None, config_id=None, deployment_id=None, description=None, fleet_configs=None, labels=None, location=None, project=None, scaling_configs=None, __props__=None, __name__=None, __opts__=None):
        """
        A game server config resource. Configs are global and immutable.

        To get more information about GameServerConfig, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: -
               (Required)
               A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: -
               (Required)
               A unique id for the deployment.
        :param pulumi.Input[str] description: -
               (Optional)
               The description of the game server config.
        :param pulumi.Input[list] fleet_configs: -
               (Required)
               The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.  Structure is documented below.
        :param pulumi.Input[dict] labels: -
               (Optional)
               The labels associated with this game server config. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: -
               (Optional)
               Location of the Deployment.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] scaling_configs: -
               (Optional)
               Optional. This contains the autoscaling settings.  Structure is documented below.

        The **fleet_configs** object supports the following:

          * `fleetSpec` (`pulumi.Input[str]`) - -
            (Required)
            The fleet spec, which is sent to Agones to configure fleet.
            The spec can be passed as inline json but it is recommended to use a file reference
            instead. File references can contain the json or yaml format of the fleet spec. Eg:
            * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
            * fleet_spec = file("fleet_configs.json")
            The format of the spec can be found :
            `https://agones.dev/site/docs/reference/fleet/`.
          * `name` (`pulumi.Input[str]`) - -
            (Required)
            The name of the FleetConfig.

        The **scaling_configs** object supports the following:

          * `fleetAutoscalerSpec` (`pulumi.Input[str]`) - -
            (Required)
            Fleet autoscaler spec, which is sent to Agones.
            Example spec can be found :
            https://agones.dev/site/docs/reference/fleetautoscaler/
          * `name` (`pulumi.Input[str]`) - -
            (Required)
            The name of the FleetConfig.
          * `schedules` (`pulumi.Input[list]`) - -
            (Optional)
            The schedules to which this scaling config applies.  Structure is documented below.
            * `cronJobDuration` (`pulumi.Input[str]`) - -
              (Optional)
              The duration for the cron job event. The duration of the event is effective
              after the cron job's start time.
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `cronSpec` (`pulumi.Input[str]`) - -
              (Optional)
              The cron definition of the scheduled event. See
              https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
              defined by the realm.
            * `endTime` (`pulumi.Input[str]`) - -
              (Optional)
              The end time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
            * `startTime` (`pulumi.Input[str]`) - -
              (Optional)
              The start time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

          * `selectors` (`pulumi.Input[list]`) - -
            (Optional)
            Labels used to identify the clusters to which this scaling config
            applies. A cluster is subject to this scaling config if its labels match
            any of the selector entries.  Structure is documented below.
            * `labels` (`pulumi.Input[dict]`) - -
              (Optional)
              The labels associated with this game server config. Each label is a
              key-value pair.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if config_id is None:
                raise TypeError("Missing required property 'config_id'")
            __props__['config_id'] = config_id
            if deployment_id is None:
                raise TypeError("Missing required property 'deployment_id'")
            __props__['deployment_id'] = deployment_id
            __props__['description'] = description
            if fleet_configs is None:
                raise TypeError("Missing required property 'fleet_configs'")
            __props__['fleet_configs'] = fleet_configs
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['scaling_configs'] = scaling_configs
            __props__['name'] = None
        super(GameServerConfig, __self__).__init__(
            'gcp:gameservices/gameServerConfig:GameServerConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, config_id=None, deployment_id=None, description=None, fleet_configs=None, labels=None, location=None, name=None, project=None, scaling_configs=None):
        """
        Get an existing GameServerConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: -
               (Required)
               A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: -
               (Required)
               A unique id for the deployment.
        :param pulumi.Input[str] description: -
               (Optional)
               The description of the game server config.
        :param pulumi.Input[list] fleet_configs: -
               (Required)
               The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.  Structure is documented below.
        :param pulumi.Input[dict] labels: -
               (Optional)
               The labels associated with this game server config. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: -
               (Optional)
               Location of the Deployment.
        :param pulumi.Input[str] name: -
               (Required)
               The name of the FleetConfig.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] scaling_configs: -
               (Optional)
               Optional. This contains the autoscaling settings.  Structure is documented below.

        The **fleet_configs** object supports the following:

          * `fleetSpec` (`pulumi.Input[str]`) - -
            (Required)
            The fleet spec, which is sent to Agones to configure fleet.
            The spec can be passed as inline json but it is recommended to use a file reference
            instead. File references can contain the json or yaml format of the fleet spec. Eg:
            * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
            * fleet_spec = file("fleet_configs.json")
            The format of the spec can be found :
            `https://agones.dev/site/docs/reference/fleet/`.
          * `name` (`pulumi.Input[str]`) - -
            (Required)
            The name of the FleetConfig.

        The **scaling_configs** object supports the following:

          * `fleetAutoscalerSpec` (`pulumi.Input[str]`) - -
            (Required)
            Fleet autoscaler spec, which is sent to Agones.
            Example spec can be found :
            https://agones.dev/site/docs/reference/fleetautoscaler/
          * `name` (`pulumi.Input[str]`) - -
            (Required)
            The name of the FleetConfig.
          * `schedules` (`pulumi.Input[list]`) - -
            (Optional)
            The schedules to which this scaling config applies.  Structure is documented below.
            * `cronJobDuration` (`pulumi.Input[str]`) - -
              (Optional)
              The duration for the cron job event. The duration of the event is effective
              after the cron job's start time.
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `cronSpec` (`pulumi.Input[str]`) - -
              (Optional)
              The cron definition of the scheduled event. See
              https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
              defined by the realm.
            * `endTime` (`pulumi.Input[str]`) - -
              (Optional)
              The end time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
            * `startTime` (`pulumi.Input[str]`) - -
              (Optional)
              The start time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

          * `selectors` (`pulumi.Input[list]`) - -
            (Optional)
            Labels used to identify the clusters to which this scaling config
            applies. A cluster is subject to this scaling config if its labels match
            any of the selector entries.  Structure is documented below.
            * `labels` (`pulumi.Input[dict]`) - -
              (Optional)
              The labels associated with this game server config. Each label is a
              key-value pair.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["config_id"] = config_id
        __props__["deployment_id"] = deployment_id
        __props__["description"] = description
        __props__["fleet_configs"] = fleet_configs
        __props__["labels"] = labels
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["scaling_configs"] = scaling_configs
        return GameServerConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

