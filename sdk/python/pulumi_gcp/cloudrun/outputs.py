# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'DomainMappingMetadata',
    'DomainMappingSpec',
    'DomainMappingStatus',
    'DomainMappingStatusCondition',
    'DomainMappingStatusResourceRecord',
    'IamBindingCondition',
    'IamMemberCondition',
    'ServiceMetadata',
    'ServiceStatus',
    'ServiceStatusCondition',
    'ServiceTemplate',
    'ServiceTemplateMetadata',
    'ServiceTemplateSpec',
    'ServiceTemplateSpecContainer',
    'ServiceTemplateSpecContainerEnv',
    'ServiceTemplateSpecContainerEnvFrom',
    'ServiceTemplateSpecContainerEnvFromConfigMapRef',
    'ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference',
    'ServiceTemplateSpecContainerEnvFromSecretRef',
    'ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference',
    'ServiceTemplateSpecContainerPort',
    'ServiceTemplateSpecContainerResources',
    'ServiceTraffic',
]

@pulumi.output_type
class DomainMappingMetadata(dict):
    @property
    @pulumi.getter
    def annotations(self) -> Optional[Mapping[str, str]]:
        """
        Annotations is a key value map stored with a resource that
        may be set by external tools to store and retrieve arbitrary metadata. More
        info: http://kubernetes.io/docs/user-guide/annotations
        """
        ...

    @property
    @pulumi.getter
    def generation(self) -> Optional[float]:
        """
        -
        A sequence number representing a specific generation of the desired state.
        """
        ...

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Map of string keys and values that can be used to organize and categorize
        (scope and select) objects. May match selectors of replication controllers
        and routes.
        More info: http://kubernetes.io/docs/user-guide/labels
        """
        ...

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        In Cloud Run the namespace must be equal to either the
        project ID or project number.
        """
        ...

    @property
    @pulumi.getter(name="resourceVersion")
    def resource_version(self) -> Optional[str]:
        """
        -
        An opaque value that represents the internal version of this object that
        can be used by clients to determine when objects have changed. May be used
        for optimistic concurrency, change detection, and the watch operation on a
        resource or set of resources. They may only be valid for a
        particular resource or set of resources.
        More info:
        https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
        """
        ...

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[str]:
        """
        -
        SelfLink is a URL representing this object.
        """
        ...

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        -
        UID is a unique id generated by the server on successful creation of a resource and is not
        allowed to change on PUT operations.
        More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingSpec(dict):
    @property
    @pulumi.getter(name="certificateMode")
    def certificate_mode(self) -> Optional[str]:
        """
        The mode of the certificate.
        """
        ...

    @property
    @pulumi.getter(name="forceOverride")
    def force_override(self) -> Optional[bool]:
        """
        If set, the mapping will override any mapping set before this spec was set.
        It is recommended that the user leaves this empty to receive an error
        warning about a potential conflict and only set it once the respective UI
        has given such a warning.
        """
        ...

    @property
    @pulumi.getter(name="routeName")
    def route_name(self) -> str:
        """
        The name of the Cloud Run Service that this DomainMapping applies to.
        The route must exist.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatus(dict):
    @property
    @pulumi.getter
    def conditions(self) -> Optional[List['outputs.DomainMappingStatusCondition']]:
        ...

    @property
    @pulumi.getter(name="mappedRouteName")
    def mapped_route_name(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[float]:
        ...

    @property
    @pulumi.getter(name="resourceRecords")
    def resource_records(self) -> Optional[List['outputs.DomainMappingStatusResourceRecord']]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatusCondition(dict):
    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def reason(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DomainMappingStatusResourceRecord(dict):
    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name should be a verified domain
        """
        ...

    @property
    @pulumi.getter
    def rrdata(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamBindingCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IamMemberCondition(dict):
    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def expression(self) -> str:
        ...

    @property
    @pulumi.getter
    def title(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceMetadata(dict):
    @property
    @pulumi.getter
    def annotations(self) -> Optional[Mapping[str, str]]:
        """
        Annotations is a key value map stored with a resource that
        may be set by external tools to store and retrieve arbitrary metadata. More
        info: http://kubernetes.io/docs/user-guide/annotations
        """
        ...

    @property
    @pulumi.getter
    def generation(self) -> Optional[float]:
        """
        -
        A sequence number representing a specific generation of the desired state.
        """
        ...

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Map of string keys and values that can be used to organize and categorize
        (scope and select) objects. May match selectors of replication controllers
        and routes.
        More info: http://kubernetes.io/docs/user-guide/labels
        """
        ...

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        In Cloud Run the namespace must be equal to either the
        project ID or project number.
        """
        ...

    @property
    @pulumi.getter(name="resourceVersion")
    def resource_version(self) -> Optional[str]:
        """
        -
        An opaque value that represents the internal version of this object that
        can be used by clients to determine when objects have changed. May be used
        for optimistic concurrency, change detection, and the watch operation on a
        resource or set of resources. They may only be valid for a
        particular resource or set of resources.
        More info:
        https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
        """
        ...

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[str]:
        """
        -
        SelfLink is a URL representing this object.
        """
        ...

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        -
        UID is a unique id generated by the server on successful creation of a resource and is not
        allowed to change on PUT operations.
        More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceStatus(dict):
    @property
    @pulumi.getter
    def conditions(self) -> Optional[List['outputs.ServiceStatusCondition']]:
        ...

    @property
    @pulumi.getter(name="latestCreatedRevisionName")
    def latest_created_revision_name(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="latestReadyRevisionName")
    def latest_ready_revision_name(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> Optional[float]:
        ...

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceStatusCondition(dict):
    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def reason(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplate(dict):
    @property
    @pulumi.getter
    def metadata(self) -> Optional['outputs.ServiceTemplateMetadata']:
        """
        Metadata associated with this Service, including name, namespace, labels,
        and annotations.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def spec(self) -> Optional['outputs.ServiceTemplateSpec']:
        """
        RevisionSpec holds the desired state of the Revision (from the client).  Structure is documented below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateMetadata(dict):
    @property
    @pulumi.getter
    def annotations(self) -> Optional[Mapping[str, str]]:
        """
        Annotations is a key value map stored with a resource that
        may be set by external tools to store and retrieve arbitrary metadata. More
        info: http://kubernetes.io/docs/user-guide/annotations
        """
        ...

    @property
    @pulumi.getter
    def generation(self) -> Optional[float]:
        """
        -
        A sequence number representing a specific generation of the desired state.
        """
        ...

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        Map of string keys and values that can be used to organize and categorize
        (scope and select) objects. May match selectors of replication controllers
        and routes.
        More info: http://kubernetes.io/docs/user-guide/labels
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the port.
        """
        ...

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        In Cloud Run the namespace must be equal to either the
        project ID or project number.
        """
        ...

    @property
    @pulumi.getter(name="resourceVersion")
    def resource_version(self) -> Optional[str]:
        """
        -
        An opaque value that represents the internal version of this object that
        can be used by clients to determine when objects have changed. May be used
        for optimistic concurrency, change detection, and the watch operation on a
        resource or set of resources. They may only be valid for a
        particular resource or set of resources.
        More info:
        https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
        """
        ...

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[str]:
        """
        -
        SelfLink is a URL representing this object.
        """
        ...

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        -
        UID is a unique id generated by the server on successful creation of a resource and is not
        allowed to change on PUT operations.
        More info: http://kubernetes.io/docs/user-guide/identifiers#uids
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpec(dict):
    @property
    @pulumi.getter(name="containerConcurrency")
    def container_concurrency(self) -> Optional[float]:
        """
        ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
        requests per container of the Revision. Values are:
        - `0` thread-safe, the system should manage the max concurrency. This is
        the default value.
        - `1` not-thread-safe. Single concurrency
        - `2-N` thread-safe, max concurrency of N
        """
        ...

    @property
    @pulumi.getter
    def containers(self) -> Optional[List['outputs.ServiceTemplateSpecContainer']]:
        """
        Container defines the unit of execution for this Revision.
        In the context of a Revision, we disallow a number of the fields of
        this Container, including: name, ports, and volumeMounts.
        The runtime contract is documented here:
        https://github.com/knative/serving/blob/master/docs/runtime-contract.md  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter(name="serviceAccountName")
    def service_account_name(self) -> Optional[str]:
        """
        Email address of the IAM service account associated with the revision of the
        service. The service account represents the identity of the running revision,
        and determines what permissions the revision has. If not provided, the revision
        will use the project's default service account.
        """
        ...

    @property
    @pulumi.getter(name="servingState")
    def serving_state(self) -> Optional[str]:
        """
        -
        ServingState holds a value describing the state the resources
        are in for this Revision.
        It is expected
        that the system will manipulate this based on routability and load.
        """
        ...

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> Optional[float]:
        """
        TimeoutSeconds holds the max duration the instance is allowed for responding to a request.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainer(dict):
    @property
    @pulumi.getter
    def args(self) -> Optional[List[str]]:
        """
        Arguments to the entrypoint.
        The docker image's CMD is used if this is not provided.
        Variable references $(VAR_NAME) are expanded using the container's
        environment. If a variable cannot be resolved, the reference in the input
        string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
        double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
        regardless of whether the variable exists or not.
        More info:
        https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
        """
        ...

    @property
    @pulumi.getter
    def commands(self) -> Optional[List[str]]:
        """
        Entrypoint array. Not executed within a shell.
        The docker image's ENTRYPOINT is used if this is not provided.
        Variable references $(VAR_NAME) are expanded using the container's
        environment. If a variable cannot be resolved, the reference in the input
        string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
        double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
        regardless of whether the variable exists or not.
        More info:
        https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
        """
        ...

    @property
    @pulumi.getter(name="envFroms")
    def env_froms(self) -> Optional[List['outputs.ServiceTemplateSpecContainerEnvFrom']]:
        """
        -
        (Optional, Deprecated)
        List of sources to populate environment variables in the container.
        All invalid keys will be reported as an event when the container is starting.
        When a key exists in multiple sources, the value associated with the last source will
        take precedence. Values defined by an Env with a duplicate key will take
        precedence.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def envs(self) -> Optional[List['outputs.ServiceTemplateSpecContainerEnv']]:
        """
        List of environment variables to set in the container.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def image(self) -> str:
        """
        Docker image name. This is most often a reference to a container located
        in the container registry, such as gcr.io/cloudrun/hello
        More info: https://kubernetes.io/docs/concepts/containers/images
        """
        ...

    @property
    @pulumi.getter
    def ports(self) -> Optional[List['outputs.ServiceTemplateSpecContainerPort']]:
        """
        List of open ports in the container.
        More Info:
        https://cloud.google.com/run/docs/reference/rest/v1/RevisionSpec#ContainerPort  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def resources(self) -> Optional['outputs.ServiceTemplateSpecContainerResources']:
        """
        Compute Resources required by this container. Used to set values such as max memory
        More info:
        https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter(name="workingDir")
    def working_dir(self) -> Optional[str]:
        """
        -
        (Optional, Deprecated)
        Container's working directory.
        If not specified, the container runtime's default will be used, which
        might be configured in the container image.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnv(dict):
    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the port.
        """
        ...

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        Variable references $(VAR_NAME) are expanded
        using the previous defined environment variables in the container and
        any route environment variables. If a variable cannot be resolved,
        the reference in the input string will be unchanged. The $(VAR_NAME)
        syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
        references will never be expanded, regardless of whether the variable
        exists or not.
        Defaults to "".
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFrom(dict):
    @property
    @pulumi.getter(name="configMapRef")
    def config_map_ref(self) -> Optional['outputs.ServiceTemplateSpecContainerEnvFromConfigMapRef']:
        """
        The ConfigMap to select from.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def prefix(self) -> Optional[str]:
        """
        An optional identifier to prepend to each key in the ConfigMap.
        """
        ...

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> Optional['outputs.ServiceTemplateSpecContainerEnvFromSecretRef']:
        """
        The Secret to select from.  Structure is documented below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromConfigMapRef(dict):
    @property
    @pulumi.getter(name="localObjectReference")
    def local_object_reference(self) -> Optional['outputs.ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference']:
        """
        The Secret to select from.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def optional(self) -> Optional[bool]:
        """
        Specify whether the Secret must be defined
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReference(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the port.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromSecretRef(dict):
    @property
    @pulumi.getter(name="localObjectReference")
    def local_object_reference(self) -> Optional['outputs.ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference']:
        """
        The Secret to select from.  Structure is documented below.
        """
        ...

    @property
    @pulumi.getter
    def optional(self) -> Optional[bool]:
        """
        Specify whether the Secret must be defined
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerEnvFromSecretRefLocalObjectReference(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the port.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerPort(dict):
    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> float:
        """
        Port number.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the port.
        """
        ...

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        Protocol used on port. Defaults to TCP.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTemplateSpecContainerResources(dict):
    @property
    @pulumi.getter
    def limits(self) -> Optional[Mapping[str, str]]:
        """
        Limits describes the maximum amount of compute resources allowed.
        The values of the map is string form of the 'quantity' k8s type:
        https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        """
        ...

    @property
    @pulumi.getter
    def requests(self) -> Optional[Mapping[str, str]]:
        """
        Requests describes the minimum amount of compute resources required.
        If Requests is omitted for a container, it defaults to Limits if that is
        explicitly specified, otherwise to an implementation-defined value.
        The values of the map is string form of the 'quantity' k8s type:
        https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceTraffic(dict):
    @property
    @pulumi.getter(name="latestRevision")
    def latest_revision(self) -> Optional[bool]:
        """
        LatestRevision may be optionally provided to indicate that the latest ready
        Revision of the Configuration should be used for this traffic target. When
        provided LatestRevision must be true if RevisionName is empty; it must be
        false when RevisionName is non-empty.
        """
        ...

    @property
    @pulumi.getter
    def percent(self) -> float:
        """
        Percent specifies percent of the traffic to this Revision or Configuration.
        """
        ...

    @property
    @pulumi.getter(name="revisionName")
    def revision_name(self) -> Optional[str]:
        """
        RevisionName of a specific revision to which to send this portion of traffic.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


