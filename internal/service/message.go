package service

import (
	"strings"

	"github.com/helderfarias/go-config-server/internal/domain"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type MessageBroker interface {
	Publish(msg string) error
}

type natsMessageBroker struct {
	cfg domain.SpringCloudConfig
}

type emptyMessageBroker struct {
}

func newFactoryMessageBroker(cfg domain.SpringCloudConfig) MessageBroker {
	if strings.TrimSpace(cfg.Spring.Nats.Servers) != "" {
		return &natsMessageBroker{cfg: cfg}
	}

	return &emptyMessageBroker{}
}

func (n *natsMessageBroker) Publish(msg string) error {
	opts := []nats.Option{nats.Name("GoConfigServer: NATS Publisher")}
	if n.cfg.Spring.Nats.Auth.Type == "token" {
		opts = append(opts, nats.Token(n.cfg.Spring.Nats.Auth.Token))
	} else if n.cfg.Spring.Nats.Auth.Type == "userinfo" {
		opts = append(opts, nats.UserInfo(n.cfg.Spring.Nats.Auth.User, n.cfg.Spring.Nats.Auth.Password))
	}

	nc, err := nats.Connect(n.cfg.Spring.Nats.Servers, opts...)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer nc.Close()

	nc.Publish(n.cfg.Spring.Nats.Subject, []byte(msg))
	nc.Flush()

	if err := nc.LastError(); err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Published")
	return nil
}

func (e *emptyMessageBroker) Publish(msg string) error {
	return nil
}
