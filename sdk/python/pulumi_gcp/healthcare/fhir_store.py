# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class FhirStore(pulumi.CustomResource):
    dataset: pulumi.Output[str]
    """
    Identifies the dataset addressed by this request. Must be in the format
    'projects/{project}/locations/{location}/datasets/{dataset}'
    """
    disable_referential_integrity: pulumi.Output[bool]
    """
    Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
    creation. The default value is false, meaning that the API will enforce referential integrity and fail the
    requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
    will skip referential integrity check. Consequently, operations that rely on references, such as
    Patient.get$everything, will not return all the results if broken references exist.
    ** Changing this property may recreate the FHIR store (removing all data) **
    """
    disable_resource_versioning: pulumi.Output[bool]
    """
    Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
    of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
    versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
    cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
    attempts to read the historical versions.
    ** Changing this property may recreate the FHIR store (removing all data) **
    """
    enable_history_import: pulumi.Output[bool]
    """
    Whether to allow the bulk import API to accept history bundles and directly insert historical resource
    versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
    occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
    will fail with an error.
    ** Changing this property may recreate the FHIR store (removing all data) **
    ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
    """
    enable_update_create: pulumi.Output[bool]
    """
    Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
    operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
    the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
    logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
    identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
    notifications.
    """
    labels: pulumi.Output[dict]
    """
    User-supplied key-value pairs used to organize FHIR stores.
    Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
    conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
    Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
    bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
    No more than 64 labels can be associated with a given store.
    An object containing a list of "key": value pairs.
    Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
    """
    name: pulumi.Output[str]
    """
    The resource name for the FhirStore.
    ** Changing this property may recreate the FHIR store (removing all data) **
    """
    notification_config: pulumi.Output[dict]
    """
    A nested object resource
    Structure is documented below.

      * `pubsubTopic` (`str`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
        PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
        It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
        was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
        project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
        Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
    """
    self_link: pulumi.Output[str]
    """
    The fully qualified name of this dataset
    """
    stream_configs: pulumi.Output[list]
    """
    A list of streaming configs that configure the destinations of streaming export for every resource mutation in
    this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
    resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
    from the list, the server stops streaming to that location. Before adding a new config, you must add the required
    bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
    the order of dozens of seconds) is expected before the results show up in the streaming destination.
    Structure is documented below.

      * `bigqueryDestination` (`dict`) - The destination BigQuery structure that contains both the dataset location and corresponding schema config.
        The output is organized in one table per resource type. The server reuses the existing tables (if any) that
        are named after the resource types, e.g. "Patient", "Observation". When there is no existing table for a given
        resource type, the server attempts to create one.
        See the [streaming config reference](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores#streamconfig) for more details.
        Structure is documented below.
        * `datasetUri` (`str`) - BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId
        * `schemaConfig` (`dict`) - The configuration for the exported BigQuery schema.
          Structure is documented below.
          * `recursiveStructureDepth` (`float`) - The depth for all recursive structures in the output analytics schema. For example, concept in the CodeSystem
            resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called
            concept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default
            value 2. The maximum depth allowed is 5.
          * `schemaType` (`str`) - Specifies the output schema type. Only ANALYTICS is supported at this time.
            * ANALYTICS: Analytics schema defined by the FHIR community.
            See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md.
            Default value is `ANALYTICS`.
            Possible values are `ANALYTICS`.

      * `resourceTypes` (`list`) - Supply a FHIR resource type (such as "Patient" or "Observation"). See
        https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
        an empty list as an intent to stream all the supported resource types in this FHIR store.
    """
    version: pulumi.Output[str]
    """
    The FHIR specification version.
    Default value is `STU3`.
    Possible values are `DSTU2`, `STU3`, and `R4`.
    """
    def __init__(__self__, resource_name, opts=None, dataset=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_history_import=None, enable_update_create=None, labels=None, name=None, notification_config=None, stream_configs=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
        standard for Healthcare information exchange

        To get more information about FhirStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.fhirStores)
        * How-to Guides
            * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[bool] disable_referential_integrity: Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
               creation. The default value is false, meaning that the API will enforce referential integrity and fail the
               requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
               will skip referential integrity check. Consequently, operations that rely on references, such as
               Patient.get$everything, will not return all the results if broken references exist.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[bool] disable_resource_versioning: Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
               of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
               versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
               cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
               attempts to read the historical versions.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[bool] enable_history_import: Whether to allow the bulk import API to accept history bundles and directly insert historical resource
               versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
               occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
               will fail with an error.
               ** Changing this property may recreate the FHIR store (removing all data) **
               ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
        :param pulumi.Input[bool] enable_update_create: Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
               operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
               the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
               logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
               identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
               notifications.
        :param pulumi.Input[dict] labels: User-supplied key-value pairs used to organize FHIR stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the FhirStore.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[dict] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[list] stream_configs: A list of streaming configs that configure the destinations of streaming export for every resource mutation in
               this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
               resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
               from the list, the server stops streaming to that location. Before adding a new config, you must add the required
               bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
               the order of dozens of seconds) is expected before the results show up in the streaming destination.
               Structure is documented below.
        :param pulumi.Input[str] version: The FHIR specification version.
               Default value is `STU3`.
               Possible values are `DSTU2`, `STU3`, and `R4`.

        The **notification_config** object supports the following:

          * `pubsubTopic` (`pulumi.Input[str]`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
            PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
            It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
            was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
            project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
            Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.

        The **stream_configs** object supports the following:

          * `bigqueryDestination` (`pulumi.Input[dict]`) - The destination BigQuery structure that contains both the dataset location and corresponding schema config.
            The output is organized in one table per resource type. The server reuses the existing tables (if any) that
            are named after the resource types, e.g. "Patient", "Observation". When there is no existing table for a given
            resource type, the server attempts to create one.
            See the [streaming config reference](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores#streamconfig) for more details.
            Structure is documented below.
            * `datasetUri` (`pulumi.Input[str]`) - BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId
            * `schemaConfig` (`pulumi.Input[dict]`) - The configuration for the exported BigQuery schema.
              Structure is documented below.
              * `recursiveStructureDepth` (`pulumi.Input[float]`) - The depth for all recursive structures in the output analytics schema. For example, concept in the CodeSystem
                resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called
                concept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default
                value 2. The maximum depth allowed is 5.
              * `schemaType` (`pulumi.Input[str]`) - Specifies the output schema type. Only ANALYTICS is supported at this time.
                * ANALYTICS: Analytics schema defined by the FHIR community.
                See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md.
                Default value is `ANALYTICS`.
                Possible values are `ANALYTICS`.

          * `resourceTypes` (`pulumi.Input[list]`) - Supply a FHIR resource type (such as "Patient" or "Observation"). See
            https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
            an empty list as an intent to stream all the supported resource types in this FHIR store.
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

            if dataset is None:
                raise TypeError("Missing required property 'dataset'")
            __props__['dataset'] = dataset
            __props__['disable_referential_integrity'] = disable_referential_integrity
            __props__['disable_resource_versioning'] = disable_resource_versioning
            __props__['enable_history_import'] = enable_history_import
            __props__['enable_update_create'] = enable_update_create
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['notification_config'] = notification_config
            __props__['stream_configs'] = stream_configs
            __props__['version'] = version
            __props__['self_link'] = None
        super(FhirStore, __self__).__init__(
            'gcp:healthcare/fhirStore:FhirStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dataset=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_history_import=None, enable_update_create=None, labels=None, name=None, notification_config=None, self_link=None, stream_configs=None, version=None):
        """
        Get an existing FhirStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[bool] disable_referential_integrity: Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
               creation. The default value is false, meaning that the API will enforce referential integrity and fail the
               requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
               will skip referential integrity check. Consequently, operations that rely on references, such as
               Patient.get$everything, will not return all the results if broken references exist.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[bool] disable_resource_versioning: Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
               of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
               versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
               cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
               attempts to read the historical versions.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[bool] enable_history_import: Whether to allow the bulk import API to accept history bundles and directly insert historical resource
               versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
               occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
               will fail with an error.
               ** Changing this property may recreate the FHIR store (removing all data) **
               ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
        :param pulumi.Input[bool] enable_update_create: Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
               operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
               the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
               logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
               identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
               notifications.
        :param pulumi.Input[dict] labels: User-supplied key-value pairs used to organize FHIR stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the FhirStore.
               ** Changing this property may recreate the FHIR store (removing all data) **
        :param pulumi.Input[dict] notification_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[list] stream_configs: A list of streaming configs that configure the destinations of streaming export for every resource mutation in
               this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
               resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
               from the list, the server stops streaming to that location. Before adding a new config, you must add the required
               bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
               the order of dozens of seconds) is expected before the results show up in the streaming destination.
               Structure is documented below.
        :param pulumi.Input[str] version: The FHIR specification version.
               Default value is `STU3`.
               Possible values are `DSTU2`, `STU3`, and `R4`.

        The **notification_config** object supports the following:

          * `pubsubTopic` (`pulumi.Input[str]`) - The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
            PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
            It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
            was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
            project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given
            Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.

        The **stream_configs** object supports the following:

          * `bigqueryDestination` (`pulumi.Input[dict]`) - The destination BigQuery structure that contains both the dataset location and corresponding schema config.
            The output is organized in one table per resource type. The server reuses the existing tables (if any) that
            are named after the resource types, e.g. "Patient", "Observation". When there is no existing table for a given
            resource type, the server attempts to create one.
            See the [streaming config reference](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores#streamconfig) for more details.
            Structure is documented below.
            * `datasetUri` (`pulumi.Input[str]`) - BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId
            * `schemaConfig` (`pulumi.Input[dict]`) - The configuration for the exported BigQuery schema.
              Structure is documented below.
              * `recursiveStructureDepth` (`pulumi.Input[float]`) - The depth for all recursive structures in the output analytics schema. For example, concept in the CodeSystem
                resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called
                concept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default
                value 2. The maximum depth allowed is 5.
              * `schemaType` (`pulumi.Input[str]`) - Specifies the output schema type. Only ANALYTICS is supported at this time.
                * ANALYTICS: Analytics schema defined by the FHIR community.
                See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md.
                Default value is `ANALYTICS`.
                Possible values are `ANALYTICS`.

          * `resourceTypes` (`pulumi.Input[list]`) - Supply a FHIR resource type (such as "Patient" or "Observation"). See
            https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
            an empty list as an intent to stream all the supported resource types in this FHIR store.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dataset"] = dataset
        __props__["disable_referential_integrity"] = disable_referential_integrity
        __props__["disable_resource_versioning"] = disable_resource_versioning
        __props__["enable_history_import"] = enable_history_import
        __props__["enable_update_create"] = enable_update_create
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["notification_config"] = notification_config
        __props__["self_link"] = self_link
        __props__["stream_configs"] = stream_configs
        __props__["version"] = version
        return FhirStore(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
