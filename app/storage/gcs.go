package storage

import (
	"context"
	"fmt"
	"go-gcs/app/config"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader

func GetStorageClinet() *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GCP_SERVICE_ACCOUNT_KEY)
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: config.GCS_BUCKET_NAME,
		projectID:  config.GCP_PROJECT_ID,
		uploadPath: "test-files/",
	}

	return uploader
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	return wc.Name, nil
}
