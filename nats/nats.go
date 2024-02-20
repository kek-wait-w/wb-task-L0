package nats

import (
	"encoding/json"
	"sync"
	"wb-l0/cache"
	"wb-l0/domain"
	logs "wb-l0/logger"

	"github.com/nats-io/stan.go"
)

func SubscribeToNATS(cache *cache.Cache) {
	var wg sync.WaitGroup
	wg.Add(1)

	sc, err := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		logs.Logger.Fatal(logs.Logger, err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("subject", func(msg *stan.Msg) {
		defer wg.Done()

		var order domain.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			logs.Logger.Debug(logs.Logger, err)
			return
		}

		cache.SetOrder(order.OrderUID, order)
		logs.Logger.Debug(logs.Logger, err)
	})
	if err != nil {
		logs.Logger.Fatal(logs.Logger, err)
	}

	wg.Wait()
}
