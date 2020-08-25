# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['GameServerDeploymentRollout']


class GameServerDeploymentRollout(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_game_server_config: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 game_server_config_overrides: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GameServerDeploymentRolloutGameServerConfigOverrideArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This represents the rollout state. This is part of the game server
        deployment.

        To get more information about GameServerDeploymentRollout, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/GameServerDeploymentRollout)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_game_server_config: This field points to the game server config that is
               applied by default to all realms and clusters. For example,
               `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
        :param pulumi.Input[str] deployment_id: The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GameServerDeploymentRolloutGameServerConfigOverrideArgs']]]] game_server_config_overrides: The game_server_config_overrides contains the per game server config
               overrides. The overrides are processed in the order they are listed. As
               soon as a match is found for a cluster, the rest of the list is not
               processed.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            if default_game_server_config is None:
                raise TypeError("Missing required property 'default_game_server_config'")
            __props__['default_game_server_config'] = default_game_server_config
            if deployment_id is None:
                raise TypeError("Missing required property 'deployment_id'")
            __props__['deployment_id'] = deployment_id
            __props__['game_server_config_overrides'] = game_server_config_overrides
            __props__['project'] = project
            __props__['name'] = None
        super(GameServerDeploymentRollout, __self__).__init__(
            'gcp:gameservices/gameServerDeploymentRollout:GameServerDeploymentRollout',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_game_server_config: Optional[pulumi.Input[str]] = None,
            deployment_id: Optional[pulumi.Input[str]] = None,
            game_server_config_overrides: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GameServerDeploymentRolloutGameServerConfigOverrideArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'GameServerDeploymentRollout':
        """
        Get an existing GameServerDeploymentRollout resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_game_server_config: This field points to the game server config that is
               applied by default to all realms and clusters. For example,
               `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
        :param pulumi.Input[str] deployment_id: The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GameServerDeploymentRolloutGameServerConfigOverrideArgs']]]] game_server_config_overrides: The game_server_config_overrides contains the per game server config
               overrides. The overrides are processed in the order they are listed. As
               soon as a match is found for a cluster, the rest of the list is not
               processed.
               Structure is documented below.
        :param pulumi.Input[str] name: The resource id of the game server deployment eg:
               'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_game_server_config"] = default_game_server_config
        __props__["deployment_id"] = deployment_id
        __props__["game_server_config_overrides"] = game_server_config_overrides
        __props__["name"] = name
        __props__["project"] = project
        return GameServerDeploymentRollout(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultGameServerConfig")
    def default_game_server_config(self) -> str:
        """
        This field points to the game server config that is
        applied by default to all realms and clusters. For example,
        `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
        """
        return pulumi.get(self, "default_game_server_config")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> str:
        """
        The deployment to rollout the new config to. Only 1 rollout must be associated with each deployment.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter(name="gameServerConfigOverrides")
    def game_server_config_overrides(self) -> Optional[List['outputs.GameServerDeploymentRolloutGameServerConfigOverride']]:
        """
        The game_server_config_overrides contains the per game server config
        overrides. The overrides are processed in the order they are listed. As
        soon as a match is found for a cluster, the rest of the list is not
        processed.
        Structure is documented below.
        """
        return pulumi.get(self, "game_server_config_overrides")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource id of the game server deployment eg:
        'projects/my-project/locations/global/gameServerDeployments/my-deployment/rollout'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

