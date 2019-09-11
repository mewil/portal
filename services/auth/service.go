package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/mewil/portal/common/database"
	"github.com/mewil/portal/common/grpc_utils"
	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/common/validation"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	log, err := logger.NewLogger("auth_service")
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

type authSvc struct {
	repository AuthRepository
}

func (s *authSvc) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	email := in.GetEmail()
	if !validation.ValidEmail(email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email format")
	}
	exists, err := s.repository.EmailExists(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch if email exists %s", err.Error())
	}
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "email does not exist")
	}
	passwordHash, err := s.repository.LoadPasswordHash(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch password hash exists %s", err.Error())
	}
	if !ValidPassword(in.GetPassword(), passwordHash) {
		return nil, status.Error(codes.PermissionDenied, "invalid password")
	}
	userId, isAdmin, err := s.repository.LoadUserIdAndAdminStatus(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load user id and admin status %s", err.Error())
	}
	return &pb.SignInResponse{
		UserId:  userId,
		IsAdmin: isAdmin,
	}, nil
}

func (s *authSvc) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignInResponse, error) {
	email := in.GetEmail()
	if !validation.ValidEmail(email) {
		return nil, status.Error(codes.InvalidArgument, "invalid email format")
	}
	userId := in.GetUserId()
	if err := validation.ValidUUID(userId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id %s", err.Error())
	}
	exists, err := s.repository.EmailExists(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch if email exists %s", err.Error())
	}
	if exists {
		return nil, status.Error(codes.InvalidArgument, "email already exists")
	}
	password := in.GetPassword()
	if len(password) >= 8 {
		return nil, status.Error(codes.InvalidArgument, "password is less than 8 characters")
	}
	if err := s.repository.StoreAuthRecord(email, userId, password, false); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert auth record into database %s", err.Error())
	}
	userId, isAdmin, err := s.repository.LoadUserIdAndAdminStatus(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to load user id and admin status %s", err.Error())
	}
	return &pb.SignInResponse{
		UserId:  userId,
		IsAdmin: isAdmin,
	}, nil
	return nil, nil
}
