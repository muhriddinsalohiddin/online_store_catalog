package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/muhriddinsalohiddin/online_store_catalog/storage/postgres"
	"github.com/muhriddinsalohiddin/online_store_catalog/storage/repo"
)

// IStorage ...
type IStorage interface {
	Catalog() repo.CatalogStorageI
}

type storagePg struct {
	db          *sqlx.DB
	catalogRepo repo.CatalogStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		catalogRepo: postgres.NewCatalogRepo(db),
	}
}

func (s storagePg) Catalog() repo.CatalogStorageI {
	return s.catalogRepo
}
