package service

import (
	_ "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	log "github.com/muhriddinsalohiddin/online_store_catalog/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_catalog/storage"
)

type CatalogService struct {
	storage storage.IStorage
	logger  log.Logger
}

func NewCatalogService(storage storage.IStorage, log log.Logger) *CatalogService {
	return &CatalogService{
		storage: storage,
		logger:  log,
	}
}
