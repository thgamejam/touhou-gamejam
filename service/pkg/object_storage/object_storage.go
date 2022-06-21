package object_storage

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"service/pkg/conf"
)

type ObjectStorage struct {
	client *minio.Client // 对象存储服务
}

// NewObjectStorage 初始化对象存储
func NewObjectStorage(c *conf.Service) (*ObjectStorage, error) {
	client, err := minio.New(
		c.Data.ObjectStorage.Domain, // 使用的域名
		&minio.Options{
			Creds: credentials.NewStaticV4(
				c.Data.ObjectStorage.AccessKeyId,
				c.Data.ObjectStorage.SecretAccessKey,
				c.Data.ObjectStorage.Token,
			),
			Secure: c.Data.ObjectStorage.Secure,
		})
	if err != nil {
		return nil, err
	}

	// TODO 绑定 Log 未完成

	oss := &ObjectStorage{
		client: client,
	}

	return oss, nil
}

// GetClient 获取对象存储客户端
func (o *ObjectStorage) GetClient() *minio.Client {
	return o.client
}

// CreateBucket 创建桶
func (o *ObjectStorage) CreateBucket(ctx context.Context, bucketName string) (err error) {
	return o.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "cn-north-1"})
}

// ExistBucket 获取对应桶是否存在
func (o *ObjectStorage) ExistBucket(ctx context.Context, bucketName string) (ok bool, err error) {
	return o.client.BucketExists(ctx, bucketName)
}

// PreSignGetURL 获取文件对象的预签名URL
// bucket: 桶
// key: 对象路径
// filename: 下载时的文件名
// expirationTime: 过期时间
func (o *ObjectStorage) PreSignGetURL(
	ctx context.Context, bucket, key, filename string, expirationTime time.Duration,
) (*url.URL, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%v\"", filename))
	preSignedURL, err := o.client.PresignedGetObject(ctx, bucket, key, expirationTime, reqParams)
	if err != nil {
		return nil, err
	}

	return preSignedURL, nil
}

// PreSignPutURL 上传文件对象的预签名URL
// bucket: 桶
// key: 对象路径
// expirationTime: 过期时间
func (o *ObjectStorage) PreSignPutURL(
	ctx context.Context, bucket, key string, expirationTime time.Duration,
) (*url.URL, error) {
	preSignedURL, err := o.client.PresignedPutObject(ctx, bucket, key, expirationTime)
	if err != nil {
		return nil, err
	}

	return preSignedURL, nil
}

func (o *ObjectStorage) PutObject(ctx context.Context, bucket, key string, file *os.File) error {
	objectStat, err := file.Stat()
	if err != nil {
		return err
	}
	_, err = o.client.PutObject(ctx, bucket, key, file, objectStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}
