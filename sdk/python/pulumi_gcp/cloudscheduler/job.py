# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Job(pulumi.CustomResource):
    app_engine_http_target: pulumi.Output[dict]
    """
    App Engine HTTP target.
    If the job providers a App Engine HTTP target the cron will
    send a request to the service instance  Structure is documented below.

      * `appEngineRouting` (`dict`) - App Engine Routing setting for the job.  Structure is documented below.
        * `instance` (`str`) - App instance.
          By default, the job is sent to an instance which is available when the job is attempted.
        * `service` (`str`) - App service.
          By default, the job is sent to the service which is the default service when the job is attempted.
        * `version` (`str`) - App version.
          By default, the job is sent to the version which is the default version when the job is attempted.

      * `body` (`str`) - HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
      * `headers` (`dict`) - This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
      * `httpMethod` (`str`) - Which HTTP method to use for the request.
      * `relativeUri` (`str`) - The relative URI.
        The relative URL must begin with "/" and must be a valid HTTP relative URL.
        It can contain a path, query string arguments, and \# fragments.
        If the relative URL is empty, then the root path "/" will be used.
        No spaces are allowed, and the maximum length allowed is 2083 characters
    """
    attempt_deadline: pulumi.Output[str]
    """
    The deadline for job attempts. If the request handler does not respond by this deadline then the request is
    cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
    execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
    The allowed duration for this deadline is:
    * For HTTP targets, between 15 seconds and 30 minutes.
    * For App Engine HTTP targets, between 15 seconds and 24 hours.
    A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
    """
    description: pulumi.Output[str]
    """
    A human-readable description for the job.
    This string must not contain more than 500 characters.
    """
    http_target: pulumi.Output[dict]
    """
    HTTP target.
    If the job providers a http_target the cron will
    send a request to the targeted url  Structure is documented below.

      * `body` (`str`) - HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
      * `headers` (`dict`) - This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
      * `httpMethod` (`str`) - Which HTTP method to use for the request.
      * `oauthToken` (`dict`) - Contains information needed for generating an OAuth token.
        This type of authorization should be used when sending requests to a GCP endpoint.  Structure is documented below.
        * `scope` (`str`) - OAuth scope to be used for generating OAuth access token. If not specified,
          "https://www.googleapis.com/auth/cloud-platform" will be used.
        * `service_account_email` (`str`) - Service account email to be used for generating OAuth token.
          The service account must be within the same project as the job.

      * `oidcToken` (`dict`) - Contains information needed for generating an OpenID Connect token.
        This type of authorization should be used when sending requests to third party endpoints or Cloud Run.  Structure is documented below.
        * `audience` (`str`) - Audience to be used when generating OIDC token. If not specified,
          the URI specified in target will be used.
        * `service_account_email` (`str`) - Service account email to be used for generating OAuth token.
          The service account must be within the same project as the job.

      * `uri` (`str`) - The full URI path that the request will be sent to.
    """
    name: pulumi.Output[str]
    """
    The name of the job.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    pubsub_target: pulumi.Output[dict]
    """
    Pub/Sub target
    If the job providers a Pub/Sub target the cron will publish
    a message to the provided topic  Structure is documented below.

      * `attributes` (`dict`) - Attributes for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
      * `data` (`str`) - The message payload for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
      * `topicName` (`str`) - The full resource name for the Cloud Pub/Sub topic to which
        messages will be published when a job is delivered. ~>**NOTE**:
        The topic name must be in the same format as required by PubSub's
        PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
    """
    region: pulumi.Output[str]
    """
    Region where the scheduler job resides
    """
    retry_config: pulumi.Output[dict]
    """
    By default, if a job does not complete successfully,
    meaning that an acknowledgement is not received from the handler,
    then it will be retried with exponential backoff according to the settings  Structure is documented below.

      * `maxBackoffDuration` (`str`) - The maximum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
      * `maxDoublings` (`float`) - The time between retries will double maxDoublings times.
        A job's retry interval starts at minBackoffDuration,
        then doubles maxDoublings times, then increases linearly,
        and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
      * `maxRetryDuration` (`str`) - The time limit for retrying a failed job, measured from time when an execution was first attempted.
        If specified with retryCount, the job will be retried until both limits are reached.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
      * `minBackoffDuration` (`str`) - The minimum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
      * `retryCount` (`float`) - The number of attempts that the system will make to run a
        job using the exponential backoff procedure described by maxDoublings.
        Values greater than 5 and negative values are not allowed.
    """
    schedule: pulumi.Output[str]
    """
    Describes the schedule on which the job will be executed.
    """
    time_zone: pulumi.Output[str]
    """
    Specifies the time zone to be used in interpreting schedule.
    The value of this field must be a time zone name from the tz database.
    """
    def __init__(__self__, resource_name, opts=None, app_engine_http_target=None, attempt_deadline=None, description=None, http_target=None, name=None, project=None, pubsub_target=None, region=None, retry_config=None, schedule=None, time_zone=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[dict] app_engine_http_target: App Engine HTTP target.
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
        :param pulumi.Input[dict] http_target: HTTP target.
               If the job providers a http_target the cron will
               send a request to the targeted url  Structure is documented below.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] pubsub_target: Pub/Sub target
               If the job providers a Pub/Sub target the cron will publish
               a message to the provided topic  Structure is documented below.
        :param pulumi.Input[str] region: Region where the scheduler job resides
        :param pulumi.Input[dict] retry_config: By default, if a job does not complete successfully,
               meaning that an acknowledgement is not received from the handler,
               then it will be retried with exponential backoff according to the settings  Structure is documented below.
        :param pulumi.Input[str] schedule: Describes the schedule on which the job will be executed.
        :param pulumi.Input[str] time_zone: Specifies the time zone to be used in interpreting schedule.
               The value of this field must be a time zone name from the tz database.

        The **app_engine_http_target** object supports the following:

          * `appEngineRouting` (`pulumi.Input[dict]`) - App Engine Routing setting for the job.  Structure is documented below.
            * `instance` (`pulumi.Input[str]`) - App instance.
              By default, the job is sent to an instance which is available when the job is attempted.
            * `service` (`pulumi.Input[str]`) - App service.
              By default, the job is sent to the service which is the default service when the job is attempted.
            * `version` (`pulumi.Input[str]`) - App version.
              By default, the job is sent to the version which is the default version when the job is attempted.

          * `body` (`pulumi.Input[str]`) - HTTP request body.
            A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
            It is an error to set body on a job with an incompatible HttpMethod.
          * `headers` (`pulumi.Input[dict]`) - This map contains the header field names and values.
            Repeated headers are not supported, but a header value can contain commas.
          * `httpMethod` (`pulumi.Input[str]`) - Which HTTP method to use for the request.
          * `relativeUri` (`pulumi.Input[str]`) - The relative URI.
            The relative URL must begin with "/" and must be a valid HTTP relative URL.
            It can contain a path, query string arguments, and \# fragments.
            If the relative URL is empty, then the root path "/" will be used.
            No spaces are allowed, and the maximum length allowed is 2083 characters

        The **http_target** object supports the following:

          * `body` (`pulumi.Input[str]`) - HTTP request body.
            A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
            It is an error to set body on a job with an incompatible HttpMethod.
          * `headers` (`pulumi.Input[dict]`) - This map contains the header field names and values.
            Repeated headers are not supported, but a header value can contain commas.
          * `httpMethod` (`pulumi.Input[str]`) - Which HTTP method to use for the request.
          * `oauthToken` (`pulumi.Input[dict]`) - Contains information needed for generating an OAuth token.
            This type of authorization should be used when sending requests to a GCP endpoint.  Structure is documented below.
            * `scope` (`pulumi.Input[str]`) - OAuth scope to be used for generating OAuth access token. If not specified,
              "https://www.googleapis.com/auth/cloud-platform" will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating OAuth token.
              The service account must be within the same project as the job.

          * `oidcToken` (`pulumi.Input[dict]`) - Contains information needed for generating an OpenID Connect token.
            This type of authorization should be used when sending requests to third party endpoints or Cloud Run.  Structure is documented below.
            * `audience` (`pulumi.Input[str]`) - Audience to be used when generating OIDC token. If not specified,
              the URI specified in target will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating OAuth token.
              The service account must be within the same project as the job.

          * `uri` (`pulumi.Input[str]`) - The full URI path that the request will be sent to.

        The **pubsub_target** object supports the following:

          * `attributes` (`pulumi.Input[dict]`) - Attributes for PubsubMessage.
            Pubsub message must contain either non-empty data, or at least one attribute.
          * `data` (`pulumi.Input[str]`) - The message payload for PubsubMessage.
            Pubsub message must contain either non-empty data, or at least one attribute.
          * `topicName` (`pulumi.Input[str]`) - The full resource name for the Cloud Pub/Sub topic to which
            messages will be published when a job is delivered. ~>**NOTE**:
            The topic name must be in the same format as required by PubSub's
            PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.

        The **retry_config** object supports the following:

          * `maxBackoffDuration` (`pulumi.Input[str]`) - The maximum amount of time to wait before retrying a job after it fails.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `maxDoublings` (`pulumi.Input[float]`) - The time between retries will double maxDoublings times.
            A job's retry interval starts at minBackoffDuration,
            then doubles maxDoublings times, then increases linearly,
            and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
          * `maxRetryDuration` (`pulumi.Input[str]`) - The time limit for retrying a failed job, measured from time when an execution was first attempted.
            If specified with retryCount, the job will be retried until both limits are reached.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `minBackoffDuration` (`pulumi.Input[str]`) - The minimum amount of time to wait before retrying a job after it fails.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `retryCount` (`pulumi.Input[float]`) - The number of attempts that the system will make to run a
            job using the exponential backoff procedure described by maxDoublings.
            Values greater than 5 and negative values are not allowed.
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
    def get(resource_name, id, opts=None, app_engine_http_target=None, attempt_deadline=None, description=None, http_target=None, name=None, project=None, pubsub_target=None, region=None, retry_config=None, schedule=None, time_zone=None):
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] app_engine_http_target: App Engine HTTP target.
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
        :param pulumi.Input[dict] http_target: HTTP target.
               If the job providers a http_target the cron will
               send a request to the targeted url  Structure is documented below.
        :param pulumi.Input[str] name: The name of the job.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] pubsub_target: Pub/Sub target
               If the job providers a Pub/Sub target the cron will publish
               a message to the provided topic  Structure is documented below.
        :param pulumi.Input[str] region: Region where the scheduler job resides
        :param pulumi.Input[dict] retry_config: By default, if a job does not complete successfully,
               meaning that an acknowledgement is not received from the handler,
               then it will be retried with exponential backoff according to the settings  Structure is documented below.
        :param pulumi.Input[str] schedule: Describes the schedule on which the job will be executed.
        :param pulumi.Input[str] time_zone: Specifies the time zone to be used in interpreting schedule.
               The value of this field must be a time zone name from the tz database.

        The **app_engine_http_target** object supports the following:

          * `appEngineRouting` (`pulumi.Input[dict]`) - App Engine Routing setting for the job.  Structure is documented below.
            * `instance` (`pulumi.Input[str]`) - App instance.
              By default, the job is sent to an instance which is available when the job is attempted.
            * `service` (`pulumi.Input[str]`) - App service.
              By default, the job is sent to the service which is the default service when the job is attempted.
            * `version` (`pulumi.Input[str]`) - App version.
              By default, the job is sent to the version which is the default version when the job is attempted.

          * `body` (`pulumi.Input[str]`) - HTTP request body.
            A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
            It is an error to set body on a job with an incompatible HttpMethod.
          * `headers` (`pulumi.Input[dict]`) - This map contains the header field names and values.
            Repeated headers are not supported, but a header value can contain commas.
          * `httpMethod` (`pulumi.Input[str]`) - Which HTTP method to use for the request.
          * `relativeUri` (`pulumi.Input[str]`) - The relative URI.
            The relative URL must begin with "/" and must be a valid HTTP relative URL.
            It can contain a path, query string arguments, and \# fragments.
            If the relative URL is empty, then the root path "/" will be used.
            No spaces are allowed, and the maximum length allowed is 2083 characters

        The **http_target** object supports the following:

          * `body` (`pulumi.Input[str]`) - HTTP request body.
            A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
            It is an error to set body on a job with an incompatible HttpMethod.
          * `headers` (`pulumi.Input[dict]`) - This map contains the header field names and values.
            Repeated headers are not supported, but a header value can contain commas.
          * `httpMethod` (`pulumi.Input[str]`) - Which HTTP method to use for the request.
          * `oauthToken` (`pulumi.Input[dict]`) - Contains information needed for generating an OAuth token.
            This type of authorization should be used when sending requests to a GCP endpoint.  Structure is documented below.
            * `scope` (`pulumi.Input[str]`) - OAuth scope to be used for generating OAuth access token. If not specified,
              "https://www.googleapis.com/auth/cloud-platform" will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating OAuth token.
              The service account must be within the same project as the job.

          * `oidcToken` (`pulumi.Input[dict]`) - Contains information needed for generating an OpenID Connect token.
            This type of authorization should be used when sending requests to third party endpoints or Cloud Run.  Structure is documented below.
            * `audience` (`pulumi.Input[str]`) - Audience to be used when generating OIDC token. If not specified,
              the URI specified in target will be used.
            * `service_account_email` (`pulumi.Input[str]`) - Service account email to be used for generating OAuth token.
              The service account must be within the same project as the job.

          * `uri` (`pulumi.Input[str]`) - The full URI path that the request will be sent to.

        The **pubsub_target** object supports the following:

          * `attributes` (`pulumi.Input[dict]`) - Attributes for PubsubMessage.
            Pubsub message must contain either non-empty data, or at least one attribute.
          * `data` (`pulumi.Input[str]`) - The message payload for PubsubMessage.
            Pubsub message must contain either non-empty data, or at least one attribute.
          * `topicName` (`pulumi.Input[str]`) - The full resource name for the Cloud Pub/Sub topic to which
            messages will be published when a job is delivered. ~>**NOTE**:
            The topic name must be in the same format as required by PubSub's
            PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.

        The **retry_config** object supports the following:

          * `maxBackoffDuration` (`pulumi.Input[str]`) - The maximum amount of time to wait before retrying a job after it fails.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `maxDoublings` (`pulumi.Input[float]`) - The time between retries will double maxDoublings times.
            A job's retry interval starts at minBackoffDuration,
            then doubles maxDoublings times, then increases linearly,
            and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
          * `maxRetryDuration` (`pulumi.Input[str]`) - The time limit for retrying a failed job, measured from time when an execution was first attempted.
            If specified with retryCount, the job will be retried until both limits are reached.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `minBackoffDuration` (`pulumi.Input[str]`) - The minimum amount of time to wait before retrying a job after it fails.
            A duration in seconds with up to nine fractional digits, terminated by 's'.
          * `retryCount` (`pulumi.Input[float]`) - The number of attempts that the system will make to run a
            job using the exponential backoff procedure described by maxDoublings.
            Values greater than 5 and negative values are not allowed.
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
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
