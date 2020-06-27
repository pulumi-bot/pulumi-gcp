// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
// Once a BigQuery job is created, it cannot be changed or deleted.
//
// ## Example Usage
// ### Bigquery Job Query
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_query_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DatasetId: bar.DatasetId,
// 			TableId:   pulumi.String("job_query_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_query"),
// 			Labels: pulumi.Map{
// 				"example-label": pulumi.String("example-value"),
// 			},
// 			Query: &bigquery.JobQueryArgs{
// 				Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
// 				Destination_table: pulumi.Map{
// 					"projectId": foo.Project,
// 					"datasetId": foo.DatasetId,
// 					"tableId":   foo.TableId,
// 				},
// 				AllowLargeResults: pulumi.Bool(true),
// 				FlattenResults:    pulumi.Bool(true),
// 				Script_options: pulumi.Map{
// 					"keyResultStatement": pulumi.String("LAST"),
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
// ### Bigquery Job Query Table Reference
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_query_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DatasetId: bar.DatasetId,
// 			TableId:   pulumi.String("job_query_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_query"),
// 			Labels: pulumi.Map{
// 				"example-label": pulumi.String("example-value"),
// 			},
// 			Query: &bigquery.JobQueryArgs{
// 				Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
// 				Destination_table: pulumi.Map{
// 					"tableId": foo.ID(),
// 				},
// 				Default_dataset: pulumi.Map{
// 					"datasetId": bar.ID(),
// 				},
// 				AllowLargeResults: pulumi.Bool(true),
// 				FlattenResults:    pulumi.Bool(true),
// 				Script_options: pulumi.Map{
// 					"keyResultStatement": pulumi.String("LAST"),
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
// ### Bigquery Job Load
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_load_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DatasetId: bar.DatasetId,
// 			TableId:   pulumi.String("job_load_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_load"),
// 			Labels: pulumi.Map{
// 				"my_job": pulumi.String("load"),
// 			},
// 			Load: &bigquery.JobLoadArgs{
// 				SourceUris: pulumi.StringArray{
// 					pulumi.String("gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv"),
// 				},
// 				Destination_table: pulumi.Map{
// 					"projectId": foo.Project,
// 					"datasetId": foo.DatasetId,
// 					"tableId":   foo.TableId,
// 				},
// 				SkipLeadingRows: pulumi.Int(1),
// 				SchemaUpdateOptions: pulumi.StringArray{
// 					pulumi.String("ALLOW_FIELD_RELAXATION"),
// 					pulumi.String("ALLOW_FIELD_ADDITION"),
// 				},
// 				WriteDisposition: pulumi.String("WRITE_APPEND"),
// 				Autodetect:       pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bigquery Job Extract
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = bigquery.NewDataset(ctx, "source_oneDataset", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_extract_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewTable(ctx, "source_oneTable", &bigquery.TableArgs{
// 			DatasetId: source_oneDataset.DatasetId,
// 			TableId:   pulumi.String("job_extract_table"),
// 			Schema:    pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "[\n", "  {\n", "    \"name\": \"name\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\"\n", "  },\n", "  {\n", "    \"name\": \"post_abbr\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\"\n", "  },\n", "  {\n", "    \"name\": \"date\",\n", "    \"type\": \"DATE\",\n", "    \"mode\": \"NULLABLE\"\n", "  }\n", "]\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		dest, err := storage.NewBucket(ctx, "dest", &storage.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_extract"),
// 			Extract: &bigquery.JobExtractArgs{
// 				DestinationUris: pulumi.StringArray{
// 					dest.Url.ApplyT(func(url string) (string, error) {
// 						return fmt.Sprintf("%v%v", url, "/extract"), nil
// 					}).(pulumi.StringOutput),
// 				},
// 				Source_table: pulumi.Map{
// 					"projectId": source_oneTable.Project,
// 					"datasetId": source_oneTable.DatasetId,
// 					"tableId":   source_oneTable.TableId,
// 				},
// 				DestinationFormat: pulumi.String("NEWLINE_DELIMITED_JSON"),
// 				Compression:       pulumi.String("GZIP"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Job struct {
	pulumi.CustomResourceState

	// Copies a table.  Structure is documented below.
	Copy JobCopyPtrOutput `pulumi:"copy"`
	// Configures an extract job.  Structure is documented below.
	Extract JobExtractPtrOutput `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrOutput `pulumi:"jobTimeoutMs"`
	// The type of the job.
	JobType pulumi.StringOutput `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Configures a load job.  Structure is documented below.
	Load JobLoadPtrOutput `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Configures a query job.  Structure is documented below.
	Query JobQueryPtrOutput `pulumi:"query"`
	// Email address of the user who ran the job.
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.JobId == nil {
		return nil, errors.New("missing required argument 'JobId'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("gcp:bigquery/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("gcp:bigquery/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Copies a table.  Structure is documented below.
	Copy *JobCopy `pulumi:"copy"`
	// Configures an extract job.  Structure is documented below.
	Extract *JobExtract `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId *string `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs *string `pulumi:"jobTimeoutMs"`
	// The type of the job.
	JobType *string `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.  Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Configures a query job.  Structure is documented below.
	Query *JobQuery `pulumi:"query"`
	// Email address of the user who ran the job.
	UserEmail *string `pulumi:"userEmail"`
}

type JobState struct {
	// Copies a table.  Structure is documented below.
	Copy JobCopyPtrInput
	// Configures an extract job.  Structure is documented below.
	Extract JobExtractPtrInput
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringPtrInput
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrInput
	// The type of the job.
	JobType pulumi.StringPtrInput
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapInput
	// Configures a load job.  Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Configures a query job.  Structure is documented below.
	Query JobQueryPtrInput
	// Email address of the user who ran the job.
	UserEmail pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Copies a table.  Structure is documented below.
	Copy *JobCopy `pulumi:"copy"`
	// Configures an extract job.  Structure is documented below.
	Extract *JobExtract `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId string `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs *string `pulumi:"jobTimeoutMs"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.  Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Configures a query job.  Structure is documented below.
	Query *JobQuery `pulumi:"query"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Copies a table.  Structure is documented below.
	Copy JobCopyPtrInput
	// Configures an extract job.  Structure is documented below.
	Extract JobExtractPtrInput
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringInput
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrInput
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapInput
	// Configures a load job.  Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Configures a query job.  Structure is documented below.
	Query JobQueryPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}
