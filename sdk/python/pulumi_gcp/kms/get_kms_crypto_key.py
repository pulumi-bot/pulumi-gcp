# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetKMSCryptoKeyResult:
    """
    A collection of values returned by getKMSCryptoKey.
    """
    def __init__(__self__, key_ring=None, labels=None, name=None, purpose=None, rotation_period=None, self_link=None, version_templates=None, id=None):
        if key_ring and not isinstance(key_ring, str):
            raise TypeError("Expected argument 'key_ring' to be a str")
        __self__.key_ring = key_ring
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if purpose and not isinstance(purpose, str):
            raise TypeError("Expected argument 'purpose' to be a str")
        __self__.purpose = purpose
        """
        Defines the cryptographic capabilities of the key.
        """
        if rotation_period and not isinstance(rotation_period, str):
            raise TypeError("Expected argument 'rotation_period' to be a str")
        __self__.rotation_period = rotation_period
        """
        Every time this period passes, generate a new CryptoKeyVersion and set it as
        the primary. The first rotation will take place after the specified period. The rotation period has the format
        of a decimal number with up to 9 fractional digits, followed by the letter s (seconds).
        """
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The self link of the created CryptoKey. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}`.
        """
        if version_templates and not isinstance(version_templates, list):
            raise TypeError("Expected argument 'version_templates' to be a list")
        __self__.version_templates = version_templates
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetKMSCryptoKeyResult(GetKMSCryptoKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSCryptoKeyResult(
            key_ring=self.key_ring,
            labels=self.labels,
            name=self.name,
            purpose=self.purpose,
            rotation_period=self.rotation_period,
            self_link=self.self_link,
            version_templates=self.version_templates,
            id=self.id)

def get_kms_crypto_key(key_ring=None,name=None,opts=None):
    """
    Provides access to a Google Cloud Platform KMS CryptoKey. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys).
    
    A CryptoKey is an interface to key material which can be used to encrypt and decrypt data. A CryptoKey belongs to a
    Google Cloud KMS KeyRing.
    
    :param str key_ring: The `self_link` of the Google Cloud Platform KeyRing to which the key belongs.
    :param str name: The CryptoKey's name.
           A CryptoKey’s name belonging to the specified Google Cloud Platform KeyRing and match the regular expression `[a-zA-Z0-9_-]{1,63}`

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_crypto_key.html.markdown.
    """
    __args__ = dict()

    __args__['keyRing'] = key_ring
    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSCryptoKey:getKMSCryptoKey', __args__, opts=opts).value

    return AwaitableGetKMSCryptoKeyResult(
        key_ring=__ret__.get('keyRing'),
        labels=__ret__.get('labels'),
        name=__ret__.get('name'),
        purpose=__ret__.get('purpose'),
        rotation_period=__ret__.get('rotationPeriod'),
        self_link=__ret__.get('selfLink'),
        version_templates=__ret__.get('versionTemplates'),
        id=__ret__.get('id'))
