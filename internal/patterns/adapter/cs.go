package adapter

type CloudStorage interface {
	UploadFile(fileName string, content string) string
	GetContent(fileName string) string
}
