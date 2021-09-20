package s3

import (
	"github.com/go-basic/uuid"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type S3File struct {
	AccessKey string
	SecretKey string
	EndPoint  string
	Bucket    string
	Item      string
}

func Download(s3 S3File) (string, error) {
	minioCli, err := minio.New(s3.EndPoint, s3.AccessKey, s3.SecretKey, false)
	if err != nil {
		logrus.Error("minio 连接创建失败！", err)
		return "", err
	}

	reader, err := minioCli.GetObject(s3.Bucket, s3.Item, minio.GetObjectOptions{})
	if err != nil {
		logrus.Error("读取远端 minio 文件失败！", err)
		return "", err
	}
	filePath := "/tmp/" + uuid.New() + ".tar"

	file, err := os.Create(filePath)
	if err != nil {
		logrus.Error("创建本地文件失败！", err)
		return "", err
	}
	defer file.Close()

	stat, err := reader.Stat()
	if err != nil {
		logrus.Error("获取远端文件信息失败！", err)
		return "", err
	}

	if _, err := io.CopyN(file, reader, stat.Size); err != nil {
		logrus.Error("下载文件失败！", err)
		return "", err
	}
	return filePath, nil
}
