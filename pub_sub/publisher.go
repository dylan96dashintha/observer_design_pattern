package pub_sub

import "sync"

type broker struct {
	subscribers map[string][]chan string
	lock        sync.RWMutex
}

func NewBroker() *broker {
	return &broker{
		subscribers: make(map[string][]chan string),
	}
}

// Subscribe to a topic
func (eb *broker) Subscribe(topic string) <-chan string {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	ch := make(chan string)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

// Publish a message to all subscribers of a topic
func (eb *broker) Publish(topic string, msg string) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	for _, ch := range eb.subscribers[topic] {
		go func(c chan string) {
			c <- msg
		}(ch)
	}
}
