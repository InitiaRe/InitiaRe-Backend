package models

type UploadRequest struct {
	Obj           []byte `json:"obj"`
	ContainerName string `json:"container_name"`
	FileName      string `json:"file_name"`
}
