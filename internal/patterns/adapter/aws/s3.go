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

	return "AWS S3: Upload Successful, " + fileName
}

func (s *S3) Content(fileName string) string {
	return "AWS S3: Content of " + fileName + "\n" + s.Storage[fileName]
}
