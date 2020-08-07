# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Instance(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    """
    Creation timestamp in RFC3339 text format.
    """
    description: pulumi.Output[str]
    """
    A description of the instance.
    """
    etag: pulumi.Output[str]
    """
    Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
    """
    file_shares: pulumi.Output[dict]
    """
    File system shares on the instance. For this version, only a
    single file share is supported.  Structure is documented below.

      * `capacityGb` (`float`) - File share capacity in GiB. This must be at least 1024 GiB
        for the standard tier, or 2560 GiB for the premium tier.
      * `name` (`str`) - The name of the fileshare (16 characters or less)
      * `nfsExportOptions` (`list`)
        * `accessMode` (`str`) - Either READ_ONLY, for allowing only read requests on the exported directory,
          or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
        * `anonGid` (`float`) - An integer representing the anonymous group id with a default value of 65534.
          Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
          if this field is specified for other squashMode settings.
        * `anonUid` (`float`) - An integer representing the anonymous user id with a default value of 65534.
          Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
          if this field is specified for other squashMode settings.
        * `ipRanges` (`list`) - List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
          Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
          The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        * `squashMode` (`str`) - Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
          for not allowing root access. The default is NO_ROOT_SQUASH.
    """
    labels: pulumi.Output[dict]
    """
    Resource labels to represent user-provided metadata.
    """
    name: pulumi.Output[str]
    """
    The name of the fileshare (16 characters or less)
    """
    networks: pulumi.Output[list]
    """
    VPC networks to which the instance is connected. For this version,
    only a single network is supported.  Structure is documented below.

      * `ip_addresses` (`list`) - -
        A list of IPv4 or IPv6 addresses.
      * `modes` (`list`) - IP versions for which the instance has
        IP addresses assigned.
      * `network` (`str`) - The name of the GCE VPC network to which the
        instance is connected.
      * `reserved_ip_range` (`str`) - A /29 CIDR block that identifies the range of IP
        addresses reserved for this instance.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    tier: pulumi.Output[str]
    """
    The service tier of the instance.
    """
    zone: pulumi.Output[str]
    """
    The name of the Filestore zone of the instance.
    """
    def __init__(__self__, resource_name, opts=None, description=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        A Google Cloud Filestore instance.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
            * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
            * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)

        ## Example Usage
        ### Filestore Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.filestore.Instance("instance",
            file_shares={
                "capacityGb": 2660,
                "name": "share1",
            },
            networks=[{
                "modes": ["MODE_IPV4"],
                "network": "default",
            }],
            tier="PREMIUM",
            zone="us-central1-b")
        ```
        ### Filestore Instance Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.filestore.Instance("instance",
            zone="us-central1-b",
            tier="BASIC_SSD",
            file_shares={
                "capacityGb": 2660,
                "name": "share1",
                "nfsExportOptions": [
                    {
                        "ipRanges": ["10.0.0.0/24"],
                        "accessMode": "READ_WRITE",
                        "squashMode": "NO_ROOT_SQUASH",
                    },
                    {
                        "ipRanges": ["10.10.0.0/24"],
                        "accessMode": "READ_ONLY",
                        "squashMode": "ROOT_SQUASH",
                        "anonUid": 123,
                        "anonGid": 456,
                    },
                ],
            },
            networks=[{
                "network": "default",
                "modes": ["MODE_IPV4"],
            }],
            opts=ResourceOptions(provider=google_beta))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the instance.
        :param pulumi.Input[dict] file_shares: File system shares on the instance. For this version, only a
               single file share is supported.  Structure is documented below.
        :param pulumi.Input[dict] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: The name of the fileshare (16 characters or less)
        :param pulumi.Input[list] networks: VPC networks to which the instance is connected. For this version,
               only a single network is supported.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] tier: The service tier of the instance.
        :param pulumi.Input[str] zone: The name of the Filestore zone of the instance.

        The **file_shares** object supports the following:

          * `capacityGb` (`pulumi.Input[float]`) - File share capacity in GiB. This must be at least 1024 GiB
            for the standard tier, or 2560 GiB for the premium tier.
          * `name` (`pulumi.Input[str]`) - The name of the fileshare (16 characters or less)
          * `nfsExportOptions` (`pulumi.Input[list]`)
            * `accessMode` (`pulumi.Input[str]`) - Either READ_ONLY, for allowing only read requests on the exported directory,
              or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
            * `anonGid` (`pulumi.Input[float]`) - An integer representing the anonymous group id with a default value of 65534.
              Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
              if this field is specified for other squashMode settings.
            * `anonUid` (`pulumi.Input[float]`) - An integer representing the anonymous user id with a default value of 65534.
              Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
              if this field is specified for other squashMode settings.
            * `ipRanges` (`pulumi.Input[list]`) - List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
              Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
              The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
            * `squashMode` (`pulumi.Input[str]`) - Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
              for not allowing root access. The default is NO_ROOT_SQUASH.

        The **networks** object supports the following:

          * `ip_addresses` (`pulumi.Input[list]`) - -
            A list of IPv4 or IPv6 addresses.
          * `modes` (`pulumi.Input[list]`) - IP versions for which the instance has
            IP addresses assigned.
          * `network` (`pulumi.Input[str]`) - The name of the GCE VPC network to which the
            instance is connected.
          * `reserved_ip_range` (`pulumi.Input[str]`) - A /29 CIDR block that identifies the range of IP
            addresses reserved for this instance.
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

            __props__['description'] = description
            if file_shares is None:
                raise TypeError("Missing required property 'file_shares'")
            __props__['file_shares'] = file_shares
            __props__['labels'] = labels
            __props__['name'] = name
            if networks is None:
                raise TypeError("Missing required property 'networks'")
            __props__['networks'] = networks
            __props__['project'] = project
            if tier is None:
                raise TypeError("Missing required property 'tier'")
            __props__['tier'] = tier
            if zone is None:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
            __props__['create_time'] = None
            __props__['etag'] = None
        super(Instance, __self__).__init__(
            'gcp:filestore/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, description=None, etag=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None):
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: A description of the instance.
        :param pulumi.Input[str] etag: Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
        :param pulumi.Input[dict] file_shares: File system shares on the instance. For this version, only a
               single file share is supported.  Structure is documented below.
        :param pulumi.Input[dict] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: The name of the fileshare (16 characters or less)
        :param pulumi.Input[list] networks: VPC networks to which the instance is connected. For this version,
               only a single network is supported.  Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] tier: The service tier of the instance.
        :param pulumi.Input[str] zone: The name of the Filestore zone of the instance.

        The **file_shares** object supports the following:

          * `capacityGb` (`pulumi.Input[float]`) - File share capacity in GiB. This must be at least 1024 GiB
            for the standard tier, or 2560 GiB for the premium tier.
          * `name` (`pulumi.Input[str]`) - The name of the fileshare (16 characters or less)
          * `nfsExportOptions` (`pulumi.Input[list]`)
            * `accessMode` (`pulumi.Input[str]`) - Either READ_ONLY, for allowing only read requests on the exported directory,
              or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
            * `anonGid` (`pulumi.Input[float]`) - An integer representing the anonymous group id with a default value of 65534.
              Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
              if this field is specified for other squashMode settings.
            * `anonUid` (`pulumi.Input[float]`) - An integer representing the anonymous user id with a default value of 65534.
              Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
              if this field is specified for other squashMode settings.
            * `ipRanges` (`pulumi.Input[list]`) - List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
              Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
              The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
            * `squashMode` (`pulumi.Input[str]`) - Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
              for not allowing root access. The default is NO_ROOT_SQUASH.

        The **networks** object supports the following:

          * `ip_addresses` (`pulumi.Input[list]`) - -
            A list of IPv4 or IPv6 addresses.
          * `modes` (`pulumi.Input[list]`) - IP versions for which the instance has
            IP addresses assigned.
          * `network` (`pulumi.Input[str]`) - The name of the GCE VPC network to which the
            instance is connected.
          * `reserved_ip_range` (`pulumi.Input[str]`) - A /29 CIDR block that identifies the range of IP
            addresses reserved for this instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["file_shares"] = file_shares
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["networks"] = networks
        __props__["project"] = project
        __props__["tier"] = tier
        __props__["zone"] = zone
        return Instance(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
