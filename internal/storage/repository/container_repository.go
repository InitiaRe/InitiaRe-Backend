package repository

import (
	"context"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/models"
	"github.com/vukyn/go-kuqery/krypto"
)

func (c *ctnRepo) Upload(ctx context.Context, req *models.UploadRequest) (string, error) {
	blobName := krypto.HashedToken()
	_, err := c.ctn.UploadBuffer(ctx, c.cfg.Storage.Container, blobName, req.Obj, &azblob.UploadBufferOptions{})
	if err != nil {
		return "", err
	}
	return filepath.Join(c.cfg.Storage.Host, c.cfg.Storage.Container, blobName), nil
}
