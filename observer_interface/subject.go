package observer_interface

// MonitoringAgent Subject interface
type MonitoringAgent interface {
	RegisterObserver(observer AlertingAgent)
	DeRegisterObserver(observer AlertingAgent)
	CheckMetrics()
}

// AlertingAgent Observer interface
type AlertingAgent interface {
	Notify(metric string, value float64)
}

type Grafana struct {
	AlertingAgent   []AlertingAgent
	memoryThreshold float64
}

func (g Grafana) RegisterObserver(observer AlertingAgent) {
	g.AlertingAgent = append(g.AlertingAgent, observer)
}

func (g Grafana) DeRegisterObserver(alert AlertingAgent) {
	for _, observer := range g.AlertingAgent {
		if observer == alert {
			
		}
	}
}
