package gcp

type CloudStorageAdapter struct {
	CloudStorageClient *CloudStorage
}

func (csa *CloudStorageAdapter) UploadFile(fileName string, content string) string {
	return csa.CloudStorageClient.Upload(fileName, content)
}

func (csa *CloudStorageAdapter) GetContent(fileName string) string {
	return csa.CloudStorageClient.Content(fileName)
}
