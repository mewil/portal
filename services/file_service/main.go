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

	fileRepository, err := NewFileRepository(
		log,
		"",
	)
	if err != nil {
		log.Fatal("failed to initialize file repository", err)
	}
	pb.RegisterFileSvcServer(s, &fileSvc{
		repository: fileRepository,
	})
	s.Serve(listener)
}
