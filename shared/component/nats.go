package sharedcomponent

import (
	"context"
	"encoding/json"
	"log"
	
	"github.com/katatrina/go12-service/shared/datatype"
	
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type natsComp struct {
	nc *nats.Conn
}

func NewNatsComp(url string) *natsComp {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	
	return &natsComp{nc: nc}
}

func (c *natsComp) Publish(ctx context.Context, topic string, evt *datatype.AppEvent) error {
	dataByte, err := json.Marshal(evt.Data)
	
	if err != nil {
		return errors.WithStack(err)
	}
	
	return c.nc.Publish(topic, dataByte)
}
