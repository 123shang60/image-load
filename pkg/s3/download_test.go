package s3

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	s3 := S3File{
		AccessKey: "admin",
		SecretKey: "123456789",
		EndPoint:  "192.168.31.10:9000",
		Bucket:    "test",
		Item:      "alpine.tar",
	}

	filePath, err := Download(s3)
	if err != nil {
		panic(err)
	}
	fmt.Println(filePath)
}
