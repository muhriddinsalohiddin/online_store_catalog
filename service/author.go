package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/logger"
)

func (c *CatalogService) CreateAuthor(ctx context.Context, in *pb.Author) (*pb.Author, error) {
	id, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("failed to generate uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to generate uuid")
	}
	in.Id = id.String()

	author, err := c.storage.Author().CreateAuthor(*in)
	if err != nil {
		c.logger.Error("failed to create author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create author")
	}
	return &author, nil
}

func (c *CatalogService) UpdateAuthor(ctx context.Context, in *pb.Author) (*pb.Author, error) {
	author, err := c.storage.Author().UpdateAuthor(*in)
	if err != nil {
		c.logger.Error("failed to update author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update author")
	}
	return &author, nil
}

func (c *CatalogService) GetAuthorById(ctx context.Context, in *pb.GetAuthorByIdReq) (*pb.Author, error) {
	author, err := c.storage.Author().GetAuthorById(*in)
	if err != nil {
		c.logger.Error("failed to get author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get author")
	}
	return &author, nil
}

func (c *CatalogService) DeleteAuthorById(ctx context.Context, in *pb.GetAuthorByIdReq) (*pb.EmptyResp, error) {
	err := c.storage.Author().DeleteAuthorById(*in)
	if err != nil {
		c.logger.Error("failed to deleate author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete author")
	}
	return &pb.EmptyResp{}, nil
}

func (c *CatalogService) ListAuthors(ctx context.Context, in *pb.ListAuthorReq) (*pb.ListAuthorResp, error) {
	authors, err := c.storage.Author().ListAuthors(*in)
	if err != nil {
		c.logger.Error("failed to list author", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to list author")
	}
	return &authors, nil
}
