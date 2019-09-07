package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	s, err := grpc_utils.NewServer(log)
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterAuthSvcServer(s, &authSvc{})
	s.Serve(listener)
}

type authSvc struct {
}

func (s *authSvc) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, nil
}

func (s *authSvc) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignInResponse, error) {
	return nil, nil
}
