package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/logger"
)

func (c *CatalogService) CreateCategory(ctx context.Context, in *pb.Category) (*pb.Category, error) {
	id, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("failed to generate uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to generate uuid")
	}
	in.Id = id.String()

	category, err := c.storage.Catalog().CreateCategory(*in)
	if err != nil {
		c.logger.Error("failed to create category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create category")
	}
	return &category, nil
}

func (c *CatalogService) UpdateCategory(ctx context.Context, in *pb.Category) (*pb.Category, error) {
	category, err := c.storage.Catalog().UpdateCategory(*in)

	if err != nil {
		c.logger.Error("failed to update category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update category")
	}
	return &category, nil
}
func (c *CatalogService) GetCategoryById(ctx context.Context, in *pb.GetCategoryByIdReq) (*pb.Category, error) {
	category, err := c.storage.Catalog().GetCategoryById(*in)

	if err != nil {
		c.logger.Error("failed to get category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get category")
	}
	return &category, nil
}
func (c *CatalogService) DeleteCategoryById(ctx context.Context, in *pb.GetCategoryByIdReq) (*pb.EmptyResp, error) {
	err := c.storage.Catalog().DeleteCategoryById(*in)

	if err != nil {
		c.logger.Error("failed to deleate category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete category")
	}
	return &pb.EmptyResp{}, nil

}
func (c *CatalogService) ListCategories(ctx context.Context, in *pb.ListCategoryReq) (*pb.ListCategoryResp, error) {
	categories, err := c.storage.Catalog().ListCategories(*in)

	if err != nil {
		c.logger.Error("failed to list category", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to list category")
	}
	return &categories, nil
}
