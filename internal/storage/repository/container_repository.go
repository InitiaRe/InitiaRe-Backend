package repository

import (
	"context"
	"io"
	"os"
	"path"

	"InitiaRe-website/internal/storage/models"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/vukyn/kuery/crypto"
)

func (c *ctnRepo) Upload(ctx context.Context, req *models.UploadRequest) (string, error) {

	// Source
	src, err := req.File.Open()
	if err != nil {
		return "", err
	}

	// Destination
	dst, err := os.Create(req.File.Filename)
	if err != nil {
		return "", err
	}

	defer func() {
		src.Close()
		dst.Close()
		os.Remove(req.File.Filename)
	}()

	// Copy
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	blobName := crypto.HashedToken() + "-" + req.File.Filename
	if _, err := c.ctn.UploadFile(ctx, c.cfg.Storage.Container, blobName, dst, &azblob.UploadBufferOptions{}); err != nil {
		return "", err
	}
	return path.Join(c.cfg.Storage.Host, c.cfg.Storage.Container, blobName), nil
}
