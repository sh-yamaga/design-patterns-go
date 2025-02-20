package aws

type S3Adapter struct {
	S3Client *S3
}

func (sa *S3Adapter) UploadFile(fileName string, content string) string {
	return sa.S3Client.Upload(fileName, content)
}

func (sa *S3Adapter) GetContent(fileName string) string {
	return sa.S3Client.Content(fileName)
}
