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

@pulumi.input_type
class TriggerBuildArgs:
    steps: pulumi.Input[List[pulumi.Input['TriggerBuildStepArgs']]] = pulumi.input_property("steps")
    """
    The operations to be performed on the workspace.  Structure is documented below.
    """
    images: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("images")
    """
    A list of images to be pushed upon the successful completion of all build steps.
    The images are pushed using the builder service account's credentials.
    The digests of the pushed images will be stored in the Build resource's results field.
    If any of the images fail to be pushed, the build status is marked FAILURE.
    """
    tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("tags")
    """
    Tags for annotation of a Build. These are not docker tags.
    """
    timeout: Optional[pulumi.Input[str]] = pulumi.input_property("timeout")
    """
    Time limit for executing this build step. If not defined,
    the step has no
    time limit and will be allowed to continue to run until either it
    completes or the build itself times out.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, steps: pulumi.Input[List[pulumi.Input['TriggerBuildStepArgs']]], images: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, timeout: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[List[pulumi.Input['TriggerBuildStepArgs']]] steps: The operations to be performed on the workspace.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] images: A list of images to be pushed upon the successful completion of all build steps.
               The images are pushed using the builder service account's credentials.
               The digests of the pushed images will be stored in the Build resource's results field.
               If any of the images fail to be pushed, the build status is marked FAILURE.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags for annotation of a Build. These are not docker tags.
        :param pulumi.Input[str] timeout: Time limit for executing this build step. If not defined,
               the step has no
               time limit and will be allowed to continue to run until either it
               completes or the build itself times out.
        """
        __self__.steps = steps
        __self__.images = images
        __self__.tags = tags
        __self__.timeout = timeout

