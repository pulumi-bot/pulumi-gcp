# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Trigger(pulumi.CustomResource):
    build: pulumi.Output[dict]
    """
    Contents of the build template. Either a filename or build template must be provided.
    Structure is documented below.

      * `images` (`list`) - A list of images to be pushed upon the successful completion of all build steps.
        The images are pushed using the builder service account's credentials.
        The digests of the pushed images will be stored in the Build resource's results field.
        If any of the images fail to be pushed, the build status is marked FAILURE.
      * `steps` (`list`) - The operations to be performed on the workspace.
        Structure is documented below.
        * `args` (`list`) - A list of arguments that will be presented to the step when it is started.
          If the image used to run the step's container has an entrypoint, the args
          are used as arguments to that entrypoint. If the image does not define an
          entrypoint, the first element in args is used as the entrypoint, and the
          remainder will be used as arguments.
        * `dir` (`str`) - Working directory to use when running this step's container.
          If this value is a relative path, it is relative to the build's working
          directory. If this value is absolute, it may be outside the build's working
          directory, in which case the contents of the path may not be persisted
          across build step executions, unless a `volume` for that path is specified.
          If the build specifies a `RepoSource` with `dir` and a step with a
          `dir`,
          which specifies an absolute path, the `RepoSource` `dir` is ignored
          for the step's execution.
        * `entrypoint` (`str`) - Entrypoint to be used instead of the build step image's
          default entrypoint.
          If unset, the image's default entrypoint is used
        * `envs` (`list`) - A list of environment variable definitions to be used when
          running a step.
          The elements are of the form "KEY=VALUE" for the environment variable
          "KEY" being given the value "VALUE".
        * `id` (`str`) - Unique identifier for this build step, used in `wait_for` to
          reference this build step as a dependency.
        * `name` (`str`) - Name of the volume to mount.
          Volume names must be unique per build step and must be valid names for
          Docker volumes. Each named volume must be used by at least two build steps.
        * `secretEnvs` (`list`) - A list of environment variables which are encrypted using
          a Cloud Key
          Management Service crypto key. These values must be specified in
          the build's `Secret`.
        * `timeout` (`str`) - Time limit for executing this build step. If not defined,
          the step has no
          time limit and will be allowed to continue to run until either it
          completes or the build itself times out.
        * `timing` (`str`) - Output only. Stores timing information for executing this
          build step.
        * `volumes` (`list`) - List of volumes to mount into the build step.
          Each volume is created as an empty volume prior to execution of the
          build step. Upon completion of the build, volumes and their contents
          are discarded.
          Using a named volume in only one step is not valid as it is
          indicative of a build request with an incorrect configuration.
          Structure is documented below.
          * `name` (`str`) - Name of the volume to mount.
            Volume names must be unique per build step and must be valid names for
            Docker volumes. Each named volume must be used by at least two build steps.
          * `path` (`str`) - Path at which to mount the volume.
            Paths must be absolute and cannot conflict with other volume paths on
            the same build step or with certain reserved volume paths.

        * `waitFors` (`list`) - The ID(s) of the step(s) that this build step depends on.
          This build step will not start until all the build steps in `wait_for`
          have completed successfully. If `wait_for` is empty, this build step
          will start when all previous build steps in the `Build.Steps` list
          have completed successfully.

      * `tags` (`list`) - Tags for annotation of a Build. These are not docker tags.
      * `timeout` (`str`) - Time limit for executing this build step. If not defined,
        the step has no
        time limit and will be allowed to continue to run until either it
        completes or the build itself times out.
    """
    create_time: pulumi.Output[str]
    """
    Time when the trigger was created.
    """
    description: pulumi.Output[str]
    """
    Human-readable description of the trigger.
    """
    disabled: pulumi.Output[bool]
    """
    Whether the trigger is disabled or not. If true, the trigger will never result in a build.
    """
    filename: pulumi.Output[str]
    """
    Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
    """
    github: pulumi.Output[dict]
    """
    Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
    One of `trigger_template` or `github` must be provided.
    Structure is documented below.

      * `name` (`str`) - Name of the volume to mount.
        Volume names must be unique per build step and must be valid names for
        Docker volumes. Each named volume must be used by at least two build steps.
      * `owner` (`str`) - Owner of the repository. For example: The owner for
        https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
      * `pullRequest` (`dict`) - filter to match changes in pull requests.  Specify only one of pullRequest or push.
        Structure is documented below.
        * `branch` (`str`) - Regex of branches to match.  Specify only one of branch or tag.
        * `commentControl` (`str`) - Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
          Possible values are `COMMENTS_DISABLED` and `COMMENTS_ENABLED`.
        * `invertRegex` (`bool`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.

      * `push` (`dict`) - filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.
        Structure is documented below.
        * `branch` (`str`) - Regex of branches to match.  Specify only one of branch or tag.
        * `invertRegex` (`bool`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        * `tag` (`str`) - Regex of tags to match.  Specify only one of branch or tag.
    """
    ignored_files: pulumi.Output[list]
    """
    ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
    extended with support for `**`.
    If ignoredFiles and changed files are both empty, then they are not
    used to determine whether or not to trigger a build.
    If ignoredFiles is not empty, then we ignore any files that match any
    of the ignored_file globs. If the change has no files that are outside
    of the ignoredFiles globs, then we do not trigger a build.
    """
    included_files: pulumi.Output[list]
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
    name: pulumi.Output[str]
    """
    Name of the volume to mount.
    Volume names must be unique per build step and must be valid names for
    Docker volumes. Each named volume must be used by at least two build steps.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    substitutions: pulumi.Output[dict]
    """
    Substitutions data for Build resource.
    """
    trigger_id: pulumi.Output[str]
    """
    The unique identifier for the trigger.
    """
    trigger_template: pulumi.Output[dict]
    """
    Template describing the types of source changes to trigger a build.
    Branch and tag names in trigger templates are interpreted as regular
    expressions. Any branch or tag change that matches that regular
    expression will trigger a build.
    One of `trigger_template` or `github` must be provided.
    Structure is documented below.

      * `branchName` (`str`) - Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
        This field is a regular expression.
      * `commitSha` (`str`) - Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
      * `dir` (`str`) - Working directory to use when running this step's container.
        If this value is a relative path, it is relative to the build's working
        directory. If this value is absolute, it may be outside the build's working
        directory, in which case the contents of the path may not be persisted
        across build step executions, unless a `volume` for that path is specified.
        If the build specifies a `RepoSource` with `dir` and a step with a
        `dir`,
        which specifies an absolute path, the `RepoSource` `dir` is ignored
        for the step's execution.
      * `invertRegex` (`bool`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
      * `project_id` (`str`) - ID of the project that owns the Cloud Source Repository. If
        omitted, the project ID requesting the build is assumed.
      * `repoName` (`str`) - Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
      * `tagName` (`str`) - Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
        This field is a regular expression.
    """
    def __init__(__self__, resource_name, opts=None, build=None, description=None, disabled=None, filename=None, github=None, ignored_files=None, included_files=None, name=None, project=None, substitutions=None, trigger_template=None, __props__=None, __name__=None, __opts__=None):
        """
        Configuration for an automated build in response to source repository changes.

        To get more information about Trigger, see:

        * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/)
        * How-to Guides
            * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] build: Contents of the build template. Either a filename or build template must be provided.
               Structure is documented below.
        :param pulumi.Input[str] description: Human-readable description of the trigger.
        :param pulumi.Input[bool] disabled: Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        :param pulumi.Input[str] filename: Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        :param pulumi.Input[dict] github: Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
        :param pulumi.Input[list] ignored_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If ignoredFiles and changed files are both empty, then they are not
               used to determine whether or not to trigger a build.
               If ignoredFiles is not empty, then we ignore any files that match any
               of the ignored_file globs. If the change has no files that are outside
               of the ignoredFiles globs, then we do not trigger a build.
        :param pulumi.Input[list] included_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
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
        :param pulumi.Input[dict] substitutions: Substitutions data for Build resource.
        :param pulumi.Input[dict] trigger_template: Template describing the types of source changes to trigger a build.
               Branch and tag names in trigger templates are interpreted as regular
               expressions. Any branch or tag change that matches that regular
               expression will trigger a build.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.

        The **build** object supports the following:

          * `images` (`pulumi.Input[list]`) - A list of images to be pushed upon the successful completion of all build steps.
            The images are pushed using the builder service account's credentials.
            The digests of the pushed images will be stored in the Build resource's results field.
            If any of the images fail to be pushed, the build status is marked FAILURE.
          * `steps` (`pulumi.Input[list]`) - The operations to be performed on the workspace.
            Structure is documented below.
            * `args` (`pulumi.Input[list]`) - A list of arguments that will be presented to the step when it is started.
              If the image used to run the step's container has an entrypoint, the args
              are used as arguments to that entrypoint. If the image does not define an
              entrypoint, the first element in args is used as the entrypoint, and the
              remainder will be used as arguments.
            * `dir` (`pulumi.Input[str]`) - Working directory to use when running this step's container.
              If this value is a relative path, it is relative to the build's working
              directory. If this value is absolute, it may be outside the build's working
              directory, in which case the contents of the path may not be persisted
              across build step executions, unless a `volume` for that path is specified.
              If the build specifies a `RepoSource` with `dir` and a step with a
              `dir`,
              which specifies an absolute path, the `RepoSource` `dir` is ignored
              for the step's execution.
            * `entrypoint` (`pulumi.Input[str]`) - Entrypoint to be used instead of the build step image's
              default entrypoint.
              If unset, the image's default entrypoint is used
            * `envs` (`pulumi.Input[list]`) - A list of environment variable definitions to be used when
              running a step.
              The elements are of the form "KEY=VALUE" for the environment variable
              "KEY" being given the value "VALUE".
            * `id` (`pulumi.Input[str]`) - Unique identifier for this build step, used in `wait_for` to
              reference this build step as a dependency.
            * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
              Volume names must be unique per build step and must be valid names for
              Docker volumes. Each named volume must be used by at least two build steps.
            * `secretEnvs` (`pulumi.Input[list]`) - A list of environment variables which are encrypted using
              a Cloud Key
              Management Service crypto key. These values must be specified in
              the build's `Secret`.
            * `timeout` (`pulumi.Input[str]`) - Time limit for executing this build step. If not defined,
              the step has no
              time limit and will be allowed to continue to run until either it
              completes or the build itself times out.
            * `timing` (`pulumi.Input[str]`) - Output only. Stores timing information for executing this
              build step.
            * `volumes` (`pulumi.Input[list]`) - List of volumes to mount into the build step.
              Each volume is created as an empty volume prior to execution of the
              build step. Upon completion of the build, volumes and their contents
              are discarded.
              Using a named volume in only one step is not valid as it is
              indicative of a build request with an incorrect configuration.
              Structure is documented below.
              * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
                Volume names must be unique per build step and must be valid names for
                Docker volumes. Each named volume must be used by at least two build steps.
              * `path` (`pulumi.Input[str]`) - Path at which to mount the volume.
                Paths must be absolute and cannot conflict with other volume paths on
                the same build step or with certain reserved volume paths.

            * `waitFors` (`pulumi.Input[list]`) - The ID(s) of the step(s) that this build step depends on.
              This build step will not start until all the build steps in `wait_for`
              have completed successfully. If `wait_for` is empty, this build step
              will start when all previous build steps in the `Build.Steps` list
              have completed successfully.

          * `tags` (`pulumi.Input[list]`) - Tags for annotation of a Build. These are not docker tags.
          * `timeout` (`pulumi.Input[str]`) - Time limit for executing this build step. If not defined,
            the step has no
            time limit and will be allowed to continue to run until either it
            completes or the build itself times out.

        The **github** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
            Volume names must be unique per build step and must be valid names for
            Docker volumes. Each named volume must be used by at least two build steps.
          * `owner` (`pulumi.Input[str]`) - Owner of the repository. For example: The owner for
            https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
          * `pullRequest` (`pulumi.Input[dict]`) - filter to match changes in pull requests.  Specify only one of pullRequest or push.
            Structure is documented below.
            * `branch` (`pulumi.Input[str]`) - Regex of branches to match.  Specify only one of branch or tag.
            * `commentControl` (`pulumi.Input[str]`) - Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
              Possible values are `COMMENTS_DISABLED` and `COMMENTS_ENABLED`.
            * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.

          * `push` (`pulumi.Input[dict]`) - filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.
            Structure is documented below.
            * `branch` (`pulumi.Input[str]`) - Regex of branches to match.  Specify only one of branch or tag.
            * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
            * `tag` (`pulumi.Input[str]`) - Regex of tags to match.  Specify only one of branch or tag.

        The **trigger_template** object supports the following:

          * `branchName` (`pulumi.Input[str]`) - Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
            This field is a regular expression.
          * `commitSha` (`pulumi.Input[str]`) - Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
          * `dir` (`pulumi.Input[str]`) - Working directory to use when running this step's container.
            If this value is a relative path, it is relative to the build's working
            directory. If this value is absolute, it may be outside the build's working
            directory, in which case the contents of the path may not be persisted
            across build step executions, unless a `volume` for that path is specified.
            If the build specifies a `RepoSource` with `dir` and a step with a
            `dir`,
            which specifies an absolute path, the `RepoSource` `dir` is ignored
            for the step's execution.
          * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
          * `project_id` (`pulumi.Input[str]`) - ID of the project that owns the Cloud Source Repository. If
            omitted, the project ID requesting the build is assumed.
          * `repoName` (`pulumi.Input[str]`) - Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
          * `tagName` (`pulumi.Input[str]`) - Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
            This field is a regular expression.
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
            __props__['trigger_template'] = trigger_template
            __props__['create_time'] = None
            __props__['trigger_id'] = None
        super(Trigger, __self__).__init__(
            'gcp:cloudbuild/trigger:Trigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, build=None, create_time=None, description=None, disabled=None, filename=None, github=None, ignored_files=None, included_files=None, name=None, project=None, substitutions=None, trigger_id=None, trigger_template=None):
        """
        Get an existing Trigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] build: Contents of the build template. Either a filename or build template must be provided.
               Structure is documented below.
        :param pulumi.Input[str] create_time: Time when the trigger was created.
        :param pulumi.Input[str] description: Human-readable description of the trigger.
        :param pulumi.Input[bool] disabled: Whether the trigger is disabled or not. If true, the trigger will never result in a build.
        :param pulumi.Input[str] filename: Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
        :param pulumi.Input[dict] github: Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.
        :param pulumi.Input[list] ignored_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
               extended with support for `**`.
               If ignoredFiles and changed files are both empty, then they are not
               used to determine whether or not to trigger a build.
               If ignoredFiles is not empty, then we ignore any files that match any
               of the ignored_file globs. If the change has no files that are outside
               of the ignoredFiles globs, then we do not trigger a build.
        :param pulumi.Input[list] included_files: ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
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
        :param pulumi.Input[dict] substitutions: Substitutions data for Build resource.
        :param pulumi.Input[str] trigger_id: The unique identifier for the trigger.
        :param pulumi.Input[dict] trigger_template: Template describing the types of source changes to trigger a build.
               Branch and tag names in trigger templates are interpreted as regular
               expressions. Any branch or tag change that matches that regular
               expression will trigger a build.
               One of `trigger_template` or `github` must be provided.
               Structure is documented below.

        The **build** object supports the following:

          * `images` (`pulumi.Input[list]`) - A list of images to be pushed upon the successful completion of all build steps.
            The images are pushed using the builder service account's credentials.
            The digests of the pushed images will be stored in the Build resource's results field.
            If any of the images fail to be pushed, the build status is marked FAILURE.
          * `steps` (`pulumi.Input[list]`) - The operations to be performed on the workspace.
            Structure is documented below.
            * `args` (`pulumi.Input[list]`) - A list of arguments that will be presented to the step when it is started.
              If the image used to run the step's container has an entrypoint, the args
              are used as arguments to that entrypoint. If the image does not define an
              entrypoint, the first element in args is used as the entrypoint, and the
              remainder will be used as arguments.
            * `dir` (`pulumi.Input[str]`) - Working directory to use when running this step's container.
              If this value is a relative path, it is relative to the build's working
              directory. If this value is absolute, it may be outside the build's working
              directory, in which case the contents of the path may not be persisted
              across build step executions, unless a `volume` for that path is specified.
              If the build specifies a `RepoSource` with `dir` and a step with a
              `dir`,
              which specifies an absolute path, the `RepoSource` `dir` is ignored
              for the step's execution.
            * `entrypoint` (`pulumi.Input[str]`) - Entrypoint to be used instead of the build step image's
              default entrypoint.
              If unset, the image's default entrypoint is used
            * `envs` (`pulumi.Input[list]`) - A list of environment variable definitions to be used when
              running a step.
              The elements are of the form "KEY=VALUE" for the environment variable
              "KEY" being given the value "VALUE".
            * `id` (`pulumi.Input[str]`) - Unique identifier for this build step, used in `wait_for` to
              reference this build step as a dependency.
            * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
              Volume names must be unique per build step and must be valid names for
              Docker volumes. Each named volume must be used by at least two build steps.
            * `secretEnvs` (`pulumi.Input[list]`) - A list of environment variables which are encrypted using
              a Cloud Key
              Management Service crypto key. These values must be specified in
              the build's `Secret`.
            * `timeout` (`pulumi.Input[str]`) - Time limit for executing this build step. If not defined,
              the step has no
              time limit and will be allowed to continue to run until either it
              completes or the build itself times out.
            * `timing` (`pulumi.Input[str]`) - Output only. Stores timing information for executing this
              build step.
            * `volumes` (`pulumi.Input[list]`) - List of volumes to mount into the build step.
              Each volume is created as an empty volume prior to execution of the
              build step. Upon completion of the build, volumes and their contents
              are discarded.
              Using a named volume in only one step is not valid as it is
              indicative of a build request with an incorrect configuration.
              Structure is documented below.
              * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
                Volume names must be unique per build step and must be valid names for
                Docker volumes. Each named volume must be used by at least two build steps.
              * `path` (`pulumi.Input[str]`) - Path at which to mount the volume.
                Paths must be absolute and cannot conflict with other volume paths on
                the same build step or with certain reserved volume paths.

            * `waitFors` (`pulumi.Input[list]`) - The ID(s) of the step(s) that this build step depends on.
              This build step will not start until all the build steps in `wait_for`
              have completed successfully. If `wait_for` is empty, this build step
              will start when all previous build steps in the `Build.Steps` list
              have completed successfully.

          * `tags` (`pulumi.Input[list]`) - Tags for annotation of a Build. These are not docker tags.
          * `timeout` (`pulumi.Input[str]`) - Time limit for executing this build step. If not defined,
            the step has no
            time limit and will be allowed to continue to run until either it
            completes or the build itself times out.

        The **github** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the volume to mount.
            Volume names must be unique per build step and must be valid names for
            Docker volumes. Each named volume must be used by at least two build steps.
          * `owner` (`pulumi.Input[str]`) - Owner of the repository. For example: The owner for
            https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
          * `pullRequest` (`pulumi.Input[dict]`) - filter to match changes in pull requests.  Specify only one of pullRequest or push.
            Structure is documented below.
            * `branch` (`pulumi.Input[str]`) - Regex of branches to match.  Specify only one of branch or tag.
            * `commentControl` (`pulumi.Input[str]`) - Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
              Possible values are `COMMENTS_DISABLED` and `COMMENTS_ENABLED`.
            * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.

          * `push` (`pulumi.Input[dict]`) - filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.
            Structure is documented below.
            * `branch` (`pulumi.Input[str]`) - Regex of branches to match.  Specify only one of branch or tag.
            * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
            * `tag` (`pulumi.Input[str]`) - Regex of tags to match.  Specify only one of branch or tag.

        The **trigger_template** object supports the following:

          * `branchName` (`pulumi.Input[str]`) - Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
            This field is a regular expression.
          * `commitSha` (`pulumi.Input[str]`) - Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
          * `dir` (`pulumi.Input[str]`) - Working directory to use when running this step's container.
            If this value is a relative path, it is relative to the build's working
            directory. If this value is absolute, it may be outside the build's working
            directory, in which case the contents of the path may not be persisted
            across build step executions, unless a `volume` for that path is specified.
            If the build specifies a `RepoSource` with `dir` and a step with a
            `dir`,
            which specifies an absolute path, the `RepoSource` `dir` is ignored
            for the step's execution.
          * `invertRegex` (`pulumi.Input[bool]`) - When true, only trigger a build if the revision regex does NOT match the git_ref regex.
          * `project_id` (`pulumi.Input[str]`) - ID of the project that owns the Cloud Source Repository. If
            omitted, the project ID requesting the build is assumed.
          * `repoName` (`pulumi.Input[str]`) - Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
          * `tagName` (`pulumi.Input[str]`) - Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
            This field is a regular expression.
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
        __props__["trigger_id"] = trigger_id
        __props__["trigger_template"] = trigger_template
        return Trigger(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
