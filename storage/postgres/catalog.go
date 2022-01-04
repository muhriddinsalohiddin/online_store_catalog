package postgres

import (
	"github.com/jmoiron/sqlx"
)

type catalogRepo struct {
	db *sqlx.DB
}

// NewCatalogRepo
func NewCatalogRepo(db *sqlx.DB) *catalogRepo {
	return &catalogRepo{db: db}
}
