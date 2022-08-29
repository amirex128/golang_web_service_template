package rabbit

import (
	"backend/internal/pkg/framework/assert"
	"container/ring"
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type Channel interface {
	/*
		Confirm puts this channel into confirm mode so that the client can ensure all
		publishing's have successfully been received by the server. After entering this
		mode, the server will send a basic.ack or basic.nack message with the deliver
		tag set to a 1 based incrementing index corresponding to every publishing
		received after the this method returns.

		Add a listener to Channel.NotifyPublish to respond to the Confirmations. If
		Channel.NotifyPublish is not called, the Confirmations will be silently
		ignored.

		The order of acknowledgments is not bound to the order of deliveries.

		Ack and Nack confirmations will arrive at some point in the future.

		Unroutable mandatory or immediate messages are acknowledged immediately after
		any Channel.NotifyReturn listeners have been notified. Other messages are
		acknowledged when all queues that should have the message routed to them have
		either have received acknowledgment of delivery or have enqueued the message,
		persisting the message if necessary.

		When noWait is true, the client will not wait for a response. A channel
		exception could occur if the server does not support this method.

	*/
	Confirm(noWait bool) error

	/*
	   NotifyPublish registers a listener for reliable publishing. Receives from this
	   chan for every publish after Channel.Confirm will be in order starting with
	   DeliveryTag 1.

	   There will be one and only one Confirmation Publishing starting with the
	   delivery tag of 1 and progressing sequentially until the total number of
	   publishing's have been seen by the server.

	   Acknowledgments will be received in the order of delivery from the
	   NotifyPublish channels even if the server acknowledges them out of order.

	   The listener chan will be closed when the Channel is closed.

	   The capacity of the chan Confirmation must be at least as large as the
	   number of outstanding publishing's. Not having enough buffered chans will
	   create a deadlock if you attempt to perform other operations on the Connection
	   or Channel while confirms are in-flight.

	   It's advisable to wait for all Confirmations to arrive before calling
	   Channel.Close() or Connection.Close().

	*/
	NotifyPublish(confirm chan amqp.Confirmation) chan amqp.Confirmation

	/*
	   Publish sends a Publishing from the client to an exchange on the server.

	   When you want a single message to be delivered to a single queue, you can
	   publish to the default exchange with the routingKey of the queue name. This is
	   because every declared queue gets an implicit route to the default exchange.

	   Since publishing's are asynchronous, any undeliverable message will get returned
	   by the server. Add a listener with Channel.NotifyReturn to handle any
	   undeliverable message when calling publish with either the mandatory or
	   immediate parameters as true.

	   publishing's can be undeliverable when the mandatory flag is true and no queue is
	   bound that matches the routing key, or when the immediate flag is true and no
	   consumer on the matched queue is ready to accept the delivery.

	   This can return an error when the channel, connection or socket is closed. The
	   error or lack of an error does not indicate whether the server has received this
	   publishing.

	   It is possible for publishing to not reach the broker if the underlying socket
	   is shutdown without pending publishing packets being flushed from the kernel
	   buffers. The easy way of making it probable that all publishing's reach the
	   server is to always call Connection.Close before terminating your publishing
	   application. The way to ensure that all publishing's reach the server is to add
	   a listener to Channel.NotifyPublish and put the channel in confirm mode with
	   Channel.Confirm. Publishing delivery tags and their corresponding
	   confirmations start at 1. Exit when all publishing's are confirmed.

	   When Publish does not return an error and the channel is in confirm mode, the
	   internal counter for DeliveryTags with the first confirmation starting at 1.

	*/
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error

	/*
	   Close initiate a clean channel closure by sending a close message with the error
	   code set to '200'.

	   It is safe to call this method multiple times.

	*/
	Close() error
}

var (
	connRng            = make(map[string]*ring.Ring, 0)
	connRngLock        = &sync.RWMutex{}
	once               = sync.Once{}
	rng                = make(map[string]*ring.Ring, 0)
	rngLock            = &sync.RWMutex{}
	kill               context.Context
	rabbitConnExpected = make([]rabbitExpected, 0)
)

var notifyClose = make(chan *amqp.Error, 10)

type ignite struct {
}

func (in *ignite) Health(ctx context.Context) error {
	select {
	case err := <-notifyClose:
		if err != nil {
			return fmt.Errorf("RabbitMQ error happen : %s", err)
		}
	default: // Do not block
	}
	return nil
}

type chnlLock struct {
	chn    Channel
	lock   *sync.Mutex
	rtrn   chan amqp.Confirmation
	wg     *sync.WaitGroup
	closed bool
}

func Initialize(ctx context.Context) {
	once.Do(func() {

		for i := range rabbitConnExpected {
			kill, _ = context.WithCancel(ctx)
			cnt := viper.GetInt("rabbit_connection_num")
			if cnt < 1 {
				cnt = 1
			}
			connString := fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
				rabbitConnExpected[i].user,
				rabbitConnExpected[i].password,
				rabbitConnExpected[i].host,
				rabbitConnExpected[i].port,
				rabbitConnExpected[i].vHost,
			)
			connRngLock.Lock()
			defer connRngLock.Unlock()
			rngLock.Lock()
			defer rngLock.Unlock()
			connRng[rabbitConnExpected[i].containerName] = ring.New(cnt)
			for j := 0; j < cnt; j++ {
				c, err := amqp.Dial(connString)
				if err == nil {
					connRng[rabbitConnExpected[i].containerName].Value = c
					connRng[rabbitConnExpected[i].containerName] = connRng[rabbitConnExpected[i].containerName].Next()
				} else {
					logrus.Errorf("error while connect to rabbit : %s", err.Error())
					return
				}
			}
			connRng[rabbitConnExpected[i].containerName] = connRng[rabbitConnExpected[i].containerName].Next()

			conn := connRng[rabbitConnExpected[i].containerName].Value.(*amqp.Connection)

			chn, err := conn.Channel()
			assert.Nil(err)
			defer func() {
				assert.Nil(chn.Close())
			}()
			assert.Nil(
				chn.ExchangeDeclare(
					viper.GetString("exchange_name"),
					viper.GetString("exchange_type"),
					true,
					false,
					false,
					false,
					amqp.Table{},
				),
			)
			//queue, err := chn.QueueDeclare(
			//	"shit", // name of the queue
			//	true,   // durable
			//	false,  // delete when unused
			//	false,  // exclusive
			//	false,  // noWait
			//	nil,    // arguments
			//)
			//if err != nil {
			//	assert.Nil(err)
			//}

			//if err = chn.QueueBind(
			//	queue.Name,                       // name of the queue
			//	queue.Name,                       // bindingKey
			//	viper.GetString("EXCHANGE_NAME"), // sourceExchange
			//	false,                            // noWait
			//	nil,                              // arguments
			//); err != nil {
			//	assert.Nil(err)
			//}

			rng[rabbitConnExpected[i].containerName] = ring.New(viper.GetInt("rabbit_publish_num"))
			for j := 0; j < viper.GetInt("rabbit_publish_num"); j++ {
				connRng[rabbitConnExpected[i].containerName] = connRng[rabbitConnExpected[i].containerName].Next()
				conn := connRng[rabbitConnExpected[i].containerName].Value.(*amqp.Connection)
				pchn, err := conn.Channel()
				assert.Nil(err)
				rtrn := make(chan amqp.Confirmation, viper.GetInt("rabbit_confirm_len"))
				err = pchn.Confirm(false)
				assert.Nil(err)
				pchn.NotifyPublish(rtrn)
				tmp := chnlLock{
					chn:    pchn,
					lock:   &sync.Mutex{},
					wg:     &sync.WaitGroup{},
					rtrn:   rtrn,
					closed: false,
				}
				go publishConfirm(&tmp)
				rng[rabbitConnExpected[i].containerName].Value = &tmp
				rng[rabbitConnExpected[i].containerName] = rng[rabbitConnExpected[i].containerName].Next()
			}

			logrus.Info("Rabbit initialized")

		}

	})
}

func publishConfirm(cl *chnlLock) {
	for range cl.rtrn {
		cl.wg.Done()
	}
}

type rabbitExpected struct {
	containerName string
	host          string
	port          int
	user          string
	password      string
	vHost         string
}

func RegisterRabbit(cnt, host, user, password, vHost string, port int) {
	rabbitConnExpected = append(rabbitConnExpected, rabbitExpected{
		containerName: cnt,
		host:          host,
		vHost:         vHost,
		user:          user,
		password:      password,
		port:          port,
	})
}