@pulumi.input_type
class TriggerBuildStepArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Name of the volume to mount.
    Volume names must be unique per build step and must be valid names for
    Docker volumes. Each named volume must be used by at least two build steps.
    """
    args: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("args")
    """
    A list of arguments that will be presented to the step when it is started.
    If the image used to run the step's container has an entrypoint, the args
    are used as arguments to that entrypoint. If the image does not define an
    entrypoint, the first element in args is used as the entrypoint, and the
    remainder will be used as arguments.
    """
    dir: Optional[pulumi.Input[str]] = pulumi.input_property("dir")
    """
    Working directory to use when running this step's container.
    If this value is a relative path, it is relative to the build's working
    directory. If this value is absolute, it may be outside the build's working
    directory, in which case the contents of the path may not be persisted
    across build step executions, unless a `volume` for that path is specified.
    If the build specifies a `RepoSource` with `dir` and a step with a
    `dir`,
    which specifies an absolute path, the `RepoSource` `dir` is ignored
    for the step's execution.
    """
    entrypoint: Optional[pulumi.Input[str]] = pulumi.input_property("entrypoint")
    """
    Entrypoint to be used instead of the build step image's
    default entrypoint.
    If unset, the image's default entrypoint is used
    """
    envs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("envs")
    """
    A list of environment variable definitions to be used when
    running a step.
    The elements are of the form "KEY=VALUE" for the environment variable
    "KEY" being given the value "VALUE".
    """
    id: Optional[pulumi.Input[str]] = pulumi.input_property("id")
    """
    Unique identifier for this build step, used in `wait_for` to
    reference this build step as a dependency.
    """
    secret_envs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("secretEnvs")
    """
    A list of environment variables which are encrypted using
    a Cloud Key
    Management Service crypto key. These values must be specified in
    the build's `Secret`.
    """
    timeout: Optional[pulumi.Input[str]] = pulumi.input_property("timeout")
    """
    Time limit for executing this build step. If not defined,
    the step has no
    time limit and will be allowed to continue to run until either it
    completes or the build itself times out.
    """
    timing: Optional[pulumi.Input[str]] = pulumi.input_property("timing")
    """
    Output only. Stores timing information for executing this
    build step.
    """
    volumes: Optional[pulumi.Input[List[pulumi.Input['TriggerBuildStepVolumeArgs']]]] = pulumi.input_property("volumes")
    """
    List of volumes to mount into the build step.
    Each volume is created as an empty volume prior to execution of the
    build step. Upon completion of the build, volumes and their contents
    are discarded.
    Using a named volume in only one step is not valid as it is
    indicative of a build request with an incorrect configuration.  Structure is documented below.
    """
    wait_fors: Optional[pulumi.Input[List[pulumi.Input[str]]]] = pulumi.input_property("waitFors")
    """
    The ID(s) of the step(s) that this build step depends on.
    This build step will not start until all the build steps in `wait_for`
    have completed successfully. If `wait_for` is empty, this build step
    will start when all previous build steps in the `Build.Steps` list
    have completed successfully.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], args: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, dir: Optional[pulumi.Input[str]] = None, entrypoint: Optional[pulumi.Input[str]] = None, envs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, id: Optional[pulumi.Input[str]] = None, secret_envs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None, timeout: Optional[pulumi.Input[str]] = None, timing: Optional[pulumi.Input[str]] = None, volumes: Optional[pulumi.Input[List[pulumi.Input['TriggerBuildStepVolumeArgs']]]] = None, wait_fors: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> None:
        """
        :param pulumi.Input[str] name: Name of the volume to mount.
               Volume names must be unique per build step and must be valid names for
               Docker volumes. Each named volume must be used by at least two build steps.
        :param pulumi.Input[List[pulumi.Input[str]]] args: A list of arguments that will be presented to the step when it is started.
               If the image used to run the step's container has an entrypoint, the args
               are used as arguments to that entrypoint. If the image does not define an
               entrypoint, the first element in args is used as the entrypoint, and the
               remainder will be used as arguments.
        :param pulumi.Input[str] dir: Working directory to use when running this step's container.
               If this value is a relative path, it is relative to the build's working
               directory. If this value is absolute, it may be outside the build's working
               directory, in which case the contents of the path may not be persisted
               across build step executions, unless a `volume` for that path is specified.
               If the build specifies a `RepoSource` with `dir` and a step with a
               `dir`,
               which specifies an absolute path, the `RepoSource` `dir` is ignored
               for the step's execution.
        :param pulumi.Input[str] entrypoint: Entrypoint to be used instead of the build step image's
               default entrypoint.
               If unset, the image's default entrypoint is used
        :param pulumi.Input[List[pulumi.Input[str]]] envs: A list of environment variable definitions to be used when
               running a step.
               The elements are of the form "KEY=VALUE" for the environment variable
               "KEY" being given the value "VALUE".
        :param pulumi.Input[str] id: Unique identifier for this build step, used in `wait_for` to
               reference this build step as a dependency.
        :param pulumi.Input[List[pulumi.Input[str]]] secret_envs: A list of environment variables which are encrypted using
               a Cloud Key
               Management Service crypto key. These values must be specified in
               the build's `Secret`.
        :param pulumi.Input[str] timeout: Time limit for executing this build step. If not defined,
               the step has no
               time limit and will be allowed to continue to run until either it
               completes or the build itself times out.
        :param pulumi.Input[str] timing: Output only. Stores timing information for executing this
               build step.
        :param pulumi.Input[List[pulumi.Input['TriggerBuildStepVolumeArgs']]] volumes: List of volumes to mount into the build step.
               Each volume is created as an empty volume prior to execution of the
               build step. Upon completion of the build, volumes and their contents
               are discarded.
               Using a named volume in only one step is not valid as it is
               indicative of a build request with an incorrect configuration.  Structure is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] wait_fors: The ID(s) of the step(s) that this build step depends on.
               This build step will not start until all the build steps in `wait_for`
               have completed successfully. If `wait_for` is empty, this build step
               will start when all previous build steps in the `Build.Steps` list
               have completed successfully.
        """
        __self__.name = name
        __self__.args = args
        __self__.dir = dir
        __self__.entrypoint = entrypoint
        __self__.envs = envs
        __self__.id = id
        __self__.secret_envs = secret_envs
        __self__.timeout = timeout
        __self__.timing = timing
        __self__.volumes = volumes
        __self__.wait_fors = wait_fors

