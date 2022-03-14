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

type ObjectStorage struct {
    client *s3.Client // 对象存储服务
}

// NewObjectStorage 初始化对象存储
func NewObjectStorage(c *conf.Service) (*ObjectStorage, error) {

    // TODO 绑定 Log 未完成

    const defaultRegion = "us-east-1"

    staticResolver := aws.EndpointResolverWithOptionsFunc(
        func(service, region string, options ...interface{}) (aws.Endpoint, error) {
            return aws.Endpoint{
                PartitionID:   "aws",
                URL:           c.Data.ObjectStorage.Url, // minio url
                SigningRegion: defaultRegion,
                //HostnameImmutable: true,
            }, nil
        },
    )

    cfg, err := config.LoadDefaultConfig(
        context.TODO(),
        func(lo *config.LoadOptions) error {
            lo.Region = defaultRegion
            lo.Credentials = credentials.NewStaticCredentialsProvider(
                c.Data.ObjectStorage.AccessKey, // Access Key
                c.Data.ObjectStorage.SecretKey, // Secret Key
                "",
            )
            lo.EndpointResolverWithOptions = staticResolver
            return nil
        },
    )
    if err != nil {
        return nil, err
    }

    client := s3.NewFromConfig(cfg)

    oss := &ObjectStorage{
        client: client,
    }

    return oss, nil
}

// GetClient 获取对象存储客户端
func (o *ObjectStorage) GetClient() *s3.Client {
    return o.client
}

// PreSignGetURL 获取文件对象的预签名URL
// bucket: 桶
// key: 对象路径
// expirationTime: 过期时间
func (o *ObjectStorage) PreSignGetURL(ctx context.Context, bucket, key string, expirationTime time.Duration) (string, error) {
    input := &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }

    psClient := s3.NewPresignClient(o.client)
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

// PreSignPutURL 上传文件对象的预签名URL
// bucket: 桶
// key: 对象路径
// expirationTime: 过期时间
func (o *ObjectStorage) PreSignPutURL(ctx context.Context, bucket, key string, expirationTime time.Duration) (string, error) {
    input := &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }

    psClient := s3.NewPresignClient(o.client)
    req, err := psClient.PresignPutObject(ctx, input,
        func(options *s3.PresignOptions) {
            options.Expires = expirationTime // 设置url过期时间
        },
    )
    if err != nil {
        return "", err
    }

    return req.URL, nil
}
