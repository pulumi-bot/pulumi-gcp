# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Key(pulumi.CustomResource):
    key_algorithm: pulumi.Output[str]
    """
    The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
    Valid values are listed at
    [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
    (only used on create)
    """
    name: pulumi.Output[str]
    """
    The name used for this key pair
    """
    pgp_key: pulumi.Output[str]
    """
    An optional PGP key to encrypt the resulting private
    key material. Only used when creating or importing a new key pair. May either be
    a base64-encoded public key or a `keybase:keybaseusername` string for looking up
    in Vault.
    """
    private_key: pulumi.Output[str]
    """
    The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
    service account keys through the CLI or web console. This is only populated when creating a new key, and when no
    `pgp_key` is provided.
    """
    private_key_encrypted: pulumi.Output[str]
    """
    The private key material, base 64 encoded and
    encrypted with the given `pgp_key`. This is only populated when creating a new
    key and `pgp_key` is supplied
    """
    private_key_fingerprint: pulumi.Output[str]
    """
    The MD5 public key fingerprint for the encrypted
    private key. This is only populated when creating a new key and `pgp_key` is supplied
    """
    private_key_type: pulumi.Output[str]
    """
    The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
    """
    public_key: pulumi.Output[str]
    """
    The public key, base64 encoded
    """
    public_key_type: pulumi.Output[str]
    """
    The output format of the public key requested. X509_PEM is the default output format.
    """
    service_account_id: pulumi.Output[str]
    """
    The Service account id of the Key Pair. This can be a string in the format
    `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
    unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.
    """
    valid_after: pulumi.Output[str]
    """
    The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
    """
    valid_before: pulumi.Output[str]
    """
    The key can be used before this timestamp.
    A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
    """
    def __init__(__self__, resource_name, opts=None, key_algorithm=None, pgp_key=None, private_key_type=None, public_key_type=None, service_account_id=None, __name__=None, __opts__=None):
        """
        Creates and manages service account key-pairs, which allow the user to establish identity of a service account outside of GCP. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_algorithm: The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
               Valid values are listed at
               [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
               (only used on create)
        :param pulumi.Input[str] pgp_key: An optional PGP key to encrypt the resulting private
               key material. Only used when creating or importing a new key pair. May either be
               a base64-encoded public key or a `keybase:keybaseusername` string for looking up
               in Vault.
        :param pulumi.Input[str] private_key_type: The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        :param pulumi.Input[str] public_key_type: The output format of the public key requested. X509_PEM is the default output format.
        :param pulumi.Input[str] service_account_id: The Service account id of the Key Pair. This can be a string in the format
               `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`, where `{ACCOUNT}` is the email address or
               unique id of the service account. If the `{ACCOUNT}` syntax is used, the project will be inferred from the account.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account_key.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['key_algorithm'] = key_algorithm

        __props__['pgp_key'] = pgp_key

        __props__['private_key_type'] = private_key_type

        __props__['public_key_type'] = public_key_type

        if service_account_id is None:
            raise TypeError("Missing required property 'service_account_id'")
        __props__['service_account_id'] = service_account_id

        __props__['name'] = None
        __props__['private_key'] = None
        __props__['private_key_encrypted'] = None
        __props__['private_key_fingerprint'] = None
        __props__['public_key'] = None
        __props__['valid_after'] = None
        __props__['valid_before'] = None

        super(Key, __self__).__init__(
            'gcp:serviceAccount/key:Key',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

