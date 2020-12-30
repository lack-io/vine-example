package main

import (
	"context"

	vine "github.com/lack-io/vine/service"
	log "github.com/lack-io/vine/service/logger"
	"github.com/lack-io/vine/service/server"
	"github.com/lack-io/vine/util/context/metadata"

	proto "github.com/lack-io/vine-example/pubsub/proto"
)

// All methods of Sub will be executed when
// a message is received
type Sub struct{}

// Method can be of any name
func (s *Sub) Process(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("[pubsub.1] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

// Alternatively a function can be used
func subEv(ctx context.Context, event *proto.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Infof("[pubsub.2] Received event %+v with metadata %+v\n", event, md)
	// do something with event
	return nil
}

func main() {
	// create a service
	service := vine.NewService(
		vine.Name("go.vine.srv.pubsub"),
	)
	// parse command line
	service.Init()

	// register subscriber
	vine.RegisterSubscriber("example.topic.pubsub.1", service.Server(), new(Sub))

	// register subscriber with queue, each message is delivered to a unique subscriber
	vine.RegisterSubscriber("example.topic.pubsub.2", service.Server(), subEv, server.SubscriberQueue("queue.pubsub"))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
