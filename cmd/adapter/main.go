package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/adapter"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/adapter/aws"
)

var s3Adapter aws.S3Adapter

func init() {
	s3 := aws.S3{Storage: make(map[string]string)}
	s3Adapter = aws.S3Adapter{S3Client: &s3}
}

func main() {
	var cs adapter.CloudStorage = &s3Adapter

	upResult := cs.UploadFile("example.txt", "example content s3")
	fmt.Println(upResult)

	content := cs.GetContent("example.txt")
	fmt.Println(content)
}
