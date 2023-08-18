package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Ho-Minh/InitiaRe-website/config"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func InitRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

type ctnRepo struct {
	cfg *config.Config
	ctn *azblob.Client
}

func InitContainerRepo(cfg *config.Config, ctn *azblob.Client) IContainerRepository {
	return &ctnRepo{
		cfg: cfg,
		ctn: ctn,
	}
}
