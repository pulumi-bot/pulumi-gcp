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
    A unique id for the deployment config.
    """
    deployment_id: pulumi.Output[str]
    """
    A unique id for the deployment.
    """
    description: pulumi.Output[str]
    """
    The description of the game server config.
    """
    fleet_configs: pulumi.Output[list]
    """
    The fleet config contains list of fleet specs. In the Single Cloud, there
    will be only one.  Structure is documented below.

      * `fleetSpec` (`str`) - The fleet spec, which is sent to Agones to configure fleet.
        The spec can be passed as inline json but it is recommended to use a file reference
        instead. File references can contain the json or yaml format of the fleet spec. Eg:
        * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
        * fleet_spec = file("fleet_configs.json")
        The format of the spec can be found :
        `https://agones.dev/site/docs/reference/fleet/`.
      * `name` (`str`) - The name of the ScalingConfig
    """
    labels: pulumi.Output[dict]
    """
    Set of labels to group by.
    """
    location: pulumi.Output[str]
    """
    Location of the Deployment.
    """
    name: pulumi.Output[str]
    """
    The name of the ScalingConfig
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    scaling_configs: pulumi.Output[list]
    """
    Optional. This contains the autoscaling settings.  Structure is documented below.

      * `fleetAutoscalerSpec` (`str`) - Fleet autoscaler spec, which is sent to Agones.
        Example spec can be found :
        https://agones.dev/site/docs/reference/fleetautoscaler/
      * `name` (`str`) - The name of the ScalingConfig
      * `schedules` (`list`) - The schedules to which this scaling config applies.  Structure is documented below.
        * `cronJobDuration` (`str`) - The duration for the cron job event. The duration of the event is effective
          after the cron job's start time.
          A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        * `cronSpec` (`str`) - The cron definition of the scheduled event. See
          https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
          defined by the realm.
        * `endTime` (`str`) - The end time of the event.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        * `startTime` (`str`) - The start time of the event.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

      * `selectors` (`list`) - Labels used to identify the clusters to which this scaling config
        applies. A cluster is subject to this scaling config if its labels match
        any of the selector entries.  Structure is documented below.
        * `labels` (`dict`) - Set of labels to group by.
    """
    def __init__(__self__, resource_name, opts=None, config_id=None, deployment_id=None, description=None, fleet_configs=None, labels=None, location=None, project=None, scaling_configs=None, __props__=None, __name__=None, __opts__=None):
        """
        A game server config resource. Configs are global and immutable.

        To get more information about GameServerConfig, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage - Game Service Config Basic


        ```python
        import pulumi
        import json
        import pulumi_gcp as gcp

        default_game_server_deployment = gcp.gameservices.GameServerDeployment("defaultGameServerDeployment",
            deployment_id="tf-test-deployment",
            description="a deployment description")
        default_game_server_config = gcp.gameservices.GameServerConfig("defaultGameServerConfig",
            config_id="tf-test-config",
            deployment_id=default_game_server_deployment.deployment_id,
            description="a config description",
            fleet_configs=[{
                "name": "something-unique",
                "fleetSpec": json.dumps({
                    "replicas": 1,
                    "scheduling": "Packed",
                    "template": {
                        "metadata": {
                            "name": "tf-test-game-server-template",
                        },
                        "spec": {
                            "template": {
                                "spec": {
                                    "containers": [{
                                        "name": "simple-udp-server",
                                        "image": "gcr.io/agones-images/udp-server:0.14",
                                    }],
                                },
                            },
                        },
                    },
                }),
            }],
            scaling_configs=[{
                "name": "scaling-config-name",
                "fleetAutoscalerSpec": json.dumps({
                    "policy": {
                        "type": "Webhook",
                        "webhook": {
                            "service": {
                                "name": "autoscaler-webhook-service",
                                "namespace": "default",
                                "path": "scale",
                            },
                        },
                    },
                }),
                "selectors": [{
                    "labels": {
                        "one": "two",
                    },
                }],
                "schedules": [{
                    "cronJobDuration": "3.500s",
                    "cronSpec": "0 0 * * 0",
                }],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[list] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.  Structure is documented below.
        :param pulumi.Input[dict] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] scaling_configs: Optional. This contains the autoscaling settings.  Structure is documented below.

        The **fleet_configs** object supports the following:

          * `fleetSpec` (`pulumi.Input[str]`) - The fleet spec, which is sent to Agones to configure fleet.
            The spec can be passed as inline json but it is recommended to use a file reference
            instead. File references can contain the json or yaml format of the fleet spec. Eg:
            * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
            * fleet_spec = file("fleet_configs.json")
            The format of the spec can be found :
            `https://agones.dev/site/docs/reference/fleet/`.
          * `name` (`pulumi.Input[str]`) - The name of the ScalingConfig

        The **scaling_configs** object supports the following:

          * `fleetAutoscalerSpec` (`pulumi.Input[str]`) - Fleet autoscaler spec, which is sent to Agones.
            Example spec can be found :
            https://agones.dev/site/docs/reference/fleetautoscaler/
          * `name` (`pulumi.Input[str]`) - The name of the ScalingConfig
          * `schedules` (`pulumi.Input[list]`) - The schedules to which this scaling config applies.  Structure is documented below.
            * `cronJobDuration` (`pulumi.Input[str]`) - The duration for the cron job event. The duration of the event is effective
              after the cron job's start time.
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `cronSpec` (`pulumi.Input[str]`) - The cron definition of the scheduled event. See
              https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
              defined by the realm.
            * `endTime` (`pulumi.Input[str]`) - The end time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
            * `startTime` (`pulumi.Input[str]`) - The start time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

          * `selectors` (`pulumi.Input[list]`) - Labels used to identify the clusters to which this scaling config
            applies. A cluster is subject to this scaling config if its labels match
            any of the selector entries.  Structure is documented below.
            * `labels` (`pulumi.Input[dict]`) - Set of labels to group by.
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
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[list] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.  Structure is documented below.
        :param pulumi.Input[dict] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] name: The name of the ScalingConfig
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[list] scaling_configs: Optional. This contains the autoscaling settings.  Structure is documented below.

        The **fleet_configs** object supports the following:

          * `fleetSpec` (`pulumi.Input[str]`) - The fleet spec, which is sent to Agones to configure fleet.
            The spec can be passed as inline json but it is recommended to use a file reference
            instead. File references can contain the json or yaml format of the fleet spec. Eg:
            * fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
            * fleet_spec = file("fleet_configs.json")
            The format of the spec can be found :
            `https://agones.dev/site/docs/reference/fleet/`.
          * `name` (`pulumi.Input[str]`) - The name of the ScalingConfig

        The **scaling_configs** object supports the following:

          * `fleetAutoscalerSpec` (`pulumi.Input[str]`) - Fleet autoscaler spec, which is sent to Agones.
            Example spec can be found :
            https://agones.dev/site/docs/reference/fleetautoscaler/
          * `name` (`pulumi.Input[str]`) - The name of the ScalingConfig
          * `schedules` (`pulumi.Input[list]`) - The schedules to which this scaling config applies.  Structure is documented below.
            * `cronJobDuration` (`pulumi.Input[str]`) - The duration for the cron job event. The duration of the event is effective
              after the cron job's start time.
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            * `cronSpec` (`pulumi.Input[str]`) - The cron definition of the scheduled event. See
              https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
              defined by the realm.
            * `endTime` (`pulumi.Input[str]`) - The end time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
            * `startTime` (`pulumi.Input[str]`) - The start time of the event.
              A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".

          * `selectors` (`pulumi.Input[list]`) - Labels used to identify the clusters to which this scaling config
            applies. A cluster is subject to this scaling config if its labels match
            any of the selector entries.  Structure is documented below.
            * `labels` (`pulumi.Input[dict]`) - Set of labels to group by.
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

