package queue

import (
    "context"
    "github.com/apache/rocketmq-client-go/v2"
    "github.com/apache/rocketmq-client-go/v2/primitive"
    "github.com/apache/rocketmq-client-go/v2/producer"
    "service/pkg/conf"
)

type Producer struct {
    producer rocketmq.Producer
}

// NewProducer 实例化生产者
func NewProducer(c *conf.Service) (*Producer, error) {
    groupName := "DEFAULT"
    if c.Data.Queue.GroupName != "" {
        groupName = c.Data.Queue.GroupName
    }
    p, err := rocketmq.NewProducer(
        producer.WithNsResolver(primitive.NewPassthroughResolver([]string{c.Data.Queue.NameServerAddr})),
        producer.WithGroupName(groupName),
        producer.WithRetry(int(c.Data.Queue.Retry)),
        )
    if err != nil {
        return nil, err
    }
    return &Producer{producer: p}, nil
}


// SendMessage 发送消息
func (p *Producer) SendMessage(ctx context.Context, m *primitive.Message) (res *primitive.SendResult,e error) {
    e = p.producer.Start()
    if e != nil {
        return
    }
    res, e = p.producer.SendSync(ctx, m)
    return
}

// Close 关闭连接
func (p *Producer) Close() (e error) {
    e = p.producer.Shutdown()
    return
}
