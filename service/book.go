package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/logger"
)

func (c *CatalogService) CreateBook(ctx context.Context, in *pb.Book) (*pb.Book, error) {
	id, err := uuid.NewV4()
	if err != nil {
		c.logger.Error("failed to generate uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to generate uuid")
	}
	in.Id = id.String()

	book, err := c.storage.Book().CreateBook(*in)
	if err != nil {
		c.logger.Error("failed to create Book", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to create Book")
	}
	return &book, nil
}

func (c *CatalogService) UpdateBook(ctx context.Context, in *pb.Book) (*pb.Book, error) {
	book, err := c.storage.Book().UpdateBook(*in)
	if err != nil {
		c.logger.Error("failed to update Book", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to update Book")
	}
	return &book, nil
}

func (c *CatalogService) GetBookById(ctx context.Context, in *pb.GetBookByIdReq) (*pb.Book, error) {
	book, err := c.storage.Book().GetBookById(*in)
	if err != nil {
		c.logger.Error("failed to get Book", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to get Book")
	}
	return &book, nil
}

func (c *CatalogService) DeletedBookById(ctx context.Context, in *pb.GetBookByIdReq) (*pb.EmptyResp, error) {
	resp, err := c.storage.Book().DeletedBookById(*in)
	if err != nil {
		c.logger.Error("failed to deleate Book", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete Book")
	}
	return &resp, nil
}

func (c *CatalogService) ListBooks(ctx context.Context, in *pb.ListBookReq) (*pb.ListBookResp, error) {
	books, err := c.storage.Book().ListBooks(*in)
	if err != nil {
		c.logger.Error("failed to list Book", logger.Error(err))
		return nil, status.Error(codes.Internal, "failed to list Book")
	}
	return &books, nil
}
