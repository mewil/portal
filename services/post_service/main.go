package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/database"
	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
)

func main() {
	log, err := logger.NewLogger("post_service")
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
	postRepository, err := NewPostRepository(
		log,
		db,
	)
	if err != nil {
		log.Fatal("failed to initialize post repository", err)
	}
	pb.RegisterPostSvcServer(s, &postSvc{
		repository: postRepository,
	})
	s.Serve(listener)
}
