// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudbuild

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Configuration for an automated build in response to source repository changes.
//
// To get more information about Trigger, see:
//
// * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.triggers)
// * How-to Guides
//     * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)
//
// > **Note:** You can retrieve the email of the Cloud Build Service Account used in jobs by using the `projects.ServiceIdentity` resource.
//
// ## Example Usage
// ### Cloudbuild Trigger Filename
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudbuild"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudbuild.NewTrigger(ctx, "filename-trigger", &cloudbuild.TriggerArgs{
// 			Filename: pulumi.String("cloudbuild.yaml"),
// 			Substitutions: pulumi.StringMap{
// 				"_BAZ": pulumi.String("qux"),
// 				"_FOO": pulumi.String("bar"),
// 			},
// 			TriggerTemplate: &cloudbuild.TriggerTriggerTemplateArgs{
// 				BranchName: pulumi.String("master"),
// 				RepoName:   pulumi.String("my-repo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Cloudbuild Trigger Build
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/cloudbuild"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudbuild.NewTrigger(ctx, "build-trigger", &cloudbuild.TriggerArgs{
// 			Build: &cloudbuild.TriggerBuildArgs{
// 				Artifacts: &cloudbuild.TriggerBuildArtifactsArgs{
// 					Images: pulumi.StringArray{
// 						pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v", "gcr.io/", "$", "PROJECT_ID/", "$", "REPO_NAME:", "$", "COMMIT_SHA")),
// 					},
// 					Objects: &cloudbuild.TriggerBuildArtifactsObjectsArgs{
// 						Location: pulumi.String("gs://bucket/path/to/somewhere/"),
// 						Paths: pulumi.StringArray{
// 							pulumi.String("path"),
// 						},
// 					},
// 				},
// 				LogsBucket: pulumi.String("gs://mybucket/logs"),
// 				Options: &cloudbuild.TriggerBuildOptionsArgs{
// 					DiskSizeGb:           pulumi.Int(100),
// 					DynamicSubstitutions: pulumi.Bool(true),
// 					Env: pulumi.StringArray{
// 						pulumi.String("ekey = evalue"),
// 					},
// 					LogStreamingOption:    pulumi.String("STREAM_OFF"),
// 					Logging:               pulumi.String("LEGACY"),
// 					MachineType:           pulumi.String("N1_HIGHCPU_8"),
// 					RequestedVerifyOption: pulumi.String("VERIFIED"),
// 					SecretEnv: pulumi.StringArray{
// 						pulumi.String("secretenv = svalue"),
// 					},
// 					SourceProvenanceHash: pulumi.StringArray{
// 						pulumi.String("MD5"),
// 					},
// 					SubstitutionOption: pulumi.String("ALLOW_LOOSE"),
// 					Volumes: cloudbuild.TriggerBuildOptionsVolumeArray{
// 						&cloudbuild.TriggerBuildOptionsVolumeArgs{
// 							Name: pulumi.String("v1"),
// 							Path: pulumi.String("v1"),
// 						},
// 					},
// 					WorkerPool: pulumi.String("pool"),
// 				},
// 				QueueTtl: pulumi.String("20s"),
// 				Secrets: cloudbuild.TriggerBuildSecretArray{
// 					&cloudbuild.TriggerBuildSecretArgs{
// 						KmsKeyName: pulumi.String("projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name"),
// 						SecretEnv: pulumi.StringMap{
// 							"PASSWORD": pulumi.String("ZW5jcnlwdGVkLXBhc3N3b3JkCg=="),
// 						},
// 					},
// 				},
// 				Source: &cloudbuild.TriggerBuildSourceArgs{
// 					StorageSource: &cloudbuild.TriggerBuildSourceStorageSourceArgs{
// 						Bucket: pulumi.String("mybucket"),
// 						Object: pulumi.String("source_code.tar.gz"),
// 					},
// 				},
// 				Steps: cloudbuild.TriggerBuildStepArray{
// 					&cloudbuild.TriggerBuildStepArgs{
// 						Args: pulumi.StringArray{
// 							pulumi.String("cp"),
// 							pulumi.String("gs://mybucket/remotefile.zip"),
// 							pulumi.String("localfile.zip"),
// 						},
// 						Name:    pulumi.String("gcr.io/cloud-builders/gsutil"),
// 						Timeout: pulumi.String("120s"),
// 					},
// 				},
// 				Substitutions: pulumi.StringMap{
// 					"_BAZ": pulumi.String("qux"),
// 					"_FOO": pulumi.String("bar"),
// 				},
// 				Tags: pulumi.StringArray{
// 					pulumi.String("build"),
// 					pulumi.String("newFeature"),
// 				},
// 			},
// 			TriggerTemplate: &cloudbuild.TriggerTriggerTemplateArgs{
// 				BranchName: pulumi.String("master"),
// 				RepoName:   pulumi.String("my-repo"),
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
// Trigger can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default projects/{{project}}/triggers/{{trigger_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{project}}/{{trigger_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{trigger_id}}
// ```
type Trigger struct {
	pulumi.CustomResourceState

	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrOutput `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human-readable description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrOutput `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrOutput `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayOutput `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayOutput `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapOutput `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The unique identifier for the trigger.
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrOutput `pulumi:"triggerTemplate"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		args = &TriggerArgs{}
	}

	var resource Trigger
	err := ctx.RegisterResource("gcp:cloudbuild/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("gcp:cloudbuild/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build *TriggerBuild `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime *string `pulumi:"createTime"`
	// Human-readable description of the trigger.
	Description *string `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename *string `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	Github *TriggerGithub `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `pulumi:"tags"`
	// The unique identifier for the trigger.
	TriggerId *string `pulumi:"triggerId"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	TriggerTemplate *TriggerTriggerTemplate `pulumi:"triggerTemplate"`
}

