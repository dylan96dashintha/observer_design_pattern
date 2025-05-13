package pub_sub

import "fmt"

// AlertingAgentChan Observer interface
type AlertingAgentChan interface {
	NotifyChan(metric string, value float64)
}

type Slack struct{}

type GoogleSpace struct{}

func (s Slack) I(metric string, value float64) {
	fmt.Printf("[Slack]  %s spike: %.2f%%\n", metric, value)
}

func (g GoogleSpace) NotifyChan(metric string, value float64) {
	fmt.Printf("[GoogleSpace]  %s spike: %.2f%%\n", metric, value)
}
