package gcp

type CloudStorageClient interface {
	Upload(fileName string, content string) string
	Content(fileName string) string
}

type CloudStorage struct {
	Bucket map[string]string
}

func (cs *CloudStorage) Upload(fileName string, content string) string {
	cs.Bucket[fileName] = content

	return "GCP CloudStorage: Upload Successful, " + fileName
}

func (cs *CloudStorage) Content(fileName string) string {
	return "GCP CloudStorage: Content of " + fileName + "\n" + cs.Bucket[fileName]
}
