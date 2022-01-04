package repo

import (
	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
)

// CatalogStorageI
type CatalogStorageI interface {
	// CRUD for Books
	CreateBook(*pb.Book) (*pb.Book, error)
	UpdateBook(*pb.Book) (*pb.Book, error)
	GetBookById(*pb.GetBookByIdReq) (*pb.Book, error)
	DeletedBookById(*pb.GetBookByIdReq) (*pb.EmptyResp, error)
	ListBooks(*pb.ListBookReq) (*pb.ListBookResp, error)
	// CRUD for Authors
	CreateAuthor(*pb.Author) (*pb.Author, error)
	UpdateAuthor(*pb.Author) (*pb.Author, error)
	GetAuthorById(*pb.GetAuthorByIdReq) (*pb.Author, error)
	DeleteAuthorById(*pb.GetAuthorByIdReq) (*pb.EmptyResp, error)
	ListAuthors(*pb.ListAuthorReq) (*pb.ListAuthorResp, error)
	// CRUD for Categories
	CreateCategory(*pb.Category) (*pb.Category, error)
	UpdateCategory(*pb.Category) (*pb.Category, error)
	GetCategoryById(*pb.GetCategoryByIdReq) (*pb.Category, error)
	DeleteCategoryById(*pb.GetCategoryByIdReq) (*pb.EmptyResp, error)
	ListCategories(*pb.ListCategoryReq) (*pb.ListCategoryResp, error)
}
