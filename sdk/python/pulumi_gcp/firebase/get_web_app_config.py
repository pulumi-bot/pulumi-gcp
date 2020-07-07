# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetWebAppConfigResult:
    """
    A collection of values returned by getWebAppConfig.
    """
    def __init__(__self__, api_key=None, auth_domain=None, database_url=None, id=None, location_id=None, measurement_id=None, messaging_sender_id=None, project=None, storage_bucket=None, web_app_id=None):
        if api_key and not isinstance(api_key, str):
            raise TypeError("Expected argument 'api_key' to be a str")
        __self__.api_key = api_key
        if auth_domain and not isinstance(auth_domain, str):
            raise TypeError("Expected argument 'auth_domain' to be a str")
        __self__.auth_domain = auth_domain
        if database_url and not isinstance(database_url, str):
            raise TypeError("Expected argument 'database_url' to be a str")
        __self__.database_url = database_url
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if location_id and not isinstance(location_id, str):
            raise TypeError("Expected argument 'location_id' to be a str")
        __self__.location_id = location_id
        if measurement_id and not isinstance(measurement_id, str):
            raise TypeError("Expected argument 'measurement_id' to be a str")
        __self__.measurement_id = measurement_id
        if messaging_sender_id and not isinstance(messaging_sender_id, str):
            raise TypeError("Expected argument 'messaging_sender_id' to be a str")
        __self__.messaging_sender_id = messaging_sender_id
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if storage_bucket and not isinstance(storage_bucket, str):
            raise TypeError("Expected argument 'storage_bucket' to be a str")
        __self__.storage_bucket = storage_bucket
        if web_app_id and not isinstance(web_app_id, str):
            raise TypeError("Expected argument 'web_app_id' to be a str")
        __self__.web_app_id = web_app_id


class AwaitableGetWebAppConfigResult(GetWebAppConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWebAppConfigResult(
            api_key=self.api_key,
            auth_domain=self.auth_domain,
            database_url=self.database_url,
            id=self.id,
            location_id=self.location_id,
            measurement_id=self.measurement_id,
            messaging_sender_id=self.messaging_sender_id,
            project=self.project,
            storage_bucket=self.storage_bucket,
            web_app_id=self.web_app_id)


def get_web_app_config(project=None, web_app_id=None, opts=None):
    """
    A Google Cloud Firebase web application configuration

    To get more information about WebApp, see:

    * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
    * How-to Guides
        * [Official Documentation](https://firebase.google.com/)


    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str web_app_id: the id of the firebase web app
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['webAppId'] = web_app_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:firebase/getWebAppConfig:getWebAppConfig', __args__, opts=opts).value

    return AwaitableGetWebAppConfigResult(
        api_key=__ret__.get('apiKey'),
        auth_domain=__ret__.get('authDomain'),
        database_url=__ret__.get('databaseUrl'),
        id=__ret__.get('id'),
        location_id=__ret__.get('locationId'),
        measurement_id=__ret__.get('measurementId'),
        messaging_sender_id=__ret__.get('messagingSenderId'),
        project=__ret__.get('project'),
        storage_bucket=__ret__.get('storageBucket'),
        web_app_id=__ret__.get('webAppId'))
