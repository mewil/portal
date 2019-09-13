package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
)

func main() {
	log, err := logger.NewLogger("file_service")
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("failed to start tcp listener", err)
	}

	s, err := grpc_utils.NewServer(log)
	if err != nil {
		log.Fatal("failed to initialize grpc server", err)
	}

	db, err := database.NewDatabase(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	authRepository, err := NewAuthRepository(
		log,
		db,
		os.Getenv("ADMIN_EMAIL"),
		os.Getenv("ADMIN_PASSWORD"),
	)
	if err != nil {
		log.Fatal("failed to initialize auth repository", err)
	}
	pb.RegisterAuthSvcServer(s, &authSvc{
		repository: authRepository,
	})
	s.Serve(listener)
}
