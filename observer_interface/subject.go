package observer_interface

// MonitoringAgent Subject interface
type MonitoringAgent interface {
	RegisterObserver(observer AlertingAgent)
	DeRegisterObserver(observer AlertingAgent)
	CheckMetrics(metric string, value float64)
}

type Grafana struct {
	AlertingAgent   []AlertingAgent
	MemoryThreshold float64
}

func (g *Grafana) RegisterObserver(observer AlertingAgent) {
	g.AlertingAgent = append(g.AlertingAgent, observer)
}

func (g *Grafana) DeRegisterObserver(alert AlertingAgent) {
	for i, observer := range g.AlertingAgent {
		if observer == alert {
			g.AlertingAgent = append(g.AlertingAgent[:i], g.AlertingAgent[i+1:]...)
			break
		}
	}
}

func (g *Grafana) CheckMetrics(metric string, value float64) {
	if g.MemoryThreshold < value {
		for _, observer := range g.AlertingAgent {
			observer.Notify(metric, value)
		}
	}
}
