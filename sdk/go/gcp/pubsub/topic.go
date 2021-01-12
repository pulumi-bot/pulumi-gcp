// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A named resource to which messages are sent by publishers.
//
// To get more information about Topic, see:
//
// * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
// * How-to Guides
//     * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)
//
// > **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
// by using the `projects.ServiceIdentity` resource.
//
// ## Example Usage
// ### Pubsub Topic Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopic(ctx, "example", &pubsub.TopicArgs{
// 			Labels: pulumi.StringMap{
// 				"foo": pulumi.String("bar"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Pubsub Topic Cmek
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/kms"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		keyRing, err := kms.NewKeyRing(ctx, "keyRing", &kms.KeyRingArgs{
// 			Location: pulumi.String("global"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		cryptoKey, err := kms.NewCryptoKey(ctx, "cryptoKey", &kms.CryptoKeyArgs{
// 			KeyRing: keyRing.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewTopic(ctx, "example", &pubsub.TopicArgs{
// 			KmsKeyName: cryptoKey.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Pubsub Topic Geo Restricted
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pubsub.NewTopic(ctx, "example", &pubsub.TopicArgs{
// 			MessageStoragePolicy: &pubsub.TopicMessageStoragePolicyArgs{
// 				AllowedPersistenceRegions: pulumi.StringArray{
// 					pulumi.String("europe-west3"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Topic can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:pubsub/topic:Topic default projects/{{project}}/topics/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/topic:Topic default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:pubsub/topic:Topic default {{name}}
// ```
type Topic struct {
	pulumi.CustomResourceState

	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyOutput `pulumi:"messageStoragePolicy"`
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
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]string `pulumi:"labels"`
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy *TopicMessageStoragePolicy `pulumi:"messageStoragePolicy"`
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type TopicState struct {
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapInput
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyPtrInput
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
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// A set of key/value label pairs to assign to this Topic.
	Labels map[string]string `pulumi:"labels"`
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy *TopicMessageStoragePolicy `pulumi:"messageStoragePolicy"`
	// Name of the topic.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
	// `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
	// The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
	KmsKeyName pulumi.StringPtrInput
	// A set of key/value label pairs to assign to this Topic.
	Labels pulumi.StringMapInput
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// Structure is documented below.
	MessageStoragePolicy TopicMessageStoragePolicyPtrInput
	// Name of the topic.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *Topic) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicPtrOutput)
}

type TopicPtrInput interface {
	pulumi.Input

	ToTopicPtrOutput() TopicPtrOutput
	ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput
}

type topicPtrType TopicArgs

func (*topicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (i *topicPtrType) ToTopicPtrOutput() TopicPtrOutput {
	return i.ToTopicPtrOutputWithContext(context.Background())
}

func (i *topicPtrType) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput).ToTopicPtrOutput()
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//          TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Topic)(nil))
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//          TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Topic)(nil))
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct {
	*pulumi.OutputState
}

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Topic)(nil))
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o.ToTopicPtrOutputWithContext(context.Background())
}

func (o TopicOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o.ApplyT(func(v Topic) *Topic {
		return &v
	}).(TopicPtrOutput)
}

type TopicPtrOutput struct {
	*pulumi.OutputState
}

func (TopicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil))
}

func (o TopicPtrOutput) ToTopicPtrOutput() TopicPtrOutput {
	return o
}

func (o TopicPtrOutput) ToTopicPtrOutputWithContext(ctx context.Context) TopicPtrOutput {
	return o
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Topic)(nil))
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Topic {
		return vs[0].([]Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Topic)(nil))
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Topic {
		return vs[0].(map[string]Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicPtrOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
