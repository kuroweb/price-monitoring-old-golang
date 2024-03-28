package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/kuroweb/price-monitoring/volumes/bff/graph/model"
	"github.com/kuroweb/price-monitoring/volumes/bff/internal"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (model.CreateProductResult, error) {
	return r.ProductService.CreateProduct(ctx, input)
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProductInput) (model.UpdateProductResult, error) {
	return r.ProductService.UpdateProduct(ctx, id, input)
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (model.DeleteProductResult, error) {
	return r.ProductService.DeleteProductById(ctx, id)
}

// CreateYahooAuctionCrawlSettingExcludeKeyword is the resolver for the createYahooAuctionCrawlSettingExcludeKeyword field.
func (r *mutationResolver) CreateYahooAuctionCrawlSettingExcludeKeyword(ctx context.Context, input model.CreateYahooAuctionCrawlSettingExcludeKeywordInput) (model.CreateYahooAuctionCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.CreateYahooAuctionCrawlSettingExcludeKeyword(ctx, input)
}

// UpdateYahooAuctionCrawlSettingExcludeKeyword is the resolver for the updateYahooAuctionCrawlSettingExcludeKeyword field.
func (r *mutationResolver) UpdateYahooAuctionCrawlSettingExcludeKeyword(ctx context.Context, input model.UpdateYahooAuctionCrawlSettingExcludeKeywordInput) (model.UpdateYahooAuctionCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.UpdateYahooAuctionCrawlSettingExcludeKeyword(ctx, input)
}

// DeleteYahooAuctionCrawlSettingExcludeKeyword is the resolver for the DeleteYahooAuctionCrawlSettingExcludeKeyword field.
func (r *mutationResolver) DeleteYahooAuctionCrawlSettingExcludeKeyword(ctx context.Context, id string, productID string) (model.DeleteYahooAuctionCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.DeleteYahooAuctionCrawlSettingExcludeKeyword(ctx, id, productID)
}

// CreateYahooAuctionCrawlSettingRequiredKeyword is the resolver for the createYahooAuctionCrawlSettingRequiredKeyword field.
func (r *mutationResolver) CreateYahooAuctionCrawlSettingRequiredKeyword(ctx context.Context, input model.CreateYahooAuctionCrawlSettingRequiredKeywordInput) (model.CreateYahooAuctionCrawlSettingRequiredKeywordResult, error) {
	panic(fmt.Errorf("not implemented: CreateYahooAuctionCrawlSettingRequiredKeyword - createYahooAuctionCrawlSettingRequiredKeyword"))
}

// UpdateYahooAuctionCrawlSettingRequiredKeyword is the resolver for the updateYahooAuctionCrawlSettingRequiredKeyword field.
func (r *mutationResolver) UpdateYahooAuctionCrawlSettingRequiredKeyword(ctx context.Context, input model.UpdateYahooAuctionCrawlSettingRequiredKeywordInput) (model.UpdateYahooAuctionCrawlSettingRequiredKeywordResult, error) {
	panic(fmt.Errorf("not implemented: UpdateYahooAuctionCrawlSettingRequiredKeyword - updateYahooAuctionCrawlSettingRequiredKeyword"))
}

// DeleteYahooAuctionCrawlSettingRequiredKeyword is the resolver for the deleteYahooAuctionCrawlSettingRequiredKeyword field.
func (r *mutationResolver) DeleteYahooAuctionCrawlSettingRequiredKeyword(ctx context.Context, id string, productID string) (model.DeleteYahooAuctionCrawlSettingExcludeKeywordResult, error) {
	panic(fmt.Errorf("not implemented: DeleteYahooAuctionCrawlSettingRequiredKeyword - deleteYahooAuctionCrawlSettingRequiredKeyword"))
}

// CreateMercariCrawlSettingExcludeKeyword is the resolver for the createMercariCrawlSettingExcludeKeyword field.
func (r *mutationResolver) CreateMercariCrawlSettingExcludeKeyword(ctx context.Context, input model.CreateMercariCrawlSettingExcludeKeywordInput) (model.CreateMercariCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.CreateMercariCrawlSettingExcludeKeyword(ctx, input)
}

// UpdateMercariCrawlSettingExcludeKeyword is the resolver for the updateMercariCrawlSettingExcludeKeyword field.
func (r *mutationResolver) UpdateMercariCrawlSettingExcludeKeyword(ctx context.Context, input model.UpdateMercariCrawlSettingExcludeKeywordInput) (model.UpdateMercariCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.UpdateMercariCrawlSettingExcludeKeyword(ctx, input)
}

// DeleteMercariCrawlSettingExcludeKeyword is the resolver for the deleteMercariCrawlSettingExcludeKeyword field.
func (r *mutationResolver) DeleteMercariCrawlSettingExcludeKeyword(ctx context.Context, id string, productID string) (model.DeleteMercariCrawlSettingExcludeKeywordResult, error) {
	return r.ProductService.DeleteMercariCrawlSettingExcludeKeyword(ctx, id, productID)
}

// CreateMercariCrawlSettingRequiredKeyword is the resolver for the createMercariCrawlSettingRequiredKeyword field.
func (r *mutationResolver) CreateMercariCrawlSettingRequiredKeyword(ctx context.Context, input model.CreateMercariCrawlSettingRequiredKeywordInput) (model.CreateMercariCrawlSettingRequiredKeywordResult, error) {
	panic(fmt.Errorf("not implemented: CreateMercariCrawlSettingRequiredKeyword - createMercariCrawlSettingRequiredKeyword"))
}

// UpdateMercariCrawlSettingRequiredKeyword is the resolver for the updateMercariCrawlSettingRequiredKeyword field.
func (r *mutationResolver) UpdateMercariCrawlSettingRequiredKeyword(ctx context.Context, input model.UpdateMercariCrawlSettingRequiredKeywordInput) (model.UpdateMercariCrawlSettingRequiredKeywordResult, error) {
	panic(fmt.Errorf("not implemented: UpdateMercariCrawlSettingRequiredKeyword - updateMercariCrawlSettingRequiredKeyword"))
}

// DeleteMercariCrawlSettingRequiredKeyword is the resolver for the deleteMercariCrawlSettingRequiredKeyword field.
func (r *mutationResolver) DeleteMercariCrawlSettingRequiredKeyword(ctx context.Context, id string, productID string) (model.DeleteMercariCrawlSettingRequiredKeywordResult, error) {
	panic(fmt.Errorf("not implemented: DeleteMercariCrawlSettingRequiredKeyword - deleteMercariCrawlSettingRequiredKeyword"))
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
