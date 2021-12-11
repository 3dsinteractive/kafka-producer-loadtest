// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"time"
)

func main() {
	ms := NewMicroservice()

	servers := "localhost:9094"
	topic := "loadtest-" + randString()

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
				prod := NewProducer(servers, ms)
				err := prod.SendMessage(topic, "", map[string]interface{}{"message_id": i})
				if err != nil {
					logE(err)
					ms.Stop()
				}
				print("%d", i)
			}
		}
	}()

	defer ms.Cleanup()
	ms.Start()
}
