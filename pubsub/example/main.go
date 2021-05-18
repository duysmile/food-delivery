package main

import (
	"200lab/food-delivery/pubsub"
	"200lab/food-delivery/pubsub/pblocal"
	"context"
	"log"
	"time"
)

func main() {
	var localPb pubsub.Pubsub = pblocal.NewPubSub()

	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPb.Subscribe(context.Background(), topic)
	sub2, _ := localPb.Subscribe(context.Background(), topic)

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPb.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Con1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Con2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()
	// close2()

	localPb.Publish(context.Background(), topic, pubsub.NewMessage(3))
	time.Sleep(time.Second * 2)
}
