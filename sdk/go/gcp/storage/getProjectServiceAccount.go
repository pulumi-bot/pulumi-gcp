// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the email address of a project's unique Google Cloud Storage service account.
//
// Each Google Cloud project has a unique service account for use with Google Cloud Storage. Only this
// special service account can be used to set up `storage.Notification` resources.
//
// For more information see
// [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v3/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		gcsAccount, err := storage.LookupProjectServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pubsub.NewTopicIAMBinding(ctx, "binding", &pubsub.TopicIAMBindingArgs{
// 			Topic: pulumi.String(google_pubsub_topic.Topic.Name),
// 			Role:  pulumi.String("roles/pubsub.publisher"),
// 			Members: pulumi.StringArray{
// 				pulumi.String(fmt.Sprintf("%v%v", "serviceAccount:", gcsAccount.EmailAddress)),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProjectServiceAccount(ctx *pulumi.Context, args *GetProjectServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetProjectServiceAccountResult, error) {
	var rv GetProjectServiceAccountResult
	err := ctx.Invoke("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectServiceAccount.
type GetProjectServiceAccountArgs struct {
	// The project the unique service account was created for. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The project the lookup originates from. This field is used if you are making the request
	// from a different account than the one you are finding the service account for.
	UserProject *string `pulumi:"userProject"`
}

// A collection of values returned by getProjectServiceAccount.
type GetProjectServiceAccountResult struct {
	// The email address of the service account. This value is often used to refer to the service account
	// in order to grant IAM permissions.
	EmailAddress string `pulumi:"emailAddress"`
	// The provider-assigned unique ID for this managed resource.
	Id          string  `pulumi:"id"`
	Project     string  `pulumi:"project"`
	UserProject *string `pulumi:"userProject"`
}
