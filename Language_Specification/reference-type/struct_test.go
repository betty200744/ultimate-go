package reference_type

import (
	"fmt"
	"sync"
	"testing"
)

var (
	agentOnce sync.Once
	agentLock = &sync.Mutex{}
	agents    *Agents
)

// 数据
type AgentService struct {
	agents *Agents
}
type GrpcService struct {
	agents *Agents
}
type Agents struct {
	lastConnCount uint32
}

func NewAgents() *Agents {
	agentOnce.Do(func() {
		agents = &Agents{}
	})
	return agents
}

func GetAgents() *Agents {
	agentLock.Lock()
	defer agentLock.Unlock()
	if agents == nil {
		agents = NewAgents()
	}
	return agents
}

func TestName(t *testing.T) {
	as := &AgentService{}
	as2 := &AgentService{}
	as.agents = GetAgents()
	as2.agents = GetAgents()
	gs := &GrpcService{}
	gs.agents = GetAgents()
	fmt.Println(as.agents, gs.agents)
}
