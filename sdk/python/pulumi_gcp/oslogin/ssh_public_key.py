# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SshPublicKey(pulumi.CustomResource):
    expiration_time_usec: pulumi.Output[str]
    """
    An expiration time in microseconds since epoch.
    """
    fingerprint: pulumi.Output[str]
    """
    The SHA-256 fingerprint of the SSH public key.
    """
    key: pulumi.Output[str]
    """
    Public key text in SSH format, defined by RFC4253 section 6.6.
    """
    user: pulumi.Output[str]
    """
    The user email.
    """
    def __init__(__self__, resource_name, opts=None, expiration_time_usec=None, key=None, user=None, __props__=None, __name__=None, __opts__=None):
        """
        The SSH public key information associated with a Google account.


        To get more information about SSHPublicKey, see:

        * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)

        ## Example Usage

        ### Os Login Ssh Key Provided User

        ```python
        import pulumi
        import pulumi_gcp as gcp

        me = gcp.organizations.get_client_open_id_user_info()
        cache = gcp.oslogin.SshPublicKey("cache",
            user=me.email,
            key=(lambda path: open(path).read())("path/to/id_rsa.pub"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] user: The user email.
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

            __props__['expiration_time_usec'] = expiration_time_usec
            if key is None:
                raise TypeError("Missing required property 'key'")
            __props__['key'] = key
            if user is None:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
            __props__['fingerprint'] = None
        super(SshPublicKey, __self__).__init__(
            'gcp:oslogin/sshPublicKey:SshPublicKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, expiration_time_usec=None, fingerprint=None, key=None, user=None):
        """
        Get an existing SshPublicKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] fingerprint: The SHA-256 fingerprint of the SSH public key.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] user: The user email.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["expiration_time_usec"] = expiration_time_usec
        __props__["fingerprint"] = fingerprint
        __props__["key"] = key
        __props__["user"] = user
        return SshPublicKey(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
