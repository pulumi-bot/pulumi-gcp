# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Job']


class Job(pulumi.CustomResource):
    app_engine_http_target: pulumi.Output[Optional['outputs.JobAppEngineHttpTarget']] = pulumi.output_property("appEngineHttpTarget")
    """
    App Engine HTTP target.
    If the job providers a App Engine HTTP target the cron will
    send a request to the service instance  Structure is documented below.
    """
    attempt_deadline: pulumi.Output[Optional[str]] = pulumi.output_property("attemptDeadline")
    """
    The deadline for job attempts. If the request handler does not respond by this deadline then the request is
    cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
    execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
    The allowed duration for this deadline is:
    * For HTTP targets, between 15 seconds and 30 minutes.
    * For App Engine HTTP targets, between 15 seconds and 24 hours.
    A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
    """
    description: pulumi.Output[Optional[str]] = pulumi.output_property("description")
    """
    A human-readable description for the job.
    This string must not contain more than 500 characters.
    """
    http_target: pulumi.Output[Optional['outputs.JobHttpTarget']] = pulumi.output_property("httpTarget")
    """
    HTTP target.
    If the job providers a http_target the cron will
    send a request to the targeted url  Structure is documented below.
    """
    name: pulumi.Output[str] = pulumi.output_property("name")
    """
    The name of the job.
    """
    project: pulumi.Output[str] = pulumi.output_property("project")
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    pubsub_target: pulumi.Output[Optional['outputs.JobPubsubTarget']] = pulumi.output_property("pubsubTarget")
    """
    Pub/Sub target
    If the job providers a Pub/Sub target the cron will publish
    a message to the provided topic  Structure is documented below.
    """
    region: pulumi.Output[str] = pulumi.output_property("region")
    """
    Region where the scheduler job resides
    """
    retry_config: pulumi.Output[Optional['outputs.JobRetryConfig']] = pulumi.output_property("retryConfig")
    """
    By default, if a job does not complete successfully,
    meaning that an acknowledgement is not received from the handler,
    then it will be retried with exponential backoff according to the settings  Structure is documented below.
    """
    schedule: pulumi.Output[Optional[str]] = pulumi.output_property("schedule")
    """
    Describes the schedule on which the job will be executed.
    """
    time_zone: pulumi.Output[Optional[str]] = pulumi.output_property("timeZone")
    """
    Specifies the time zone to be used in interpreting schedule.
    The value of this field must be a time zone name from the tz database.
    """
    # pylint: disable=no-self-argument
    def __init__(__self__, resource_name, opts: Optional[pulumi.ResourceOptions] = None, app_engine_http_target: Optional[pulumi.Input[pulumi.InputType['JobAppEngineHttpTargetArgs']]] = None, attempt_deadline: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, http_target: Optional[pulumi.Input[pulumi.InputType['JobHttpTargetArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, pubsub_target: Optional[pulumi.Input[pulumi.InputType['JobPubsubTargetArgs']]] = None, region: Optional[pulumi.Input[str]] = None, retry_config: Optional[pulumi.Input[pulumi.InputType['JobRetryConfigArgs']]] = None, schedule: Optional[pulumi.Input[str]] = None, time_zone: Optional[pulumi.Input[str]] = None, __props__=None, __name__=None, __opts__=None) -> None:
        """
        A scheduled job that can publish a pubsub message or a http request
        every X interval of time, using crontab format string.

        To use Cloud Scheduler your project must contain an App Engine app
        that is located in one of the supported regions. If your project
        does not have an App Engine app, you must create one.

        To get more information about Job, see:

        * [API documentation](https://cloud.google.com/scheduler/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/scheduler/)

        ## Example Usage
        ### Scheduler Job Http

        ```python
        import pulumi
        import pulumi_gcp as gcp

        job = gcp.cloudscheduler.Job("job",
            attempt_deadline="320s",
            description="test http job",
            http_target={
                "httpMethod": "POST",
                "uri": "https://example.com/ping",
            },
            retry_config={
                "retryCount": 1,
            },
            schedule="*/8 * * * *",
            time_zone="America/New_York")
        ```
        ### Scheduler Job App Engine

        ```python
        import pulumi
        import pulumi_gcp as gcp

        job = gcp.cloudscheduler.Job("job",
            app_engine_http_target={
                "appEngineRouting": {
                    "instance": "my-instance-001",
                    "service": "web",
                    "version": "prod",
                },
                "httpMethod": "POST",
                "relativeUri": "/ping",
            },
            attempt_deadline="320s",
            description="test app engine job",
            retry_config={
                "maxDoublings": 2,
                "maxRetryDuration": "10s",
                "minBackoffDuration": "1s",
                "retryCount": 3,
            },
            schedule="*/4 * * * *",
            time_zone="Europe/London")
        ```
        ### Scheduler Job Oauth

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.get_default_service_account()
        job = gcp.cloudscheduler.Job("job",
            description="test http job",
            schedule="*/8 * * * *",
            time_zone="America/New_York",
            attempt_deadline="320s",
            http_target={
                "httpMethod": "GET",
                "uri": "https://cloudscheduler.googleapis.com/v1/projects/my-project-name/locations/us-west1/jobs",
                "oauthToken": {
                    "service_account_email": default.email,
                },
            })
        ```
        ### Scheduler Job Oidc

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.get_default_service_account()
        job = gcp.cloudscheduler.Job("job",
            description="test http job",
            schedule="*/8 * * * *",
            time_zone="America/New_York",
            attempt_deadline="320s",
            http_target={
                "httpMethod": "GET",
                "uri": "https://example.com/ping",
                "oidcToken": {
                    "service_account_email": default.email,
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['JobAppEngineHttpTargetArgs']] app_engine_http_target: App Engine HTTP target.
               If the job providers a App Engine HTTP target the cron will
               send a request to the service instance  Structure is documented below.
        :param pulumi.Input[str] attempt_deadline: The deadline for job attempts. If the request handler does not respond by this deadline then the request is
               cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
               execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
               The allowed duration for this deadline is:
               * For HTTP targets, between 15 seconds and 30 minutes.
               * For App Engine HTTP targets, between 15 seconds and 24 hours.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
        :param pulumi.Input[str] description: A human-readable description for the job.
               This string must not contain more than 500 characters.
        :param pulumi.Input[pulumi.InputType['JobHttpTargetArgs']] http_target: HTTP target.
               If the job providers a http_target the cron will
               send a request to the targeted url  Structure is documented below.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['JobPubsubTargetArgs']] pubsub_target: Pub/Sub target
               If the job providers a Pub/Sub target the cron will publish
               a message to the provided topic  Structure is documented below.
        :param pulumi.Input[str] region: Region where the scheduler job resides
        :param pulumi.Input[pulumi.InputType['JobRetryConfigArgs']] retry_config: By default, if a job does not complete successfully,
               meaning that an acknowledgement is not received from the handler,
               then it will be retried with exponential backoff according to the settings  Structure is documented below.
        :param pulumi.Input[str] schedule: Describes the schedule on which the job will be executed.
        :param pulumi.Input[str] time_zone: Specifies the time zone to be used in interpreting schedule.
               The value of this field must be a time zone name from the tz database.
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

            __props__['app_engine_http_target'] = app_engine_http_target
            __props__['attempt_deadline'] = attempt_deadline
            __props__['description'] = description
            __props__['http_target'] = http_target
            __props__['name'] = name
            __props__['project'] = project
            __props__['pubsub_target'] = pubsub_target
            __props__['region'] = region
            __props__['retry_config'] = retry_config
            __props__['schedule'] = schedule
            __props__['time_zone'] = time_zone
        super(Job, __self__).__init__(
            'gcp:cloudscheduler/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str, id: str, opts: Optional[pulumi.ResourceOptions] = None, app_engine_http_target: Optional[pulumi.Input[pulumi.InputType['JobAppEngineHttpTargetArgs']]] = None, attempt_deadline: Optional[pulumi.Input[str]] = None, description: Optional[pulumi.Input[str]] = None, http_target: Optional[pulumi.Input[pulumi.InputType['JobHttpTargetArgs']]] = None, name: Optional[pulumi.Input[str]] = None, project: Optional[pulumi.Input[str]] = None, pubsub_target: Optional[pulumi.Input[pulumi.InputType['JobPubsubTargetArgs']]] = None, region: Optional[pulumi.Input[str]] = None, retry_config: Optional[pulumi.Input[pulumi.InputType['JobRetryConfigArgs']]] = None, schedule: Optional[pulumi.Input[str]] = None, time_zone: Optional[pulumi.Input[str]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['JobAppEngineHttpTargetArgs']] app_engine_http_target: App Engine HTTP target.
               If the job providers a App Engine HTTP target the cron will
               send a request to the service instance  Structure is documented below.
        :param pulumi.Input[str] attempt_deadline: The deadline for job attempts. If the request handler does not respond by this deadline then the request is
               cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
               execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
               The allowed duration for this deadline is:
               * For HTTP targets, between 15 seconds and 30 minutes.
               * For App Engine HTTP targets, between 15 seconds and 24 hours.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
        :param pulumi.Input[str] description: A human-readable description for the job.
               This string must not contain more than 500 characters.
        :param pulumi.Input[pulumi.InputType['JobHttpTargetArgs']] http_target: HTTP target.
               If the job providers a http_target the cron will
               send a request to the targeted url  Structure is documented below.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['JobPubsubTargetArgs']] pubsub_target: Pub/Sub target
               If the job providers a Pub/Sub target the cron will publish
               a message to the provided topic  Structure is documented below.
        :param pulumi.Input[str] region: Region where the scheduler job resides
        :param pulumi.Input[pulumi.InputType['JobRetryConfigArgs']] retry_config: By default, if a job does not complete successfully,
               meaning that an acknowledgement is not received from the handler,
               then it will be retried with exponential backoff according to the settings  Structure is documented below.
        :param pulumi.Input[str] schedule: Describes the schedule on which the job will be executed.
        :param pulumi.Input[str] time_zone: Specifies the time zone to be used in interpreting schedule.
               The value of this field must be a time zone name from the tz database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_engine_http_target"] = app_engine_http_target
        __props__["attempt_deadline"] = attempt_deadline
        __props__["description"] = description
        __props__["http_target"] = http_target
        __props__["name"] = name
        __props__["project"] = project
        __props__["pubsub_target"] = pubsub_target
        __props__["region"] = region
        __props__["retry_config"] = retry_config
        __props__["schedule"] = schedule
        __props__["time_zone"] = time_zone
        return Job(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

