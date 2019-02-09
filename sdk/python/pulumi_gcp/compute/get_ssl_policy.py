# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSSLPolicyResult(object):
    """
    A collection of values returned by getSSLPolicy.
    """
    def __init__(__self__, creation_timestamp=None, custom_features=None, description=None, enabled_features=None, fingerprint=None, min_tls_version=None, profile=None, self_link=None, id=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError('Expected argument creation_timestamp to be a str')
        __self__.creation_timestamp = creation_timestamp
        if custom_features and not isinstance(custom_features, list):
            raise TypeError('Expected argument custom_features to be a list')
        __self__.custom_features = custom_features
        """
        If the `profile` is `CUSTOM`, these are the custom encryption
        ciphers supported by the profile. If the `profile` is *not* `CUSTOM`, this
        attribute will be empty.
        """
        if description and not isinstance(description, str):
            raise TypeError('Expected argument description to be a str')
        __self__.description = description
        """
        Description of this SSL Policy.
        """
        if enabled_features and not isinstance(enabled_features, list):
            raise TypeError('Expected argument enabled_features to be a list')
        __self__.enabled_features = enabled_features
        """
        The set of enabled encryption ciphers as a result of the policy config
        """
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError('Expected argument fingerprint to be a str')
        __self__.fingerprint = fingerprint
        """
        Fingerprint of this resource.
        """
        if min_tls_version and not isinstance(min_tls_version, str):
            raise TypeError('Expected argument min_tls_version to be a str')
        __self__.min_tls_version = min_tls_version
        """
        The minimum supported TLS version of this policy.
        """
        if profile and not isinstance(profile, str):
            raise TypeError('Expected argument profile to be a str')
        __self__.profile = profile
        """
        The Google-curated or custom profile used by this policy.
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError('Expected argument self_link to be a str')
        __self__.self_link = self_link
        """
        The URI of the created resource.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_ssl_policy(name=None, project=None):
    """
    Gets an SSL Policy within GCE from its name, for use with Target HTTPS and Target SSL Proxies.
        For more information see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies).
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __ret__ = await pulumi.runtime.invoke('gcp:compute/getSSLPolicy:getSSLPolicy', __args__)

    return GetSSLPolicyResult(
        creation_timestamp=__ret__.get('creationTimestamp'),
        custom_features=__ret__.get('customFeatures'),
        description=__ret__.get('description'),
        enabled_features=__ret__.get('enabledFeatures'),
        fingerprint=__ret__.get('fingerprint'),
        min_tls_version=__ret__.get('minTlsVersion'),
        profile=__ret__.get('profile'),
        self_link=__ret__.get('selfLink'),
        id=__ret__.get('id'))