type TriggerState struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrInput
	// Time when the trigger was created.
	CreateTime pulumi.StringPtrInput
	// Human-readable description of the trigger.
	Description pulumi.StringPtrInput
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrInput
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayInput
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayInput
	// The unique identifier for the trigger.
	TriggerId pulumi.StringPtrInput
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build *TriggerBuild `pulumi:"build"`
	// Human-readable description of the trigger.
	Description *string `pulumi:"description"`
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename *string `pulumi:"filename"`
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	Github *TriggerGithub `pulumi:"github"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `pulumi:"tags"`
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	TriggerTemplate *TriggerTriggerTemplate `pulumi:"triggerTemplate"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Contents of the build template. Either a filename or build template must be provided.
	// Structure is documented below.
	Build TriggerBuildPtrInput
	// Human-readable description of the trigger.
	Description pulumi.StringPtrInput
	// Whether the trigger is disabled or not. If true, the trigger will never result in a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
	Filename pulumi.StringPtrInput
	// Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	Github TriggerGithubPtrInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If ignoredFiles and changed files are both empty, then they are not
	// used to determine whether or not to trigger a build.
	// If ignoredFiles is not empty, then we ignore any files that match any
	// of the ignoredFile globs. If the change has no files that are outside
	// of the ignoredFiles globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
	// extended with support for `**`.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is empty, then as far as this filter is concerned, we
	// should trigger the build.
	// If any of the files altered in the commit pass the ignoredFiles filter
	// and includedFiles is not empty, then we make sure that at least one of
	// those files matches a includedFiles glob. If not, then we do not trigger
	// a build.
	IncludedFiles pulumi.StringArrayInput
	// Name of the volume to mount.
	// Volume names must be unique per build step and must be valid names for Docker volumes.
	// Each named volume must be used by at least two build steps.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Substitutions to use in a triggered build. Should only be used with triggers.run
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a Build. These are not docker tags.
	Tags pulumi.StringArrayInput
	// Template describing the types of source changes to trigger a build.
	// Branch and tag names in trigger templates are interpreted as regular
	// expressions. Any branch or tag change that matches that regular
	// expression will trigger a build.
	// One of `triggerTemplate` or `github` must be provided.
	// Structure is documented below.
	TriggerTemplate TriggerTriggerTemplatePtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct {
	*pulumi.OutputState
}

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TriggerOutput{})
}
