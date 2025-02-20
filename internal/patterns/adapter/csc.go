package adapter

type StorageClient interface {
	UploadFile(fileName string, content string) string
	GetContent(fileName string) string
}
