package main

import (
	"io"

	"github.com/mewil/portal/common/logger"
	"github.com/mewil/portal/pb"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
)

type FileRepository interface {
	PutFile(file io.Reader, size int, fileID string) error
	GetFile(fileID string) (io.Reader, error)
	GetFileStats(fileID string) (*pb.FileStats, error)
}

const defaultMinioRegion = "us-east-1"

func NewFileRepository(log logger.Logger, storeAddr, accessID, secretKey, bucketName string) (FileRepository, error) {
	client, err := minio.New(storeAddr, accessID, secretKey, false)
	if err != nil {
		return nil, err
	}
	r := repository{
		log:        log.(*zap.SugaredLogger).Named("repository"),
		store:      client,
		bucketName: bucketName,
	}
	found, err := r.store.BucketExists(r.bucketName)
	if err != nil {
		return nil, err
	}
	if !found {
		if err = r.store.MakeBucket(r.bucketName, defaultMinioRegion); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type repository struct {
	log        logger.Logger
	store      *minio.Client
	bucketName string
}

func (r *repository) PutFile(file io.Reader, size int, fileID string) error {
	_, err := r.store.PutObject(r.bucketName, fileID, file, int64(size), minio.PutObjectOptions{})
	return err
}

func (r *repository) GetFile(fileID string) (io.Reader, error) {
	return r.store.GetObject(r.bucketName, fileID, minio.GetObjectOptions{})
}

func (r *repository) GetFileStats(fileID string) (*pb.FileStats, error) {
	stats, err := r.store.StatObject(r.bucketName, fileID, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	return &pb.FileStats{
		ETag:     stats.ETag,
		FileSize: stats.Size,
	}, nil
}
