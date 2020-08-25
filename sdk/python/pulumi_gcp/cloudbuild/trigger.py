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

__all__ = ['Trigger']


class Trigger(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[pulumi.InputType['TriggerBuildArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 github: Optional[pulumi.Input[pulumi.InputType['TriggerGithubArgs']]] = None,
                 ignored_files: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 included_files: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 trigger_template: Optional[pulumi.Input[pulumi.InputType['TriggerTriggerTemplateArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configuration for an automated build in response to source repository changes.

        To get more information about Trigger, see:

        * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/)
        * How-to Guides
            * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TriggerBuildArgs']] build: Contents of the build template. Either a filename or build template must be provided.
               Structure is documented below.
        :param pulumi.Input[str] description: Human-readable description of the trigger.
        :param pulumi.Input[bool] disabled: Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        :param pulumi.Input[str] filename: Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        :param pulumi.Input[pulumi.InputType['TriggerGithubArgs']] github: Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] ignored_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If ignoredFiles and changed files are both empty, then they are not
               used to determine whether or not to trigger a build.
               If ignoredFiles is not empty, then we ignore any files that match any
               of the ignored_file globs. If the change has no files that are outside
               of the ignoredFiles globs, then we do not trigger a build.
        :param pulumi.Input[List[pulumi.Input[str]]] included_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If any of the files altered in the commit pass the ignoredFiles filter
               and includedFiles is empty, then as far as this filter is concerned, we
               should trigger the build.
               If any of the files altered in the commit pass the ignoredFiles filter
               and includedFiles is not empty, then we make sure that at least one of
               those files matches a includedFiles glob. If not, then we do not trigger
               a build.
        :param pulumi.Input[str] name: Name of the volume to mount.
               Volume names must be unique per build step and must be valid names for
               Docker volumes. Each named volume must be used by at least two build steps.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] substitutions: Substitutions to use in a triggered build. Should only be used with triggers.run
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags for annotation of a Build. These are not docker tags.
        :param pulumi.Input[pulumi.InputType['TriggerTriggerTemplateArgs']] trigger_template: Template describing the types of source changes to trigger a build.
               Branch and tag names in trigger templates are interpreted as regular
               expressions. Any branch or tag change that matches that regular
               expression will trigger a build.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
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

            __props__['build'] = build
            __props__['description'] = description
            __props__['disabled'] = disabled
            __props__['filename'] = filename
            __props__['github'] = github
            __props__['ignored_files'] = ignored_files
            __props__['included_files'] = included_files
            __props__['name'] = name
            __props__['project'] = project
            __props__['substitutions'] = substitutions
            __props__['tags'] = tags
            __props__['trigger_template'] = trigger_template
            __props__['create_time'] = None
            __props__['trigger_id'] = None
        super(Trigger, __self__).__init__(
            'gcp:cloudbuild/trigger:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            build: Optional[pulumi.Input[pulumi.InputType['TriggerBuildArgs']]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            filename: Optional[pulumi.Input[str]] = None,
            github: Optional[pulumi.Input[pulumi.InputType['TriggerGithubArgs']]] = None,
            ignored_files: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            included_files: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            substitutions: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            trigger_id: Optional[pulumi.Input[str]] = None,
            trigger_template: Optional[pulumi.Input[pulumi.InputType['TriggerTriggerTemplateArgs']]] = None) -> 'Trigger':
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TriggerBuildArgs']] build: Contents of the build template. Either a filename or build template must be provided.
               Structure is documented below.
        :param pulumi.Input[str] create_time: Time when the trigger was created.
        :param pulumi.Input[str] description: Human-readable description of the trigger.
        :param pulumi.Input[bool] disabled: Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        :param pulumi.Input[str] filename: Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        :param pulumi.Input[pulumi.InputType['TriggerGithubArgs']] github: Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] ignored_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If ignoredFiles and changed files are both empty, then they are not
               used to determine whether or not to trigger a build.
               If ignoredFiles is not empty, then we ignore any files that match any
               of the ignored_file globs. If the change has no files that are outside
               of the ignoredFiles globs, then we do not trigger a build.
        :param pulumi.Input[List[pulumi.Input[str]]] included_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If any of the files altered in the commit pass the ignoredFiles filter
               and includedFiles is empty, then as far as this filter is concerned, we
               should trigger the build.
               If any of the files altered in the commit pass the ignoredFiles filter
               and includedFiles is not empty, then we make sure that at least one of
               those files matches a includedFiles glob. If not, then we do not trigger
               a build.
        :param pulumi.Input[str] name: Name of the volume to mount.
               Volume names must be unique per build step and must be valid names for
               Docker volumes. Each named volume must be used by at least two build steps.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] substitutions: Substitutions to use in a triggered build. Should only be used with triggers.run
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags for annotation of a Build. These are not docker tags.
        :param pulumi.Input[str] trigger_id: The unique identifier for the trigger.
        :param pulumi.Input[pulumi.InputType['TriggerTriggerTemplateArgs']] trigger_template: Template describing the types of source changes to trigger a build.
               Branch and tag names in trigger templates are interpreted as regular
               expressions. Any branch or tag change that matches that regular
               expression will trigger a build.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["build"] = build
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["filename"] = filename
        __props__["github"] = github
        __props__["ignored_files"] = ignored_files
        __props__["included_files"] = included_files
        __props__["name"] = name
        __props__["project"] = project
        __props__["substitutions"] = substitutions
        __props__["tags"] = tags
        __props__["trigger_id"] = trigger_id
        __props__["trigger_template"] = trigger_template
        return Trigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def build(self) -> Optional['outputs.TriggerBuild']:
        """
        Contents of the build template. Either a filename or build template must be provided.
        Structure is documented below.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the trigger was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Human-readable description of the trigger.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        """
        Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filename(self) -> Optional[str]:
        """
        Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def github(self) -> Optional['outputs.TriggerGithub']:
        """
        Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
        One of `trigger_template` or `github` must be provided.
        Structure is documented below.
        """
        return pulumi.get(self, "github")

    @property
    @pulumi.getter(name="ignoredFiles")
    def ignored_files(self) -> Optional[List[str]]:
        """
        ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        extended with support for `**`.
        If ignoredFiles and changed files are both empty, then they are not
        used to determine whether or not to trigger a build.
        If ignoredFiles is not empty, then we ignore any files that match any
        of the ignored_file globs. If the change has no files that are outside
        of the ignoredFiles globs, then we do not trigger a build.
        """
        return pulumi.get(self, "ignored_files")

    @property
    @pulumi.getter(name="includedFiles")
    def included_files(self) -> Optional[List[str]]:
        """
        ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
        extended with support for `**`.
        If any of the files altered in the commit pass the ignoredFiles filter
        and includedFiles is empty, then as far as this filter is concerned, we
        should trigger the build.
        If any of the files altered in the commit pass the ignoredFiles filter
        and includedFiles is not empty, then we make sure that at least one of
        those files matches a includedFiles glob. If not, then we do not trigger
        a build.
        """
        return pulumi.get(self, "included_files")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the volume to mount.
        Volume names must be unique per build step and must be valid names for
        Docker volumes. Each named volume must be used by at least two build steps.
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

    @property
    @pulumi.getter
    def substitutions(self) -> Optional[Mapping[str, str]]:
        """
        Substitutions to use in a triggered build. Should only be used with triggers.run
        """
        return pulumi.get(self, "substitutions")

    @property
    @pulumi.getter
    def tags(self) -> Optional[List[str]]:
        """
        Tags for annotation of a Build. These are not docker tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> str:
        """
        The unique identifier for the trigger.
        """
        return pulumi.get(self, "trigger_id")

    @property
    @pulumi.getter(name="triggerTemplate")
    def trigger_template(self) -> Optional['outputs.TriggerTriggerTemplate']:
        """
        Template describing the types of source changes to trigger a build.
        Branch and tag names in trigger templates are interpreted as regular
        expressions. Any branch or tag change that matches that regular
        expression will trigger a build.
        One of `trigger_template` or `github` must be provided.
        Structure is documented below.
        """
        return pulumi.get(self, "trigger_template")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

