// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A user-defined function or a stored procedure that belongs to a Dataset
//
// To get more information about Routine, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)
// * How-to Guides
//     * [Routines Intro](https://cloud.google.com/bigquery/docs/reference/rest/v2/routines)
//
// ## Example Usage
// ### Big Query Routine Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := bigquery.NewDataset(ctx, "test", &bigquery.DatasetArgs{
// 			DatasetId: pulumi.String("dataset_id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewRoutine(ctx, "sproc", &bigquery.RoutineArgs{
// 			DatasetId:      test.DatasetId,
// 			RoutineId:      pulumi.String("routine_id"),
// 			RoutineType:    pulumi.String("PROCEDURE"),
// 			Language:       pulumi.String("SQL"),
// 			DefinitionBody: pulumi.String("CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);"),
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
// Routine can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:bigquery/routine:Routine default projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/routine:Routine default {{project}}/{{dataset_id}}/{{routine_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/routine:Routine default {{dataset_id}}/{{routine_id}}
// ```
type Routine struct {
	pulumi.CustomResourceState

	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments RoutineArgumentArrayOutput `pulumi:"arguments"`
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime pulumi.IntOutput `pulumi:"creationTime"`
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringOutput `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined.
	// Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
	DeterminismLevel pulumi.StringPtrOutput `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayOutput `pulumi:"importedLibraries"`
	// The language of the routine.
	// Possible values are `SQL` and `JAVASCRIPT`.
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntOutput `pulumi:"lastModifiedTime"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType pulumi.StringPtrOutput `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId pulumi.StringOutput `pulumi:"routineId"`
	// The type of routine.
	// Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
	RoutineType pulumi.StringPtrOutput `pulumi:"routineType"`
}

// NewRoutine registers a new resource with the given unique name, arguments, and options.
func NewRoutine(ctx *pulumi.Context,
	name string, args *RoutineArgs, opts ...pulumi.ResourceOption) (*Routine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.DefinitionBody == nil {
		return nil, errors.New("invalid value for required argument 'DefinitionBody'")
	}
	if args.RoutineId == nil {
		return nil, errors.New("invalid value for required argument 'RoutineId'")
	}
	var resource Routine
	err := ctx.RegisterResource("gcp:bigquery/routine:Routine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutine gets an existing Routine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutineState, opts ...pulumi.ResourceOption) (*Routine, error) {
	var resource Routine
	err := ctx.ReadResource("gcp:bigquery/routine:Routine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Routine resources.
type routineState struct {
	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments []RoutineArgument `pulumi:"arguments"`
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime *int `pulumi:"creationTime"`
	// The ID of the dataset containing this routine
	DatasetId *string `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody *string `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description *string `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined.
	// Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
	DeterminismLevel *string `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries []string `pulumi:"importedLibraries"`
	// The language of the routine.
	// Possible values are `SQL` and `JAVASCRIPT`.
	Language *string `pulumi:"language"`
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime *int `pulumi:"lastModifiedTime"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType *string `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId *string `pulumi:"routineId"`
	// The type of routine.
	// Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
	RoutineType *string `pulumi:"routineType"`
}

type RoutineState struct {
	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments RoutineArgumentArrayInput
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime pulumi.IntPtrInput
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringPtrInput
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringPtrInput
	// The description of the routine if defined.
	Description pulumi.StringPtrInput
	// The determinism level of the JavaScript UDF if defined.
	// Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
	DeterminismLevel pulumi.StringPtrInput
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayInput
	// The language of the routine.
	// Possible values are `SQL` and `JAVASCRIPT`.
	Language pulumi.StringPtrInput
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType pulumi.StringPtrInput
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId pulumi.StringPtrInput
	// The type of routine.
	// Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
	RoutineType pulumi.StringPtrInput
}

func (RoutineState) ElementType() reflect.Type {
	return reflect.TypeOf((*routineState)(nil)).Elem()
}

type routineArgs struct {
	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments []RoutineArgument `pulumi:"arguments"`
	// The ID of the dataset containing this routine
	DatasetId string `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody string `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description *string `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined.
	// Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
	DeterminismLevel *string `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries []string `pulumi:"importedLibraries"`
	// The language of the routine.
	// Possible values are `SQL` and `JAVASCRIPT`.
	Language *string `pulumi:"language"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType *string `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId string `pulumi:"routineId"`
	// The type of routine.
	// Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
	RoutineType *string `pulumi:"routineType"`
}

// The set of arguments for constructing a Routine resource.
type RoutineArgs struct {
	// Input/output argument of a function or a stored procedure.
	// Structure is documented below.
	Arguments RoutineArgumentArrayInput
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringInput
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringInput
	// The description of the routine if defined.
	Description pulumi.StringPtrInput
	// The determinism level of the JavaScript UDF if defined.
	// Possible values are `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, and `NOT_DETERMINISTIC`.
	DeterminismLevel pulumi.StringPtrInput
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayInput
	// The language of the routine.
	// Possible values are `SQL` and `JAVASCRIPT`.
	Language pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	ReturnType pulumi.StringPtrInput
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId pulumi.StringInput
	// The type of routine.
	// Possible values are `SCALAR_FUNCTION` and `PROCEDURE`.
	RoutineType pulumi.StringPtrInput
}

func (RoutineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routineArgs)(nil)).Elem()
}

type RoutineInput interface {
	pulumi.Input

	ToRoutineOutput() RoutineOutput
	ToRoutineOutputWithContext(ctx context.Context) RoutineOutput
}

func (Routine) ElementType() reflect.Type {
	return reflect.TypeOf((*Routine)(nil)).Elem()
}

func (i Routine) ToRoutineOutput() RoutineOutput {
	return i.ToRoutineOutputWithContext(context.Background())
}

func (i Routine) ToRoutineOutputWithContext(ctx context.Context) RoutineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutineOutput)
}

type RoutineOutput struct {
	*pulumi.OutputState
}

func (RoutineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutineOutput)(nil)).Elem()
}

func (o RoutineOutput) ToRoutineOutput() RoutineOutput {
	return o
}

func (o RoutineOutput) ToRoutineOutputWithContext(ctx context.Context) RoutineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoutineOutput{})
}
