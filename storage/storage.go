package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/muhriddinsalohiddin/online_store_catalog/storage/postgres"
	"github.com/muhriddinsalohiddin/online_store_catalog/storage/repo"
)

// IStorage ...
type IStorage interface {
	Book() repo.BookStorageI
	Category() repo.CategoryStorageI
	Author() repo.AuthorStorageI
}

type storagePg struct {
	db          *sqlx.DB
	bookRepo repo.BookStorageI
	categoryRepo repo.CategoryStorageI
	authorRepo repo.AuthorStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		bookRepo: postgres.NewCatalogRepo(db),
		categoryRepo: postgres.NewCatalogRepo(db),
		authorRepo: postgres.NewCatalogRepo(db),
	}
}

func (s storagePg) Book() repo.BookStorageI {
	return s.bookRepo
}
func (s storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}
func (s storagePg) Author() repo.AuthorStorageI {
	return s.authorRepo
}
