package consumer

import (
	"context"
	"encoding/json"
	"github.com/murat96k/kitaptar.kz/internal/user/handler/consumer/dto"
	"github.com/murat96k/kitaptar.kz/internal/user/service"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type UserVerificationCallback struct {
	service service.Manager
}

func NewUserVerificationCallback(service service.Manager) *UserVerificationCallback {
	return &UserVerificationCallback{service: service}
}

func (c *UserVerificationCallback) Callback(message <-chan *sarama.ConsumerMessage, errors <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				log.Printf("failed to unmarshall record value err: %v", err)
			} else {
				log.Printf("user code: %s", userCode.Code)
				//nolint
				go c.service.SendConfirmCode(userCode.Email, userCode.Code)

				err = c.service.Cache.CodeCache.SetCode(context.Background(), userCode.Email, userCode.Code, time.Minute*5)
				if err != nil {
					log.Printf("Redis set code cache error: %v", err)
					return
				}
			}
		case err := <-errors:
			log.Printf("failed consume err: %v", err)
		}
	}
}
