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

__all__ = ['RoutineArgs', 'Routine']

@pulumi.input_type
class RoutineArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 definition_body: pulumi.Input[str],
                 routine_id: pulumi.Input[str],
                 arguments: Optional[pulumi.Input[Sequence[pulumi.Input['RoutineArgumentArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 determinism_level: Optional[pulumi.Input[str]] = None,
                 imported_libraries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 return_type: Optional[pulumi.Input[str]] = None,
                 routine_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Routine resource.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this routine
        :param pulumi.Input[str] definition_body: The body of the routine. For functions, this is the expression in the AS clause.
               If language=SQL, it is the substring inside (but excluding) the parentheses.
        :param pulumi.Input[str] routine_id: The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        :param pulumi.Input[Sequence[pulumi.Input['RoutineArgumentArgs']]] arguments: Input/output argument of a function or a stored procedure.
               Structure is documented below.
        :param pulumi.Input[str] description: The description of the routine if defined.
        :param pulumi.Input[str] determinism_level: The determinism level of the JavaScript UDF if defined.
               Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] imported_libraries: Optional. If language = "JAVASCRIPT", this field stores the path of the
               imported JAVASCRIPT libraries.
        :param pulumi.Input[str] language: The language of the routine.
               Possible values are `SQL` and `JAVASCRIPT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] return_type: A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
               If absent, the return type is inferred from definitionBody at query time in each query
               that references this routine. If present, then the evaluated result will be cast to
               the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
               string, any changes to the string will create a diff, even if the JSON itself hasn't
               changed. If the API returns a different value for the same schema, e.g. it switche
               d the order of values or replaced STRUCT field type with RECORD field type, we currently
               cannot suppress the recurring diff this causes. As a workaround, we recommend using
               the schema as returned by the API.
        :param pulumi.Input[str] routine_type: The type of routine.
               Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "definition_body", definition_body)
        pulumi.set(__self__, "routine_id", routine_id)
        if arguments is not None:
            pulumi.set(__self__, "arguments", arguments)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if determinism_level is not None:
            pulumi.set(__self__, "determinism_level", determinism_level)
        if imported_libraries is not None:
            pulumi.set(__self__, "imported_libraries", imported_libraries)
        if language is not None:
            pulumi.set(__self__, "language", language)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if return_type is not None:
            pulumi.set(__self__, "return_type", return_type)
        if routine_type is not None:
            pulumi.set(__self__, "routine_type", routine_type)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        """
        The ID of the dataset containing this routine
        """
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter(name="definitionBody")
    def definition_body(self) -> pulumi.Input[str]:
        """
        The body of the routine. For functions, this is the expression in the AS clause.
        If language=SQL, it is the substring inside (but excluding) the parentheses.
        """
        return pulumi.get(self, "definition_body")

    @definition_body.setter
    def definition_body(self, value: pulumi.Input[str]):
        pulumi.set(self, "definition_body", value)

    @property
    @pulumi.getter(name="routineId")
    def routine_id(self) -> pulumi.Input[str]:
        """
        The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        """
        return pulumi.get(self, "routine_id")

    @routine_id.setter
    def routine_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "routine_id", value)

    @property
    @pulumi.getter
    def arguments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RoutineArgumentArgs']]]]:
        """
        Input/output argument of a function or a stored procedure.
        Structure is documented below.
        """
        return pulumi.get(self, "arguments")

    @arguments.setter
    def arguments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RoutineArgumentArgs']]]]):
        pulumi.set(self, "arguments", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the routine if defined.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="determinismLevel")
    def determinism_level(self) -> Optional[pulumi.Input[str]]:
        """
        The determinism level of the JavaScript UDF if defined.
        Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
        """
        return pulumi.get(self, "determinism_level")

    @determinism_level.setter
    def determinism_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "determinism_level", value)

    @property
    @pulumi.getter(name="importedLibraries")
    def imported_libraries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. If language = "JAVASCRIPT", this field stores the path of the
        imported JAVASCRIPT libraries.
        """
        return pulumi.get(self, "imported_libraries")

    @imported_libraries.setter
    def imported_libraries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "imported_libraries", value)

    @property
    @pulumi.getter
    def language(self) -> Optional[pulumi.Input[str]]:
        """
        The language of the routine.
        Possible values are `SQL` and `JAVASCRIPT`.
        """
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="returnType")
    def return_type(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
        If absent, the return type is inferred from definitionBody at query time in each query
        that references this routine. If present, then the evaluated result will be cast to
        the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
        string, any changes to the string will create a diff, even if the JSON itself hasn't
        changed. If the API returns a different value for the same schema, e.g. it switche
        d the order of values or replaced STRUCT field type with RECORD field type, we currently
        cannot suppress the recurring diff this causes. As a workaround, we recommend using
        the schema as returned by the API.
        """
        return pulumi.get(self, "return_type")

    @return_type.setter
    def return_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "return_type", value)

    @property
    @pulumi.getter(name="routineType")
    def routine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of routine.
        Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
        """
        return pulumi.get(self, "routine_type")

    @routine_type.setter
    def routine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routine_type", value)


class Routine(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutineArgumentArgs']]]]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 definition_body: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 determinism_level: Optional[pulumi.Input[str]] = None,
                 imported_libraries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 return_type: Optional[pulumi.Input[str]] = None,
                 routine_id: Optional[pulumi.Input[str]] = None,
                 routine_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A user-defined function or a stored procedure that belongs to a Dataset

        To get more information about Routine, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)
        * How-to Guides
            * [Routines Intro](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)

        ## Example Usage
        ### Big Query Routine Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test = gcp.bigquery.Dataset("test", dataset_id="dataset_id")
        sproc = gcp.bigquery.Routine("sproc",
            dataset_id=test.dataset_id,
            routine_id="routine_id",
            routine_type="PROCEDURE",
            language="SQL",
            definition_body="CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);")
        ```

        ## Import

        Routine can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default {{project}}/{{dataset_id}}/{{routine_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default {{dataset_id}}/{{routine_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutineArgumentArgs']]]] arguments: Input/output argument of a function or a stored procedure.
               Structure is documented below.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this routine
        :param pulumi.Input[str] definition_body: The body of the routine. For functions, this is the expression in the AS clause.
               If language=SQL, it is the substring inside (but excluding) the parentheses.
        :param pulumi.Input[str] description: The description of the routine if defined.
        :param pulumi.Input[str] determinism_level: The determinism level of the JavaScript UDF if defined.
               Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] imported_libraries: Optional. If language = "JAVASCRIPT", this field stores the path of the
               imported JAVASCRIPT libraries.
        :param pulumi.Input[str] language: The language of the routine.
               Possible values are `SQL` and `JAVASCRIPT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] return_type: A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
               If absent, the return type is inferred from definitionBody at query time in each query
               that references this routine. If present, then the evaluated result will be cast to
               the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
               string, any changes to the string will create a diff, even if the JSON itself hasn't
               changed. If the API returns a different value for the same schema, e.g. it switche
               d the order of values or replaced STRUCT field type with RECORD field type, we currently
               cannot suppress the recurring diff this causes. As a workaround, we recommend using
               the schema as returned by the API.
        :param pulumi.Input[str] routine_id: The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        :param pulumi.Input[str] routine_type: The type of routine.
               Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoutineArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A user-defined function or a stored procedure that belongs to a Dataset

        To get more information about Routine, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)
        * How-to Guides
            * [Routines Intro](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)

        ## Example Usage
        ### Big Query Routine Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test = gcp.bigquery.Dataset("test", dataset_id="dataset_id")
        sproc = gcp.bigquery.Routine("sproc",
            dataset_id=test.dataset_id,
            routine_id="routine_id",
            routine_type="PROCEDURE",
            language="SQL",
            definition_body="CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);")
        ```

        ## Import

        Routine can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default {{project}}/{{dataset_id}}/{{routine_id}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/routine:Routine default {{dataset_id}}/{{routine_id}}
        ```

        :param str resource_name: The name of the resource.
        :param RoutineArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoutineArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arguments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutineArgumentArgs']]]]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 definition_body: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 determinism_level: Optional[pulumi.Input[str]] = None,
                 imported_libraries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 return_type: Optional[pulumi.Input[str]] = None,
                 routine_id: Optional[pulumi.Input[str]] = None,
                 routine_type: Optional[pulumi.Input[str]] = None,
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

            __props__['arguments'] = arguments
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            if definition_body is None and not opts.urn:
                raise TypeError("Missing required property 'definition_body'")
            __props__['definition_body'] = definition_body
            __props__['description'] = description
            __props__['determinism_level'] = determinism_level
            __props__['imported_libraries'] = imported_libraries
            __props__['language'] = language
            __props__['project'] = project
            __props__['return_type'] = return_type
            if routine_id is None and not opts.urn:
                raise TypeError("Missing required property 'routine_id'")
            __props__['routine_id'] = routine_id
            __props__['routine_type'] = routine_type
            __props__['creation_time'] = None
            __props__['last_modified_time'] = None
        super(Routine, __self__).__init__(
            'gcp:bigquery/routine:Routine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arguments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutineArgumentArgs']]]]] = None,
            creation_time: Optional[pulumi.Input[int]] = None,
            dataset_id: Optional[pulumi.Input[str]] = None,
            definition_body: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            determinism_level: Optional[pulumi.Input[str]] = None,
            imported_libraries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            language: Optional[pulumi.Input[str]] = None,
            last_modified_time: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            return_type: Optional[pulumi.Input[str]] = None,
            routine_id: Optional[pulumi.Input[str]] = None,
            routine_type: Optional[pulumi.Input[str]] = None) -> 'Routine':
        """
        Get an existing Routine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RoutineArgumentArgs']]]] arguments: Input/output argument of a function or a stored procedure.
               Structure is documented below.
        :param pulumi.Input[int] creation_time: The time when this routine was created, in milliseconds since the epoch.
        :param pulumi.Input[str] dataset_id: The ID of the dataset containing this routine
        :param pulumi.Input[str] definition_body: The body of the routine. For functions, this is the expression in the AS clause.
               If language=SQL, it is the substring inside (but excluding) the parentheses.
        :param pulumi.Input[str] description: The description of the routine if defined.
        :param pulumi.Input[str] determinism_level: The determinism level of the JavaScript UDF if defined.
               Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] imported_libraries: Optional. If language = "JAVASCRIPT", this field stores the path of the
               imported JAVASCRIPT libraries.
        :param pulumi.Input[str] language: The language of the routine.
               Possible values are `SQL` and `JAVASCRIPT`.
        :param pulumi.Input[int] last_modified_time: The time when this routine was modified, in milliseconds since the epoch.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] return_type: A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
               If absent, the return type is inferred from definitionBody at query time in each query
               that references this routine. If present, then the evaluated result will be cast to
               the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
               string, any changes to the string will create a diff, even if the JSON itself hasn't
               changed. If the API returns a different value for the same schema, e.g. it switche
               d the order of values or replaced STRUCT field type with RECORD field type, we currently
               cannot suppress the recurring diff this causes. As a workaround, we recommend using
               the schema as returned by the API.
        :param pulumi.Input[str] routine_id: The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        :param pulumi.Input[str] routine_type: The type of routine.
               Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arguments"] = arguments
        __props__["creation_time"] = creation_time
        __props__["dataset_id"] = dataset_id
        __props__["definition_body"] = definition_body
        __props__["description"] = description
        __props__["determinism_level"] = determinism_level
        __props__["imported_libraries"] = imported_libraries
        __props__["language"] = language
        __props__["last_modified_time"] = last_modified_time
        __props__["project"] = project
        __props__["return_type"] = return_type
        __props__["routine_id"] = routine_id
        __props__["routine_type"] = routine_type
        return Routine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arguments(self) -> pulumi.Output[Optional[Sequence['outputs.RoutineArgument']]]:
        """
        Input/output argument of a function or a stored procedure.
        Structure is documented below.
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[int]:
        """
        The time when this routine was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Output[str]:
        """
        The ID of the dataset containing this routine
        """
        return pulumi.get(self, "dataset_id")

    @property
    @pulumi.getter(name="definitionBody")
    def definition_body(self) -> pulumi.Output[str]:
        """
        The body of the routine. For functions, this is the expression in the AS clause.
        If language=SQL, it is the substring inside (but excluding) the parentheses.
        """
        return pulumi.get(self, "definition_body")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the routine if defined.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="determinismLevel")
    def determinism_level(self) -> pulumi.Output[Optional[str]]:
        """
        The determinism level of the JavaScript UDF if defined.
        Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
        """
        return pulumi.get(self, "determinism_level")

    @property
    @pulumi.getter(name="importedLibraries")
    def imported_libraries(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Optional. If language = "JAVASCRIPT", this field stores the path of the
        imported JAVASCRIPT libraries.
        """
        return pulumi.get(self, "imported_libraries")

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[Optional[str]]:
        """
        The language of the routine.
        Possible values are `SQL` and `JAVASCRIPT`.
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> pulumi.Output[int]:
        """
        The time when this routine was modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="returnType")
    def return_type(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
        If absent, the return type is inferred from definitionBody at query time in each query
        that references this routine. If present, then the evaluated result will be cast to
        the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
        string, any changes to the string will create a diff, even if the JSON itself hasn't
        changed. If the API returns a different value for the same schema, e.g. it switche
        d the order of values or replaced STRUCT field type with RECORD field type, we currently
        cannot suppress the recurring diff this causes. As a workaround, we recommend using
        the schema as returned by the API.
        """
        return pulumi.get(self, "return_type")

    @property
    @pulumi.getter(name="routineId")
    def routine_id(self) -> pulumi.Output[str]:
        """
        The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
        """
        return pulumi.get(self, "routine_id")

    @property
    @pulumi.getter(name="routineType")
    def routine_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of routine.
        Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
        """
        return pulumi.get(self, "routine_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

