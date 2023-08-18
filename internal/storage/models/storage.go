package models

import "mime/multipart"

type Response struct {
	Id          int    `json:"id,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
	Type        string `json:"type,omitempty"`
	CreatedBy   int    `json:"created_by,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type SaveRequest struct {
	Id          int    `json:"id"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Token       string `json:"token"`
	LifeTime    int    `json:"life_time"`
}

type UploadRequest struct {
	File *multipart.FileHeader `json:"file"`
}
