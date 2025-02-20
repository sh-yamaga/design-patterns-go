package aws

type S3 struct {
	Storage map[string]string
}

type S3Client interface {
	Upload(fileName string, content string) string
	Content(fileName string) string
}

func (s *S3) Upload(fileName string, content string) string {
	s.Storage[fileName] = content

	return "Upload Successful: " + fileName
}

func (s *S3) Content(fileName string) string {
	return s.Storage[fileName]
}
