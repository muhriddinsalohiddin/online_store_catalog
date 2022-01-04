package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/muhriddinsalohiddin/online_store_catalog/config"
	pb "github.com/muhriddinsalohiddin/online_store_catalog/genproto/catalog_service"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/db"
	"github.com/muhriddinsalohiddin/online_store_catalog/pkg/logger"
	"github.com/muhriddinsalohiddin/online_store_catalog/service"
	"github.com/muhriddinsalohiddin/online_store_catalog/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "catalog-service")

	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres err", logger.Error(err))
	}

	pgStorage := storage.NewStoragePg(connDB)

	catalogService := service.NewCatalogService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Connection GRPC error", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCatalogServiceServer(s, catalogService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening", logger.Error(err))
	}
}
