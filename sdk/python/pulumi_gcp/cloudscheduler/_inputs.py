# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'JobAppEngineHttpTargetArgs',
    'JobAppEngineHttpTargetAppEngineRoutingArgs',
    'JobHttpTargetArgs',
    'JobHttpTargetOauthTokenArgs',
    'JobHttpTargetOidcTokenArgs',
    'JobPubsubTargetArgs',
    'JobRetryConfigArgs',
]

@pulumi.input_type
class JobAppEngineHttpTargetArgs:
    def __init__(__self__, *,
                 relative_uri: pulumi.Input[str],
                 app_engine_routing: Optional[pulumi.Input['JobAppEngineHttpTargetAppEngineRoutingArgs']] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 http_method: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] relative_uri: The relative URI.
               The relative URL must begin with "/" and must be a valid HTTP relative URL.
               It can contain a path, query string arguments, and \# fragments.
               If the relative URL is empty, then the root path "/" will be used.
               No spaces are allowed, and the maximum length allowed is 2083 characters
        :param pulumi.Input['JobAppEngineHttpTargetAppEngineRoutingArgs'] app_engine_routing: App Engine Routing setting for the job.
               Structure is documented below.
        :param pulumi.Input[str] body: HTTP request body.
               A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
               It is an error to set body on a job with an incompatible HttpMethod.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: This map contains the header field names and values.
               Repeated headers are not supported, but a header value can contain commas.
        :param pulumi.Input[str] http_method: Which HTTP method to use for the request.
        """
        pulumi.set(__self__, "relative_uri", relative_uri)
        if app_engine_routing is not None:
            pulumi.set(__self__, "app_engine_routing", app_engine_routing)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)

    @property
    @pulumi.getter(name="relativeUri")
    def relative_uri(self) -> pulumi.Input[str]:
        """
        The relative URI.
        The relative URL must begin with "/" and must be a valid HTTP relative URL.
        It can contain a path, query string arguments, and \# fragments.
        If the relative URL is empty, then the root path "/" will be used.
        No spaces are allowed, and the maximum length allowed is 2083 characters
        """
        return pulumi.get(self, "relative_uri")

    @relative_uri.setter
    def relative_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "relative_uri", value)

    @property
    @pulumi.getter(name="appEngineRouting")
    def app_engine_routing(self) -> Optional[pulumi.Input['JobAppEngineHttpTargetAppEngineRoutingArgs']]:
        """
        App Engine Routing setting for the job.
        Structure is documented below.
        """
        return pulumi.get(self, "app_engine_routing")

    @app_engine_routing.setter
    def app_engine_routing(self, value: Optional[pulumi.Input['JobAppEngineHttpTargetAppEngineRoutingArgs']]):
        pulumi.set(self, "app_engine_routing", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Which HTTP method to use for the request.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)


@pulumi.input_type
class JobAppEngineHttpTargetAppEngineRoutingArgs:
    def __init__(__self__, *,
                 instance: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance: App instance.
               By default, the job is sent to an instance which is available when the job is attempted.
        :param pulumi.Input[str] service: App service.
               By default, the job is sent to the service which is the default service when the job is attempted.
        :param pulumi.Input[str] version: App version.
               By default, the job is sent to the version which is the default version when the job is attempted.
        """
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        App instance.
        By default, the job is sent to an instance which is available when the job is attempted.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        App service.
        By default, the job is sent to the service which is the default service when the job is attempted.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        App version.
        By default, the job is sent to the version which is the default version when the job is attempted.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class JobHttpTargetArgs:
    def __init__(__self__, *,
                 uri: pulumi.Input[str],
                 body: Optional[pulumi.Input[str]] = None,
                 headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 oauth_token: Optional[pulumi.Input['JobHttpTargetOauthTokenArgs']] = None,
                 oidc_token: Optional[pulumi.Input['JobHttpTargetOidcTokenArgs']] = None):
        """
        :param pulumi.Input[str] uri: The full URI path that the request will be sent to.
        :param pulumi.Input[str] body: HTTP request body.
               A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
               It is an error to set body on a job with an incompatible HttpMethod.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] headers: This map contains the header field names and values.
               Repeated headers are not supported, but a header value can contain commas.
        :param pulumi.Input[str] http_method: Which HTTP method to use for the request.
        :param pulumi.Input['JobHttpTargetOauthTokenArgs'] oauth_token: Contains information needed for generating an OAuth token.
               This type of authorization should be used when sending requests to a GCP endpoint.
               Structure is documented below.
        :param pulumi.Input['JobHttpTargetOidcTokenArgs'] oidc_token: Contains information needed for generating an OpenID Connect token.
               This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
               Structure is documented below.
        """
        pulumi.set(__self__, "uri", uri)
        if body is not None:
            pulumi.set(__self__, "body", body)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if oauth_token is not None:
            pulumi.set(__self__, "oauth_token", oauth_token)
        if oidc_token is not None:
            pulumi.set(__self__, "oidc_token", oidc_token)

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Input[str]:
        """
        The full URI path that the request will be sent to.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "uri", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP request body.
        A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        It is an error to set body on a job with an incompatible HttpMethod.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        This map contains the header field names and values.
        Repeated headers are not supported, but a header value can contain commas.
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Which HTTP method to use for the request.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="oauthToken")
    def oauth_token(self) -> Optional[pulumi.Input['JobHttpTargetOauthTokenArgs']]:
        """
        Contains information needed for generating an OAuth token.
        This type of authorization should be used when sending requests to a GCP endpoint.
        Structure is documented below.
        """
        return pulumi.get(self, "oauth_token")

    @oauth_token.setter
    def oauth_token(self, value: Optional[pulumi.Input['JobHttpTargetOauthTokenArgs']]):
        pulumi.set(self, "oauth_token", value)

    @property
    @pulumi.getter(name="oidcToken")
    def oidc_token(self) -> Optional[pulumi.Input['JobHttpTargetOidcTokenArgs']]:
        """
        Contains information needed for generating an OpenID Connect token.
        This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
        Structure is documented below.
        """
        return pulumi.get(self, "oidc_token")

    @oidc_token.setter
    def oidc_token(self, value: Optional[pulumi.Input['JobHttpTargetOidcTokenArgs']]):
        pulumi.set(self, "oidc_token", value)


@pulumi.input_type
class JobHttpTargetOauthTokenArgs:
    def __init__(__self__, *,
                 service_account_email: pulumi.Input[str],
                 scope: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] service_account_email: Service account email to be used for generating OAuth token.
               The service account must be within the same project as the job.
        :param pulumi.Input[str] scope: OAuth scope to be used for generating OAuth access token. If not specified,
               "https://www.googleapis.com/auth/cloud-platform" will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        Service account email to be used for generating OAuth token.
        The service account must be within the same project as the job.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        OAuth scope to be used for generating OAuth access token. If not specified,
        "https://www.googleapis.com/auth/cloud-platform" will be used.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class JobHttpTargetOidcTokenArgs:
    def __init__(__self__, *,
                 service_account_email: pulumi.Input[str],
                 audience: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] service_account_email: Service account email to be used for generating OAuth token.
               The service account must be within the same project as the job.
        :param pulumi.Input[str] audience: Audience to be used when generating OIDC token. If not specified,
               the URI specified in target will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        Service account email to be used for generating OAuth token.
        The service account must be within the same project as the job.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def audience(self) -> Optional[pulumi.Input[str]]:
        """
        Audience to be used when generating OIDC token. If not specified,
        the URI specified in target will be used.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audience", value)


