package api

import (
	"bytes"
	"context"
	"io"

	"github.com/mewil/portal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FrontendSvc) FileSvcGetFileStats(ctx context.Context, newFileSvcClient FileSvcInjector, fileID string) (*pb.FileStats, error) {
	req := &pb.FileRequest{FileID: fileID}
	return newFileSvcClient().GetFileStats(ctx, req)
}

func (s *FrontendSvc) FileSvcGetFile(ctx context.Context, newFileSvcClient FileSvcInjector, fileID string) ([]byte, error) {
	req := &pb.FileRequest{FileID: fileID}
	stream, err := newFileSvcClient().GetFile(ctx, req)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})
	for {
		filePart, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, status.Error(codes.Internal, "failed to receive file")
		}
		buf.Write(filePart.GetContent())
	}
	println(buf.Len())
	return buf.Bytes(), nil
}
