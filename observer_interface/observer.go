package observer_interface

import "fmt"

// AlertingAgent Observer interface
type AlertingAgent interface {
	Notify(metric string, value float64)
}

type Slack struct{}

type GoogleSpace struct{}

func (s Slack) Notify(metric string, value float64) {
	fmt.Printf("[Slack]  %s spike: %.2f%%\n", metric, value)
}

func (g GoogleSpace) Notify(metric string, value float64) {
	fmt.Printf("[GoogleSpace]  %s spike: %.2f%%\n", metric, value)
}
