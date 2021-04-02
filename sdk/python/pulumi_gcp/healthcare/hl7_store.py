# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Hl7StoreArgs', 'Hl7Store']

@pulumi.input_type
class Hl7StoreArgs:
    def __init__(__self__, *,
                 dataset: pulumi.Input[str],
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input['Hl7StoreNotificationConfigArgs']] = None,
                 notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input['Hl7StoreNotificationConfigsArgs']]]] = None,
                 parser_config: Optional[pulumi.Input['Hl7StoreParserConfigArgs']] = None):
        """
        The set of arguments for constructing a Hl7Store resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize HL7v2 stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the Hl7V2Store.
               ** Changing this property may recreate the Hl7v2 store (removing all data) **
        :param pulumi.Input['Hl7StoreNotificationConfigArgs'] notification_config: -
               (Optional, Deprecated)
               A nested object resource
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['Hl7StoreNotificationConfigsArgs']]] notification_configs: A list of notification configs. Each configuration uses a filter to determine whether to publish a
               message (both Ingest & Create) on the corresponding notification destination. Only the message name
               is sent as part of the notification. Supplied by the client.
               Structure is documented below.
        :param pulumi.Input['Hl7StoreParserConfigArgs'] parser_config: A nested object resource
               Structure is documented below.
        """
        pulumi.set(__self__, "dataset", dataset)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_config is not None:
            warnings.warn("""This field has been replaced by notificationConfigs""", DeprecationWarning)
            pulumi.log.warn("""notification_config is deprecated: This field has been replaced by notificationConfigs""")
        if notification_config is not None:
            pulumi.set(__self__, "notification_config", notification_config)
        if notification_configs is not None:
            pulumi.set(__self__, "notification_configs", notification_configs)
        if parser_config is not None:
            pulumi.set(__self__, "parser_config", parser_config)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key-value pairs used to organize HL7v2 stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the Hl7V2Store.
        ** Changing this property may recreate the Hl7v2 store (removing all data) **
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> Optional[pulumi.Input['Hl7StoreNotificationConfigArgs']]:
        """
        -
        (Optional, Deprecated)
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @notification_config.setter
    def notification_config(self, value: Optional[pulumi.Input['Hl7StoreNotificationConfigArgs']]):
        pulumi.set(self, "notification_config", value)

    @property
    @pulumi.getter(name="notificationConfigs")
    def notification_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['Hl7StoreNotificationConfigsArgs']]]]:
        """
        A list of notification configs. Each configuration uses a filter to determine whether to publish a
        message (both Ingest & Create) on the corresponding notification destination. Only the message name
        is sent as part of the notification. Supplied by the client.
        Structure is documented below.
        """
        return pulumi.get(self, "notification_configs")

    @notification_configs.setter
    def notification_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['Hl7StoreNotificationConfigsArgs']]]]):
        pulumi.set(self, "notification_configs", value)

    @property
    @pulumi.getter(name="parserConfig")
    def parser_config(self) -> Optional[pulumi.Input['Hl7StoreParserConfigArgs']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "parser_config")

    @parser_config.setter
    def parser_config(self, value: Optional[pulumi.Input['Hl7StoreParserConfigArgs']]):
        pulumi.set(self, "parser_config", value)


class Hl7Store(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigArgs']]] = None,
                 notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigsArgs']]]]] = None,
                 parser_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreParserConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/hl7V2/STU3/)
        standard for Healthcare information exchange

        To get more information about Hl7V2Store, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.hl7V2Stores)
        * How-to Guides
            * [Creating a HL7v2 Store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2)

        ## Example Usage
        ### Healthcare Hl7 V2 Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic")
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            notification_configs=[gcp.healthcare.Hl7StoreNotificationConfigsArgs(
                pubsub_topic=topic.id,
            )],
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Hl7 V2 Store Parser Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            parser_config=gcp.healthcare.Hl7StoreParserConfigArgs(
                allow_null_header=False,
                segment_terminator="Jw==",
                schema=\"\"\"{
          "schemas": [{
            "messageSchemaConfigs": {
              "ADT_A01": {
                "name": "ADT_A01",
                "minOccurs": 1,
                "maxOccurs": 1,
                "members": [{
                    "segment": {
                      "type": "MSH",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "EVN",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "PID",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "ZPD",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "OBX"
                    }
                  },
                  {
                    "group": {
                      "name": "PROCEDURE",
                      "members": [{
                          "segment": {
                            "type": "PR1",
                            "minOccurs": 1,
                            "maxOccurs": 1
                          }
                        },
                        {
                          "segment": {
                            "type": "ROL"
                          }
                        }
                      ]
                    }
                  },
                  {
                    "segment": {
                      "type": "PDA",
                      "maxOccurs": 1
                    }
                  }
                ]
              }
            }
          }],
          "types": [{
            "type": [{
                "name": "ZPD",
                "primitive": "VARIES"
              }

            ]
          }],
          "ignoreMinOccurs": true
        }
        \"\"\",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Healthcare Hl7 V2 Store Unschematized

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            parser_config=gcp.healthcare.Hl7StoreParserConfigArgs(
                allow_null_header=False,
                segment_terminator="Jw==",
                version="V2",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Hl7V2Store can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/hl7V2Stores/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize HL7v2 stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the Hl7V2Store.
               ** Changing this property may recreate the Hl7v2 store (removing all data) **
        :param pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigArgs']] notification_config: -
               (Optional, Deprecated)
               A nested object resource
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigsArgs']]]] notification_configs: A list of notification configs. Each configuration uses a filter to determine whether to publish a
               message (both Ingest & Create) on the corresponding notification destination. Only the message name
               is sent as part of the notification. Supplied by the client.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['Hl7StoreParserConfigArgs']] parser_config: A nested object resource
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Hl7StoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/hl7V2/STU3/)
        standard for Healthcare information exchange

        To get more information about Hl7V2Store, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.hl7V2Stores)
        * How-to Guides
            * [Creating a HL7v2 Store](https://cloud.google.com/healthcare/docs/how-tos/hl7v2)

        ## Example Usage
        ### Healthcare Hl7 V2 Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        topic = gcp.pubsub.Topic("topic")
        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            notification_configs=[gcp.healthcare.Hl7StoreNotificationConfigsArgs(
                pubsub_topic=topic.id,
            )],
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Hl7 V2 Store Parser Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            parser_config=gcp.healthcare.Hl7StoreParserConfigArgs(
                allow_null_header=False,
                segment_terminator="Jw==",
                schema=\"\"\"{
          "schemas": [{
            "messageSchemaConfigs": {
              "ADT_A01": {
                "name": "ADT_A01",
                "minOccurs": 1,
                "maxOccurs": 1,
                "members": [{
                    "segment": {
                      "type": "MSH",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "EVN",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "PID",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "ZPD",
                      "minOccurs": 1,
                      "maxOccurs": 1
                    }
                  },
                  {
                    "segment": {
                      "type": "OBX"
                    }
                  },
                  {
                    "group": {
                      "name": "PROCEDURE",
                      "members": [{
                          "segment": {
                            "type": "PR1",
                            "minOccurs": 1,
                            "maxOccurs": 1
                          }
                        },
                        {
                          "segment": {
                            "type": "ROL"
                          }
                        }
                      ]
                    }
                  },
                  {
                    "segment": {
                      "type": "PDA",
                      "maxOccurs": 1
                    }
                  }
                ]
              }
            }
          }],
          "types": [{
            "type": [{
                "name": "ZPD",
                "primitive": "VARIES"
              }

            ]
          }],
          "ignoreMinOccurs": true
        }
        \"\"\",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Healthcare Hl7 V2 Store Unschematized

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1",
        opts=pulumi.ResourceOptions(provider=google_beta))
        store = gcp.healthcare.Hl7Store("store",
            dataset=dataset.id,
            parser_config=gcp.healthcare.Hl7StoreParserConfigArgs(
                allow_null_header=False,
                segment_terminator="Jw==",
                version="V2",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Hl7V2Store can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/hl7V2Stores/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/hl7Store:Hl7Store default {{dataset}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param Hl7StoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Hl7StoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigArgs']]] = None,
                 notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigsArgs']]]]] = None,
                 parser_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreParserConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if dataset is None and not opts.urn:
                raise TypeError("Missing required property 'dataset'")
            __props__['dataset'] = dataset
            __props__['labels'] = labels
            __props__['name'] = name
            if notification_config is not None and not opts.urn:
                warnings.warn("""This field has been replaced by notificationConfigs""", DeprecationWarning)
                pulumi.log.warn("""notification_config is deprecated: This field has been replaced by notificationConfigs""")
            __props__['notification_config'] = notification_config
            __props__['notification_configs'] = notification_configs
            __props__['parser_config'] = parser_config
            __props__['self_link'] = None
        super(Hl7Store, __self__).__init__(
            'gcp:healthcare/hl7Store:Hl7Store',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigArgs']]] = None,
            notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigsArgs']]]]] = None,
            parser_config: Optional[pulumi.Input[pulumi.InputType['Hl7StoreParserConfigArgs']]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'Hl7Store':
        """
        Get an existing Hl7Store resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize HL7v2 stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name for the Hl7V2Store.
               ** Changing this property may recreate the Hl7v2 store (removing all data) **
        :param pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigArgs']] notification_config: -
               (Optional, Deprecated)
               A nested object resource
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['Hl7StoreNotificationConfigsArgs']]]] notification_configs: A list of notification configs. Each configuration uses a filter to determine whether to publish a
               message (both Ingest & Create) on the corresponding notification destination. Only the message name
               is sent as part of the notification. Supplied by the client.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['Hl7StoreParserConfigArgs']] parser_config: A nested object resource
               Structure is documented below.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dataset"] = dataset
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["notification_config"] = notification_config
        __props__["notification_configs"] = notification_configs
        __props__["parser_config"] = parser_config
        __props__["self_link"] = self_link
        return Hl7Store(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-supplied key-value pairs used to organize HL7v2 stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the Hl7V2Store.
        ** Changing this property may recreate the Hl7v2 store (removing all data) **
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> pulumi.Output[Optional['outputs.Hl7StoreNotificationConfig']]:
        """
        -
        (Optional, Deprecated)
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "notification_config")

    @property
    @pulumi.getter(name="notificationConfigs")
    def notification_configs(self) -> pulumi.Output[Optional[Sequence['outputs.Hl7StoreNotificationConfigs']]]:
        """
        A list of notification configs. Each configuration uses a filter to determine whether to publish a
        message (both Ingest & Create) on the corresponding notification destination. Only the message name
        is sent as part of the notification. Supplied by the client.
        Structure is documented below.
        """
        return pulumi.get(self, "notification_configs")

    @property
    @pulumi.getter(name="parserConfig")
    def parser_config(self) -> pulumi.Output[Optional['outputs.Hl7StoreParserConfig']]:
        """
        A nested object resource
        Structure is documented below.
        """
        return pulumi.get(self, "parser_config")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

