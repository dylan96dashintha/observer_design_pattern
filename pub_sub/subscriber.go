package pub_sub

import "fmt"

// Slack listens for alerts
func Slack(sub <-chan string) {
	for msg := range sub {
		fmt.Println("[Slack] spike, msg: ", msg)
	}
}

// GoogleSpace listens for alerts
func GoogleSpace(sub <-chan string) {
	for msg := range sub {
		fmt.Println("[GoogleSpace] spike, msg:", msg)
	}
}
