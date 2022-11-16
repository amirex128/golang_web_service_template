package rabbit

import (
	"context"
	"encoding/json"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/assert"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/random"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/safe"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"os"
	"time"
)

type jsonDelivery struct {
	delivery *amqp.Delivery
}

func (jd jsonDelivery) Decode(v interface{}) error {
	err := json.Unmarshal(jd.delivery.Body, v)
	if err != nil {
		logrus.Debugf("Convert %s ====> %T , err was %s", string(jd.delivery.Body), v, err.Error())
	}
	return err
}

func (jd jsonDelivery) Ack(multiple bool) error {
	return jd.delivery.Ack(multiple)
}

func (jd jsonDelivery) Nack(multiple, requeue bool) error {
	return jd.delivery.Nack(multiple, requeue)
}

func (jd jsonDelivery) Reject(requeue bool) error {
	return jd.delivery.Reject(requeue)
}

// Delivery is the job to consumer
type Delivery interface {
	Decode(v interface{}) error
	// Ack delegates an acknowledgement through the Acknowledger interface that the client or server has finished work on a delivery.
	Ack(multiple bool) error
	// Nack negatively acknowledge the delivery of message(s) identified by the delivery tag from either the client or server.
	Nack(multiple, requeue bool) error
	// Reject delegates a negatively acknowledgement through the Acknowledger interface.
	Reject(requeue bool) error
}

// Consumer is the side the workers on it
type Consumer interface {
	// Topic return the topic that this worker want to listen to it
	Topic() string
	// Queue is the queue that this want to listen to
	Queue() string
	// Consume return a channel to put jobs into
	Consume(context.Context) chan<- Delivery
}

func RegisterConsumer(consumer Consumer, cnt string) error {
	connRngLock.Lock()
	connRng[cnt] = connRng[cnt].Next()
	conn := connRng[cnt].Value.(*amqp.Connection)
	connRngLock.Unlock()
	c, err := conn.Channel()
	if err != nil {
		return err
	}
	//err = c.ExchangeDeclare(
	//	viper.GetString("exchange_name"), // name
	//	viper.GetString("exchange_type"), // type
	//	true,                             // durable
	//	false,                            // auto-deleted
	//	false,                            // internal
	//	false,                            // no-wait
	//	nil,                              // arguments
	//)
	//
	//if err != nil {
	//	return err
	//}
	qu := consumer.Queue()
	q, err := c.QueueDeclare(qu, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// prefetch count
	// **WARNING**
	// If ignore this, then there is a problem with rabbit. prefetch all jobs for this worker then.
	// the next worker get nothing at all!
	// **WARNING**
	// TODO : limit on workers must match with this prefetch
	err = c.Qos(viper.GetInt("prefetch_count"), 0, false)
	if err != nil {
		return err
	}

	//topic := consumer.Topic()

	err = c.QueueBind(
		q.Name,                           // queue name
		"",                               // routing key
		viper.GetString("exchange_name"), // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}
	safe.ContinuesGoRoutine(kill, func(cnl context.CancelFunc) time.Duration {
		consumerTag := <-random.ID
		delivery, err := c.Consume(q.Name, consumerTag, false, false, false, false, nil)
		if err != nil {
			if val, ok := err.(*amqp.Error); ok {
				if val.Code == amqp.ChannelError {
					notifyClose <- amqp.ErrClosed
					cnl()
					os.Exit(1)
				}
			}
			cnl()
			return 0
		}
		consume(kill, cnl, consumer.Consume(kill), c, delivery, consumerTag)
		return time.Second
	})
	return nil
}

func consume(ctx context.Context, cnl context.CancelFunc, consumer chan<- Delivery, c *amqp.Channel, delivery <-chan amqp.Delivery, consumerTag string) {
	done := ctx.Done()

	cErr := c.NotifyClose(make(chan *amqp.Error))
bigLoop:
	for {
		select {
		case job, ok := <-delivery:
			assert.True(ok, "[BUG] Channel is closed! why??")
			consumer <- &jsonDelivery{delivery: &job}
		case <-done:
			logrus.Debug("closing channel")
			// break the continues loop
			cnl()
			_ = c.Cancel(consumerTag, true)
			break bigLoop
		case e := <-cErr:
			logrus.Errorf("%T => %+v", *e, *e)
		}
	}
}
