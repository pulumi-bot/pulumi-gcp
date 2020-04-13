// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A named resource to which messages are sent by publishers.
//
//
// To get more information about Topic, see:
//
// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
// * How-to Guides
//     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
type Topic struct {
	pulumi.CustomResourceState

	// -
	// (Optional)
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// -
	// (Optional)
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// -
	// (Optional)
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.  Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyOutput `pulumi:"messageStoragePolicy"`
	// -
	// (Required)
	// Name of the topic.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		args = &TopicArgs{}
	}
	var resource Topic
	err := ctx.RegisterResource("gcp:pubsub/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("gcp:pubsub/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// -
	// (Optional)
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// -
	// (Optional)
	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]string `pulumi:"labels"`
	// -
	// (Optional)
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.  Structure is documented below.
	MessageStoragePolicy *TopicMessageStoragePolicy `pulumi:"messageStoragePolicy"`
	// -
	// (Required)
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type TopicState struct {
	// -
	// (Optional)
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrInput
	// -
	// (Optional)
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapInput
	// -
	// (Optional)
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.  Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyPtrInput
	// -
	// (Required)
	// Name of the topic.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// -
	// (Optional)
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// -
	// (Optional)
	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]string `pulumi:"labels"`
	// -
	// (Optional)
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.  Structure is documented below.
	MessageStoragePolicy *TopicMessageStoragePolicy `pulumi:"messageStoragePolicy"`
	// -
	// (Required)
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// -
	// (Optional)
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrInput
	// -
	// (Optional)
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapInput
	// -
	// (Optional)
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.  Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyPtrInput
	// -
	// (Required)
	// Name of the topic.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}
