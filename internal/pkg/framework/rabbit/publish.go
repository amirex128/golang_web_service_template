package rabbit

import (
	"github.com/amirex128/selloora_backend/internal/pkg/framework/random"
	"errors"
	"os"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func Publish(in Job, cnt string) error {
	var err error
	rngLock.Lock()
	rng[cnt] = rng[cnt].Next()
	v := rng[cnt].Value.(*chnlLock)
	rngLock.Unlock()
	v.lock.Lock()
	defer v.lock.Unlock()
	if v.closed {
		return errors.New("waiting for finalize, can not publish")
	}

	msg, err := in.Encode()
	if err != nil {
		return err
	}

	pub := amqp.Publishing{
		CorrelationId: <-random.ID,
		Body:          msg,
	}

	v.wg.Add(1)
	defer func() {
		// If the result is error, release the lock, there is no message to confirm!
		if err != nil {
			v.wg.Done()
		}
	}()
	topic := in.Topic()
	err = v.chn.Publish(viper.GetString("exchange_name"), topic, true, false, pub)
	if err != nil {
		if val, ok := err.(*amqp.Error); ok {
			if val.Code == amqp.ChannelError {
				notifyClose <- amqp.ErrClosed
				os.Exit(1)
			}
		}
	}
	return err
}
