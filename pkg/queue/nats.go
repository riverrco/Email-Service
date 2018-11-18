package queue

import (
	"github.com/nats-io/go-nats-streaming"
	"time"
)

type NatsClient struct {
	client stan.Conn
}

func NewNats(client stan.Conn) *NatsClient{
	return &NatsClient{
		client: client,
	}
}

func (n *NatsClient) Sub(ch string, q string, listener func(msg *stan.Msg) ) error {
	_, err := n.client.QueueSubscribe(
		ch,
		q,
		listener,
		stan.DurableName("notification-service"),
		stan.SetManualAckMode(),
		stan.AckWait(5 * time.Second),
	)

	if err != nil {
		return err
	}
	return nil
}

