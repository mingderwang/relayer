package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"runtime"

	"github.com/iqlusioninc/relayer/server/graph/generated"
	"github.com/iqlusioninc/relayer/server/graph/model"
)

func (r *mutationResolver) Post(ctx context.Context, text string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Version(ctx context.Context) (*model.Version, error) {
	initConfig()
	v := &model.Version{
		Version:   Version,
		Commit:    Commit,
		CosmosSdk: SDKCommit,
		Gaia:      GaiaCommit,
		Go:        fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
	return v, nil
}

func (r *queryResolver) Config(ctx context.Context) (*model.Config, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Path(ctx context.Context, name string) (*model.PathStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Start(ctx context.Context, pathName string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	//config   *Config
	// following vars are create only when make, see Makefile
	// Version defines the application version (defined at compile time)
	Version = ""
	// Commit defines the application commit hash (defined at compile time)
	Commit = ""
	// SDKCommit defines the CosmosSDK commit hash (defined at compile time)
	SDKCommit = ""
	// GaiaCommit defines the Gaia commit hash (defined at compile time)
	GaiaCommit = ""
)

func (r *subscriptionResolver) MessageAdded(ctx context.Context, roomName string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}
