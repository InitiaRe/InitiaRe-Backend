package repository

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

type ctnRepo struct {
	ctn *azblob.Client
}

func NewContainerRepo(ctn *azblob.Client) IContainerRepository {
	return &ctnRepo{
		ctn: ctn,
	}
}
