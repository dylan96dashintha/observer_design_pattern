package pub_sub

// MonitoringAgent Subject interface
type MonitoringAgent interface {
	RegisterObserver(observer chan AlertingAgentChan)
	DeRegisterObserver(observer chan AlertingAgentChan)
	CheckMetrics(metric string, value float64)
}

type Grafana struct {
	AlertingAgentChannelList []chan AlertingAgentChan
	MemoryThreshold          float64
}

func (g *Grafana) RegisterObserver(observerChan chan AlertingAgentChan) {
	g.AlertingAgentChannelList = append(g.AlertingAgentChannelList, observerChan)
}

func (g *Grafana) DeRegisterObserver(alertChan chan AlertingAgentChan) {
	for i, observer := range g.AlertingAgentChannelList {
		if observer == alertChan {
			g.AlertingAgentChannelList = append(g.AlertingAgentChannelList[:i], g.AlertingAgentChannelList[i+1:]...)
			break
		}
	}
}

func (g *Grafana) CheckMetrics(metric string, value float64) {
	if g.MemoryThreshold < value {
	}
}
