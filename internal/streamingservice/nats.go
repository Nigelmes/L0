package streamingservice

import (
	"encoding/json"
	"github.com/Nigelmes/L0/internal/config"
	"github.com/Nigelmes/L0/internal/models"
	"github.com/Nigelmes/L0/internal/repository"
	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	durable = "backend"
	subject = "wb"
)

type NatsStream struct {
	client stan.Conn
}

func NewNatsStream(cfg *config.Config) *NatsStream {
	sc, err := stan.Connect(
		cfg.NatsStreamingCfg.StanClusterId, cfg.NatsStreamingCfg.ClientId,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		logrus.Fatalf("error connection nats-streaming, %s", err.Error())
	}
	logrus.Println("nats-streaming connection successful")
	return &NatsStream{client: sc}
}

func (n *NatsStream) RunNatsSteaming(repo *repository.Repository) {
	_, err := n.client.Subscribe(
		subject, func(m *stan.Msg) {
			var order models.Order
			err := json.Unmarshal(m.Data, &order)
			if err != nil {
				logrus.Error(err)
				return
			}
			err = repo.OrderRepo.Create(order)
			if err != nil {
				logrus.Error(err)
				return
			}
			repo.CacheRepo.Set(order)
		}, stan.StartAtTimeDelta(time.Minute*10),
		stan.DurableName(durable),
	)
	if err != nil {
		logrus.Error(err)
	}
}

func (n *NatsStream) ShutDown() error {
	return n.client.Close()
}
