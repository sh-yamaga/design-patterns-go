package aws

type S3Client interface {
	Upload(fileName string, content string) string
	Content(fileName string) string
}

type S3 struct {
	Bucket map[string]string
}

func (s *S3) Upload(fileName string, content string) string {
	s.Bucket[fileName] = content

	return "AWS S3: Upload Successful, " + fileName
}

func (s *S3) Content(fileName string) string {
	return "AWS S3: Content of " + fileName + "\n" + s.Bucket[fileName]
}
