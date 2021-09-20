package s3

import (
	"github.com/123shang60/image-load/pkg/common"
	"io"
	"os"

	"github.com/go-basic/uuid"
	"github.com/minio/minio-go"
)

type S3File struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	EndPoint  string `json:"end_point"`
	Bucket    string `json:"bucket"`
	Item      string `json:"item"`
}

func Download(s3 S3File) (string, error) {
	minioCli, err := minio.New(s3.EndPoint, s3.AccessKey, s3.SecretKey, false)
	if err != nil {
		common.Logger().Error("minio 连接创建失败！", err)
		return "", err
	}

	reader, err := minioCli.GetObject(s3.Bucket, s3.Item, minio.GetObjectOptions{})
	if err != nil {
		common.Logger().Error("读取远端 minio 文件失败！", err)
		return "", err
	}
	filePath := "/tmp/" + uuid.New() + ".tar"

	file, err := os.Create(filePath)
	if err != nil {
		common.Logger().Error("创建本地文件失败！", err)
		return "", err
	}
	defer file.Close()

	stat, err := reader.Stat()
	if err != nil {
		common.Logger().Error("获取远端文件信息失败！", err)
		return "", err
	}

	if _, err := io.CopyN(file, reader, stat.Size); err != nil {
		common.Logger().Error("下载文件失败！", err)
		return "", err
	}
	return filePath, nil
}
