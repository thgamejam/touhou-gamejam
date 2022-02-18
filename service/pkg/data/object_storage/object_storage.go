package object_storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"service/pkg/conf"
	"time"
)

type Client struct {
	S3Client *s3.Client // 对象存储服务

	bucket                  *string
	smallFileExpirationTime time.Duration //小文件到期时间
	largeFileExpirationTime time.Duration //大文件到期时间
}

// NewObjectStorage 初始化对象存储
func NewObjectStorage(c *conf.Data_ObjectStorage) (*Client, error) {

	// TODO 绑定 Log 未完成

	const defaultRegion = "us-east-1"

	staticResolver := aws.EndpointResolverFunc(
		func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               c.Url, // minio url
				SigningRegion:     defaultRegion,
				HostnameImmutable: true,
			}, nil
		},
	)

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		func(lo *config.LoadOptions) error {
			lo.Region = defaultRegion
			lo.Credentials = credentials.NewStaticCredentialsProvider(
				c.AccessKey, // Access Key
				c.SecretKey, // Secret Key
				"",
			)
			lo.EndpointResolver = staticResolver
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	oss := &Client{
		S3Client:                client,
		smallFileExpirationTime: c.SmallFileExpirationTime.AsDuration(),
		largeFileExpirationTime: c.LargeFileExpirationTime.AsDuration(),
	}

	oss.bucket = aws.String(c.Bucket)

	return oss, nil
}

// GetURL 获取文件对象的预签名URL
// key: 对象路径
// expirationTime: 过期时间
func (o *Client) GetURL(ctx context.Context, key string, expirationTime time.Duration) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: o.bucket,
		Key:    aws.String(key),
	}

	psClient := s3.NewPresignClient(o.S3Client)
	req, err := psClient.PresignGetObject(ctx, input,
		func(options *s3.PresignOptions) {
			options.Expires = expirationTime // 设置url过期时间
		},
	)
	if err != nil {
		return "", err
	}

	return req.URL, nil
}

// GetSmallObjectURL 获取小对象的URL
func (o *Client) GetSmallObjectURL(ctx context.Context, key string) (string, error) {
	return o.GetURL(ctx, key, o.smallFileExpirationTime)
}

// GetLargeObjectURL 获取大对象的URL
func (o *Client) GetLargeObjectURL(ctx context.Context, key string) (string, error) {
	return o.GetURL(ctx, key, o.largeFileExpirationTime)
}
