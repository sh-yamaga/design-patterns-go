package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/adapter"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/adapter/aws"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/adapter/gcp"
)

var s3Adapter aws.S3Adapter
var csAdapter gcp.CloudStorageAdapter

func init() {
	s3 := aws.S3{Bucket: make(map[string]string)}
	s3Adapter = aws.S3Adapter{S3Client: &s3}

	cs := gcp.CloudStorage{Bucket: make(map[string]string)}
	csAdapter = gcp.CloudStorageAdapter{CloudStorageClient: &cs}
}

func main() {
	var sc adapter.StorageClient

	// aws
	sc = &s3Adapter
	fmt.Println(sc.UploadFile("example.txt", "example content s3"))
	//　Output:
	//　AWS S3: Upload Successful, example.txt

	fmt.Println(sc.GetContent("example.txt"))
	//　Output:
	//　AWS S3: Content of example.txt
	//　example content s3

	// gcp
	sc = &csAdapter
	fmt.Println(sc.UploadFile("example.txt", "example content CloudStorage"))
	//　Output:
	//　GCP CloudStorage: Upload Successful, example.txt

	fmt.Println(sc.GetContent("example.txt"))
	//　Output:
	//　GCP CloudStorage: Content of example.txt
	//　example content CloudStorage
}