@pulumi.input_type
class TriggerBuildStepVolumeArgs:
    name: pulumi.Input[str] = pulumi.input_property("name")
    """
    Name of the volume to mount.
    Volume names must be unique per build step and must be valid names for
    Docker volumes. Each named volume must be used by at least two build steps.
    """
    path: pulumi.Input[str] = pulumi.input_property("path")
    """
    Path at which to mount the volume.
    Paths must be absolute and cannot conflict with other volume paths on
    the same build step or with certain reserved volume paths.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: pulumi.Input[str], path: pulumi.Input[str]) -> None:
        """
        :param pulumi.Input[str] name: Name of the volume to mount.
               Volume names must be unique per build step and must be valid names for
               Docker volumes. Each named volume must be used by at least two build steps.
        :param pulumi.Input[str] path: Path at which to mount the volume.
               Paths must be absolute and cannot conflict with other volume paths on
               the same build step or with certain reserved volume paths.
        """
        __self__.name = name
        __self__.path = path

@pulumi.input_type
class TriggerGithubArgs:
    name: Optional[pulumi.Input[str]] = pulumi.input_property("name")
    """
    Name of the volume to mount.
    Volume names must be unique per build step and must be valid names for
    Docker volumes. Each named volume must be used by at least two build steps.
    """
    owner: Optional[pulumi.Input[str]] = pulumi.input_property("owner")
    """
    Owner of the repository. For example: The owner for
    https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
    """
    pull_request: Optional[pulumi.Input['TriggerGithubPullRequestArgs']] = pulumi.input_property("pullRequest")
    """
    filter to match changes in pull requests.  Specify only one of pullRequest or push.  Structure is documented below.
    """
    push: Optional[pulumi.Input['TriggerGithubPushArgs']] = pulumi.input_property("push")
    """
    filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.  Structure is documented below.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, name: Optional[pulumi.Input[str]] = None, owner: Optional[pulumi.Input[str]] = None, pull_request: Optional[pulumi.Input['TriggerGithubPullRequestArgs']] = None, push: Optional[pulumi.Input['TriggerGithubPushArgs']] = None) -> None:
        """
        :param pulumi.Input[str] name: Name of the volume to mount.
               Volume names must be unique per build step and must be valid names for
               Docker volumes. Each named volume must be used by at least two build steps.
        :param pulumi.Input[str] owner: Owner of the repository. For example: The owner for
               https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
        :param pulumi.Input['TriggerGithubPullRequestArgs'] pull_request: filter to match changes in pull requests.  Specify only one of pullRequest or push.  Structure is documented below.
        :param pulumi.Input['TriggerGithubPushArgs'] push: filter to match changes in refs, like branches or tags.  Specify only one of pullRequest or push.  Structure is documented below.
        """
        __self__.name = name
        __self__.owner = owner
        __self__.pull_request = pull_request
        __self__.push = push

