package main

import (
	"bytes"
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/mewil/portal/common/validation"
	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fileSvc struct {
	repository FileRepository
}

const fileBufferSize = 1024

func (s *fileSvc) Upload(stream pb.FileSvc_UploadServer) error {
	buf := bytes.NewBuffer([]byte{})
	for {
		filePart, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return status.Errorf(codes.InvalidArgument, "failed to read file %s", err.Error())
		}
		buf.Write(filePart.GetContent())
	}
	fileID, err := uuid.NewUUID()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to generate UUID for new file %s", err.Error())
	}
	err = s.repository.PutFile(buf, buf.Len(), fileID.String())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to upload file %s", err.Error())
	}
	return stream.SendAndClose(&pb.UploadStatus{
		FileID: fileID.String(),
		Status: pb.UploadStatusCode_Ok,
	})
}

func (s *fileSvc) GetFile(req *pb.FileRequest, stream pb.FileSvc_GetFileServer) error {
	fileID := req.GetFileID()
	if err := validation.ValidUUID(fileID); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid file id %s", err.Error())
	}
	reader, err := s.repository.GetFile(fileID)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to fetch file %s", err.Error())
	}
	buf := make([]byte, 0, fileBufferSize)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return status.Errorf(codes.Internal, "failed to read file %s", err.Error())
		}
		if err = stream.Send(&pb.FilePart{
			Content: buf[:n],
		}); err != nil {
			return status.Errorf(codes.Internal, "failed to send file part %s", err.Error())
		}
	}
	return nil
}

func (s *fileSvc) GetFileStats(ctx context.Context, req *pb.FileRequest) (*pb.FileStats, error) {
	fileID := req.GetFileID()
	if err := validation.ValidUUID(fileID); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid file id %s", err.Error())
	}
	return s.repository.GetFileStats(fileID)
}
