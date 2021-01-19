// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
// (https://www.dicomstandard.org/about/) standard for Healthcare information exchange
//
// To get more information about DicomStore, see:
//
// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
// * How-to Guides
//     * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)
//
// ## Example Usage
// ### Healthcare Dicom Store Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/healthcare"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		topic, err := pubsub.NewTopic(ctx, "topic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		dataset, err := healthcare.NewDataset(ctx, "dataset", &healthcare.DatasetArgs{
// 			Location: pulumi.String("us-central1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = healthcare.NewDicomStore(ctx, "_default", &healthcare.DicomStoreArgs{
// 			Dataset: dataset.ID(),
// 			NotificationConfig: &healthcare.DicomStoreNotificationConfigArgs{
// 				PubsubTopic: topic.ID(),
// 			},
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("labelvalue1"),
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
// DicomStore can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/dicomStores/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/{{name}}
// ```
type DicomStore struct {
	pulumi.CustomResourceState

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the DicomStore.
	// ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringOutput `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig DicomStoreNotificationConfigPtrOutput `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewDicomStore registers a new resource with the given unique name, arguments, and options.
func NewDicomStore(ctx *pulumi.Context,
	name string, args *DicomStoreArgs, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	var resource DicomStore
	err := ctx.RegisterResource("gcp:healthcare/dicomStore:DicomStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDicomStore gets an existing DicomStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDicomStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DicomStoreState, opts ...pulumi.ResourceOption) (*DicomStore, error) {
	var resource DicomStore
	err := ctx.ReadResource("gcp:healthcare/dicomStore:DicomStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DicomStore resources.
type dicomStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset *string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the DicomStore.
	// ** Changing this property may recreate the Dicom store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig *DicomStoreNotificationConfig `pulumi:"notificationConfig"`
	// The fully qualified name of this dataset
	SelfLink *string `pulumi:"selfLink"`
}

type DicomStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize DICOM stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the DicomStore.
	// ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	// Structure is documented below.
	NotificationConfig DicomStoreNotificationConfigPtrInput
	// The fully qualified name of this dataset
	SelfLink pulumi.StringPtrInput
}

func (DicomStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreState)(nil)).Elem()
}

type dicomStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset string `pulumi:"dataset"`
	// User-supplied key-value pairs used to organize DICOM stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the DicomStore.
	// ** Changing this property may recreate the Dicom store (removing all data) **
	Name *string `pulumi:"name"`
	// A nested object resource
	// Structure is documented below.
	NotificationConfig *DicomStoreNotificationConfig `pulumi:"notificationConfig"`
}

// The set of arguments for constructing a DicomStore resource.
type DicomStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset pulumi.StringInput
	// User-supplied key-value pairs used to organize DICOM stores.
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given store.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The resource name for the DicomStore.
	// ** Changing this property may recreate the Dicom store (removing all data) **
	Name pulumi.StringPtrInput
	// A nested object resource
	// Structure is documented below.
	NotificationConfig DicomStoreNotificationConfigPtrInput
}

func (DicomStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dicomStoreArgs)(nil)).Elem()
}

type DicomStoreInput interface {
	pulumi.Input

	ToDicomStoreOutput() DicomStoreOutput
	ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput
}

func (*DicomStore) ElementType() reflect.Type {
	return reflect.TypeOf((*DicomStore)(nil))
}

func (i *DicomStore) ToDicomStoreOutput() DicomStoreOutput {
	return i.ToDicomStoreOutputWithContext(context.Background())
}

func (i *DicomStore) ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DicomStoreOutput)
}

type DicomStoreOutput struct {
	*pulumi.OutputState
}

func (DicomStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DicomStore)(nil))
}

func (o DicomStoreOutput) ToDicomStoreOutput() DicomStoreOutput {
	return o
}

func (o DicomStoreOutput) ToDicomStoreOutputWithContext(ctx context.Context) DicomStoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DicomStoreOutput{})
}
