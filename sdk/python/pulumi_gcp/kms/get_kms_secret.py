# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetKMSSecretResult(object):
    """
    A collection of values returned by getKMSSecret.
    """
    def __init__(__self__, plaintext=None, id=None):
        if plaintext and not isinstance(plaintext, basestring):
            raise TypeError('Expected argument plaintext to be a basestring')
        __self__.plaintext = plaintext
        """
        Contains the result of decrypting the provided ciphertext.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_kms_secret(ciphertext=None, crypto_key=None):
    """
    This data source allows you to use data encrypted with Google Cloud KMS
    within your resource definitions.
    
    For more information see
    [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
    
    ~> **NOTE**: Using this data provider will allow you to conceal secret data within your
    resource definitions, but it does not take care of protecting that data in the
    logging output, plan output, or state output.  Please take care to secure your secret
    data outside of resource definitions.
    """
    __args__ = dict()

    __args__['ciphertext'] = ciphertext
    __args__['cryptoKey'] = crypto_key
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSSecret:getKMSSecret', __args__)

    return GetKMSSecretResult(
        plaintext=__ret__.get('plaintext'),
        id=__ret__.get('id'))
