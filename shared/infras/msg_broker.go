package sharedinfras

import (
	"context"
	
	"github.com/katatrina/go12-service/shared/datatype"
)

type IMsgBroker interface {
	Publish(ctx context.Context, topic string, evt *datatype.AppEvent) error
}
