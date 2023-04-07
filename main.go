package main

import (
	"context"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/beslow/rocketmq_consumer/config"
	"github.com/beslow/rocketmq_consumer/logger"
)

func main() {
	sig := make(chan os.Signal)

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{config.NameServerAddr()})),
	)

	err := c.Subscribe("resume", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			logger.Infof("subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		logger.Errorf("subscribe fail, err: %v\n", err)
	}

	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		logger.Errorf("start push consumer fail, err: %v\n", err)
		os.Exit(-1)
	}

	<-sig

	err = c.Shutdown()
	if err != nil {
		logger.Errorf("shutdown Consumer error: %s", err.Error())
	}
}
