package queue

import (
    "github.com/apache/rocketmq-client-go/v2"
    "github.com/apache/rocketmq-client-go/v2/primitive"
    "github.com/apache/rocketmq-client-go/v2/producer"
    "service/pkg/conf"
)

// NewProducer 实例化生产者
func NewProducer(c *conf.Queue) (rocketmq.Producer, error) {
    groupName := "DEFAULT"
    if c.GroupName != "" {
        groupName = c.GroupName
    }
    p, err := rocketmq.NewProducer(
        producer.WithNsResolver(primitive.NewPassthroughResolver([]string{c.NameServerAddr})),
        producer.WithGroupName(groupName),
        producer.WithRetry(int(c.Retry)),
        )
    if err != nil {
        return nil, err
    }
    err = p.Start()
    if err != nil {
        return nil, err
    }
    return p, nil
}
