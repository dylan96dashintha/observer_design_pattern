package main

import (
	"github.com/observer_design_pattern/observer_interface"
	"github.com/observer_design_pattern/pub_sub"
	"time"
)

func main() {
	// observer design pattern
	slackObj := observer_interface.Slack{}
	googleSpaceObj := observer_interface.GoogleSpace{}
	grafanaServer := observer_interface.Grafana{
		MemoryThreshold: 50,
	}
	grafanaServer.RegisterObserver(&slackObj)
	grafanaServer.RegisterObserver(&googleSpaceObj)

	grafanaServer.CheckMetrics("high memory usage", 70)

	//pub sub design pattern
	brokerObj := pub_sub.NewBroker()
	alertChanSlack := brokerObj.Subscribe("CPU alert")
	alertChanGSpace := brokerObj.Subscribe("CPU alert")
	go pub_sub.Slack(alertChanSlack)
	go pub_sub.GoogleSpace(alertChanGSpace)

	go brokerObj.Publish("CPU alert", "high memory usage")
	time.Sleep(15 * time.Second)
}
