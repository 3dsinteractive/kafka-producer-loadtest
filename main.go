// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"time"
)

func main() {
	ms := NewMicroservice()

	servers := "localhost:9094"
	topic := "loadtest-" + randString()
	// groupID := "loadtest-consumer"
	// timeout := time.Duration(-1)

	// ms.Consume(servers, topic, groupID, timeout, func(ctx IContext) error {
	// 	msg := ctx.ReadInput()
	// 	ctx.Log(msg)
	// 	return nil
	// })

	prod := NewProducer(servers, ms)
	i := 0
	go func() {
		t := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-t.C:
				ms.Stop()
				return
			default:
				i++
				prod.SendMessage(topic, "", map[string]interface{}{"message_id": i})
				print("%d", i)
			}
		}
	}()

	defer ms.Cleanup()
	ms.Start()
}