@pulumi.input_type
class JobPubsubTargetArgs:
    def __init__(__self__, *,
                 topic_name: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 data: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] topic_name: The full resource name for the Cloud Pub/Sub topic to which
               messages will be published when a job is delivered. ~>**NOTE:**
               The topic name must be in the same format as required by PubSub's
               PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Attributes for PubsubMessage.
               Pubsub message must contain either non-empty data, or at least one attribute.
        :param pulumi.Input[str] data: The message payload for PubsubMessage.
               Pubsub message must contain either non-empty data, or at least one attribute.
        """
        pulumi.set(__self__, "topic_name", topic_name)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if data is not None:
            pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The full resource name for the Cloud Pub/Sub topic to which
        messages will be published when a job is delivered. ~>**NOTE:**
        The topic name must be in the same format as required by PubSub's
        PublishRequest.name, e.g. `projects/my-project/topics/my-topic`.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Attributes for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The message payload for PubsubMessage.
        Pubsub message must contain either non-empty data, or at least one attribute.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class JobRetryConfigArgs:
    def __init__(__self__, *,
                 max_backoff_duration: Optional[pulumi.Input[str]] = None,
                 max_doublings: Optional[pulumi.Input[float]] = None,
                 max_retry_duration: Optional[pulumi.Input[str]] = None,
                 min_backoff_duration: Optional[pulumi.Input[str]] = None,
                 retry_count: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] max_backoff_duration: The maximum amount of time to wait before retrying a job after it fails.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param pulumi.Input[float] max_doublings: The time between retries will double maxDoublings times.
               A job's retry interval starts at minBackoffDuration,
               then doubles maxDoublings times, then increases linearly,
               and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
        :param pulumi.Input[str] max_retry_duration: The time limit for retrying a failed job, measured from time when an execution was first attempted.
               If specified with retryCount, the job will be retried until both limits are reached.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param pulumi.Input[str] min_backoff_duration: The minimum amount of time to wait before retrying a job after it fails.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
        :param pulumi.Input[float] retry_count: The number of attempts that the system will make to run a
               job using the exponential backoff procedure described by maxDoublings.
               Values greater than 5 and negative values are not allowed.
        """
        if max_backoff_duration is not None:
            pulumi.set(__self__, "max_backoff_duration", max_backoff_duration)
        if max_doublings is not None:
            pulumi.set(__self__, "max_doublings", max_doublings)
        if max_retry_duration is not None:
            pulumi.set(__self__, "max_retry_duration", max_retry_duration)
        if min_backoff_duration is not None:
            pulumi.set(__self__, "min_backoff_duration", min_backoff_duration)
        if retry_count is not None:
            pulumi.set(__self__, "retry_count", retry_count)

    @property
    @pulumi.getter(name="maxBackoffDuration")
    def max_backoff_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "max_backoff_duration")

    @max_backoff_duration.setter
    def max_backoff_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_backoff_duration", value)

    @property
    @pulumi.getter(name="maxDoublings")
    def max_doublings(self) -> Optional[pulumi.Input[float]]:
        """
        The time between retries will double maxDoublings times.
        A job's retry interval starts at minBackoffDuration,
        then doubles maxDoublings times, then increases linearly,
        and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
        """
        return pulumi.get(self, "max_doublings")

    @max_doublings.setter
    def max_doublings(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_doublings", value)

    @property
    @pulumi.getter(name="maxRetryDuration")
    def max_retry_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The time limit for retrying a failed job, measured from time when an execution was first attempted.
        If specified with retryCount, the job will be retried until both limits are reached.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "max_retry_duration")

    @max_retry_duration.setter
    def max_retry_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_retry_duration", value)

    @property
    @pulumi.getter(name="minBackoffDuration")
    def min_backoff_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The minimum amount of time to wait before retrying a job after it fails.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        """
        return pulumi.get(self, "min_backoff_duration")

    @min_backoff_duration.setter
    def min_backoff_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "min_backoff_duration", value)

    @property
    @pulumi.getter(name="retryCount")
    def retry_count(self) -> Optional[pulumi.Input[float]]:
        """
        The number of attempts that the system will make to run a
        job using the exponential backoff procedure described by maxDoublings.
        Values greater than 5 and negative values are not allowed.
        """
        return pulumi.get(self, "retry_count")

    @retry_count.setter
    def retry_count(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "retry_count", value)


