package google

import (
	"context"
	"fmt"
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
}

// var uploader *ClientUploader

func InitStorage(fPath string, bucketName string, projectID string) *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", fPath) // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
	}

}

func (c *ClientUploader) UploadFile(file multipart.File, uploadPath, object string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	link := fmt.Sprintf("https://storage.googleapis.com/be10-petdopter/%s%s", uploadPath, object)
	return link, nil
}
