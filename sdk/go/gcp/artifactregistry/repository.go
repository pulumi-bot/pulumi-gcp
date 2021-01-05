// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A repository for storing artifacts
//
// To get more information about Repository, see:
//
// * [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
//
// ## Example Usage
// ### Artifact Registry Repository Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactregistry.NewRepository(ctx, "my-repo", &artifactregistry.RepositoryArgs{
// 			Location:     pulumi.String("us-central1"),
// 			RepositoryId: pulumi.String("my-repository"),
// 			Description:  pulumi.String("example docker repository"),
// 			Format:       pulumi.String("DOCKER"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Artifact Registry Repository Cmek
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactregistry.NewRepository(ctx, "my-repo", &artifactregistry.RepositoryArgs{
// 			Location:     pulumi.String("us-central1"),
// 			RepositoryId: pulumi.String("my-repository"),
// 			Description:  pulumi.String("example docker repository with cmek"),
// 			Format:       pulumi.String("DOCKER"),
// 			KmsKeyName:   pulumi.String("kms-key"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Artifact Registry Repository Iam
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/artifactregistry"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/serviceAccount"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactregistry.NewRepository(ctx, "my-repo", &artifactregistry.RepositoryArgs{
// 			Location:     pulumi.String("us-central1"),
// 			RepositoryId: pulumi.String("my-repository"),
// 			Description:  pulumi.String("example docker repository with iam"),
// 			Format:       pulumi.String("DOCKER"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = serviceAccount.NewAccount(ctx, "test-account", &serviceAccount.AccountArgs{
// 			AccountId:   pulumi.String("my-account"),
// 			DisplayName: pulumi.String("Test Service Account"),
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = artifactregistry.NewRepositoryIamMember(ctx, "test-iam", &artifactregistry.RepositoryIamMemberArgs{
// 			Location:   my_repo.Location,
// 			Repository: my_repo.Name,
// 			Role:       pulumi.String("roles/artifactregistry.reader"),
// 			Member: test_account.Email.ApplyT(func(email string) (string, error) {
// 				return fmt.Sprintf("%v%v", "serviceAccount:", email), nil
// 			}).(pulumi.StringOutput),
// 		}, pulumi.Provider(google_beta))
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
// Repository can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repository:Repository default projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repository:Repository default {{project}}/{{location}}/{{repository_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repository:Repository default {{location}}/{{repository_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:artifactregistry/repository:Repository default {{repository_id}}
// ```
type Repository struct {
	pulumi.CustomResourceState

	// The time when the repository was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The format of packages that are stored in the repository.
	// Possible values are `DOCKER`.
	Format pulumi.StringOutput `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// The time when the repository was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	var resource Repository
	err := ctx.RegisterResource("gcp:artifactregistry/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("gcp:artifactregistry/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// The time when the repository was created.
	CreateTime *string `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description *string `pulumi:"description"`
	// The format of packages that are stored in the repository.
	// Possible values are `DOCKER`.
	Format *string `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location *string `pulumi:"location"`
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId *string `pulumi:"repositoryId"`
	// The time when the repository was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type RepositoryState struct {
	// The time when the repository was created.
	CreateTime pulumi.StringPtrInput
	// The user-provided description of the repository.
	Description pulumi.StringPtrInput
	// The format of packages that are stored in the repository.
	// Possible values are `DOCKER`.
	Format pulumi.StringPtrInput
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrInput
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringPtrInput
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1"
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringPtrInput
	// The time when the repository was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The user-provided description of the repository.
	Description *string `pulumi:"description"`
	// The format of packages that are stored in the repository.
	// Possible values are `DOCKER`.
	Format string `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location this repository is located in.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId string `pulumi:"repositoryId"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The user-provided description of the repository.
	Description pulumi.StringPtrInput
	// The format of packages that are stored in the repository.
	// Possible values are `DOCKER`.
	Format pulumi.StringInput
	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	// This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrInput
	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels pulumi.StringMapInput
	// The name of the location this repository is located in.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The last part of the repository name, for example:
	// "repo1"
	RepositoryId pulumi.StringInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct {
	*pulumi.OutputState
}

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil))
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RepositoryOutput{})
}
