// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cloudfunctions

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information about a Google Cloud Function. For more information see
// the [official documentation](https://cloud.google.com/functions/docs/)
// and [API](https://cloud.google.com/functions/docs/apis).
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_cloudfunctions_function.html.markdown.
func LookupFunction(ctx *pulumi.Context, args *LookupFunctionArgs, opts ...pulumi.InvokeOption) (*LookupFunctionResult, error) {
	var rv LookupFunctionResult
	err := ctx.Invoke("gcp:cloudfunctions/getFunction:getFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunction.
type LookupFunctionArgs struct {
	// The name of a Cloud Function.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the resource belongs. If it
	// is not provided, the provider region is used.
	Region *string `pulumi:"region"`
}


// A collection of values returned by getFunction.
type LookupFunctionResult struct {
	// Available memory (in MB) to the function.
	AvailableMemoryMb int `pulumi:"availableMemoryMb"`
	// Description of the function.
	Description string `pulumi:"description"`
	// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
	EntryPoint string `pulumi:"entryPoint"`
	EnvironmentVariables map[string]interface{} `pulumi:"environmentVariables"`
	// A source that fires events in response to a condition in another service. Structure is documented below.
	EventTriggers []GetFunctionEventTrigger `pulumi:"eventTriggers"`
	// If function is triggered by HTTP, trigger URL is set here.
	HttpsTriggerUrl string `pulumi:"httpsTriggerUrl"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A map of labels applied to this function.
	Labels map[string]interface{} `pulumi:"labels"`
	MaxInstances int `pulumi:"maxInstances"`
	// The name of the Cloud Function.
	Name string `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region *string `pulumi:"region"`
	// The runtime in which the function is running.
	Runtime string `pulumi:"runtime"`
	// The service account email to be assumed by the cloud function.
	ServiceAccountEmail string `pulumi:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	SourceArchiveBucket string `pulumi:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	SourceArchiveObject string `pulumi:"sourceArchiveObject"`
	SourceRepositories []GetFunctionSourceRepository `pulumi:"sourceRepositories"`
	// Function execution timeout (in seconds).
	Timeout int `pulumi:"timeout"`
	// If function is triggered by HTTP, this boolean is set.
	TriggerHttp bool `pulumi:"triggerHttp"`
	VpcConnector string `pulumi:"vpcConnector"`
}

