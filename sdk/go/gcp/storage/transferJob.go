// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Transfer Job in Google Cloud Storage Transfer.
//
// To get more information about Google Cloud Storage Transfer, see:
//
// * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs#TransferJob)
// * How-to Guides
//     * [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
//
// ## Example Usage
//
// Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Project
// 		_default, err := storage.GetTransferProjectServieAccount(ctx, &storage.GetTransferProjectServieAccountArgs{
// 			Project: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucket(ctx, "s3-backup-bucketBucket", &storage.BucketArgs{
// 			StorageClass: pulumi.String("NEARLINE"),
// 			Project:      pulumi.Any(_var.Project),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucketIAMMember(ctx, "s3-backup-bucketBucketIAMMember", &storage.BucketIAMMemberArgs{
// 			Bucket: s3_backup_bucketBucket.Name,
// 			Role:   pulumi.String("roles/storage.admin"),
// 			Member: pulumi.String(fmt.Sprintf("%v%v", "serviceAccount:", _default.Email)),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			s3_backup_bucketBucket,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewTransferJob(ctx, "s3-bucket-nightly-backup", &storage.TransferJobArgs{
// 			Description: pulumi.String("Nightly backup of S3 bucket"),
// 			Project:     pulumi.Any(_var.Project),
// 			TransferSpec: &storage.TransferJobTransferSpecArgs{
// 				ObjectConditions: &storage.TransferJobTransferSpecObjectConditionsArgs{
// 					MaxTimeElapsedSinceLastModification: pulumi.String("600s"),
// 					ExcludePrefixes: pulumi.StringArray{
// 						pulumi.String("requests.gz"),
// 					},
// 				},
// 				TransferOptions: &storage.TransferJobTransferSpecTransferOptionsArgs{
// 					DeleteObjectsUniqueInSink: pulumi.Bool(false),
// 				},
// 				AwsS3DataSource: &storage.TransferJobTransferSpecAwsS3DataSourceArgs{
// 					BucketName: pulumi.Any(_var.Aws_s3_bucket),
// 					AwsAccessKey: &storage.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs{
// 						AccessKeyId:     pulumi.Any(_var.Aws_access_key),
// 						SecretAccessKey: pulumi.Any(_var.Aws_secret_key),
// 					},
// 				},
// 				GcsDataSink: &storage.TransferJobTransferSpecGcsDataSinkArgs{
// 					BucketName: s3_backup_bucketBucket.Name,
// 				},
// 			},
// 			Schedule: &storage.TransferJobScheduleArgs{
// 				ScheduleStartDate: &storage.TransferJobScheduleScheduleStartDateArgs{
// 					Year:  pulumi.Int(2018),
// 					Month: pulumi.Int(10),
// 					Day:   pulumi.Int(1),
// 				},
// 				ScheduleEndDate: &storage.TransferJobScheduleScheduleEndDateArgs{
// 					Year:  pulumi.Int(2019),
// 					Month: pulumi.Int(1),
// 					Day:   pulumi.Int(15),
// 				},
// 				StartTimeOfDay: &storage.TransferJobScheduleStartTimeOfDayArgs{
// 					Hours:   pulumi.Int(23),
// 					Minutes: pulumi.Int(30),
// 					Seconds: pulumi.Int(0),
// 					Nanos:   pulumi.Int(0),
// 				},
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			s3_backup_bucketBucketIAMMember,
// 		}))
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
// Storage buckets can be imported using the Transfer Job's `project` and `name` without the `transferJob/` prefix, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/transferJob:TransferJob nightly-backup-transfer-job my-project-1asd32/8422144862922355674
// ```
type TransferJob struct {
	pulumi.CustomResourceState

	// When the Transfer Job was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringOutput `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description pulumi.StringOutput `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringOutput `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobScheduleOutput `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecOutput `pulumi:"transferSpec"`
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.TransferSpec == nil {
		return nil, errors.New("invalid value for required argument 'TransferSpec'")
	}
	var resource TransferJob
	err := ctx.RegisterResource("gcp:storage/transferJob:TransferJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferJobState, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	var resource TransferJob
	err := ctx.ReadResource("gcp:storage/transferJob:TransferJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferJob resources.
type transferJobState struct {
	// When the Transfer Job was created.
	CreationTime *string `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime *string `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description *string `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime *string `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule *TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec *TransferJobTransferSpec `pulumi:"transferSpec"`
}

type TransferJobState struct {
	// When the Transfer Job was created.
	CreationTime pulumi.StringPtrInput
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringPtrInput
	// Unique description to identify the Transfer Job.
	Description pulumi.StringPtrInput
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringPtrInput
	// The name of the Transfer Job.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobSchedulePtrInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecPtrInput
}

func (TransferJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobState)(nil)).Elem()
}

type transferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description string `pulumi:"description"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpec `pulumi:"transferSpec"`
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobScheduleInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecInput
}

func (TransferJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobArgs)(nil)).Elem()
}

type TransferJobInput interface {
	pulumi.Input

	ToTransferJobOutput() TransferJobOutput
	ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput
}

func (*TransferJob) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJob)(nil))
}

func (i *TransferJob) ToTransferJobOutput() TransferJobOutput {
	return i.ToTransferJobOutputWithContext(context.Background())
}

func (i *TransferJob) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobOutput)
}

type TransferJobOutput struct {
	*pulumi.OutputState
}

func (TransferJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransferJob)(nil))
}

func (o TransferJobOutput) ToTransferJobOutput() TransferJobOutput {
	return o
}

func (o TransferJobOutput) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransferJobOutput{})
}