@pulumi.input_type
class TriggerGithubPullRequestArgs:
    branch: pulumi.Input[str] = pulumi.input_property("branch")
    """
    Regex of branches to match.  Specify only one of branch or tag.
    """
    comment_control: Optional[pulumi.Input[str]] = pulumi.input_property("commentControl")
    """
    Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
    """
    invert_regex: Optional[pulumi.Input[bool]] = pulumi.input_property("invertRegex")
    """
    When true, only trigger a build if the revision regex does NOT match the git_ref regex.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, branch: pulumi.Input[str], comment_control: Optional[pulumi.Input[str]] = None, invert_regex: Optional[pulumi.Input[bool]] = None) -> None:
        """
        :param pulumi.Input[str] branch: Regex of branches to match.  Specify only one of branch or tag.
        :param pulumi.Input[str] comment_control: Whether to block builds on a "/gcbrun" comment from a repository owner or collaborator.
        :param pulumi.Input[bool] invert_regex: When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        """
        __self__.branch = branch
        __self__.comment_control = comment_control
        __self__.invert_regex = invert_regex

@pulumi.input_type
class TriggerGithubPushArgs:
    branch: Optional[pulumi.Input[str]] = pulumi.input_property("branch")
    """
    Regex of branches to match.  Specify only one of branch or tag.
    """
    invert_regex: Optional[pulumi.Input[bool]] = pulumi.input_property("invertRegex")
    """
    When true, only trigger a build if the revision regex does NOT match the git_ref regex.
    """
    tag: Optional[pulumi.Input[str]] = pulumi.input_property("tag")
    """
    Regex of tags to match.  Specify only one of branch or tag.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, branch: Optional[pulumi.Input[str]] = None, invert_regex: Optional[pulumi.Input[bool]] = None, tag: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] branch: Regex of branches to match.  Specify only one of branch or tag.
        :param pulumi.Input[bool] invert_regex: When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        :param pulumi.Input[str] tag: Regex of tags to match.  Specify only one of branch or tag.
        """
        __self__.branch = branch
        __self__.invert_regex = invert_regex
        __self__.tag = tag

@pulumi.input_type
class TriggerTriggerTemplateArgs:
    branch_name: Optional[pulumi.Input[str]] = pulumi.input_property("branchName")
    """
    Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
    This field is a regular expression.
    """
    commit_sha: Optional[pulumi.Input[str]] = pulumi.input_property("commitSha")
    """
    Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
    """
    dir: Optional[pulumi.Input[str]] = pulumi.input_property("dir")
    """
    Working directory to use when running this step's container.
    If this value is a relative path, it is relative to the build's working
    directory. If this value is absolute, it may be outside the build's working
    directory, in which case the contents of the path may not be persisted
    across build step executions, unless a `volume` for that path is specified.
    If the build specifies a `RepoSource` with `dir` and a step with a
    `dir`,
    which specifies an absolute path, the `RepoSource` `dir` is ignored
    for the step's execution.
    """
    invert_regex: Optional[pulumi.Input[bool]] = pulumi.input_property("invertRegex")
    """
    When true, only trigger a build if the revision regex does NOT match the git_ref regex.
    """
    project_id: Optional[pulumi.Input[str]] = pulumi.input_property("projectId")
    """
    ID of the project that owns the Cloud Source Repository. If
    omitted, the project ID requesting the build is assumed.
    """
    repo_name: Optional[pulumi.Input[str]] = pulumi.input_property("repoName")
    """
    Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
    """
    tag_name: Optional[pulumi.Input[str]] = pulumi.input_property("tagName")
    """
    Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
    This field is a regular expression.
    """

    # pylint: disable=no-self-argument
    def __init__(__self__, *, branch_name: Optional[pulumi.Input[str]] = None, commit_sha: Optional[pulumi.Input[str]] = None, dir: Optional[pulumi.Input[str]] = None, invert_regex: Optional[pulumi.Input[bool]] = None, project_id: Optional[pulumi.Input[str]] = None, repo_name: Optional[pulumi.Input[str]] = None, tag_name: Optional[pulumi.Input[str]] = None) -> None:
        """
        :param pulumi.Input[str] branch_name: Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
               This field is a regular expression.
        :param pulumi.Input[str] commit_sha: Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided.
        :param pulumi.Input[str] dir: Working directory to use when running this step's container.
               If this value is a relative path, it is relative to the build's working
               directory. If this value is absolute, it may be outside the build's working
               directory, in which case the contents of the path may not be persisted
               across build step executions, unless a `volume` for that path is specified.
               If the build specifies a `RepoSource` with `dir` and a step with a
               `dir`,
               which specifies an absolute path, the `RepoSource` `dir` is ignored
               for the step's execution.
        :param pulumi.Input[bool] invert_regex: When true, only trigger a build if the revision regex does NOT match the git_ref regex.
        :param pulumi.Input[str] project_id: ID of the project that owns the Cloud Source Repository. If
               omitted, the project ID requesting the build is assumed.
        :param pulumi.Input[str] repo_name: Name of the Cloud Source Repository. If omitted, the name "default" is assumed.
        :param pulumi.Input[str] tag_name: Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
               This field is a regular expression.
        """
        __self__.branch_name = branch_name
        __self__.commit_sha = commit_sha
        __self__.dir = dir
        __self__.invert_regex = invert_regex
        __self__.project_id = project_id
        __self__.repo_name = repo_name
        __self__.tag_name = tag_name

