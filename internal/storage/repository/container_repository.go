package repository

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/models"
)

func (c *ctnRepo) Upload(ctx context.Context, req *models.UploadRequest) (string, error) {
	_, err := c.ctn.UploadBuffer(ctx, req.ContainerName, req.FileName, req.Obj, &azblob.UploadBufferOptions{}); 
	if err != nil {
		return "", err
	}
	return "", nil
}
