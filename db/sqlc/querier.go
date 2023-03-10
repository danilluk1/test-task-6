// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateProductsCategory(ctx context.Context, arg CreateProductsCategoryParams) (ProductsCategory, error)
	CreateShop(ctx context.Context, arg CreateShopParams) (Shop, error)
	CreateShopCategory(ctx context.Context, arg CreateShopCategoryParams) (ShopsCategory, error)
	DeleteProduct(ctx context.Context, id int32) error
	DeleteProductsCategoriesRelationshipByProductCategoryId(ctx context.Context, productCategoryID int32) error
	DeleteProductsCategoriesRelationshipByProductId(ctx context.Context, productID int32) error
	DeleteProductsCategory(ctx context.Context, id int32) error
	DeleteShop(ctx context.Context, id int32) error
	DeleteShopsCategoriesRelationshipByShopId(ctx context.Context, shopID int32) error
	DeleteShopsCategoriesRelationshipByShopsCategoryId(ctx context.Context, shopCategoryID int32) error
	DeleteShopsCategory(ctx context.Context, id int32) error
	GetProduct(ctx context.Context, id int32) (Product, error)
	GetProductsCategories(ctx context.Context, arg GetProductsCategoriesParams) ([]ProductsCategory, error)
	GetShop(ctx context.Context, id int32) (Shop, error)
	GetShopCategory(ctx context.Context, id int32) (ShopsCategory, error)
	GetShopsCategories(ctx context.Context, arg GetShopsCategoriesParams) ([]ShopsCategory, error)
	InsertNewProductsCategoriesRelationship(ctx context.Context, arg InsertNewProductsCategoriesRelationshipParams) (ProductsProductsCategory, error)
	InsertNewShopCategoriesRelationship(ctx context.Context, arg InsertNewShopCategoriesRelationshipParams) (ShopsShopsCategory, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListShops(ctx context.Context, arg ListShopsParams) ([]Shop, error)
	UpdateShop(ctx context.Context, arg UpdateShopParams) (Shop, error)
	UpdateShopCategory(ctx context.Context, arg UpdateShopCategoryParams) (ShopsCategory, error)
}

var _ Querier = (*Queries)(nil)
