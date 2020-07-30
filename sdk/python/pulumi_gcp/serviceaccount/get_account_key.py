# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetAccountKeyResult:
    """
    A collection of values returned by getAccountKey.
    """
    def __init__(__self__, id=None, key_algorithm=None, name=None, project=None, public_key=None, public_key_type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if key_algorithm and not isinstance(key_algorithm, str):
            raise TypeError("Expected argument 'key_algorithm' to be a str")
        __self__.key_algorithm = key_algorithm
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if public_key and not isinstance(public_key, str):
            raise TypeError("Expected argument 'public_key' to be a str")
        __self__.public_key = public_key
        """
        The public key, base64 encoded
        """
        if public_key_type and not isinstance(public_key_type, str):
            raise TypeError("Expected argument 'public_key_type' to be a str")
        __self__.public_key_type = public_key_type


class AwaitableGetAccountKeyResult(GetAccountKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountKeyResult(
            id=self.id,
            key_algorithm=self.key_algorithm,
            name=self.name,
            project=self.project,
            public_key=self.public_key,
            public_key_type=self.public_key_type)


def get_account_key(name=None, project=None, public_key_type=None, opts=None):
    """
    Get service account public key. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys/get).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    myaccount = gcp.service_account.Account("myaccount", account_id="dev-foo-account")
    mykey_key = gcp.service_account.Key("mykeyKey", service_account_id=myaccount.name)
    mykey_account_key = mykey_key.name.apply(lambda name: gcp.serviceAccount.get_account_key(gcp.service_account.GetAccountKeyArgsArgs(
        name=name,
        public_key_type="TYPE_X509_PEM_FILE",
    )))
    ```


    :param str name: The name of the service account key. This must have format
           `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{KEYID}`, where `{ACCOUNT}`
           is the email address or unique id of the service account.
    :param str project: The ID of the project that the service account will be created in.
           Defaults to the provider project configuration.
    :param str public_key_type: The output format of the public key requested. X509_PEM is the default output format.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['publicKeyType'] = public_key_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:serviceAccount/getAccountKey:getAccountKey', __args__, opts=opts).value

    return AwaitableGetAccountKeyResult(
        id=__ret__.get('id'),
        key_algorithm=__ret__.get('keyAlgorithm'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        public_key=__ret__.get('publicKey'),
        public_key_type=__ret__.get('publicKeyType'))
