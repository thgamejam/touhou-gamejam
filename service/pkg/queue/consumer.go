package queue

import (
    "github.com/apache/rocketmq-client-go/v2"
    "github.com/apache/rocketmq-client-go/v2/consumer"
    "github.com/apache/rocketmq-client-go/v2/primitive"
    "service/pkg/conf"
)

func NewConsumer(c *conf.Queue) (rocketmq.PushConsumer, error) {
    groupName := "DEFAULT"
    if c.GroupName != "" {
        groupName = c.GroupName
    }
    p, err := rocketmq.NewPushConsumer(
        consumer.WithGroupName(groupName),
        consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{c.NameServerAddr})),
        )
    if err != nil {
        return nil, err
    }
    return p, nil
}